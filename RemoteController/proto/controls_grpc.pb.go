// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package controls

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoombaControlerClient is the client API for RoombaControler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoombaControlerClient interface {
	RoombaControl(ctx context.Context, opts ...grpc.CallOption) (RoombaControler_RoombaControlClient, error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResponse, error)
}

type roombaControlerClient struct {
	cc grpc.ClientConnInterface
}

func NewRoombaControlerClient(cc grpc.ClientConnInterface) RoombaControlerClient {
	return &roombaControlerClient{cc}
}

func (c *roombaControlerClient) RoombaControl(ctx context.Context, opts ...grpc.CallOption) (RoombaControler_RoombaControlClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoombaControler_serviceDesc.Streams[0], "/RoombaControler/RoombaControl", opts...)
	if err != nil {
		return nil, err
	}
	x := &roombaControlerRoombaControlClient{stream}
	return x, nil
}

type RoombaControler_RoombaControlClient interface {
	Send(*RoombaControlRequest) error
	Recv() (*RoombaControlResponse, error)
	grpc.ClientStream
}

type roombaControlerRoombaControlClient struct {
	grpc.ClientStream
}

func (x *roombaControlerRoombaControlClient) Send(m *RoombaControlRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *roombaControlerRoombaControlClient) Recv() (*RoombaControlResponse, error) {
	m := new(RoombaControlResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *roombaControlerClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResponse, error) {
	out := new(AckResponse)
	err := c.cc.Invoke(ctx, "/RoombaControler/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoombaControlerServer is the server API for RoombaControler service.
// All implementations must embed UnimplementedRoombaControlerServer
// for forward compatibility
type RoombaControlerServer interface {
	RoombaControl(RoombaControler_RoombaControlServer) error
	Ack(context.Context, *AckRequest) (*AckResponse, error)
	mustEmbedUnimplementedRoombaControlerServer()
}

// UnimplementedRoombaControlerServer must be embedded to have forward compatible implementations.
type UnimplementedRoombaControlerServer struct {
}

func (*UnimplementedRoombaControlerServer) RoombaControl(RoombaControler_RoombaControlServer) error {
	return status.Errorf(codes.Unimplemented, "method RoombaControl not implemented")
}
func (*UnimplementedRoombaControlerServer) Ack(context.Context, *AckRequest) (*AckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (*UnimplementedRoombaControlerServer) mustEmbedUnimplementedRoombaControlerServer() {}

func RegisterRoombaControlerServer(s *grpc.Server, srv RoombaControlerServer) {
	s.RegisterService(&_RoombaControler_serviceDesc, srv)
}

func _RoombaControler_RoombaControl_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RoombaControlerServer).RoombaControl(&roombaControlerRoombaControlServer{stream})
}

type RoombaControler_RoombaControlServer interface {
	Send(*RoombaControlResponse) error
	Recv() (*RoombaControlRequest, error)
	grpc.ServerStream
}

type roombaControlerRoombaControlServer struct {
	grpc.ServerStream
}

func (x *roombaControlerRoombaControlServer) Send(m *RoombaControlResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *roombaControlerRoombaControlServer) Recv() (*RoombaControlRequest, error) {
	m := new(RoombaControlRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RoombaControler_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoombaControlerServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoombaControler/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoombaControlerServer).Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoombaControler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoombaControler",
	HandlerType: (*RoombaControlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ack",
			Handler:    _RoombaControler_Ack_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RoombaControl",
			Handler:       _RoombaControler_RoombaControl_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "controls.proto",
}
