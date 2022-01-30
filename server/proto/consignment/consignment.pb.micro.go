// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: consignment.proto

package consignment

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

// Api Endpoints for ShippingService service

func NewShippingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ShippingService service

type ShippingService interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*CreateResponse, error)
	GetConsignment(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type shippingService struct {
	c    client.Client
	name string
}

func NewShippingService(name string, c client.Client) ShippingService {
	return &shippingService{
		c:    c,
		name: name,
	}
}

func (c *shippingService) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "ShippingService.CreateConsignment", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) GetConsignment(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "ShippingService.GetConsignment", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *CreateResponse) error
	GetConsignment(context.Context, *GetRequest, *GetResponse) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) error {
	type shippingService interface {
		CreateConsignment(ctx context.Context, in *Consignment, out *CreateResponse) error
		GetConsignment(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type ShippingService struct {
		shippingService
	}
	h := &shippingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ShippingService{h}, opts...))
}

type shippingServiceHandler struct {
	ShippingServiceHandler
}

func (h *shippingServiceHandler) CreateConsignment(ctx context.Context, in *Consignment, out *CreateResponse) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *shippingServiceHandler) GetConsignment(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ShippingServiceHandler.GetConsignment(ctx, in, out)
}
