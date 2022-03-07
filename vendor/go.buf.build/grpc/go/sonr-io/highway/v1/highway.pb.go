// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: v1/highway.proto

// Package Highway is used for defining a Highway node and its properties.

package highwayv1

import (
	bucket "go.buf.build/grpc/go/sonr-io/sonr/bucket"
	channel "go.buf.build/grpc/go/sonr-io/sonr/channel"
	object "go.buf.build/grpc/go/sonr-io/sonr/object"
	registry "go.buf.build/grpc/go/sonr-io/sonr/registry"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v1_highway_proto protoreflect.FileDescriptor

var file_v1_highway_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x15,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x74,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x9c, 0x13, 0x0a, 0x0e, 0x48, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x2d, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x2b,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69,
	0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x30, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d,
	0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x2e, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x67, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x25, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x2d, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0b, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x2b,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x25,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x2d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x73,
	0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x2d,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x20, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x62, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5c, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x1a, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x62, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x62, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5c, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x21,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x1a, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x62, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x73, 0x67,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1f, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x49,
	0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1b, 0x2e, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x44, 0x69, 0x64, 0x12, 0x1b, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x12,
	0x1d, 0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x44, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x69, 0x67, 0x68,
	0x77, 0x61, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_v1_highway_proto_goTypes = []interface{}{
	(*AccessNameRequest)(nil),                   // 0: highway.v1.AccessNameRequest
	(*registry.MsgRegisterName)(nil),            // 1: sonrio.sonr.registry.MsgRegisterName
	(*registry.MsgUpdateName)(nil),              // 2: sonrio.sonr.registry.MsgUpdateName
	(*AccessServiceRequest)(nil),                // 3: highway.v1.AccessServiceRequest
	(*registry.MsgRegisterService)(nil),         // 4: sonrio.sonr.registry.MsgRegisterService
	(*registry.MsgUpdateService)(nil),           // 5: sonrio.sonr.registry.MsgUpdateService
	(*channel.MsgCreateChannel)(nil),            // 6: sonrio.sonr.channel.MsgCreateChannel
	(*channel.MsgReadChannel)(nil),              // 7: sonrio.sonr.channel.MsgReadChannel
	(*channel.MsgUpdateChannel)(nil),            // 8: sonrio.sonr.channel.MsgUpdateChannel
	(*channel.MsgDeleteChannel)(nil),            // 9: sonrio.sonr.channel.MsgDeleteChannel
	(*ListenChannelRequest)(nil),                // 10: highway.v1.ListenChannelRequest
	(*bucket.MsgCreateBucket)(nil),              // 11: sonrio.sonr.bucket.MsgCreateBucket
	(*bucket.MsgReadBucket)(nil),                // 12: sonrio.sonr.bucket.MsgReadBucket
	(*bucket.MsgUpdateBucket)(nil),              // 13: sonrio.sonr.bucket.MsgUpdateBucket
	(*bucket.MsgDeleteBucket)(nil),              // 14: sonrio.sonr.bucket.MsgDeleteBucket
	(*ListenBucketRequest)(nil),                 // 15: highway.v1.ListenBucketRequest
	(*object.MsgCreateObject)(nil),              // 16: sonrio.sonr.object.MsgCreateObject
	(*object.MsgReadObject)(nil),                // 17: sonrio.sonr.object.MsgReadObject
	(*object.MsgUpdateObject)(nil),              // 18: sonrio.sonr.object.MsgUpdateObject
	(*object.MsgDeleteObject)(nil),              // 19: sonrio.sonr.object.MsgDeleteObject
	(*UploadBlobRequest)(nil),                   // 20: highway.v1.UploadBlobRequest
	(*DownloadBlobRequest)(nil),                 // 21: highway.v1.DownloadBlobRequest
	(*SyncBlobRequest)(nil),                     // 22: highway.v1.SyncBlobRequest
	(*DeleteBlobRequest)(nil),                   // 23: highway.v1.DeleteBlobRequest
	(*ParseDidRequest)(nil),                     // 24: highway.v1.ParseDidRequest
	(*ResolveDidRequest)(nil),                   // 25: highway.v1.ResolveDidRequest
	(*AccessNameResponse)(nil),                  // 26: highway.v1.AccessNameResponse
	(*registry.MsgRegisterNameResponse)(nil),    // 27: sonrio.sonr.registry.MsgRegisterNameResponse
	(*registry.MsgUpdateNameResponse)(nil),      // 28: sonrio.sonr.registry.MsgUpdateNameResponse
	(*AccessServiceResponse)(nil),               // 29: highway.v1.AccessServiceResponse
	(*registry.MsgRegisterServiceResponse)(nil), // 30: sonrio.sonr.registry.MsgRegisterServiceResponse
	(*registry.MsgUpdateServiceResponse)(nil),   // 31: sonrio.sonr.registry.MsgUpdateServiceResponse
	(*channel.MsgCreateChannelResponse)(nil),    // 32: sonrio.sonr.channel.MsgCreateChannelResponse
	(*channel.MsgReadChannelResponse)(nil),      // 33: sonrio.sonr.channel.MsgReadChannelResponse
	(*channel.MsgUpdateChannelResponse)(nil),    // 34: sonrio.sonr.channel.MsgUpdateChannelResponse
	(*channel.MsgDeleteChannelResponse)(nil),    // 35: sonrio.sonr.channel.MsgDeleteChannelResponse
	(*channel.ChannelMessage)(nil),              // 36: sonrio.sonr.channel.ChannelMessage
	(*bucket.MsgCreateBucketResponse)(nil),      // 37: sonrio.sonr.bucket.MsgCreateBucketResponse
	(*bucket.MsgReadBucketResponse)(nil),        // 38: sonrio.sonr.bucket.MsgReadBucketResponse
	(*bucket.MsgUpdateBucketResponse)(nil),      // 39: sonrio.sonr.bucket.MsgUpdateBucketResponse
	(*bucket.MsgDeleteBucketResponse)(nil),      // 40: sonrio.sonr.bucket.MsgDeleteBucketResponse
	(*ListenBucketResponse)(nil),                // 41: highway.v1.ListenBucketResponse
	(*object.MsgCreateObjectResponse)(nil),      // 42: sonrio.sonr.object.MsgCreateObjectResponse
	(*object.MsgReadObjectResponse)(nil),        // 43: sonrio.sonr.object.MsgReadObjectResponse
	(*object.MsgUpdateObjectResponse)(nil),      // 44: sonrio.sonr.object.MsgUpdateObjectResponse
	(*object.MsgDeleteObjectResponse)(nil),      // 45: sonrio.sonr.object.MsgDeleteObjectResponse
	(*UploadBlobResponse)(nil),                  // 46: highway.v1.UploadBlobResponse
	(*DownloadBlobResponse)(nil),                // 47: highway.v1.DownloadBlobResponse
	(*SyncBlobResponse)(nil),                    // 48: highway.v1.SyncBlobResponse
	(*DeleteBlobResponse)(nil),                  // 49: highway.v1.DeleteBlobResponse
	(*ParseDidResponse)(nil),                    // 50: highway.v1.ParseDidResponse
	(*ResolveDidResponse)(nil),                  // 51: highway.v1.ResolveDidResponse
}
var file_v1_highway_proto_depIdxs = []int32{
	0,  // 0: highway.v1.HighwayService.AccessName:input_type -> highway.v1.AccessNameRequest
	1,  // 1: highway.v1.HighwayService.RegisterName:input_type -> sonrio.sonr.registry.MsgRegisterName
	2,  // 2: highway.v1.HighwayService.UpdateName:input_type -> sonrio.sonr.registry.MsgUpdateName
	3,  // 3: highway.v1.HighwayService.AccessService:input_type -> highway.v1.AccessServiceRequest
	4,  // 4: highway.v1.HighwayService.RegisterService:input_type -> sonrio.sonr.registry.MsgRegisterService
	5,  // 5: highway.v1.HighwayService.UpdateService:input_type -> sonrio.sonr.registry.MsgUpdateService
	6,  // 6: highway.v1.HighwayService.CreateChannel:input_type -> sonrio.sonr.channel.MsgCreateChannel
	7,  // 7: highway.v1.HighwayService.ReadChannel:input_type -> sonrio.sonr.channel.MsgReadChannel
	8,  // 8: highway.v1.HighwayService.UpdateChannel:input_type -> sonrio.sonr.channel.MsgUpdateChannel
	9,  // 9: highway.v1.HighwayService.DeleteChannel:input_type -> sonrio.sonr.channel.MsgDeleteChannel
	10, // 10: highway.v1.HighwayService.ListenChannel:input_type -> highway.v1.ListenChannelRequest
	11, // 11: highway.v1.HighwayService.CreateBucket:input_type -> sonrio.sonr.bucket.MsgCreateBucket
	12, // 12: highway.v1.HighwayService.ReadBucket:input_type -> sonrio.sonr.bucket.MsgReadBucket
	13, // 13: highway.v1.HighwayService.UpdateBucket:input_type -> sonrio.sonr.bucket.MsgUpdateBucket
	14, // 14: highway.v1.HighwayService.DeleteBucket:input_type -> sonrio.sonr.bucket.MsgDeleteBucket
	15, // 15: highway.v1.HighwayService.ListenBucket:input_type -> highway.v1.ListenBucketRequest
	16, // 16: highway.v1.HighwayService.CreateObject:input_type -> sonrio.sonr.object.MsgCreateObject
	17, // 17: highway.v1.HighwayService.ReadObject:input_type -> sonrio.sonr.object.MsgReadObject
	18, // 18: highway.v1.HighwayService.UpdateObject:input_type -> sonrio.sonr.object.MsgUpdateObject
	19, // 19: highway.v1.HighwayService.DeleteObject:input_type -> sonrio.sonr.object.MsgDeleteObject
	20, // 20: highway.v1.HighwayService.UploadBlob:input_type -> highway.v1.UploadBlobRequest
	21, // 21: highway.v1.HighwayService.DownloadBlob:input_type -> highway.v1.DownloadBlobRequest
	22, // 22: highway.v1.HighwayService.SyncBlob:input_type -> highway.v1.SyncBlobRequest
	23, // 23: highway.v1.HighwayService.DeleteBlob:input_type -> highway.v1.DeleteBlobRequest
	24, // 24: highway.v1.HighwayService.ParseDid:input_type -> highway.v1.ParseDidRequest
	25, // 25: highway.v1.HighwayService.ResolveDid:input_type -> highway.v1.ResolveDidRequest
	26, // 26: highway.v1.HighwayService.AccessName:output_type -> highway.v1.AccessNameResponse
	27, // 27: highway.v1.HighwayService.RegisterName:output_type -> sonrio.sonr.registry.MsgRegisterNameResponse
	28, // 28: highway.v1.HighwayService.UpdateName:output_type -> sonrio.sonr.registry.MsgUpdateNameResponse
	29, // 29: highway.v1.HighwayService.AccessService:output_type -> highway.v1.AccessServiceResponse
	30, // 30: highway.v1.HighwayService.RegisterService:output_type -> sonrio.sonr.registry.MsgRegisterServiceResponse
	31, // 31: highway.v1.HighwayService.UpdateService:output_type -> sonrio.sonr.registry.MsgUpdateServiceResponse
	32, // 32: highway.v1.HighwayService.CreateChannel:output_type -> sonrio.sonr.channel.MsgCreateChannelResponse
	33, // 33: highway.v1.HighwayService.ReadChannel:output_type -> sonrio.sonr.channel.MsgReadChannelResponse
	34, // 34: highway.v1.HighwayService.UpdateChannel:output_type -> sonrio.sonr.channel.MsgUpdateChannelResponse
	35, // 35: highway.v1.HighwayService.DeleteChannel:output_type -> sonrio.sonr.channel.MsgDeleteChannelResponse
	36, // 36: highway.v1.HighwayService.ListenChannel:output_type -> sonrio.sonr.channel.ChannelMessage
	37, // 37: highway.v1.HighwayService.CreateBucket:output_type -> sonrio.sonr.bucket.MsgCreateBucketResponse
	38, // 38: highway.v1.HighwayService.ReadBucket:output_type -> sonrio.sonr.bucket.MsgReadBucketResponse
	39, // 39: highway.v1.HighwayService.UpdateBucket:output_type -> sonrio.sonr.bucket.MsgUpdateBucketResponse
	40, // 40: highway.v1.HighwayService.DeleteBucket:output_type -> sonrio.sonr.bucket.MsgDeleteBucketResponse
	41, // 41: highway.v1.HighwayService.ListenBucket:output_type -> highway.v1.ListenBucketResponse
	42, // 42: highway.v1.HighwayService.CreateObject:output_type -> sonrio.sonr.object.MsgCreateObjectResponse
	43, // 43: highway.v1.HighwayService.ReadObject:output_type -> sonrio.sonr.object.MsgReadObjectResponse
	44, // 44: highway.v1.HighwayService.UpdateObject:output_type -> sonrio.sonr.object.MsgUpdateObjectResponse
	45, // 45: highway.v1.HighwayService.DeleteObject:output_type -> sonrio.sonr.object.MsgDeleteObjectResponse
	46, // 46: highway.v1.HighwayService.UploadBlob:output_type -> highway.v1.UploadBlobResponse
	47, // 47: highway.v1.HighwayService.DownloadBlob:output_type -> highway.v1.DownloadBlobResponse
	48, // 48: highway.v1.HighwayService.SyncBlob:output_type -> highway.v1.SyncBlobResponse
	49, // 49: highway.v1.HighwayService.DeleteBlob:output_type -> highway.v1.DeleteBlobResponse
	50, // 50: highway.v1.HighwayService.ParseDid:output_type -> highway.v1.ParseDidResponse
	51, // 51: highway.v1.HighwayService.ResolveDid:output_type -> highway.v1.ResolveDidResponse
	26, // [26:52] is the sub-list for method output_type
	0,  // [0:26] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_highway_proto_init() }
func file_v1_highway_proto_init() {
	if File_v1_highway_proto != nil {
		return
	}
	file_v1_request_proto_init()
	file_v1_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_highway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_highway_proto_goTypes,
		DependencyIndexes: file_v1_highway_proto_depIdxs,
	}.Build()
	File_v1_highway_proto = out.File
	file_v1_highway_proto_rawDesc = nil
	file_v1_highway_proto_goTypes = nil
	file_v1_highway_proto_depIdxs = nil
}