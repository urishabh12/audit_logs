// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/audit.proto

package audit_logs

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

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID  int64  `protobuf:"varint,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	UserID    int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	EntityID  int64  `protobuf:"varint,4,opt,name=EntityID,proto3" json:"EntityID,omitempty"`
	Entity    string `protobuf:"bytes,5,opt,name=Entity,proto3" json:"Entity,omitempty"`
	Action    string `protobuf:"bytes,6,opt,name=Action,proto3" json:"Action,omitempty"`
	Data      string `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetTenantID() int64 {
	if x != nil {
		return x.TenantID
	}
	return 0
}

func (x *LogRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LogRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogRequest) GetEntityID() int64 {
	if x != nil {
		return x.EntityID
	}
	return 0
}

func (x *LogRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *LogRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *LogRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type LogByEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID  int64  `protobuf:"varint,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	Entity    string `protobuf:"bytes,5,opt,name=Entity,proto3" json:"Entity,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *LogByEntityRequest) Reset() {
	*x = LogByEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogByEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogByEntityRequest) ProtoMessage() {}

func (x *LogByEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogByEntityRequest.ProtoReflect.Descriptor instead.
func (*LogByEntityRequest) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{1}
}

func (x *LogByEntityRequest) GetTenantID() int64 {
	if x != nil {
		return x.TenantID
	}
	return 0
}

func (x *LogByEntityRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *LogByEntityRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type LogByEntityPagedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID       int64  `protobuf:"varint,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	Entity         string `protobuf:"bytes,2,opt,name=Entity,proto3" json:"Entity,omitempty"`
	StartTimestamp int64  `protobuf:"varint,3,opt,name=StartTimestamp,proto3" json:"StartTimestamp,omitempty"`
	EndTimestamp   int64  `protobuf:"varint,4,opt,name=EndTimestamp,proto3" json:"EndTimestamp,omitempty"`
	PageSize       int32  `protobuf:"varint,5,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
}

func (x *LogByEntityPagedRequest) Reset() {
	*x = LogByEntityPagedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogByEntityPagedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogByEntityPagedRequest) ProtoMessage() {}

func (x *LogByEntityPagedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogByEntityPagedRequest.ProtoReflect.Descriptor instead.
func (*LogByEntityPagedRequest) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{2}
}

func (x *LogByEntityPagedRequest) GetTenantID() int64 {
	if x != nil {
		return x.TenantID
	}
	return 0
}

func (x *LogByEntityPagedRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *LogByEntityPagedRequest) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *LogByEntityPagedRequest) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

func (x *LogByEntityPagedRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type LogByEntityIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID  int64  `protobuf:"varint,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	Entity    string `protobuf:"bytes,2,opt,name=Entity,proto3" json:"Entity,omitempty"`
	EntityID  int64  `protobuf:"varint,3,opt,name=EntityID,proto3" json:"EntityID,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *LogByEntityIDRequest) Reset() {
	*x = LogByEntityIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogByEntityIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogByEntityIDRequest) ProtoMessage() {}

func (x *LogByEntityIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogByEntityIDRequest.ProtoReflect.Descriptor instead.
func (*LogByEntityIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{3}
}

func (x *LogByEntityIDRequest) GetTenantID() int64 {
	if x != nil {
		return x.TenantID
	}
	return 0
}

func (x *LogByEntityIDRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *LogByEntityIDRequest) GetEntityID() int64 {
	if x != nil {
		return x.EntityID
	}
	return 0
}

func (x *LogByEntityIDRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type LogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*AuditLog `protobuf:"bytes,1,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (x *LogsResponse) Reset() {
	*x = LogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsResponse) ProtoMessage() {}

func (x *LogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsResponse.ProtoReflect.Descriptor instead.
func (*LogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{4}
}

func (x *LogsResponse) GetLogs() []*AuditLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type AuditLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID  int64  `protobuf:"varint,1,opt,name=TenantID,proto3" json:"TenantID,omitempty"`
	UserID    int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	EntityID  int64  `protobuf:"varint,4,opt,name=EntityID,proto3" json:"EntityID,omitempty"`
	Entity    string `protobuf:"bytes,5,opt,name=Entity,proto3" json:"Entity,omitempty"`
	Action    string `protobuf:"bytes,6,opt,name=Action,proto3" json:"Action,omitempty"`
	Data      string `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{5}
}

func (x *AuditLog) GetTenantID() int64 {
	if x != nil {
		return x.TenantID
	}
	return 0
}

func (x *AuditLog) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *AuditLog) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AuditLog) GetEntityID() int64 {
	if x != nil {
		return x.EntityID
	}
	return 0
}

func (x *AuditLog) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *AuditLog) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AuditLog) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{6}
}

func (x *LogResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_proto_audit_proto protoreflect.FileDescriptor

var file_proto_audit_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x0a, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x12, 0x4c,
	0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0xb5, 0x01, 0x0a, 0x17, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x14,
	0x4c, 0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x33, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9c, 0x02,
	0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x19, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x72, 0x69, 0x73, 0x68,
	0x61, 0x62, 0x68, 0x31, 0x32, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_audit_proto_rawDescOnce sync.Once
	file_proto_audit_proto_rawDescData = file_proto_audit_proto_rawDesc
)

func file_proto_audit_proto_rawDescGZIP() []byte {
	file_proto_audit_proto_rawDescOnce.Do(func() {
		file_proto_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_audit_proto_rawDescData)
	})
	return file_proto_audit_proto_rawDescData
}

var file_proto_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_audit_proto_goTypes = []interface{}{
	(*LogRequest)(nil),              // 0: audit.LogRequest
	(*LogByEntityRequest)(nil),      // 1: audit.LogByEntityRequest
	(*LogByEntityPagedRequest)(nil), // 2: audit.LogByEntityPagedRequest
	(*LogByEntityIDRequest)(nil),    // 3: audit.LogByEntityIDRequest
	(*LogsResponse)(nil),            // 4: audit.LogsResponse
	(*AuditLog)(nil),                // 5: audit.AuditLog
	(*LogResponse)(nil),             // 6: audit.LogResponse
}
var file_proto_audit_proto_depIdxs = []int32{
	5, // 0: audit.LogsResponse.Logs:type_name -> audit.AuditLog
	0, // 1: audit.AuditService.Log:input_type -> audit.LogRequest
	1, // 2: audit.AuditService.GetLogByEntity:input_type -> audit.LogByEntityRequest
	2, // 3: audit.AuditService.GetLogByEntityPaginated:input_type -> audit.LogByEntityPagedRequest
	3, // 4: audit.AuditService.GetLogByEntityID:input_type -> audit.LogByEntityIDRequest
	6, // 5: audit.AuditService.Log:output_type -> audit.LogResponse
	4, // 6: audit.AuditService.GetLogByEntity:output_type -> audit.LogsResponse
	4, // 7: audit.AuditService.GetLogByEntityPaginated:output_type -> audit.LogsResponse
	4, // 8: audit.AuditService.GetLogByEntityID:output_type -> audit.LogsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_audit_proto_init() }
func file_proto_audit_proto_init() {
	if File_proto_audit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_audit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_proto_audit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogByEntityRequest); i {
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
		file_proto_audit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogByEntityPagedRequest); i {
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
		file_proto_audit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogByEntityIDRequest); i {
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
		file_proto_audit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsResponse); i {
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
		file_proto_audit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLog); i {
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
		file_proto_audit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
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
			RawDescriptor: file_proto_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_audit_proto_goTypes,
		DependencyIndexes: file_proto_audit_proto_depIdxs,
		MessageInfos:      file_proto_audit_proto_msgTypes,
	}.Build()
	File_proto_audit_proto = out.File
	file_proto_audit_proto_rawDesc = nil
	file_proto_audit_proto_goTypes = nil
	file_proto_audit_proto_depIdxs = nil
}
