/// This file contains service for the Node RPC Server

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: v1/request.proto

// Package Highway is used for defining a Highway node and its accessible API Endpoints

package highwayv1

import (
	_ "go.buf.build/grpc/go/sonr-io/sonr/bucket"
	_ "go.buf.build/grpc/go/sonr-io/sonr/object"
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

// MsgAccessName represents a request payload to get details from the ".snr" name of a peer
type MsgAccessName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the peer to get the details from
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The public key of the peer to get the details from
	PublicKey string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"` // optional
}

func (x *MsgAccessName) Reset() {
	*x = MsgAccessName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAccessName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAccessName) ProtoMessage() {}

func (x *MsgAccessName) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAccessName.ProtoReflect.Descriptor instead.
func (*MsgAccessName) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *MsgAccessName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MsgAccessName) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// MsgCheckName checks the chain to see ifa  name is available
type MsgCheckName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account address of the name owner
	NameToRegister string `protobuf:"bytes,1,opt,name=nameToRegister,proto3" json:"nameToRegister,omitempty"`
	// key of account that can access names
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *MsgCheckName) Reset() {
	*x = MsgCheckName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCheckName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCheckName) ProtoMessage() {}

func (x *MsgCheckName) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCheckName.ProtoReflect.Descriptor instead.
func (*MsgCheckName) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCheckName) GetNameToRegister() string {
	if x != nil {
		return x.NameToRegister
	}
	return ""
}

func (x *MsgCheckName) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type MsgWebToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The JWT
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *MsgWebToken) Reset() {
	*x = MsgWebToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgWebToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgWebToken) ProtoMessage() {}

func (x *MsgWebToken) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgWebToken.ProtoReflect.Descriptor instead.
func (*MsgWebToken) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{2}
}

func (x *MsgWebToken) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

// MsgAccessService represents a request payload to get the service details of a peer
type MsgAccessService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the peer to get the service details of
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// The metadata for any service information required
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // optional
}

func (x *MsgAccessService) Reset() {
	*x = MsgAccessService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAccessService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAccessService) ProtoMessage() {}

func (x *MsgAccessService) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAccessService.ProtoReflect.Descriptor instead.
func (*MsgAccessService) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{3}
}

func (x *MsgAccessService) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgAccessService) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgListenChannel represents a request payload to subscribe to a channel
type MsgListenChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the channel
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Metadata is additional metadata for the channel
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // optional
}

func (x *MsgListenChannel) Reset() {
	*x = MsgListenChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgListenChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgListenChannel) ProtoMessage() {}

func (x *MsgListenChannel) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgListenChannel.ProtoReflect.Descriptor instead.
func (*MsgListenChannel) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{4}
}

func (x *MsgListenChannel) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgListenChannel) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUploadBlob represents a request payload to upload a blob
type MsgUploadBlob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Label is the label of the blob
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Path is the path of the blob
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Bucket or object DID where the blob is being uploaded to
	RefDid string `protobuf:"bytes,3,opt,name=ref_did,json=refDid,proto3" json:"ref_did,omitempty"`
	// Size is the size of the blob
	Size int64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// LastModified is the last modified time of the blob
	LastModified int64 `protobuf:"varint,5,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
}

func (x *MsgUploadBlob) Reset() {
	*x = MsgUploadBlob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUploadBlob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUploadBlob) ProtoMessage() {}

func (x *MsgUploadBlob) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUploadBlob.ProtoReflect.Descriptor instead.
func (*MsgUploadBlob) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{5}
}

func (x *MsgUploadBlob) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MsgUploadBlob) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MsgUploadBlob) GetRefDid() string {
	if x != nil {
		return x.RefDid
	}
	return ""
}

func (x *MsgUploadBlob) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MsgUploadBlob) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

// MsgDownloadBlob represents a request payload to download a blob
type MsgDownloadBlob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the identifier of the blob
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Out Path is the download path of the blob
	OutPath string `protobuf:"bytes,2,opt,name=out_path,json=outPath,proto3" json:"out_path,omitempty"`
}

func (x *MsgDownloadBlob) Reset() {
	*x = MsgDownloadBlob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDownloadBlob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDownloadBlob) ProtoMessage() {}

func (x *MsgDownloadBlob) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDownloadBlob.ProtoReflect.Descriptor instead.
func (*MsgDownloadBlob) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{6}
}

func (x *MsgDownloadBlob) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgDownloadBlob) GetOutPath() string {
	if x != nil {
		return x.OutPath
	}
	return ""
}

// MsgSyncBlob represents a request payload to sync a blob
type MsgSyncBlob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the identifier of the blob
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Path is the location of the blob
	Path    string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"` // optional
}

func (x *MsgSyncBlob) Reset() {
	*x = MsgSyncBlob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSyncBlob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSyncBlob) ProtoMessage() {}

func (x *MsgSyncBlob) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSyncBlob.ProtoReflect.Descriptor instead.
func (*MsgSyncBlob) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{7}
}

func (x *MsgSyncBlob) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgSyncBlob) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MsgSyncBlob) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

// MsgDeleteBlob represents a request payload to delete a blob
type MsgDeleteBlob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the identifier of the blob
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Metadata is the metadata of the blob thats being deleted
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Public key of the node that is deleting the blob
	PublicKey string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *MsgDeleteBlob) Reset() {
	*x = MsgDeleteBlob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeleteBlob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeleteBlob) ProtoMessage() {}

func (x *MsgDeleteBlob) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeleteBlob.ProtoReflect.Descriptor instead.
func (*MsgDeleteBlob) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{8}
}

func (x *MsgDeleteBlob) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgDeleteBlob) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgDeleteBlob) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// MsgParseDid represents a request payload to convert a string to a DID object
type MsgParseDid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the DID of the DID
	DidString string `protobuf:"bytes,1,opt,name=did_string,json=didString,proto3" json:"did_string,omitempty"`
	// Metadata is the metadata of the blob thats being deleted
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MsgParseDid) Reset() {
	*x = MsgParseDid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgParseDid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgParseDid) ProtoMessage() {}

func (x *MsgParseDid) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgParseDid.ProtoReflect.Descriptor instead.
func (*MsgParseDid) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{9}
}

func (x *MsgParseDid) GetDidString() string {
	if x != nil {
		return x.DidString
	}
	return ""
}

func (x *MsgParseDid) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgResolveDid represents a request payload to resolve a DID
type MsgResolveDid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the DID of the DID
	DidString string `protobuf:"bytes,1,opt,name=did_string,json=didString,proto3" json:"did_string,omitempty"`
	// Metadata is the metadata of the blob thats being deleted
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MsgResolveDid) Reset() {
	*x = MsgResolveDid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgResolveDid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgResolveDid) ProtoMessage() {}

func (x *MsgResolveDid) ProtoReflect() protoreflect.Message {
	mi := &file_v1_request_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgResolveDid.ProtoReflect.Descriptor instead.
func (*MsgResolveDid) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{10}
}

func (x *MsgResolveDid) GetDidString() string {
	if x != nil {
		return x.DidString
	}
	return ""
}

func (x *MsgResolveDid) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_v1_request_proto protoreflect.FileDescriptor

var file_v1_request_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x42, 0x0a, 0x0d, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x50, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x61, 0x6d,
	0x65, 0x54, 0x6f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x1f, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x57, 0x65, 0x62, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x10, 0x4d, 0x73, 0x67, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x4d, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb0, 0x01, 0x0a, 0x10, 0x4d, 0x73,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64,
	0x12, 0x4d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x68, 0x69, 0x67, 0x68,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8b, 0x01, 0x0a,
	0x0d, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x5f,
	0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x66, 0x44, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x0f, 0x4d, 0x73,
	0x67, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4d, 0x0a, 0x0b, 0x4d, 0x73,
	0x67, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xc9, 0x01, 0x0a, 0x0d, 0x4d, 0x73,
	0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x4a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb3, 0x01, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x44, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x64, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x48, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e,
	0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x44, 0x69, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb7, 0x01, 0x0a, 0x0d,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x4a, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6f,
	0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_request_proto_rawDescOnce sync.Once
	file_v1_request_proto_rawDescData = file_v1_request_proto_rawDesc
)

func file_v1_request_proto_rawDescGZIP() []byte {
	file_v1_request_proto_rawDescOnce.Do(func() {
		file_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_request_proto_rawDescData)
	})
	return file_v1_request_proto_rawDescData
}

var file_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_v1_request_proto_goTypes = []interface{}{
	(*MsgAccessName)(nil),    // 0: sonrio.highway.v1.MsgAccessName
	(*MsgCheckName)(nil),     // 1: sonrio.highway.v1.MsgCheckName
	(*MsgWebToken)(nil),      // 2: sonrio.highway.v1.MsgWebToken
	(*MsgAccessService)(nil), // 3: sonrio.highway.v1.MsgAccessService
	(*MsgListenChannel)(nil), // 4: sonrio.highway.v1.MsgListenChannel
	(*MsgUploadBlob)(nil),    // 5: sonrio.highway.v1.MsgUploadBlob
	(*MsgDownloadBlob)(nil),  // 6: sonrio.highway.v1.MsgDownloadBlob
	(*MsgSyncBlob)(nil),      // 7: sonrio.highway.v1.MsgSyncBlob
	(*MsgDeleteBlob)(nil),    // 8: sonrio.highway.v1.MsgDeleteBlob
	(*MsgParseDid)(nil),      // 9: sonrio.highway.v1.MsgParseDid
	(*MsgResolveDid)(nil),    // 10: sonrio.highway.v1.MsgResolveDid
	nil,                      // 11: sonrio.highway.v1.MsgAccessService.MetadataEntry
	nil,                      // 12: sonrio.highway.v1.MsgListenChannel.MetadataEntry
	nil,                      // 13: sonrio.highway.v1.MsgDeleteBlob.MetadataEntry
	nil,                      // 14: sonrio.highway.v1.MsgParseDid.MetadataEntry
	nil,                      // 15: sonrio.highway.v1.MsgResolveDid.MetadataEntry
}
var file_v1_request_proto_depIdxs = []int32{
	11, // 0: sonrio.highway.v1.MsgAccessService.metadata:type_name -> sonrio.highway.v1.MsgAccessService.MetadataEntry
	12, // 1: sonrio.highway.v1.MsgListenChannel.metadata:type_name -> sonrio.highway.v1.MsgListenChannel.MetadataEntry
	13, // 2: sonrio.highway.v1.MsgDeleteBlob.metadata:type_name -> sonrio.highway.v1.MsgDeleteBlob.MetadataEntry
	14, // 3: sonrio.highway.v1.MsgParseDid.metadata:type_name -> sonrio.highway.v1.MsgParseDid.MetadataEntry
	15, // 4: sonrio.highway.v1.MsgResolveDid.metadata:type_name -> sonrio.highway.v1.MsgResolveDid.MetadataEntry
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_v1_request_proto_init() }
func file_v1_request_proto_init() {
	if File_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAccessName); i {
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
		file_v1_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCheckName); i {
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
		file_v1_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgWebToken); i {
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
		file_v1_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAccessService); i {
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
		file_v1_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgListenChannel); i {
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
		file_v1_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUploadBlob); i {
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
		file_v1_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDownloadBlob); i {
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
		file_v1_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSyncBlob); i {
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
		file_v1_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeleteBlob); i {
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
		file_v1_request_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgParseDid); i {
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
		file_v1_request_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgResolveDid); i {
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
			RawDescriptor: file_v1_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_request_proto_goTypes,
		DependencyIndexes: file_v1_request_proto_depIdxs,
		MessageInfos:      file_v1_request_proto_msgTypes,
	}.Build()
	File_v1_request_proto = out.File
	file_v1_request_proto_rawDesc = nil
	file_v1_request_proto_goTypes = nil
	file_v1_request_proto_depIdxs = nil
}
