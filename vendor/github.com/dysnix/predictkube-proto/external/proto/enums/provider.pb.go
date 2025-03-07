// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: proto/enums/provider.proto

package enums

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

type PeriodType int32

const (
	PeriodType_Daily   PeriodType = 0
	PeriodType_Weekly  PeriodType = 1
	PeriodType_Monthly PeriodType = 2
)

// Enum value maps for PeriodType.
var (
	PeriodType_name = map[int32]string{
		0: "Daily",
		1: "Weekly",
		2: "Monthly",
	}
	PeriodType_value = map[string]int32{
		"Daily":   0,
		"Weekly":  1,
		"Monthly": 2,
	}
)

func (x PeriodType) Enum() *PeriodType {
	p := new(PeriodType)
	*p = x
	return p
}

func (x PeriodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeriodType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_provider_proto_enumTypes[0].Descriptor()
}

func (PeriodType) Type() protoreflect.EnumType {
	return &file_proto_enums_provider_proto_enumTypes[0]
}

func (x PeriodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeriodType.Descriptor instead.
func (PeriodType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_provider_proto_rawDescGZIP(), []int{0}
}

type MetricsType int32

const (
	MetricsType_Cpu           MetricsType = 0
	MetricsType_Memory        MetricsType = 1
	MetricsType_Disk          MetricsType = 2
	MetricsType_Network       MetricsType = 3
	MetricsType_Logs          MetricsType = 4
	MetricsType_Nginx         MetricsType = 5
	MetricsType_ReplicasCount MetricsType = 6
	MetricsType_Geth          MetricsType = 7
)

// Enum value maps for MetricsType.
var (
	MetricsType_name = map[int32]string{
		0: "Cpu",
		1: "Memory",
		2: "Disk",
		3: "Network",
		4: "Logs",
		5: "Nginx",
		6: "ReplicasCount",
		7: "Geth",
	}
	MetricsType_value = map[string]int32{
		"Cpu":           0,
		"Memory":        1,
		"Disk":          2,
		"Network":       3,
		"Logs":          4,
		"Nginx":         5,
		"ReplicasCount": 6,
		"Geth":          7,
	}
)

func (x MetricsType) Enum() *MetricsType {
	p := new(MetricsType)
	*p = x
	return p
}

func (x MetricsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricsType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_provider_proto_enumTypes[1].Descriptor()
}

func (MetricsType) Type() protoreflect.EnumType {
	return &file_proto_enums_provider_proto_enumTypes[1]
}

func (x MetricsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricsType.Descriptor instead.
func (MetricsType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_provider_proto_rawDescGZIP(), []int{1}
}

type ResourceType int32

const (
	ResourceType_Deployment  ResourceType = 0
	ResourceType_StatefulSet ResourceType = 1
	ResourceType_DaemonSet   ResourceType = 2
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "Deployment",
		1: "StatefulSet",
		2: "DaemonSet",
	}
	ResourceType_value = map[string]int32{
		"Deployment":  0,
		"StatefulSet": 1,
		"DaemonSet":   2,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_provider_proto_enumTypes[2].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_proto_enums_provider_proto_enumTypes[2]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_provider_proto_rawDescGZIP(), []int{2}
}

type ValueType int32

const (
	ValueType_Matrix ValueType = 0
	ValueType_Vector ValueType = 1
	ValueType_Scalar ValueType = 2
	ValueType_String ValueType = 3
)

// Enum value maps for ValueType.
var (
	ValueType_name = map[int32]string{
		0: "Matrix",
		1: "Vector",
		2: "Scalar",
		3: "String",
	}
	ValueType_value = map[string]int32{
		"Matrix": 0,
		"Vector": 1,
		"Scalar": 2,
		"String": 3,
	}
)

func (x ValueType) Enum() *ValueType {
	p := new(ValueType)
	*p = x
	return p
}

func (x ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_provider_proto_enumTypes[3].Descriptor()
}

func (ValueType) Type() protoreflect.EnumType {
	return &file_proto_enums_provider_proto_enumTypes[3]
}

func (x ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValueType.Descriptor instead.
func (ValueType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_provider_proto_rawDescGZIP(), []int{3}
}

type KeyType int32

const (
	KeyType_HS256 KeyType = 0
	KeyType_HS384 KeyType = 1
	KeyType_HS512 KeyType = 2
)

// Enum value maps for KeyType.
var (
	KeyType_name = map[int32]string{
		0: "HS256",
		1: "HS384",
		2: "HS512",
	}
	KeyType_value = map[string]int32{
		"HS256": 0,
		"HS384": 1,
		"HS512": 2,
	}
)

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}

func (x KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_provider_proto_enumTypes[4].Descriptor()
}

func (KeyType) Type() protoreflect.EnumType {
	return &file_proto_enums_provider_proto_enumTypes[4]
}

func (x KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyType.Descriptor instead.
func (KeyType) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_provider_proto_rawDescGZIP(), []int{4}
}

var File_proto_enums_provider_proto protoreflect.FileDescriptor

var file_proto_enums_provider_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2a, 0x30, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x6c, 0x79, 0x10, 0x02, 0x2a, 0x6b, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x69, 0x73,
	0x6b, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x03,
	0x12, 0x08, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x67,
	0x69, 0x6e, 0x78, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x65, 0x74, 0x68,
	0x10, 0x07, 0x2a, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65,
	0x74, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x10, 0x02, 0x2a, 0x3b, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x2a,
	0x2a, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x53,
	0x32, 0x35, 0x36, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x53, 0x33, 0x38, 0x34, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x48, 0x53, 0x35, 0x31, 0x32, 0x10, 0x02, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x79, 0x73, 0x6e, 0x69, 0x78,
	0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_enums_provider_proto_rawDescOnce sync.Once
	file_proto_enums_provider_proto_rawDescData = file_proto_enums_provider_proto_rawDesc
)

func file_proto_enums_provider_proto_rawDescGZIP() []byte {
	file_proto_enums_provider_proto_rawDescOnce.Do(func() {
		file_proto_enums_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_enums_provider_proto_rawDescData)
	})
	return file_proto_enums_provider_proto_rawDescData
}

var file_proto_enums_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_proto_enums_provider_proto_goTypes = []any{
	(PeriodType)(0),   // 0: enums.PeriodType
	(MetricsType)(0),  // 1: enums.MetricsType
	(ResourceType)(0), // 2: enums.ResourceType
	(ValueType)(0),    // 3: enums.ValueType
	(KeyType)(0),      // 4: enums.KeyType
}
var file_proto_enums_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_enums_provider_proto_init() }
func file_proto_enums_provider_proto_init() {
	if File_proto_enums_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_enums_provider_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_enums_provider_proto_goTypes,
		DependencyIndexes: file_proto_enums_provider_proto_depIdxs,
		EnumInfos:         file_proto_enums_provider_proto_enumTypes,
	}.Build()
	File_proto_enums_provider_proto = out.File
	file_proto_enums_provider_proto_rawDesc = nil
	file_proto_enums_provider_proto_goTypes = nil
	file_proto_enums_provider_proto_depIdxs = nil
}
