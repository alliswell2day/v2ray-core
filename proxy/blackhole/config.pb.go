package blackhole

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	serial "v2ray.com/core/common/serial"
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

type NoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoneResponse) Reset() {
	*x = NoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoneResponse) ProtoMessage() {}

func (x *NoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoneResponse.ProtoReflect.Descriptor instead.
func (*NoneResponse) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_blackhole_config_proto_rawDescGZIP(), []int{0}
}

type HTTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HTTPResponse) Reset() {
	*x = HTTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPResponse) ProtoMessage() {}

func (x *HTTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPResponse.ProtoReflect.Descriptor instead.
func (*HTTPResponse) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_blackhole_config_proto_rawDescGZIP(), []int{1}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *serial.TypedMessage `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_v2ray_com_core_proxy_blackhole_config_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetResponse() *serial.TypedMessage {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_v2ray_com_core_proxy_blackhole_config_proto protoreflect.FileDescriptor

var file_v2ray_com_core_proxy_blackhole_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x6c, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x6c, 0x65, 0x1a, 0x30, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x4e,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4a, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x6c, 0x65, 0x50, 0x01, 0x5a, 0x09, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x6c, 0x65, 0xaa, 0x02, 0x1a, 0x56, 0x32, 0x52, 0x61, 0x79,
	0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x68, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2ray_com_core_proxy_blackhole_config_proto_rawDescOnce sync.Once
	file_v2ray_com_core_proxy_blackhole_config_proto_rawDescData = file_v2ray_com_core_proxy_blackhole_config_proto_rawDesc
)

func file_v2ray_com_core_proxy_blackhole_config_proto_rawDescGZIP() []byte {
	file_v2ray_com_core_proxy_blackhole_config_proto_rawDescOnce.Do(func() {
		file_v2ray_com_core_proxy_blackhole_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2ray_com_core_proxy_blackhole_config_proto_rawDescData)
	})
	return file_v2ray_com_core_proxy_blackhole_config_proto_rawDescData
}

var file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v2ray_com_core_proxy_blackhole_config_proto_goTypes = []interface{}{
	(*NoneResponse)(nil),        // 0: v2ray.core.proxy.blackhole.NoneResponse
	(*HTTPResponse)(nil),        // 1: v2ray.core.proxy.blackhole.HTTPResponse
	(*Config)(nil),              // 2: v2ray.core.proxy.blackhole.Config
	(*serial.TypedMessage)(nil), // 3: v2ray.core.common.serial.TypedMessage
}
var file_v2ray_com_core_proxy_blackhole_config_proto_depIdxs = []int32{
	3, // 0: v2ray.core.proxy.blackhole.Config.response:type_name -> v2ray.core.common.serial.TypedMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2ray_com_core_proxy_blackhole_config_proto_init() }
func file_v2ray_com_core_proxy_blackhole_config_proto_init() {
	if File_v2ray_com_core_proxy_blackhole_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoneResponse); i {
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
		file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPResponse); i {
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
		file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_v2ray_com_core_proxy_blackhole_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2ray_com_core_proxy_blackhole_config_proto_goTypes,
		DependencyIndexes: file_v2ray_com_core_proxy_blackhole_config_proto_depIdxs,
		MessageInfos:      file_v2ray_com_core_proxy_blackhole_config_proto_msgTypes,
	}.Build()
	File_v2ray_com_core_proxy_blackhole_config_proto = out.File
	file_v2ray_com_core_proxy_blackhole_config_proto_rawDesc = nil
	file_v2ray_com_core_proxy_blackhole_config_proto_goTypes = nil
	file_v2ray_com_core_proxy_blackhole_config_proto_depIdxs = nil
}
