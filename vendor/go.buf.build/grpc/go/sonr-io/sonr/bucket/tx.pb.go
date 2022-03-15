// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: bucket/tx.proto

package bucket

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

type MsgCreateBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Label       string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Kind        string `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *MsgCreateBucket) Reset() {
	*x = MsgCreateBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBucket) ProtoMessage() {}

func (x *MsgCreateBucket) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateBucket.ProtoReflect.Descriptor instead.
func (*MsgCreateBucket) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateBucket) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *MsgCreateBucket) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MsgCreateBucket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MsgCreateBucket) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type MsgCreateBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCreateBucketResponse) Reset() {
	*x = MsgCreateBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateBucketResponse) ProtoMessage() {}

func (x *MsgCreateBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateBucketResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateBucketResponse) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{1}
}

type MsgReadBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Did     string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
}

func (x *MsgReadBucket) Reset() {
	*x = MsgReadBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgReadBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReadBucket) ProtoMessage() {}

func (x *MsgReadBucket) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReadBucket.ProtoReflect.Descriptor instead.
func (*MsgReadBucket) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgReadBucket) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *MsgReadBucket) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

type MsgReadBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgReadBucketResponse) Reset() {
	*x = MsgReadBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgReadBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReadBucketResponse) ProtoMessage() {}

func (x *MsgReadBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReadBucketResponse.ProtoReflect.Descriptor instead.
func (*MsgReadBucketResponse) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{3}
}

type MsgUpdateBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Did         string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Label       string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MsgUpdateBucket) Reset() {
	*x = MsgUpdateBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateBucket) ProtoMessage() {}

func (x *MsgUpdateBucket) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdateBucket.ProtoReflect.Descriptor instead.
func (*MsgUpdateBucket) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgUpdateBucket) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *MsgUpdateBucket) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgUpdateBucket) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MsgUpdateBucket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type MsgUpdateBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateBucketResponse) Reset() {
	*x = MsgUpdateBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateBucketResponse) ProtoMessage() {}

func (x *MsgUpdateBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdateBucketResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateBucketResponse) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{5}
}

type MsgDeleteBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Did       string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	PublicKey string `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *MsgDeleteBucket) Reset() {
	*x = MsgDeleteBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeleteBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteBucket) ProtoMessage() {}

func (x *MsgDeleteBucket) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteBucket.ProtoReflect.Descriptor instead.
func (*MsgDeleteBucket) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{6}
}

func (x *MsgDeleteBucket) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *MsgDeleteBucket) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgDeleteBucket) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type MsgDeleteBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDeleteBucketResponse) Reset() {
	*x = MsgDeleteBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeleteBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteBucketResponse) ProtoMessage() {}

func (x *MsgDeleteBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteBucketResponse.ProtoReflect.Descriptor instead.
func (*MsgDeleteBucketResponse) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{7}
}

type MsgCreateWhichIs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index        string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Did          string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
	DocumentJson string `protobuf:"bytes,4,opt,name=documentJson,proto3" json:"documentJson,omitempty"`
}

func (x *MsgCreateWhichIs) Reset() {
	*x = MsgCreateWhichIs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateWhichIs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateWhichIs) ProtoMessage() {}

func (x *MsgCreateWhichIs) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateWhichIs.ProtoReflect.Descriptor instead.
func (*MsgCreateWhichIs) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{8}
}

func (x *MsgCreateWhichIs) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *MsgCreateWhichIs) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *MsgCreateWhichIs) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgCreateWhichIs) GetDocumentJson() string {
	if x != nil {
		return x.DocumentJson
	}
	return ""
}

type MsgCreateWhichIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCreateWhichIsResponse) Reset() {
	*x = MsgCreateWhichIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateWhichIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateWhichIsResponse) ProtoMessage() {}

func (x *MsgCreateWhichIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateWhichIsResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateWhichIsResponse) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{9}
}

type MsgUpdateWhichIs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index        string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Did          string `protobuf:"bytes,3,opt,name=did,proto3" json:"did,omitempty"`
	DocumentJson string `protobuf:"bytes,4,opt,name=documentJson,proto3" json:"documentJson,omitempty"`
}

func (x *MsgUpdateWhichIs) Reset() {
	*x = MsgUpdateWhichIs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateWhichIs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateWhichIs) ProtoMessage() {}

func (x *MsgUpdateWhichIs) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdateWhichIs.ProtoReflect.Descriptor instead.
func (*MsgUpdateWhichIs) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{10}
}

func (x *MsgUpdateWhichIs) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *MsgUpdateWhichIs) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *MsgUpdateWhichIs) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgUpdateWhichIs) GetDocumentJson() string {
	if x != nil {
		return x.DocumentJson
	}
	return ""
}

type MsgUpdateWhichIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateWhichIsResponse) Reset() {
	*x = MsgUpdateWhichIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateWhichIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateWhichIsResponse) ProtoMessage() {}

func (x *MsgUpdateWhichIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdateWhichIsResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateWhichIsResponse) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{11}
}

type MsgDeleteWhichIs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *MsgDeleteWhichIs) Reset() {
	*x = MsgDeleteWhichIs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeleteWhichIs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteWhichIs) ProtoMessage() {}

func (x *MsgDeleteWhichIs) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteWhichIs.ProtoReflect.Descriptor instead.
func (*MsgDeleteWhichIs) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{12}
}

func (x *MsgDeleteWhichIs) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *MsgDeleteWhichIs) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

type MsgDeleteWhichIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDeleteWhichIsResponse) Reset() {
	*x = MsgDeleteWhichIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_tx_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeleteWhichIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteWhichIsResponse) ProtoMessage() {}

func (x *MsgDeleteWhichIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_tx_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteWhichIsResponse.ProtoReflect.Descriptor instead.
func (*MsgDeleteWhichIsResponse) Descriptor() ([]byte, []int) {
	return file_bucket_tx_proto_rawDescGZIP(), []int{13}
}

var File_bucket_tx_proto protoreflect.FileDescriptor

var file_bucket_tx_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x15, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x77, 0x68,
	0x69, 0x63, 0x68, 0x5f, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x0f,
	0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3b, 0x0a, 0x0d, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a,
	0x17, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x78, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x68, 0x69,
	0x63, 0x68, 0x49, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x4d, 0x73,
	0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x78, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e,
	0x22, 0x1a, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x68, 0x69,
	0x63, 0x68, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x10,
	0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x1a, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x68, 0x69,
	0x63, 0x68, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb6, 0x05, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x60, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61,
	0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x1a, 0x2c, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63,
	0x68, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x12, 0x24, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68,
	0x49, 0x73, 0x1a, 0x2c, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x63, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49,
	0x73, 0x12, 0x24, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x1a, 0x2c, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x68, 0x69, 0x63, 0x68, 0x49, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6f,
	0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bucket_tx_proto_rawDescOnce sync.Once
	file_bucket_tx_proto_rawDescData = file_bucket_tx_proto_rawDesc
)

func file_bucket_tx_proto_rawDescGZIP() []byte {
	file_bucket_tx_proto_rawDescOnce.Do(func() {
		file_bucket_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_bucket_tx_proto_rawDescData)
	})
	return file_bucket_tx_proto_rawDescData
}

var file_bucket_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_bucket_tx_proto_goTypes = []interface{}{
	(*MsgCreateBucket)(nil),          // 0: sonrio.sonr.bucket.MsgCreateBucket
	(*MsgCreateBucketResponse)(nil),  // 1: sonrio.sonr.bucket.MsgCreateBucketResponse
	(*MsgReadBucket)(nil),            // 2: sonrio.sonr.bucket.MsgReadBucket
	(*MsgReadBucketResponse)(nil),    // 3: sonrio.sonr.bucket.MsgReadBucketResponse
	(*MsgUpdateBucket)(nil),          // 4: sonrio.sonr.bucket.MsgUpdateBucket
	(*MsgUpdateBucketResponse)(nil),  // 5: sonrio.sonr.bucket.MsgUpdateBucketResponse
	(*MsgDeleteBucket)(nil),          // 6: sonrio.sonr.bucket.MsgDeleteBucket
	(*MsgDeleteBucketResponse)(nil),  // 7: sonrio.sonr.bucket.MsgDeleteBucketResponse
	(*MsgCreateWhichIs)(nil),         // 8: sonrio.sonr.bucket.MsgCreateWhichIs
	(*MsgCreateWhichIsResponse)(nil), // 9: sonrio.sonr.bucket.MsgCreateWhichIsResponse
	(*MsgUpdateWhichIs)(nil),         // 10: sonrio.sonr.bucket.MsgUpdateWhichIs
	(*MsgUpdateWhichIsResponse)(nil), // 11: sonrio.sonr.bucket.MsgUpdateWhichIsResponse
	(*MsgDeleteWhichIs)(nil),         // 12: sonrio.sonr.bucket.MsgDeleteWhichIs
	(*MsgDeleteWhichIsResponse)(nil), // 13: sonrio.sonr.bucket.MsgDeleteWhichIsResponse
}
var file_bucket_tx_proto_depIdxs = []int32{
	0,  // 0: sonrio.sonr.bucket.Msg.CreateBucket:input_type -> sonrio.sonr.bucket.MsgCreateBucket
	2,  // 1: sonrio.sonr.bucket.Msg.ReadBucket:input_type -> sonrio.sonr.bucket.MsgReadBucket
	4,  // 2: sonrio.sonr.bucket.Msg.UpdateBucket:input_type -> sonrio.sonr.bucket.MsgUpdateBucket
	6,  // 3: sonrio.sonr.bucket.Msg.DeleteBucket:input_type -> sonrio.sonr.bucket.MsgDeleteBucket
	8,  // 4: sonrio.sonr.bucket.Msg.CreateWhichIs:input_type -> sonrio.sonr.bucket.MsgCreateWhichIs
	10, // 5: sonrio.sonr.bucket.Msg.UpdateWhichIs:input_type -> sonrio.sonr.bucket.MsgUpdateWhichIs
	12, // 6: sonrio.sonr.bucket.Msg.DeleteWhichIs:input_type -> sonrio.sonr.bucket.MsgDeleteWhichIs
	1,  // 7: sonrio.sonr.bucket.Msg.CreateBucket:output_type -> sonrio.sonr.bucket.MsgCreateBucketResponse
	3,  // 8: sonrio.sonr.bucket.Msg.ReadBucket:output_type -> sonrio.sonr.bucket.MsgReadBucketResponse
	5,  // 9: sonrio.sonr.bucket.Msg.UpdateBucket:output_type -> sonrio.sonr.bucket.MsgUpdateBucketResponse
	7,  // 10: sonrio.sonr.bucket.Msg.DeleteBucket:output_type -> sonrio.sonr.bucket.MsgDeleteBucketResponse
	9,  // 11: sonrio.sonr.bucket.Msg.CreateWhichIs:output_type -> sonrio.sonr.bucket.MsgCreateWhichIsResponse
	11, // 12: sonrio.sonr.bucket.Msg.UpdateWhichIs:output_type -> sonrio.sonr.bucket.MsgUpdateWhichIsResponse
	13, // 13: sonrio.sonr.bucket.Msg.DeleteWhichIs:output_type -> sonrio.sonr.bucket.MsgDeleteWhichIsResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_bucket_tx_proto_init() }
func file_bucket_tx_proto_init() {
	if File_bucket_tx_proto != nil {
		return
	}
	file_bucket_which_is_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bucket_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBucket); i {
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
		file_bucket_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateBucketResponse); i {
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
		file_bucket_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgReadBucket); i {
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
		file_bucket_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgReadBucketResponse); i {
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
		file_bucket_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateBucket); i {
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
		file_bucket_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateBucketResponse); i {
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
		file_bucket_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeleteBucket); i {
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
		file_bucket_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeleteBucketResponse); i {
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
		file_bucket_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateWhichIs); i {
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
		file_bucket_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateWhichIsResponse); i {
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
		file_bucket_tx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateWhichIs); i {
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
		file_bucket_tx_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateWhichIsResponse); i {
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
		file_bucket_tx_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeleteWhichIs); i {
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
		file_bucket_tx_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeleteWhichIsResponse); i {
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
			RawDescriptor: file_bucket_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bucket_tx_proto_goTypes,
		DependencyIndexes: file_bucket_tx_proto_depIdxs,
		MessageInfos:      file_bucket_tx_proto_msgTypes,
	}.Build()
	File_bucket_tx_proto = out.File
	file_bucket_tx_proto_rawDesc = nil
	file_bucket_tx_proto_goTypes = nil
	file_bucket_tx_proto_depIdxs = nil
}