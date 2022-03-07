/// This file contains service for the Node RPC Server

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: v1/request.proto

// Package Highway is used for defining a Highway node and its properties.

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

// AccessNameRequest is a request to get details from the ".snr" name of a peer
type AccessNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the peer to get the details from
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The public key of the peer to get the details from
	PublicKey string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"` // optional
}

func (x *AccessNameRequest) Reset() {
	*x = AccessNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessNameRequest) ProtoMessage() {}

func (x *AccessNameRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AccessNameRequest.ProtoReflect.Descriptor instead.
func (*AccessNameRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *AccessNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccessNameRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// AccessServiceRequest is a request to get the service details of a peer
type AccessServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the peer to get the service details of
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// The metadata for any service information required
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // optional
}

func (x *AccessServiceRequest) Reset() {
	*x = AccessServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessServiceRequest) ProtoMessage() {}

func (x *AccessServiceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AccessServiceRequest.ProtoReflect.Descriptor instead.
func (*AccessServiceRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{1}
}

func (x *AccessServiceRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *AccessServiceRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ListenChannelRequest is the request to subscribe to a channel
type ListenChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the channel
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Metadata is additional metadata for the channel
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // optional
}

func (x *ListenChannelRequest) Reset() {
	*x = ListenChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenChannelRequest) ProtoMessage() {}

func (x *ListenChannelRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListenChannelRequest.ProtoReflect.Descriptor instead.
func (*ListenChannelRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{2}
}

func (x *ListenChannelRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *ListenChannelRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ListenBucketRequest is the request to subscribe to a bucket
type ListenBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the DID of the bucket
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Metadata is the metadata of the bucket thats being listened to
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // optional
}

func (x *ListenBucketRequest) Reset() {
	*x = ListenBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenBucketRequest) ProtoMessage() {}

func (x *ListenBucketRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListenBucketRequest.ProtoReflect.Descriptor instead.
func (*ListenBucketRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{3}
}

func (x *ListenBucketRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *ListenBucketRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// UploadBlobRequest is the request to upload a blob
type UploadBlobRequest struct {
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

func (x *UploadBlobRequest) Reset() {
	*x = UploadBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBlobRequest) ProtoMessage() {}

func (x *UploadBlobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UploadBlobRequest.ProtoReflect.Descriptor instead.
func (*UploadBlobRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{4}
}

func (x *UploadBlobRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UploadBlobRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UploadBlobRequest) GetRefDid() string {
	if x != nil {
		return x.RefDid
	}
	return ""
}

func (x *UploadBlobRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UploadBlobRequest) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

// DownloadBlobRequest is the request to download a blob
type DownloadBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the identifier of the blob
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Out Path is the download path of the blob
	OutPath string `protobuf:"bytes,2,opt,name=out_path,json=outPath,proto3" json:"out_path,omitempty"`
}

func (x *DownloadBlobRequest) Reset() {
	*x = DownloadBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadBlobRequest) ProtoMessage() {}

func (x *DownloadBlobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownloadBlobRequest.ProtoReflect.Descriptor instead.
func (*DownloadBlobRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadBlobRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *DownloadBlobRequest) GetOutPath() string {
	if x != nil {
		return x.OutPath
	}
	return ""
}

// SyncBlobRequest is the request to sync a blob
type SyncBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the identifier of the blob
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Path is the location of the blob
	Path    string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Timeout int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"` // optional
}

func (x *SyncBlobRequest) Reset() {
	*x = SyncBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncBlobRequest) ProtoMessage() {}

func (x *SyncBlobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SyncBlobRequest.ProtoReflect.Descriptor instead.
func (*SyncBlobRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{6}
}

func (x *SyncBlobRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *SyncBlobRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SyncBlobRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

// DeleteBlobRequest is the request to delete a blob
type DeleteBlobRequest struct {
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

func (x *DeleteBlobRequest) Reset() {
	*x = DeleteBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlobRequest) ProtoMessage() {}

func (x *DeleteBlobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteBlobRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlobRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBlobRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *DeleteBlobRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DeleteBlobRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// ParseDidRequest is the request to convert a string to a DID object
type ParseDidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the DID of the DID
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Metadata is the metadata of the blob thats being deleted
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ParseDidRequest) Reset() {
	*x = ParseDidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseDidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseDidRequest) ProtoMessage() {}

func (x *ParseDidRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ParseDidRequest.ProtoReflect.Descriptor instead.
func (*ParseDidRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{8}
}

func (x *ParseDidRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *ParseDidRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ResolveDidRequest is the request to resolve a DID
type ResolveDidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DID is the DID of the DID
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Metadata is the metadata of the blob thats being deleted
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResolveDidRequest) Reset() {
	*x = ResolveDidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_request_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveDidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveDidRequest) ProtoMessage() {}

func (x *ResolveDidRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ResolveDidRequest.ProtoReflect.Descriptor instead.
func (*ResolveDidRequest) Descriptor() ([]byte, []int) {
	return file_v1_request_proto_rawDescGZIP(), []int{9}
}

func (x *ResolveDidRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *ResolveDidRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_v1_request_proto protoreflect.FileDescriptor

var file_v1_request_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x13,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x22, 0xb1, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xb1, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12,
	0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xaf, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x69, 0x64, 0x12, 0x49, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65,
	0x66, 0x5f, 0x64, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x66,
	0x44, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x42, 0x0a, 0x13,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x51, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xa7, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x73, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x69, 0x64, 0x12, 0x47, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x3b, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*AccessNameRequest)(nil),    // 0: highway.v1.AccessNameRequest
	(*AccessServiceRequest)(nil), // 1: highway.v1.AccessServiceRequest
	(*ListenChannelRequest)(nil), // 2: highway.v1.ListenChannelRequest
	(*ListenBucketRequest)(nil),  // 3: highway.v1.ListenBucketRequest
	(*UploadBlobRequest)(nil),    // 4: highway.v1.UploadBlobRequest
	(*DownloadBlobRequest)(nil),  // 5: highway.v1.DownloadBlobRequest
	(*SyncBlobRequest)(nil),      // 6: highway.v1.SyncBlobRequest
	(*DeleteBlobRequest)(nil),    // 7: highway.v1.DeleteBlobRequest
	(*ParseDidRequest)(nil),      // 8: highway.v1.ParseDidRequest
	(*ResolveDidRequest)(nil),    // 9: highway.v1.ResolveDidRequest
	nil,                          // 10: highway.v1.AccessServiceRequest.MetadataEntry
	nil,                          // 11: highway.v1.ListenChannelRequest.MetadataEntry
	nil,                          // 12: highway.v1.ListenBucketRequest.MetadataEntry
	nil,                          // 13: highway.v1.DeleteBlobRequest.MetadataEntry
	nil,                          // 14: highway.v1.ParseDidRequest.MetadataEntry
	nil,                          // 15: highway.v1.ResolveDidRequest.MetadataEntry
}
var file_v1_request_proto_depIdxs = []int32{
	10, // 0: highway.v1.AccessServiceRequest.metadata:type_name -> highway.v1.AccessServiceRequest.MetadataEntry
	11, // 1: highway.v1.ListenChannelRequest.metadata:type_name -> highway.v1.ListenChannelRequest.MetadataEntry
	12, // 2: highway.v1.ListenBucketRequest.metadata:type_name -> highway.v1.ListenBucketRequest.MetadataEntry
	13, // 3: highway.v1.DeleteBlobRequest.metadata:type_name -> highway.v1.DeleteBlobRequest.MetadataEntry
	14, // 4: highway.v1.ParseDidRequest.metadata:type_name -> highway.v1.ParseDidRequest.MetadataEntry
	15, // 5: highway.v1.ResolveDidRequest.metadata:type_name -> highway.v1.ResolveDidRequest.MetadataEntry
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_v1_request_proto_init() }
func file_v1_request_proto_init() {
	if File_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessNameRequest); i {
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
			switch v := v.(*AccessServiceRequest); i {
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
			switch v := v.(*ListenChannelRequest); i {
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
			switch v := v.(*ListenBucketRequest); i {
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
			switch v := v.(*UploadBlobRequest); i {
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
			switch v := v.(*DownloadBlobRequest); i {
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
			switch v := v.(*SyncBlobRequest); i {
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
			switch v := v.(*DeleteBlobRequest); i {
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
			switch v := v.(*ParseDidRequest); i {
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
			switch v := v.(*ResolveDidRequest); i {
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