package p2p

import (
	"bufio"
	"context"
	"net"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	inet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	ps "github.com/libp2p/go-libp2p-peerstore"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	multiaddr "github.com/multiformats/go-multiaddr"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// HostNode is the node for host
type HostNode struct {
	publicKey  crypto.PubKey
	privateKey crypto.PrivKey
	host       host.Host
	gossipSub  *pubsub.PubSub
	ctx        context.Context
	cancel     context.CancelFunc
	grpcServer *grpc.Server
	streamCh   chan inet.Stream
	readWriter *bufio.ReadWriter
}

// NewHostNode creates a host node
func NewHostNode(listenAddress multiaddr.Multiaddr, publicKey crypto.PubKey, privateKey crypto.PrivKey) (*HostNode, error) {
	ctx, cancel := context.WithCancel(context.Background())
	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrs(listenAddress),
		libp2p.Identity(privateKey),
	)
	if err != nil {
		logger.WithField("Function", "NewHostNode").Warn(err)
		cancel()
		return nil, err
	}

	for _, a := range host.Addrs() {
		logger.WithField("address", a.String()).Debug("binding to port")
	}

	if err != nil {
		cancel()
		return nil, err
	}

	g, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		cancel()
		return nil, err
	}

	grpcServer := grpc.NewServer()
	stream := make(chan inet.Stream)
	hostNode := HostNode{
		publicKey:  publicKey,
		privateKey: privateKey,
		host:       host,
		gossipSub:  g,
		ctx:        ctx,
		cancel:     cancel,
		grpcServer: grpcServer,
		streamCh:   stream,
		//readWriter: bufio.NewReadWriter(bufio.NewReader(inet.Stream(stream)), bufio.NewWriter(inet.Stream(stream))),
	}

	host.SetStreamHandler("/grpc/0.0.1", hostNode.HandleStream)

	return &hostNode, nil
}

// Connect connects to a peer
func (node *HostNode) Connect(peerInfo *peerstore.PeerInfo) (*PeerNode, error) {
	for _, p := range node.GetHost().Peerstore().PeersWithAddrs() {
		if p == peerInfo.ID {
			return nil, nil
		}
	}

	logger.WithField("addrs", peerInfo.Addrs).WithField("id", peerInfo.ID).Debug("attempting to connect to a peer")

	ctx, cancel := context.WithTimeout(node.ctx, 10*time.Second)
	defer cancel()

	node.host.Peerstore().AddAddrs(peerInfo.ID, peerInfo.Addrs, ps.PermanentAddrTTL)

	// grpcConn
	_, err := node.Dial(ctx, peerInfo.ID, grpc.WithInsecure(), grpc.WithBlock())
	//err := node.host.Connect(ctx, *peerInfo)

	if err != nil {
		logger.WithField("error", err).Warn("failed to connect to peer")
		return nil, err
	}

	stream, err := node.host.NewStream(context.Background(), peerInfo.ID, "")

	if err != nil {
		logger.WithField("error", err).Warn("failed to open stream")
		return nil, err
	}

	readWriter := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	peerNode := PeerNode{
		stream:     stream,
		readWriter: readWriter,
	}

	return &peerNode, err
}

// Run runs the main loop of the host node
func (node *HostNode) Run() {
	go node.grpcServer.Serve(newGrpcListener(node))
}

// GetGRPCServer returns the grpc server.
func (node *HostNode) GetGRPCServer() *grpc.Server {
	return node.grpcServer
}

// HandleStream handles an incoming stream.
func (node *HostNode) HandleStream(stream inet.Stream) {
	select {
	case <-node.ctx.Done():
		return
	case node.streamCh <- stream:
	}
}

// GetDialOption returns the WithDialer option to dial via libp2p.
// note: ctx should be the root context.
func (node *HostNode) GetDialOption(ctx context.Context) grpc.DialOption {
	return grpc.WithDialer(func(peerIdStr string, timeout time.Duration) (net.Conn, error) {
		subCtx, subCtxCancel := context.WithTimeout(ctx, timeout)
		defer subCtxCancel()

		id, err := peer.IDB58Decode(peerIdStr)
		if err != nil {
			logger.WithField("Function", "GetDialOption").WithField("PeerID", peerIdStr).Warn(err)
			return nil, err
		}

		err = node.host.Connect(subCtx, ps.PeerInfo{
			ID: id,
		})
		if err != nil {
			logger.WithField("Function", "GetDialOption").Warn(err)
			return nil, err
		}

		stream, err := node.host.NewStream(ctx, id, "/grpc/0.0.1")
		if err != nil {
			logger.WithField("Function", "GetDialOption").Warn(err)
			return nil, err
		}

		return &streamConn{Stream: stream}, nil
	})
}

// Dial attempts to open a GRPC connection over libp2p to a peer.
// Note that the context is used as the **stream context** not just the dial context.
func (node *HostNode) Dial(ctx context.Context, peerID peer.ID, dialOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	dialOpsPrepended := append([]grpc.DialOption{node.GetDialOption(ctx)}, dialOpts...)
	return grpc.DialContext(ctx, peerID.Pretty(), dialOpsPrepended...)
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

// Discover discovers the peers
func (node *HostNode) Discover(options *DiscoveryOptions) error {
	return startDiscovery(node, options)
}

// GetConnectedPeerCount returns the connected peer count
func (node *HostNode) GetConnectedPeerCount() int {
	return node.host.Peerstore().Peers().Len()
}

// Broadcast broadcasts a message to the network for a topic.
func (node *HostNode) Broadcast(topic string, data []byte) error {
	return node.gossipSub.Publish(topic, data)
}

// SubscribeMessage registers a handler for a network topic.
func (node *HostNode) SubscribeMessage(topic string, handler func([]byte) error) (*pubsub.Subscription, error) {
	subscription, err := node.gossipSub.Subscribe(topic)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			msg, err := subscription.Next(node.ctx)
			if err != nil {
				logger.WithField("error", err).Warn("error when getting next topic message")
				return
			}

			err = handler(msg.Data)
			if err != nil {
				logger.WithField("topic", topic).WithField("error", err).Warn("error when handling message")
			}
		}
	}()

	return subscription, nil
}

// UnsubscribeMessage cancels a subscription to a topic.
func (node *HostNode) UnsubscribeMessage(subscription *pubsub.Subscription) {
	subscription.Cancel()
}
