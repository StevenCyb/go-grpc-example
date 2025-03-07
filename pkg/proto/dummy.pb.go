// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.3
// source: demo.proto

package proto

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

type DemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DemoRequest) Reset() {
	*x = DemoRequest{}
	mi := &file_demo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoRequest) ProtoMessage() {}

func (x *DemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoRequest.ProtoReflect.Descriptor instead.
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

func (x *DemoRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result uint32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DemoResponse) Reset() {
	*x = DemoResponse{}
	mi := &file_demo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoResponse) ProtoMessage() {}

func (x *DemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoResponse.ProtoReflect.Descriptor instead.
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *DemoResponse) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64,
	0x75, 0x6d, 0x6d, 0x79, 0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x8a, 0x03,
	0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x0d, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d,
	0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x13,
	0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3f, 0x0a,
	0x0e, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d,
	0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x42,
	0x0a, 0x0f, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x12,
	0x13, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d,
	0x6d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d,
	0x55, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x13, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData = file_demo_proto_rawDesc
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_proto_rawDescData)
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demo_proto_goTypes = []any{
	(*DemoRequest)(nil),  // 0: demo.DemoRequest
	(*DemoResponse)(nil), // 1: demo.DemoResponse
}
var file_demo_proto_depIdxs = []int32{
	0, // 0: demo.DemoService.OneReqOneResp:input_type -> demo.DemoRequest
	0, // 1: demo.DemoService.OneReqManyResp:input_type -> demo.DemoRequest
	0, // 2: demo.DemoService.ManyReqOneResp:input_type -> demo.DemoRequest
	0, // 3: demo.DemoService.ManyReqManyResp:input_type -> demo.DemoRequest
	0, // 4: demo.DemoService.UnaryFail:input_type -> demo.DemoRequest
	0, // 5: demo.DemoService.UnaryDeadline:input_type -> demo.DemoRequest
	1, // 6: demo.DemoService.OneReqOneResp:output_type -> demo.DemoResponse
	1, // 7: demo.DemoService.OneReqManyResp:output_type -> demo.DemoResponse
	1, // 8: demo.DemoService.ManyReqOneResp:output_type -> demo.DemoResponse
	1, // 9: demo.DemoService.ManyReqManyResp:output_type -> demo.DemoResponse
	1, // 10: demo.DemoService.UnaryFail:output_type -> demo.DemoResponse
	1, // 11: demo.DemoService.UnaryDeadline:output_type -> demo.DemoResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_rawDesc = nil
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}
