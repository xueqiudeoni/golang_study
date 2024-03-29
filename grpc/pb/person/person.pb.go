// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.9.0
// source: person/person.proto

package person

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Mbody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Mbody) Reset() {
	*x = Mbody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mbody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mbody) ProtoMessage() {}

func (x *Mbody) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mbody.ProtoReflect.Descriptor instead.
func (*Mbody) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0}
}

func (x *Mbody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PersonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age   int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Mbody *Mbody `protobuf:"bytes,3,opt,name=Mbody,proto3" json:"Mbody,omitempty"`
}

func (x *PersonReq) Reset() {
	*x = PersonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonReq) ProtoMessage() {}

func (x *PersonReq) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonReq.ProtoReflect.Descriptor instead.
func (*PersonReq) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1}
}

func (x *PersonReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonReq) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PersonReq) GetMbody() *Mbody {
	if x != nil {
		return x.Mbody
	}
	return nil
}

type PersonRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age   int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Mbody *Mbody `protobuf:"bytes,3,opt,name=Mbody,proto3" json:"Mbody,omitempty"`
}

func (x *PersonRes) Reset() {
	*x = PersonRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRes) ProtoMessage() {}

func (x *PersonRes) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRes.ProtoReflect.Descriptor instead.
func (*PersonRes) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{2}
}

func (x *PersonRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonRes) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PersonRes) GetMbody() *Mbody {
	if x != nil {
		return x.Mbody
	}
	return nil
}

var File_person_person_proto protoreflect.FileDescriptor

var file_person_person_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x05, 0x4d,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x4d,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x4d, 0x62, 0x6f, 0x64, 0x79, 0x52, 0x05, 0x4d, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x56, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x4d, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x4d, 0x62, 0x6f, 0x64,
	0x79, 0x52, 0x05, 0x4d, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x6e, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x27, 0x12, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x61, 0x67, 0x65, 0x7d, 0x2f, 0x7b, 0x4d, 0x62, 0x6f,
	0x64, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_person_proto_rawDescOnce sync.Once
	file_person_person_proto_rawDescData = file_person_person_proto_rawDesc
)

func file_person_person_proto_rawDescGZIP() []byte {
	file_person_person_proto_rawDescOnce.Do(func() {
		file_person_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_person_proto_rawDescData)
	})
	return file_person_person_proto_rawDescData
}

var file_person_person_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_person_person_proto_goTypes = []interface{}{
	(*Mbody)(nil),     // 0: person.Mbody
	(*PersonReq)(nil), // 1: person.PersonReq
	(*PersonRes)(nil), // 2: person.PersonRes
}
var file_person_person_proto_depIdxs = []int32{
	0, // 0: person.PersonReq.Mbody:type_name -> person.Mbody
	0, // 1: person.PersonRes.Mbody:type_name -> person.Mbody
	1, // 2: person.SearchService.Search:input_type -> person.PersonReq
	2, // 3: person.SearchService.Search:output_type -> person.PersonRes
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_person_person_proto_init() }
func file_person_person_proto_init() {
	if File_person_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mbody); i {
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
		file_person_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonReq); i {
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
		file_person_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRes); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_person_proto_goTypes,
		DependencyIndexes: file_person_person_proto_depIdxs,
		MessageInfos:      file_person_person_proto_msgTypes,
	}.Build()
	File_person_person_proto = out.File
	file_person_person_proto_rawDesc = nil
	file_person_person_proto_goTypes = nil
	file_person_person_proto_depIdxs = nil
}
