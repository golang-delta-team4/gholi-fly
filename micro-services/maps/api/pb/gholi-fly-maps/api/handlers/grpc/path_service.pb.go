// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: path_service.proto

package pathservice

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

// Request for getting a path by ID
type GetPathRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Path ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPathRequest) Reset() {
	*x = GetPathRequest{}
	mi := &file_path_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPathRequest) ProtoMessage() {}

func (x *GetPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_path_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPathRequest.ProtoReflect.Descriptor instead.
func (*GetPathRequest) Descriptor() ([]byte, []int) {
	return file_path_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPathRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response for returning path details
type GetPathResponse struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	Id                    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceTerminalId      string                 `protobuf:"bytes,2,opt,name=source_terminal_id,json=sourceTerminalId,proto3" json:"source_terminal_id,omitempty"`
	DestinationTerminalId string                 `protobuf:"bytes,3,opt,name=destination_terminal_id,json=destinationTerminalId,proto3" json:"destination_terminal_id,omitempty"`
	DistanceKm            float32                `protobuf:"fixed32,4,opt,name=distance_km,json=distanceKm,proto3" json:"distance_km,omitempty"`
	RouteCode             string                 `protobuf:"bytes,5,opt,name=route_code,json=routeCode,proto3" json:"route_code,omitempty"`
	VehicleType           string                 `protobuf:"bytes,6,opt,name=vehicle_type,json=vehicleType,proto3" json:"vehicle_type,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *GetPathResponse) Reset() {
	*x = GetPathResponse{}
	mi := &file_path_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPathResponse) ProtoMessage() {}

func (x *GetPathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_path_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPathResponse.ProtoReflect.Descriptor instead.
func (*GetPathResponse) Descriptor() ([]byte, []int) {
	return file_path_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPathResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPathResponse) GetSourceTerminalId() string {
	if x != nil {
		return x.SourceTerminalId
	}
	return ""
}

func (x *GetPathResponse) GetDestinationTerminalId() string {
	if x != nil {
		return x.DestinationTerminalId
	}
	return ""
}

func (x *GetPathResponse) GetDistanceKm() float32 {
	if x != nil {
		return x.DistanceKm
	}
	return 0
}

func (x *GetPathResponse) GetRouteCode() string {
	if x != nil {
		return x.RouteCode
	}
	return ""
}

func (x *GetPathResponse) GetVehicleType() string {
	if x != nil {
		return x.VehicleType
	}
	return ""
}

var File_path_service_proto protoreflect.FileDescriptor

var file_path_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xea, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6b, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x6d, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x32, 0x53, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x74,
	0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x68, 0x6f, 0x6c, 0x69, 0x2d, 0x66,
	0x6c, 0x79, 0x2d, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x61, 0x74, 0x68, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_path_service_proto_rawDescOnce sync.Once
	file_path_service_proto_rawDescData = file_path_service_proto_rawDesc
)

func file_path_service_proto_rawDescGZIP() []byte {
	file_path_service_proto_rawDescOnce.Do(func() {
		file_path_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_path_service_proto_rawDescData)
	})
	return file_path_service_proto_rawDescData
}

var file_path_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_path_service_proto_goTypes = []any{
	(*GetPathRequest)(nil),  // 0: pathservice.GetPathRequest
	(*GetPathResponse)(nil), // 1: pathservice.GetPathResponse
}
var file_path_service_proto_depIdxs = []int32{
	0, // 0: pathservice.PathService.GetPath:input_type -> pathservice.GetPathRequest
	1, // 1: pathservice.PathService.GetPath:output_type -> pathservice.GetPathResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_path_service_proto_init() }
func file_path_service_proto_init() {
	if File_path_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_path_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_path_service_proto_goTypes,
		DependencyIndexes: file_path_service_proto_depIdxs,
		MessageInfos:      file_path_service_proto_msgTypes,
	}.Build()
	File_path_service_proto = out.File
	file_path_service_proto_rawDesc = nil
	file_path_service_proto_goTypes = nil
	file_path_service_proto_depIdxs = nil
}
