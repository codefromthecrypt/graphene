package p2p

import (
	"bufio"
	"context"
	"errors"
	"sync"
	"time"

	streammux "github.com/libp2p/go-stream-muxer"

	"github.com/golang/protobuf/proto"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	inet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ps "github.com/libp2p/go-libp2p-peerstore"
	protocol "github.com/libp2p/go-libp2p-protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	multiaddr "github.com/multiformats/go-multiaddr"
	"github.com/phoreproject/synapse/pb"
	logger "github.com/sirupsen/logrus"
)

// MessageHandler is a function to handle messages.
type MessageHandler func(peer *Peer, message proto.Message) error

type messageHandlerAndID struct {
	handler MessageHandler
	id      uint64
}

// Message is a single message from a single peer.
type Message struct {
	From    *Peer
	Message proto.Message
}

// HostNode is the node for p2p host
// It's the low level P2P communication layer, the App class handles high level protocols
// The RPC communication is hanlded by App, not HostNode
type HostNode struct {
	publicKey  crypto.PubKey
	privateKey crypto.PrivKey

	host      host.Host
	gossipSub *pubsub.PubSub
	ctx       context.Context
	cancel    context.CancelFunc
	maxPeers  int

	timeoutInterval time.Duration

	// discovery handles peer discovery (mDNS, DHT, etc)
	discovery *Discovery

	// peerChan is a channel that handles incoming peers
	peerChan chan ps.PeerInfo

	// All peers that connected successfully with correct handshake
	peerList     []*Peer
	peerListLock *sync.Mutex

	// a messageHandler is called when a message with certain name is received
	messageHandlerMap map[string][]messageHandlerAndID
	handlerLock       *sync.RWMutex
	currentID         uint64
}

var protocolID = protocol.ID("/grpc/phore/0.0.1")

// NewHostNode creates a host node
func NewHostNode(listenAddress multiaddr.Multiaddr, publicKey crypto.PubKey, privateKey crypto.PrivKey, options DiscoveryOptions, timeoutInterval time.Duration, maxPeers int) (*HostNode, error) {
	ctx, cancel := context.WithCancel(context.Background())
	h, err := libp2p.New(
		ctx,
		libp2p.ListenAddrs(listenAddress),
		libp2p.Identity(privateKey),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	addrs, err := peerstore.InfoToP2pAddrs(&peerstore.PeerInfo{
		ID: h.ID(),
		Addrs: []multiaddr.Multiaddr{
			listenAddress,
		},
	})

	for _, a := range addrs {
		logger.WithField("addr", a).Info("binding to address")
	}

	// setup gossipsub protocol
	g, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		cancel()
		return nil, err
	}

	hostNode := &HostNode{
		publicKey:         publicKey,
		privateKey:        privateKey,
		host:              h,
		gossipSub:         g,
		ctx:               ctx,
		cancel:            cancel,
		messageHandlerMap: make(map[string][]messageHandlerAndID),
		handlerLock:       new(sync.RWMutex),
		currentID:         0,
		timeoutInterval:   timeoutInterval,
		peerListLock:      new(sync.Mutex),
		maxPeers:          maxPeers,
	}

	discovery := NewDiscovery(ctx, hostNode, options)
	hostNode.discovery = discovery

	// setup phore protocol
	h.SetStreamHandler(protocolID, hostNode.handleStream)

	return hostNode, nil
}

// handleStream handles an incoming stream.
func (node *HostNode) handleStream(stream inet.Stream) {
	_, err := node.setupPeerNode(stream, false)
	if err != nil {
		logger.Error("setup", err)
		return
	}
}

func (node *HostNode) handleMessage(peer *Peer, message proto.Message) error {
	logger.WithFields(logger.Fields{
		"peerID":  peer.ID.String(),
		"message": proto.MessageName(message),
	}).Debug("received message")

	err := peer.handleMessage(message)
	if err != nil {
		return err
	}

	node.handlerLock.RLock()
	handlerMap, found := node.messageHandlerMap[proto.MessageName(message)]
	node.handlerLock.RUnlock()
	if found {
		for _, handler := range handlerMap {
			err := handler.handler(peer, message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Handler is a message handler representation.
type Handler struct {
	ID          uint64
	MessageName string
}

// RegisterMessageHandler registers a message handler
func (node *HostNode) RegisterMessageHandler(messageName string, handler MessageHandler) Handler {
	node.handlerLock.Lock()
	defer node.handlerLock.Unlock()
	_, ok := node.messageHandlerMap[messageName]
	if !ok {
		node.messageHandlerMap[messageName] = make([]messageHandlerAndID, 0)
	}

	node.messageHandlerMap[messageName] = append(node.messageHandlerMap[messageName], messageHandlerAndID{handler, node.currentID})

	node.currentID++

	return Handler{node.currentID - 1, messageName}
}

// RemoveMessageHandler deregisters a message handler.
func (node *HostNode) RemoveMessageHandler(handler Handler) {
	node.handlerLock.Lock()
	defer node.handlerLock.Unlock()
	oldHandlerMap := node.messageHandlerMap[handler.MessageName]
	newHandlerMap := make([]messageHandlerAndID, 0, len(oldHandlerMap)-1)

	for i := range oldHandlerMap {
		if oldHandlerMap[i].id != handler.ID {
			newHandlerMap = append(newHandlerMap, oldHandlerMap[i])
		}
	}

	node.messageHandlerMap[handler.MessageName] = newHandlerMap
}

// Connect connects to a peer that we're not already connected to.
func (node *HostNode) Connect(peerInfo peerstore.PeerInfo) (*Peer, error) {
	if peerInfo.ID == node.GetHost().ID() {
		return nil, errors.New("cannot connect to self")
	}

	if node.IsPeerConnected(peerInfo) {
		return nil, nil
	}

	err := node.host.Connect(node.ctx, peerInfo)
	if err != nil {
		return nil, err
	}

	node.host.Peerstore().AddAddrs(peerInfo.ID, peerInfo.Addrs, ps.PermanentAddrTTL)

	stream, err := node.host.NewStream(context.Background(), peerInfo.ID, protocolID)

	if err != nil {
		logger.WithField("Function", "Connect").WithField("error", err).Warn("failed to open stream")
		return nil, err
	}

	return node.setupPeerNode(stream, true)
}

// IsPeerConnected checks if a peer is connected
func (node *HostNode) IsPeerConnected(peerInfo peerstore.PeerInfo) bool {
	for _, p := range node.peerList {
		if p.ID == peerInfo.ID {
			return true
		}
	}
	return false
}

// Run runs the main loop of the host node
func (node *HostNode) setupPeerNode(stream inet.Stream, outbound bool) (*Peer, error) {
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	peerNode := newPeer(rw, outbound, stream.Conn().RemotePeer(), node, node.timeoutInterval)

	node.peerListLock.Lock()
	node.peerList = append(node.peerList, peerNode)
	node.peerListLock.Unlock()

	logger.WithField("peer", peerNode.ID.Pretty()).WithField("outbound", peerNode.Outbound).Info("connected to peer")

	if outbound {
		peerIDBytes, err := node.host.ID().MarshalBinary()
		if err != nil {
			return nil, err
		}

		err = peerNode.SendMessage(&pb.VersionMessage{
			Version: ClientVersion,
			PeerID:  peerIDBytes,
		})
		if err != nil {
			return nil, err
		}
	}

	// TODO: switch handlers to be the responsibility of PeerNode so they can be cleaned up nicer

	handlers := make([]Handler, 0)

	handlers = append(handlers, node.RegisterMessageHandler("pb.VersionMessage", func(peer *Peer, message proto.Message) error {
		return peer.HandleVersionMessage(message.(*pb.VersionMessage))
	}))

	handlers = append(handlers, node.RegisterMessageHandler("pb.VerackMessage", func(peer *Peer, message proto.Message) error {
		return peer.handleVerackMessage(message.(*pb.VerackMessage))
	}))

	handlers = append(handlers, node.RegisterMessageHandler("pb.PingMessage", func(peer *Peer, message proto.Message) error {
		return peer.handlePingMessage(message.(*pb.PingMessage))
	}))

	handlers = append(handlers, node.RegisterMessageHandler("pb.PongMessage", func(peer *Peer, message proto.Message) error {
		return peer.handlePongMessage(message.(*pb.PongMessage))
	}))

	go func() {
		err := processMessages(rw.Reader, func(message proto.Message) error {
			return node.handleMessage(peerNode, message)
		})
		if err != nil {
			for _, h := range handlers {
				node.RemoveMessageHandler(h)
			}

			node.removePeer(peerNode)

			if err != streammux.ErrReset {
				logger.WithField("peer", peerNode.ID.Pretty()).Error(err)
			}
		}
	}()

	return peerNode, nil
}

// GetPublicKey returns the public key
func (node *HostNode) GetPublicKey() *crypto.PubKey {
	return &node.publicKey
}

// GetContext returns the context
func (node *HostNode) GetContext() context.Context {
	return node.ctx
}

// GetHost returns the host
func (node *HostNode) GetHost() host.Host {
	return node.host
}

// Broadcast broadcasts a message to the network for a topic.
func (node *HostNode) Broadcast(topic string, data []byte) error {
	return node.gossipSub.Publish(topic, data)
}

// SubscribeMessage registers a handler for a network topic.
func (node *HostNode) SubscribeMessage(topic string, handler func([]byte, peer.ID)) (*pubsub.Subscription, error) {
	subscription, err := node.gossipSub.Subscribe(topic)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			msg, err := subscription.Next(node.ctx)
			if err != nil {
				logger.WithField("error", err).Warn("error when getting next topic message")
				continue
			}

			handler(msg.Data, msg.GetFrom())
		}
	}()

	return subscription, nil
}

// UnsubscribeMessage cancels a subscription to a topic.
func (node *HostNode) UnsubscribeMessage(subscription *pubsub.Subscription) {
	subscription.Cancel()
}

func (node *HostNode) removePeer(peer *Peer) {
	node.peerListLock.Lock()
	defer node.peerListLock.Unlock()
	for i, p := range node.peerList {
		if p == peer {
			node.peerList = append(node.peerList[:i], node.peerList[i+1:]...)
			break
		}
	}
	node.host.Peerstore().ClearAddrs(peer.ID)
}

// DisconnectPeer disconnects a peer
func (node *HostNode) DisconnectPeer(peer *Peer) error {
	err := peer.disconnect()
	if err != nil {
		return err
	}
	node.removePeer(peer)
	return nil
}

// FindPeerByID finds a peer node by ID, returns nil if not found
func (node *HostNode) FindPeerByID(id peer.ID) (*Peer, bool) {
	node.peerListLock.Lock()
	defer node.peerListLock.Unlock()
	for _, p := range node.peerList {
		if p.ID == id {
			return p, true
		}
	}
	return nil, false
}

// PeerDiscovered is run when peers are discovered.
func (node *HostNode) PeerDiscovered(pi peerstore.PeerInfo) {
	if node.maxPeers <= len(node.peerList) {
		return
	}

	_, err := node.Connect(pi)
	if err != nil {
		logger.WithField("err", err).Debug("could not connect to peer")
	}
}

// Connected checks if the host node is connected.
func (node *HostNode) Connected() bool {
	node.peerListLock.Lock()
	defer node.peerListLock.Unlock()
	for _, p := range node.peerList {
		if p.Connecting == false {
			return true
		}
	}
	return false
}

// PeersConnected checks how many peers are connected.
func (node *HostNode) PeersConnected() int {
	node.peerListLock.Lock()
	defer node.peerListLock.Unlock()
	peersConnected := 0
	for _, p := range node.peerList {
		if p.Connecting == false {
			peersConnected++
		}
	}
	return peersConnected
}

// GetPeerList returns a list of all peers.
func (node *HostNode) GetPeerList() []*Peer {
	node.peerListLock.Lock()
	defer node.peerListLock.Unlock()
	return node.peerList
}

// GetPeerByID gets a peer by ID or returns nil if we aren't
// connected.
func (node *HostNode) GetPeerByID(id peer.ID) *Peer {
	node.peerListLock.Lock()
	defer node.peerListLock.Unlock()
	for _, p := range node.peerList {
		if id == p.ID {
			return p
		}
	}
	return nil
}

// StartDiscovery starts the host node discovering peers.
func (node *HostNode) StartDiscovery() error {
	return node.discovery.StartDiscovery()
}
