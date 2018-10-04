// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maintenance.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InstallRequest struct {
	Account              *InstallRequest_Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *InstallRequest) Reset()         { *m = InstallRequest{} }
func (m *InstallRequest) String() string { return proto.CompactTextString(m) }
func (*InstallRequest) ProtoMessage()    {}
func (*InstallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0}
}
func (m *InstallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallRequest.Unmarshal(m, b)
}
func (m *InstallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallRequest.Marshal(b, m, deterministic)
}
func (m *InstallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallRequest.Merge(m, src)
}
func (m *InstallRequest) XXX_Size() int {
	return xxx_messageInfo_InstallRequest.Size(m)
}
func (m *InstallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstallRequest proto.InternalMessageInfo

func (m *InstallRequest) GetAccount() *InstallRequest_Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type InstallRequest_Account struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Users                []*InstallRequest_User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *InstallRequest_Account) Reset()         { *m = InstallRequest_Account{} }
func (m *InstallRequest_Account) String() string { return proto.CompactTextString(m) }
func (*InstallRequest_Account) ProtoMessage()    {}
func (*InstallRequest_Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0, 0}
}
func (m *InstallRequest_Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallRequest_Account.Unmarshal(m, b)
}
func (m *InstallRequest_Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallRequest_Account.Marshal(b, m, deterministic)
}
func (m *InstallRequest_Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallRequest_Account.Merge(m, src)
}
func (m *InstallRequest_Account) XXX_Size() int {
	return xxx_messageInfo_InstallRequest_Account.Size(m)
}
func (m *InstallRequest_Account) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallRequest_Account.DiscardUnknown(m)
}

var xxx_messageInfo_InstallRequest_Account proto.InternalMessageInfo

func (m *InstallRequest_Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallRequest_Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstallRequest_Account) GetUsers() []*InstallRequest_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type InstallRequest_User struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tokens               []*InstallRequest_Token `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *InstallRequest_User) Reset()         { *m = InstallRequest_User{} }
func (m *InstallRequest_User) String() string { return proto.CompactTextString(m) }
func (*InstallRequest_User) ProtoMessage()    {}
func (*InstallRequest_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0, 1}
}
func (m *InstallRequest_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallRequest_User.Unmarshal(m, b)
}
func (m *InstallRequest_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallRequest_User.Marshal(b, m, deterministic)
}
func (m *InstallRequest_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallRequest_User.Merge(m, src)
}
func (m *InstallRequest_User) XXX_Size() int {
	return xxx_messageInfo_InstallRequest_User.Size(m)
}
func (m *InstallRequest_User) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallRequest_User.DiscardUnknown(m)
}

var xxx_messageInfo_InstallRequest_User proto.InternalMessageInfo

func (m *InstallRequest_User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallRequest_User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstallRequest_User) GetTokens() []*InstallRequest_Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type InstallRequest_Token struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InstallRequest_Token) Reset()         { *m = InstallRequest_Token{} }
func (m *InstallRequest_Token) String() string { return proto.CompactTextString(m) }
func (*InstallRequest_Token) ProtoMessage()    {}
func (*InstallRequest_Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0, 2}
}
func (m *InstallRequest_Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallRequest_Token.Unmarshal(m, b)
}
func (m *InstallRequest_Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallRequest_Token.Marshal(b, m, deterministic)
}
func (m *InstallRequest_Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallRequest_Token.Merge(m, src)
}
func (m *InstallRequest_Token) XXX_Size() int {
	return xxx_messageInfo_InstallRequest_Token.Size(m)
}
func (m *InstallRequest_Token) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallRequest_Token.DiscardUnknown(m)
}

var xxx_messageInfo_InstallRequest_Token proto.InternalMessageInfo

func (m *InstallRequest_Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallRequest_Token) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

type InstallResponse struct {
	Account              *InstallResponse_Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *InstallResponse) Reset()         { *m = InstallResponse{} }
func (m *InstallResponse) String() string { return proto.CompactTextString(m) }
func (*InstallResponse) ProtoMessage()    {}
func (*InstallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{1}
}
func (m *InstallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallResponse.Unmarshal(m, b)
}
func (m *InstallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallResponse.Marshal(b, m, deterministic)
}
func (m *InstallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallResponse.Merge(m, src)
}
func (m *InstallResponse) XXX_Size() int {
	return xxx_messageInfo_InstallResponse.Size(m)
}
func (m *InstallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InstallResponse proto.InternalMessageInfo

func (m *InstallResponse) GetAccount() *InstallResponse_Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type InstallResponse_Account struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp    `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp    `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Users                []*InstallResponse_User `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *InstallResponse_Account) Reset()         { *m = InstallResponse_Account{} }
func (m *InstallResponse_Account) String() string { return proto.CompactTextString(m) }
func (*InstallResponse_Account) ProtoMessage()    {}
func (*InstallResponse_Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{1, 0}
}
func (m *InstallResponse_Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallResponse_Account.Unmarshal(m, b)
}
func (m *InstallResponse_Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallResponse_Account.Marshal(b, m, deterministic)
}
func (m *InstallResponse_Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallResponse_Account.Merge(m, src)
}
func (m *InstallResponse_Account) XXX_Size() int {
	return xxx_messageInfo_InstallResponse_Account.Size(m)
}
func (m *InstallResponse_Account) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallResponse_Account.DiscardUnknown(m)
}

var xxx_messageInfo_InstallResponse_Account proto.InternalMessageInfo

func (m *InstallResponse_Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallResponse_Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstallResponse_Account) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *InstallResponse_Account) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *InstallResponse_Account) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *InstallResponse_Account) GetUsers() []*InstallResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type InstallResponse_User struct {
	Id                   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp     `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp     `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp     `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Tokens               []*InstallResponse_Token `protobuf:"bytes,6,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *InstallResponse_User) Reset()         { *m = InstallResponse_User{} }
func (m *InstallResponse_User) String() string { return proto.CompactTextString(m) }
func (*InstallResponse_User) ProtoMessage()    {}
func (*InstallResponse_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{1, 1}
}
func (m *InstallResponse_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallResponse_User.Unmarshal(m, b)
}
func (m *InstallResponse_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallResponse_User.Marshal(b, m, deterministic)
}
func (m *InstallResponse_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallResponse_User.Merge(m, src)
}
func (m *InstallResponse_User) XXX_Size() int {
	return xxx_messageInfo_InstallResponse_User.Size(m)
}
func (m *InstallResponse_User) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallResponse_User.DiscardUnknown(m)
}

var xxx_messageInfo_InstallResponse_User proto.InternalMessageInfo

func (m *InstallResponse_User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallResponse_User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstallResponse_User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *InstallResponse_User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *InstallResponse_User) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *InstallResponse_User) GetTokens() []*InstallResponse_Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type InstallResponse_Token struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Revoked              bool                 `protobuf:"varint,2,opt,name=revoked,proto3" json:"revoked,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InstallResponse_Token) Reset()         { *m = InstallResponse_Token{} }
func (m *InstallResponse_Token) String() string { return proto.CompactTextString(m) }
func (*InstallResponse_Token) ProtoMessage()    {}
func (*InstallResponse_Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{1, 2}
}
func (m *InstallResponse_Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallResponse_Token.Unmarshal(m, b)
}
func (m *InstallResponse_Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallResponse_Token.Marshal(b, m, deterministic)
}
func (m *InstallResponse_Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallResponse_Token.Merge(m, src)
}
func (m *InstallResponse_Token) XXX_Size() int {
	return xxx_messageInfo_InstallResponse_Token.Size(m)
}
func (m *InstallResponse_Token) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallResponse_Token.DiscardUnknown(m)
}

var xxx_messageInfo_InstallResponse_Token proto.InternalMessageInfo

func (m *InstallResponse_Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallResponse_Token) GetRevoked() bool {
	if m != nil {
		return m.Revoked
	}
	return false
}

func (m *InstallResponse_Token) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

func (m *InstallResponse_Token) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *InstallResponse_Token) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *InstallResponse_Token) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallRequest)(nil), "grpc.InstallRequest")
	proto.RegisterType((*InstallRequest_Account)(nil), "grpc.InstallRequest.Account")
	proto.RegisterType((*InstallRequest_User)(nil), "grpc.InstallRequest.User")
	proto.RegisterType((*InstallRequest_Token)(nil), "grpc.InstallRequest.Token")
	proto.RegisterType((*InstallResponse)(nil), "grpc.InstallResponse")
	proto.RegisterType((*InstallResponse_Account)(nil), "grpc.InstallResponse.Account")
	proto.RegisterType((*InstallResponse_User)(nil), "grpc.InstallResponse.User")
	proto.RegisterType((*InstallResponse_Token)(nil), "grpc.InstallResponse.Token")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MaintenanceClient is the client API for Maintenance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaintenanceClient interface {
	Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*InstallResponse, error)
}

type maintenanceClient struct {
	cc *grpc.ClientConn
}

func NewMaintenanceClient(cc *grpc.ClientConn) MaintenanceClient {
	return &maintenanceClient{cc}
}

func (c *maintenanceClient) Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*InstallResponse, error) {
	out := new(InstallResponse)
	err := c.cc.Invoke(ctx, "/grpc.Maintenance/Install", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaintenanceServer is the server API for Maintenance service.
type MaintenanceServer interface {
	Install(context.Context, *InstallRequest) (*InstallResponse, error)
}

func RegisterMaintenanceServer(s *grpc.Server, srv MaintenanceServer) {
	s.RegisterService(&_Maintenance_serviceDesc, srv)
}

func _Maintenance_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Maintenance/Install",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Install(ctx, req.(*InstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Maintenance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Maintenance",
	HandlerType: (*MaintenanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Install",
			Handler:    _Maintenance_Install_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maintenance.proto",
}

func init() { proto.RegisterFile("maintenance.proto", fileDescriptor_6053ae89a3b3f561) }

var fileDescriptor_6053ae89a3b3f561 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0x75, 0xbe, 0xdc, 0xa5, 0x7d, 0x27, 0x8a, 0x30, 0x20, 0x1d, 0xa1, 0x88, 0xea, 0xa6,
	0x8a, 0x21, 0x41, 0x57, 0x09, 0x54, 0xb6, 0x1b, 0x11, 0x62, 0xb1, 0xca, 0xc2, 0x50, 0xe4, 0x26,
	0x8f, 0x93, 0xd5, 0xc4, 0x0e, 0xb1, 0x83, 0x98, 0xf9, 0x04, 0x27, 0x31, 0xf3, 0xa9, 0x58, 0x19,
	0xf9, 0x20, 0x28, 0xb6, 0x53, 0x68, 0x1a, 0x74, 0xe4, 0xb6, 0x6e, 0xb1, 0xfd, 0x7b, 0xef, 0x59,
	0xbf, 0xfc, 0x65, 0xb8, 0x57, 0x70, 0x21, 0x0d, 0x4a, 0x2e, 0x53, 0x8c, 0xcb, 0x4a, 0x19, 0x45,
	0x83, 0x75, 0x55, 0xa6, 0xd1, 0xe1, 0x5a, 0xa9, 0x75, 0x8e, 0x09, 0x2f, 0x45, 0xc2, 0xa5, 0x54,
	0x86, 0x1b, 0xa1, 0xa4, 0x76, 0x4c, 0xf4, 0xd4, 0x9f, 0xda, 0xd5, 0x45, 0xfd, 0x31, 0x31, 0xa2,
	0x40, 0x6d, 0x78, 0x51, 0x3a, 0x60, 0xf1, 0x93, 0xc0, 0xc1, 0x6b, 0xa9, 0x0d, 0xcf, 0x73, 0x86,
	0x9f, 0x6a, 0xd4, 0x86, 0xbe, 0x80, 0x90, 0xa7, 0xa9, 0xaa, 0xa5, 0x99, 0x8f, 0x8e, 0x46, 0xc7,
	0xb3, 0xe5, 0x61, 0xdc, 0x4c, 0x8a, 0xaf, 0x63, 0xf1, 0xca, 0x31, 0xac, 0x85, 0xa3, 0x73, 0x08,
	0xfd, 0x1e, 0x3d, 0x00, 0x22, 0x32, 0x5b, 0xbd, 0xcf, 0x88, 0xc8, 0x28, 0x85, 0x40, 0xf2, 0x02,
	0xe7, 0xc4, 0xee, 0xd8, 0x6f, 0x9a, 0xc0, 0xa4, 0xd6, 0x58, 0xe9, 0xf9, 0xf8, 0x68, 0x7c, 0x3c,
	0x5b, 0x3e, 0xea, 0x1d, 0xf2, 0x4e, 0x63, 0xc5, 0x1c, 0x17, 0x9d, 0x43, 0xd0, 0x2c, 0xff, 0xab,
	0xf9, 0x12, 0xa6, 0x46, 0x5d, 0xa2, 0x6c, 0xbb, 0x47, 0xbd, 0xdd, 0xcf, 0x1a, 0x84, 0x79, 0x32,
	0x62, 0x30, 0xb1, 0x1b, 0x37, 0x06, 0x9c, 0x02, 0xe0, 0x97, 0x52, 0x54, 0x98, 0x7d, 0xe0, 0xc6,
	0x8e, 0xb1, 0x0d, 0xad, 0xd9, 0xb8, 0x35, 0x1b, 0x9f, 0xb5, 0x66, 0xd9, 0xbe, 0xa7, 0x57, 0x66,
	0xb1, 0x09, 0xe1, 0xee, 0xd5, 0x50, 0x5d, 0x2a, 0xa9, 0x91, 0xbe, 0xec, 0xfa, 0x7d, 0xd2, 0xb9,
	0x9c, 0xe3, 0x6e, 0x0a, 0xde, 0x90, 0x61, 0x86, 0x4f, 0x01, 0xd2, 0x0a, 0xb9, 0x71, 0xf7, 0x1e,
	0x6f, 0xbf, 0xb7, 0xa7, 0x57, 0xa6, 0x29, 0xad, 0xcb, 0xac, 0x2d, 0x0d, 0xb6, 0x97, 0x7a, 0xda,
	0x95, 0x66, 0x98, 0xa3, 0x2f, 0x9d, 0x6c, 0x2f, 0xf5, 0xf4, 0xca, 0xd0, 0xe7, 0x6d, 0x24, 0xa6,
	0xbd, 0x3f, 0xcd, 0x7b, 0xf9, 0x3b, 0x13, 0x1b, 0x32, 0x20, 0x14, 0xb7, 0xce, 0xc7, 0xc9, 0x55,
	0x8a, 0x9d, 0x90, 0xc7, 0xfd, 0x42, 0xae, 0xc7, 0xf8, 0x3b, 0xf9, 0x57, 0x8e, 0xe7, 0x10, 0x56,
	0xf8, 0x59, 0x5d, 0x62, 0x66, 0xb5, 0xec, 0xb1, 0x76, 0xd9, 0x49, 0xf8, 0x78, 0x40, 0xc2, 0x3b,
	0x52, 0x83, 0xdd, 0xa5, 0x4e, 0x76, 0x97, 0x3a, 0x1d, 0x20, 0x75, 0xf9, 0x1e, 0x66, 0x6f, 0xff,
	0xbc, 0xa5, 0xf4, 0x0d, 0x84, 0xde, 0x27, 0x7d, 0xd0, 0xf7, 0x48, 0x44, 0x0f, 0x7b, 0xa5, 0x2f,
	0xee, 0x7f, 0xfd, 0xf1, 0xeb, 0x1b, 0xb9, 0xb3, 0xd8, 0x4b, 0x84, 0x3b, 0x79, 0x35, 0x7a, 0x76,
	0x31, 0xb5, 0xa3, 0x4f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x56, 0xe0, 0xe3, 0xae, 0x05,
	0x00, 0x00,
}
