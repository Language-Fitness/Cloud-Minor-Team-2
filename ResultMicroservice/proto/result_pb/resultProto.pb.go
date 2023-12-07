// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: resultProto.proto

package result_pb

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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Deleted bool   `protobuf:"varint,2,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resultProto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_resultProto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_resultProto_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Response) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type DeleteByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	BearerToken string `protobuf:"bytes,2,opt,name=BearerToken,proto3" json:"BearerToken,omitempty"`
}

func (x *DeleteByUserRequest) Reset() {
	*x = DeleteByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resultProto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByUserRequest) ProtoMessage() {}

func (x *DeleteByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resultProto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteByUserRequest) Descriptor() ([]byte, []int) {
	return file_resultProto_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteByUserRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *DeleteByUserRequest) GetBearerToken() string {
	if x != nil {
		return x.BearerToken
	}
	return ""
}

type DeleteByClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassID     string `protobuf:"bytes,1,opt,name=ClassID,proto3" json:"ClassID,omitempty"`
	BearerToken string `protobuf:"bytes,2,opt,name=BearerToken,proto3" json:"BearerToken,omitempty"`
}

func (x *DeleteByClassRequest) Reset() {
	*x = DeleteByClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resultProto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByClassRequest) ProtoMessage() {}

func (x *DeleteByClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resultProto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByClassRequest.ProtoReflect.Descriptor instead.
func (*DeleteByClassRequest) Descriptor() ([]byte, []int) {
	return file_resultProto_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteByClassRequest) GetClassID() string {
	if x != nil {
		return x.ClassID
	}
	return ""
}

func (x *DeleteByClassRequest) GetBearerToken() string {
	if x != nil {
		return x.BearerToken
	}
	return ""
}

type DeleteByModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleID    string `protobuf:"bytes,1,opt,name=ModuleID,proto3" json:"ModuleID,omitempty"`
	BearerToken string `protobuf:"bytes,2,opt,name=BearerToken,proto3" json:"BearerToken,omitempty"`
}

func (x *DeleteByModuleRequest) Reset() {
	*x = DeleteByModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resultProto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteByModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteByModuleRequest) ProtoMessage() {}

func (x *DeleteByModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resultProto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteByModuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteByModuleRequest) Descriptor() ([]byte, []int) {
	return file_resultProto_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteByModuleRequest) GetModuleID() string {
	if x != nil {
		return x.ModuleID
	}
	return ""
}

func (x *DeleteByModuleRequest) GetBearerToken() string {
	if x != nil {
		return x.BearerToken
	}
	return ""
}

var File_resultProto_proto protoreflect.FileDescriptor

var file_resultProto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x4f, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x52, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x55,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xca, 0x02, 0x0a, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x12, 0x53, 0x6f, 0x66,
	0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x11, 0x53, 0x6f, 0x66, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x10, 0x53, 0x6f, 0x66, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resultProto_proto_rawDescOnce sync.Once
	file_resultProto_proto_rawDescData = file_resultProto_proto_rawDesc
)

func file_resultProto_proto_rawDescGZIP() []byte {
	file_resultProto_proto_rawDescOnce.Do(func() {
		file_resultProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_resultProto_proto_rawDescData)
	})
	return file_resultProto_proto_rawDescData
}

var file_resultProto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_resultProto_proto_goTypes = []interface{}{
	(*Response)(nil),              // 0: Response
	(*DeleteByUserRequest)(nil),   // 1: DeleteByUserRequest
	(*DeleteByClassRequest)(nil),  // 2: DeleteByClassRequest
	(*DeleteByModuleRequest)(nil), // 3: DeleteByModuleRequest
}
var file_resultProto_proto_depIdxs = []int32{
	3, // 0: GrpcResult.DeleteByModule:input_type -> DeleteByModuleRequest
	3, // 1: GrpcResult.SoftDeleteByModule:input_type -> DeleteByModuleRequest
	2, // 2: GrpcResult.DeleteByClass:input_type -> DeleteByClassRequest
	2, // 3: GrpcResult.SoftDeleteByClass:input_type -> DeleteByClassRequest
	1, // 4: GrpcResult.DeleteByUser:input_type -> DeleteByUserRequest
	1, // 5: GrpcResult.SoftDeleteByUser:input_type -> DeleteByUserRequest
	0, // 6: GrpcResult.DeleteByModule:output_type -> Response
	0, // 7: GrpcResult.SoftDeleteByModule:output_type -> Response
	0, // 8: GrpcResult.DeleteByClass:output_type -> Response
	0, // 9: GrpcResult.SoftDeleteByClass:output_type -> Response
	0, // 10: GrpcResult.DeleteByUser:output_type -> Response
	0, // 11: GrpcResult.SoftDeleteByUser:output_type -> Response
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resultProto_proto_init() }
func file_resultProto_proto_init() {
	if File_resultProto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resultProto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_resultProto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByUserRequest); i {
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
		file_resultProto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByClassRequest); i {
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
		file_resultProto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteByModuleRequest); i {
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
			RawDescriptor: file_resultProto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resultProto_proto_goTypes,
		DependencyIndexes: file_resultProto_proto_depIdxs,
		MessageInfos:      file_resultProto_proto_msgTypes,
	}.Build()
	File_resultProto_proto = out.File
	file_resultProto_proto_rawDesc = nil
	file_resultProto_proto_goTypes = nil
	file_resultProto_proto_depIdxs = nil
}
