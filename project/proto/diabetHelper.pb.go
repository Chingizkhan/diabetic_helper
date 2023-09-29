// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: diabetHelper.proto

package diabetHelper

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

type AddSLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddSLRequest) Reset() {
	*x = AddSLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diabetHelper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSLRequest) ProtoMessage() {}

func (x *AddSLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diabetHelper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSLRequest.ProtoReflect.Descriptor instead.
func (*AddSLRequest) Descriptor() ([]byte, []int) {
	return file_diabetHelper_proto_rawDescGZIP(), []int{0}
}

func (x *AddSLRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddSLRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateSLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateSLRequest) Reset() {
	*x = UpdateSLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diabetHelper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSLRequest) ProtoMessage() {}

func (x *UpdateSLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diabetHelper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSLRequest.ProtoReflect.Descriptor instead.
func (*UpdateSLRequest) Descriptor() ([]byte, []int) {
	return file_diabetHelper_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateSLRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateSLRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type FindSLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination      *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	UserId          string      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Value           string      `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	CreatedAtStart  int64       `protobuf:"varint,5,opt,name=created_at_start,json=createdAtStart,proto3" json:"created_at_start,omitempty"`
	CreatedAtFinish int64       `protobuf:"varint,6,opt,name=created_at_finish,json=createdAtFinish,proto3" json:"created_at_finish,omitempty"`
}

func (x *FindSLRequest) Reset() {
	*x = FindSLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diabetHelper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSLRequest) ProtoMessage() {}

func (x *FindSLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diabetHelper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSLRequest.ProtoReflect.Descriptor instead.
func (*FindSLRequest) Descriptor() ([]byte, []int) {
	return file_diabetHelper_proto_rawDescGZIP(), []int{2}
}

func (x *FindSLRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *FindSLRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FindSLRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *FindSLRequest) GetCreatedAtStart() int64 {
	if x != nil {
		return x.CreatedAtStart
	}
	return 0
}

func (x *FindSLRequest) GetCreatedAtFinish() int64 {
	if x != nil {
		return x.CreatedAtFinish
	}
	return 0
}

type SugarLevels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PagesCount uint32        `protobuf:"varint,1,opt,name=pages_count,json=pagesCount,proto3" json:"pages_count,omitempty"`
	TotalItems uint32        `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	PerPage    uint32        `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Data       []*SugarLevel `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SugarLevels) Reset() {
	*x = SugarLevels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diabetHelper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SugarLevels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SugarLevels) ProtoMessage() {}

func (x *SugarLevels) ProtoReflect() protoreflect.Message {
	mi := &file_diabetHelper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SugarLevels.ProtoReflect.Descriptor instead.
func (*SugarLevels) Descriptor() ([]byte, []int) {
	return file_diabetHelper_proto_rawDescGZIP(), []int{3}
}

func (x *SugarLevels) GetPagesCount() uint32 {
	if x != nil {
		return x.PagesCount
	}
	return 0
}

func (x *SugarLevels) GetTotalItems() uint32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *SugarLevels) GetPerPage() uint32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *SugarLevels) GetData() []*SugarLevel {
	if x != nil {
		return x.Data
	}
	return nil
}

type SugarLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Value     string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SugarLevel) Reset() {
	*x = SugarLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diabetHelper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SugarLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SugarLevel) ProtoMessage() {}

func (x *SugarLevel) ProtoReflect() protoreflect.Message {
	mi := &file_diabetHelper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SugarLevel.ProtoReflect.Descriptor instead.
func (*SugarLevel) Descriptor() ([]byte, []int) {
	return file_diabetHelper_proto_rawDescGZIP(), []int{4}
}

func (x *SugarLevel) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SugarLevel) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SugarLevel) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SugarLevel) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Ok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Ok) Reset() {
	*x = Ok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diabetHelper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_diabetHelper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ok.ProtoReflect.Descriptor instead.
func (*Ok) Descriptor() ([]byte, []int) {
	return file_diabetHelper_proto_rawDescGZIP(), []int{5}
}

func (x *Ok) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diabetHelper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_diabetHelper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_diabetHelper_proto_rawDescGZIP(), []int{6}
}

func (x *Pagination) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_diabetHelper_proto protoreflect.FileDescriptor

var file_diabetHelper_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70,
	0x65, 0x72, 0x22, 0x3d, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x53, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x40, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x61, 0x62,
	0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x67, 0x61, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x53,
	0x75, 0x67, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x79, 0x0a, 0x0a, 0x53, 0x75, 0x67, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1a, 0x0a, 0x02, 0x4f, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4e, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0xd4, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x61, 0x62, 0x65,
	0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x53, 0x4c,
	0x12, 0x1a, 0x2e, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x53, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64,
	0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x67, 0x61,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x4c, 0x12, 0x1d, 0x2e, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x67, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x40, 0x0a, 0x06, 0x46,
	0x69, 0x6e, 0x64, 0x53, 0x4c, 0x12, 0x1b, 0x2e, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x2e, 0x53, 0x75, 0x67, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x42, 0x10, 0x5a,
	0x0e, 0x2f, 0x3b, 0x64, 0x69, 0x61, 0x62, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_diabetHelper_proto_rawDescOnce sync.Once
	file_diabetHelper_proto_rawDescData = file_diabetHelper_proto_rawDesc
)

func file_diabetHelper_proto_rawDescGZIP() []byte {
	file_diabetHelper_proto_rawDescOnce.Do(func() {
		file_diabetHelper_proto_rawDescData = protoimpl.X.CompressGZIP(file_diabetHelper_proto_rawDescData)
	})
	return file_diabetHelper_proto_rawDescData
}

var file_diabetHelper_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_diabetHelper_proto_goTypes = []interface{}{
	(*AddSLRequest)(nil),    // 0: diabetHelper.AddSLRequest
	(*UpdateSLRequest)(nil), // 1: diabetHelper.UpdateSLRequest
	(*FindSLRequest)(nil),   // 2: diabetHelper.FindSLRequest
	(*SugarLevels)(nil),     // 3: diabetHelper.SugarLevels
	(*SugarLevel)(nil),      // 4: diabetHelper.SugarLevel
	(*Ok)(nil),              // 5: diabetHelper.Ok
	(*Pagination)(nil),      // 6: diabetHelper.Pagination
}
var file_diabetHelper_proto_depIdxs = []int32{
	6, // 0: diabetHelper.FindSLRequest.pagination:type_name -> diabetHelper.Pagination
	4, // 1: diabetHelper.SugarLevels.data:type_name -> diabetHelper.SugarLevel
	0, // 2: diabetHelper.DiabetHelper.AddSL:input_type -> diabetHelper.AddSLRequest
	1, // 3: diabetHelper.DiabetHelper.UpdateSL:input_type -> diabetHelper.UpdateSLRequest
	2, // 4: diabetHelper.DiabetHelper.FindSL:input_type -> diabetHelper.FindSLRequest
	4, // 5: diabetHelper.DiabetHelper.AddSL:output_type -> diabetHelper.SugarLevel
	4, // 6: diabetHelper.DiabetHelper.UpdateSL:output_type -> diabetHelper.SugarLevel
	3, // 7: diabetHelper.DiabetHelper.FindSL:output_type -> diabetHelper.SugarLevels
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_diabetHelper_proto_init() }
func file_diabetHelper_proto_init() {
	if File_diabetHelper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_diabetHelper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSLRequest); i {
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
		file_diabetHelper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSLRequest); i {
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
		file_diabetHelper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSLRequest); i {
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
		file_diabetHelper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SugarLevels); i {
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
		file_diabetHelper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SugarLevel); i {
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
		file_diabetHelper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ok); i {
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
		file_diabetHelper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			RawDescriptor: file_diabetHelper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_diabetHelper_proto_goTypes,
		DependencyIndexes: file_diabetHelper_proto_depIdxs,
		MessageInfos:      file_diabetHelper_proto_msgTypes,
	}.Build()
	File_diabetHelper_proto = out.File
	file_diabetHelper_proto_rawDesc = nil
	file_diabetHelper_proto_goTypes = nil
	file_diabetHelper_proto_depIdxs = nil
}
