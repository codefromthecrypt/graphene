// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BeaconVersionMessage struct {
	Version              uint64   `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	PeerID               []byte   `protobuf:"bytes,2,opt,name=PeerID,proto3" json:"PeerID,omitempty"`
	PeerInfo             []byte   `protobuf:"bytes,3,opt,name=PeerInfo,proto3" json:"PeerInfo,omitempty"`
	GenesisHash          []byte   `protobuf:"bytes,4,opt,name=GenesisHash,proto3" json:"GenesisHash,omitempty"`
	Height               uint64   `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconVersionMessage) Reset()         { *m = BeaconVersionMessage{} }
func (m *BeaconVersionMessage) String() string { return proto.CompactTextString(m) }
func (*BeaconVersionMessage) ProtoMessage()    {}
func (*BeaconVersionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{0}
}
func (m *BeaconVersionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconVersionMessage.Unmarshal(m, b)
}
func (m *BeaconVersionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconVersionMessage.Marshal(b, m, deterministic)
}
func (dst *BeaconVersionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconVersionMessage.Merge(dst, src)
}
func (m *BeaconVersionMessage) XXX_Size() int {
	return xxx_messageInfo_BeaconVersionMessage.Size(m)
}
func (m *BeaconVersionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconVersionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconVersionMessage proto.InternalMessageInfo

func (m *BeaconVersionMessage) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BeaconVersionMessage) GetPeerID() []byte {
	if m != nil {
		return m.PeerID
	}
	return nil
}

func (m *BeaconVersionMessage) GetPeerInfo() []byte {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

func (m *BeaconVersionMessage) GetGenesisHash() []byte {
	if m != nil {
		return m.GenesisHash
	}
	return nil
}

func (m *BeaconVersionMessage) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ShardVersionMessage struct {
	Version              uint64   `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	PeerID               []byte   `protobuf:"bytes,3,opt,name=PeerID,proto3" json:"PeerID,omitempty"`
	PeerInfo             []byte   `protobuf:"bytes,4,opt,name=PeerInfo,proto3" json:"PeerInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardVersionMessage) Reset()         { *m = ShardVersionMessage{} }
func (m *ShardVersionMessage) String() string { return proto.CompactTextString(m) }
func (*ShardVersionMessage) ProtoMessage()    {}
func (*ShardVersionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{1}
}
func (m *ShardVersionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardVersionMessage.Unmarshal(m, b)
}
func (m *ShardVersionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardVersionMessage.Marshal(b, m, deterministic)
}
func (dst *ShardVersionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardVersionMessage.Merge(dst, src)
}
func (m *ShardVersionMessage) XXX_Size() int {
	return xxx_messageInfo_ShardVersionMessage.Size(m)
}
func (m *ShardVersionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardVersionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ShardVersionMessage proto.InternalMessageInfo

func (m *ShardVersionMessage) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ShardVersionMessage) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ShardVersionMessage) GetPeerID() []byte {
	if m != nil {
		return m.PeerID
	}
	return nil
}

func (m *ShardVersionMessage) GetPeerInfo() []byte {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

type RelayerVersionMessage struct {
	Version              uint64   `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Provider             bool     `protobuf:"varint,2,opt,name=Provider,proto3" json:"Provider,omitempty"`
	PeerID               []byte   `protobuf:"bytes,3,opt,name=PeerID,proto3" json:"PeerID,omitempty"`
	PeerInfo             []byte   `protobuf:"bytes,4,opt,name=PeerInfo,proto3" json:"PeerInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelayerVersionMessage) Reset()         { *m = RelayerVersionMessage{} }
func (m *RelayerVersionMessage) String() string { return proto.CompactTextString(m) }
func (*RelayerVersionMessage) ProtoMessage()    {}
func (*RelayerVersionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{2}
}
func (m *RelayerVersionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelayerVersionMessage.Unmarshal(m, b)
}
func (m *RelayerVersionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelayerVersionMessage.Marshal(b, m, deterministic)
}
func (dst *RelayerVersionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayerVersionMessage.Merge(dst, src)
}
func (m *RelayerVersionMessage) XXX_Size() int {
	return xxx_messageInfo_RelayerVersionMessage.Size(m)
}
func (m *RelayerVersionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayerVersionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RelayerVersionMessage proto.InternalMessageInfo

func (m *RelayerVersionMessage) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RelayerVersionMessage) GetProvider() bool {
	if m != nil {
		return m.Provider
	}
	return false
}

func (m *RelayerVersionMessage) GetPeerID() []byte {
	if m != nil {
		return m.PeerID
	}
	return nil
}

func (m *RelayerVersionMessage) GetPeerInfo() []byte {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

type PingMessage struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{3}
}
func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (dst *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(dst, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type PongMessage struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongMessage) Reset()         { *m = PongMessage{} }
func (m *PongMessage) String() string { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()    {}
func (*PongMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{4}
}
func (m *PongMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongMessage.Unmarshal(m, b)
}
func (m *PongMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongMessage.Marshal(b, m, deterministic)
}
func (dst *PongMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongMessage.Merge(dst, src)
}
func (m *PongMessage) XXX_Size() int {
	return xxx_messageInfo_PongMessage.Size(m)
}
func (m *PongMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PongMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PongMessage proto.InternalMessageInfo

func (m *PongMessage) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type RejectMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RejectMessage) Reset()         { *m = RejectMessage{} }
func (m *RejectMessage) String() string { return proto.CompactTextString(m) }
func (*RejectMessage) ProtoMessage()    {}
func (*RejectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{5}
}
func (m *RejectMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RejectMessage.Unmarshal(m, b)
}
func (m *RejectMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RejectMessage.Marshal(b, m, deterministic)
}
func (dst *RejectMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RejectMessage.Merge(dst, src)
}
func (m *RejectMessage) XXX_Size() int {
	return xxx_messageInfo_RejectMessage.Size(m)
}
func (m *RejectMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RejectMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RejectMessage proto.InternalMessageInfo

func (m *RejectMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AttestationMempoolItem struct {
	AttestationHash      []byte   `protobuf:"bytes,1,opt,name=AttestationHash,proto3" json:"AttestationHash,omitempty"`
	Participation        []byte   `protobuf:"bytes,2,opt,name=Participation,proto3" json:"Participation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestationMempoolItem) Reset()         { *m = AttestationMempoolItem{} }
func (m *AttestationMempoolItem) String() string { return proto.CompactTextString(m) }
func (*AttestationMempoolItem) ProtoMessage()    {}
func (*AttestationMempoolItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{6}
}
func (m *AttestationMempoolItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationMempoolItem.Unmarshal(m, b)
}
func (m *AttestationMempoolItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationMempoolItem.Marshal(b, m, deterministic)
}
func (dst *AttestationMempoolItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationMempoolItem.Merge(dst, src)
}
func (m *AttestationMempoolItem) XXX_Size() int {
	return xxx_messageInfo_AttestationMempoolItem.Size(m)
}
func (m *AttestationMempoolItem) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationMempoolItem.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationMempoolItem proto.InternalMessageInfo

func (m *AttestationMempoolItem) GetAttestationHash() []byte {
	if m != nil {
		return m.AttestationHash
	}
	return nil
}

func (m *AttestationMempoolItem) GetParticipation() []byte {
	if m != nil {
		return m.Participation
	}
	return nil
}

// This will advertise which mempool items the client needs/has.
type GetAttestationMempoolMessage struct {
	Attestations         []*AttestationMempoolItem `protobuf:"bytes,1,rep,name=Attestations,proto3" json:"Attestations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetAttestationMempoolMessage) Reset()         { *m = GetAttestationMempoolMessage{} }
func (m *GetAttestationMempoolMessage) String() string { return proto.CompactTextString(m) }
func (*GetAttestationMempoolMessage) ProtoMessage()    {}
func (*GetAttestationMempoolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{7}
}
func (m *GetAttestationMempoolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAttestationMempoolMessage.Unmarshal(m, b)
}
func (m *GetAttestationMempoolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAttestationMempoolMessage.Marshal(b, m, deterministic)
}
func (dst *GetAttestationMempoolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAttestationMempoolMessage.Merge(dst, src)
}
func (m *GetAttestationMempoolMessage) XXX_Size() int {
	return xxx_messageInfo_GetAttestationMempoolMessage.Size(m)
}
func (m *GetAttestationMempoolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAttestationMempoolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetAttestationMempoolMessage proto.InternalMessageInfo

func (m *GetAttestationMempoolMessage) GetAttestations() []*AttestationMempoolItem {
	if m != nil {
		return m.Attestations
	}
	return nil
}

type AttestationMempoolMessage struct {
	Attestations         []*Attestation `protobuf:"bytes,1,rep,name=Attestations,proto3" json:"Attestations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AttestationMempoolMessage) Reset()         { *m = AttestationMempoolMessage{} }
func (m *AttestationMempoolMessage) String() string { return proto.CompactTextString(m) }
func (*AttestationMempoolMessage) ProtoMessage()    {}
func (*AttestationMempoolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{8}
}
func (m *AttestationMempoolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationMempoolMessage.Unmarshal(m, b)
}
func (m *AttestationMempoolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationMempoolMessage.Marshal(b, m, deterministic)
}
func (dst *AttestationMempoolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationMempoolMessage.Merge(dst, src)
}
func (m *AttestationMempoolMessage) XXX_Size() int {
	return xxx_messageInfo_AttestationMempoolMessage.Size(m)
}
func (m *AttestationMempoolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationMempoolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationMempoolMessage proto.InternalMessageInfo

func (m *AttestationMempoolMessage) GetAttestations() []*Attestation {
	if m != nil {
		return m.Attestations
	}
	return nil
}

type GetBlocksMessage struct {
	LocatorHashes        [][]byte `protobuf:"bytes,1,rep,name=LocatorHashes,proto3" json:"LocatorHashes,omitempty"`
	HashStop             []byte   `protobuf:"bytes,2,opt,name=HashStop,proto3" json:"HashStop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlocksMessage) Reset()         { *m = GetBlocksMessage{} }
func (m *GetBlocksMessage) String() string { return proto.CompactTextString(m) }
func (*GetBlocksMessage) ProtoMessage()    {}
func (*GetBlocksMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{9}
}
func (m *GetBlocksMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlocksMessage.Unmarshal(m, b)
}
func (m *GetBlocksMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlocksMessage.Marshal(b, m, deterministic)
}
func (dst *GetBlocksMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlocksMessage.Merge(dst, src)
}
func (m *GetBlocksMessage) XXX_Size() int {
	return xxx_messageInfo_GetBlocksMessage.Size(m)
}
func (m *GetBlocksMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlocksMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlocksMessage proto.InternalMessageInfo

func (m *GetBlocksMessage) GetLocatorHashes() [][]byte {
	if m != nil {
		return m.LocatorHashes
	}
	return nil
}

func (m *GetBlocksMessage) GetHashStop() []byte {
	if m != nil {
		return m.HashStop
	}
	return nil
}

// Response to GetBlockMessage
type BeaconBlockMessage struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
	LatestBlockHash      []byte   `protobuf:"bytes,2,opt,name=LatestBlockHash,proto3" json:"LatestBlockHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconBlockMessage) Reset()         { *m = BeaconBlockMessage{} }
func (m *BeaconBlockMessage) String() string { return proto.CompactTextString(m) }
func (*BeaconBlockMessage) ProtoMessage()    {}
func (*BeaconBlockMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{10}
}
func (m *BeaconBlockMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconBlockMessage.Unmarshal(m, b)
}
func (m *BeaconBlockMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconBlockMessage.Marshal(b, m, deterministic)
}
func (dst *BeaconBlockMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconBlockMessage.Merge(dst, src)
}
func (m *BeaconBlockMessage) XXX_Size() int {
	return xxx_messageInfo_BeaconBlockMessage.Size(m)
}
func (m *BeaconBlockMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconBlockMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconBlockMessage proto.InternalMessageInfo

func (m *BeaconBlockMessage) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *BeaconBlockMessage) GetLatestBlockHash() []byte {
	if m != nil {
		return m.LatestBlockHash
	}
	return nil
}

type GetShardBlocksMessage struct {
	StartHash            []byte   `protobuf:"bytes,1,opt,name=StartHash,proto3" json:"StartHash,omitempty"`
	LocatorHashes        [][]byte `protobuf:"bytes,2,rep,name=LocatorHashes,proto3" json:"LocatorHashes,omitempty"`
	HashStop             []byte   `protobuf:"bytes,3,opt,name=HashStop,proto3" json:"HashStop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShardBlocksMessage) Reset()         { *m = GetShardBlocksMessage{} }
func (m *GetShardBlocksMessage) String() string { return proto.CompactTextString(m) }
func (*GetShardBlocksMessage) ProtoMessage()    {}
func (*GetShardBlocksMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{11}
}
func (m *GetShardBlocksMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShardBlocksMessage.Unmarshal(m, b)
}
func (m *GetShardBlocksMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShardBlocksMessage.Marshal(b, m, deterministic)
}
func (dst *GetShardBlocksMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShardBlocksMessage.Merge(dst, src)
}
func (m *GetShardBlocksMessage) XXX_Size() int {
	return xxx_messageInfo_GetShardBlocksMessage.Size(m)
}
func (m *GetShardBlocksMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShardBlocksMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetShardBlocksMessage proto.InternalMessageInfo

func (m *GetShardBlocksMessage) GetStartHash() []byte {
	if m != nil {
		return m.StartHash
	}
	return nil
}

func (m *GetShardBlocksMessage) GetLocatorHashes() [][]byte {
	if m != nil {
		return m.LocatorHashes
	}
	return nil
}

func (m *GetShardBlocksMessage) GetHashStop() []byte {
	if m != nil {
		return m.HashStop
	}
	return nil
}

type ShardBlockMessage struct {
	Blocks               []*ShardBlock `protobuf:"bytes,1,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
	LatestBlockHash      []byte        `protobuf:"bytes,2,opt,name=LatestBlockHash,proto3" json:"LatestBlockHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ShardBlockMessage) Reset()         { *m = ShardBlockMessage{} }
func (m *ShardBlockMessage) String() string { return proto.CompactTextString(m) }
func (*ShardBlockMessage) ProtoMessage()    {}
func (*ShardBlockMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{12}
}
func (m *ShardBlockMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardBlockMessage.Unmarshal(m, b)
}
func (m *ShardBlockMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardBlockMessage.Marshal(b, m, deterministic)
}
func (dst *ShardBlockMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardBlockMessage.Merge(dst, src)
}
func (m *ShardBlockMessage) XXX_Size() int {
	return xxx_messageInfo_ShardBlockMessage.Size(m)
}
func (m *ShardBlockMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardBlockMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ShardBlockMessage proto.InternalMessageInfo

func (m *ShardBlockMessage) GetBlocks() []*ShardBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *ShardBlockMessage) GetLatestBlockHash() []byte {
	if m != nil {
		return m.LatestBlockHash
	}
	return nil
}

type GetPackagesMessage struct {
	TipStateRoot         []byte   `protobuf:"bytes,1,opt,name=TipStateRoot,proto3" json:"TipStateRoot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPackagesMessage) Reset()         { *m = GetPackagesMessage{} }
func (m *GetPackagesMessage) String() string { return proto.CompactTextString(m) }
func (*GetPackagesMessage) ProtoMessage()    {}
func (*GetPackagesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{13}
}
func (m *GetPackagesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPackagesMessage.Unmarshal(m, b)
}
func (m *GetPackagesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPackagesMessage.Marshal(b, m, deterministic)
}
func (dst *GetPackagesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPackagesMessage.Merge(dst, src)
}
func (m *GetPackagesMessage) XXX_Size() int {
	return xxx_messageInfo_GetPackagesMessage.Size(m)
}
func (m *GetPackagesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPackagesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetPackagesMessage proto.InternalMessageInfo

func (m *GetPackagesMessage) GetTipStateRoot() []byte {
	if m != nil {
		return m.TipStateRoot
	}
	return nil
}

type PackageMessage struct {
	Package              *TransactionPackage `protobuf:"bytes,1,opt,name=Package,proto3" json:"Package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PackageMessage) Reset()         { *m = PackageMessage{} }
func (m *PackageMessage) String() string { return proto.CompactTextString(m) }
func (*PackageMessage) ProtoMessage()    {}
func (*PackageMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{14}
}
func (m *PackageMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageMessage.Unmarshal(m, b)
}
func (m *PackageMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageMessage.Marshal(b, m, deterministic)
}
func (dst *PackageMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageMessage.Merge(dst, src)
}
func (m *PackageMessage) XXX_Size() int {
	return xxx_messageInfo_PackageMessage.Size(m)
}
func (m *PackageMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PackageMessage proto.InternalMessageInfo

func (m *PackageMessage) GetPackage() *TransactionPackage {
	if m != nil {
		return m.Package
	}
	return nil
}

type SubmitTransaction struct {
	Transaction          *ShardTransaction `protobuf:"bytes,1,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SubmitTransaction) Reset()         { *m = SubmitTransaction{} }
func (m *SubmitTransaction) String() string { return proto.CompactTextString(m) }
func (*SubmitTransaction) ProtoMessage()    {}
func (*SubmitTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e115988d37105cfa, []int{15}
}
func (m *SubmitTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransaction.Unmarshal(m, b)
}
func (m *SubmitTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransaction.Marshal(b, m, deterministic)
}
func (dst *SubmitTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransaction.Merge(dst, src)
}
func (m *SubmitTransaction) XXX_Size() int {
	return xxx_messageInfo_SubmitTransaction.Size(m)
}
func (m *SubmitTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransaction proto.InternalMessageInfo

func (m *SubmitTransaction) GetTransaction() *ShardTransaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func init() {
	proto.RegisterType((*BeaconVersionMessage)(nil), "pb.BeaconVersionMessage")
	proto.RegisterType((*ShardVersionMessage)(nil), "pb.ShardVersionMessage")
	proto.RegisterType((*RelayerVersionMessage)(nil), "pb.RelayerVersionMessage")
	proto.RegisterType((*PingMessage)(nil), "pb.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "pb.PongMessage")
	proto.RegisterType((*RejectMessage)(nil), "pb.RejectMessage")
	proto.RegisterType((*AttestationMempoolItem)(nil), "pb.AttestationMempoolItem")
	proto.RegisterType((*GetAttestationMempoolMessage)(nil), "pb.GetAttestationMempoolMessage")
	proto.RegisterType((*AttestationMempoolMessage)(nil), "pb.AttestationMempoolMessage")
	proto.RegisterType((*GetBlocksMessage)(nil), "pb.GetBlocksMessage")
	proto.RegisterType((*BeaconBlockMessage)(nil), "pb.BeaconBlockMessage")
	proto.RegisterType((*GetShardBlocksMessage)(nil), "pb.GetShardBlocksMessage")
	proto.RegisterType((*ShardBlockMessage)(nil), "pb.ShardBlockMessage")
	proto.RegisterType((*GetPackagesMessage)(nil), "pb.GetPackagesMessage")
	proto.RegisterType((*PackageMessage)(nil), "pb.PackageMessage")
	proto.RegisterType((*SubmitTransaction)(nil), "pb.SubmitTransaction")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor_p2p_e115988d37105cfa) }

var fileDescriptor_p2p_e115988d37105cfa = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x6f, 0x12, 0x41,
	0x10, 0xce, 0x41, 0x4b, 0xcb, 0x00, 0xad, 0x3d, 0x29, 0x39, 0x49, 0x1f, 0x70, 0x6d, 0x0c, 0xbe,
	0x10, 0x43, 0x13, 0xe3, 0x93, 0x89, 0xc4, 0x84, 0x36, 0x56, 0x43, 0x16, 0xe2, 0xa3, 0xc9, 0x72,
	0x1d, 0xe1, 0x2c, 0xdc, 0x5e, 0xf6, 0x46, 0x8d, 0xf1, 0xc9, 0x3f, 0xe2, 0x6f, 0x35, 0xbb, 0xb7,
	0x07, 0x7b, 0xd7, 0x36, 0xd5, 0xbe, 0xed, 0xf7, 0xcd, 0xec, 0x37, 0xdf, 0xc7, 0xdc, 0x02, 0xf5,
	0x64, 0x98, 0x0c, 0x12, 0x25, 0x49, 0xfa, 0x95, 0x64, 0xde, 0x6d, 0x86, 0x72, 0xbd, 0x96, 0x71,
	0xc6, 0xb0, 0x3f, 0x1e, 0xb4, 0x47, 0x28, 0x42, 0x19, 0x7f, 0x42, 0x95, 0x46, 0x32, 0xfe, 0x80,
	0x69, 0x2a, 0x16, 0xe8, 0x07, 0xb0, 0x67, 0x99, 0xc0, 0xeb, 0x79, 0xfd, 0x1d, 0x9e, 0x43, 0xbf,
	0x03, 0xb5, 0x09, 0xa2, 0xba, 0x78, 0x17, 0x54, 0x7a, 0x5e, 0xbf, 0xc9, 0x2d, 0xf2, 0xbb, 0xb0,
	0x6f, 0x4e, 0xf1, 0x17, 0x19, 0x54, 0x4d, 0x65, 0x83, 0xfd, 0x1e, 0x34, 0xc6, 0x18, 0x63, 0x1a,
	0xa5, 0xe7, 0x22, 0x5d, 0x06, 0x3b, 0xa6, 0xec, 0x52, 0x5a, 0xf5, 0x1c, 0xa3, 0xc5, 0x92, 0x82,
	0x5d, 0x33, 0xce, 0x22, 0xf6, 0x0b, 0x1e, 0x4f, 0x97, 0x42, 0x5d, 0xfd, 0x8f, 0x3d, 0x2b, 0x54,
	0x71, 0x85, 0x1c, 0xdb, 0xd5, 0x3b, 0x6d, 0xef, 0x14, 0x6d, 0xb3, 0xdf, 0x1e, 0x1c, 0x73, 0x5c,
	0x89, 0x9f, 0xa8, 0xfe, 0x79, 0xbe, 0xd6, 0x53, 0xf2, 0x7b, 0x74, 0x85, 0xca, 0x38, 0xd8, 0xe7,
	0x1b, 0xfc, 0x20, 0x0f, 0xcf, 0xa0, 0x31, 0x89, 0xe2, 0x45, 0x3e, 0xb8, 0x0d, 0xbb, 0x1f, 0x65,
	0x1c, 0xa2, 0x1d, 0x9b, 0x01, 0xd3, 0x24, 0xef, 0x6b, 0x7a, 0x01, 0x2d, 0x8e, 0x5f, 0x31, 0x24,
	0x27, 0x84, 0x3d, 0x9a, 0xc6, 0x3a, 0xcf, 0x21, 0x5b, 0x42, 0xe7, 0x2d, 0x11, 0xa6, 0x24, 0xc8,
	0x84, 0x5e, 0x27, 0x52, 0xae, 0x2e, 0x08, 0xd7, 0x7e, 0x1f, 0x0e, 0x9d, 0x8a, 0xd9, 0xa6, 0x67,
	0x1c, 0x97, 0x69, 0xff, 0x14, 0x5a, 0x13, 0xa1, 0x28, 0x0a, 0xa3, 0xc4, 0x90, 0xf6, 0x73, 0x29,
	0x92, 0xec, 0x33, 0x9c, 0x8c, 0x91, 0x6e, 0x0e, 0xcb, 0x3d, 0xbe, 0x81, 0xa6, 0x53, 0x4c, 0x03,
	0xaf, 0x57, 0xed, 0x37, 0x86, 0xdd, 0x41, 0x32, 0x1f, 0xdc, 0xee, 0x90, 0x17, 0xfa, 0xd9, 0x04,
	0x9e, 0xdc, 0x2d, 0x7e, 0x76, 0xab, 0xf8, 0x61, 0x49, 0xbc, 0xa4, 0x38, 0x83, 0x47, 0x63, 0xa4,
	0xd1, 0x4a, 0x86, 0xd7, 0x69, 0x2e, 0x74, 0x0a, 0xad, 0x4b, 0x19, 0x0a, 0x92, 0x4a, 0x47, 0xc7,
	0x4c, 0xa9, 0xc9, 0x8b, 0xa4, 0x5e, 0xb3, 0x3e, 0x4d, 0x49, 0x26, 0xf6, 0xc7, 0xd8, 0x60, 0x26,
	0xc0, 0xcf, 0xde, 0xa1, 0x11, 0xce, 0x75, 0x9f, 0x42, 0x2d, 0x1b, 0x64, 0xad, 0xd5, 0xb5, 0x35,
	0xc3, 0x70, 0x5b, 0xd0, 0x0b, 0xb9, 0x14, 0xda, 0x9e, 0xc1, 0x66, 0x21, 0x99, 0x76, 0x99, 0x66,
	0x3f, 0xe0, 0x78, 0x8c, 0x64, 0x5e, 0x53, 0xd1, 0xfd, 0x09, 0xd4, 0xa7, 0x24, 0x14, 0x39, 0xdb,
	0xdc, 0x12, 0x37, 0xb3, 0x55, 0xee, 0xcb, 0x56, 0x2d, 0x65, 0x43, 0x38, 0xda, 0x4e, 0xcd, 0x87,
	0x3e, 0x2f, 0x45, 0x3b, 0xd0, 0xd1, 0xb6, 0x6d, 0x0f, 0xc8, 0xf7, 0x1a, 0xfc, 0x31, 0xd2, 0x44,
	0x84, 0xd7, 0x62, 0x81, 0x9b, 0x70, 0x0c, 0x9a, 0xb3, 0x28, 0x99, 0x92, 0x20, 0xe4, 0x52, 0x92,
	0xcd, 0x57, 0xe0, 0xd8, 0x08, 0x0e, 0xec, 0xb5, 0xfc, 0xd6, 0x4b, 0xd8, 0xb3, 0x8c, 0xb9, 0xd0,
	0x18, 0x76, 0xb4, 0xbd, 0x99, 0x12, 0x71, 0x2a, 0x42, 0xfd, 0x1d, 0xd8, 0x2a, 0xcf, 0xdb, 0xd8,
	0x7b, 0x38, 0x9a, 0x7e, 0x9b, 0xaf, 0x23, 0x72, 0x9a, 0xfc, 0x57, 0xd0, 0x70, 0xa0, 0x95, 0x6a,
	0x6f, 0x92, 0x3a, 0x35, 0xee, 0x36, 0xce, 0x6b, 0xe6, 0xdf, 0xf9, 0xec, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x26, 0x6f, 0x27, 0xf3, 0xbc, 0x05, 0x00, 0x00,
}
