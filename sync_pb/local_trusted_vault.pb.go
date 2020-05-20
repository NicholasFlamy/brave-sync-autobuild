// Code generated by protoc-gen-go. DO NOT EDIT.
// source: local_trusted_vault.proto

package sync_pb

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

type LocalTrustedVaultKey struct {
	// The actual key.
	KeyMaterial          []byte   `protobuf:"bytes,1,opt,name=key_material,json=keyMaterial" json:"key_material,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalTrustedVaultKey) Reset()         { *m = LocalTrustedVaultKey{} }
func (m *LocalTrustedVaultKey) String() string { return proto.CompactTextString(m) }
func (*LocalTrustedVaultKey) ProtoMessage()    {}
func (*LocalTrustedVaultKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2b0fbfa463d1617, []int{0}
}

func (m *LocalTrustedVaultKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTrustedVaultKey.Unmarshal(m, b)
}
func (m *LocalTrustedVaultKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTrustedVaultKey.Marshal(b, m, deterministic)
}
func (m *LocalTrustedVaultKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTrustedVaultKey.Merge(m, src)
}
func (m *LocalTrustedVaultKey) XXX_Size() int {
	return xxx_messageInfo_LocalTrustedVaultKey.Size(m)
}
func (m *LocalTrustedVaultKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTrustedVaultKey.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTrustedVaultKey proto.InternalMessageInfo

func (m *LocalTrustedVaultKey) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

type LocalTrustedVaultPerUser struct {
	// User identifier.
	GaiaId []byte `protobuf:"bytes,1,opt,name=gaia_id,json=gaiaId" json:"gaia_id,omitempty"`
	// All keys known for a user.
	Key []*LocalTrustedVaultKey `protobuf:"bytes,2,rep,name=key" json:"key,omitempty"`
	// The version corresponding to the last element in |key|.
	LastKeyVersion       *int32   `protobuf:"varint,3,opt,name=last_key_version,json=lastKeyVersion" json:"last_key_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalTrustedVaultPerUser) Reset()         { *m = LocalTrustedVaultPerUser{} }
func (m *LocalTrustedVaultPerUser) String() string { return proto.CompactTextString(m) }
func (*LocalTrustedVaultPerUser) ProtoMessage()    {}
func (*LocalTrustedVaultPerUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2b0fbfa463d1617, []int{1}
}

func (m *LocalTrustedVaultPerUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTrustedVaultPerUser.Unmarshal(m, b)
}
func (m *LocalTrustedVaultPerUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTrustedVaultPerUser.Marshal(b, m, deterministic)
}
func (m *LocalTrustedVaultPerUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTrustedVaultPerUser.Merge(m, src)
}
func (m *LocalTrustedVaultPerUser) XXX_Size() int {
	return xxx_messageInfo_LocalTrustedVaultPerUser.Size(m)
}
func (m *LocalTrustedVaultPerUser) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTrustedVaultPerUser.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTrustedVaultPerUser proto.InternalMessageInfo

func (m *LocalTrustedVaultPerUser) GetGaiaId() []byte {
	if m != nil {
		return m.GaiaId
	}
	return nil
}

func (m *LocalTrustedVaultPerUser) GetKey() []*LocalTrustedVaultKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LocalTrustedVaultPerUser) GetLastKeyVersion() int32 {
	if m != nil && m.LastKeyVersion != nil {
		return *m.LastKeyVersion
	}
	return 0
}

type LocalTrustedVault struct {
	User                 []*LocalTrustedVaultPerUser `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LocalTrustedVault) Reset()         { *m = LocalTrustedVault{} }
func (m *LocalTrustedVault) String() string { return proto.CompactTextString(m) }
func (*LocalTrustedVault) ProtoMessage()    {}
func (*LocalTrustedVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2b0fbfa463d1617, []int{2}
}

func (m *LocalTrustedVault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTrustedVault.Unmarshal(m, b)
}
func (m *LocalTrustedVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTrustedVault.Marshal(b, m, deterministic)
}
func (m *LocalTrustedVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTrustedVault.Merge(m, src)
}
func (m *LocalTrustedVault) XXX_Size() int {
	return xxx_messageInfo_LocalTrustedVault.Size(m)
}
func (m *LocalTrustedVault) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTrustedVault.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTrustedVault proto.InternalMessageInfo

func (m *LocalTrustedVault) GetUser() []*LocalTrustedVaultPerUser {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalTrustedVaultKey)(nil), "sync_pb.LocalTrustedVaultKey")
	proto.RegisterType((*LocalTrustedVaultPerUser)(nil), "sync_pb.LocalTrustedVaultPerUser")
	proto.RegisterType((*LocalTrustedVault)(nil), "sync_pb.LocalTrustedVault")
}

func init() {
	proto.RegisterFile("local_trusted_vault.proto", fileDescriptor_c2b0fbfa463d1617)
}

var fileDescriptor_c2b0fbfa463d1617 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x49, 0xa3, 0x16, 0xa6, 0x45, 0x34, 0x08, 0xc6, 0x83, 0xb0, 0xdd, 0x53, 0x4e, 0x2b,
	0x08, 0x1e, 0xbc, 0x7a, 0x52, 0xab, 0x20, 0x41, 0x7b, 0x0d, 0xb1, 0x3b, 0x48, 0xd8, 0xd8, 0x94,
	0x24, 0x5b, 0xc8, 0x9f, 0xf0, 0x37, 0x4b, 0x76, 0x73, 0x5b, 0x3d, 0xce, 0xe3, 0x7b, 0xf3, 0x3e,
	0xb8, 0xb2, 0x6e, 0xab, 0xad, 0x8a, 0xbe, 0x0f, 0x11, 0x5b, 0x75, 0xd0, 0xbd, 0x8d, 0xcd, 0xde,
	0xbb, 0xe8, 0xd8, 0x3c, 0xa4, 0xdd, 0x56, 0xed, 0x3f, 0xeb, 0x7b, 0xb8, 0x78, 0xc9, 0xd4, 0xfb,
	0x08, 0x6d, 0x32, 0xb3, 0xc6, 0xc4, 0x56, 0xb0, 0xec, 0x30, 0xa9, 0x6f, 0x1d, 0xd1, 0x1b, 0x6d,
	0x39, 0xa9, 0x88, 0x58, 0xca, 0x45, 0x87, 0xe9, 0xb5, 0x44, 0xf5, 0x0f, 0x01, 0x3e, 0xe9, 0xbe,
	0xa1, 0xff, 0x08, 0xe8, 0xd9, 0x25, 0xcc, 0xbf, 0xb4, 0xd1, 0xca, 0xb4, 0xa5, 0x7a, 0x92, 0xcf,
	0xa7, 0x96, 0xdd, 0x00, 0xed, 0x30, 0xf1, 0x59, 0x45, 0xc5, 0xe2, 0xf6, 0xba, 0x29, 0x1e, 0xcd,
	0x5f, 0x12, 0x32, 0x93, 0x4c, 0xc0, 0x99, 0xd5, 0x21, 0xaa, 0xac, 0x73, 0x40, 0x1f, 0x8c, 0xdb,
	0x71, 0x5a, 0x11, 0x71, 0x2c, 0x4f, 0x73, 0xbe, 0xc6, 0xb4, 0x19, 0xd3, 0xfa, 0x19, 0xce, 0x27,
	0x6f, 0xd8, 0x1d, 0x1c, 0xf5, 0x01, 0x3d, 0x27, 0xc3, 0xe0, 0xea, 0xff, 0xc1, 0x62, 0x2e, 0x07,
	0xfc, 0x61, 0xf6, 0x48, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x61, 0x40, 0xdf, 0xde, 0x40, 0x01,
	0x00, 0x00,
}
