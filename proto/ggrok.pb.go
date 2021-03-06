// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ggrok.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ggrok.proto

It has these top-level messages:
	ClientServer
	InitRequest
	HTTPResponse
	ServerClient
	InitResponse
	HTTPRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// container message for Client -> Server communication. clients send InitRequest messages to initiate a steam with Server
// clients send HTTPResponse messages encapsulating the the proxied response into a proto message.
type ClientServer struct {
	// Types that are valid to be assigned to Clientmsg:
	//	*ClientServer_Initreq
	//	*ClientServer_Httpresp
	Clientmsg isClientServer_Clientmsg `protobuf_oneof:"clientmsg"`
}

func (m *ClientServer) Reset()                    { *m = ClientServer{} }
func (m *ClientServer) String() string            { return proto1.CompactTextString(m) }
func (*ClientServer) ProtoMessage()               {}
func (*ClientServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isClientServer_Clientmsg interface {
	isClientServer_Clientmsg()
}

type ClientServer_Initreq struct {
	Initreq *InitRequest `protobuf:"bytes,1,opt,name=initreq,oneof"`
}
type ClientServer_Httpresp struct {
	Httpresp *HTTPResponse `protobuf:"bytes,2,opt,name=httpresp,oneof"`
}

func (*ClientServer_Initreq) isClientServer_Clientmsg()  {}
func (*ClientServer_Httpresp) isClientServer_Clientmsg() {}

func (m *ClientServer) GetClientmsg() isClientServer_Clientmsg {
	if m != nil {
		return m.Clientmsg
	}
	return nil
}

func (m *ClientServer) GetInitreq() *InitRequest {
	if x, ok := m.GetClientmsg().(*ClientServer_Initreq); ok {
		return x.Initreq
	}
	return nil
}

func (m *ClientServer) GetHttpresp() *HTTPResponse {
	if x, ok := m.GetClientmsg().(*ClientServer_Httpresp); ok {
		return x.Httpresp
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientServer) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _ClientServer_OneofMarshaler, _ClientServer_OneofUnmarshaler, _ClientServer_OneofSizer, []interface{}{
		(*ClientServer_Initreq)(nil),
		(*ClientServer_Httpresp)(nil),
	}
}

func _ClientServer_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*ClientServer)
	// clientmsg
	switch x := m.Clientmsg.(type) {
	case *ClientServer_Initreq:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Initreq); err != nil {
			return err
		}
	case *ClientServer_Httpresp:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Httpresp); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientServer.Clientmsg has unexpected type %T", x)
	}
	return nil
}

func _ClientServer_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*ClientServer)
	switch tag {
	case 1: // clientmsg.initreq
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(InitRequest)
		err := b.DecodeMessage(msg)
		m.Clientmsg = &ClientServer_Initreq{msg}
		return true, err
	case 2: // clientmsg.httpresp
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(HTTPResponse)
		err := b.DecodeMessage(msg)
		m.Clientmsg = &ClientServer_Httpresp{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientServer_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*ClientServer)
	// clientmsg
	switch x := m.Clientmsg.(type) {
	case *ClientServer_Initreq:
		s := proto1.Size(x.Initreq)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *ClientServer_Httpresp:
		s := proto1.Size(x.Httpresp)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// InitRequest is the initial client request sent to the server.
type InitRequest struct {
	// opportunistically ask for a hostname. Server may still decide to respond
	// with a random ID for hostname.
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto1.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InitRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// HTTPResponse holds the bytes returned from a call to http.Response.Write(). See: https://golang.org/src/net/http/response.go?s=7618:7661#L229
type HTTPResponse struct {
	// a unique ID representing this Request -> Response http transaction
	RequestUUID string `protobuf:"bytes,1,opt,name=requestUUID" json:"requestUUID,omitempty"`
	// a byte array consiting of the results from http.Response.Write method
	Response []byte `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (m *HTTPResponse) Reset()                    { *m = HTTPResponse{} }
func (m *HTTPResponse) String() string            { return proto1.CompactTextString(m) }
func (*HTTPResponse) ProtoMessage()               {}
func (*HTTPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HTTPResponse) GetRequestUUID() string {
	if m != nil {
		return m.RequestUUID
	}
	return ""
}

func (m *HTTPResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

// Container message for Server -> Client communication. Servers respond to client's InitRequest messages.
// Servers encapulate http.Request and send them to client
type ServerClient struct {
	// Types that are valid to be assigned to Servermsg:
	//	*ServerClient_Initresp
	//	*ServerClient_Httpreq
	Servermsg isServerClient_Servermsg `protobuf_oneof:"servermsg"`
}

func (m *ServerClient) Reset()                    { *m = ServerClient{} }
func (m *ServerClient) String() string            { return proto1.CompactTextString(m) }
func (*ServerClient) ProtoMessage()               {}
func (*ServerClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isServerClient_Servermsg interface {
	isServerClient_Servermsg()
}

type ServerClient_Initresp struct {
	Initresp *InitResponse `protobuf:"bytes,1,opt,name=initresp,oneof"`
}
type ServerClient_Httpreq struct {
	Httpreq *HTTPRequest `protobuf:"bytes,2,opt,name=httpreq,oneof"`
}

func (*ServerClient_Initresp) isServerClient_Servermsg() {}
func (*ServerClient_Httpreq) isServerClient_Servermsg()  {}

func (m *ServerClient) GetServermsg() isServerClient_Servermsg {
	if m != nil {
		return m.Servermsg
	}
	return nil
}

func (m *ServerClient) GetInitresp() *InitResponse {
	if x, ok := m.GetServermsg().(*ServerClient_Initresp); ok {
		return x.Initresp
	}
	return nil
}

func (m *ServerClient) GetHttpreq() *HTTPRequest {
	if x, ok := m.GetServermsg().(*ServerClient_Httpreq); ok {
		return x.Httpreq
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerClient) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _ServerClient_OneofMarshaler, _ServerClient_OneofUnmarshaler, _ServerClient_OneofSizer, []interface{}{
		(*ServerClient_Initresp)(nil),
		(*ServerClient_Httpreq)(nil),
	}
}

func _ServerClient_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*ServerClient)
	// servermsg
	switch x := m.Servermsg.(type) {
	case *ServerClient_Initresp:
		b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Initresp); err != nil {
			return err
		}
	case *ServerClient_Httpreq:
		b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Httpreq); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerClient.Servermsg has unexpected type %T", x)
	}
	return nil
}

func _ServerClient_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*ServerClient)
	switch tag {
	case 1: // servermsg.initresp
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(InitResponse)
		err := b.DecodeMessage(msg)
		m.Servermsg = &ServerClient_Initresp{msg}
		return true, err
	case 2: // servermsg.httpreq
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(HTTPRequest)
		err := b.DecodeMessage(msg)
		m.Servermsg = &ServerClient_Httpreq{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerClient_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*ServerClient)
	// servermsg
	switch x := m.Servermsg.(type) {
	case *ServerClient_Initresp:
		s := proto1.Size(x.Initresp)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *ServerClient_Httpreq:
		s := proto1.Size(x.Httpreq)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// InitResponse is the response to a client's initial request.
type InitResponse struct {
	// server defined hostname for the current stream
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *InitResponse) Reset()                    { *m = InitResponse{} }
func (m *InitResponse) String() string            { return proto1.CompactTextString(m) }
func (*InitResponse) ProtoMessage()               {}
func (*InitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InitResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// HTTPRequest holds the bytes returned from a call to http.Request.Write(). See: https://golang.org/src/net/http/request.go?s=16549:16591#L467
type HTTPRequest struct {
	// a unique ID representing this Request -> Response http transaction
	RequestUUID string `protobuf:"bytes,1,opt,name=requestUUID" json:"requestUUID,omitempty"`
	// a byte array consisting of the results from http.Request.Write method
	Request []byte `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (m *HTTPRequest) Reset()                    { *m = HTTPRequest{} }
func (m *HTTPRequest) String() string            { return proto1.CompactTextString(m) }
func (*HTTPRequest) ProtoMessage()               {}
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HTTPRequest) GetRequestUUID() string {
	if m != nil {
		return m.RequestUUID
	}
	return ""
}

func (m *HTTPRequest) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto1.RegisterType((*ClientServer)(nil), "proto.ClientServer")
	proto1.RegisterType((*InitRequest)(nil), "proto.InitRequest")
	proto1.RegisterType((*HTTPResponse)(nil), "proto.HTTPResponse")
	proto1.RegisterType((*ServerClient)(nil), "proto.ServerClient")
	proto1.RegisterType((*InitResponse)(nil), "proto.InitResponse")
	proto1.RegisterType((*HTTPRequest)(nil), "proto.HTTPRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FFProxy service

type FFProxyClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (FFProxy_StreamClient, error)
}

type fFProxyClient struct {
	cc *grpc.ClientConn
}

func NewFFProxyClient(cc *grpc.ClientConn) FFProxyClient {
	return &fFProxyClient{cc}
}

func (c *fFProxyClient) Stream(ctx context.Context, opts ...grpc.CallOption) (FFProxy_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FFProxy_serviceDesc.Streams[0], c.cc, "/proto.FFProxy/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &fFProxyStreamClient{stream}
	return x, nil
}

type FFProxy_StreamClient interface {
	Send(*ClientServer) error
	Recv() (*ServerClient, error)
	grpc.ClientStream
}

type fFProxyStreamClient struct {
	grpc.ClientStream
}

func (x *fFProxyStreamClient) Send(m *ClientServer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fFProxyStreamClient) Recv() (*ServerClient, error) {
	m := new(ServerClient)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FFProxy service

type FFProxyServer interface {
	Stream(FFProxy_StreamServer) error
}

func RegisterFFProxyServer(s *grpc.Server, srv FFProxyServer) {
	s.RegisterService(&_FFProxy_serviceDesc, srv)
}

func _FFProxy_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FFProxyServer).Stream(&fFProxyStreamServer{stream})
}

type FFProxy_StreamServer interface {
	Send(*ServerClient) error
	Recv() (*ClientServer, error)
	grpc.ServerStream
}

type fFProxyStreamServer struct {
	grpc.ServerStream
}

func (x *fFProxyStreamServer) Send(m *ServerClient) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fFProxyStreamServer) Recv() (*ClientServer, error) {
	m := new(ClientServer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FFProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FFProxy",
	HandlerType: (*FFProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _FFProxy_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ggrok.proto",
}

func init() { proto1.RegisterFile("ggrok.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x1b, 0x24, 0x9a, 0xf6, 0x9c, 0xc9, 0x5d, 0xa2, 0x4e, 0x55, 0xa6, 0xc2, 0x10, 0x41,
	0x59, 0x98, 0x29, 0xaa, 0x12, 0x89, 0xa1, 0x72, 0xdb, 0x0f, 0x50, 0xd0, 0x29, 0x8d, 0x20, 0x71,
	0x62, 0x1b, 0x04, 0x0b, 0x9f, 0x1d, 0xf9, 0x5f, 0x70, 0x17, 0xc4, 0x64, 0xbd, 0xf3, 0xbb, 0x77,
	0xfe, 0x9d, 0x0c, 0xa4, 0xaa, 0x04, 0x7f, 0xcd, 0x3b, 0xc1, 0x15, 0xa7, 0x97, 0xe6, 0xc8, 0xbe,
	0x21, 0x59, 0xbf, 0xd5, 0xd8, 0xaa, 0x1d, 0x8a, 0x0f, 0x14, 0x34, 0x87, 0xb8, 0x6e, 0x6b, 0x25,
	0xb0, 0x4f, 0xa3, 0x45, 0xb4, 0x24, 0x2b, 0x6a, 0xfd, 0x79, 0xd9, 0xd6, 0x8a, 0x61, 0xff, 0x8e,
	0x52, 0x15, 0x23, 0xe6, 0x4d, 0xf4, 0x16, 0x26, 0x27, 0xa5, 0x3a, 0x81, 0xb2, 0x4b, 0x2f, 0x4c,
	0xc3, 0xcc, 0x35, 0x14, 0xfb, 0xfd, 0x96, 0xa1, 0xec, 0x78, 0x2b, 0xb1, 0x18, 0xb1, 0xc1, 0xf6,
	0x40, 0x60, 0xfa, 0x62, 0x46, 0x36, 0xb2, 0xca, 0xae, 0x80, 0x04, 0xc9, 0x74, 0x0e, 0x93, 0x13,
	0x97, 0xaa, 0x3d, 0x36, 0x68, 0xe6, 0x4f, 0xd9, 0xa0, 0xb3, 0x27, 0x48, 0xc2, 0x4c, 0xba, 0x00,
	0x22, 0x6c, 0xdb, 0xe1, 0x50, 0x3e, 0x3a, 0x7b, 0x58, 0xd2, 0x69, 0xc2, 0xb9, 0xcd, 0xe3, 0x12,
	0x36, 0x68, 0x0d, 0x6e, 0x91, 0x2d, 0xbe, 0x06, 0xb1, 0x4c, 0xb2, 0x73, 0xe4, 0xb3, 0x33, 0xf2,
	0x5f, 0x10, 0x6f, 0xd3, 0xbb, 0xb2, 0x50, 0xbd, 0x43, 0xa7, 0x67, 0xe8, 0xc3, 0xae, 0x9c, 0x49,
	0x83, 0x4b, 0x33, 0x52, 0x83, 0x5f, 0x43, 0x12, 0x06, 0xff, 0x49, 0x5e, 0x02, 0x09, 0x22, 0xff,
	0x01, 0x9e, 0x42, 0xec, 0xa4, 0xe3, 0xf6, 0x72, 0xb5, 0x86, 0x78, 0xb3, 0xd9, 0x0a, 0xfe, 0xf9,
	0x45, 0xef, 0x61, 0xbc, 0x53, 0x02, 0x8f, 0x0d, 0xf5, 0xa4, 0xe1, 0x4f, 0x98, 0xfb, 0x62, 0xb8,
	0xa5, 0x6c, 0xb4, 0x8c, 0x6e, 0xa2, 0xe7, 0xb1, 0xb9, 0xb9, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x15, 0xfe, 0x8d, 0xf5, 0x51, 0x02, 0x00, 0x00,
}
