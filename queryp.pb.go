// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: queryp.proto

package queryp

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type FilterLogic int32

const (
	FilterLogicAnd FilterLogic = 0
	FilterLogicOr  FilterLogic = 1
)

var FilterLogic_name = map[int32]string{
	0: "FilterLogicAnd",
	1: "FilterLogicOr",
}

var FilterLogic_value = map[string]int32{
	"FilterLogicAnd": 0,
	"FilterLogicOr":  1,
}

func (FilterLogic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b02c08be017f0785, []int{0}
}

type FilterOp int32

const (
	FilterOpEquals           FilterOp = 0
	FilterOpNotEquals        FilterOp = 1
	FilterOpLessThan         FilterOp = 2
	FilterOpLessThanEqual    FilterOp = 3
	FilterOpGreaterThan      FilterOp = 4
	FilterOpGreaterThanEqual FilterOp = 5
	FilterOpLike             FilterOp = 6
	FilterOpNotLike          FilterOp = 7
	FilterOpILike            FilterOp = 8
	FilterOpNotILike         FilterOp = 9
	FilterOpRegexp           FilterOp = 10
	FilterOpNotRegexp        FilterOp = 11
	FilterOpIRegexp          FilterOp = 12
	FilterOpNotIRegexp       FilterOp = 13
)

var FilterOp_name = map[int32]string{
	0:  "FilterOpEquals",
	1:  "FilterOpNotEquals",
	2:  "FilterOpLessThan",
	3:  "FilterOpLessThanEqual",
	4:  "FilterOpGreaterThan",
	5:  "FilterOpGreaterThanEqual",
	6:  "FilterOpLike",
	7:  "FilterOpNotLike",
	8:  "FilterOpILike",
	9:  "FilterOpNotILike",
	10: "FilterOpRegexp",
	11: "FilterOpNotRegexp",
	12: "FilterOpIRegexp",
	13: "FilterOpNotIRegexp",
}

var FilterOp_value = map[string]int32{
	"FilterOpEquals":           0,
	"FilterOpNotEquals":        1,
	"FilterOpLessThan":         2,
	"FilterOpLessThanEqual":    3,
	"FilterOpGreaterThan":      4,
	"FilterOpGreaterThanEqual": 5,
	"FilterOpLike":             6,
	"FilterOpNotLike":          7,
	"FilterOpILike":            8,
	"FilterOpNotILike":         9,
	"FilterOpRegexp":           10,
	"FilterOpNotRegexp":        11,
	"FilterOpIRegexp":          12,
	"FilterOpNotIRegexp":       13,
}

func (FilterOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b02c08be017f0785, []int{1}
}

type FilterType int32

const (
	FilterTypeNotFound FilterType = 0
	FilterTypeSimple   FilterType = 1
	FilterTypeString   FilterType = 2
	FilterTypeNumeric  FilterType = 3
	FilterTypeTime     FilterType = 4
	FilterTypeBool     FilterType = 5
)

var FilterType_name = map[int32]string{
	0: "FilterTypeNotFound",
	1: "FilterTypeSimple",
	2: "FilterTypeString",
	3: "FilterTypeNumeric",
	4: "FilterTypeTime",
	5: "FilterTypeBool",
}

var FilterType_value = map[string]int32{
	"FilterTypeNotFound": 0,
	"FilterTypeSimple":   1,
	"FilterTypeString":   2,
	"FilterTypeNumeric":  3,
	"FilterTypeTime":     4,
	"FilterTypeBool":     5,
}

func (FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b02c08be017f0785, []int{2}
}

type QueryParameters struct {
	Filter  []FilterTerm `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter"`
	Sort    []SortTerm   `protobuf:"bytes,2,rep,name=sort,proto3" json:"sort"`
	Options Options      `protobuf:"bytes,3,rep,name=options,proto3,casttype=Options" json:"options" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Limit   int64        `protobuf:"varint,4,opt,name=limit,proto3" json:"limit"`
	Offset  int64        `protobuf:"varint,5,opt,name=offset,proto3" json:"offset"`
}

func (m *QueryParameters) Reset()      { *m = QueryParameters{} }
func (*QueryParameters) ProtoMessage() {}
func (*QueryParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_b02c08be017f0785, []int{0}
}
func (m *QueryParameters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParameters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParameters.Merge(m, src)
}
func (m *QueryParameters) XXX_Size() int {
	return m.Size()
}
func (m *QueryParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParameters.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParameters proto.InternalMessageInfo

func (m *QueryParameters) GetFilter() []FilterTerm {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *QueryParameters) GetSort() []SortTerm {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *QueryParameters) GetOptions() Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *QueryParameters) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryParameters) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SortTerm struct {
	Field Field `protobuf:"bytes,1,opt,name=field,proto3,casttype=Field" json:"field"`
	Desc  bool  `protobuf:"varint,2,opt,name=desc,proto3" json:"desc"`
}

func (m *SortTerm) Reset()      { *m = SortTerm{} }
func (*SortTerm) ProtoMessage() {}
func (*SortTerm) Descriptor() ([]byte, []int) {
	return fileDescriptor_b02c08be017f0785, []int{1}
}
func (m *SortTerm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SortTerm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SortTerm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SortTerm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortTerm.Merge(m, src)
}
func (m *SortTerm) XXX_Size() int {
	return m.Size()
}
func (m *SortTerm) XXX_DiscardUnknown() {
	xxx_messageInfo_SortTerm.DiscardUnknown(m)
}

var xxx_messageInfo_SortTerm proto.InternalMessageInfo

func (m *SortTerm) GetField() Field {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *SortTerm) GetDesc() bool {
	if m != nil {
		return m.Desc
	}
	return false
}

type FilterTerm struct {
	Logic     FilterLogic  `protobuf:"varint,1,opt,name=logic,proto3,enum=queryp.FilterLogic" json:"logic"`
	Op        FilterOp     `protobuf:"varint,2,opt,name=op,proto3,enum=queryp.FilterOp" json:"op"`
	Field     Field        `protobuf:"bytes,3,opt,name=field,proto3,casttype=Field" json:"field"`
	Value     string       `protobuf:"bytes,4,opt,name=value,proto3" json:"value"`
	SubFilter []FilterTerm `protobuf:"bytes,10,rep,name=sub_filter,json=subFilter,proto3" json:"sub_filter,omitempty"`
}

func (m *FilterTerm) Reset()      { *m = FilterTerm{} }
func (*FilterTerm) ProtoMessage() {}
func (*FilterTerm) Descriptor() ([]byte, []int) {
	return fileDescriptor_b02c08be017f0785, []int{2}
}
func (m *FilterTerm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterTerm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterTerm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterTerm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterTerm.Merge(m, src)
}
func (m *FilterTerm) XXX_Size() int {
	return m.Size()
}
func (m *FilterTerm) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterTerm.DiscardUnknown(m)
}

var xxx_messageInfo_FilterTerm proto.InternalMessageInfo

func (m *FilterTerm) GetLogic() FilterLogic {
	if m != nil {
		return m.Logic
	}
	return FilterLogicAnd
}

func (m *FilterTerm) GetOp() FilterOp {
	if m != nil {
		return m.Op
	}
	return FilterOpEquals
}

func (m *FilterTerm) GetField() Field {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *FilterTerm) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FilterTerm) GetSubFilter() []FilterTerm {
	if m != nil {
		return m.SubFilter
	}
	return nil
}

func init() {
	proto.RegisterEnum("queryp.FilterLogic", FilterLogic_name, FilterLogic_value)
	proto.RegisterEnum("queryp.FilterOp", FilterOp_name, FilterOp_value)
	proto.RegisterEnum("queryp.FilterType", FilterType_name, FilterType_value)
	proto.RegisterType((*QueryParameters)(nil), "queryp.QueryParameters")
	proto.RegisterMapType((Options)(nil), "queryp.QueryParameters.OptionsEntry")
	proto.RegisterType((*SortTerm)(nil), "queryp.SortTerm")
	proto.RegisterType((*FilterTerm)(nil), "queryp.FilterTerm")
}

func init() { proto.RegisterFile("queryp.proto", fileDescriptor_b02c08be017f0785) }

var fileDescriptor_b02c08be017f0785 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x4f, 0xdb, 0x48,
	0x14, 0xc7, 0x6d, 0x27, 0x0e, 0xc9, 0x23, 0x04, 0x33, 0xfc, 0x58, 0x6f, 0x84, 0xc6, 0x51, 0xb4,
	0x87, 0x08, 0xed, 0x06, 0x89, 0xdd, 0xc3, 0x2a, 0x87, 0x4a, 0xb5, 0x04, 0x15, 0x12, 0x22, 0xd4,
	0xa0, 0x1e, 0x7a, 0xa9, 0x92, 0x30, 0x09, 0x16, 0x76, 0xc6, 0xd8, 0xe3, 0xb6, 0xe9, 0x89, 0x23,
	0xc7, 0x5e, 0x7b, 0x6e, 0x0f, 0x3d, 0xf6, 0xcf, 0xe0, 0x56, 0x8e, 0x1c, 0xaa, 0x08, 0xc2, 0xa5,
	0xca, 0x89, 0x33, 0xa7, 0xca, 0x33, 0x36, 0x36, 0xa8, 0x55, 0x2f, 0x9e, 0x79, 0x9f, 0x79, 0xdf,
	0x99, 0xf7, 0x9d, 0x79, 0x32, 0x94, 0x4f, 0x42, 0xe2, 0x8f, 0xbc, 0xa6, 0xe7, 0x53, 0x46, 0x51,
	0x41, 0x44, 0xd5, 0x7f, 0x06, 0x36, 0x3b, 0x0a, 0xbb, 0xcd, 0x1e, 0x75, 0xd7, 0x07, 0x74, 0x40,
	0xd7, 0xf9, 0x72, 0x37, 0xec, 0xf3, 0x88, 0x07, 0x7c, 0x26, 0x64, 0xf5, 0x2b, 0x05, 0xe6, 0x9f,
	0x47, 0xca, 0xbd, 0x8e, 0xdf, 0x71, 0x09, 0x23, 0x7e, 0x80, 0x5a, 0x50, 0xe8, 0xdb, 0x0e, 0x23,
	0xbe, 0x2e, 0xd7, 0x72, 0x8d, 0xd9, 0x0d, 0xd4, 0x8c, 0x4f, 0xda, 0xe2, 0xf4, 0x80, 0xf8, 0xae,
	0x59, 0x39, 0x1f, 0x1b, 0xd2, 0x74, 0x6c, 0xc4, 0x99, 0x56, 0x3c, 0xa2, 0x0d, 0xc8, 0x07, 0xd4,
	0x67, 0xba, 0xc2, 0x95, 0x5a, 0xa2, 0xdc, 0xa7, 0x3e, 0xe3, 0xba, 0x72, 0xac, 0xe3, 0x59, 0x16,
	0xff, 0xa2, 0x17, 0x30, 0x43, 0x3d, 0x66, 0xd3, 0x61, 0xa0, 0xe7, 0xb8, 0xec, 0xaf, 0x44, 0xf6,
	0xa8, 0xb2, 0x66, 0x5b, 0xa4, 0x6d, 0x0e, 0x99, 0x3f, 0x32, 0x57, 0xa6, 0x63, 0x23, 0x11, 0xde,
	0x8d, 0x8d, 0x99, 0x78, 0xd1, 0x4a, 0x18, 0x32, 0x40, 0x75, 0x6c, 0xd7, 0x66, 0x7a, 0xbe, 0x26,
	0x37, 0x72, 0x66, 0x69, 0x3a, 0x36, 0x04, 0xb0, 0xc4, 0x80, 0xea, 0x50, 0xa0, 0xfd, 0x7e, 0x40,
	0x98, 0xae, 0xf2, 0x0c, 0x88, 0x0c, 0x09, 0x62, 0xc5, 0x63, 0xb5, 0x05, 0xe5, 0xec, 0xa9, 0x48,
	0x83, 0xdc, 0x31, 0x19, 0xe9, 0x72, 0x4d, 0x6e, 0x94, 0xac, 0x68, 0x8a, 0x96, 0x40, 0x7d, 0xdd,
	0x71, 0x42, 0xa2, 0x2b, 0x9c, 0x89, 0xa0, 0xa5, 0xfc, 0x2f, 0xb7, 0xf2, 0xa7, 0xdf, 0x6a, 0x52,
	0xdd, 0x82, 0x62, 0x62, 0x1f, 0x35, 0x40, 0xed, 0xdb, 0xc4, 0x39, 0x14, 0x7a, 0x13, 0x45, 0x25,
	0x71, 0x70, 0x37, 0x36, 0xd4, 0xad, 0x68, 0x62, 0x89, 0x18, 0xad, 0x42, 0xfe, 0x90, 0x04, 0x3d,
	0xbe, 0x69, 0xd1, 0x2c, 0x46, 0x57, 0x16, 0xc5, 0x16, 0xff, 0xd6, 0xcf, 0x14, 0x80, 0xf4, 0x35,
	0xd0, 0x7f, 0xa0, 0x3a, 0x74, 0x60, 0xf7, 0xf8, 0xb6, 0x95, 0x8d, 0xc5, 0x87, 0x0f, 0xb6, 0x13,
	0x2d, 0xc5, 0xf6, 0xa3, 0xa9, 0x25, 0x06, 0xd4, 0x00, 0x85, 0x7a, 0xfc, 0x80, 0x4a, 0xfa, 0x52,
	0x42, 0xd2, 0xf6, 0xcc, 0xc2, 0x74, 0x6c, 0x28, 0xd4, 0xb3, 0x14, 0xea, 0xa5, 0x65, 0xe7, 0x7e,
	0x57, 0xb6, 0x91, 0x5c, 0x46, 0x9e, 0x67, 0xf2, 0x43, 0x39, 0x88, 0xef, 0x05, 0xb5, 0x01, 0x82,
	0xb0, 0xfb, 0x2a, 0x6e, 0x30, 0xf8, 0x65, 0x83, 0xad, 0xc6, 0x8d, 0xb2, 0x94, 0x66, 0xff, 0x4d,
	0x5d, 0x9b, 0x11, 0xd7, 0x63, 0x23, 0xab, 0x14, 0x84, 0x5d, 0x91, 0xbc, 0xf6, 0x04, 0x66, 0x33,
	0x36, 0x11, 0x82, 0x4a, 0x26, 0x7c, 0x3a, 0x3c, 0xd4, 0x24, 0xb4, 0x00, 0x73, 0x19, 0xd6, 0xf6,
	0x35, 0xb9, 0x5a, 0x3c, 0xfb, 0x88, 0xa5, 0x2f, 0x9f, 0xb0, 0xb4, 0xf6, 0x55, 0x81, 0x62, 0x62,
	0x3a, 0x55, 0xb7, 0xbd, 0xcd, 0x93, 0xb0, 0xe3, 0x04, 0x9a, 0x84, 0x96, 0x61, 0x21, 0x61, 0xbb,
	0x94, 0xc5, 0x58, 0x46, 0x4b, 0xa0, 0x25, 0x78, 0x87, 0x04, 0xc1, 0xc1, 0x51, 0x67, 0xa8, 0x29,
	0xe8, 0x4f, 0x58, 0x7e, 0x4c, 0xb9, 0x42, 0xcb, 0xa1, 0x3f, 0x60, 0x31, 0x59, 0x7a, 0xe6, 0x93,
	0x4e, 0x64, 0x34, 0xd2, 0xe4, 0xd1, 0x2a, 0xe8, 0x3f, 0x59, 0x10, 0x32, 0x15, 0x69, 0x50, 0xbe,
	0xdf, 0xd1, 0x3e, 0x26, 0x5a, 0x01, 0x2d, 0xc2, 0x7c, 0xa6, 0x20, 0x0e, 0x67, 0x52, 0x8f, 0x6d,
	0x6f, 0x9b, 0xa3, 0x62, 0xb6, 0xc2, 0x5d, 0xca, 0x04, 0x2d, 0x65, 0x2d, 0x5a, 0x64, 0x40, 0xde,
	0x7a, 0x1a, 0x3c, 0xb2, 0x18, 0xe3, 0xd9, 0xec, 0x41, 0xdb, 0x31, 0x2c, 0xa3, 0x15, 0x40, 0xd9,
	0x5d, 0x63, 0x3e, 0x97, 0xb9, 0xd1, 0x0f, 0xf2, 0x7d, 0x73, 0x8e, 0x3c, 0x92, 0x0a, 0xa2, 0x68,
	0x97, 0xb2, 0x2d, 0x1a, 0xf2, 0x57, 0xb9, 0x2f, 0x2f, 0xe2, 0xfb, 0xb6, 0xeb, 0x39, 0x24, 0x7b,
	0xad, 0x9c, 0x32, 0xdf, 0x1e, 0x0e, 0x34, 0x25, 0x2d, 0x90, 0xef, 0x11, 0xba, 0xc4, 0xb7, 0x7b,
	0x5a, 0x2e, 0xf5, 0x12, 0xe1, 0x03, 0xdb, 0x25, 0x5a, 0xfe, 0x21, 0x33, 0x29, 0x75, 0x34, 0x35,
	0xad, 0xcd, 0xdc, 0xbb, 0xb8, 0xc6, 0xd2, 0xe5, 0x35, 0x96, 0x6e, 0xaf, 0xb1, 0x7c, 0x3a, 0xc1,
	0xf2, 0xe7, 0x09, 0x96, 0xcf, 0x27, 0x58, 0xbe, 0x98, 0x60, 0xf9, 0x6a, 0x82, 0xe5, 0xef, 0x13,
	0x2c, 0xdd, 0x4e, 0xb0, 0xfc, 0xfe, 0x06, 0x4b, 0x17, 0x37, 0x58, 0xba, 0xbc, 0xc1, 0xd2, 0xcb,
	0x6a, 0xe6, 0xa7, 0x1a, 0x0c, 0xe9, 0x9b, 0x77, 0x9d, 0xde, 0xd1, 0xba, 0xe8, 0xd9, 0x6e, 0x81,
	0xff, 0x48, 0xff, 0xfd, 0x11, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xdb, 0xb4, 0x95, 0x8f, 0x05, 0x00,
	0x00,
}

func (this *QueryParameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryParameters)
	if !ok {
		that2, ok := that.(QueryParameters)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Filter) != len(that1.Filter) {
		return false
	}
	for i := range this.Filter {
		if !this.Filter[i].Equal(&that1.Filter[i]) {
			return false
		}
	}
	if len(this.Sort) != len(that1.Sort) {
		return false
	}
	for i := range this.Sort {
		if !this.Sort[i].Equal(&that1.Sort[i]) {
			return false
		}
	}
	if len(this.Options) != len(that1.Options) {
		return false
	}
	for i := range this.Options {
		if this.Options[i] != that1.Options[i] {
			return false
		}
	}
	if this.Limit != that1.Limit {
		return false
	}
	if this.Offset != that1.Offset {
		return false
	}
	return true
}
func (this *SortTerm) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SortTerm)
	if !ok {
		that2, ok := that.(SortTerm)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field != that1.Field {
		return false
	}
	if this.Desc != that1.Desc {
		return false
	}
	return true
}
func (this *FilterTerm) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FilterTerm)
	if !ok {
		that2, ok := that.(FilterTerm)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Logic != that1.Logic {
		return false
	}
	if this.Op != that1.Op {
		return false
	}
	if this.Field != that1.Field {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if len(this.SubFilter) != len(that1.SubFilter) {
		return false
	}
	for i := range this.SubFilter {
		if !this.SubFilter[i].Equal(&that1.SubFilter[i]) {
			return false
		}
	}
	return true
}
func (this *QueryParameters) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&queryp.QueryParameters{")
	if this.Filter != nil {
		vs := make([]FilterTerm, len(this.Filter))
		for i := range vs {
			vs[i] = this.Filter[i]
		}
		s = append(s, "Filter: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	if this.Sort != nil {
		vs := make([]SortTerm, len(this.Sort))
		for i := range vs {
			vs[i] = this.Sort[i]
		}
		s = append(s, "Sort: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	keysForOptions := make([]string, 0, len(this.Options))
	for k, _ := range this.Options {
		keysForOptions = append(keysForOptions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForOptions)
	mapStringForOptions := "Options{"
	for _, k := range keysForOptions {
		mapStringForOptions += fmt.Sprintf("%#v: %#v,", k, this.Options[k])
	}
	mapStringForOptions += "}"
	if this.Options != nil {
		s = append(s, "Options: "+mapStringForOptions+",\n")
	}
	s = append(s, "Limit: "+fmt.Sprintf("%#v", this.Limit)+",\n")
	s = append(s, "Offset: "+fmt.Sprintf("%#v", this.Offset)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SortTerm) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&queryp.SortTerm{")
	s = append(s, "Field: "+fmt.Sprintf("%#v", this.Field)+",\n")
	s = append(s, "Desc: "+fmt.Sprintf("%#v", this.Desc)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *FilterTerm) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&queryp.FilterTerm{")
	s = append(s, "Logic: "+fmt.Sprintf("%#v", this.Logic)+",\n")
	s = append(s, "Op: "+fmt.Sprintf("%#v", this.Op)+",\n")
	s = append(s, "Field: "+fmt.Sprintf("%#v", this.Field)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	if this.SubFilter != nil {
		vs := make([]FilterTerm, len(this.SubFilter))
		for i := range vs {
			vs[i] = this.SubFilter[i]
		}
		s = append(s, "SubFilter: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQueryp(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *QueryParameters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParameters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParameters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		i = encodeVarintQueryp(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x28
	}
	if m.Limit != 0 {
		i = encodeVarintQueryp(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Options) > 0 {
		for k := range m.Options {
			v := m.Options[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintQueryp(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQueryp(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQueryp(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Sort) > 0 {
		for iNdEx := len(m.Sort) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sort[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Filter) > 0 {
		for iNdEx := len(m.Filter) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Filter[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SortTerm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SortTerm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SortTerm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Desc {
		i--
		if m.Desc {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintQueryp(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FilterTerm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterTerm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterTerm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubFilter) > 0 {
		for iNdEx := len(m.SubFilter) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubFilter[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintQueryp(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintQueryp(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Op != 0 {
		i = encodeVarintQueryp(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x10
	}
	if m.Logic != 0 {
		i = encodeVarintQueryp(dAtA, i, uint64(m.Logic))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryp(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParameters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Filter) > 0 {
		for _, e := range m.Filter {
			l = e.Size()
			n += 1 + l + sovQueryp(uint64(l))
		}
	}
	if len(m.Sort) > 0 {
		for _, e := range m.Sort {
			l = e.Size()
			n += 1 + l + sovQueryp(uint64(l))
		}
	}
	if len(m.Options) > 0 {
		for k, v := range m.Options {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQueryp(uint64(len(k))) + 1 + len(v) + sovQueryp(uint64(len(v)))
			n += mapEntrySize + 1 + sovQueryp(uint64(mapEntrySize))
		}
	}
	if m.Limit != 0 {
		n += 1 + sovQueryp(uint64(m.Limit))
	}
	if m.Offset != 0 {
		n += 1 + sovQueryp(uint64(m.Offset))
	}
	return n
}

func (m *SortTerm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovQueryp(uint64(l))
	}
	if m.Desc {
		n += 2
	}
	return n
}

func (m *FilterTerm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Logic != 0 {
		n += 1 + sovQueryp(uint64(m.Logic))
	}
	if m.Op != 0 {
		n += 1 + sovQueryp(uint64(m.Op))
	}
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovQueryp(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovQueryp(uint64(l))
	}
	if len(m.SubFilter) > 0 {
		for _, e := range m.SubFilter {
			l = e.Size()
			n += 1 + l + sovQueryp(uint64(l))
		}
	}
	return n
}

func sovQueryp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryp(x uint64) (n int) {
	return sovQueryp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SortTerm) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SortTerm{`,
		`Field:` + fmt.Sprintf("%v", this.Field) + `,`,
		`Desc:` + fmt.Sprintf("%v", this.Desc) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FilterTerm) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForSubFilter := "[]FilterTerm{"
	for _, f := range this.SubFilter {
		repeatedStringForSubFilter += strings.Replace(strings.Replace(f.String(), "FilterTerm", "FilterTerm", 1), `&`, ``, 1) + ","
	}
	repeatedStringForSubFilter += "}"
	s := strings.Join([]string{`&FilterTerm{`,
		`Logic:` + fmt.Sprintf("%v", this.Logic) + `,`,
		`Op:` + fmt.Sprintf("%v", this.Op) + `,`,
		`Field:` + fmt.Sprintf("%v", this.Field) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`SubFilter:` + repeatedStringForSubFilter + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQueryp(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueryParameters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filter = append(m.Filter, FilterTerm{})
			if err := m.Filter[len(m.Filter)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sort = append(m.Sort, SortTerm{})
			if err := m.Sort[len(m.Sort)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = make(Options)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQueryp
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQueryp
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQueryp
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQueryp
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQueryp
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthQueryp
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthQueryp
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQueryp(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQueryp
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Options[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SortTerm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SortTerm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SortTerm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = Field(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Desc = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQueryp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FilterTerm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FilterTerm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterTerm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logic", wireType)
			}
			m.Logic = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Logic |= FilterLogic(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= FilterOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = Field(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubFilter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubFilter = append(m.SubFilter, FilterTerm{})
			if err := m.SubFilter[len(m.SubFilter)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQueryp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQueryp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryp = fmt.Errorf("proto: unexpected end of group")
)
