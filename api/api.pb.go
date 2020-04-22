// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message for a collector ID.
type CollectorIDRequest struct {
	Orchestrator         string   `protobuf:"bytes,2,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectorIDRequest) Reset()         { *m = CollectorIDRequest{} }
func (m *CollectorIDRequest) String() string { return proto.CompactTextString(m) }
func (*CollectorIDRequest) ProtoMessage()    {}
func (*CollectorIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *CollectorIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectorIDRequest.Unmarshal(m, b)
}
func (m *CollectorIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectorIDRequest.Marshal(b, m, deterministic)
}
func (m *CollectorIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectorIDRequest.Merge(m, src)
}
func (m *CollectorIDRequest) XXX_Size() int {
	return xxx_messageInfo_CollectorIDRequest.Size(m)
}
func (m *CollectorIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectorIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectorIDRequest proto.InternalMessageInfo

func (m *CollectorIDRequest) GetOrchestrator() string {
	if m != nil {
		return m.Orchestrator
	}
	return ""
}

// The response message from a collector ID request.
type CollectorIDReply struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectorIDReply) Reset()         { *m = CollectorIDReply{} }
func (m *CollectorIDReply) String() string { return proto.CompactTextString(m) }
func (*CollectorIDReply) ProtoMessage()    {}
func (*CollectorIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *CollectorIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectorIDReply.Unmarshal(m, b)
}
func (m *CollectorIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectorIDReply.Marshal(b, m, deterministic)
}
func (m *CollectorIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectorIDReply.Merge(m, src)
}
func (m *CollectorIDReply) XXX_Size() int {
	return xxx_messageInfo_CollectorIDReply.Size(m)
}
func (m *CollectorIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectorIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_CollectorIDReply proto.InternalMessageInfo

func (m *CollectorIDReply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CollectorIDRequest)(nil), "api.CollectorIDRequest")
	proto.RegisterType((*CollectorIDReply)(nil), "api.CollectorIDReply")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0xb2, 0xe0, 0x12, 0x72, 0xce,
	0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0x2f, 0xf2, 0x74, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x52, 0xe2, 0xe2, 0xc9, 0x2f, 0x4a, 0xce, 0x48, 0x2d, 0x2e, 0x29, 0x4a, 0x2c, 0xc9, 0x2f,
	0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x11, 0x53, 0x52, 0xe2, 0x12, 0x40, 0xd1, 0x59,
	0x90, 0x53, 0x29, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4,
	0x94, 0x99, 0x62, 0x14, 0xc9, 0x25, 0x06, 0x57, 0x13, 0x9c, 0x5a, 0xe2, 0x9c, 0x9f, 0x57, 0x52,
	0x04, 0xe2, 0x17, 0x09, 0xd9, 0x73, 0x71, 0x23, 0xe9, 0x16, 0x12, 0xd7, 0x03, 0xb9, 0x0b, 0xd3,
	0x25, 0x52, 0xa2, 0x98, 0x12, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x4f, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x31, 0xf4, 0x7e, 0x42, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CollectorSetControllerClient is the client API for CollectorSetController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectorSetControllerClient interface {
	// Retrieves a collector ID.
	CollectorID(ctx context.Context, in *CollectorIDRequest, opts ...grpc.CallOption) (*CollectorIDReply, error)
}

type collectorSetControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorSetControllerClient(cc grpc.ClientConnInterface) CollectorSetControllerClient {
	return &collectorSetControllerClient{cc}
}

func (c *collectorSetControllerClient) CollectorID(ctx context.Context, in *CollectorIDRequest, opts ...grpc.CallOption) (*CollectorIDReply, error) {
	out := new(CollectorIDReply)
	err := c.cc.Invoke(ctx, "/api.CollectorSetController/CollectorID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorSetControllerServer is the server API for CollectorSetController service.
type CollectorSetControllerServer interface {
	// Retrieves a collector ID.
	CollectorID(context.Context, *CollectorIDRequest) (*CollectorIDReply, error)
}

// UnimplementedCollectorSetControllerServer can be embedded to have forward compatible implementations.
type UnimplementedCollectorSetControllerServer struct {
}

func (*UnimplementedCollectorSetControllerServer) CollectorID(ctx context.Context, req *CollectorIDRequest) (*CollectorIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectorID not implemented")
}

func RegisterCollectorSetControllerServer(s *grpc.Server, srv CollectorSetControllerServer) {
	s.RegisterService(&_CollectorSetController_serviceDesc, srv)
}

func _CollectorSetController_CollectorID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectorIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorSetControllerServer).CollectorID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CollectorSetController/CollectorID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorSetControllerServer).CollectorID(ctx, req.(*CollectorIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorSetController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CollectorSetController",
	HandlerType: (*CollectorSetControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectorID",
			Handler:    _CollectorSetController_CollectorID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
