// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: guilds_service/guilds/v1/guilds.proto

package guildsv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Guild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon        string  `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Splash      string  `protobuf:"bytes,4,opt,name=splash,proto3" json:"splash,omitempty"`
	Banner      string  `protobuf:"bytes,5,opt,name=banner,proto3" json:"banner,omitempty"`
	Description string  `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	OwnerId     int64   `protobuf:"varint,7,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Roles       []int64 `protobuf:"varint,8,rep,packed,name=roles,proto3" json:"roles,omitempty"`
	CreatedAt   int64   `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   int64   `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Guild) Reset() {
	*x = Guild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guild) ProtoMessage() {}

func (x *Guild) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guild.ProtoReflect.Descriptor instead.
func (*Guild) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{0}
}

func (x *Guild) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Guild) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Guild) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Guild) GetSplash() string {
	if x != nil {
		return x.Splash
	}
	return ""
}

func (x *Guild) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *Guild) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Guild) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Guild) GetRoles() []int64 {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Guild) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Guild) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Icon    string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"` // TODO limit size
	OwnerId int64  `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CreateRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guild *Guild `protobuf:"bytes,1,opt,name=guild,proto3" json:"guild,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetGuild() *Guild {
	if x != nil {
		return x.Guild
	}
	return nil
}

type GetByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guild *Guild `protobuf:"bytes,1,opt,name=guild,proto3" json:"guild,omitempty"`
}

func (x *GetByIdResponse) Reset() {
	*x = GetByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdResponse) ProtoMessage() {}

func (x *GetByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdResponse.ProtoReflect.Descriptor instead.
func (*GetByIdResponse) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIdResponse) GetGuild() *Guild {
	if x != nil {
		return x.Guild
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Icon        string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`     // TODO limit size
	Splash      string `protobuf:"bytes,5,opt,name=splash,proto3" json:"splash,omitempty"` // TODO limit size
	Banner      string `protobuf:"bytes,6,opt,name=banner,proto3" json:"banner,omitempty"` // TODO limit size
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	OwnerId     int64  `protobuf:"varint,8,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UpdateRequest) GetSplash() string {
	if x != nil {
		return x.Splash
	}
	return ""
}

func (x *UpdateRequest) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *UpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guild *Guild `protobuf:"bytes,1,opt,name=guild,proto3" json:"guild,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetGuild() *Guild {
	if x != nil {
		return x.Guild
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guilds_service_guilds_v1_guilds_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_guilds_service_guilds_v1_guilds_proto protoreflect.FileDescriptor

var file_guilds_service_guilds_v1_guilds_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x80, 0x02, 0x0a, 0x05, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xa4, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x64, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3c, 0xba, 0x48, 0x39, 0xd0, 0x01, 0x01, 0x72, 0x34, 0x32, 0x32, 0x5e, 0x64, 0x61,
	0x74, 0x61, 0x3a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x28, 0x6a, 0x70, 0x65, 0x67, 0x7c, 0x70,
	0x6e, 0x67, 0x7c, 0x67, 0x69, 0x66, 0x29, 0x3b, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x2c, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2b, 0x2f, 0x3d, 0x5d, 0x2b, 0x24, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x67,
	0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x05, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x69,
	0x6c, 0x64, 0x52, 0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x95, 0x03, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x64, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3c, 0xba, 0x48, 0x39, 0xd0, 0x01, 0x01, 0x72, 0x34, 0x32, 0x32, 0x5e, 0x64, 0x61,
	0x74, 0x61, 0x3a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x28, 0x6a, 0x70, 0x65, 0x67, 0x7c, 0x70,
	0x6e, 0x67, 0x7c, 0x67, 0x69, 0x66, 0x29, 0x3b, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x2c, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2b, 0x2f, 0x3d, 0x5d, 0x2b, 0x24, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xba, 0x48, 0x35, 0xd0, 0x01, 0x01, 0x72, 0x30, 0x32,
	0x2e, 0x5e, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x28, 0x6a, 0x70,
	0x65, 0x67, 0x7c, 0x70, 0x6e, 0x67, 0x29, 0x3b, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x2c, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2b, 0x2f, 0x3d, 0x5d, 0x2b, 0x24, 0x52,
	0x06, 0x73, 0x70, 0x6c, 0x61, 0x73, 0x68, 0x12, 0x50, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xba, 0x48, 0x35, 0xd0, 0x01, 0x01, 0x72,
	0x30, 0x32, 0x2e, 0x5e, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x28,
	0x6a, 0x70, 0x65, 0x67, 0x7c, 0x70, 0x6e, 0x67, 0x29, 0x3b, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34,
	0x2c, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2b, 0x2f, 0x3d, 0x5d, 0x2b,
	0x24, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d,
	0xba, 0x48, 0x0a, 0xd0, 0x01, 0x01, 0x72, 0x05, 0x10, 0x02, 0x18, 0xe8, 0x07, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xba, 0x48,
	0x07, 0xd0, 0x01, 0x01, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x38, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x28, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x96, 0x02, 0x0a, 0x0d, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x19, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa2, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x33, 0x6e, 0x65, 0x2f, 0x64, 0x73, 0x5f,
	0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47,
	0x58, 0x58, 0xaa, 0x02, 0x09, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x09, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x47, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_guilds_service_guilds_v1_guilds_proto_rawDescOnce sync.Once
	file_guilds_service_guilds_v1_guilds_proto_rawDescData = file_guilds_service_guilds_v1_guilds_proto_rawDesc
)

func file_guilds_service_guilds_v1_guilds_proto_rawDescGZIP() []byte {
	file_guilds_service_guilds_v1_guilds_proto_rawDescOnce.Do(func() {
		file_guilds_service_guilds_v1_guilds_proto_rawDescData = protoimpl.X.CompressGZIP(file_guilds_service_guilds_v1_guilds_proto_rawDescData)
	})
	return file_guilds_service_guilds_v1_guilds_proto_rawDescData
}

var file_guilds_service_guilds_v1_guilds_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_guilds_service_guilds_v1_guilds_proto_goTypes = []interface{}{
	(*Guild)(nil),           // 0: guilds.v1.Guild
	(*CreateRequest)(nil),   // 1: guilds.v1.CreateRequest
	(*CreateResponse)(nil),  // 2: guilds.v1.CreateResponse
	(*GetByIdRequest)(nil),  // 3: guilds.v1.GetByIdRequest
	(*GetByIdResponse)(nil), // 4: guilds.v1.GetByIdResponse
	(*UpdateRequest)(nil),   // 5: guilds.v1.UpdateRequest
	(*UpdateResponse)(nil),  // 6: guilds.v1.UpdateResponse
	(*DeleteRequest)(nil),   // 7: guilds.v1.DeleteRequest
	(*DeleteResponse)(nil),  // 8: guilds.v1.DeleteResponse
}
var file_guilds_service_guilds_v1_guilds_proto_depIdxs = []int32{
	0, // 0: guilds.v1.CreateResponse.guild:type_name -> guilds.v1.Guild
	0, // 1: guilds.v1.GetByIdResponse.guild:type_name -> guilds.v1.Guild
	0, // 2: guilds.v1.UpdateResponse.guild:type_name -> guilds.v1.Guild
	1, // 3: guilds.v1.GuildsService.Create:input_type -> guilds.v1.CreateRequest
	3, // 4: guilds.v1.GuildsService.GetById:input_type -> guilds.v1.GetByIdRequest
	5, // 5: guilds.v1.GuildsService.Update:input_type -> guilds.v1.UpdateRequest
	7, // 6: guilds.v1.GuildsService.Delete:input_type -> guilds.v1.DeleteRequest
	2, // 7: guilds.v1.GuildsService.Create:output_type -> guilds.v1.CreateResponse
	4, // 8: guilds.v1.GuildsService.GetById:output_type -> guilds.v1.GetByIdResponse
	6, // 9: guilds.v1.GuildsService.Update:output_type -> guilds.v1.UpdateResponse
	8, // 10: guilds.v1.GuildsService.Delete:output_type -> guilds.v1.DeleteResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_guilds_service_guilds_v1_guilds_proto_init() }
func file_guilds_service_guilds_v1_guilds_proto_init() {
	if File_guilds_service_guilds_v1_guilds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guild); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdRequest); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdResponse); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_guilds_service_guilds_v1_guilds_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_guilds_service_guilds_v1_guilds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_guilds_service_guilds_v1_guilds_proto_goTypes,
		DependencyIndexes: file_guilds_service_guilds_v1_guilds_proto_depIdxs,
		MessageInfos:      file_guilds_service_guilds_v1_guilds_proto_msgTypes,
	}.Build()
	File_guilds_service_guilds_v1_guilds_proto = out.File
	file_guilds_service_guilds_v1_guilds_proto_rawDesc = nil
	file_guilds_service_guilds_v1_guilds_proto_goTypes = nil
	file_guilds_service_guilds_v1_guilds_proto_depIdxs = nil
}