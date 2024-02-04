// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.1
// source: proxy/wireguard/config.proto

package wireguard

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

type DeviceConfig_DomainStrategy int32

const (
	DeviceConfig_FORCE_IP   DeviceConfig_DomainStrategy = 0
	DeviceConfig_FORCE_IP4  DeviceConfig_DomainStrategy = 1
	DeviceConfig_FORCE_IP6  DeviceConfig_DomainStrategy = 2
	DeviceConfig_FORCE_IP46 DeviceConfig_DomainStrategy = 3
	DeviceConfig_FORCE_IP64 DeviceConfig_DomainStrategy = 4
)

// Enum value maps for DeviceConfig_DomainStrategy.
var (
	DeviceConfig_DomainStrategy_name = map[int32]string{
		0: "FORCE_IP",
		1: "FORCE_IP4",
		2: "FORCE_IP6",
		3: "FORCE_IP46",
		4: "FORCE_IP64",
	}
	DeviceConfig_DomainStrategy_value = map[string]int32{
		"FORCE_IP":   0,
		"FORCE_IP4":  1,
		"FORCE_IP6":  2,
		"FORCE_IP46": 3,
		"FORCE_IP64": 4,
	}
)

func (x DeviceConfig_DomainStrategy) Enum() *DeviceConfig_DomainStrategy {
	p := new(DeviceConfig_DomainStrategy)
	*p = x
	return p
}

func (x DeviceConfig_DomainStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceConfig_DomainStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_wireguard_config_proto_enumTypes[0].Descriptor()
}

func (DeviceConfig_DomainStrategy) Type() protoreflect.EnumType {
	return &file_proxy_wireguard_config_proto_enumTypes[0]
}

func (x DeviceConfig_DomainStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceConfig_DomainStrategy.Descriptor instead.
func (DeviceConfig_DomainStrategy) EnumDescriptor() ([]byte, []int) {
	return file_proxy_wireguard_config_proto_rawDescGZIP(), []int{1, 0}
}

type PeerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey    string   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PreSharedKey string   `protobuf:"bytes,2,opt,name=pre_shared_key,json=preSharedKey,proto3" json:"pre_shared_key,omitempty"`
	Endpoint     string   `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	KeepAlive    uint32   `protobuf:"varint,4,opt,name=keep_alive,json=keepAlive,proto3" json:"keep_alive,omitempty"`
	AllowedIps   []string `protobuf:"bytes,5,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
}

func (x *PeerConfig) Reset() {
	*x = PeerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_wireguard_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerConfig) ProtoMessage() {}

func (x *PeerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_wireguard_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerConfig.ProtoReflect.Descriptor instead.
func (*PeerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_wireguard_config_proto_rawDescGZIP(), []int{0}
}

func (x *PeerConfig) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *PeerConfig) GetPreSharedKey() string {
	if x != nil {
		return x.PreSharedKey
	}
	return ""
}

func (x *PeerConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *PeerConfig) GetKeepAlive() uint32 {
	if x != nil {
		return x.KeepAlive
	}
	return 0
}

func (x *PeerConfig) GetAllowedIps() []string {
	if x != nil {
		return x.AllowedIps
	}
	return nil
}

type DeviceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretKey      string                      `protobuf:"bytes,1,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Endpoint       []string                    `protobuf:"bytes,2,rep,name=endpoint,proto3" json:"endpoint,omitempty"`
	Peers          []*PeerConfig               `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
	Mtu            int32                       `protobuf:"varint,4,opt,name=mtu,proto3" json:"mtu,omitempty"`
	NumWorkers     int32                       `protobuf:"varint,5,opt,name=num_workers,json=numWorkers,proto3" json:"num_workers,omitempty"`
	Reserved       []byte                      `protobuf:"bytes,6,opt,name=reserved,proto3" json:"reserved,omitempty"`
	DomainStrategy DeviceConfig_DomainStrategy `protobuf:"varint,7,opt,name=domain_strategy,json=domainStrategy,proto3,enum=xray.proxy.wireguard.DeviceConfig_DomainStrategy" json:"domain_strategy,omitempty"`
	IsClient       bool                        `protobuf:"varint,8,opt,name=is_client,json=isClient,proto3" json:"is_client,omitempty"`
	KernelMode     bool                        `protobuf:"varint,9,opt,name=kernel_mode,json=kernelMode,proto3" json:"kernel_mode,omitempty"`
}

func (x *DeviceConfig) Reset() {
	*x = DeviceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_wireguard_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceConfig) ProtoMessage() {}

func (x *DeviceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_wireguard_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceConfig.ProtoReflect.Descriptor instead.
func (*DeviceConfig) Descriptor() ([]byte, []int) {
	return file_proxy_wireguard_config_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *DeviceConfig) GetEndpoint() []string {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *DeviceConfig) GetPeers() []*PeerConfig {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *DeviceConfig) GetMtu() int32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *DeviceConfig) GetNumWorkers() int32 {
	if x != nil {
		return x.NumWorkers
	}
	return 0
}

func (x *DeviceConfig) GetReserved() []byte {
	if x != nil {
		return x.Reserved
	}
	return nil
}

func (x *DeviceConfig) GetDomainStrategy() DeviceConfig_DomainStrategy {
	if x != nil {
		return x.DomainStrategy
	}
	return DeviceConfig_FORCE_IP
}

func (x *DeviceConfig) GetIsClient() bool {
	if x != nil {
		return x.IsClient
	}
	return false
}

func (x *DeviceConfig) GetKernelMode() bool {
	if x != nil {
		return x.KernelMode
	}
	return false
}

var File_proxy_wireguard_config_proto protoreflect.FileDescriptor

var file_proxy_wireguard_config_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69,
	0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x49, 0x70, 0x73, 0x22, 0xc8, 0x03, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x36, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x77, 0x69, 0x72,
	0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75,
	0x6d, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6e, 0x75, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x5a, 0x0a, 0x0f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x31, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x77, 0x69,
	0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x0e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64,
	0x65, 0x22, 0x5c, 0x0a, 0x0e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x36, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x34, 0x10, 0x04, 0x42,
	0x5e, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x50, 0x01, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78,
	0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x77,
	0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0xaa, 0x02, 0x14, 0x58, 0x72, 0x61, 0x79, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x47, 0x75, 0x61, 0x72, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_wireguard_config_proto_rawDescOnce sync.Once
	file_proxy_wireguard_config_proto_rawDescData = file_proxy_wireguard_config_proto_rawDesc
)

func file_proxy_wireguard_config_proto_rawDescGZIP() []byte {
	file_proxy_wireguard_config_proto_rawDescOnce.Do(func() {
		file_proxy_wireguard_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_wireguard_config_proto_rawDescData)
	})
	return file_proxy_wireguard_config_proto_rawDescData
}

var file_proxy_wireguard_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proxy_wireguard_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proxy_wireguard_config_proto_goTypes = []interface{}{
	(DeviceConfig_DomainStrategy)(0), // 0: xray.proxy.wireguard.DeviceConfig.DomainStrategy
	(*PeerConfig)(nil),               // 1: xray.proxy.wireguard.PeerConfig
	(*DeviceConfig)(nil),             // 2: xray.proxy.wireguard.DeviceConfig
}
var file_proxy_wireguard_config_proto_depIdxs = []int32{
	1, // 0: xray.proxy.wireguard.DeviceConfig.peers:type_name -> xray.proxy.wireguard.PeerConfig
	0, // 1: xray.proxy.wireguard.DeviceConfig.domain_strategy:type_name -> xray.proxy.wireguard.DeviceConfig.DomainStrategy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proxy_wireguard_config_proto_init() }
func file_proxy_wireguard_config_proto_init() {
	if File_proxy_wireguard_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_wireguard_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerConfig); i {
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
		file_proxy_wireguard_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceConfig); i {
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
			RawDescriptor: file_proxy_wireguard_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_wireguard_config_proto_goTypes,
		DependencyIndexes: file_proxy_wireguard_config_proto_depIdxs,
		EnumInfos:         file_proxy_wireguard_config_proto_enumTypes,
		MessageInfos:      file_proxy_wireguard_config_proto_msgTypes,
	}.Build()
	File_proxy_wireguard_config_proto = out.File
	file_proxy_wireguard_config_proto_rawDesc = nil
	file_proxy_wireguard_config_proto_goTypes = nil
	file_proxy_wireguard_config_proto_depIdxs = nil
}
