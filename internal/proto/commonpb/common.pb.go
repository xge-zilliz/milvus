// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package commonpb

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

type ErrorCode int32

const (
	ErrorCode_Success               ErrorCode = 0
	ErrorCode_UnexpectedError       ErrorCode = 1
	ErrorCode_ConnectFailed         ErrorCode = 2
	ErrorCode_PermissionDenied      ErrorCode = 3
	ErrorCode_CollectionNotExists   ErrorCode = 4
	ErrorCode_IllegalArgument       ErrorCode = 5
	ErrorCode_IllegalDimension      ErrorCode = 7
	ErrorCode_IllegalIndexType      ErrorCode = 8
	ErrorCode_IllegalCollectionName ErrorCode = 9
	ErrorCode_IllegalTOPK           ErrorCode = 10
	ErrorCode_IllegalRowRecord      ErrorCode = 11
	ErrorCode_IllegalVectorID       ErrorCode = 12
	ErrorCode_IllegalSearchResult   ErrorCode = 13
	ErrorCode_FileNotFound          ErrorCode = 14
	ErrorCode_MetaFailed            ErrorCode = 15
	ErrorCode_CacheFailed           ErrorCode = 16
	ErrorCode_CannotCreateFolder    ErrorCode = 17
	ErrorCode_CannotCreateFile      ErrorCode = 18
	ErrorCode_CannotDeleteFolder    ErrorCode = 19
	ErrorCode_CannotDeleteFile      ErrorCode = 20
	ErrorCode_BuildIndexError       ErrorCode = 21
	ErrorCode_IllegalNLIST          ErrorCode = 22
	ErrorCode_IllegalMetricType     ErrorCode = 23
	ErrorCode_OutOfMemory           ErrorCode = 24
	ErrorCode_IndexNotExist         ErrorCode = 25
	// internal error code.
	ErrorCode_DDRequestRace ErrorCode = 1000
)

var ErrorCode_name = map[int32]string{
	0:    "Success",
	1:    "UnexpectedError",
	2:    "ConnectFailed",
	3:    "PermissionDenied",
	4:    "CollectionNotExists",
	5:    "IllegalArgument",
	7:    "IllegalDimension",
	8:    "IllegalIndexType",
	9:    "IllegalCollectionName",
	10:   "IllegalTOPK",
	11:   "IllegalRowRecord",
	12:   "IllegalVectorID",
	13:   "IllegalSearchResult",
	14:   "FileNotFound",
	15:   "MetaFailed",
	16:   "CacheFailed",
	17:   "CannotCreateFolder",
	18:   "CannotCreateFile",
	19:   "CannotDeleteFolder",
	20:   "CannotDeleteFile",
	21:   "BuildIndexError",
	22:   "IllegalNLIST",
	23:   "IllegalMetricType",
	24:   "OutOfMemory",
	25:   "IndexNotExist",
	1000: "DDRequestRace",
}

var ErrorCode_value = map[string]int32{
	"Success":               0,
	"UnexpectedError":       1,
	"ConnectFailed":         2,
	"PermissionDenied":      3,
	"CollectionNotExists":   4,
	"IllegalArgument":       5,
	"IllegalDimension":      7,
	"IllegalIndexType":      8,
	"IllegalCollectionName": 9,
	"IllegalTOPK":           10,
	"IllegalRowRecord":      11,
	"IllegalVectorID":       12,
	"IllegalSearchResult":   13,
	"FileNotFound":          14,
	"MetaFailed":            15,
	"CacheFailed":           16,
	"CannotCreateFolder":    17,
	"CannotCreateFile":      18,
	"CannotDeleteFolder":    19,
	"CannotDeleteFile":      20,
	"BuildIndexError":       21,
	"IllegalNLIST":          22,
	"IllegalMetricType":     23,
	"OutOfMemory":           24,
	"IndexNotExist":         25,
	"DDRequestRace":         1000,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type IndexState int32

const (
	IndexState_IndexStateNone IndexState = 0
	IndexState_Unissued       IndexState = 1
	IndexState_InProgress     IndexState = 2
	IndexState_Finished       IndexState = 3
	IndexState_Failed         IndexState = 4
)

var IndexState_name = map[int32]string{
	0: "IndexStateNone",
	1: "Unissued",
	2: "InProgress",
	3: "Finished",
	4: "Failed",
}

var IndexState_value = map[string]int32{
	"IndexStateNone": 0,
	"Unissued":       1,
	"InProgress":     2,
	"Finished":       3,
	"Failed":         4,
}

func (x IndexState) String() string {
	return proto.EnumName(IndexState_name, int32(x))
}

func (IndexState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type SegmentState int32

const (
	SegmentState_SegmentStateNone SegmentState = 0
	SegmentState_NotExist         SegmentState = 1
	SegmentState_Growing          SegmentState = 2
	SegmentState_Sealed           SegmentState = 3
	SegmentState_Flushed          SegmentState = 4
)

var SegmentState_name = map[int32]string{
	0: "SegmentStateNone",
	1: "NotExist",
	2: "Growing",
	3: "Sealed",
	4: "Flushed",
}

var SegmentState_value = map[string]int32{
	"SegmentStateNone": 0,
	"NotExist":         1,
	"Growing":          2,
	"Sealed":           3,
	"Flushed":          4,
}

func (x SegmentState) String() string {
	return proto.EnumName(SegmentState_name, int32(x))
}

func (SegmentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type MsgType int32

const (
	MsgType_Undefined MsgType = 0
	// DEFINITION REQUESTS: COLLECTION
	MsgType_CreateCollection   MsgType = 100
	MsgType_DropCollection     MsgType = 101
	MsgType_HasCollection      MsgType = 102
	MsgType_DescribeCollection MsgType = 103
	MsgType_ShowCollections    MsgType = 104
	MsgType_GetSystemConfigs   MsgType = 105
	MsgType_LoadCollection     MsgType = 106
	MsgType_ReleaseCollection  MsgType = 107
	// DEFINITION REQUESTS: PARTITION
	MsgType_CreatePartition   MsgType = 200
	MsgType_DropPartition     MsgType = 201
	MsgType_HasPartition      MsgType = 202
	MsgType_DescribePartition MsgType = 203
	MsgType_ShowPartitions    MsgType = 204
	MsgType_LoadPartitions    MsgType = 205
	MsgType_ReleasePartitions MsgType = 206
	// DEFINE REQUESTS: SEGMENT
	MsgType_ShowSegments    MsgType = 250
	MsgType_DescribeSegment MsgType = 251
	// DEFINITION REQUESTS: INDEX
	MsgType_CreateIndex   MsgType = 300
	MsgType_DescribeIndex MsgType = 301
	MsgType_DropIndex     MsgType = 302
	// MANIPULATION REQUESTS
	MsgType_Insert MsgType = 400
	MsgType_Delete MsgType = 401
	MsgType_Flush  MsgType = 402
	// QUERY
	MsgType_Search                  MsgType = 500
	MsgType_SearchResult            MsgType = 501
	MsgType_GetIndexState           MsgType = 502
	MsgType_GetCollectionStatistics MsgType = 503
	MsgType_GetPartitionStatistics  MsgType = 504
	// DATA SERVICE
	MsgType_SegmentInfo MsgType = 600
	// SYSTEM CONTROL
	MsgType_TimeTick          MsgType = 1200
	MsgType_QueryNodeStats    MsgType = 1201
	MsgType_LoadIndex         MsgType = 1202
	MsgType_RequestID         MsgType = 1203
	MsgType_RequestTSO        MsgType = 1204
	MsgType_AllocateSegment   MsgType = 1205
	MsgType_SegmentStatistics MsgType = 1206
	MsgType_SegmentFlushDone  MsgType = 1207
)

var MsgType_name = map[int32]string{
	0:    "Undefined",
	100:  "CreateCollection",
	101:  "DropCollection",
	102:  "HasCollection",
	103:  "DescribeCollection",
	104:  "ShowCollections",
	105:  "GetSystemConfigs",
	106:  "LoadCollection",
	107:  "ReleaseCollection",
	200:  "CreatePartition",
	201:  "DropPartition",
	202:  "HasPartition",
	203:  "DescribePartition",
	204:  "ShowPartitions",
	205:  "LoadPartitions",
	206:  "ReleasePartitions",
	250:  "ShowSegments",
	251:  "DescribeSegment",
	300:  "CreateIndex",
	301:  "DescribeIndex",
	302:  "DropIndex",
	400:  "Insert",
	401:  "Delete",
	402:  "Flush",
	500:  "Search",
	501:  "SearchResult",
	502:  "GetIndexState",
	503:  "GetCollectionStatistics",
	504:  "GetPartitionStatistics",
	600:  "SegmentInfo",
	1200: "TimeTick",
	1201: "QueryNodeStats",
	1202: "LoadIndex",
	1203: "RequestID",
	1204: "RequestTSO",
	1205: "AllocateSegment",
	1206: "SegmentStatistics",
	1207: "SegmentFlushDone",
}

var MsgType_value = map[string]int32{
	"Undefined":               0,
	"CreateCollection":        100,
	"DropCollection":          101,
	"HasCollection":           102,
	"DescribeCollection":      103,
	"ShowCollections":         104,
	"GetSystemConfigs":        105,
	"LoadCollection":          106,
	"ReleaseCollection":       107,
	"CreatePartition":         200,
	"DropPartition":           201,
	"HasPartition":            202,
	"DescribePartition":       203,
	"ShowPartitions":          204,
	"LoadPartitions":          205,
	"ReleasePartitions":       206,
	"ShowSegments":            250,
	"DescribeSegment":         251,
	"CreateIndex":             300,
	"DescribeIndex":           301,
	"DropIndex":               302,
	"Insert":                  400,
	"Delete":                  401,
	"Flush":                   402,
	"Search":                  500,
	"SearchResult":            501,
	"GetIndexState":           502,
	"GetCollectionStatistics": 503,
	"GetPartitionStatistics":  504,
	"SegmentInfo":             600,
	"TimeTick":                1200,
	"QueryNodeStats":          1201,
	"LoadIndex":               1202,
	"RequestID":               1203,
	"RequestTSO":              1204,
	"AllocateSegment":         1205,
	"SegmentStatistics":       1206,
	"SegmentFlushDone":        1207,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

type Status struct {
	ErrorCode            ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=milvus.proto.common.ErrorCode" json:"error_code,omitempty"`
	Reason               string    `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_Success
}

func (m *Status) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type KeyValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Blob struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blob) Reset()         { *m = Blob{} }
func (m *Blob) String() string { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()    {}
func (*Blob) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *Blob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blob.Unmarshal(m, b)
}
func (m *Blob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blob.Marshal(b, m, deterministic)
}
func (m *Blob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blob.Merge(m, src)
}
func (m *Blob) XXX_Size() int {
	return xxx_messageInfo_Blob.Size(m)
}
func (m *Blob) XXX_DiscardUnknown() {
	xxx_messageInfo_Blob.DiscardUnknown(m)
}

var xxx_messageInfo_Blob proto.InternalMessageInfo

func (m *Blob) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Address struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Address) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type MsgBase struct {
	MsgType              MsgType  `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=milvus.proto.common.MsgType" json:"msg_type,omitempty"`
	MsgID                int64    `protobuf:"varint,2,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Timestamp            uint64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SourceID             int64    `protobuf:"varint,4,opt,name=sourceID,proto3" json:"sourceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBase) Reset()         { *m = MsgBase{} }
func (m *MsgBase) String() string { return proto.CompactTextString(m) }
func (*MsgBase) ProtoMessage()    {}
func (*MsgBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *MsgBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBase.Unmarshal(m, b)
}
func (m *MsgBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBase.Marshal(b, m, deterministic)
}
func (m *MsgBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBase.Merge(m, src)
}
func (m *MsgBase) XXX_Size() int {
	return xxx_messageInfo_MsgBase.Size(m)
}
func (m *MsgBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBase proto.InternalMessageInfo

func (m *MsgBase) GetMsgType() MsgType {
	if m != nil {
		return m.MsgType
	}
	return MsgType_Undefined
}

func (m *MsgBase) GetMsgID() int64 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *MsgBase) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MsgBase) GetSourceID() int64 {
	if m != nil {
		return m.SourceID
	}
	return 0
}

// Don't Modify This. @czs
type MsgHeader struct {
	Base                 *MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgHeader) Reset()         { *m = MsgHeader{} }
func (m *MsgHeader) String() string { return proto.CompactTextString(m) }
func (*MsgHeader) ProtoMessage()    {}
func (*MsgHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *MsgHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgHeader.Unmarshal(m, b)
}
func (m *MsgHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgHeader.Marshal(b, m, deterministic)
}
func (m *MsgHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHeader.Merge(m, src)
}
func (m *MsgHeader) XXX_Size() int {
	return xxx_messageInfo_MsgHeader.Size(m)
}
func (m *MsgHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHeader proto.InternalMessageInfo

func (m *MsgHeader) GetBase() *MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func init() {
	proto.RegisterEnum("milvus.proto.common.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("milvus.proto.common.IndexState", IndexState_name, IndexState_value)
	proto.RegisterEnum("milvus.proto.common.SegmentState", SegmentState_name, SegmentState_value)
	proto.RegisterEnum("milvus.proto.common.MsgType", MsgType_name, MsgType_value)
	proto.RegisterType((*Status)(nil), "milvus.proto.common.Status")
	proto.RegisterType((*KeyValuePair)(nil), "milvus.proto.common.KeyValuePair")
	proto.RegisterType((*Blob)(nil), "milvus.proto.common.Blob")
	proto.RegisterType((*Address)(nil), "milvus.proto.common.Address")
	proto.RegisterType((*MsgBase)(nil), "milvus.proto.common.MsgBase")
	proto.RegisterType((*MsgHeader)(nil), "milvus.proto.common.MsgHeader")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x25, 0xd9, 0x32, 0x47, 0xb2, 0xbc, 0x5e, 0xff, 0xc4, 0x69, 0x8d, 0x22, 0xf0, 0x29,
	0x30, 0x10, 0xbb, 0x68, 0xd0, 0xf6, 0x94, 0x43, 0x2c, 0xc5, 0x8e, 0x90, 0xf8, 0xa7, 0x94, 0x13,
	0x04, 0xbd, 0x04, 0x34, 0x39, 0xa6, 0xb7, 0x21, 0x77, 0xd5, 0xdd, 0x65, 0x12, 0xbd, 0x45, 0x9b,
	0x77, 0xe8, 0xad, 0x2d, 0xfa, 0xdf, 0x57, 0xe8, 0xff, 0xb9, 0x8f, 0xd0, 0x07, 0xe8, 0x1f, 0xd2,
	0x4b, 0x31, 0x4b, 0x4a, 0x64, 0x81, 0xf4, 0xb6, 0xf3, 0xcd, 0xec, 0x37, 0xdf, 0xcc, 0xec, 0x90,
	0xd0, 0x8d, 0x54, 0x96, 0x29, 0xb9, 0x33, 0xd6, 0xca, 0x2a, 0xbe, 0x92, 0x89, 0xf4, 0x49, 0x6e,
	0x0a, 0x6b, 0xa7, 0x70, 0x6d, 0x3d, 0x82, 0xf9, 0x91, 0x0d, 0x6d, 0x6e, 0xf8, 0x4d, 0x00, 0xd4,
	0x5a, 0xe9, 0x47, 0x91, 0x8a, 0x71, 0xc3, 0xbb, 0xea, 0x5d, 0xeb, 0xbd, 0xf1, 0xda, 0xce, 0x4b,
	0xee, 0xec, 0xdc, 0xa6, 0xb0, 0xbe, 0x8a, 0x31, 0xf0, 0x71, 0x7a, 0xe4, 0xeb, 0x30, 0xaf, 0x31,
	0x34, 0x4a, 0x6e, 0x34, 0xae, 0x7a, 0xd7, 0xfc, 0xa0, 0xb4, 0xb6, 0xde, 0x82, 0xee, 0x5d, 0x9c,
	0x3c, 0x08, 0xd3, 0x1c, 0x4f, 0x42, 0xa1, 0x39, 0x83, 0xe6, 0x63, 0x9c, 0x38, 0x7e, 0x3f, 0xa0,
	0x23, 0x5f, 0x85, 0xb9, 0x27, 0xe4, 0x2e, 0x2f, 0x16, 0xc6, 0xd6, 0x26, 0xb4, 0xf6, 0x52, 0x75,
	0x56, 0x79, 0xe9, 0x46, 0x77, 0xea, 0xbd, 0x0e, 0xed, 0x5b, 0x71, 0xac, 0xd1, 0x18, 0xde, 0x83,
	0x86, 0x18, 0x97, 0x7c, 0x0d, 0x31, 0xe6, 0x1c, 0x5a, 0x63, 0xa5, 0xad, 0x63, 0x6b, 0x06, 0xee,
	0xbc, 0xf5, 0xdc, 0x83, 0xf6, 0xa1, 0x49, 0xf6, 0x42, 0x83, 0xfc, 0x6d, 0x58, 0xc8, 0x4c, 0xf2,
	0xc8, 0x4e, 0xc6, 0xd3, 0x2a, 0x37, 0x5f, 0x5a, 0xe5, 0xa1, 0x49, 0x4e, 0x27, 0x63, 0x0c, 0xda,
	0x59, 0x71, 0x20, 0x25, 0x99, 0x49, 0x86, 0x83, 0x92, 0xb9, 0x30, 0xf8, 0x26, 0xf8, 0x56, 0x64,
	0x68, 0x6c, 0x98, 0x8d, 0x37, 0x9a, 0x57, 0xbd, 0x6b, 0xad, 0xa0, 0x02, 0xf8, 0x2b, 0xb0, 0x60,
	0x54, 0xae, 0x23, 0x1c, 0x0e, 0x36, 0x5a, 0xee, 0xda, 0xcc, 0xde, 0xba, 0x09, 0xfe, 0xa1, 0x49,
	0xee, 0x60, 0x18, 0xa3, 0xe6, 0xaf, 0x43, 0xeb, 0x2c, 0x34, 0x85, 0xa2, 0xce, 0xff, 0x2b, 0xa2,
	0x0a, 0x02, 0x17, 0xb9, 0xfd, 0x51, 0x0b, 0xfc, 0xd9, 0x24, 0x78, 0x07, 0xda, 0xa3, 0x3c, 0x8a,
	0xd0, 0x18, 0x76, 0x89, 0xaf, 0xc0, 0xd2, 0x7d, 0x89, 0xcf, 0xc6, 0x18, 0x59, 0x8c, 0x5d, 0x0c,
	0xf3, 0xf8, 0x32, 0x2c, 0xf6, 0x95, 0x94, 0x18, 0xd9, 0xfd, 0x50, 0xa4, 0x18, 0xb3, 0x06, 0x5f,
	0x05, 0x76, 0x82, 0x3a, 0x13, 0xc6, 0x08, 0x25, 0x07, 0x28, 0x05, 0xc6, 0xac, 0xc9, 0x2f, 0xc3,
	0x4a, 0x5f, 0xa5, 0x29, 0x46, 0x56, 0x28, 0x79, 0xa4, 0xec, 0xed, 0x67, 0xc2, 0x58, 0xc3, 0x5a,
	0x44, 0x3b, 0x4c, 0x53, 0x4c, 0xc2, 0xf4, 0x96, 0x4e, 0xf2, 0x0c, 0xa5, 0x65, 0x73, 0xc4, 0x51,
	0x82, 0x03, 0x91, 0xa1, 0x24, 0x26, 0xd6, 0xae, 0xa1, 0x43, 0x19, 0xe3, 0x33, 0xea, 0x1f, 0x5b,
	0xe0, 0x57, 0x60, 0xad, 0x44, 0x6b, 0x09, 0xc2, 0x0c, 0x99, 0xcf, 0x97, 0xa0, 0x53, 0xba, 0x4e,
	0x8f, 0x4f, 0xee, 0x32, 0xa8, 0x31, 0x04, 0xea, 0x69, 0x80, 0x91, 0xd2, 0x31, 0xeb, 0xd4, 0x24,
	0x3c, 0xc0, 0xc8, 0x2a, 0x3d, 0x1c, 0xb0, 0x2e, 0x09, 0x2e, 0xc1, 0x11, 0x86, 0x3a, 0xba, 0x08,
	0xd0, 0xe4, 0xa9, 0x65, 0x8b, 0x9c, 0x41, 0x77, 0x5f, 0xa4, 0x78, 0xa4, 0xec, 0xbe, 0xca, 0x65,
	0xcc, 0x7a, 0xbc, 0x07, 0x70, 0x88, 0x36, 0x2c, 0x3b, 0xb0, 0x44, 0x69, 0xfb, 0x61, 0x74, 0x81,
	0x25, 0xc0, 0xf8, 0x3a, 0xf0, 0x7e, 0x28, 0xa5, 0xb2, 0x7d, 0x8d, 0xa1, 0xc5, 0x7d, 0x95, 0xc6,
	0xa8, 0xd9, 0x32, 0xc9, 0xf9, 0x0f, 0x2e, 0x52, 0x64, 0xbc, 0x8a, 0x1e, 0x60, 0x8a, 0xb3, 0xe8,
	0x95, 0x2a, 0xba, 0xc4, 0x29, 0x7a, 0x95, 0xc4, 0xef, 0xe5, 0x22, 0x8d, 0x5d, 0x4b, 0x8a, 0xb1,
	0xac, 0x91, 0xc6, 0x52, 0xfc, 0xd1, 0xbd, 0xe1, 0xe8, 0x94, 0xad, 0xf3, 0x35, 0x58, 0x2e, 0x91,
	0x43, 0xb4, 0x5a, 0x44, 0xae, 0x79, 0x97, 0x49, 0xea, 0x71, 0x6e, 0x8f, 0xcf, 0x0f, 0x31, 0x53,
	0x7a, 0xc2, 0x36, 0x68, 0xa0, 0x8e, 0x69, 0x3a, 0x22, 0x76, 0x85, 0x73, 0x58, 0x1c, 0x0c, 0x02,
	0x7c, 0x3f, 0x47, 0x63, 0x83, 0x30, 0x42, 0xf6, 0x5b, 0x7b, 0xfb, 0x21, 0x80, 0x0b, 0xa3, 0x35,
	0x47, 0xce, 0xa1, 0x57, 0x59, 0x47, 0x4a, 0x22, 0xbb, 0xc4, 0xbb, 0xb0, 0x70, 0x5f, 0x0a, 0x63,
	0x72, 0x8c, 0x99, 0x47, 0x2d, 0x1a, 0xca, 0x13, 0xad, 0x12, 0xda, 0x2e, 0xd6, 0x20, 0xef, 0xbe,
	0x90, 0xc2, 0x5c, 0xb8, 0xc7, 0x01, 0x30, 0x5f, 0xf6, 0xaa, 0xb5, 0xfd, 0x10, 0xba, 0x23, 0x4c,
	0xe8, 0x1d, 0x14, 0xdc, 0xab, 0xc0, 0xea, 0x76, 0xc5, 0x3e, 0x53, 0xe8, 0xd1, 0x3b, 0x3d, 0xd0,
	0xea, 0xa9, 0x90, 0x09, 0x6b, 0x10, 0xd9, 0x08, 0xc3, 0xd4, 0x11, 0x77, 0xa0, 0xbd, 0x9f, 0xe6,
	0x2e, 0x4b, 0x6b, 0xfb, 0xc5, 0x9c, 0xdb, 0x57, 0xb7, 0x76, 0x8b, 0xe0, 0xdf, 0x97, 0x31, 0x9e,
	0x0b, 0x89, 0x31, 0xbb, 0xe4, 0x5a, 0xeb, 0x46, 0x50, 0x3d, 0x21, 0x16, 0x53, 0x59, 0x03, 0xad,
	0xc6, 0x35, 0x0c, 0xa9, 0x3f, 0x77, 0x42, 0x53, 0x83, 0xce, 0x69, 0x5e, 0x03, 0x34, 0x91, 0x16,
	0x67, 0xf5, 0xeb, 0x09, 0x4d, 0x66, 0x74, 0xa1, 0x9e, 0x56, 0x98, 0x61, 0x17, 0x94, 0xe9, 0x00,
	0xed, 0x68, 0x62, 0x2c, 0x66, 0x7d, 0x25, 0xcf, 0x45, 0x62, 0x98, 0xa0, 0x4c, 0xf7, 0x54, 0x18,
	0xd7, 0xae, 0xbf, 0x47, 0x13, 0x0b, 0x30, 0xc5, 0xd0, 0xd4, 0x59, 0x1f, 0xf3, 0x55, 0x58, 0x2a,
	0xa4, 0x9e, 0x84, 0xda, 0x0a, 0x07, 0x7e, 0xe7, 0xb9, 0x19, 0x69, 0x35, 0xae, 0xb0, 0xef, 0x69,
	0x37, 0xbb, 0x77, 0x42, 0x53, 0x41, 0x3f, 0x78, 0x7c, 0x1d, 0x96, 0xa7, 0x52, 0x2b, 0xfc, 0x47,
	0x8f, 0xaf, 0x40, 0x8f, 0xa4, 0xce, 0x30, 0xc3, 0x7e, 0x72, 0x20, 0x89, 0xaa, 0x81, 0x3f, 0x3b,
	0x86, 0x52, 0x55, 0x0d, 0xff, 0xc5, 0x25, 0x23, 0x86, 0x72, 0x54, 0x86, 0xbd, 0xf0, 0x48, 0xe9,
	0x34, 0x59, 0x09, 0xb3, 0x7f, 0x3c, 0xce, 0xa0, 0x53, 0xe8, 0x77, 0x2f, 0x86, 0x7d, 0xdc, 0x70,
	0xda, 0xcb, 0xb8, 0x02, 0xfb, 0xa4, 0xc1, 0x7b, 0xe0, 0x53, 0x3d, 0x85, 0xfd, 0x69, 0x83, 0x77,
	0x60, 0x7e, 0x28, 0x0d, 0x6a, 0xcb, 0x3e, 0xa0, 0xa9, 0xce, 0x17, 0x2b, 0xc0, 0x3e, 0xa4, 0xb7,
	0x33, 0xe7, 0x46, 0xcc, 0x9e, 0x3b, 0x47, 0xb1, 0xac, 0xec, 0xf7, 0xa6, 0x53, 0x54, 0xdf, 0xdc,
	0x3f, 0x9a, 0x94, 0xe9, 0x00, 0x6d, 0xf5, 0x54, 0xd9, 0x9f, 0x4d, 0xbe, 0x09, 0x97, 0x0f, 0xd0,
	0x56, 0x2d, 0x26, 0x87, 0x30, 0x56, 0x44, 0x86, 0xfd, 0xd5, 0xe4, 0xaf, 0xc2, 0xfa, 0x01, 0xda,
	0x59, 0xa9, 0x35, 0xe7, 0xdf, 0x4d, 0x2a, 0xa5, 0x2c, 0x6c, 0x28, 0xcf, 0x15, 0xfb, 0xb5, 0xc5,
	0x17, 0x61, 0xe1, 0x54, 0x64, 0x78, 0x2a, 0xa2, 0xc7, 0xec, 0x33, 0x9f, 0x3a, 0xf8, 0x4e, 0x8e,
	0x7a, 0x72, 0xa4, 0x62, 0xa4, 0xab, 0x86, 0x7d, 0xee, 0x53, 0x69, 0xd4, 0xd6, 0xa2, 0xb4, 0x2f,
	0x9c, 0x5d, 0x2e, 0xd7, 0x70, 0xc0, 0xbe, 0xa4, 0x8f, 0x16, 0x94, 0xf6, 0xe9, 0xe8, 0x98, 0x7d,
	0xe5, 0x53, 0x1f, 0x6f, 0xa5, 0xa9, 0x8a, 0x42, 0x3b, 0xeb, 0xe3, 0xd7, 0x3e, 0x0d, 0xa2, 0xb6,
	0x17, 0xa5, 0xa8, 0x6f, 0x7c, 0xbe, 0x36, 0xdb, 0x17, 0xd7, 0x96, 0x01, 0xed, 0xcb, 0xb7, 0xfe,
	0xde, 0x9b, 0xef, 0xde, 0x48, 0x84, 0xbd, 0xc8, 0xcf, 0xe8, 0xbb, 0xbf, 0x5b, 0xfc, 0x08, 0xae,
	0x0b, 0x55, 0x9e, 0x76, 0x85, 0xb4, 0xa8, 0x65, 0x98, 0xee, 0xba, 0x7f, 0xc3, 0x6e, 0xf1, 0x6f,
	0x18, 0x9f, 0x9d, 0xcd, 0x3b, 0xfb, 0xc6, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x22, 0x4d,
	0xc1, 0xf5, 0x07, 0x00, 0x00,
}
