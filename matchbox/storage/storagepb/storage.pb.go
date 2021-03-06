// Code generated by protoc-gen-go. DO NOT EDIT.
// source: matchbox/storage/storagepb/storage.proto

package storagepb

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

// Group selects one or more machines and matches them to a Profile.
type Group struct {
	// machine readable Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// human readable name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Profile id
	Profile string `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	// Selectors to match machines
	Selector map[string]string `protobuf:"bytes,4,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// JSON encoded metadata
	Metadata             []byte   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed97bf224c67cd0, []int{0}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *Group) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *Group) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Profile defines the boot and provisioning behavior of a group of machines.
type Profile struct {
	// profile id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// human readable name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// ignition id
	IgnitionId string `protobuf:"bytes,3,opt,name=ignition_id,json=ignitionId,proto3" json:"ignition_id,omitempty"`
	// cloud config id
	CloudId string `protobuf:"bytes,4,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// support network boot / PXE
	Boot *NetBoot `protobuf:"bytes,5,opt,name=boot,proto3" json:"boot,omitempty"`
	// generic config id
	GenericId            string   `protobuf:"bytes,6,opt,name=generic_id,json=genericId,proto3" json:"generic_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed97bf224c67cd0, []int{1}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetIgnitionId() string {
	if m != nil {
		return m.IgnitionId
	}
	return ""
}

func (m *Profile) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

func (m *Profile) GetBoot() *NetBoot {
	if m != nil {
		return m.Boot
	}
	return nil
}

func (m *Profile) GetGenericId() string {
	if m != nil {
		return m.GenericId
	}
	return ""
}

// NetBoot describes network or PXE boot settings for a machine.
type NetBoot struct {
	// the URL of the kernel image
	Kernel string `protobuf:"bytes,1,opt,name=kernel,proto3" json:"kernel,omitempty"`
	// the init RAM filesystem URLs
	Initrd []string `protobuf:"bytes,2,rep,name=initrd,proto3" json:"initrd,omitempty"`
	// kernel args
	Args                 []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetBoot) Reset()         { *m = NetBoot{} }
func (m *NetBoot) String() string { return proto.CompactTextString(m) }
func (*NetBoot) ProtoMessage()    {}
func (*NetBoot) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ed97bf224c67cd0, []int{2}
}

func (m *NetBoot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetBoot.Unmarshal(m, b)
}
func (m *NetBoot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetBoot.Marshal(b, m, deterministic)
}
func (m *NetBoot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetBoot.Merge(m, src)
}
func (m *NetBoot) XXX_Size() int {
	return xxx_messageInfo_NetBoot.Size(m)
}
func (m *NetBoot) XXX_DiscardUnknown() {
	xxx_messageInfo_NetBoot.DiscardUnknown(m)
}

var xxx_messageInfo_NetBoot proto.InternalMessageInfo

func (m *NetBoot) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *NetBoot) GetInitrd() []string {
	if m != nil {
		return m.Initrd
	}
	return nil
}

func (m *NetBoot) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "storagepb.Group")
	proto.RegisterMapType((map[string]string)(nil), "storagepb.Group.SelectorEntry")
	proto.RegisterType((*Profile)(nil), "storagepb.Profile")
	proto.RegisterType((*NetBoot)(nil), "storagepb.NetBoot")
}

func init() {
	proto.RegisterFile("matchbox/storage/storagepb/storage.proto", fileDescriptor_5ed97bf224c67cd0)
}

var fileDescriptor_5ed97bf224c67cd0 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0x1f, 0x89, 0xed, 0x49, 0x5b, 0x82, 0x28, 0xc5, 0x0d, 0xb4, 0x35, 0x39, 0x14, 0x9f,
	0x1c, 0x48, 0x0f, 0x6d, 0xd3, 0x5b, 0xa0, 0x94, 0xf4, 0x50, 0x8a, 0x7b, 0x2b, 0x85, 0x22, 0x5b,
	0x5a, 0x47, 0xc4, 0xd6, 0x18, 0x59, 0x5e, 0x36, 0xff, 0x6a, 0x7f, 0xcf, 0xfe, 0x9a, 0xc5, 0x8a,
	0x6c, 0x76, 0x0f, 0x0b, 0x7b, 0xf2, 0x7b, 0x6f, 0xc6, 0x9a, 0xf7, 0x86, 0x81, 0xb4, 0xa1, 0xba,
	0x3c, 0x16, 0x78, 0xb3, 0xe9, 0x34, 0x2a, 0x5a, 0xf1, 0xf1, 0xdb, 0x16, 0x23, 0xca, 0x5a, 0x85,
	0x1a, 0x49, 0x34, 0x15, 0xd6, 0x77, 0x0e, 0xcc, 0x7e, 0x28, 0xec, 0x5b, 0xf2, 0x0a, 0x5c, 0xc1,
	0x62, 0x27, 0x71, 0xd2, 0x28, 0x77, 0x05, 0x23, 0x04, 0x7c, 0x49, 0x1b, 0x1e, 0xbb, 0x46, 0x31,
	0x98, 0xc4, 0x10, 0xb4, 0x0a, 0xaf, 0x44, 0xcd, 0x63, 0xcf, 0xc8, 0x23, 0x25, 0x3b, 0x08, 0x3b,
	0x5e, 0xf3, 0x52, 0xa3, 0x8a, 0xfd, 0xc4, 0x4b, 0x17, 0xdb, 0xf7, 0xd9, 0x34, 0x25, 0x33, 0x13,
	0xb2, 0x3f, 0xb6, 0xe1, 0xbb, 0xd4, 0xea, 0x9c, 0x4f, 0xfd, 0x64, 0x05, 0x61, 0xc3, 0x35, 0x65,
	0x54, 0xd3, 0x78, 0x96, 0x38, 0xe9, 0x8b, 0x7c, 0xe2, 0xab, 0x6f, 0xf0, 0xf2, 0xd1, 0x6f, 0x64,
	0x09, 0xde, 0x89, 0x9f, 0xad, 0xcf, 0x01, 0x92, 0xd7, 0x30, 0xbb, 0xa6, 0x75, 0x3f, 0x3a, 0xbd,
	0x90, 0x9d, 0xfb, 0xc5, 0x59, 0xdf, 0x3a, 0x10, 0xfc, 0xb6, 0x06, 0x9f, 0x13, 0xef, 0x03, 0x2c,
	0x44, 0x25, 0x85, 0x16, 0x28, 0xff, 0x0b, 0x66, 0x23, 0xc2, 0x28, 0x1d, 0x18, 0x79, 0x0b, 0x61,
	0x59, 0x63, 0xcf, 0x86, 0xaa, 0x7f, 0x59, 0x80, 0xe1, 0x07, 0x46, 0x3e, 0x82, 0x5f, 0x20, 0x6a,
	0x13, 0x60, 0xb1, 0x25, 0x0f, 0xc2, 0xff, 0xe2, 0x7a, 0x8f, 0xa8, 0x73, 0x53, 0x27, 0xef, 0x00,
	0x2a, 0x2e, 0xb9, 0x12, 0xe5, 0xf0, 0xc8, 0xdc, 0x3c, 0x12, 0x59, 0xe5, 0xc0, 0xd6, 0xff, 0x20,
	0xb0, 0xfd, 0xe4, 0x0d, 0xcc, 0x4f, 0x5c, 0x49, 0x5e, 0x5b, 0xd7, 0x96, 0x0d, 0xba, 0x90, 0x42,
	0x2b, 0x16, 0xbb, 0x89, 0x37, 0xe8, 0x17, 0x36, 0x24, 0xa2, 0xaa, 0xea, 0xcc, 0xfa, 0xa3, 0xdc,
	0xe0, 0x9f, 0x7e, 0xe8, 0x2d, 0xfd, 0x3c, 0x28, 0x1b, 0x56, 0x0b, 0xc9, 0xf7, 0x5f, 0xff, 0x7e,
	0xae, 0x84, 0x3e, 0xf6, 0x45, 0x56, 0x62, 0xb3, 0x69, 0xb1, 0xe3, 0x82, 0xa1, 0xdc, 0x4c, 0x87,
	0xf3, 0xf4, 0x05, 0x15, 0x73, 0x73, 0x3a, 0x9f, 0xee, 0x03, 0x00, 0x00, 0xff, 0xff, 0x37, 0x8a,
	0x89, 0x62, 0x66, 0x02, 0x00, 0x00,
}
