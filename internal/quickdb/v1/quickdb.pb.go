// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.7
// source: proto/quickdb/v1/quickdb.proto

package v1

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

type CreateEphemeralDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region     string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Engine     string   `protobuf:"bytes,2,opt,name=engine,proto3" json:"engine,omitempty"`
	ServerId   string   `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Migrations []string `protobuf:"bytes,4,rep,name=migrations,proto3" json:"migrations,omitempty"`
}

func (x *CreateEphemeralDatabaseRequest) Reset() {
	*x = CreateEphemeralDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEphemeralDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEphemeralDatabaseRequest) ProtoMessage() {}

func (x *CreateEphemeralDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEphemeralDatabaseRequest.ProtoReflect.Descriptor instead.
func (*CreateEphemeralDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_quickdb_v1_quickdb_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEphemeralDatabaseRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateEphemeralDatabaseRequest) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *CreateEphemeralDatabaseRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CreateEphemeralDatabaseRequest) GetMigrations() []string {
	if x != nil {
		return x.Migrations
	}
	return nil
}

type CreateEphemeralDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseId string `protobuf:"bytes,1,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
	Uri        string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *CreateEphemeralDatabaseResponse) Reset() {
	*x = CreateEphemeralDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEphemeralDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEphemeralDatabaseResponse) ProtoMessage() {}

func (x *CreateEphemeralDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEphemeralDatabaseResponse.ProtoReflect.Descriptor instead.
func (*CreateEphemeralDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_quickdb_v1_quickdb_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEphemeralDatabaseResponse) GetDatabaseId() string {
	if x != nil {
		return x.DatabaseId
	}
	return ""
}

func (x *CreateEphemeralDatabaseResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type DropEphemeralDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseId string `protobuf:"bytes,1,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
}

func (x *DropEphemeralDatabaseRequest) Reset() {
	*x = DropEphemeralDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropEphemeralDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropEphemeralDatabaseRequest) ProtoMessage() {}

func (x *DropEphemeralDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropEphemeralDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DropEphemeralDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_quickdb_v1_quickdb_proto_rawDescGZIP(), []int{2}
}

func (x *DropEphemeralDatabaseRequest) GetDatabaseId() string {
	if x != nil {
		return x.DatabaseId
	}
	return ""
}

type DropEphemeralDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DropEphemeralDatabaseResponse) Reset() {
	*x = DropEphemeralDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropEphemeralDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropEphemeralDatabaseResponse) ProtoMessage() {}

func (x *DropEphemeralDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropEphemeralDatabaseResponse.ProtoReflect.Descriptor instead.
func (*DropEphemeralDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_quickdb_v1_quickdb_proto_rawDescGZIP(), []int{3}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Contents []byte `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_proto_quickdb_v1_quickdb_proto_rawDescGZIP(), []int{4}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

type UploadArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SqlcVersion string  `protobuf:"bytes,1,opt,name=sqlc_version,json=sqlcVersion,proto3" json:"sqlc_version,omitempty"`
	Inputs      []*File `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs     []*File `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *UploadArchiveRequest) Reset() {
	*x = UploadArchiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadArchiveRequest) ProtoMessage() {}

func (x *UploadArchiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadArchiveRequest.ProtoReflect.Descriptor instead.
func (*UploadArchiveRequest) Descriptor() ([]byte, []int) {
	return file_proto_quickdb_v1_quickdb_proto_rawDescGZIP(), []int{5}
}

func (x *UploadArchiveRequest) GetSqlcVersion() string {
	if x != nil {
		return x.SqlcVersion
	}
	return ""
}

func (x *UploadArchiveRequest) GetInputs() []*File {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *UploadArchiveRequest) GetOutputs() []*File {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type UploadArchiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checksum []byte `protobuf:"bytes,1,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *UploadArchiveResponse) Reset() {
	*x = UploadArchiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadArchiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadArchiveResponse) ProtoMessage() {}

func (x *UploadArchiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quickdb_v1_quickdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadArchiveResponse.ProtoReflect.Descriptor instead.
func (*UploadArchiveResponse) Descriptor() ([]byte, []int) {
	return file_proto_quickdb_v1_quickdb_proto_rawDescGZIP(), []int{6}
}

func (x *UploadArchiveResponse) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

var File_proto_quickdb_v1_quickdb_proto protoreflect.FileDescriptor

var file_proto_quickdb_v1_quickdb_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x71, 0x6c, 0x63, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x8d, 0x01, 0x0a,
	0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x54, 0x0a, 0x1f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x22, 0x3f, 0x0a, 0x1c, 0x44, 0x72, 0x6f, 0x70, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65,
	0x72, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x72, 0x6f, 0x70, 0x45, 0x70, 0x68, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xaf, 0x01, 0x0a,
	0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x71, 0x6c, 0x63, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x71, 0x6c,
	0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x73, 0x71, 0x6c, 0x63, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x71, 0x6c,
	0x63, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x33,
	0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x32, 0xa1, 0x03, 0x0a, 0x05, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x12, 0x92, 0x01,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3a, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x73, 0x71, 0x6c, 0x63, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71, 0x75, 0x69, 0x63,
	0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x70, 0x68,
	0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x73,
	0x71, 0x6c, 0x63, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x15, 0x44, 0x72, 0x6f, 0x70, 0x45, 0x70, 0x68, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x38, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x71, 0x6c, 0x63, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71,
	0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x45, 0x70,
	0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x73, 0x71, 0x6c, 0x63, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x74, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x12, 0x30, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x71, 0x6c, 0x63,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x71,
	0x6c, 0x63, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quickdb_v1_quickdb_proto_rawDescOnce sync.Once
	file_proto_quickdb_v1_quickdb_proto_rawDescData = file_proto_quickdb_v1_quickdb_proto_rawDesc
)

func file_proto_quickdb_v1_quickdb_proto_rawDescGZIP() []byte {
	file_proto_quickdb_v1_quickdb_proto_rawDescOnce.Do(func() {
		file_proto_quickdb_v1_quickdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quickdb_v1_quickdb_proto_rawDescData)
	})
	return file_proto_quickdb_v1_quickdb_proto_rawDescData
}

var file_proto_quickdb_v1_quickdb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_quickdb_v1_quickdb_proto_goTypes = []interface{}{
	(*CreateEphemeralDatabaseRequest)(nil),  // 0: remote.sqlc.dev.quickdb.v1.CreateEphemeralDatabaseRequest
	(*CreateEphemeralDatabaseResponse)(nil), // 1: remote.sqlc.dev.quickdb.v1.CreateEphemeralDatabaseResponse
	(*DropEphemeralDatabaseRequest)(nil),    // 2: remote.sqlc.dev.quickdb.v1.DropEphemeralDatabaseRequest
	(*DropEphemeralDatabaseResponse)(nil),   // 3: remote.sqlc.dev.quickdb.v1.DropEphemeralDatabaseResponse
	(*File)(nil),                            // 4: remote.sqlc.dev.quickdb.v1.File
	(*UploadArchiveRequest)(nil),            // 5: remote.sqlc.dev.quickdb.v1.UploadArchiveRequest
	(*UploadArchiveResponse)(nil),           // 6: remote.sqlc.dev.quickdb.v1.UploadArchiveResponse
}
var file_proto_quickdb_v1_quickdb_proto_depIdxs = []int32{
	4, // 0: remote.sqlc.dev.quickdb.v1.UploadArchiveRequest.inputs:type_name -> remote.sqlc.dev.quickdb.v1.File
	4, // 1: remote.sqlc.dev.quickdb.v1.UploadArchiveRequest.outputs:type_name -> remote.sqlc.dev.quickdb.v1.File
	0, // 2: remote.sqlc.dev.quickdb.v1.Quick.CreateEphemeralDatabase:input_type -> remote.sqlc.dev.quickdb.v1.CreateEphemeralDatabaseRequest
	2, // 3: remote.sqlc.dev.quickdb.v1.Quick.DropEphemeralDatabase:input_type -> remote.sqlc.dev.quickdb.v1.DropEphemeralDatabaseRequest
	5, // 4: remote.sqlc.dev.quickdb.v1.Quick.UploadArchive:input_type -> remote.sqlc.dev.quickdb.v1.UploadArchiveRequest
	1, // 5: remote.sqlc.dev.quickdb.v1.Quick.CreateEphemeralDatabase:output_type -> remote.sqlc.dev.quickdb.v1.CreateEphemeralDatabaseResponse
	3, // 6: remote.sqlc.dev.quickdb.v1.Quick.DropEphemeralDatabase:output_type -> remote.sqlc.dev.quickdb.v1.DropEphemeralDatabaseResponse
	6, // 7: remote.sqlc.dev.quickdb.v1.Quick.UploadArchive:output_type -> remote.sqlc.dev.quickdb.v1.UploadArchiveResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_quickdb_v1_quickdb_proto_init() }
func file_proto_quickdb_v1_quickdb_proto_init() {
	if File_proto_quickdb_v1_quickdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quickdb_v1_quickdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEphemeralDatabaseRequest); i {
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
		file_proto_quickdb_v1_quickdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEphemeralDatabaseResponse); i {
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
		file_proto_quickdb_v1_quickdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropEphemeralDatabaseRequest); i {
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
		file_proto_quickdb_v1_quickdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropEphemeralDatabaseResponse); i {
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
		file_proto_quickdb_v1_quickdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_proto_quickdb_v1_quickdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadArchiveRequest); i {
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
		file_proto_quickdb_v1_quickdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadArchiveResponse); i {
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
			RawDescriptor: file_proto_quickdb_v1_quickdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quickdb_v1_quickdb_proto_goTypes,
		DependencyIndexes: file_proto_quickdb_v1_quickdb_proto_depIdxs,
		MessageInfos:      file_proto_quickdb_v1_quickdb_proto_msgTypes,
	}.Build()
	File_proto_quickdb_v1_quickdb_proto = out.File
	file_proto_quickdb_v1_quickdb_proto_rawDesc = nil
	file_proto_quickdb_v1_quickdb_proto_goTypes = nil
	file_proto_quickdb_v1_quickdb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QuickClient is the client API for Quick service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuickClient interface {
	CreateEphemeralDatabase(ctx context.Context, in *CreateEphemeralDatabaseRequest, opts ...grpc.CallOption) (*CreateEphemeralDatabaseResponse, error)
	DropEphemeralDatabase(ctx context.Context, in *DropEphemeralDatabaseRequest, opts ...grpc.CallOption) (*DropEphemeralDatabaseResponse, error)
	UploadArchive(ctx context.Context, in *UploadArchiveRequest, opts ...grpc.CallOption) (*UploadArchiveResponse, error)
}

type quickClient struct {
	cc grpc.ClientConnInterface
}

func NewQuickClient(cc grpc.ClientConnInterface) QuickClient {
	return &quickClient{cc}
}

func (c *quickClient) CreateEphemeralDatabase(ctx context.Context, in *CreateEphemeralDatabaseRequest, opts ...grpc.CallOption) (*CreateEphemeralDatabaseResponse, error) {
	out := new(CreateEphemeralDatabaseResponse)
	err := c.cc.Invoke(ctx, "/remote.sqlc.dev.quickdb.v1.Quick/CreateEphemeralDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickClient) DropEphemeralDatabase(ctx context.Context, in *DropEphemeralDatabaseRequest, opts ...grpc.CallOption) (*DropEphemeralDatabaseResponse, error) {
	out := new(DropEphemeralDatabaseResponse)
	err := c.cc.Invoke(ctx, "/remote.sqlc.dev.quickdb.v1.Quick/DropEphemeralDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quickClient) UploadArchive(ctx context.Context, in *UploadArchiveRequest, opts ...grpc.CallOption) (*UploadArchiveResponse, error) {
	out := new(UploadArchiveResponse)
	err := c.cc.Invoke(ctx, "/remote.sqlc.dev.quickdb.v1.Quick/UploadArchive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuickServer is the server API for Quick service.
type QuickServer interface {
	CreateEphemeralDatabase(context.Context, *CreateEphemeralDatabaseRequest) (*CreateEphemeralDatabaseResponse, error)
	DropEphemeralDatabase(context.Context, *DropEphemeralDatabaseRequest) (*DropEphemeralDatabaseResponse, error)
	UploadArchive(context.Context, *UploadArchiveRequest) (*UploadArchiveResponse, error)
}

// UnimplementedQuickServer can be embedded to have forward compatible implementations.
type UnimplementedQuickServer struct {
}

func (*UnimplementedQuickServer) CreateEphemeralDatabase(context.Context, *CreateEphemeralDatabaseRequest) (*CreateEphemeralDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEphemeralDatabase not implemented")
}
func (*UnimplementedQuickServer) DropEphemeralDatabase(context.Context, *DropEphemeralDatabaseRequest) (*DropEphemeralDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropEphemeralDatabase not implemented")
}
func (*UnimplementedQuickServer) UploadArchive(context.Context, *UploadArchiveRequest) (*UploadArchiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadArchive not implemented")
}

func RegisterQuickServer(s *grpc.Server, srv QuickServer) {
	s.RegisterService(&_Quick_serviceDesc, srv)
}

func _Quick_CreateEphemeralDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEphemeralDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickServer).CreateEphemeralDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.sqlc.dev.quickdb.v1.Quick/CreateEphemeralDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickServer).CreateEphemeralDatabase(ctx, req.(*CreateEphemeralDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quick_DropEphemeralDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropEphemeralDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickServer).DropEphemeralDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.sqlc.dev.quickdb.v1.Quick/DropEphemeralDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickServer).DropEphemeralDatabase(ctx, req.(*DropEphemeralDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quick_UploadArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadArchiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickServer).UploadArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.sqlc.dev.quickdb.v1.Quick/UploadArchive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickServer).UploadArchive(ctx, req.(*UploadArchiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Quick_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.sqlc.dev.quickdb.v1.Quick",
	HandlerType: (*QuickServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEphemeralDatabase",
			Handler:    _Quick_CreateEphemeralDatabase_Handler,
		},
		{
			MethodName: "DropEphemeralDatabase",
			Handler:    _Quick_DropEphemeralDatabase_Handler,
		},
		{
			MethodName: "UploadArchive",
			Handler:    _Quick_UploadArchive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/quickdb/v1/quickdb.proto",
}
