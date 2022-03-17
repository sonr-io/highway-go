// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: bucket/bucket.proto

package bucket

import (
	object "go.buf.build/grpc/go/sonr-io/sonr/object"
	registry "go.buf.build/grpc/go/sonr-io/sonr/registry"
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

// BucketType is the type of a bucket.
type BucketType int32

const (
	// BucketTypeUnspecified is the default value.
	BucketType_BUCKET_TYPE_UNSPECIFIED BucketType = 0
	// BucketTypeApp is an App specific bucket. For Assets regarding the service.
	BucketType_BUCKET_TYPE_APP BucketType = 1
	// BucketTypeUser is a User specific bucket. For any remote user data that is required
	// to be stored in the Network.
	BucketType_BUCKET_TYPE_USER BucketType = 2
)

// Enum value maps for BucketType.
var (
	BucketType_name = map[int32]string{
		0: "BUCKET_TYPE_UNSPECIFIED",
		1: "BUCKET_TYPE_APP",
		2: "BUCKET_TYPE_USER",
	}
	BucketType_value = map[string]int32{
		"BUCKET_TYPE_UNSPECIFIED": 0,
		"BUCKET_TYPE_APP":         1,
		"BUCKET_TYPE_USER":        2,
	}
)

func (x BucketType) Enum() *BucketType {
	p := new(BucketType)
	*p = x
	return p
}

func (x BucketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BucketType) Descriptor() protoreflect.EnumDescriptor {
	return file_bucket_bucket_proto_enumTypes[0].Descriptor()
}

func (BucketType) Type() protoreflect.EnumType {
	return &file_bucket_bucket_proto_enumTypes[0]
}

func (x BucketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BucketType.Descriptor instead.
func (BucketType) EnumDescriptor() ([]byte, []int) {
	return file_bucket_bucket_proto_rawDescGZIP(), []int{0}
}

// EventType is the type of event being performed on a Bucket.
type BucketEventType int32

const (
	// EventTypeUnspecified is the default value.
	BucketEventType_BUCKET_EVENT_TYPE_UNSPECIFIED BucketEventType = 0
	// EventTypeGet is a get event being performed on a Bucket record.
	BucketEventType_BUCKET_EVENT_TYPE_GET BucketEventType = 1
	// EventTypeSet is a set event on the Bucket store.
	BucketEventType_BUCKET_EVENT_TYPE_SET BucketEventType = 2
	// EventTypeDelete is a delete event on the Bucket store.
	BucketEventType_BUCKET_EVENT_TYPE_DELETE BucketEventType = 3
)

// Enum value maps for BucketEventType.
var (
	BucketEventType_name = map[int32]string{
		0: "BUCKET_EVENT_TYPE_UNSPECIFIED",
		1: "BUCKET_EVENT_TYPE_GET",
		2: "BUCKET_EVENT_TYPE_SET",
		3: "BUCKET_EVENT_TYPE_DELETE",
	}
	BucketEventType_value = map[string]int32{
		"BUCKET_EVENT_TYPE_UNSPECIFIED": 0,
		"BUCKET_EVENT_TYPE_GET":         1,
		"BUCKET_EVENT_TYPE_SET":         2,
		"BUCKET_EVENT_TYPE_DELETE":      3,
	}
)

func (x BucketEventType) Enum() *BucketEventType {
	p := new(BucketEventType)
	*p = x
	return p
}

func (x BucketEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BucketEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_bucket_bucket_proto_enumTypes[1].Descriptor()
}

func (BucketEventType) Type() protoreflect.EnumType {
	return &file_bucket_bucket_proto_enumTypes[1]
}

func (x BucketEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BucketEventType.Descriptor instead.
func (BucketEventType) EnumDescriptor() ([]byte, []int) {
	return file_bucket_bucket_proto_rawDescGZIP(), []int{1}
}

// Bucket is a collection of objects.
type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Label is human-readable name of the bucket.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Description is a human-readable description of the bucket.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Type is the kind of bucket for either App specific or User specific data.
	Type BucketType `protobuf:"varint,3,opt,name=type,proto3,enum=sonrio.sonr.bucket.BucketType" json:"type,omitempty"`
	// Did is the identifier of the bucket.
	Did string `protobuf:"bytes,4,opt,name=did,proto3" json:"did,omitempty"`
	// Objects are stored in a tree structure.
	Objects []*object.ObjectDoc `protobuf:"bytes,5,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_bucket_bucket_proto_rawDescGZIP(), []int{0}
}

func (x *Bucket) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Bucket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Bucket) GetType() BucketType {
	if x != nil {
		return x.Type
	}
	return BucketType_BUCKET_TYPE_UNSPECIFIED
}

func (x *Bucket) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Bucket) GetObjects() []*object.ObjectDoc {
	if x != nil {
		return x.Objects
	}
	return nil
}

// BucketEvent is the base event type for all Bucket events.
type BucketEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Owner is the peer that originated the event.
	Peer *registry.Peer `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	// Type is the type of event being performed on a Bucket.
	Type BucketEventType `protobuf:"varint,2,opt,name=type,proto3,enum=sonrio.sonr.bucket.BucketEventType" json:"type,omitempty"`
	// Object is the entry being modified in the Bucket.
	Object *object.ObjectDoc `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	// Metadata is the metadata associated with the event.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BucketEvent) Reset() {
	*x = BucketEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_bucket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketEvent) ProtoMessage() {}

func (x *BucketEvent) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_bucket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketEvent.ProtoReflect.Descriptor instead.
func (*BucketEvent) Descriptor() ([]byte, []int) {
	return file_bucket_bucket_proto_rawDescGZIP(), []int{1}
}

func (x *BucketEvent) GetPeer() *registry.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *BucketEvent) GetType() BucketEventType {
	if x != nil {
		return x.Type
	}
	return BucketEventType_BUCKET_EVENT_TYPE_UNSPECIFIED
}

func (x *BucketEvent) GetObject() *object.ObjectDoc {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *BucketEvent) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_bucket_bucket_proto protoreflect.FileDescriptor

var file_bucket_bucket_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x13, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x07,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x07, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x0b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x54, 0x0a,
	0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x42,
	0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x55, 0x43, 0x4b,
	0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x10, 0x02, 0x2a, 0x88, 0x01, 0x0a, 0x0f, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x42, 0x55, 0x43, 0x4b, 0x45,
	0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x55,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x02,
	0x12, 0x1c, 0x0a, 0x18, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6f, 0x6e, 0x72, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bucket_bucket_proto_rawDescOnce sync.Once
	file_bucket_bucket_proto_rawDescData = file_bucket_bucket_proto_rawDesc
)

func file_bucket_bucket_proto_rawDescGZIP() []byte {
	file_bucket_bucket_proto_rawDescOnce.Do(func() {
		file_bucket_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_bucket_bucket_proto_rawDescData)
	})
	return file_bucket_bucket_proto_rawDescData
}

var file_bucket_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bucket_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bucket_bucket_proto_goTypes = []interface{}{
	(BucketType)(0),          // 0: sonrio.sonr.bucket.BucketType
	(BucketEventType)(0),     // 1: sonrio.sonr.bucket.BucketEventType
	(*Bucket)(nil),           // 2: sonrio.sonr.bucket.Bucket
	(*BucketEvent)(nil),      // 3: sonrio.sonr.bucket.BucketEvent
	nil,                      // 4: sonrio.sonr.bucket.BucketEvent.MetadataEntry
	(*object.ObjectDoc)(nil), // 5: sonrio.sonr.object.ObjectDoc
	(*registry.Peer)(nil),    // 6: sonrio.sonr.registry.Peer
}
var file_bucket_bucket_proto_depIdxs = []int32{
	0, // 0: sonrio.sonr.bucket.Bucket.type:type_name -> sonrio.sonr.bucket.BucketType
	5, // 1: sonrio.sonr.bucket.Bucket.objects:type_name -> sonrio.sonr.object.ObjectDoc
	6, // 2: sonrio.sonr.bucket.BucketEvent.peer:type_name -> sonrio.sonr.registry.Peer
	1, // 3: sonrio.sonr.bucket.BucketEvent.type:type_name -> sonrio.sonr.bucket.BucketEventType
	5, // 4: sonrio.sonr.bucket.BucketEvent.object:type_name -> sonrio.sonr.object.ObjectDoc
	4, // 5: sonrio.sonr.bucket.BucketEvent.metadata:type_name -> sonrio.sonr.bucket.BucketEvent.MetadataEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bucket_bucket_proto_init() }
func file_bucket_bucket_proto_init() {
	if File_bucket_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bucket_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucket); i {
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
		file_bucket_bucket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketEvent); i {
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
			RawDescriptor: file_bucket_bucket_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bucket_bucket_proto_goTypes,
		DependencyIndexes: file_bucket_bucket_proto_depIdxs,
		EnumInfos:         file_bucket_bucket_proto_enumTypes,
		MessageInfos:      file_bucket_bucket_proto_msgTypes,
	}.Build()
	File_bucket_bucket_proto = out.File
	file_bucket_bucket_proto_rawDesc = nil
	file_bucket_bucket_proto_goTypes = nil
	file_bucket_bucket_proto_depIdxs = nil
}
