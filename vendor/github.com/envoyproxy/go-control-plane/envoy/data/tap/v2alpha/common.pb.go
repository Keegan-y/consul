// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/common.proto

package envoy_data_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Body struct {
	// Types that are valid to be assigned to BodyType:
	//	*Body_AsBytes
	//	*Body_AsString
	BodyType             isBody_BodyType `protobuf_oneof:"body_type"`
	Truncated            bool            `protobuf:"varint,3,opt,name=truncated,proto3" json:"truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Body) Reset()         { *m = Body{} }
func (m *Body) String() string { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()    {}
func (*Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_a560f1e2899ebe7a, []int{0}
}

func (m *Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Body.Unmarshal(m, b)
}
func (m *Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Body.Marshal(b, m, deterministic)
}
func (m *Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Body.Merge(m, src)
}
func (m *Body) XXX_Size() int {
	return xxx_messageInfo_Body.Size(m)
}
func (m *Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Body.DiscardUnknown(m)
}

var xxx_messageInfo_Body proto.InternalMessageInfo

type isBody_BodyType interface {
	isBody_BodyType()
}

type Body_AsBytes struct {
	AsBytes []byte `protobuf:"bytes,1,opt,name=as_bytes,json=asBytes,proto3,oneof"`
}

type Body_AsString struct {
	AsString string `protobuf:"bytes,2,opt,name=as_string,json=asString,proto3,oneof"`
}

func (*Body_AsBytes) isBody_BodyType() {}

func (*Body_AsString) isBody_BodyType() {}

func (m *Body) GetBodyType() isBody_BodyType {
	if m != nil {
		return m.BodyType
	}
	return nil
}

func (m *Body) GetAsBytes() []byte {
	if x, ok := m.GetBodyType().(*Body_AsBytes); ok {
		return x.AsBytes
	}
	return nil
}

func (m *Body) GetAsString() string {
	if x, ok := m.GetBodyType().(*Body_AsString); ok {
		return x.AsString
	}
	return ""
}

func (m *Body) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Body) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Body_AsBytes)(nil),
		(*Body_AsString)(nil),
	}
}

func init() {
	proto.RegisterType((*Body)(nil), "envoy.data.tap.v2alpha.Body")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/common.proto", fileDescriptor_a560f1e2899ebe7a)
}

var fileDescriptor_a560f1e2899ebe7a = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x14, 0x45, 0x77, 0x56, 0x59, 0x93, 0x59, 0x0b, 0x49, 0x21, 0x41, 0x5d, 0x08, 0x6a, 0x91, 0x6a,
	0x06, 0xb4, 0xb6, 0x19, 0x9b, 0x2d, 0x97, 0xf8, 0x01, 0xe1, 0x65, 0x27, 0x68, 0xc0, 0xcc, 0x1b,
	0xe6, 0xbd, 0x2c, 0x4e, 0xe7, 0x77, 0xf9, 0x05, 0xb6, 0xfe, 0x91, 0x64, 0x14, 0x6c, 0x6c, 0x0f,
	0xe7, 0xc2, 0xb9, 0xf2, 0xa6, 0x77, 0x07, 0x8c, 0xda, 0x02, 0x83, 0x66, 0xf0, 0xfa, 0x70, 0x07,
	0xaf, 0xfe, 0x05, 0xf4, 0x1e, 0xc7, 0x11, 0x9d, 0xf2, 0x01, 0x19, 0x8b, 0xf3, 0x24, 0xa9, 0x59,
	0x52, 0x0c, 0x5e, 0xfd, 0x4a, 0x17, 0x9b, 0xc9, 0x7a, 0xd0, 0xe0, 0x1c, 0x32, 0xf0, 0x80, 0x8e,
	0x34, 0x31, 0xf0, 0x44, 0x3f, 0xb3, 0xeb, 0x51, 0x1e, 0x1b, 0xb4, 0xb1, 0xb8, 0x94, 0x19, 0x50,
	0xdb, 0x45, 0xee, 0xa9, 0x14, 0x95, 0xa8, 0x4f, 0xb7, 0x8b, 0xe6, 0x04, 0xc8, 0xcc, 0xa0, 0xd8,
	0xc8, 0x1c, 0xa8, 0x25, 0x0e, 0x83, 0x7b, 0x2e, 0x97, 0x95, 0xa8, 0xf3, 0xed, 0xa2, 0xc9, 0x80,
	0x9e, 0x12, 0x29, 0xae, 0x64, 0xce, 0x61, 0x72, 0x7b, 0xe0, 0xde, 0x96, 0x47, 0x95, 0xa8, 0xb3,
	0xe6, 0x0f, 0x98, 0xb5, 0xcc, 0x3b, 0xb4, 0xb1, 0xe5, 0xe8, 0x7b, 0xf3, 0xf0, 0xf1, 0xfe, 0xf9,
	0xb5, 0x5a, 0x9e, 0x09, 0x79, 0x3b, 0xa0, 0x4a, 0xc9, 0x3e, 0xe0, 0x5b, 0x54, 0xff, 0xd7, 0x9b,
	0xf5, 0x63, 0xfa, 0xb8, 0x9b, 0x5b, 0x77, 0xa2, 0x5b, 0xa5, 0xe8, 0xfb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe5, 0xb6, 0x6b, 0xba, 0x12, 0x01, 0x00, 0x00,
}