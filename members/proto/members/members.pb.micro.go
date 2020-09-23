// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/members/members.proto

package members

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Members service

func NewMembersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Members service

type MembersService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Members_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Members_PingPongService, error)
}

type membersService struct {
	c    client.Client
	name string
}

func NewMembersService(name string, c client.Client) MembersService {
	return &membersService{
		c:    c,
		name: name,
	}
}

func (c *membersService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Members.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membersService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Members_StreamService, error) {
	req := c.c.NewRequest(c.name, "Members.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &membersServiceStream{stream}, nil
}

type Members_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type membersServiceStream struct {
	stream client.Stream
}

func (x *membersServiceStream) Close() error {
	return x.stream.Close()
}

func (x *membersServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *membersServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *membersServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *membersServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *membersService) PingPong(ctx context.Context, opts ...client.CallOption) (Members_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Members.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &membersServicePingPong{stream}, nil
}

type Members_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type membersServicePingPong struct {
	stream client.Stream
}

func (x *membersServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *membersServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *membersServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *membersServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *membersServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *membersServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Members service

type MembersHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Members_StreamStream) error
	PingPong(context.Context, Members_PingPongStream) error
}

func RegisterMembersHandler(s server.Server, hdlr MembersHandler, opts ...server.HandlerOption) error {
	type members interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Members struct {
		members
	}
	h := &membersHandler{hdlr}
	return s.Handle(s.NewHandler(&Members{h}, opts...))
}

type membersHandler struct {
	MembersHandler
}

func (h *membersHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.MembersHandler.Call(ctx, in, out)
}

func (h *membersHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MembersHandler.Stream(ctx, m, &membersStreamStream{stream})
}

type Members_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type membersStreamStream struct {
	stream server.Stream
}

func (x *membersStreamStream) Close() error {
	return x.stream.Close()
}

func (x *membersStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *membersStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *membersStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *membersStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *membersHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.MembersHandler.PingPong(ctx, &membersPingPongStream{stream})
}

type Members_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type membersPingPongStream struct {
	stream server.Stream
}

func (x *membersPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *membersPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *membersPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *membersPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *membersPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *membersPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
