// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: chat_app/chatpb/chat.proto

package chatpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ChatRoomNo int64  `protobuf:"varint,3,opt,name=chatRoomNo,proto3" json:"chatRoomNo,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_app_chatpb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_chat_app_chatpb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_chat_app_chatpb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetChatRoomNo() int64 {
	if x != nil {
		return x.ChatRoomNo
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChatRoom  int64  `protobuf:"varint,4,opt,name=chatRoom,proto3" json:"chatRoom,omitempty"`
	IsYTLink  bool   `protobuf:"varint,5,opt,name=isYTLink,proto3" json:"isYTLink,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_app_chatpb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_app_chatpb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_app_chatpb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Message) GetChatRoom() int64 {
	if x != nil {
		return x.ChatRoom
	}
	return 0
}

func (x *Message) GetIsYTLink() bool {
	if x != nil {
		return x.IsYTLink
	}
	return false
}

type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Active       bool  `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	WantsNewRoom bool  `protobuf:"varint,3,opt,name=wantsNewRoom,proto3" json:"wantsNewRoom,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_app_chatpb_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_chat_app_chatpb_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_chat_app_chatpb_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Connect) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Connect) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Connect) GetWantsNewRoom() bool {
	if x != nil {
		return x.WantsNewRoom
	}
	return false
}

type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_app_chatpb_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_chat_app_chatpb_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_chat_app_chatpb_chat_proto_rawDescGZIP(), []int{3}
}

var File_chat_app_chatpb_chat_proto protoreflect.FileDescriptor

var file_chat_app_chatpb_chat_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x4a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x6f, 0x22, 0x89,
	0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x59, 0x54, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x59, 0x54, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x65, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x6f,
	0x6d, 0x22, 0x07, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x32, 0x6b, 0x0a, 0x09, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_app_chatpb_chat_proto_rawDescOnce sync.Once
	file_chat_app_chatpb_chat_proto_rawDescData = file_chat_app_chatpb_chat_proto_rawDesc
)

func file_chat_app_chatpb_chat_proto_rawDescGZIP() []byte {
	file_chat_app_chatpb_chat_proto_rawDescOnce.Do(func() {
		file_chat_app_chatpb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_app_chatpb_chat_proto_rawDescData)
	})
	return file_chat_app_chatpb_chat_proto_rawDescData
}

var file_chat_app_chatpb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chat_app_chatpb_chat_proto_goTypes = []interface{}{
	(*User)(nil),    // 0: chat.User
	(*Message)(nil), // 1: chat.Message
	(*Connect)(nil), // 2: chat.Connect
	(*Close)(nil),   // 3: chat.Close
}
var file_chat_app_chatpb_chat_proto_depIdxs = []int32{
	0, // 0: chat.Connect.user:type_name -> chat.User
	2, // 1: chat.Broadcast.CreateStream:input_type -> chat.Connect
	1, // 2: chat.Broadcast.BroadcastMessage:input_type -> chat.Message
	1, // 3: chat.Broadcast.CreateStream:output_type -> chat.Message
	3, // 4: chat.Broadcast.BroadcastMessage:output_type -> chat.Close
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_app_chatpb_chat_proto_init() }
func file_chat_app_chatpb_chat_proto_init() {
	if File_chat_app_chatpb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_app_chatpb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_chat_app_chatpb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_app_chatpb_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
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
		file_chat_app_chatpb_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
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
			RawDescriptor: file_chat_app_chatpb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_app_chatpb_chat_proto_goTypes,
		DependencyIndexes: file_chat_app_chatpb_chat_proto_depIdxs,
		MessageInfos:      file_chat_app_chatpb_chat_proto_msgTypes,
	}.Build()
	File_chat_app_chatpb_chat_proto = out.File
	file_chat_app_chatpb_chat_proto_rawDesc = nil
	file_chat_app_chatpb_chat_proto_goTypes = nil
	file_chat_app_chatpb_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BroadcastClient is the client API for Broadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BroadcastClient interface {
	CreateStream(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Broadcast_CreateStreamClient, error)
	BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error)
}

type broadcastClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastClient(cc grpc.ClientConnInterface) BroadcastClient {
	return &broadcastClient{cc}
}

func (c *broadcastClient) CreateStream(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Broadcast_CreateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Broadcast_serviceDesc.Streams[0], "/chat.Broadcast/CreateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &broadcastCreateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broadcast_CreateStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type broadcastCreateStreamClient struct {
	grpc.ClientStream
}

func (x *broadcastCreateStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *broadcastClient) BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error) {
	out := new(Close)
	err := c.cc.Invoke(ctx, "/chat.Broadcast/BroadcastMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServer is the server API for Broadcast service.
type BroadcastServer interface {
	CreateStream(*Connect, Broadcast_CreateStreamServer) error
	BroadcastMessage(context.Context, *Message) (*Close, error)
}

// UnimplementedBroadcastServer can be embedded to have forward compatible implementations.
type UnimplementedBroadcastServer struct {
}

func (*UnimplementedBroadcastServer) CreateStream(*Connect, Broadcast_CreateStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateStream not implemented")
}
func (*UnimplementedBroadcastServer) BroadcastMessage(context.Context, *Message) (*Close, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastMessage not implemented")
}

func RegisterBroadcastServer(s *grpc.Server, srv BroadcastServer) {
	s.RegisterService(&_Broadcast_serviceDesc, srv)
}

func _Broadcast_CreateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Connect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BroadcastServer).CreateStream(m, &broadcastCreateStreamServer{stream})
}

type Broadcast_CreateStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type broadcastCreateStreamServer struct {
	grpc.ServerStream
}

func (x *broadcastCreateStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Broadcast_BroadcastMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).BroadcastMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.Broadcast/BroadcastMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).BroadcastMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Broadcast",
	HandlerType: (*BroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastMessage",
			Handler:    _Broadcast_BroadcastMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateStream",
			Handler:       _Broadcast_CreateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat_app/chatpb/chat.proto",
}
