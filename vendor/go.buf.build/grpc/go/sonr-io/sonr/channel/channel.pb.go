// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: channel/channel.proto

package channel

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

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Label is human-readable name of the channel.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Description is a human-readable description of the channel.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Did is the identifier of the channel.
	Did string `protobuf:"bytes,4,opt,name=did,proto3" json:"did,omitempty"`
	// RegisterdObject is the object that is registered as the payload for the channel.
	RegisteredObject *object.ObjectDoc `protobuf:"bytes,5,opt,name=registered_object,json=registeredObject,proto3" json:"registered_object,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_channel_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_channel_channel_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Channel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Channel) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Channel) GetRegisteredObject() *object.ObjectDoc {
	if x != nil {
		return x.RegisteredObject
	}
	return nil
}

// ChannelMessage is a message sent to a channel.
type ChannelMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Owner is the peer that originated the message.
	Peer *registry.Peer `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	// Did is the identifier of the channel.
	Did string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	// Data is the data being sent.
	Object *object.ObjectDoc `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"` // optional
	// Metadata is the metadata associated with the message.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChannelMessage) Reset() {
	*x = ChannelMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessage) ProtoMessage() {}

func (x *ChannelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_channel_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessage.ProtoReflect.Descriptor instead.
func (*ChannelMessage) Descriptor() ([]byte, []int) {
	return file_channel_channel_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelMessage) GetPeer() *registry.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *ChannelMessage) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *ChannelMessage) GetObject() *object.ObjectDoc {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *ChannelMessage) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_channel_channel_proto protoreflect.FileDescriptor

var file_channel_channel_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x13, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x65, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x11,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x95, 0x02, 0x0a, 0x0e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x70,
	0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x35, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x4d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_channel_channel_proto_rawDescOnce sync.Once
	file_channel_channel_proto_rawDescData = file_channel_channel_proto_rawDesc
)

func file_channel_channel_proto_rawDescGZIP() []byte {
	file_channel_channel_proto_rawDescOnce.Do(func() {
		file_channel_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_channel_proto_rawDescData)
	})
	return file_channel_channel_proto_rawDescData
}

var file_channel_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_channel_channel_proto_goTypes = []interface{}{
	(*Channel)(nil),          // 0: sonrio.sonr.channel.Channel
	(*ChannelMessage)(nil),   // 1: sonrio.sonr.channel.ChannelMessage
	nil,                      // 2: sonrio.sonr.channel.ChannelMessage.MetadataEntry
	(*object.ObjectDoc)(nil), // 3: sonrio.sonr.object.ObjectDoc
	(*registry.Peer)(nil),    // 4: sonrio.sonr.registry.Peer
}
var file_channel_channel_proto_depIdxs = []int32{
	3, // 0: sonrio.sonr.channel.Channel.registered_object:type_name -> sonrio.sonr.object.ObjectDoc
	4, // 1: sonrio.sonr.channel.ChannelMessage.peer:type_name -> sonrio.sonr.registry.Peer
	3, // 2: sonrio.sonr.channel.ChannelMessage.object:type_name -> sonrio.sonr.object.ObjectDoc
	2, // 3: sonrio.sonr.channel.ChannelMessage.metadata:type_name -> sonrio.sonr.channel.ChannelMessage.MetadataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_channel_channel_proto_init() }
func file_channel_channel_proto_init() {
	if File_channel_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_channel_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelMessage); i {
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
			RawDescriptor: file_channel_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_channel_channel_proto_goTypes,
		DependencyIndexes: file_channel_channel_proto_depIdxs,
		MessageInfos:      file_channel_channel_proto_msgTypes,
	}.Build()
	File_channel_channel_proto = out.File
	file_channel_channel_proto_rawDesc = nil
	file_channel_channel_proto_goTypes = nil
	file_channel_channel_proto_depIdxs = nil
}
