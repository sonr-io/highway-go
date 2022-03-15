// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: object/query.proto

package object

import (
	v1beta1 "go.buf.build/grpc/go/cosmos/cosmos-sdk/cosmos/base/query/v1beta1"
	_ "go.buf.build/grpc/go/cosmos/gogo-proto/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryParamsRequest) Reset() {
	*x = QueryParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsRequest) ProtoMessage() {}

func (x *QueryParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_object_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsRequest.ProtoReflect.Descriptor instead.
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return file_object_query_proto_rawDescGZIP(), []int{0}
}

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params holds all the parameters of this module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *QueryParamsResponse) Reset() {
	*x = QueryParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsResponse) ProtoMessage() {}

func (x *QueryParamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_object_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsResponse.ProtoReflect.Descriptor instead.
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return file_object_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryParamsResponse) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type QueryGetWhatIsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *QueryGetWhatIsRequest) Reset() {
	*x = QueryGetWhatIsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGetWhatIsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGetWhatIsRequest) ProtoMessage() {}

func (x *QueryGetWhatIsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_object_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGetWhatIsRequest.ProtoReflect.Descriptor instead.
func (*QueryGetWhatIsRequest) Descriptor() ([]byte, []int) {
	return file_object_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryGetWhatIsRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

type QueryGetWhatIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WhatIs *WhatIs `protobuf:"bytes,1,opt,name=whatIs,proto3" json:"whatIs,omitempty"`
}

func (x *QueryGetWhatIsResponse) Reset() {
	*x = QueryGetWhatIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGetWhatIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGetWhatIsResponse) ProtoMessage() {}

func (x *QueryGetWhatIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_object_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGetWhatIsResponse.ProtoReflect.Descriptor instead.
func (*QueryGetWhatIsResponse) Descriptor() ([]byte, []int) {
	return file_object_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryGetWhatIsResponse) GetWhatIs() *WhatIs {
	if x != nil {
		return x.WhatIs
	}
	return nil
}

type QueryAllWhatIsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *v1beta1.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllWhatIsRequest) Reset() {
	*x = QueryAllWhatIsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllWhatIsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllWhatIsRequest) ProtoMessage() {}

func (x *QueryAllWhatIsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_object_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllWhatIsRequest.ProtoReflect.Descriptor instead.
func (*QueryAllWhatIsRequest) Descriptor() ([]byte, []int) {
	return file_object_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryAllWhatIsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type QueryAllWhatIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WhatIs     []*WhatIs             `protobuf:"bytes,1,rep,name=whatIs,proto3" json:"whatIs,omitempty"`
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllWhatIsResponse) Reset() {
	*x = QueryAllWhatIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_object_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllWhatIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllWhatIsResponse) ProtoMessage() {}

func (x *QueryAllWhatIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_object_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllWhatIsResponse.ProtoReflect.Descriptor instead.
func (*QueryAllWhatIsResponse) Descriptor() ([]byte, []int) {
	return file_object_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryAllWhatIsResponse) GetWhatIs() []*WhatIs {
	if x != nil {
		return x.WhatIs
	}
	return nil
}

func (x *QueryAllWhatIsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_object_query_proto protoreflect.FileDescriptor

var file_object_query_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x13, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2d, 0x0a, 0x15, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x47, 0x65, 0x74, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x52, 0x0a, 0x16, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x47, 0x65, 0x74, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x77, 0x68, 0x61, 0x74, 0x49, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x77, 0x68, 0x61, 0x74, 0x49, 0x73, 0x22, 0x5f, 0x0a,
	0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9b,
	0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x57, 0x68, 0x61, 0x74, 0x49,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x77, 0x68, 0x61,
	0x74, 0x49, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x57,
	0x68, 0x61, 0x74, 0x49, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x77, 0x68, 0x61,
	0x74, 0x49, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa1, 0x03, 0x0a,
	0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x7d, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x26, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69,
	0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x06, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73,
	0x12, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x65, 0x74, 0x57, 0x68,
	0x61, 0x74, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x65, 0x74, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12,
	0x24, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x73, 0x2f, 0x7b, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x7d, 0x12, 0x88, 0x01, 0x0a, 0x09, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73,
	0x41, 0x6c, 0x6c, 0x12, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c,
	0x6c, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x57, 0x68, 0x61, 0x74,
	0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6e,
	0x72, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x73,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_object_query_proto_rawDescOnce sync.Once
	file_object_query_proto_rawDescData = file_object_query_proto_rawDesc
)

func file_object_query_proto_rawDescGZIP() []byte {
	file_object_query_proto_rawDescOnce.Do(func() {
		file_object_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_object_query_proto_rawDescData)
	})
	return file_object_query_proto_rawDescData
}

var file_object_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_object_query_proto_goTypes = []interface{}{
	(*QueryParamsRequest)(nil),     // 0: sonrio.sonr.object.QueryParamsRequest
	(*QueryParamsResponse)(nil),    // 1: sonrio.sonr.object.QueryParamsResponse
	(*QueryGetWhatIsRequest)(nil),  // 2: sonrio.sonr.object.QueryGetWhatIsRequest
	(*QueryGetWhatIsResponse)(nil), // 3: sonrio.sonr.object.QueryGetWhatIsResponse
	(*QueryAllWhatIsRequest)(nil),  // 4: sonrio.sonr.object.QueryAllWhatIsRequest
	(*QueryAllWhatIsResponse)(nil), // 5: sonrio.sonr.object.QueryAllWhatIsResponse
	(*Params)(nil),                 // 6: sonrio.sonr.object.Params
	(*WhatIs)(nil),                 // 7: sonrio.sonr.object.WhatIs
	(*v1beta1.PageRequest)(nil),    // 8: cosmos.base.query.v1beta1.PageRequest
	(*v1beta1.PageResponse)(nil),   // 9: cosmos.base.query.v1beta1.PageResponse
}
var file_object_query_proto_depIdxs = []int32{
	6, // 0: sonrio.sonr.object.QueryParamsResponse.params:type_name -> sonrio.sonr.object.Params
	7, // 1: sonrio.sonr.object.QueryGetWhatIsResponse.whatIs:type_name -> sonrio.sonr.object.WhatIs
	8, // 2: sonrio.sonr.object.QueryAllWhatIsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	7, // 3: sonrio.sonr.object.QueryAllWhatIsResponse.whatIs:type_name -> sonrio.sonr.object.WhatIs
	9, // 4: sonrio.sonr.object.QueryAllWhatIsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	0, // 5: sonrio.sonr.object.Query.Params:input_type -> sonrio.sonr.object.QueryParamsRequest
	2, // 6: sonrio.sonr.object.Query.WhatIs:input_type -> sonrio.sonr.object.QueryGetWhatIsRequest
	4, // 7: sonrio.sonr.object.Query.WhatIsAll:input_type -> sonrio.sonr.object.QueryAllWhatIsRequest
	1, // 8: sonrio.sonr.object.Query.Params:output_type -> sonrio.sonr.object.QueryParamsResponse
	3, // 9: sonrio.sonr.object.Query.WhatIs:output_type -> sonrio.sonr.object.QueryGetWhatIsResponse
	5, // 10: sonrio.sonr.object.Query.WhatIsAll:output_type -> sonrio.sonr.object.QueryAllWhatIsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_object_query_proto_init() }
func file_object_query_proto_init() {
	if File_object_query_proto != nil {
		return
	}
	file_object_params_proto_init()
	file_object_what_is_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_object_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsRequest); i {
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
		file_object_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsResponse); i {
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
		file_object_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGetWhatIsRequest); i {
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
		file_object_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGetWhatIsResponse); i {
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
		file_object_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllWhatIsRequest); i {
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
		file_object_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllWhatIsResponse); i {
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
			RawDescriptor: file_object_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_object_query_proto_goTypes,
		DependencyIndexes: file_object_query_proto_depIdxs,
		MessageInfos:      file_object_query_proto_msgTypes,
	}.Build()
	File_object_query_proto = out.File
	file_object_query_proto_rawDesc = nil
	file_object_query_proto_goTypes = nil
	file_object_query_proto_depIdxs = nil
}