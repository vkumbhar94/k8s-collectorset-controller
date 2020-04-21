// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/topic_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Use topics to target or exclude placements in the Google Display Network
// based on the category into which the placement falls (for example,
// "Pets & Animals/Pets/Dogs").
type TopicConstant struct {
	// The resource name of the topic constant.
	// topic constant resource names have the form:
	//
	// `topicConstants/{topic_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the topic.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Resource name of parent of the topic constant.
	TopicConstantParent *wrappers.StringValue `protobuf:"bytes,3,opt,name=topic_constant_parent,json=topicConstantParent,proto3" json:"topic_constant_parent,omitempty"`
	// The category to target or exclude. Each subsequent element in the array
	// describes a more specific sub-category. For example,
	// {"Pets & Animals", "Pets", "Dogs"} represents the
	// "Pets & Animals/Pets/Dogs" category. List of available topic categories at
	// https://developers.google.com/adwords/api/docs/appendix/verticals
	Path                 []*wrappers.StringValue `protobuf:"bytes,4,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TopicConstant) Reset()         { *m = TopicConstant{} }
func (m *TopicConstant) String() string { return proto.CompactTextString(m) }
func (*TopicConstant) ProtoMessage()    {}
func (*TopicConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8e2bc7b19f8427, []int{0}
}

func (m *TopicConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicConstant.Unmarshal(m, b)
}
func (m *TopicConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicConstant.Marshal(b, m, deterministic)
}
func (m *TopicConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicConstant.Merge(m, src)
}
func (m *TopicConstant) XXX_Size() int {
	return xxx_messageInfo_TopicConstant.Size(m)
}
func (m *TopicConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicConstant.DiscardUnknown(m)
}

var xxx_messageInfo_TopicConstant proto.InternalMessageInfo

func (m *TopicConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *TopicConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TopicConstant) GetTopicConstantParent() *wrappers.StringValue {
	if m != nil {
		return m.TopicConstantParent
	}
	return nil
}

func (m *TopicConstant) GetPath() []*wrappers.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicConstant)(nil), "google.ads.googleads.v2.resources.TopicConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/topic_constant.proto", fileDescriptor_bc8e2bc7b19f8427)
}

var fileDescriptor_bc8e2bc7b19f8427 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xfb, 0x30,
	0x00, 0xc7, 0x69, 0x37, 0x7e, 0xf0, 0xab, 0xee, 0x52, 0x11, 0xca, 0x1c, 0xb2, 0x29, 0x83, 0x81,
	0x90, 0x4a, 0x95, 0x1d, 0xe2, 0xa9, 0xf3, 0x30, 0xf4, 0x20, 0x65, 0x4a, 0x0f, 0x52, 0x18, 0x59,
	0x1b, 0x6b, 0x60, 0x4b, 0x42, 0x92, 0xcd, 0xf7, 0xf1, 0xe8, 0xa3, 0xf8, 0x1e, 0x5e, 0x7c, 0x09,
	0xa5, 0x4d, 0x13, 0x1c, 0x82, 0xbb, 0x7d, 0x69, 0x3e, 0xdf, 0x3f, 0x4d, 0xbc, 0x71, 0xc9, 0x58,
	0xb9, 0xc4, 0x21, 0x2a, 0x64, 0xa8, 0x65, 0xa5, 0x36, 0x51, 0x28, 0xb0, 0x64, 0x6b, 0x91, 0x63,
	0x19, 0x2a, 0xc6, 0x49, 0x3e, 0xcf, 0x19, 0x95, 0x0a, 0x51, 0x05, 0xb8, 0x60, 0x8a, 0xf9, 0x03,
	0x0d, 0x03, 0x54, 0x48, 0x60, 0x7d, 0x60, 0x13, 0x01, 0xeb, 0xeb, 0x1e, 0x37, 0xd1, 0xb5, 0x61,
	0xb1, 0x7e, 0x0a, 0x5f, 0x04, 0xe2, 0x1c, 0x0b, 0xa9, 0x23, 0xba, 0x3d, 0x53, 0xcd, 0x49, 0x88,
	0x28, 0x65, 0x0a, 0x29, 0xc2, 0x68, 0x73, 0x7a, 0xf2, 0xe1, 0x78, 0x9d, 0x87, 0xaa, 0xf9, 0xba,
	0x29, 0xf6, 0x4f, 0xbd, 0x8e, 0x09, 0x9f, 0x53, 0xb4, 0xc2, 0x81, 0xd3, 0x77, 0x46, 0xff, 0x67,
	0xfb, 0xe6, 0xe3, 0x1d, 0x5a, 0x61, 0xff, 0xcc, 0x73, 0x49, 0x11, 0xb8, 0x7d, 0x67, 0xb4, 0x17,
	0x1d, 0x35, 0xcb, 0x80, 0x59, 0x00, 0x6e, 0xa8, 0x1a, 0x5f, 0xa6, 0x68, 0xb9, 0xc6, 0x33, 0x97,
	0x14, 0x7e, 0xe2, 0x1d, 0x6e, 0xff, 0xdc, 0x9c, 0x23, 0x81, 0xa9, 0x0a, 0x5a, 0xb5, 0xbf, 0xf7,
	0xcb, 0x7f, 0xaf, 0x04, 0xa1, 0xa5, 0x0e, 0x38, 0x50, 0x3f, 0xd7, 0x25, 0xb5, 0xd1, 0x3f, 0xf7,
	0xda, 0x1c, 0xa9, 0xe7, 0xa0, 0xdd, 0x6f, 0xed, 0x0c, 0xa8, 0xc9, 0xc9, 0x97, 0xe3, 0x0d, 0x73,
	0xb6, 0x02, 0x3b, 0xef, 0x73, 0xe2, 0x6f, 0x5d, 0x47, 0x52, 0x45, 0x26, 0xce, 0xe3, 0x6d, 0x63,
	0x2c, 0xd9, 0x12, 0xd1, 0x12, 0x30, 0x51, 0x86, 0x25, 0xa6, 0x75, 0xa1, 0x79, 0x50, 0x4e, 0xe4,
	0x1f, 0xef, 0x7b, 0x65, 0xd5, 0xab, 0xdb, 0x9a, 0xc6, 0xf1, 0x9b, 0x3b, 0x98, 0xea, 0xc8, 0xb8,
	0x90, 0x40, 0xcb, 0x4a, 0xa5, 0x11, 0x98, 0x19, 0xf2, 0xdd, 0x30, 0x59, 0x5c, 0xc8, 0xcc, 0x32,
	0x59, 0x1a, 0x65, 0x96, 0xf9, 0x74, 0x87, 0xfa, 0x00, 0xc2, 0xb8, 0x90, 0x10, 0x5a, 0x0a, 0xc2,
	0x34, 0x82, 0xd0, 0x72, 0x8b, 0x7f, 0xf5, 0xd8, 0x8b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0xbb, 0xd3, 0x4f, 0x8b, 0x02, 0x00, 0x00,
}
