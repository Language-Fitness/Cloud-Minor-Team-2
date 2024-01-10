// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: SagaActions.proto

package pb

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

// Enum for SagaObjectTypes
type SagaObjectType int32

const (
	SagaObjectType_SCHOOL   SagaObjectType = 0
	SagaObjectType_CLASS    SagaObjectType = 1
	SagaObjectType_EXERCISE SagaObjectType = 2
	SagaObjectType_RESULT   SagaObjectType = 3
	SagaObjectType_MODULE   SagaObjectType = 4
)

// Enum value maps for SagaObjectType.
var (
	SagaObjectType_name = map[int32]string{
		0: "SCHOOL",
		1: "CLASS",
		2: "EXERCISE",
		3: "RESULT",
		4: "MODULE",
	}
	SagaObjectType_value = map[string]int32{
		"SCHOOL":   0,
		"CLASS":    1,
		"EXERCISE": 2,
		"RESULT":   3,
		"MODULE":   4,
	}
)

func (x SagaObjectType) Enum() *SagaObjectType {
	p := new(SagaObjectType)
	*p = x
	return p
}

func (x SagaObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SagaObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_SagaActions_proto_enumTypes[0].Descriptor()
}

func (SagaObjectType) Type() protoreflect.EnumType {
	return &file_SagaActions_proto_enumTypes[0]
}

func (x SagaObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SagaObjectType.Descriptor instead.
func (SagaObjectType) EnumDescriptor() ([]byte, []int) {
	return file_SagaActions_proto_rawDescGZIP(), []int{0}
}

// Enum for SagaObjectStatus
type SagaObjectStatus int32

const (
	SagaObjectStatus_EXIST   SagaObjectStatus = 0
	SagaObjectStatus_DELETED SagaObjectStatus = 1
)

// Enum value maps for SagaObjectStatus.
var (
	SagaObjectStatus_name = map[int32]string{
		0: "EXIST",
		1: "DELETED",
	}
	SagaObjectStatus_value = map[string]int32{
		"EXIST":   0,
		"DELETED": 1,
	}
)

func (x SagaObjectStatus) Enum() *SagaObjectStatus {
	p := new(SagaObjectStatus)
	*p = x
	return p
}

func (x SagaObjectStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SagaObjectStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_SagaActions_proto_enumTypes[1].Descriptor()
}

func (SagaObjectStatus) Type() protoreflect.EnumType {
	return &file_SagaActions_proto_enumTypes[1]
}

func (x SagaObjectStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SagaObjectStatus.Descriptor instead.
func (SagaObjectStatus) EnumDescriptor() ([]byte, []int) {
	return file_SagaActions_proto_rawDescGZIP(), []int{1}
}

// Request message for getting a user
type ObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId     string           `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	BearerToken  string           `protobuf:"bytes,2,opt,name=bearer_token,json=bearerToken,proto3" json:"bearer_token,omitempty"`
	ObjectType   SagaObjectType   `protobuf:"varint,3,opt,name=object_type,json=objectType,proto3,enum=proto.SagaObjectType" json:"object_type,omitempty"`
	ObjectStatus SagaObjectStatus `protobuf:"varint,4,opt,name=object_status,json=objectStatus,proto3,enum=proto.SagaObjectStatus" json:"object_status,omitempty"`
}

func (x *ObjectRequest) Reset() {
	*x = ObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SagaActions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRequest) ProtoMessage() {}

func (x *ObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_SagaActions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRequest.ProtoReflect.Descriptor instead.
func (*ObjectRequest) Descriptor() ([]byte, []int) {
	return file_SagaActions_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *ObjectRequest) GetBearerToken() string {
	if x != nil {
		return x.BearerToken
	}
	return ""
}

func (x *ObjectRequest) GetObjectType() SagaObjectType {
	if x != nil {
		return x.ObjectType
	}
	return SagaObjectType_SCHOOL
}

func (x *ObjectRequest) GetObjectStatus() SagaObjectStatus {
	if x != nil {
		return x.ObjectStatus
	}
	return SagaObjectStatus_EXIST
}

type ObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*SagaObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ObjectResponse) Reset() {
	*x = ObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SagaActions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectResponse) ProtoMessage() {}

func (x *ObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_SagaActions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectResponse.ProtoReflect.Descriptor instead.
func (*ObjectResponse) Descriptor() ([]byte, []int) {
	return file_SagaActions_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectResponse) GetObjects() []*SagaObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

type SagaObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId     string           `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	ObjectType   SagaObjectType   `protobuf:"varint,2,opt,name=object_type,json=objectType,proto3,enum=proto.SagaObjectType" json:"object_type,omitempty"`
	ObjectStatus SagaObjectStatus `protobuf:"varint,3,opt,name=object_status,json=objectStatus,proto3,enum=proto.SagaObjectStatus" json:"object_status,omitempty"`
}

func (x *SagaObject) Reset() {
	*x = SagaObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SagaActions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SagaObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SagaObject) ProtoMessage() {}

func (x *SagaObject) ProtoReflect() protoreflect.Message {
	mi := &file_SagaActions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SagaObject.ProtoReflect.Descriptor instead.
func (*SagaObject) Descriptor() ([]byte, []int) {
	return file_SagaActions_proto_rawDescGZIP(), []int{2}
}

func (x *SagaObject) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *SagaObject) GetObjectType() SagaObjectType {
	if x != nil {
		return x.ObjectType
	}
	return SagaObjectType_SCHOOL
}

func (x *SagaObject) GetObjectStatus() SagaObjectStatus {
	if x != nil {
		return x.ObjectStatus
	}
	return SagaObjectStatus_EXIST
}

var File_SagaActions_proto protoreflect.FileDescriptor

var file_SagaActions_proto_rawDesc = []byte{
	0x0a, 0x11, 0x53, 0x61, 0x67, 0x61, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x0d, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x0b,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x3d, 0x0a, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61,
	0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0a, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2a, 0x4d, 0x0a, 0x0e, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x45, 0x58, 0x45, 0x52, 0x43, 0x49, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45,
	0x10, 0x04, 0x2a, 0x2a, 0x0a, 0x10, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x32, 0x8f,
	0x02, 0x0a, 0x0f, 0x67, 0x52, 0x50, 0x43, 0x53, 0x61, 0x67, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x45, 0x0a,
	0x16, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x61, 0x67, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x55, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SagaActions_proto_rawDescOnce sync.Once
	file_SagaActions_proto_rawDescData = file_SagaActions_proto_rawDesc
)

func file_SagaActions_proto_rawDescGZIP() []byte {
	file_SagaActions_proto_rawDescOnce.Do(func() {
		file_SagaActions_proto_rawDescData = protoimpl.X.CompressGZIP(file_SagaActions_proto_rawDescData)
	})
	return file_SagaActions_proto_rawDescData
}

var file_SagaActions_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_SagaActions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_SagaActions_proto_goTypes = []interface{}{
	(SagaObjectType)(0),    // 0: proto.SagaObjectType
	(SagaObjectStatus)(0),  // 1: proto.SagaObjectStatus
	(*ObjectRequest)(nil),  // 2: proto.ObjectRequest
	(*ObjectResponse)(nil), // 3: proto.ObjectResponse
	(*SagaObject)(nil),     // 4: proto.SagaObject
}
var file_SagaActions_proto_depIdxs = []int32{
	0, // 0: proto.ObjectRequest.object_type:type_name -> proto.SagaObjectType
	1, // 1: proto.ObjectRequest.object_status:type_name -> proto.SagaObjectStatus
	4, // 2: proto.ObjectResponse.objects:type_name -> proto.SagaObject
	0, // 3: proto.SagaObject.object_type:type_name -> proto.SagaObjectType
	1, // 4: proto.SagaObject.object_status:type_name -> proto.SagaObjectStatus
	2, // 5: proto.gRPCSagaService.FindSagaObject:input_type -> proto.ObjectRequest
	2, // 6: proto.gRPCSagaService.FindSagaObjectChildren:input_type -> proto.ObjectRequest
	2, // 7: proto.gRPCSagaService.DeleteObject:input_type -> proto.ObjectRequest
	2, // 8: proto.gRPCSagaService.UnDeleteObject:input_type -> proto.ObjectRequest
	4, // 9: proto.gRPCSagaService.FindSagaObject:output_type -> proto.SagaObject
	3, // 10: proto.gRPCSagaService.FindSagaObjectChildren:output_type -> proto.ObjectResponse
	3, // 11: proto.gRPCSagaService.DeleteObject:output_type -> proto.ObjectResponse
	3, // 12: proto.gRPCSagaService.UnDeleteObject:output_type -> proto.ObjectResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_SagaActions_proto_init() }
func file_SagaActions_proto_init() {
	if File_SagaActions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SagaActions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectRequest); i {
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
		file_SagaActions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectResponse); i {
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
		file_SagaActions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SagaObject); i {
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
			RawDescriptor: file_SagaActions_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_SagaActions_proto_goTypes,
		DependencyIndexes: file_SagaActions_proto_depIdxs,
		EnumInfos:         file_SagaActions_proto_enumTypes,
		MessageInfos:      file_SagaActions_proto_msgTypes,
	}.Build()
	File_SagaActions_proto = out.File
	file_SagaActions_proto_rawDesc = nil
	file_SagaActions_proto_goTypes = nil
	file_SagaActions_proto_depIdxs = nil
}