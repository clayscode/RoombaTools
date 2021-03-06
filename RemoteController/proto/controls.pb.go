// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.8.0
// source: controls.proto

package controls

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

type RoombaControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rotatedegrees float64 `protobuf:"fixed64,1,opt,name=rotatedegrees,proto3" json:"rotatedegrees,omitempty"`
	Speed         float64 `protobuf:"fixed64,2,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *RoombaControlRequest) Reset() {
	*x = RoombaControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoombaControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoombaControlRequest) ProtoMessage() {}

func (x *RoombaControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoombaControlRequest.ProtoReflect.Descriptor instead.
func (*RoombaControlRequest) Descriptor() ([]byte, []int) {
	return file_controls_proto_rawDescGZIP(), []int{0}
}

func (x *RoombaControlRequest) GetRotatedegrees() float64 {
	if x != nil {
		return x.Rotatedegrees
	}
	return 0
}

func (x *RoombaControlRequest) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

type RoombaControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RoombaControlResponse) Reset() {
	*x = RoombaControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoombaControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoombaControlResponse) ProtoMessage() {}

func (x *RoombaControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoombaControlResponse.ProtoReflect.Descriptor instead.
func (*RoombaControlResponse) Descriptor() ([]byte, []int) {
	return file_controls_proto_rawDescGZIP(), []int{1}
}

func (x *RoombaControlResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RoombaControlResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack bool `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *AckRequest) Reset() {
	*x = AckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckRequest) ProtoMessage() {}

func (x *AckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckRequest.ProtoReflect.Descriptor instead.
func (*AckRequest) Descriptor() ([]byte, []int) {
	return file_controls_proto_rawDescGZIP(), []int{2}
}

func (x *AckRequest) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

type AckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack bool `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *AckResponse) Reset() {
	*x = AckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckResponse) ProtoMessage() {}

func (x *AckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckResponse.ProtoReflect.Descriptor instead.
func (*AckResponse) Descriptor() ([]byte, []int) {
	return file_controls_proto_rawDescGZIP(), []int{3}
}

func (x *AckResponse) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

var File_controls_proto protoreflect.FileDescriptor

var file_controls_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x52, 0x0a, 0x14, 0x52, 0x6f, 0x6f, 0x6d, 0x62, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x22, 0x47, 0x0a, 0x15, 0x52, 0x6f, 0x6f, 0x6d, 0x62, 0x61, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x1e, 0x0a,
	0x0a, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x22, 0x1f, 0x0a,
	0x0b, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x32, 0x77,
	0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x62, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x65,
	0x72, 0x12, 0x42, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x62, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x12, 0x15, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x62, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x62, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x20, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x0b, 0x2e, 0x41,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controls_proto_rawDescOnce sync.Once
	file_controls_proto_rawDescData = file_controls_proto_rawDesc
)

func file_controls_proto_rawDescGZIP() []byte {
	file_controls_proto_rawDescOnce.Do(func() {
		file_controls_proto_rawDescData = protoimpl.X.CompressGZIP(file_controls_proto_rawDescData)
	})
	return file_controls_proto_rawDescData
}

var file_controls_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controls_proto_goTypes = []interface{}{
	(*RoombaControlRequest)(nil),  // 0: RoombaControlRequest
	(*RoombaControlResponse)(nil), // 1: RoombaControlResponse
	(*AckRequest)(nil),            // 2: AckRequest
	(*AckResponse)(nil),           // 3: AckResponse
}
var file_controls_proto_depIdxs = []int32{
	0, // 0: RoombaControler.RoombaControl:input_type -> RoombaControlRequest
	2, // 1: RoombaControler.Ack:input_type -> AckRequest
	1, // 2: RoombaControler.RoombaControl:output_type -> RoombaControlResponse
	3, // 3: RoombaControler.Ack:output_type -> AckResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_controls_proto_init() }
func file_controls_proto_init() {
	if File_controls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoombaControlRequest); i {
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
		file_controls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoombaControlResponse); i {
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
		file_controls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckRequest); i {
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
		file_controls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckResponse); i {
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
			RawDescriptor: file_controls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controls_proto_goTypes,
		DependencyIndexes: file_controls_proto_depIdxs,
		MessageInfos:      file_controls_proto_msgTypes,
	}.Build()
	File_controls_proto = out.File
	file_controls_proto_rawDesc = nil
	file_controls_proto_goTypes = nil
	file_controls_proto_depIdxs = nil
}
