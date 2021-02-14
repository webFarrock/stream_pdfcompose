// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/pdfcompose_stream.proto

package stream_pdfcompose

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Upfiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upfile1 []byte `protobuf:"bytes,1,opt,name=Upfile1,proto3" json:"Upfile1,omitempty"`
	Upfile2 []byte `protobuf:"bytes,2,opt,name=Upfile2,proto3" json:"Upfile2,omitempty"`
	Upfile3 []byte `protobuf:"bytes,3,opt,name=Upfile3,proto3" json:"Upfile3,omitempty"`
}

func (x *Upfiles) Reset() {
	*x = Upfiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pdfcompose_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upfiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upfiles) ProtoMessage() {}

func (x *Upfiles) ProtoReflect() protoreflect.Message {
	mi := &file_api_pdfcompose_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upfiles.ProtoReflect.Descriptor instead.
func (*Upfiles) Descriptor() ([]byte, []int) {
	return file_api_pdfcompose_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Upfiles) GetUpfile1() []byte {
	if x != nil {
		return x.Upfile1
	}
	return nil
}

func (x *Upfiles) GetUpfile2() []byte {
	if x != nil {
		return x.Upfile2
	}
	return nil
}

func (x *Upfiles) GetUpfile3() []byte {
	if x != nil {
		return x.Upfile3
	}
	return nil
}

type PdfResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	File    []byte `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *PdfResult) Reset() {
	*x = PdfResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pdfcompose_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PdfResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PdfResult) ProtoMessage() {}

func (x *PdfResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_pdfcompose_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PdfResult.ProtoReflect.Descriptor instead.
func (*PdfResult) Descriptor() ([]byte, []int) {
	return file_api_pdfcompose_stream_proto_rawDescGZIP(), []int{1}
}

func (x *PdfResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PdfResult) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

var File_api_pdfcompose_stream_proto protoreflect.FileDescriptor

var file_api_pdfcompose_stream_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x07, 0x55, 0x70, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x55, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x31, 0x12, 0x18,
	0x0a, 0x07, 0x55, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x55, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x70, 0x66, 0x69,
	0x6c, 0x65, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x55, 0x70, 0x66, 0x69, 0x6c,
	0x65, 0x33, 0x22, 0x39, 0x0a, 0x09, 0x50, 0x64, 0x66, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x4b, 0x0a,
	0x0a, 0x50, 0x64, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x64, 0x66, 0x12, 0x13, 0x2e, 0x70, 0x64, 0x66, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x15, 0x2e,
	0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x50, 0x64, 0x66, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x70, 0x64,
	0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3b,
	0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_pdfcompose_stream_proto_rawDescOnce sync.Once
	file_api_pdfcompose_stream_proto_rawDescData = file_api_pdfcompose_stream_proto_rawDesc
)

func file_api_pdfcompose_stream_proto_rawDescGZIP() []byte {
	file_api_pdfcompose_stream_proto_rawDescOnce.Do(func() {
		file_api_pdfcompose_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pdfcompose_stream_proto_rawDescData)
	})
	return file_api_pdfcompose_stream_proto_rawDescData
}

var file_api_pdfcompose_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_pdfcompose_stream_proto_goTypes = []interface{}{
	(*Upfiles)(nil),   // 0: pdfcompose.Upfiles
	(*PdfResult)(nil), // 1: pdfcompose.PdfResult
}
var file_api_pdfcompose_stream_proto_depIdxs = []int32{
	0, // 0: pdfcompose.PdfCompose.CreatePdf:input_type -> pdfcompose.Upfiles
	1, // 1: pdfcompose.PdfCompose.CreatePdf:output_type -> pdfcompose.PdfResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_pdfcompose_stream_proto_init() }
func file_api_pdfcompose_stream_proto_init() {
	if File_api_pdfcompose_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_pdfcompose_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upfiles); i {
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
		file_api_pdfcompose_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PdfResult); i {
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
			RawDescriptor: file_api_pdfcompose_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pdfcompose_stream_proto_goTypes,
		DependencyIndexes: file_api_pdfcompose_stream_proto_depIdxs,
		MessageInfos:      file_api_pdfcompose_stream_proto_msgTypes,
	}.Build()
	File_api_pdfcompose_stream_proto = out.File
	file_api_pdfcompose_stream_proto_rawDesc = nil
	file_api_pdfcompose_stream_proto_goTypes = nil
	file_api_pdfcompose_stream_proto_depIdxs = nil
}
