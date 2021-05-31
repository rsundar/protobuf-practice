// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: simple.proto

package simple

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

type City struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ZipCode     string `protobuf:"bytes,2,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	CountryName string `protobuf:"bytes,3,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
}

func (x *City) Reset() {
	*x = City{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *City) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*City) ProtoMessage() {}

func (x *City) ProtoReflect() protoreflect.Message {
	mi := &file_simple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use City.ProtoReflect.Descriptor instead.
func (*City) Descriptor() ([]byte, []int) {
	return file_simple_proto_rawDescGZIP(), []int{0}
}

func (x *City) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *City) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *City) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

type Street struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetName string `protobuf:"bytes,1,opt,name=street_name,json=streetName,proto3" json:"street_name,omitempty"`
	City       *City  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *Street) Reset() {
	*x = Street{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Street) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Street) ProtoMessage() {}

func (x *Street) ProtoReflect() protoreflect.Message {
	mi := &file_simple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Street.ProtoReflect.Descriptor instead.
func (*Street) Descriptor() ([]byte, []int) {
	return file_simple_proto_rawDescGZIP(), []int{1}
}

func (x *Street) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *Street) GetCity() *City {
	if x != nil {
		return x.City
	}
	return nil
}

type Building struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildingName   string  `protobuf:"bytes,1,opt,name=building_name,json=buildingName,proto3" json:"building_name,omitempty"`
	BuildingNumber string  `protobuf:"bytes,2,opt,name=building_number,json=buildingNumber,proto3" json:"building_number,omitempty"`
	Street         *Street `protobuf:"bytes,3,opt,name=street,proto3" json:"street,omitempty"`
}

func (x *Building) Reset() {
	*x = Building{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Building) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Building) ProtoMessage() {}

func (x *Building) ProtoReflect() protoreflect.Message {
	mi := &file_simple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Building.ProtoReflect.Descriptor instead.
func (*Building) Descriptor() ([]byte, []int) {
	return file_simple_proto_rawDescGZIP(), []int{2}
}

func (x *Building) GetBuildingName() string {
	if x != nil {
		return x.BuildingName
	}
	return ""
}

func (x *Building) GetBuildingNumber() string {
	if x != nil {
		return x.BuildingNumber
	}
	return ""
}

func (x *Building) GetStreet() *Street {
	if x != nil {
		return x.Street
	}
	return nil
}

var File_simple_proto protoreflect.FileDescriptor

var file_simple_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58,
	0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69,
	0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x79,
	0x0a, 0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x42, 0x16, 0x5a, 0x14, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simple_proto_rawDescOnce sync.Once
	file_simple_proto_rawDescData = file_simple_proto_rawDesc
)

func file_simple_proto_rawDescGZIP() []byte {
	file_simple_proto_rawDescOnce.Do(func() {
		file_simple_proto_rawDescData = protoimpl.X.CompressGZIP(file_simple_proto_rawDescData)
	})
	return file_simple_proto_rawDescData
}

var file_simple_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_simple_proto_goTypes = []interface{}{
	(*City)(nil),     // 0: City
	(*Street)(nil),   // 1: Street
	(*Building)(nil), // 2: Building
}
var file_simple_proto_depIdxs = []int32{
	0, // 0: Street.city:type_name -> City
	1, // 1: Building.street:type_name -> Street
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_simple_proto_init() }
func file_simple_proto_init() {
	if File_simple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*City); i {
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
		file_simple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Street); i {
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
		file_simple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Building); i {
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
			RawDescriptor: file_simple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_simple_proto_goTypes,
		DependencyIndexes: file_simple_proto_depIdxs,
		MessageInfos:      file_simple_proto_msgTypes,
	}.Build()
	File_simple_proto = out.File
	file_simple_proto_rawDesc = nil
	file_simple_proto_goTypes = nil
	file_simple_proto_depIdxs = nil
}
