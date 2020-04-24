// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: oplog_test/oplog_test.proto

// define a test proto package

package oplog_test

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

// TestUser for gorm test user model
type TestUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *TestUser) Reset() {
	*x = TestUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oplog_test_oplog_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestUser) ProtoMessage() {}

func (x *TestUser) ProtoReflect() protoreflect.Message {
	mi := &file_oplog_test_oplog_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestUser.ProtoReflect.Descriptor instead.
func (*TestUser) Descriptor() ([]byte, []int) {
	return file_oplog_test_oplog_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestUser) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *TestUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// TestCar for gorm test car model
type TestCar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	Model string `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Mpg   int32  `protobuf:"varint,5,opt,name=mpg,proto3" json:"mpg,omitempty"`
}

func (x *TestCar) Reset() {
	*x = TestCar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oplog_test_oplog_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCar) ProtoMessage() {}

func (x *TestCar) ProtoReflect() protoreflect.Message {
	mi := &file_oplog_test_oplog_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCar.ProtoReflect.Descriptor instead.
func (*TestCar) Descriptor() ([]byte, []int) {
	return file_oplog_test_oplog_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestCar) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestCar) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *TestCar) GetMpg() int32 {
	if x != nil {
		return x.Mpg
	}
	return 0
}

// TestRental for gorm test rental model
type TestRental struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CarId  uint32 `protobuf:"varint,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *TestRental) Reset() {
	*x = TestRental{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oplog_test_oplog_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRental) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRental) ProtoMessage() {}

func (x *TestRental) ProtoReflect() protoreflect.Message {
	mi := &file_oplog_test_oplog_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRental.ProtoReflect.Descriptor instead.
func (*TestRental) Descriptor() ([]byte, []int) {
	return file_oplog_test_oplog_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestRental) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TestRental) GetCarId() uint32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

var File_oplog_test_oplog_test_proto protoreflect.FileDescriptor

var file_oplog_test_oplog_test_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6f, 0x70, 0x6c,
	0x6f, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x74, 0x6f,
	0x77, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x6f,
	0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x67, 0x0a, 0x08,
	0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x41, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x70, 0x67, 0x22, 0x3c, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2f, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x3b, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oplog_test_oplog_test_proto_rawDescOnce sync.Once
	file_oplog_test_oplog_test_proto_rawDescData = file_oplog_test_oplog_test_proto_rawDesc
)

func file_oplog_test_oplog_test_proto_rawDescGZIP() []byte {
	file_oplog_test_oplog_test_proto_rawDescOnce.Do(func() {
		file_oplog_test_oplog_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_oplog_test_oplog_test_proto_rawDescData)
	})
	return file_oplog_test_oplog_test_proto_rawDescData
}

var file_oplog_test_oplog_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oplog_test_oplog_test_proto_goTypes = []interface{}{
	(*TestUser)(nil),   // 0: hashicorp.watchtower.controller.oplog.test.v1.TestUser
	(*TestCar)(nil),    // 1: hashicorp.watchtower.controller.oplog.test.v1.TestCar
	(*TestRental)(nil), // 2: hashicorp.watchtower.controller.oplog.test.v1.TestRental
}
var file_oplog_test_oplog_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oplog_test_oplog_test_proto_init() }
func file_oplog_test_oplog_test_proto_init() {
	if File_oplog_test_oplog_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oplog_test_oplog_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestUser); i {
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
		file_oplog_test_oplog_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCar); i {
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
		file_oplog_test_oplog_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRental); i {
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
			RawDescriptor: file_oplog_test_oplog_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oplog_test_oplog_test_proto_goTypes,
		DependencyIndexes: file_oplog_test_oplog_test_proto_depIdxs,
		MessageInfos:      file_oplog_test_oplog_test_proto_msgTypes,
	}.Build()
	File_oplog_test_oplog_test_proto = out.File
	file_oplog_test_oplog_test_proto_rawDesc = nil
	file_oplog_test_oplog_test_proto_goTypes = nil
	file_oplog_test_oplog_test_proto_depIdxs = nil
}
