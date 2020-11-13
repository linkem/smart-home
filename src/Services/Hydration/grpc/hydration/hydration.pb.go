// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: proto/hydration.proto

package hydration

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

type Hydration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Soil float32 `protobuf:"fixed32,1,opt,name=soil,proto3" json:"soil,omitempty"`
	Hum  float32 `protobuf:"fixed32,2,opt,name=hum,proto3" json:"hum,omitempty"`
	Temp float32 `protobuf:"fixed32,3,opt,name=temp,proto3" json:"temp,omitempty"`
	// google.protobuf.Timestamp createdDateUtc = 4;
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Hydration) Reset() {
	*x = Hydration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hydration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hydration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hydration) ProtoMessage() {}

func (x *Hydration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hydration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hydration.ProtoReflect.Descriptor instead.
func (*Hydration) Descriptor() ([]byte, []int) {
	return file_proto_hydration_proto_rawDescGZIP(), []int{0}
}

func (x *Hydration) GetSoil() float32 {
	if x != nil {
		return x.Soil
	}
	return 0
}

func (x *Hydration) GetHum() float32 {
	if x != nil {
		return x.Hum
	}
	return 0
}

func (x *Hydration) GetTemp() float32 {
	if x != nil {
		return x.Temp
	}
	return 0
}

func (x *Hydration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type HydrationGrouped struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Soil float32 `protobuf:"fixed32,1,opt,name=soil,proto3" json:"soil,omitempty"`
	Hum  float32 `protobuf:"fixed32,2,opt,name=hum,proto3" json:"hum,omitempty"`
	Temp float32 `protobuf:"fixed32,3,opt,name=temp,proto3" json:"temp,omitempty"`
	// google.protobuf.Timestamp createdDateUtc = 4;
	Samples int32 `protobuf:"varint,5,opt,name=samples,proto3" json:"samples,omitempty"`
}

func (x *HydrationGrouped) Reset() {
	*x = HydrationGrouped{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hydration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HydrationGrouped) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HydrationGrouped) ProtoMessage() {}

func (x *HydrationGrouped) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hydration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HydrationGrouped.ProtoReflect.Descriptor instead.
func (*HydrationGrouped) Descriptor() ([]byte, []int) {
	return file_proto_hydration_proto_rawDescGZIP(), []int{1}
}

func (x *HydrationGrouped) GetSoil() float32 {
	if x != nil {
		return x.Soil
	}
	return 0
}

func (x *HydrationGrouped) GetHum() float32 {
	if x != nil {
		return x.Hum
	}
	return 0
}

func (x *HydrationGrouped) GetTemp() float32 {
	if x != nil {
		return x.Temp
	}
	return 0
}

func (x *HydrationGrouped) GetSamples() int32 {
	if x != nil {
		return x.Samples
	}
	return 0
}

type HydrationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HydrationsRequest) Reset() {
	*x = HydrationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hydration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HydrationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HydrationsRequest) ProtoMessage() {}

func (x *HydrationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hydration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HydrationsRequest.ProtoReflect.Descriptor instead.
func (*HydrationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_hydration_proto_rawDescGZIP(), []int{2}
}

type HydrationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hydrations []*Hydration `protobuf:"bytes,1,rep,name=hydrations,proto3" json:"hydrations,omitempty"`
}

func (x *HydrationsResponse) Reset() {
	*x = HydrationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hydration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HydrationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HydrationsResponse) ProtoMessage() {}

func (x *HydrationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hydration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HydrationsResponse.ProtoReflect.Descriptor instead.
func (*HydrationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_hydration_proto_rawDescGZIP(), []int{3}
}

func (x *HydrationsResponse) GetHydrations() []*Hydration {
	if x != nil {
		return x.Hydrations
	}
	return nil
}

type HydrationGroupedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HydrationGroupedRequest) Reset() {
	*x = HydrationGroupedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hydration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HydrationGroupedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HydrationGroupedRequest) ProtoMessage() {}

func (x *HydrationGroupedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hydration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HydrationGroupedRequest.ProtoReflect.Descriptor instead.
func (*HydrationGroupedRequest) Descriptor() ([]byte, []int) {
	return file_proto_hydration_proto_rawDescGZIP(), []int{4}
}

type HydrationGroupedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hydrations []*HydrationGrouped `protobuf:"bytes,1,rep,name=hydrations,proto3" json:"hydrations,omitempty"`
}

func (x *HydrationGroupedResponse) Reset() {
	*x = HydrationGroupedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hydration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HydrationGroupedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HydrationGroupedResponse) ProtoMessage() {}

func (x *HydrationGroupedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hydration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HydrationGroupedResponse.ProtoReflect.Descriptor instead.
func (*HydrationGroupedResponse) Descriptor() ([]byte, []int) {
	return file_proto_hydration_proto_rawDescGZIP(), []int{5}
}

func (x *HydrationGroupedResponse) GetHydrations() []*HydrationGrouped {
	if x != nil {
		return x.Hydrations
	}
	return nil
}

var File_proto_hydration_proto protoreflect.FileDescriptor

var file_proto_hydration_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x09, 0x48, 0x79, 0x64, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x04, 0x73, 0x6f, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x68, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66,
	0x0a, 0x10, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x73, 0x6f, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x68, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x12, 0x48,
	0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x19, 0x0a,
	0x17, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x18, 0x48, 0x79, 0x64, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x48, 0x79, 0x64, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52, 0x0a, 0x68, 0x79, 0x64,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x9d, 0x01, 0x0a, 0x10, 0x48, 0x79, 0x64, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x2e,
	0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x18, 0x2e, 0x48, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x48, 0x79, 0x64,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x68, 0x79, 0x64, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_hydration_proto_rawDescOnce sync.Once
	file_proto_hydration_proto_rawDescData = file_proto_hydration_proto_rawDesc
)

func file_proto_hydration_proto_rawDescGZIP() []byte {
	file_proto_hydration_proto_rawDescOnce.Do(func() {
		file_proto_hydration_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hydration_proto_rawDescData)
	})
	return file_proto_hydration_proto_rawDescData
}

var file_proto_hydration_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_hydration_proto_goTypes = []interface{}{
	(*Hydration)(nil),                // 0: Hydration
	(*HydrationGrouped)(nil),         // 1: HydrationGrouped
	(*HydrationsRequest)(nil),        // 2: HydrationsRequest
	(*HydrationsResponse)(nil),       // 3: HydrationsResponse
	(*HydrationGroupedRequest)(nil),  // 4: HydrationGroupedRequest
	(*HydrationGroupedResponse)(nil), // 5: HydrationGroupedResponse
}
var file_proto_hydration_proto_depIdxs = []int32{
	0, // 0: HydrationsResponse.hydrations:type_name -> Hydration
	1, // 1: HydrationGroupedResponse.hydrations:type_name -> HydrationGrouped
	2, // 2: HydrationService.GetHydrations:input_type -> HydrationsRequest
	4, // 3: HydrationService.GetGroupedHydrations:input_type -> HydrationGroupedRequest
	3, // 4: HydrationService.GetHydrations:output_type -> HydrationsResponse
	5, // 5: HydrationService.GetGroupedHydrations:output_type -> HydrationGroupedResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_hydration_proto_init() }
func file_proto_hydration_proto_init() {
	if File_proto_hydration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hydration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hydration); i {
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
		file_proto_hydration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HydrationGrouped); i {
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
		file_proto_hydration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HydrationsRequest); i {
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
		file_proto_hydration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HydrationsResponse); i {
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
		file_proto_hydration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HydrationGroupedRequest); i {
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
		file_proto_hydration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HydrationGroupedResponse); i {
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
			RawDescriptor: file_proto_hydration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hydration_proto_goTypes,
		DependencyIndexes: file_proto_hydration_proto_depIdxs,
		MessageInfos:      file_proto_hydration_proto_msgTypes,
	}.Build()
	File_proto_hydration_proto = out.File
	file_proto_hydration_proto_rawDesc = nil
	file_proto_hydration_proto_goTypes = nil
	file_proto_hydration_proto_depIdxs = nil
}
