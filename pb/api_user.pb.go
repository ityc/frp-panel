// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v4.25.1
// source: api_user.proto

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

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{0}
}

type GetUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *Status `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	UserInfo *User   `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3,oneof" json:"user_info,omitempty"`
}

func (x *GetUserInfoResponse) Reset() {
	*x = GetUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResponse) ProtoMessage() {}

func (x *GetUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResponse.ProtoReflect.Descriptor instead.
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetUserInfoResponse) GetUserInfo() *User {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UpdateUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *User `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3,oneof" json:"user_info,omitempty"`
}

func (x *UpdateUserInfoRequest) Reset() {
	*x = UpdateUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoRequest) ProtoMessage() {}

func (x *UpdateUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserInfoRequest) GetUserInfo() *User {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UpdateUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *UpdateUserInfoResponse) Reset() {
	*x = UpdateUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoResponse) ProtoMessage() {}

func (x *UpdateUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserInfoResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetPlatformInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlatformInfoRequest) Reset() {
	*x = GetPlatformInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlatformInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlatformInfoRequest) ProtoMessage() {}

func (x *GetPlatformInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlatformInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPlatformInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{4}
}

type GetPlatformInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                  *Status `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	TotalClientCount        int32   `protobuf:"varint,2,opt,name=total_client_count,json=totalClientCount,proto3" json:"total_client_count,omitempty"`
	TotalServerCount        int32   `protobuf:"varint,3,opt,name=total_server_count,json=totalServerCount,proto3" json:"total_server_count,omitempty"`
	UnconfiguredClientCount int32   `protobuf:"varint,4,opt,name=unconfigured_client_count,json=unconfiguredClientCount,proto3" json:"unconfigured_client_count,omitempty"`
	UnconfiguredServerCount int32   `protobuf:"varint,5,opt,name=unconfigured_server_count,json=unconfiguredServerCount,proto3" json:"unconfigured_server_count,omitempty"`
	ConfiguredClientCount   int32   `protobuf:"varint,6,opt,name=configured_client_count,json=configuredClientCount,proto3" json:"configured_client_count,omitempty"`
	ConfiguredServerCount   int32   `protobuf:"varint,7,opt,name=configured_server_count,json=configuredServerCount,proto3" json:"configured_server_count,omitempty"`
	GlobalSecret            string  `protobuf:"bytes,8,opt,name=global_secret,json=globalSecret,proto3" json:"global_secret,omitempty"`
	MasterRpcHost           string  `protobuf:"bytes,9,opt,name=master_rpc_host,json=masterRpcHost,proto3" json:"master_rpc_host,omitempty"`
}

func (x *GetPlatformInfoResponse) Reset() {
	*x = GetPlatformInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlatformInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlatformInfoResponse) ProtoMessage() {}

func (x *GetPlatformInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlatformInfoResponse.ProtoReflect.Descriptor instead.
func (*GetPlatformInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlatformInfoResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetPlatformInfoResponse) GetTotalClientCount() int32 {
	if x != nil {
		return x.TotalClientCount
	}
	return 0
}

func (x *GetPlatformInfoResponse) GetTotalServerCount() int32 {
	if x != nil {
		return x.TotalServerCount
	}
	return 0
}

func (x *GetPlatformInfoResponse) GetUnconfiguredClientCount() int32 {
	if x != nil {
		return x.UnconfiguredClientCount
	}
	return 0
}

func (x *GetPlatformInfoResponse) GetUnconfiguredServerCount() int32 {
	if x != nil {
		return x.UnconfiguredServerCount
	}
	return 0
}

func (x *GetPlatformInfoResponse) GetConfiguredClientCount() int32 {
	if x != nil {
		return x.ConfiguredClientCount
	}
	return 0
}

func (x *GetPlatformInfoResponse) GetConfiguredServerCount() int32 {
	if x != nil {
		return x.ConfiguredServerCount
	}
	return 0
}

func (x *GetPlatformInfoResponse) GetGlobalSecret() string {
	if x != nil {
		return x.GlobalSecret
	}
	return ""
}

func (x *GetPlatformInfoResponse) GetMasterRpcHost() string {
	if x != nil {
		return x.MasterRpcHost
	}
	return ""
}

var File_api_user_proto protoreflect.FileDescriptor

var file_api_user_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8b,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x48, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0x50, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xe2, 0x03, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x19, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x17, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x70, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_proto_rawDescOnce sync.Once
	file_api_user_proto_rawDescData = file_api_user_proto_rawDesc
)

func file_api_user_proto_rawDescGZIP() []byte {
	file_api_user_proto_rawDescOnce.Do(func() {
		file_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_proto_rawDescData)
	})
	return file_api_user_proto_rawDescData
}

var file_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_user_proto_goTypes = []interface{}{
	(*GetUserInfoRequest)(nil),      // 0: api_user.GetUserInfoRequest
	(*GetUserInfoResponse)(nil),     // 1: api_user.GetUserInfoResponse
	(*UpdateUserInfoRequest)(nil),   // 2: api_user.UpdateUserInfoRequest
	(*UpdateUserInfoResponse)(nil),  // 3: api_user.UpdateUserInfoResponse
	(*GetPlatformInfoRequest)(nil),  // 4: api_user.GetPlatformInfoRequest
	(*GetPlatformInfoResponse)(nil), // 5: api_user.GetPlatformInfoResponse
	(*Status)(nil),                  // 6: common.Status
	(*User)(nil),                    // 7: common.User
}
var file_api_user_proto_depIdxs = []int32{
	6, // 0: api_user.GetUserInfoResponse.status:type_name -> common.Status
	7, // 1: api_user.GetUserInfoResponse.user_info:type_name -> common.User
	7, // 2: api_user.UpdateUserInfoRequest.user_info:type_name -> common.User
	6, // 3: api_user.UpdateUserInfoResponse.status:type_name -> common.Status
	6, // 4: api_user.GetPlatformInfoResponse.status:type_name -> common.Status
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_user_proto_init() }
func file_api_user_proto_init() {
	if File_api_user_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRequest); i {
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
		file_api_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResponse); i {
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
		file_api_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoRequest); i {
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
		file_api_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoResponse); i {
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
		file_api_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlatformInfoRequest); i {
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
		file_api_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlatformInfoResponse); i {
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
	file_api_user_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_api_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_api_user_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_api_user_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_proto_goTypes,
		DependencyIndexes: file_api_user_proto_depIdxs,
		MessageInfos:      file_api_user_proto_msgTypes,
	}.Build()
	File_api_user_proto = out.File
	file_api_user_proto_rawDesc = nil
	file_api_user_proto_goTypes = nil
	file_api_user_proto_depIdxs = nil
}
