// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tags.proto

package tags

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Outside struct {
	*Inside `protobuf:"bytes,1,opt,name=Inside,embedded=Inside" json:""`
	Field2  *string `protobuf:"bytes,2,opt,name=Field2" json:"MyField2" xml:",comment"`
	// Types that are valid to be assigned to Filed:
	//	*Outside_Field3
	Filed                isOutside_Filed `protobuf_oneof:"filed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Outside) Reset()         { *m = Outside{} }
func (m *Outside) String() string { return proto.CompactTextString(m) }
func (*Outside) ProtoMessage()    {}
func (*Outside) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d9cbcae1e528f6, []int{0}
}
func (m *Outside) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outside.Unmarshal(m, b)
}
func (m *Outside) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outside.Marshal(b, m, deterministic)
}
func (m *Outside) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outside.Merge(m, src)
}
func (m *Outside) XXX_Size() int {
	return xxx_messageInfo_Outside.Size(m)
}
func (m *Outside) XXX_DiscardUnknown() {
	xxx_messageInfo_Outside.DiscardUnknown(m)
}

var xxx_messageInfo_Outside proto.InternalMessageInfo

type isOutside_Filed interface {
	isOutside_Filed()
}

type Outside_Field3 struct {
	Field3 string `protobuf:"bytes,3,opt,name=Field3,oneof" json:"MyField3" xml:",comment"`
}

func (*Outside_Field3) isOutside_Filed() {}

func (m *Outside) GetFiled() isOutside_Filed {
	if m != nil {
		return m.Filed
	}
	return nil
}

func (m *Outside) GetField2() string {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return ""
}

func (m *Outside) GetField3() string {
	if x, ok := m.GetFiled().(*Outside_Field3); ok {
		return x.Field3
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Outside) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Outside_Field3)(nil),
	}
}

type Inside struct {
	Field1               *string  `protobuf:"bytes,1,opt,name=Field1" json:"MyField1" xml:",chardata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inside) Reset()         { *m = Inside{} }
func (m *Inside) String() string { return proto.CompactTextString(m) }
func (*Inside) ProtoMessage()    {}
func (*Inside) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d9cbcae1e528f6, []int{1}
}
func (m *Inside) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inside.Unmarshal(m, b)
}
func (m *Inside) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inside.Marshal(b, m, deterministic)
}
func (m *Inside) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inside.Merge(m, src)
}
func (m *Inside) XXX_Size() int {
	return xxx_messageInfo_Inside.Size(m)
}
func (m *Inside) XXX_DiscardUnknown() {
	xxx_messageInfo_Inside.DiscardUnknown(m)
}

var xxx_messageInfo_Inside proto.InternalMessageInfo

func (m *Inside) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

func init() {
	proto.RegisterType((*Outside)(nil), "tags.Outside")
	proto.RegisterType((*Inside)(nil), "tags.Inside")
}

func init() { proto.RegisterFile("tags.proto", fileDescriptor_e7d9cbcae1e528f6) }

var fileDescriptor_e7d9cbcae1e528f6 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x4c, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x92, 0x49, 0xa5,
	0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x34, 0x29, 0x6d, 0x61, 0xe4, 0x62, 0xf7, 0x2f, 0x2d,
	0x29, 0xce, 0x4c, 0x49, 0x15, 0xd2, 0xe3, 0x62, 0xf3, 0xcc, 0x03, 0xb1, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x78, 0xf4, 0xc0, 0xa6, 0x43, 0xc4, 0x9c, 0x38, 0x2e, 0xdc, 0x93, 0x67, 0x7c,
	0x75, 0x4f, 0x9e, 0x21, 0x08, 0xaa, 0x4a, 0xc8, 0x8c, 0x8b, 0xcd, 0x2d, 0x33, 0x35, 0x27, 0xc5,
	0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0xd3, 0x49, 0xee, 0xd5, 0x3d, 0x79, 0x0e, 0xdf, 0x4a, 0x88,
	0xd8, 0xa7, 0x7b, 0xf2, 0x7c, 0x15, 0xb9, 0x39, 0x56, 0x4a, 0x3a, 0xc9, 0xf9, 0xb9, 0xb9, 0xa9,
	0x79, 0x25, 0x4a, 0x41, 0x50, 0xd5, 0x42, 0x16, 0x50, 0x7d, 0xc6, 0x12, 0xcc, 0x18, 0xfa, 0x8c,
	0x31, 0xf5, 0x79, 0x30, 0x40, 0x75, 0x1a, 0x3b, 0xb1, 0x73, 0xb1, 0xa6, 0x65, 0xe6, 0xa4, 0xa6,
	0x28, 0x39, 0xc2, 0x9c, 0x2a, 0x64, 0x0e, 0x35, 0xcc, 0x10, 0xec, 0x68, 0x4e, 0x27, 0x79, 0x24,
	0xc3, 0x0c, 0x3f, 0xdd, 0x93, 0xe7, 0x87, 0x1a, 0x96, 0x91, 0x58, 0x94, 0x92, 0x58, 0x92, 0x08,
	0x73, 0x85, 0xa1, 0x13, 0xcb, 0x8f, 0x87, 0x72, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23,
	0x9e, 0x05, 0xb9, 0x41, 0x01, 0x00, 0x00,
}

func NewPopulatedOutside(r randyTags, easy bool) *Outside {
	this := &Outside{}
	if r.Intn(5) != 0 {
		this.Inside = NewPopulatedInside(r, easy)
	}
	if r.Intn(5) != 0 {
		v1 := string(randStringTags(r))
		this.Field2 = &v1
	}
	oneofNumber_Filed := []int32{3}[r.Intn(1)]
	switch oneofNumber_Filed {
	case 3:
		this.Filed = NewPopulatedOutside_Field3(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 4)
	}
	return this
}

func NewPopulatedOutside_Field3(r randyTags, easy bool) *Outside_Field3 {
	this := &Outside_Field3{}
	this.Field3 = string(randStringTags(r))
	return this
}
func NewPopulatedInside(r randyTags, easy bool) *Inside {
	this := &Inside{}
	if r.Intn(5) != 0 {
		v2 := string(randStringTags(r))
		this.Field1 = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 2)
	}
	return this
}

type randyTags interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTags(r randyTags) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTags(r randyTags) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTags(r)
	}
	return string(tmps)
}
func randUnrecognizedTags(r randyTags, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTags(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTags(dAtA []byte, r randyTags, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTags(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTags(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTags(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
