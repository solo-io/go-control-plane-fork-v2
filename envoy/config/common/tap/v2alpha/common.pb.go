// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: envoy/config/common/tap/v2alpha/common.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/service/tap/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Common configuration for all tap extensions.
type CommonExtensionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	ConfigType isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
}

func (x *CommonExtensionConfig) Reset() {
	*x = CommonExtensionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_tap_v2alpha_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonExtensionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonExtensionConfig) ProtoMessage() {}

func (x *CommonExtensionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_tap_v2alpha_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonExtensionConfig.ProtoReflect.Descriptor instead.
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_tap_v2alpha_common_proto_rawDescGZIP(), []int{0}
}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := x.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (x *CommonExtensionConfig) GetStaticConfig() *v2alpha.TapConfig {
	if x, ok := x.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
}

type CommonExtensionConfig_AdminConfig struct {
	// If specified, the tap filter will be configured via an admin handler.
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof"`
}

type CommonExtensionConfig_StaticConfig struct {
	// If specified, the tap filter will be configured via a static configuration that cannot be
	// changed.
	StaticConfig *v2alpha.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}

// Configuration for the admin handler. See :ref:`here <config_http_filters_tap_admin_handler>` for
// more information.
type AdminConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque configuration ID. When requests are made to the admin handler, the passed opaque ID is
	// matched to the configured filter opaque ID to determine which filter to configure.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *AdminConfig) Reset() {
	*x = AdminConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_tap_v2alpha_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminConfig) ProtoMessage() {}

func (x *AdminConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_tap_v2alpha_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminConfig.ProtoReflect.Descriptor instead.
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_tap_v2alpha_common_proto_rawDescGZIP(), []int{1}
}

func (x *AdminConfig) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

var File_envoy_config_common_tap_v2alpha_common_proto protoreflect.FileDescriptor

var file_envoy_config_common_tap_v2alpha_common_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a,
	0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74,
	0x61, 0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcb, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x0c, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x54, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x12, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x33, 0x0a,
	0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x42, 0xb4, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x20, 0x12, 0x1e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x01, 0x0a, 0x2d, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x61,
	0x70, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_config_common_tap_v2alpha_common_proto_rawDescOnce sync.Once
	file_envoy_config_common_tap_v2alpha_common_proto_rawDescData = file_envoy_config_common_tap_v2alpha_common_proto_rawDesc
)

func file_envoy_config_common_tap_v2alpha_common_proto_rawDescGZIP() []byte {
	file_envoy_config_common_tap_v2alpha_common_proto_rawDescOnce.Do(func() {
		file_envoy_config_common_tap_v2alpha_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_common_tap_v2alpha_common_proto_rawDescData)
	})
	return file_envoy_config_common_tap_v2alpha_common_proto_rawDescData
}

var file_envoy_config_common_tap_v2alpha_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_common_tap_v2alpha_common_proto_goTypes = []interface{}{
	(*CommonExtensionConfig)(nil), // 0: envoy.config.common.tap.v2alpha.CommonExtensionConfig
	(*AdminConfig)(nil),           // 1: envoy.config.common.tap.v2alpha.AdminConfig
	(*v2alpha.TapConfig)(nil),     // 2: envoy.service.tap.v2alpha.TapConfig
}
var file_envoy_config_common_tap_v2alpha_common_proto_depIdxs = []int32{
	1, // 0: envoy.config.common.tap.v2alpha.CommonExtensionConfig.admin_config:type_name -> envoy.config.common.tap.v2alpha.AdminConfig
	2, // 1: envoy.config.common.tap.v2alpha.CommonExtensionConfig.static_config:type_name -> envoy.service.tap.v2alpha.TapConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_common_tap_v2alpha_common_proto_init() }
func file_envoy_config_common_tap_v2alpha_common_proto_init() {
	if File_envoy_config_common_tap_v2alpha_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_common_tap_v2alpha_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonExtensionConfig); i {
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
		file_envoy_config_common_tap_v2alpha_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminConfig); i {
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
	file_envoy_config_common_tap_v2alpha_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_common_tap_v2alpha_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_common_tap_v2alpha_common_proto_goTypes,
		DependencyIndexes: file_envoy_config_common_tap_v2alpha_common_proto_depIdxs,
		MessageInfos:      file_envoy_config_common_tap_v2alpha_common_proto_msgTypes,
	}.Build()
	File_envoy_config_common_tap_v2alpha_common_proto = out.File
	file_envoy_config_common_tap_v2alpha_common_proto_rawDesc = nil
	file_envoy_config_common_tap_v2alpha_common_proto_goTypes = nil
	file_envoy_config_common_tap_v2alpha_common_proto_depIdxs = nil
}
