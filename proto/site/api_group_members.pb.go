// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_group_members.proto

package site // import "github.com/duckchat/duckchat-gateway/proto/site"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/yayashanzei/duckchat-gateway/proto/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApiGroupMembersUserProfile struct {
	Profile              *core.PublicUserProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Type                 core.GroupMemberType    `protobuf:"varint,2,opt,name=type,proto3,enum=core.GroupMemberType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ApiGroupMembersUserProfile) Reset()         { *m = ApiGroupMembersUserProfile{} }
func (m *ApiGroupMembersUserProfile) String() string { return proto.CompactTextString(m) }
func (*ApiGroupMembersUserProfile) ProtoMessage()    {}
func (*ApiGroupMembersUserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_members_9ef6d8e346dd226b, []int{0}
}
func (m *ApiGroupMembersUserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupMembersUserProfile.Unmarshal(m, b)
}
func (m *ApiGroupMembersUserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupMembersUserProfile.Marshal(b, m, deterministic)
}
func (dst *ApiGroupMembersUserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupMembersUserProfile.Merge(dst, src)
}
func (m *ApiGroupMembersUserProfile) XXX_Size() int {
	return xxx_messageInfo_ApiGroupMembersUserProfile.Size(m)
}
func (m *ApiGroupMembersUserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupMembersUserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupMembersUserProfile proto.InternalMessageInfo

func (m *ApiGroupMembersUserProfile) GetProfile() *core.PublicUserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *ApiGroupMembersUserProfile) GetType() core.GroupMemberType {
	if m != nil {
		return m.Type
	}
	return core.GroupMemberType_GroupMemberGuest
}

// *
//
// action: api.group.members
//
type ApiGroupMembersRequest struct {
	GroupId              string   `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiGroupMembersRequest) Reset()         { *m = ApiGroupMembersRequest{} }
func (m *ApiGroupMembersRequest) String() string { return proto.CompactTextString(m) }
func (*ApiGroupMembersRequest) ProtoMessage()    {}
func (*ApiGroupMembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_members_9ef6d8e346dd226b, []int{1}
}
func (m *ApiGroupMembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupMembersRequest.Unmarshal(m, b)
}
func (m *ApiGroupMembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupMembersRequest.Marshal(b, m, deterministic)
}
func (dst *ApiGroupMembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupMembersRequest.Merge(dst, src)
}
func (m *ApiGroupMembersRequest) XXX_Size() int {
	return xxx_messageInfo_ApiGroupMembersRequest.Size(m)
}
func (m *ApiGroupMembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupMembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupMembersRequest proto.InternalMessageInfo

func (m *ApiGroupMembersRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *ApiGroupMembersRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ApiGroupMembersRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ApiGroupMembersResponse struct {
	List                 []*ApiGroupMembersUserProfile `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	TotalCount           int32                         `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ApiGroupMembersResponse) Reset()         { *m = ApiGroupMembersResponse{} }
func (m *ApiGroupMembersResponse) String() string { return proto.CompactTextString(m) }
func (*ApiGroupMembersResponse) ProtoMessage()    {}
func (*ApiGroupMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_group_members_9ef6d8e346dd226b, []int{2}
}
func (m *ApiGroupMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiGroupMembersResponse.Unmarshal(m, b)
}
func (m *ApiGroupMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiGroupMembersResponse.Marshal(b, m, deterministic)
}
func (dst *ApiGroupMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiGroupMembersResponse.Merge(dst, src)
}
func (m *ApiGroupMembersResponse) XXX_Size() int {
	return xxx_messageInfo_ApiGroupMembersResponse.Size(m)
}
func (m *ApiGroupMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiGroupMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiGroupMembersResponse proto.InternalMessageInfo

func (m *ApiGroupMembersResponse) GetList() []*ApiGroupMembersUserProfile {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ApiGroupMembersResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ApiGroupMembersUserProfile)(nil), "site.ApiGroupMembersUserProfile")
	proto.RegisterType((*ApiGroupMembersRequest)(nil), "site.ApiGroupMembersRequest")
	proto.RegisterType((*ApiGroupMembersResponse)(nil), "site.ApiGroupMembersResponse")
}

func init() {
	proto.RegisterFile("site/api_group_members.proto", fileDescriptor_api_group_members_9ef6d8e346dd226b)
}

var fileDescriptor_api_group_members_9ef6d8e346dd226b = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x7f, 0x71, 0x0b, 0x56, 0x56, 0x6d, 0x43, 0x11, 0x09, 0xbd, 0xaa, 0x17, 0x6e,
	0xb0, 0xfa, 0x02, 0xea, 0x85, 0x78, 0x21, 0x94, 0xa8, 0x20, 0xa5, 0x50, 0x37, 0xdb, 0x69, 0xbb,
	0x98, 0x74, 0xd7, 0xfd, 0x41, 0xd2, 0x47, 0xf4, 0xa9, 0x24, 0xbb, 0x56, 0x82, 0xe2, 0xdd, 0xcc,
	0x99, 0x6f, 0xce, 0x19, 0x06, 0x9d, 0x68, 0x6e, 0x20, 0xa6, 0x92, 0xcf, 0x57, 0x4a, 0x58, 0x39,
	0xcf, 0x21, 0x4f, 0x41, 0x69, 0x22, 0x95, 0x30, 0x02, 0x37, 0xca, 0xe9, 0xe0, 0x80, 0x09, 0x05,
	0xb1, 0x9b, 0x7b, 0x7d, 0xd0, 0x75, 0x8a, 0xd5, 0xa0, 0xbc, 0x30, 0xdc, 0xa2, 0xc1, 0xb5, 0xe4,
	0x77, 0x25, 0xf2, 0xe0, 0x1d, 0x9e, 0x35, 0xa8, 0x89, 0x12, 0x4b, 0x9e, 0x01, 0xbe, 0x40, 0x6d,
	0xe9, 0xcb, 0x30, 0x88, 0x82, 0x51, 0x67, 0xdc, 0x27, 0xa5, 0x01, 0x99, 0xd8, 0x34, 0xe3, 0xac,
	0x42, 0x26, 0x3b, 0x0e, 0x9f, 0xa1, 0x86, 0x29, 0x24, 0x84, 0xb5, 0x28, 0x18, 0xed, 0x8f, 0x8f,
	0x3d, 0x5f, 0xf1, 0x7f, 0x2a, 0x24, 0x24, 0x0e, 0x19, 0xbe, 0xa2, 0xde, 0xaf, 0xec, 0x04, 0xde,
	0x2d, 0x68, 0x83, 0x43, 0xd4, 0x76, 0x57, 0xdf, 0x2f, 0x5c, 0xee, 0x5e, 0xb2, 0x6b, 0x71, 0x0f,
	0xb5, 0xc4, 0x72, 0xa9, 0xc1, 0xb8, 0x80, 0x66, 0xf2, 0xdd, 0xe1, 0x23, 0xd4, 0x64, 0xc2, 0x6e,
	0x4c, 0x58, 0x77, 0xb2, 0x6f, 0x86, 0x02, 0xf5, 0xff, 0x24, 0x68, 0x29, 0x36, 0x1a, 0xf0, 0x15,
	0x6a, 0x64, 0x5c, 0x9b, 0x30, 0x88, 0xea, 0xa3, 0xce, 0x38, 0x22, 0xe5, 0xc3, 0xc8, 0xff, 0xaf,
	0x48, 0x1c, 0x8d, 0x4f, 0x11, 0x32, 0xc2, 0xd0, 0xec, 0xd6, 0x65, 0xf9, 0x13, 0x2a, 0xca, 0xcd,
	0x0b, 0x3a, 0x64, 0x22, 0x27, 0x5b, 0x9a, 0x15, 0xfe, 0xc1, 0xce, 0x77, 0x1a, 0xaf, 0xb8, 0x59,
	0xdb, 0x94, 0x30, 0x91, 0xc7, 0x0b, 0xcb, 0xde, 0xd8, 0x9a, 0x9a, 0x9f, 0xe2, 0x7c, 0x45, 0x0d,
	0x7c, 0xd0, 0x22, 0x76, 0x0b, 0x71, 0xb9, 0xf0, 0x59, 0xeb, 0x4e, 0x69, 0x56, 0xcc, 0x26, 0xa5,
	0x32, 0x7b, 0xe4, 0x06, 0xd2, 0x96, 0x9b, 0x5e, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x34, 0x4f,
	0x10, 0x81, 0xf8, 0x01, 0x00, 0x00,
}
