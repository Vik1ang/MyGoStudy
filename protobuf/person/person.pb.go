// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: person/person.proto

package person

import (
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

type Home struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person     `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	V       *Home_Visitor `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *Home) Reset() {
	*x = Home{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Home) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home) ProtoMessage() {}

func (x *Home) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Home.ProtoReflect.Descriptor instead.
func (*Home) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0}
}

func (x *Home) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *Home) GetV() *Home_Visitor {
	if x != nil {
		return x.V
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32             `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Sex     bool              `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	Test    []string          `protobuf:"bytes,4,rep,name=test,proto3" json:"test,omitempty"`
	TestMap map[string]string `protobuf:"bytes,5,rep,name=test_map,json=testMap,proto3" json:"test_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *Person) GetTest() []string {
	if x != nil {
		return x.Test
	}
	return nil
}

func (x *Person) GetTestMap() map[string]string {
	if x != nil {
		return x.TestMap
	}
	return nil
}

type Home_Visitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Home_Visitor) Reset() {
	*x = Home_Visitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Home_Visitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home_Visitor) ProtoMessage() {}

func (x *Home_Visitor) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Home_Visitor.ProtoReflect.Descriptor instead.
func (*Home_Visitor) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Home_Visitor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_person_person_proto protoreflect.FileDescriptor

var file_person_person_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x73, 0x0a,
	0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12,
	0x22, 0x0a, 0x01, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x52, 0x01, 0x76, 0x1a, 0x1d, 0x0a, 0x07, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x61,
	0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x22, 0x5a,
	0x20, 0x4d, 0x79, 0x47, 0x6f, 0x53, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f,
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

var file_person_person_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_person_person_proto_goTypes = []interface{}{
	(*Home)(nil),         // 0: person.Home
	(*Person)(nil),       // 1: person.Person
	(*Home_Visitor)(nil), // 2: person.Home.Visitor
	nil,                  // 3: person.Person.TestMapEntry
}
var file_person_person_proto_depIdxs = []int32{
	1, // 0: person.Home.persons:type_name -> person.Person
	2, // 1: person.Home.v:type_name -> person.Home.Visitor
	3, // 2: person.Person.test_map:type_name -> person.Person.TestMapEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_person_person_proto_init() }
func file_person_person_proto_init() {
	if File_person_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Home); i {
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
			switch v := v.(*Person); i {
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
			switch v := v.(*Home_Visitor); i {
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
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
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