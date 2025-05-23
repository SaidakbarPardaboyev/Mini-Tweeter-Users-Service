// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: followings.proto

package users_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FollowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FollowerId    string                 `protobuf:"bytes,1,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
	FollowingId   string                 `protobuf:"bytes,2,opt,name=following_id,json=followingId,proto3" json:"following_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	mi := &file_followings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRequest) GetFollowerId() string {
	if x != nil {
		return x.FollowerId
	}
	return ""
}

func (x *FollowRequest) GetFollowingId() string {
	if x != nil {
		return x.FollowingId
	}
	return ""
}

type FollowResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	mi := &file_followings_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{1}
}

func (x *FollowResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UnFollowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FollowerId    string                 `protobuf:"bytes,1,opt,name=follower_id,json=followerId,proto3" json:"follower_id,omitempty"`
	FollowingId   string                 `protobuf:"bytes,2,opt,name=following_id,json=followingId,proto3" json:"following_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnFollowRequest) Reset() {
	*x = UnFollowRequest{}
	mi := &file_followings_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnFollowRequest) ProtoMessage() {}

func (x *UnFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnFollowRequest.ProtoReflect.Descriptor instead.
func (*UnFollowRequest) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{2}
}

func (x *UnFollowRequest) GetFollowerId() string {
	if x != nil {
		return x.FollowerId
	}
	return ""
}

func (x *UnFollowRequest) GetFollowingId() string {
	if x != nil {
		return x.FollowingId
	}
	return ""
}

type UnFollowResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnFollowResponse) Reset() {
	*x = UnFollowResponse{}
	mi := &file_followings_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnFollowResponse) ProtoMessage() {}

func (x *UnFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnFollowResponse.ProtoReflect.Descriptor instead.
func (*UnFollowResponse) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{3}
}

func (x *UnFollowResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request message for GetFollowing
type GetListFollowingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Search        *string                `protobuf:"bytes,4,opt,name=search,proto3,oneof" json:"search,omitempty"` // Merged field for name, surname, Followingname
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListFollowingRequest) Reset() {
	*x = GetListFollowingRequest{}
	mi := &file_followings_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListFollowingRequest) ProtoMessage() {}

func (x *GetListFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListFollowingRequest.ProtoReflect.Descriptor instead.
func (*GetListFollowingRequest) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{4}
}

func (x *GetListFollowingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetListFollowingRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListFollowingRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListFollowingRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

// Request message for GetFollowing
type GetListFollowerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Search        *string                `protobuf:"bytes,4,opt,name=search,proto3,oneof" json:"search,omitempty"` // Merged field for name, surname, Followingname
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListFollowerRequest) Reset() {
	*x = GetListFollowerRequest{}
	mi := &file_followings_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListFollowerRequest) ProtoMessage() {}

func (x *GetListFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListFollowerRequest.ProtoReflect.Descriptor instead.
func (*GetListFollowerRequest) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{5}
}

func (x *GetListFollowerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetListFollowerRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListFollowerRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListFollowerRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

// Response message for GetFollowing
type FollowingList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Followings    []*User                `protobuf:"bytes,1,rep,name=Followings,proto3" json:"Followings,omitempty"`
	Count         int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FollowingList) Reset() {
	*x = FollowingList{}
	mi := &file_followings_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingList) ProtoMessage() {}

func (x *FollowingList) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingList.ProtoReflect.Descriptor instead.
func (*FollowingList) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{6}
}

func (x *FollowingList) GetFollowings() []*User {
	if x != nil {
		return x.Followings
	}
	return nil
}

func (x *FollowingList) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Response message for GetFollowing
type FollowerList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Followers     []*User                `protobuf:"bytes,1,rep,name=Followers,proto3" json:"Followers,omitempty"`
	Count         int64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FollowerList) Reset() {
	*x = FollowerList{}
	mi := &file_followings_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerList) ProtoMessage() {}

func (x *FollowerList) ProtoReflect() protoreflect.Message {
	mi := &file_followings_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerList.ProtoReflect.Descriptor instead.
func (*FollowerList) Descriptor() ([]byte, []int) {
	return file_followings_proto_rawDescGZIP(), []int{7}
}

func (x *FollowerList) GetFollowers() []*User {
	if x != nil {
		return x.Followers
	}
	return nil
}

func (x *FollowerList) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_followings_proto protoreflect.FileDescriptor

var file_followings_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53,
	0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x55, 0x0a, 0x0f, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x10, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x7b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x22, 0x7a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5a, 0x0a,
	0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0c, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x09, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xd9, 0x02, 0x0a, 0x10, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x08, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x18,
	0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_followings_proto_rawDescOnce sync.Once
	file_followings_proto_rawDescData []byte
)

func file_followings_proto_rawDescGZIP() []byte {
	file_followings_proto_rawDescOnce.Do(func() {
		file_followings_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_followings_proto_rawDesc), len(file_followings_proto_rawDesc)))
	})
	return file_followings_proto_rawDescData
}

var file_followings_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_followings_proto_goTypes = []any{
	(*FollowRequest)(nil),           // 0: users_service.FollowRequest
	(*FollowResponse)(nil),          // 1: users_service.FollowResponse
	(*UnFollowRequest)(nil),         // 2: users_service.UnFollowRequest
	(*UnFollowResponse)(nil),        // 3: users_service.UnFollowResponse
	(*GetListFollowingRequest)(nil), // 4: users_service.GetListFollowingRequest
	(*GetListFollowerRequest)(nil),  // 5: users_service.GetListFollowerRequest
	(*FollowingList)(nil),           // 6: users_service.FollowingList
	(*FollowerList)(nil),            // 7: users_service.FollowerList
	(*User)(nil),                    // 8: users_service.User
}
var file_followings_proto_depIdxs = []int32{
	8, // 0: users_service.FollowingList.Followings:type_name -> users_service.User
	8, // 1: users_service.FollowerList.Followers:type_name -> users_service.User
	0, // 2: users_service.FollowingService.Follow:input_type -> users_service.FollowRequest
	2, // 3: users_service.FollowingService.UnFollow:input_type -> users_service.UnFollowRequest
	4, // 4: users_service.FollowingService.GetListFollowings:input_type -> users_service.GetListFollowingRequest
	5, // 5: users_service.FollowingService.GetListFollowers:input_type -> users_service.GetListFollowerRequest
	1, // 6: users_service.FollowingService.Follow:output_type -> users_service.FollowResponse
	3, // 7: users_service.FollowingService.UnFollow:output_type -> users_service.UnFollowResponse
	6, // 8: users_service.FollowingService.GetListFollowings:output_type -> users_service.FollowingList
	7, // 9: users_service.FollowingService.GetListFollowers:output_type -> users_service.FollowerList
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_followings_proto_init() }
func file_followings_proto_init() {
	if File_followings_proto != nil {
		return
	}
	file_users_proto_init()
	file_followings_proto_msgTypes[4].OneofWrappers = []any{}
	file_followings_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_followings_proto_rawDesc), len(file_followings_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_followings_proto_goTypes,
		DependencyIndexes: file_followings_proto_depIdxs,
		MessageInfos:      file_followings_proto_msgTypes,
	}.Build()
	File_followings_proto = out.File
	file_followings_proto_goTypes = nil
	file_followings_proto_depIdxs = nil
}
