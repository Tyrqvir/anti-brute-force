// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: antiBruteForceService.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccessCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip       uint32 `protobuf:"varint,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *AccessCheckRequest) Reset() {
	*x = AccessCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antiBruteForceService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessCheckRequest) ProtoMessage() {}

func (x *AccessCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antiBruteForceService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessCheckRequest.ProtoReflect.Descriptor instead.
func (*AccessCheckRequest) Descriptor() ([]byte, []int) {
	return file_antiBruteForceService_proto_rawDescGZIP(), []int{0}
}

func (x *AccessCheckRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AccessCheckRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AccessCheckRequest) GetIp() uint32 {
	if x != nil {
		return x.Ip
	}
	return 0
}

type AccessCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *AccessCheckResponse) Reset() {
	*x = AccessCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antiBruteForceService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessCheckResponse) ProtoMessage() {}

func (x *AccessCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_antiBruteForceService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessCheckResponse.ProtoReflect.Descriptor instead.
func (*AccessCheckResponse) Descriptor() ([]byte, []int) {
	return file_antiBruteForceService_proto_rawDescGZIP(), []int{1}
}

func (x *AccessCheckResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ResetBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Ip    uint32 `protobuf:"varint,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ResetBucketRequest) Reset() {
	*x = ResetBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antiBruteForceService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetBucketRequest) ProtoMessage() {}

func (x *ResetBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antiBruteForceService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetBucketRequest.ProtoReflect.Descriptor instead.
func (*ResetBucketRequest) Descriptor() ([]byte, []int) {
	return file_antiBruteForceService_proto_rawDescGZIP(), []int{2}
}

func (x *ResetBucketRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *ResetBucketRequest) GetIp() uint32 {
	if x != nil {
		return x.Ip
	}
	return 0
}

type AddToBlackListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *AddToBlackListRequest) Reset() {
	*x = AddToBlackListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antiBruteForceService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToBlackListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToBlackListRequest) ProtoMessage() {}

func (x *AddToBlackListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antiBruteForceService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToBlackListRequest.ProtoReflect.Descriptor instead.
func (*AddToBlackListRequest) Descriptor() ([]byte, []int) {
	return file_antiBruteForceService_proto_rawDescGZIP(), []int{3}
}

func (x *AddToBlackListRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type ListCases struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubNetwork string `protobuf:"bytes,1,opt,name=subNetwork,proto3" json:"subNetwork,omitempty"`
}

func (x *ListCases) Reset() {
	*x = ListCases{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antiBruteForceService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCases) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCases) ProtoMessage() {}

func (x *ListCases) ProtoReflect() protoreflect.Message {
	mi := &file_antiBruteForceService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCases.ProtoReflect.Descriptor instead.
func (*ListCases) Descriptor() ([]byte, []int) {
	return file_antiBruteForceService_proto_rawDescGZIP(), []int{4}
}

func (x *ListCases) GetSubNetwork() string {
	if x != nil {
		return x.SubNetwork
	}
	return ""
}

type ExistInListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok int64 `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ExistInListResponse) Reset() {
	*x = ExistInListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antiBruteForceService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistInListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistInListResponse) ProtoMessage() {}

func (x *ExistInListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_antiBruteForceService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistInListResponse.ProtoReflect.Descriptor instead.
func (*ExistInListResponse) Descriptor() ([]byte, []int) {
	return file_antiBruteForceService_proto_rawDescGZIP(), []int{5}
}

func (x *ExistInListResponse) GetOk() int64 {
	if x != nil {
		return x.Ok
	}
	return 0
}

var File_antiBruteForceService_proto protoreflect.FileDescriptor

var file_antiBruteForceService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x6e, 0x74, 0x69, 0x42, 0x72, 0x75, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x12, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x70, 0x22, 0x25, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x3a, 0x0a, 0x12, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x70, 0x22, 0x2d, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x22, 0x2b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x22, 0x25, 0x0a, 0x13, 0x45, 0x78, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xdb, 0x03, 0x0a, 0x15, 0x41, 0x6e, 0x74,
	0x69, 0x42, 0x72, 0x75, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x13, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39,
	0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x39, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x10, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x14, 0x2e, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x10, 0x45, 0x78, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73,
	0x1a, 0x14, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x6e, 0x74, 0x69, 0x2d, 0x62, 0x72, 0x75, 0x74, 0x65, 0x2d, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_antiBruteForceService_proto_rawDescOnce sync.Once
	file_antiBruteForceService_proto_rawDescData = file_antiBruteForceService_proto_rawDesc
)

func file_antiBruteForceService_proto_rawDescGZIP() []byte {
	file_antiBruteForceService_proto_rawDescOnce.Do(func() {
		file_antiBruteForceService_proto_rawDescData = protoimpl.X.CompressGZIP(file_antiBruteForceService_proto_rawDescData)
	})
	return file_antiBruteForceService_proto_rawDescData
}

var file_antiBruteForceService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_antiBruteForceService_proto_goTypes = []interface{}{
	(*AccessCheckRequest)(nil),    // 0: AccessCheckRequest
	(*AccessCheckResponse)(nil),   // 1: AccessCheckResponse
	(*ResetBucketRequest)(nil),    // 2: ResetBucketRequest
	(*AddToBlackListRequest)(nil), // 3: AddToBlackListRequest
	(*ListCases)(nil),             // 4: ListCases
	(*ExistInListResponse)(nil),   // 5: ExistInListResponse
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_antiBruteForceService_proto_depIdxs = []int32{
	0, // 0: AntiBruteForceService.AccessCheck:input_type -> AccessCheckRequest
	2, // 1: AntiBruteForceService.ResetBucket:input_type -> ResetBucketRequest
	4, // 2: AntiBruteForceService.AddToBlackList:input_type -> ListCases
	4, // 3: AntiBruteForceService.RemoveFromBlackList:input_type -> ListCases
	4, // 4: AntiBruteForceService.AddToWhiteList:input_type -> ListCases
	4, // 5: AntiBruteForceService.RemoveFromWhiteList:input_type -> ListCases
	4, // 6: AntiBruteForceService.ExistInWhiteList:input_type -> ListCases
	4, // 7: AntiBruteForceService.ExistInBlackList:input_type -> ListCases
	1, // 8: AntiBruteForceService.AccessCheck:output_type -> AccessCheckResponse
	6, // 9: AntiBruteForceService.ResetBucket:output_type -> google.protobuf.Empty
	6, // 10: AntiBruteForceService.AddToBlackList:output_type -> google.protobuf.Empty
	6, // 11: AntiBruteForceService.RemoveFromBlackList:output_type -> google.protobuf.Empty
	6, // 12: AntiBruteForceService.AddToWhiteList:output_type -> google.protobuf.Empty
	6, // 13: AntiBruteForceService.RemoveFromWhiteList:output_type -> google.protobuf.Empty
	5, // 14: AntiBruteForceService.ExistInWhiteList:output_type -> ExistInListResponse
	5, // 15: AntiBruteForceService.ExistInBlackList:output_type -> ExistInListResponse
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_antiBruteForceService_proto_init() }
func file_antiBruteForceService_proto_init() {
	if File_antiBruteForceService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_antiBruteForceService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessCheckRequest); i {
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
		file_antiBruteForceService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessCheckResponse); i {
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
		file_antiBruteForceService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetBucketRequest); i {
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
		file_antiBruteForceService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToBlackListRequest); i {
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
		file_antiBruteForceService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCases); i {
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
		file_antiBruteForceService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistInListResponse); i {
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
			RawDescriptor: file_antiBruteForceService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_antiBruteForceService_proto_goTypes,
		DependencyIndexes: file_antiBruteForceService_proto_depIdxs,
		MessageInfos:      file_antiBruteForceService_proto_msgTypes,
	}.Build()
	File_antiBruteForceService_proto = out.File
	file_antiBruteForceService_proto_rawDesc = nil
	file_antiBruteForceService_proto_goTypes = nil
	file_antiBruteForceService_proto_depIdxs = nil
}
