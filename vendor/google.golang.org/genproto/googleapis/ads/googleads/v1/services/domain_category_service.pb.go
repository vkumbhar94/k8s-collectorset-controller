// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/domain_category_service.proto

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

// Request message for
// [DomainCategoryService.GetDomainCategory][google.ads.googleads.v1.services.DomainCategoryService.GetDomainCategory].
type GetDomainCategoryRequest struct {
	// Resource name of the domain category to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDomainCategoryRequest) Reset()         { *m = GetDomainCategoryRequest{} }
func (m *GetDomainCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetDomainCategoryRequest) ProtoMessage()    {}
func (*GetDomainCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e4bcffa65a4ff8, []int{0}
}

func (m *GetDomainCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDomainCategoryRequest.Unmarshal(m, b)
}
func (m *GetDomainCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDomainCategoryRequest.Marshal(b, m, deterministic)
}
func (m *GetDomainCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDomainCategoryRequest.Merge(m, src)
}
func (m *GetDomainCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetDomainCategoryRequest.Size(m)
}
func (m *GetDomainCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDomainCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDomainCategoryRequest proto.InternalMessageInfo

func (m *GetDomainCategoryRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDomainCategoryRequest)(nil), "google.ads.googleads.v1.services.GetDomainCategoryRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/domain_category_service.proto", fileDescriptor_15e4bcffa65a4ff8)
}

var fileDescriptor_15e4bcffa65a4ff8 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0xb9, 0x70, 0xe1, 0x86, 0x7b, 0x17, 0x37, 0x20, 0x96, 0xe8, 0xa2, 0xd4, 0x2e, 0xa4,
	0x8b, 0x19, 0x52, 0x05, 0x61, 0x44, 0x25, 0x55, 0xa8, 0x2b, 0x29, 0x15, 0xba, 0x90, 0x40, 0x19,
	0x93, 0x21, 0x04, 0x9a, 0x99, 0x9a, 0x33, 0x2d, 0x88, 0xb8, 0xd0, 0x57, 0xf0, 0x0d, 0x5c, 0xfa,
	0x10, 0x3e, 0x40, 0xb7, 0xbe, 0x82, 0x2b, 0xf7, 0xee, 0x25, 0x9d, 0x99, 0xd4, 0x62, 0x43, 0x77,
	0x1f, 0x73, 0xbe, 0x9f, 0x73, 0xbe, 0xc4, 0x39, 0x4e, 0x84, 0x48, 0x46, 0x0c, 0xd3, 0x18, 0xb0,
	0x82, 0x05, 0x9a, 0xfa, 0x18, 0x58, 0x3e, 0x4d, 0x23, 0x06, 0x38, 0x16, 0x19, 0x4d, 0xf9, 0x30,
	0xa2, 0x92, 0x25, 0x22, 0xbf, 0x1d, 0xea, 0x01, 0x1a, 0xe7, 0x42, 0x0a, 0xb7, 0xae, 0x44, 0x88,
	0xc6, 0x80, 0x4a, 0x3d, 0x9a, 0xfa, 0xc8, 0xe8, 0xbd, 0x83, 0xaa, 0x84, 0x9c, 0x81, 0x98, 0xe4,
	0x2b, 0x22, 0x94, 0xb5, 0xb7, 0x6d, 0x84, 0xe3, 0x14, 0x53, 0xce, 0x85, 0xa4, 0x32, 0x15, 0x1c,
	0xf4, 0x74, 0xf3, 0xdb, 0x34, 0x1a, 0xa5, 0x8c, 0x4b, 0x35, 0x68, 0x9c, 0x38, 0xb5, 0x2e, 0x93,
	0x67, 0x73, 0xcb, 0x53, 0xed, 0xd8, 0x67, 0x37, 0x13, 0x06, 0xd2, 0xdd, 0x71, 0xfe, 0x99, 0xd4,
	0x21, 0xa7, 0x19, 0xab, 0x59, 0x75, 0x6b, 0xf7, 0x4f, 0xff, 0xaf, 0x79, 0xbc, 0xa0, 0x19, 0x6b,
	0x7f, 0x5a, 0xce, 0xc6, 0xb2, 0xfc, 0x52, 0xdd, 0xe2, 0xbe, 0x5a, 0xce, 0xff, 0x1f, 0xde, 0x2e,
	0x41, 0xeb, 0x3a, 0x40, 0x55, 0x0b, 0x79, 0x7e, 0xa5, 0xb6, 0x6c, 0x07, 0x2d, 0x2b, 0x1b, 0xe4,
	0xf1, 0xed, 0xfd, 0xc9, 0xde, 0x77, 0xdb, 0x45, 0x87, 0x77, 0x4b, 0xe7, 0x1c, 0x45, 0x13, 0x90,
	0x22, 0x63, 0x39, 0xe0, 0x96, 0x2e, 0x55, 0xcb, 0x52, 0x06, 0xb8, 0x75, 0xef, 0x6d, 0xcd, 0x82,
	0xda, 0x22, 0x45, 0xa3, 0x71, 0x0a, 0x28, 0x12, 0x59, 0xe7, 0xc1, 0x76, 0x9a, 0x91, 0xc8, 0xd6,
	0x5e, 0xd3, 0xf1, 0x56, 0xb6, 0xd3, 0x2b, 0xda, 0xef, 0x59, 0x57, 0xe7, 0x5a, 0x9f, 0x88, 0x11,
	0xe5, 0x09, 0x12, 0x79, 0x82, 0x13, 0xc6, 0xe7, 0xdf, 0x06, 0x2f, 0x12, 0xab, 0x7f, 0xb8, 0x43,
	0x03, 0x9e, 0xed, 0x5f, 0xdd, 0x20, 0x78, 0xb1, 0xeb, 0x5d, 0x65, 0x18, 0xc4, 0x80, 0x14, 0x2c,
	0xd0, 0xc0, 0x47, 0x3a, 0x18, 0x66, 0x86, 0x12, 0x06, 0x31, 0x84, 0x25, 0x25, 0x1c, 0xf8, 0xa1,
	0xa1, 0x7c, 0xd8, 0x4d, 0xf5, 0x4e, 0x48, 0x10, 0x03, 0x21, 0x25, 0x89, 0x90, 0x81, 0x4f, 0x88,
	0xa1, 0x5d, 0xff, 0x9e, 0xef, 0xb9, 0xf7, 0x15, 0x00, 0x00, 0xff, 0xff, 0x90, 0x91, 0x15, 0xd7,
	0x17, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DomainCategoryServiceClient is the client API for DomainCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DomainCategoryServiceClient interface {
	// Returns the requested domain category.
	GetDomainCategory(ctx context.Context, in *GetDomainCategoryRequest, opts ...grpc.CallOption) (*resources.DomainCategory, error)
}

type domainCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainCategoryServiceClient(cc grpc.ClientConnInterface) DomainCategoryServiceClient {
	return &domainCategoryServiceClient{cc}
}

func (c *domainCategoryServiceClient) GetDomainCategory(ctx context.Context, in *GetDomainCategoryRequest, opts ...grpc.CallOption) (*resources.DomainCategory, error) {
	out := new(resources.DomainCategory)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.DomainCategoryService/GetDomainCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainCategoryServiceServer is the server API for DomainCategoryService service.
type DomainCategoryServiceServer interface {
	// Returns the requested domain category.
	GetDomainCategory(context.Context, *GetDomainCategoryRequest) (*resources.DomainCategory, error)
}

// UnimplementedDomainCategoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDomainCategoryServiceServer struct {
}

func (*UnimplementedDomainCategoryServiceServer) GetDomainCategory(ctx context.Context, req *GetDomainCategoryRequest) (*resources.DomainCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainCategory not implemented")
}

func RegisterDomainCategoryServiceServer(s *grpc.Server, srv DomainCategoryServiceServer) {
	s.RegisterService(&_DomainCategoryService_serviceDesc, srv)
}

func _DomainCategoryService_GetDomainCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainCategoryServiceServer).GetDomainCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.DomainCategoryService/GetDomainCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainCategoryServiceServer).GetDomainCategory(ctx, req.(*GetDomainCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DomainCategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.DomainCategoryService",
	HandlerType: (*DomainCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDomainCategory",
			Handler:    _DomainCategoryService_GetDomainCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/domain_category_service.proto",
}
