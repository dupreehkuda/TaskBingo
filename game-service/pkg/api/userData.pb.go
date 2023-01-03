// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.9
// source: userData.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetUserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *GetUserDataRequest) Reset() {
	*x = GetUserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataRequest) ProtoMessage() {}

func (x *GetUserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataRequest.ProtoReflect.Descriptor instead.
func (*GetUserDataRequest) Descriptor() ([]byte, []int) {
	return file_userData_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserDataRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type GetUserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login      string        `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	City       string        `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Wins       int32         `protobuf:"varint,3,opt,name=wins,proto3" json:"wins,omitempty"`
	Loses      int32         `protobuf:"varint,4,opt,name=loses,proto3" json:"loses,omitempty"`
	Scoreboard int32         `protobuf:"varint,5,opt,name=scoreboard,proto3" json:"scoreboard,omitempty"`
	Friends    []*FriendInfo `protobuf:"bytes,6,rep,name=friends,proto3" json:"friends,omitempty"`
	Packs      []string      `protobuf:"bytes,7,rep,name=packs,proto3" json:"packs,omitempty"`
}

func (x *GetUserDataResponse) Reset() {
	*x = GetUserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataResponse) ProtoMessage() {}

func (x *GetUserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataResponse.ProtoReflect.Descriptor instead.
func (*GetUserDataResponse) Descriptor() ([]byte, []int) {
	return file_userData_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserDataResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetUserDataResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetUserDataResponse) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *GetUserDataResponse) GetLoses() int32 {
	if x != nil {
		return x.Loses
	}
	return 0
}

func (x *GetUserDataResponse) GetScoreboard() int32 {
	if x != nil {
		return x.Scoreboard
	}
	return 0
}

func (x *GetUserDataResponse) GetFriends() []*FriendInfo {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *GetUserDataResponse) GetPacks() []string {
	if x != nil {
		return x.Packs
	}
	return nil
}

type FriendInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	City  string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Score string `protobuf:"bytes,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *FriendInfo) Reset() {
	*x = FriendInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendInfo) ProtoMessage() {}

func (x *FriendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_userData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendInfo.ProtoReflect.Descriptor instead.
func (*FriendInfo) Descriptor() ([]byte, []int) {
	return file_userData_proto_rawDescGZIP(), []int{2}
}

func (x *FriendInfo) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *FriendInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *FriendInfo) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

type RegisterUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_userData_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterUserRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *RegisterUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_userData_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterUserResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LoginUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userData_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userData_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_userData_proto_rawDescGZIP(), []int{5}
}

func (x *LoginUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userData_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userData_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_userData_proto_rawDescGZIP(), []int{6}
}

func (x *LoginUserResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_userData_proto protoreflect.FileDescriptor

var file_userData_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x22, 0xca, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x73, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x29, 0x0a,
	0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x63, 0x6b,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x4c,
	0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x71, 0x0a, 0x13,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x2c, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x44, 0x0a,
	0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd0,
	0x01, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userData_proto_rawDescOnce sync.Once
	file_userData_proto_rawDescData = file_userData_proto_rawDesc
)

func file_userData_proto_rawDescGZIP() []byte {
	file_userData_proto_rawDescOnce.Do(func() {
		file_userData_proto_rawDescData = protoimpl.X.CompressGZIP(file_userData_proto_rawDescData)
	})
	return file_userData_proto_rawDescData
}

var file_userData_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_userData_proto_goTypes = []interface{}{
	(*GetUserDataRequest)(nil),   // 0: api.GetUserDataRequest
	(*GetUserDataResponse)(nil),  // 1: api.GetUserDataResponse
	(*FriendInfo)(nil),           // 2: api.FriendInfo
	(*RegisterUserRequest)(nil),  // 3: api.RegisterUserRequest
	(*RegisterUserResponse)(nil), // 4: api.RegisterUserResponse
	(*LoginUserRequest)(nil),     // 5: api.LoginUserRequest
	(*LoginUserResponse)(nil),    // 6: api.LoginUserResponse
}
var file_userData_proto_depIdxs = []int32{
	2, // 0: api.GetUserDataResponse.friends:type_name -> api.FriendInfo
	0, // 1: api.Users.GetUserData:input_type -> api.GetUserDataRequest
	3, // 2: api.Users.RegisterUser:input_type -> api.RegisterUserRequest
	5, // 3: api.Users.LoginUser:input_type -> api.LoginUserRequest
	1, // 4: api.Users.GetUserData:output_type -> api.GetUserDataResponse
	4, // 5: api.Users.RegisterUser:output_type -> api.RegisterUserResponse
	6, // 6: api.Users.LoginUser:output_type -> api.LoginUserResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_userData_proto_init() }
func file_userData_proto_init() {
	if File_userData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataRequest); i {
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
		file_userData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataResponse); i {
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
		file_userData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendInfo); i {
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
		file_userData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserRequest); i {
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
		file_userData_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserResponse); i {
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
		file_userData_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserRequest); i {
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
		file_userData_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserResponse); i {
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
			RawDescriptor: file_userData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userData_proto_goTypes,
		DependencyIndexes: file_userData_proto_depIdxs,
		MessageInfos:      file_userData_proto_msgTypes,
	}.Build()
	File_userData_proto = out.File
	file_userData_proto_rawDesc = nil
	file_userData_proto_goTypes = nil
	file_userData_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error) {
	out := new(GetUserDataResponse)
	err := c.cc.Invoke(ctx, "/api.Users/GetUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, "/api.Users/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/api.Users/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error)
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (*UnimplementedUsersServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (*UnimplementedUsersServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Users/GetUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserData(ctx, req.(*GetUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Users/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Users/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserData",
			Handler:    _Users_GetUserData_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _Users_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Users_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userData.proto",
}
