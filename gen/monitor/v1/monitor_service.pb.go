// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: monitor/v1/monitor_service.proto

package monitorv1

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

type GetLatencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	NumOfRequests int32  `protobuf:"varint,2,opt,name=num_of_requests,json=numOfRequests,proto3" json:"num_of_requests,omitempty"`
}

func (x *GetLatencyRequest) Reset() {
	*x = GetLatencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_v1_monitor_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatencyRequest) ProtoMessage() {}

func (x *GetLatencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_v1_monitor_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatencyRequest.ProtoReflect.Descriptor instead.
func (*GetLatencyRequest) Descriptor() ([]byte, []int) {
	return file_monitor_v1_monitor_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetLatencyRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetLatencyRequest) GetNumOfRequests() int32 {
	if x != nil {
		return x.NumOfRequests
	}
	return 0
}

type GetLatencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latency float32 `protobuf:"fixed32,1,opt,name=latency,proto3" json:"latency,omitempty"`
}

func (x *GetLatencyResponse) Reset() {
	*x = GetLatencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_v1_monitor_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatencyResponse) ProtoMessage() {}

func (x *GetLatencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_v1_monitor_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatencyResponse.ProtoReflect.Descriptor instead.
func (*GetLatencyResponse) Descriptor() ([]byte, []int) {
	return file_monitor_v1_monitor_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetLatencyResponse) GetLatency() float32 {
	if x != nil {
		return x.Latency
	}
	return 0
}

var File_monitor_v1_monitor_service_proto protoreflect.FileDescriptor

var file_monitor_v1_monitor_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x4d,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x66, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x6e, 0x75, 0x6d, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x2e, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x32, 0x5f, 0x0a,
	0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa3,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x13, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x61, 0x6d, 0x65, 0x6a, 0x6e, 0x72, 0x2f, 0x73, 0x65,
	0x73, 0x68, 0x61, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_v1_monitor_service_proto_rawDescOnce sync.Once
	file_monitor_v1_monitor_service_proto_rawDescData = file_monitor_v1_monitor_service_proto_rawDesc
)

func file_monitor_v1_monitor_service_proto_rawDescGZIP() []byte {
	file_monitor_v1_monitor_service_proto_rawDescOnce.Do(func() {
		file_monitor_v1_monitor_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_v1_monitor_service_proto_rawDescData)
	})
	return file_monitor_v1_monitor_service_proto_rawDescData
}

var file_monitor_v1_monitor_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_monitor_v1_monitor_service_proto_goTypes = []interface{}{
	(*GetLatencyRequest)(nil),  // 0: monitor.v1.GetLatencyRequest
	(*GetLatencyResponse)(nil), // 1: monitor.v1.GetLatencyResponse
}
var file_monitor_v1_monitor_service_proto_depIdxs = []int32{
	0, // 0: monitor.v1.MonitorService.GetLatency:input_type -> monitor.v1.GetLatencyRequest
	1, // 1: monitor.v1.MonitorService.GetLatency:output_type -> monitor.v1.GetLatencyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_monitor_v1_monitor_service_proto_init() }
func file_monitor_v1_monitor_service_proto_init() {
	if File_monitor_v1_monitor_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_v1_monitor_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatencyRequest); i {
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
		file_monitor_v1_monitor_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatencyResponse); i {
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
			RawDescriptor: file_monitor_v1_monitor_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitor_v1_monitor_service_proto_goTypes,
		DependencyIndexes: file_monitor_v1_monitor_service_proto_depIdxs,
		MessageInfos:      file_monitor_v1_monitor_service_proto_msgTypes,
	}.Build()
	File_monitor_v1_monitor_service_proto = out.File
	file_monitor_v1_monitor_service_proto_rawDesc = nil
	file_monitor_v1_monitor_service_proto_goTypes = nil
	file_monitor_v1_monitor_service_proto_depIdxs = nil
}
