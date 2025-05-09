// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: spotify/connectstate/devices/connect_devices.proto

package devices

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

type DeviceType int32

const (
	DeviceType_UNKNOWN         DeviceType = 0
	DeviceType_COMPUTER        DeviceType = 1
	DeviceType_TABLET          DeviceType = 2
	DeviceType_SMARTPHONE      DeviceType = 3
	DeviceType_SPEAKER         DeviceType = 4
	DeviceType_TV              DeviceType = 5
	DeviceType_AVR             DeviceType = 6
	DeviceType_STB             DeviceType = 7
	DeviceType_AUDIO_DONGLE    DeviceType = 8
	DeviceType_GAME_CONSOLE    DeviceType = 9
	DeviceType_CAST_VIDEO      DeviceType = 10
	DeviceType_CAST_AUDIO      DeviceType = 11
	DeviceType_AUTOMOBILE      DeviceType = 12
	DeviceType_SMARTWATCH      DeviceType = 13
	DeviceType_CHROMEBOOK      DeviceType = 14
	DeviceType_UNKNOWN_SPOTIFY DeviceType = 100
	DeviceType_CAR_THING       DeviceType = 101
	DeviceType_OBSERVER        DeviceType = 102
	DeviceType_HOME_THING      DeviceType = 103
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "COMPUTER",
		2:   "TABLET",
		3:   "SMARTPHONE",
		4:   "SPEAKER",
		5:   "TV",
		6:   "AVR",
		7:   "STB",
		8:   "AUDIO_DONGLE",
		9:   "GAME_CONSOLE",
		10:  "CAST_VIDEO",
		11:  "CAST_AUDIO",
		12:  "AUTOMOBILE",
		13:  "SMARTWATCH",
		14:  "CHROMEBOOK",
		100: "UNKNOWN_SPOTIFY",
		101: "CAR_THING",
		102: "OBSERVER",
		103: "HOME_THING",
	}
	DeviceType_value = map[string]int32{
		"UNKNOWN":         0,
		"COMPUTER":        1,
		"TABLET":          2,
		"SMARTPHONE":      3,
		"SPEAKER":         4,
		"TV":              5,
		"AVR":             6,
		"STB":             7,
		"AUDIO_DONGLE":    8,
		"GAME_CONSOLE":    9,
		"CAST_VIDEO":      10,
		"CAST_AUDIO":      11,
		"AUTOMOBILE":      12,
		"SMARTWATCH":      13,
		"CHROMEBOOK":      14,
		"UNKNOWN_SPOTIFY": 100,
		"CAR_THING":       101,
		"OBSERVER":        102,
		"HOME_THING":      103,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_spotify_connectstate_devices_connect_devices_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_spotify_connectstate_devices_connect_devices_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_spotify_connectstate_devices_connect_devices_proto_rawDescGZIP(), []int{0}
}

type DeviceAlias struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName   string                 `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	IsGroup       bool                   `protobuf:"varint,3,opt,name=is_group,json=isGroup,proto3" json:"is_group,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceAlias) Reset() {
	*x = DeviceAlias{}
	mi := &file_spotify_connectstate_devices_connect_devices_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceAlias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceAlias) ProtoMessage() {}

func (x *DeviceAlias) ProtoReflect() protoreflect.Message {
	mi := &file_spotify_connectstate_devices_connect_devices_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceAlias.ProtoReflect.Descriptor instead.
func (*DeviceAlias) Descriptor() ([]byte, []int) {
	return file_spotify_connectstate_devices_connect_devices_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceAlias) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeviceAlias) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DeviceAlias) GetIsGroup() bool {
	if x != nil {
		return x.IsGroup
	}
	return false
}

var File_spotify_connectstate_devices_connect_devices_proto protoreflect.FileDescriptor

var file_spotify_connectstate_devices_connect_devices_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x22, 0x5b, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2a,
	0xa0, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x50, 0x48,
	0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x45, 0x41, 0x4b, 0x45, 0x52,
	0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x54, 0x56, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x56,
	0x52, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x54, 0x42, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c,
	0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x44, 0x4f, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x08, 0x12, 0x10,
	0x0a, 0x0c, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x4f, 0x4c, 0x45, 0x10, 0x09,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x0a,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x0b,
	0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x0c,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x57, 0x41, 0x54, 0x43, 0x48, 0x10, 0x0d,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x52, 0x4f, 0x4d, 0x45, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x0e,
	0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x50, 0x4f, 0x54,
	0x49, 0x46, 0x59, 0x10, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x52, 0x5f, 0x54, 0x48, 0x49,
	0x4e, 0x47, 0x10, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x10, 0x66, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x54, 0x48, 0x49, 0x4e, 0x47,
	0x10, 0x67, 0x42, 0x8f, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x13, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x67, 0x69,
	0x61, 0x6e, 0x6c, 0x75, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x53, 0x43, 0x44, 0xaa, 0x02, 0x1c, 0x53, 0x70, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x1c, 0x53, 0x70, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x28, 0x53, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5c,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x3a, 0x3a, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x73, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spotify_connectstate_devices_connect_devices_proto_rawDescOnce sync.Once
	file_spotify_connectstate_devices_connect_devices_proto_rawDescData = file_spotify_connectstate_devices_connect_devices_proto_rawDesc
)

func file_spotify_connectstate_devices_connect_devices_proto_rawDescGZIP() []byte {
	file_spotify_connectstate_devices_connect_devices_proto_rawDescOnce.Do(func() {
		file_spotify_connectstate_devices_connect_devices_proto_rawDescData = protoimpl.X.CompressGZIP(file_spotify_connectstate_devices_connect_devices_proto_rawDescData)
	})
	return file_spotify_connectstate_devices_connect_devices_proto_rawDescData
}

var file_spotify_connectstate_devices_connect_devices_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spotify_connectstate_devices_connect_devices_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spotify_connectstate_devices_connect_devices_proto_goTypes = []any{
	(DeviceType)(0),     // 0: spotify.connectstate.devices.DeviceType
	(*DeviceAlias)(nil), // 1: spotify.connectstate.devices.DeviceAlias
}
var file_spotify_connectstate_devices_connect_devices_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spotify_connectstate_devices_connect_devices_proto_init() }
func file_spotify_connectstate_devices_connect_devices_proto_init() {
	if File_spotify_connectstate_devices_connect_devices_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spotify_connectstate_devices_connect_devices_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spotify_connectstate_devices_connect_devices_proto_goTypes,
		DependencyIndexes: file_spotify_connectstate_devices_connect_devices_proto_depIdxs,
		EnumInfos:         file_spotify_connectstate_devices_connect_devices_proto_enumTypes,
		MessageInfos:      file_spotify_connectstate_devices_connect_devices_proto_msgTypes,
	}.Build()
	File_spotify_connectstate_devices_connect_devices_proto = out.File
	file_spotify_connectstate_devices_connect_devices_proto_rawDesc = nil
	file_spotify_connectstate_devices_connect_devices_proto_goTypes = nil
	file_spotify_connectstate_devices_connect_devices_proto_depIdxs = nil
}
