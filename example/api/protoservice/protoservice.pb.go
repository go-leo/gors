// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: example/api/protoservice/protoservice.proto

package protoservice

import (
	protodemo "github.com/go-leo/gors/example/api/protodemo"
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

type HelloRequest1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age          int32                   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Salary       float64                 `protobuf:"fixed64,3,opt,name=salary,proto3" json:"salary,omitempty"`
	Token        string                  `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	HelloRequest *protodemo.HelloRequest `protobuf:"bytes,5,opt,name=helloRequest,proto3" json:"helloRequest,omitempty"`
}

func (x *HelloRequest1) Reset() {
	*x = HelloRequest1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_protoservice_protoservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest1) ProtoMessage() {}

func (x *HelloRequest1) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_protoservice_protoservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest1.ProtoReflect.Descriptor instead.
func (*HelloRequest1) Descriptor() ([]byte, []int) {
	return file_example_api_protoservice_protoservice_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest1) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *HelloRequest1) GetSalary() float64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *HelloRequest1) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *HelloRequest1) GetHelloRequest() *protodemo.HelloRequest {
	if x != nil {
		return x.HelloRequest
	}
	return nil
}

type HelloReply1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply1) Reset() {
	*x = HelloReply1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_api_protoservice_protoservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply1) ProtoMessage() {}

func (x *HelloReply1) ProtoReflect() protoreflect.Message {
	mi := &file_example_api_protoservice_protoservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply1.ProtoReflect.Descriptor instead.
func (*HelloReply1) Descriptor() ([]byte, []int) {
	return file_example_api_protoservice_protoservice_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply1) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_example_api_protoservice_protoservice_proto protoreflect.FileDescriptor

var file_example_api_protoservice_protoservice_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x25, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65,
	0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0c, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x52,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42,
	0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x31, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x31,
	0x22, 0x00, 0x32, 0x53, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x32, 0x12, 0x42, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x31, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x31, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x67, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_api_protoservice_protoservice_proto_rawDescOnce sync.Once
	file_example_api_protoservice_protoservice_proto_rawDescData = file_example_api_protoservice_protoservice_proto_rawDesc
)

func file_example_api_protoservice_protoservice_proto_rawDescGZIP() []byte {
	file_example_api_protoservice_protoservice_proto_rawDescOnce.Do(func() {
		file_example_api_protoservice_protoservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_api_protoservice_protoservice_proto_rawDescData)
	})
	return file_example_api_protoservice_protoservice_proto_rawDescData
}

var file_example_api_protoservice_protoservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_api_protoservice_protoservice_proto_goTypes = []interface{}{
	(*HelloRequest1)(nil),          // 0: protoservice.HelloRequest1
	(*HelloReply1)(nil),            // 1: protoservice.HelloReply1
	(*protodemo.HelloRequest)(nil), // 2: protodemo.HelloRequest
}
var file_example_api_protoservice_protoservice_proto_depIdxs = []int32{
	2, // 0: protoservice.HelloRequest1.helloRequest:type_name -> protodemo.HelloRequest
	0, // 1: protoservice.ProtoService.Method:input_type -> protoservice.HelloRequest1
	0, // 2: protoservice.ProtoService2.Method:input_type -> protoservice.HelloRequest1
	1, // 3: protoservice.ProtoService.Method:output_type -> protoservice.HelloReply1
	1, // 4: protoservice.ProtoService2.Method:output_type -> protoservice.HelloReply1
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_api_protoservice_protoservice_proto_init() }
func file_example_api_protoservice_protoservice_proto_init() {
	if File_example_api_protoservice_protoservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_api_protoservice_protoservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest1); i {
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
		file_example_api_protoservice_protoservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply1); i {
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
			RawDescriptor: file_example_api_protoservice_protoservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_example_api_protoservice_protoservice_proto_goTypes,
		DependencyIndexes: file_example_api_protoservice_protoservice_proto_depIdxs,
		MessageInfos:      file_example_api_protoservice_protoservice_proto_msgTypes,
	}.Build()
	File_example_api_protoservice_protoservice_proto = out.File
	file_example_api_protoservice_protoservice_proto_rawDesc = nil
	file_example_api_protoservice_protoservice_proto_goTypes = nil
	file_example_api_protoservice_protoservice_proto_depIdxs = nil
}
