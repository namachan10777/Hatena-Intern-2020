// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: title_fetcher.proto

package title_fetcher

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

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_title_fetcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_title_fetcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_title_fetcher_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FetchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *FetchReply) Reset() {
	*x = FetchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_title_fetcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchReply) ProtoMessage() {}

func (x *FetchReply) ProtoReflect() protoreflect.Message {
	mi := &file_title_fetcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchReply.ProtoReflect.Descriptor instead.
func (*FetchReply) Descriptor() ([]byte, []int) {
	return file_title_fetcher_proto_rawDescGZIP(), []int{1}
}

func (x *FetchReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_title_fetcher_proto protoreflect.FileDescriptor

var file_title_fetcher_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x22, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0x4f, 0x0a, 0x0c, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x05, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x74, 0x65, 0x6e, 0x61,
	0x2f, 0x48, 0x61, 0x74, 0x65, 0x6e, 0x61, 0x2d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2d, 0x32,
	0x30, 0x32, 0x30, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_title_fetcher_proto_rawDescOnce sync.Once
	file_title_fetcher_proto_rawDescData = file_title_fetcher_proto_rawDesc
)

func file_title_fetcher_proto_rawDescGZIP() []byte {
	file_title_fetcher_proto_rawDescOnce.Do(func() {
		file_title_fetcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_title_fetcher_proto_rawDescData)
	})
	return file_title_fetcher_proto_rawDescData
}

var file_title_fetcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_title_fetcher_proto_goTypes = []interface{}{
	(*FetchRequest)(nil), // 0: title_fetcher.FetchRequest
	(*FetchReply)(nil),   // 1: title_fetcher.FetchReply
}
var file_title_fetcher_proto_depIdxs = []int32{
	0, // 0: title_fetcher.TitleFetcher.Fetch:input_type -> title_fetcher.FetchRequest
	1, // 1: title_fetcher.TitleFetcher.Fetch:output_type -> title_fetcher.FetchReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_title_fetcher_proto_init() }
func file_title_fetcher_proto_init() {
	if File_title_fetcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_title_fetcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_title_fetcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchReply); i {
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
			RawDescriptor: file_title_fetcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_title_fetcher_proto_goTypes,
		DependencyIndexes: file_title_fetcher_proto_depIdxs,
		MessageInfos:      file_title_fetcher_proto_msgTypes,
	}.Build()
	File_title_fetcher_proto = out.File
	file_title_fetcher_proto_rawDesc = nil
	file_title_fetcher_proto_goTypes = nil
	file_title_fetcher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TitleFetcherClient is the client API for TitleFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TitleFetcherClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchReply, error)
}

type titleFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewTitleFetcherClient(cc grpc.ClientConnInterface) TitleFetcherClient {
	return &titleFetcherClient{cc}
}

func (c *titleFetcherClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchReply, error) {
	out := new(FetchReply)
	err := c.cc.Invoke(ctx, "/title_fetcher.TitleFetcher/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TitleFetcherServer is the server API for TitleFetcher service.
type TitleFetcherServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchReply, error)
}

// UnimplementedTitleFetcherServer can be embedded to have forward compatible implementations.
type UnimplementedTitleFetcherServer struct {
}

func (*UnimplementedTitleFetcherServer) Fetch(context.Context, *FetchRequest) (*FetchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

func RegisterTitleFetcherServer(s *grpc.Server, srv TitleFetcherServer) {
	s.RegisterService(&_TitleFetcher_serviceDesc, srv)
}

func _TitleFetcher_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TitleFetcherServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/title_fetcher.TitleFetcher/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TitleFetcherServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TitleFetcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "title_fetcher.TitleFetcher",
	HandlerType: (*TitleFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _TitleFetcher_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "title_fetcher.proto",
}
