// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proof.proto

package pbtypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AccumulatorProof struct {
	// The bitmap indicating which siblings are default. 1 means non-default and
	// 0 means default. The LSB corresponds to the sibling at the bottom of the
	// accumulator. The leftmost 1-bit corresponds to the sibling at the level
	// just below root level in the accumulator, since this one is always
	// non-default.
	Bitmap uint64 `protobuf:"varint,1,opt,name=bitmap,proto3" json:"bitmap,omitempty"`
	// The non-default siblings. The ones near the root are at the beginning of
	// the list.
	NonDefaultSiblings   [][]byte `protobuf:"bytes,2,rep,name=non_default_siblings,json=nonDefaultSiblings,proto3" json:"non_default_siblings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccumulatorProof) Reset()         { *m = AccumulatorProof{} }
func (m *AccumulatorProof) String() string { return proto.CompactTextString(m) }
func (*AccumulatorProof) ProtoMessage()    {}
func (*AccumulatorProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_473d204b28f447f0, []int{0}
}

func (m *AccumulatorProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccumulatorProof.Unmarshal(m, b)
}
func (m *AccumulatorProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccumulatorProof.Marshal(b, m, deterministic)
}
func (m *AccumulatorProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccumulatorProof.Merge(m, src)
}
func (m *AccumulatorProof) XXX_Size() int {
	return xxx_messageInfo_AccumulatorProof.Size(m)
}
func (m *AccumulatorProof) XXX_DiscardUnknown() {
	xxx_messageInfo_AccumulatorProof.DiscardUnknown(m)
}

var xxx_messageInfo_AccumulatorProof proto.InternalMessageInfo

func (m *AccumulatorProof) GetBitmap() uint64 {
	if m != nil {
		return m.Bitmap
	}
	return 0
}

func (m *AccumulatorProof) GetNonDefaultSiblings() [][]byte {
	if m != nil {
		return m.NonDefaultSiblings
	}
	return nil
}

type SparseMerkleProof struct {
	// This proof can be used to authenticate whether a given leaf exists in the
	// tree or not. In Rust:
	//   - If this is `Some(HashValue, HashValue)`
	//     - If the first `HashValue` equals requested key, this is an inclusion
	//       proof and the second `HashValue` equals the hash of the
	//       corresponding account blob.
	//     - Otherwise this is a non-inclusion proof. The first `HashValue` is
	//       the only key that exists in the subtree and the second `HashValue`
	//       equals the hash of the corresponding account blob.
	//   - If this is `None`, this is also a non-inclusion proof which indicates
	//     the subtree is empty.
	//
	// In protobuf, this leaf field should either be
	//   - empty, which corresponds to None in the Rust structure.
	//   - exactly 64 bytes, which corresponds to Some<(HashValue, HashValue)>
	//     in the Rust structure.
	Leaf []byte `protobuf:"bytes,1,opt,name=leaf,proto3" json:"leaf,omitempty"`
	// The bitmap indicating which siblings are default. 1 means non-default and
	// 0 means default. The MSB of the first byte corresponds to the sibling at
	// the top of the Sparse Merkle Tree. The rightmost 1-bit of the last byte
	// corresponds to the sibling at the bottom, since this one is always
	// non-default.
	Bitmap []byte `protobuf:"bytes,2,opt,name=bitmap,proto3" json:"bitmap,omitempty"`
	// The non-default siblings. The ones near the root are at the beginning of
	// the list.
	NonDefaultSiblings   [][]byte `protobuf:"bytes,3,rep,name=non_default_siblings,json=nonDefaultSiblings,proto3" json:"non_default_siblings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SparseMerkleProof) Reset()         { *m = SparseMerkleProof{} }
func (m *SparseMerkleProof) String() string { return proto.CompactTextString(m) }
func (*SparseMerkleProof) ProtoMessage()    {}
func (*SparseMerkleProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_473d204b28f447f0, []int{1}
}

func (m *SparseMerkleProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SparseMerkleProof.Unmarshal(m, b)
}
func (m *SparseMerkleProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SparseMerkleProof.Marshal(b, m, deterministic)
}
func (m *SparseMerkleProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SparseMerkleProof.Merge(m, src)
}
func (m *SparseMerkleProof) XXX_Size() int {
	return xxx_messageInfo_SparseMerkleProof.Size(m)
}
func (m *SparseMerkleProof) XXX_DiscardUnknown() {
	xxx_messageInfo_SparseMerkleProof.DiscardUnknown(m)
}

var xxx_messageInfo_SparseMerkleProof proto.InternalMessageInfo

func (m *SparseMerkleProof) GetLeaf() []byte {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *SparseMerkleProof) GetBitmap() []byte {
	if m != nil {
		return m.Bitmap
	}
	return nil
}

func (m *SparseMerkleProof) GetNonDefaultSiblings() [][]byte {
	if m != nil {
		return m.NonDefaultSiblings
	}
	return nil
}

type AccumulatorConsistencyProof struct {
	// The root hashes of the subtrees that represent new leaves. Note that none
	// of these hashes should be default hash.
	Subtrees             [][]byte `protobuf:"bytes,1,rep,name=subtrees,proto3" json:"subtrees,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccumulatorConsistencyProof) Reset()         { *m = AccumulatorConsistencyProof{} }
func (m *AccumulatorConsistencyProof) String() string { return proto.CompactTextString(m) }
func (*AccumulatorConsistencyProof) ProtoMessage()    {}
func (*AccumulatorConsistencyProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_473d204b28f447f0, []int{2}
}

func (m *AccumulatorConsistencyProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccumulatorConsistencyProof.Unmarshal(m, b)
}
func (m *AccumulatorConsistencyProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccumulatorConsistencyProof.Marshal(b, m, deterministic)
}
func (m *AccumulatorConsistencyProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccumulatorConsistencyProof.Merge(m, src)
}
func (m *AccumulatorConsistencyProof) XXX_Size() int {
	return xxx_messageInfo_AccumulatorConsistencyProof.Size(m)
}
func (m *AccumulatorConsistencyProof) XXX_DiscardUnknown() {
	xxx_messageInfo_AccumulatorConsistencyProof.DiscardUnknown(m)
}

var xxx_messageInfo_AccumulatorConsistencyProof proto.InternalMessageInfo

func (m *AccumulatorConsistencyProof) GetSubtrees() [][]byte {
	if m != nil {
		return m.Subtrees
	}
	return nil
}

// The complete proof used to authenticate a signed transaction.
type SignedTransactionProof struct {
	LedgerInfoToTransactionInfoProof *AccumulatorProof `protobuf:"bytes,1,opt,name=ledger_info_to_transaction_info_proof,json=ledgerInfoToTransactionInfoProof,proto3" json:"ledger_info_to_transaction_info_proof,omitempty"`
	TransactionInfo                  *TransactionInfo  `protobuf:"bytes,2,opt,name=transaction_info,json=transactionInfo,proto3" json:"transaction_info,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}          `json:"-"`
	XXX_unrecognized                 []byte            `json:"-"`
	XXX_sizecache                    int32             `json:"-"`
}

func (m *SignedTransactionProof) Reset()         { *m = SignedTransactionProof{} }
func (m *SignedTransactionProof) String() string { return proto.CompactTextString(m) }
func (*SignedTransactionProof) ProtoMessage()    {}
func (*SignedTransactionProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_473d204b28f447f0, []int{3}
}

func (m *SignedTransactionProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransactionProof.Unmarshal(m, b)
}
func (m *SignedTransactionProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransactionProof.Marshal(b, m, deterministic)
}
func (m *SignedTransactionProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransactionProof.Merge(m, src)
}
func (m *SignedTransactionProof) XXX_Size() int {
	return xxx_messageInfo_SignedTransactionProof.Size(m)
}
func (m *SignedTransactionProof) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransactionProof.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransactionProof proto.InternalMessageInfo

func (m *SignedTransactionProof) GetLedgerInfoToTransactionInfoProof() *AccumulatorProof {
	if m != nil {
		return m.LedgerInfoToTransactionInfoProof
	}
	return nil
}

func (m *SignedTransactionProof) GetTransactionInfo() *TransactionInfo {
	if m != nil {
		return m.TransactionInfo
	}
	return nil
}

// The complete proof used to authenticate an account state.
type AccountStateProof struct {
	LedgerInfoToTransactionInfoProof *AccumulatorProof  `protobuf:"bytes,1,opt,name=ledger_info_to_transaction_info_proof,json=ledgerInfoToTransactionInfoProof,proto3" json:"ledger_info_to_transaction_info_proof,omitempty"`
	TransactionInfo                  *TransactionInfo   `protobuf:"bytes,2,opt,name=transaction_info,json=transactionInfo,proto3" json:"transaction_info,omitempty"`
	TransactionInfoToAccountProof    *SparseMerkleProof `protobuf:"bytes,3,opt,name=transaction_info_to_account_proof,json=transactionInfoToAccountProof,proto3" json:"transaction_info_to_account_proof,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}           `json:"-"`
	XXX_unrecognized                 []byte             `json:"-"`
	XXX_sizecache                    int32              `json:"-"`
}

func (m *AccountStateProof) Reset()         { *m = AccountStateProof{} }
func (m *AccountStateProof) String() string { return proto.CompactTextString(m) }
func (*AccountStateProof) ProtoMessage()    {}
func (*AccountStateProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_473d204b28f447f0, []int{4}
}

func (m *AccountStateProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountStateProof.Unmarshal(m, b)
}
func (m *AccountStateProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountStateProof.Marshal(b, m, deterministic)
}
func (m *AccountStateProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountStateProof.Merge(m, src)
}
func (m *AccountStateProof) XXX_Size() int {
	return xxx_messageInfo_AccountStateProof.Size(m)
}
func (m *AccountStateProof) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountStateProof.DiscardUnknown(m)
}

var xxx_messageInfo_AccountStateProof proto.InternalMessageInfo

func (m *AccountStateProof) GetLedgerInfoToTransactionInfoProof() *AccumulatorProof {
	if m != nil {
		return m.LedgerInfoToTransactionInfoProof
	}
	return nil
}

func (m *AccountStateProof) GetTransactionInfo() *TransactionInfo {
	if m != nil {
		return m.TransactionInfo
	}
	return nil
}

func (m *AccountStateProof) GetTransactionInfoToAccountProof() *SparseMerkleProof {
	if m != nil {
		return m.TransactionInfoToAccountProof
	}
	return nil
}

// The complete proof used to authenticate an event.
type EventProof struct {
	LedgerInfoToTransactionInfoProof *AccumulatorProof `protobuf:"bytes,1,opt,name=ledger_info_to_transaction_info_proof,json=ledgerInfoToTransactionInfoProof,proto3" json:"ledger_info_to_transaction_info_proof,omitempty"`
	TransactionInfo                  *TransactionInfo  `protobuf:"bytes,2,opt,name=transaction_info,json=transactionInfo,proto3" json:"transaction_info,omitempty"`
	TransactionInfoToEventProof      *AccumulatorProof `protobuf:"bytes,3,opt,name=transaction_info_to_event_proof,json=transactionInfoToEventProof,proto3" json:"transaction_info_to_event_proof,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}          `json:"-"`
	XXX_unrecognized                 []byte            `json:"-"`
	XXX_sizecache                    int32             `json:"-"`
}

func (m *EventProof) Reset()         { *m = EventProof{} }
func (m *EventProof) String() string { return proto.CompactTextString(m) }
func (*EventProof) ProtoMessage()    {}
func (*EventProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_473d204b28f447f0, []int{5}
}

func (m *EventProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventProof.Unmarshal(m, b)
}
func (m *EventProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventProof.Marshal(b, m, deterministic)
}
func (m *EventProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventProof.Merge(m, src)
}
func (m *EventProof) XXX_Size() int {
	return xxx_messageInfo_EventProof.Size(m)
}
func (m *EventProof) XXX_DiscardUnknown() {
	xxx_messageInfo_EventProof.DiscardUnknown(m)
}

var xxx_messageInfo_EventProof proto.InternalMessageInfo

func (m *EventProof) GetLedgerInfoToTransactionInfoProof() *AccumulatorProof {
	if m != nil {
		return m.LedgerInfoToTransactionInfoProof
	}
	return nil
}

func (m *EventProof) GetTransactionInfo() *TransactionInfo {
	if m != nil {
		return m.TransactionInfo
	}
	return nil
}

func (m *EventProof) GetTransactionInfoToEventProof() *AccumulatorProof {
	if m != nil {
		return m.TransactionInfoToEventProof
	}
	return nil
}

func init() {
	proto.RegisterType((*AccumulatorProof)(nil), "types.AccumulatorProof")
	proto.RegisterType((*SparseMerkleProof)(nil), "types.SparseMerkleProof")
	proto.RegisterType((*AccumulatorConsistencyProof)(nil), "types.AccumulatorConsistencyProof")
	proto.RegisterType((*SignedTransactionProof)(nil), "types.SignedTransactionProof")
	proto.RegisterType((*AccountStateProof)(nil), "types.AccountStateProof")
	proto.RegisterType((*EventProof)(nil), "types.EventProof")
}

func init() { proto.RegisterFile("proof.proto", fileDescriptor_473d204b28f447f0) }

var fileDescriptor_473d204b28f447f0 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0x99, 0x99, 0x75, 0x91, 0x9a, 0x05, 0x77, 0x83, 0x8c, 0x61, 0x17, 0x71, 0x0c, 0x08,
	0x7b, 0xd0, 0x44, 0xc6, 0x83, 0xec, 0x71, 0xfc, 0x73, 0xf0, 0x20, 0x48, 0x32, 0x27, 0x51, 0x42,
	0x77, 0x52, 0xc9, 0x34, 0x66, 0xba, 0x62, 0x77, 0x45, 0xd8, 0x37, 0xf1, 0x01, 0x7c, 0x15, 0xdf,
	0x4b, 0xa6, 0x3b, 0xb8, 0xb3, 0x19, 0x17, 0xcf, 0x73, 0x4b, 0x55, 0x51, 0xdf, 0x57, 0xdf, 0x8f,
	0xd0, 0x30, 0x6d, 0x0d, 0x51, 0x15, 0xb7, 0x86, 0x98, 0x82, 0x7b, 0x7c, 0xdd, 0xa2, 0x3d, 0x9f,
	0xb1, 0x11, 0xda, 0x8a, 0x82, 0x15, 0xe9, 0x5c, 0xe9, 0x8a, 0xfc, 0x38, 0xfa, 0x02, 0xa7, 0xcb,
	0xa2, 0xe8, 0x36, 0x5d, 0x23, 0x98, 0xcc, 0xa7, 0xed, 0x62, 0x30, 0x83, 0x63, 0xa9, 0x78, 0x23,
	0xda, 0x70, 0x34, 0x1f, 0x5d, 0x1e, 0xa5, 0x7d, 0x15, 0xbc, 0x84, 0x87, 0x9a, 0x74, 0x5e, 0x62,
	0x25, 0xba, 0x86, 0x73, 0xab, 0x64, 0xa3, 0x74, 0x6d, 0xc3, 0xf1, 0x7c, 0x72, 0x79, 0x92, 0x06,
	0x9a, 0xf4, 0x3b, 0x3f, 0xca, 0xfa, 0x49, 0xf4, 0x1d, 0xce, 0xb2, 0x56, 0x18, 0x8b, 0x1f, 0xd1,
	0x7c, 0x6b, 0xd0, 0xcb, 0x07, 0x70, 0xd4, 0xa0, 0xa8, 0x9c, 0xf8, 0x49, 0xea, 0xbe, 0x77, 0x2c,
	0xc7, 0xae, 0xfb, 0x3f, 0xcb, 0xc9, 0x9d, 0x96, 0x57, 0x70, 0xb1, 0x13, 0xe8, 0x2d, 0x69, 0xab,
	0x2c, 0xa3, 0x2e, 0xae, 0xbd, 0xf9, 0x39, 0xdc, 0xb7, 0x9d, 0x64, 0x83, 0x68, 0xc3, 0x91, 0x13,
	0xf9, 0x5b, 0x47, 0xbf, 0x47, 0x30, 0xcb, 0x54, 0xad, 0xb1, 0x5c, 0xdd, 0xc0, 0xf2, 0x6b, 0x6b,
	0x78, 0xd6, 0x60, 0x59, 0xa3, 0x71, 0xec, 0x72, 0xa6, 0x7c, 0xc8, 0x33, 0x77, 0xd0, 0x5d, 0xa8,
	0xe9, 0xe2, 0x51, 0xec, 0xa8, 0xc7, 0x43, 0xb4, 0xe9, 0xdc, 0xab, 0x7c, 0xd0, 0x15, 0xad, 0x68,
	0xc7, 0x65, 0xdb, 0xf0, 0x4e, 0x4b, 0x38, 0x1d, 0x4a, 0x3b, 0x26, 0xd3, 0xc5, 0xac, 0x17, 0x1d,
	0xac, 0xa5, 0x0f, 0xf8, 0x76, 0x23, 0xfa, 0x35, 0x86, 0xb3, 0x65, 0x51, 0x50, 0xa7, 0x39, 0x63,
	0xc1, 0x78, 0x78, 0x11, 0x02, 0x09, 0x4f, 0xf7, 0xae, 0x63, 0xca, 0x85, 0x4f, 0xd5, 0x1f, 0x3a,
	0x71, 0x9a, 0x61, 0xaf, 0xb9, 0xf7, 0xa3, 0xa5, 0x8f, 0x07, 0xaa, 0x2b, 0xea, 0xa9, 0xb8, 0x71,
	0xf4, 0x73, 0x0c, 0xf0, 0xfe, 0x07, 0xf6, 0xe5, 0x61, 0xf1, 0xf9, 0x0a, 0x4f, 0xfe, 0xc5, 0x07,
	0xb7, 0x71, 0x6e, 0xd1, 0xb9, 0xf3, 0xcc, 0x8b, 0x3d, 0x38, 0x37, 0x2c, 0xde, 0xc4, 0x9f, 0x9f,
	0xd7, 0x8a, 0xd7, 0x9d, 0x8c, 0x0b, 0xda, 0x24, 0xbc, 0xc6, 0xd7, 0x8b, 0xab, 0xa4, 0xa6, 0x17,
	0x8d, 0x92, 0x46, 0x24, 0x35, 0x6a, 0x34, 0x82, 0xb1, 0x4c, 0x5a, 0xe9, 0xd4, 0xe5, 0xb1, 0x7b,
	0x4c, 0x5e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xd2, 0xcd, 0x5d, 0x7a, 0x04, 0x00, 0x00,
}
