// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: urlservice/proto/service.proto

package proto

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

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlservice_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_urlservice_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_urlservice_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *Origin) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Short struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Short) Reset() {
	*x = Short{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlservice_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Short) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Short) ProtoMessage() {}

func (x *Short) ProtoReflect() protoreflect.Message {
	mi := &file_urlservice_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Short.ProtoReflect.Descriptor instead.
func (*Short) Descriptor() ([]byte, []int) {
	return file_urlservice_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *Short) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short     string                 `protobuf:"bytes,1,opt,name=short,proto3" json:"short,omitempty"`
	Origin    string                 `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Views     uint64                 `protobuf:"varint,4,opt,name=views,proto3" json:"views,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlservice_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_urlservice_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_urlservice_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *URL) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *URL) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *URL) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *URL) GetViews() uint64 {
	if x != nil {
		return x.Views
	}
	return 0
}

var File_urlservice_proto_service_proto protoreflect.FileDescriptor

var file_urlservice_proto_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x75, 0x72, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1a, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x17, 0x0a,
	0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x32, 0x5f, 0x0a, 0x0a,
	0x55, 0x52, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x07, 0x2e, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x1a, 0x06, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x06, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x1a,
	0x07, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x06, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x1a, 0x04, 0x2e, 0x55, 0x52, 0x4c, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x72, 0x61,
	0x6c, 0x65, 0x66, 0x2f, 0x73, 0x68, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_urlservice_proto_service_proto_rawDescOnce sync.Once
	file_urlservice_proto_service_proto_rawDescData = file_urlservice_proto_service_proto_rawDesc
)

func file_urlservice_proto_service_proto_rawDescGZIP() []byte {
	file_urlservice_proto_service_proto_rawDescOnce.Do(func() {
		file_urlservice_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_urlservice_proto_service_proto_rawDescData)
	})
	return file_urlservice_proto_service_proto_rawDescData
}

var file_urlservice_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_urlservice_proto_service_proto_goTypes = []interface{}{
	(*Origin)(nil),                // 0: Origin
	(*Short)(nil),                 // 1: Short
	(*URL)(nil),                   // 2: URL
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_urlservice_proto_service_proto_depIdxs = []int32{
	3, // 0: URL.createdAt:type_name -> google.protobuf.Timestamp
	0, // 1: URLService.CreateShort:input_type -> Origin
	1, // 2: URLService.GetOrigin:input_type -> Short
	1, // 3: URLService.Get:input_type -> Short
	1, // 4: URLService.CreateShort:output_type -> Short
	0, // 5: URLService.GetOrigin:output_type -> Origin
	2, // 6: URLService.Get:output_type -> URL
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_urlservice_proto_service_proto_init() }
func file_urlservice_proto_service_proto_init() {
	if File_urlservice_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_urlservice_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
		file_urlservice_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Short); i {
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
		file_urlservice_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
			RawDescriptor: file_urlservice_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_urlservice_proto_service_proto_goTypes,
		DependencyIndexes: file_urlservice_proto_service_proto_depIdxs,
		MessageInfos:      file_urlservice_proto_service_proto_msgTypes,
	}.Build()
	File_urlservice_proto_service_proto = out.File
	file_urlservice_proto_service_proto_rawDesc = nil
	file_urlservice_proto_service_proto_goTypes = nil
	file_urlservice_proto_service_proto_depIdxs = nil
}
