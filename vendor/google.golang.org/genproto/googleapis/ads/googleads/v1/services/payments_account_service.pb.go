// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/payments_account_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for fetching all accessible payments accounts.
type ListPaymentsAccountsRequest struct {
	// The ID of the customer to apply the PaymentsAccount list operation to.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPaymentsAccountsRequest) Reset()         { *m = ListPaymentsAccountsRequest{} }
func (m *ListPaymentsAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsRequest) ProtoMessage()    {}
func (*ListPaymentsAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68acb38ba4f095b7, []int{0}
}

func (m *ListPaymentsAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Marshal(b, m, deterministic)
}
func (m *ListPaymentsAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsRequest.Merge(m, src)
}
func (m *ListPaymentsAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Size(m)
}
func (m *ListPaymentsAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsRequest proto.InternalMessageInfo

func (m *ListPaymentsAccountsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for [PaymentsAccountService.ListPaymentsAccounts][google.ads.googleads.v1.services.PaymentsAccountService.ListPaymentsAccounts].
type ListPaymentsAccountsResponse struct {
	// The list of accessible payments accounts.
	PaymentsAccounts     []*resources.PaymentsAccount `protobuf:"bytes,1,rep,name=payments_accounts,json=paymentsAccounts,proto3" json:"payments_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListPaymentsAccountsResponse) Reset()         { *m = ListPaymentsAccountsResponse{} }
func (m *ListPaymentsAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsResponse) ProtoMessage()    {}
func (*ListPaymentsAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68acb38ba4f095b7, []int{1}
}

func (m *ListPaymentsAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Marshal(b, m, deterministic)
}
func (m *ListPaymentsAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsResponse.Merge(m, src)
}
func (m *ListPaymentsAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Size(m)
}
func (m *ListPaymentsAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsResponse proto.InternalMessageInfo

func (m *ListPaymentsAccountsResponse) GetPaymentsAccounts() []*resources.PaymentsAccount {
	if m != nil {
		return m.PaymentsAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPaymentsAccountsRequest)(nil), "google.ads.googleads.v1.services.ListPaymentsAccountsRequest")
	proto.RegisterType((*ListPaymentsAccountsResponse)(nil), "google.ads.googleads.v1.services.ListPaymentsAccountsResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/payments_account_service.proto", fileDescriptor_68acb38ba4f095b7)
}

var fileDescriptor_68acb38ba4f095b7 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x25, 0xf9, 0x40, 0x70, 0xba, 0xd1, 0x20, 0x5a, 0xd2, 0x82, 0xa1, 0x74, 0x51, 0x5c, 0xcc,
	0x98, 0x0a, 0x22, 0x23, 0xad, 0xa4, 0x9b, 0x2a, 0xb8, 0x28, 0x15, 0xba, 0x90, 0x40, 0x88, 0xc9,
	0x10, 0x02, 0xcd, 0x4c, 0xcc, 0x9d, 0x14, 0x44, 0x44, 0x28, 0xf8, 0x04, 0xbe, 0x81, 0x4b, 0x1f,
	0xa5, 0xe0, 0xca, 0x57, 0x70, 0xe5, 0x53, 0x48, 0x3a, 0x99, 0xb4, 0x84, 0xd4, 0xc2, 0xb7, 0x3b,
	0xe4, 0x9e, 0x9f, 0x3b, 0x67, 0x26, 0xe8, 0x55, 0x22, 0x44, 0xb2, 0x65, 0x24, 0x8c, 0x81, 0x28,
	0x58, 0xa1, 0x9d, 0x4b, 0x80, 0x15, 0xbb, 0x34, 0x62, 0x40, 0xf2, 0xf0, 0x53, 0xc6, 0xb8, 0x84,
	0x20, 0x8c, 0x22, 0x51, 0x72, 0x19, 0xd4, 0x13, 0x9c, 0x17, 0x42, 0x0a, 0xcb, 0x51, 0x2a, 0x1c,
	0xc6, 0x80, 0x1b, 0x03, 0xbc, 0x73, 0xb1, 0x36, 0xb0, 0x5f, 0x5c, 0x8a, 0x28, 0x18, 0x88, 0xb2,
	0xe8, 0xca, 0x50, 0xde, 0xf6, 0x50, 0x2b, 0xf3, 0x94, 0x84, 0x9c, 0x0b, 0x19, 0xca, 0x54, 0x70,
	0xa8, 0xa7, 0x8f, 0xce, 0xa6, 0xd1, 0x36, 0x65, 0x5a, 0x36, 0x9a, 0xa3, 0xc1, 0xdb, 0x14, 0xe4,
	0xaa, 0x36, 0xf5, 0x94, 0x27, 0xac, 0xd9, 0xc7, 0x92, 0x81, 0xb4, 0x1e, 0xa3, 0x5e, 0x54, 0x82,
	0x14, 0x19, 0x2b, 0x82, 0x34, 0xee, 0x1b, 0x8e, 0x31, 0xb9, 0xbb, 0x46, 0xfa, 0xd3, 0x9b, 0x78,
	0xf4, 0x15, 0x0d, 0xbb, 0xf5, 0x90, 0x0b, 0x0e, 0xcc, 0x0a, 0xd0, 0xfd, 0xf6, 0xc2, 0xd0, 0x37,
	0x9c, 0x9b, 0x49, 0x6f, 0x3a, 0xc5, 0x97, 0xea, 0x68, 0x0e, 0x8b, 0x5b, 0xbe, 0xeb, 0x7b, 0x79,
	0x2b, 0x68, 0xfa, 0xcd, 0x44, 0x0f, 0x5b, 0xac, 0x77, 0xaa, 0x4d, 0xeb, 0x97, 0x81, 0x1e, 0x74,
	0x2d, 0x67, 0xcd, 0xf0, 0xb5, 0x8b, 0xc0, 0xff, 0x29, 0xc5, 0x9e, 0xdf, 0x56, 0xae, 0x3a, 0x19,
	0x3d, 0xdf, 0xff, 0xfe, 0xf3, 0xdd, 0x7c, 0x6a, 0xe1, 0xea, 0x62, 0x75, 0x97, 0x40, 0x3e, 0x9f,
	0x35, 0x3d, 0x7b, 0xf2, 0x85, 0xb4, 0x8f, 0x6a, 0x0f, 0x0e, 0x5e, 0xff, 0x14, 0x57, 0xa3, 0x3c,
	0x05, 0x1c, 0x89, 0x6c, 0xb1, 0x37, 0xd1, 0x38, 0x12, 0xd9, 0xd5, 0xd5, 0x16, 0x83, 0xee, 0xb6,
	0x56, 0xd5, 0x73, 0x58, 0x19, 0xef, 0x5f, 0xd7, 0x06, 0x89, 0xd8, 0x86, 0x3c, 0xc1, 0xa2, 0x48,
	0x48, 0xc2, 0xf8, 0xf1, 0xb1, 0x90, 0x53, 0xe4, 0xe5, 0x7f, 0xe0, 0xa5, 0x06, 0x3f, 0xcc, 0x9b,
	0xa5, 0xe7, 0xfd, 0x34, 0x9d, 0xa5, 0x32, 0xf4, 0x62, 0xc0, 0x0a, 0x56, 0x68, 0xe3, 0xe2, 0x3a,
	0x18, 0x0e, 0x9a, 0xe2, 0x7b, 0x31, 0xf8, 0x0d, 0xc5, 0xdf, 0xb8, 0xbe, 0xa6, 0xfc, 0x35, 0xc7,
	0xea, 0x3b, 0xa5, 0x5e, 0x0c, 0x94, 0x36, 0x24, 0x4a, 0x37, 0x2e, 0xa5, 0x9a, 0xf6, 0xe1, 0xce,
	0x71, 0xcf, 0x67, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x99, 0x8e, 0xe2, 0x0d, 0xaa, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentsAccountServiceClient is the client API for PaymentsAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentsAccountServiceClient interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error)
}

type paymentsAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsAccountServiceClient(cc grpc.ClientConnInterface) PaymentsAccountServiceClient {
	return &paymentsAccountServiceClient{cc}
}

func (c *paymentsAccountServiceClient) ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error) {
	out := new(ListPaymentsAccountsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.PaymentsAccountService/ListPaymentsAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsAccountServiceServer is the server API for PaymentsAccountService service.
type PaymentsAccountServiceServer interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error)
}

// UnimplementedPaymentsAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentsAccountServiceServer struct {
}

func (*UnimplementedPaymentsAccountServiceServer) ListPaymentsAccounts(ctx context.Context, req *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPaymentsAccounts not implemented")
}

func RegisterPaymentsAccountServiceServer(s *grpc.Server, srv PaymentsAccountServiceServer) {
	s.RegisterService(&_PaymentsAccountService_serviceDesc, srv)
}

func _PaymentsAccountService_ListPaymentsAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.PaymentsAccountService/ListPaymentsAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, req.(*ListPaymentsAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentsAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.PaymentsAccountService",
	HandlerType: (*PaymentsAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPaymentsAccounts",
			Handler:    _PaymentsAccountService_ListPaymentsAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/payments_account_service.proto",
}
