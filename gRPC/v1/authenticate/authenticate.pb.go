// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: gRPC/v1/authenticate/authenticate.proto

package authenticate

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_v1_authenticate_authenticate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_v1_authenticate_authenticate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gRPC_v1_authenticate_authenticate_proto_rawDescGZIP(), []int{0}
}

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId string `protobuf:"bytes,1,opt,name=ChallengeId,proto3" json:"ChallengeId,omitempty"`
	Signature   string `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_v1_authenticate_authenticate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_v1_authenticate_authenticate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_gRPC_v1_authenticate_authenticate_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateRequest) GetChallengeId() string {
	if x != nil {
		return x.ChallengeId
	}
	return ""
}

func (x *AuthenticateRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type AuthenticatePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int64  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Token   string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AuthenticatePayload) Reset() {
	*x = AuthenticatePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_v1_authenticate_authenticate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatePayload) ProtoMessage() {}

func (x *AuthenticatePayload) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_v1_authenticate_authenticate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatePayload.ProtoReflect.Descriptor instead.
func (*AuthenticatePayload) Descriptor() ([]byte, []int) {
	return file_gRPC_v1_authenticate_authenticate_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticatePayload) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuthenticatePayload) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AuthenticatePayload) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuthenticatePayload) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_gRPC_v1_authenticate_authenticate_proto protoreflect.FileDescriptor

var file_gRPC_v1_authenticate_authenticate_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x55, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x77, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x32, 0x6b, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x4c,
	0x61, 0x7a, 0x61, 0x72, 0x75, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x52,
	0x50, 0x43, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_v1_authenticate_authenticate_proto_rawDescOnce sync.Once
	file_gRPC_v1_authenticate_authenticate_proto_rawDescData = file_gRPC_v1_authenticate_authenticate_proto_rawDesc
)

func file_gRPC_v1_authenticate_authenticate_proto_rawDescGZIP() []byte {
	file_gRPC_v1_authenticate_authenticate_proto_rawDescOnce.Do(func() {
		file_gRPC_v1_authenticate_authenticate_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_v1_authenticate_authenticate_proto_rawDescData)
	})
	return file_gRPC_v1_authenticate_authenticate_proto_rawDescData
}

var file_gRPC_v1_authenticate_authenticate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gRPC_v1_authenticate_authenticate_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: authenticate.Empty
	(*AuthenticateRequest)(nil), // 1: authenticate.AuthenticateRequest
	(*AuthenticatePayload)(nil), // 2: authenticate.AuthenticatePayload
}
var file_gRPC_v1_authenticate_authenticate_proto_depIdxs = []int32{
	1, // 0: authenticate.AuthenticateService.Authenticate:input_type -> authenticate.AuthenticateRequest
	2, // 1: authenticate.AuthenticateService.Authenticate:output_type -> authenticate.AuthenticatePayload
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPC_v1_authenticate_authenticate_proto_init() }
func file_gRPC_v1_authenticate_authenticate_proto_init() {
	if File_gRPC_v1_authenticate_authenticate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPC_v1_authenticate_authenticate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_gRPC_v1_authenticate_authenticate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_gRPC_v1_authenticate_authenticate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatePayload); i {
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
			RawDescriptor: file_gRPC_v1_authenticate_authenticate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_v1_authenticate_authenticate_proto_goTypes,
		DependencyIndexes: file_gRPC_v1_authenticate_authenticate_proto_depIdxs,
		MessageInfos:      file_gRPC_v1_authenticate_authenticate_proto_msgTypes,
	}.Build()
	File_gRPC_v1_authenticate_authenticate_proto = out.File
	file_gRPC_v1_authenticate_authenticate_proto_rawDesc = nil
	file_gRPC_v1_authenticate_authenticate_proto_goTypes = nil
	file_gRPC_v1_authenticate_authenticate_proto_depIdxs = nil
}
