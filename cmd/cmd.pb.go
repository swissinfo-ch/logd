/*
Copyright © 2024 JOSEPH INNES <avianpneuma@gmail.com>
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: cmd.proto

package cmd

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Name int32

const (
	Name_WRITE Name = 0
	Name_TAIL  Name = 1
	Name_PING  Name = 2
	Name_QUERY Name = 3
)

// Enum value maps for Name.
var (
	Name_name = map[int32]string{
		0: "WRITE",
		1: "TAIL",
		2: "PING",
		3: "QUERY",
	}
	Name_value = map[string]int32{
		"WRITE": 0,
		"TAIL":  1,
		"PING":  2,
		"QUERY": 3,
	}
)

func (x Name) Enum() *Name {
	p := new(Name)
	*p = x
	return p
}

func (x Name) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Name) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[0].Descriptor()
}

func (Name) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[0]
}

func (x Name) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Name.Descriptor instead.
func (Name) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

type Lvl int32

const (
	Lvl_TRACE Lvl = 0
	Lvl_DEBUG Lvl = 1
	Lvl_INFO  Lvl = 2
	Lvl_WARN  Lvl = 3
	Lvl_ERROR Lvl = 4
	Lvl_FATAL Lvl = 5
)

// Enum value maps for Lvl.
var (
	Lvl_name = map[int32]string{
		0: "TRACE",
		1: "DEBUG",
		2: "INFO",
		3: "WARN",
		4: "ERROR",
		5: "FATAL",
	}
	Lvl_value = map[string]int32{
		"TRACE": 0,
		"DEBUG": 1,
		"INFO":  2,
		"WARN":  3,
		"ERROR": 4,
		"FATAL": 5,
	}
)

func (x Lvl) Enum() *Lvl {
	p := new(Lvl)
	*p = x
	return p
}

func (x Lvl) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lvl) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[1].Descriptor()
}

func (Lvl) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[1]
}

func (x Lvl) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lvl.Descriptor instead.
func (Lvl) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

type HttpMethod int32

const (
	HttpMethod_UNKNOWN HttpMethod = 0
	HttpMethod_DELETE  HttpMethod = 1
	HttpMethod_GET     HttpMethod = 2
	HttpMethod_HEAD    HttpMethod = 3
	HttpMethod_OPTIONS HttpMethod = 4
	HttpMethod_PATCH   HttpMethod = 5
	HttpMethod_POST    HttpMethod = 6
	HttpMethod_PUT     HttpMethod = 7
)

// Enum value maps for HttpMethod.
var (
	HttpMethod_name = map[int32]string{
		0: "UNKNOWN",
		1: "DELETE",
		2: "GET",
		3: "HEAD",
		4: "OPTIONS",
		5: "PATCH",
		6: "POST",
		7: "PUT",
	}
	HttpMethod_value = map[string]int32{
		"UNKNOWN": 0,
		"DELETE":  1,
		"GET":     2,
		"HEAD":    3,
		"OPTIONS": 4,
		"PATCH":   5,
		"POST":    6,
		"PUT":     7,
	}
)

func (x HttpMethod) Enum() *HttpMethod {
	p := new(HttpMethod)
	*p = x
	return p
}

func (x HttpMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[2].Descriptor()
}

func (HttpMethod) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[2]
}

func (x HttpMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpMethod.Descriptor instead.
func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{2}
}

type Cmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name Name     `protobuf:"varint,1,opt,name=name,proto3,enum=Name" json:"name,omitempty"`
	Msg  *Msg     `protobuf:"bytes,2,opt,name=msg,proto3,oneof" json:"msg,omitempty"`
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *Cmd) Reset() {
	*x = Cmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmd) ProtoMessage() {}

func (x *Cmd) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmd.ProtoReflect.Descriptor instead.
func (*Cmd) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *Cmd) GetName() Name {
	if x != nil {
		return x.Name
	}
	return Name_WRITE
}

func (x *Cmd) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Cmd) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
	Env            string                 `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	Svc            string                 `protobuf:"bytes,4,opt,name=svc,proto3" json:"svc,omitempty"`
	Fn             string                 `protobuf:"bytes,5,opt,name=fn,proto3" json:"fn,omitempty"`
	Lvl            *Lvl                   `protobuf:"varint,6,opt,name=lvl,proto3,enum=Lvl,oneof" json:"lvl,omitempty"`
	Txt            *string                `protobuf:"bytes,7,opt,name=txt,proto3,oneof" json:"txt,omitempty"`
	StackTrace     *string                `protobuf:"bytes,8,opt,name=stackTrace,proto3,oneof" json:"stackTrace,omitempty"`
	HttpMethod     *HttpMethod            `protobuf:"varint,9,opt,name=httpMethod,proto3,enum=HttpMethod,oneof" json:"httpMethod,omitempty"`
	Url            *string                `protobuf:"bytes,10,opt,name=url,proto3,oneof" json:"url,omitempty"`
	ResponseStatus *int32                 `protobuf:"varint,11,opt,name=responseStatus,proto3,oneof" json:"responseStatus,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetT() *timestamppb.Timestamp {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *Msg) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Msg) GetSvc() string {
	if x != nil {
		return x.Svc
	}
	return ""
}

func (x *Msg) GetFn() string {
	if x != nil {
		return x.Fn
	}
	return ""
}

func (x *Msg) GetLvl() Lvl {
	if x != nil && x.Lvl != nil {
		return *x.Lvl
	}
	return Lvl_TRACE
}

func (x *Msg) GetTxt() string {
	if x != nil && x.Txt != nil {
		return *x.Txt
	}
	return ""
}

func (x *Msg) GetStackTrace() string {
	if x != nil && x.StackTrace != nil {
		return *x.StackTrace
	}
	return ""
}

func (x *Msg) GetHttpMethod() HttpMethod {
	if x != nil && x.HttpMethod != nil {
		return *x.HttpMethod
	}
	return HttpMethod_UNKNOWN
}

func (x *Msg) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *Msg) GetResponseStatus() int32 {
	if x != nil && x.ResponseStatus != nil {
		return *x.ResponseStatus
	}
	return 0
}

var File_cmd_proto protoreflect.FileDescriptor

var file_cmd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x03,
	0x43, 0x6d, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4d, 0x73,
	0x67, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0xfb, 0x02, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12,
	0x28, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x01, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x76, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x76, 0x63, 0x12, 0x0e, 0x0a,
	0x02, 0x66, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x66, 0x6e, 0x12, 0x1b, 0x0a,
	0x03, 0x6c, 0x76, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x04, 0x2e, 0x4c, 0x76, 0x6c,
	0x48, 0x00, 0x52, 0x03, 0x6c, 0x76, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x74, 0x78,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x74, 0x78, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x03, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x6c, 0x76, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x78, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x68, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75,
	0x72, 0x6c, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x30, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x41, 0x49, 0x4c,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x03, 0x2a, 0x45, 0x0a, 0x03, 0x4c, 0x76, 0x6c, 0x12, 0x09,
	0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42,
	0x55, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x05, 0x2a, 0x63,
	0x0a, 0x0a, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x53, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x05,
	0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55,
	0x54, 0x10, 0x07, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_proto_rawDescOnce sync.Once
	file_cmd_proto_rawDescData = file_cmd_proto_rawDesc
)

func file_cmd_proto_rawDescGZIP() []byte {
	file_cmd_proto_rawDescOnce.Do(func() {
		file_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_proto_rawDescData)
	})
	return file_cmd_proto_rawDescData
}

var file_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_proto_goTypes = []interface{}{
	(Name)(0),                     // 0: Name
	(Lvl)(0),                      // 1: Lvl
	(HttpMethod)(0),               // 2: HttpMethod
	(*Cmd)(nil),                   // 3: Cmd
	(*Msg)(nil),                   // 4: Msg
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_cmd_proto_depIdxs = []int32{
	0, // 0: Cmd.name:type_name -> Name
	4, // 1: Cmd.msg:type_name -> Msg
	5, // 2: Msg.t:type_name -> google.protobuf.Timestamp
	1, // 3: Msg.lvl:type_name -> Lvl
	2, // 4: Msg.httpMethod:type_name -> HttpMethod
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cmd_proto_init() }
func file_cmd_proto_init() {
	if File_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_cmd_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_cmd_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_proto_goTypes,
		DependencyIndexes: file_cmd_proto_depIdxs,
		EnumInfos:         file_cmd_proto_enumTypes,
		MessageInfos:      file_cmd_proto_msgTypes,
	}.Build()
	File_cmd_proto = out.File
	file_cmd_proto_rawDesc = nil
	file_cmd_proto_goTypes = nil
	file_cmd_proto_depIdxs = nil
}
