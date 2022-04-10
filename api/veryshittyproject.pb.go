// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: veryshittyproject.proto

package vspservice

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

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_veryshittyproject_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_veryshittyproject_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_veryshittyproject_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *UUID  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Username    string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	FirstName   string `protobuf:"bytes,6,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_veryshittyproject_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_veryshittyproject_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_veryshittyproject_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type LoginTelegramDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginTelegramDetailsRequest) Reset() {
	*x = LoginTelegramDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_veryshittyproject_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginTelegramDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginTelegramDetailsRequest) ProtoMessage() {}

func (x *LoginTelegramDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_veryshittyproject_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginTelegramDetailsRequest.ProtoReflect.Descriptor instead.
func (*LoginTelegramDetailsRequest) Descriptor() ([]byte, []int) {
	return file_veryshittyproject_proto_rawDescGZIP(), []int{2}
}

type LoginTelegramDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *LoginTelegramDetailsResponse) Reset() {
	*x = LoginTelegramDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_veryshittyproject_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginTelegramDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginTelegramDetailsResponse) ProtoMessage() {}

func (x *LoginTelegramDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_veryshittyproject_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginTelegramDetailsResponse.ProtoReflect.Descriptor instead.
func (*LoginTelegramDetailsResponse) Descriptor() ([]byte, []int) {
	return file_veryshittyproject_proto_rawDescGZIP(), []int{3}
}

func (x *LoginTelegramDetailsResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type LoginTelegramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	PhotoUrl  string `protobuf:"bytes,5,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	AuthDate  string `protobuf:"bytes,6,opt,name=auth_date,json=authDate,proto3" json:"auth_date,omitempty"`
	Hash      string `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *LoginTelegramRequest) Reset() {
	*x = LoginTelegramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_veryshittyproject_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginTelegramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginTelegramRequest) ProtoMessage() {}

func (x *LoginTelegramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_veryshittyproject_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginTelegramRequest.ProtoReflect.Descriptor instead.
func (*LoginTelegramRequest) Descriptor() ([]byte, []int) {
	return file_veryshittyproject_proto_rawDescGZIP(), []int{4}
}

func (x *LoginTelegramRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginTelegramRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *LoginTelegramRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *LoginTelegramRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginTelegramRequest) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *LoginTelegramRequest) GetAuthDate() string {
	if x != nil {
		return x.AuthDate
	}
	return ""
}

func (x *LoginTelegramRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type LoginTelegramResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *LoginTelegramResponse) Reset() {
	*x = LoginTelegramResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_veryshittyproject_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginTelegramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginTelegramResponse) ProtoMessage() {}

func (x *LoginTelegramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_veryshittyproject_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginTelegramResponse.ProtoReflect.Descriptor instead.
func (*LoginTelegramResponse) Descriptor() ([]byte, []int) {
	return file_veryshittyproject_proto_rawDescGZIP(), []int{5}
}

func (x *LoginTelegramResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_veryshittyproject_proto protoreflect.FileDescriptor

var file_veryshittyproject_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x65, 0x72, 0x79, 0x73, 0x68, 0x69, 0x74, 0x74, 0x79, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x1c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0xcc, 0x01, 0x0a,
	0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x32, 0x0a, 0x15, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32,
	0x9a, 0x01, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x53, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x1c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x15, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x79, 0x73,
	0x68, 0x69, 0x74, 0x74, 0x79, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x76, 0x73, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_veryshittyproject_proto_rawDescOnce sync.Once
	file_veryshittyproject_proto_rawDescData = file_veryshittyproject_proto_rawDesc
)

func file_veryshittyproject_proto_rawDescGZIP() []byte {
	file_veryshittyproject_proto_rawDescOnce.Do(func() {
		file_veryshittyproject_proto_rawDescData = protoimpl.X.CompressGZIP(file_veryshittyproject_proto_rawDescData)
	})
	return file_veryshittyproject_proto_rawDescData
}

var file_veryshittyproject_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_veryshittyproject_proto_goTypes = []interface{}{
	(*UUID)(nil),                         // 0: UUID
	(*User)(nil),                         // 1: User
	(*LoginTelegramDetailsRequest)(nil),  // 2: LoginTelegramDetailsRequest
	(*LoginTelegramDetailsResponse)(nil), // 3: LoginTelegramDetailsResponse
	(*LoginTelegramRequest)(nil),         // 4: LoginTelegramRequest
	(*LoginTelegramResponse)(nil),        // 5: LoginTelegramResponse
}
var file_veryshittyproject_proto_depIdxs = []int32{
	0, // 0: User.id:type_name -> UUID
	1, // 1: LoginTelegramResponse.user:type_name -> User
	2, // 2: API.LoginTelegramDetails:input_type -> LoginTelegramDetailsRequest
	4, // 3: API.LoginTelegram:input_type -> LoginTelegramRequest
	3, // 4: API.LoginTelegramDetails:output_type -> LoginTelegramDetailsResponse
	5, // 5: API.LoginTelegram:output_type -> LoginTelegramResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_veryshittyproject_proto_init() }
func file_veryshittyproject_proto_init() {
	if File_veryshittyproject_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_veryshittyproject_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_veryshittyproject_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_veryshittyproject_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginTelegramDetailsRequest); i {
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
		file_veryshittyproject_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginTelegramDetailsResponse); i {
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
		file_veryshittyproject_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginTelegramRequest); i {
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
		file_veryshittyproject_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginTelegramResponse); i {
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
			RawDescriptor: file_veryshittyproject_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_veryshittyproject_proto_goTypes,
		DependencyIndexes: file_veryshittyproject_proto_depIdxs,
		MessageInfos:      file_veryshittyproject_proto_msgTypes,
	}.Build()
	File_veryshittyproject_proto = out.File
	file_veryshittyproject_proto_rawDesc = nil
	file_veryshittyproject_proto_goTypes = nil
	file_veryshittyproject_proto_depIdxs = nil
}
