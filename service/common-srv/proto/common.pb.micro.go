// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: common.proto

package fotune_srv_common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Api Endpoints for Common service

func NewCommonEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Common service

type CommonService interface {
	Carousel(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*CarouselList, error)
	CustomerService(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ContractAddr, error)
	GetAppVersion(ctx context.Context, in *VersionReq, opts ...client.CallOption) (*AppVersion, error)
	GetUserRateRank(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*UserRateRankResp, error)
	GetUserRateYearRank(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*UserRateRankResp, error)
	PushProfitImageOss(ctx context.Context, in *PushImageReq, opts ...client.CallOption) (*ImageResp, error)
}

type commonService struct {
	c    client.Client
	name string
}

func NewCommonService(name string, c client.Client) CommonService {
	return &commonService{
		c:    c,
		name: name,
	}
}

func (c *commonService) Carousel(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*CarouselList, error) {
	req := c.c.NewRequest(c.name, "Common.Carousel", in)
	out := new(CarouselList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonService) CustomerService(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ContractAddr, error) {
	req := c.c.NewRequest(c.name, "Common.CustomerService", in)
	out := new(ContractAddr)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonService) GetAppVersion(ctx context.Context, in *VersionReq, opts ...client.CallOption) (*AppVersion, error) {
	req := c.c.NewRequest(c.name, "Common.GetAppVersion", in)
	out := new(AppVersion)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonService) GetUserRateRank(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*UserRateRankResp, error) {
	req := c.c.NewRequest(c.name, "Common.GetUserRateRank", in)
	out := new(UserRateRankResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonService) GetUserRateYearRank(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*UserRateRankResp, error) {
	req := c.c.NewRequest(c.name, "Common.GetUserRateYearRank", in)
	out := new(UserRateRankResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonService) PushProfitImageOss(ctx context.Context, in *PushImageReq, opts ...client.CallOption) (*ImageResp, error) {
	req := c.c.NewRequest(c.name, "Common.PushProfitImageOss", in)
	out := new(ImageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Common service

type CommonHandler interface {
	Carousel(context.Context, *empty.Empty, *CarouselList) error
	CustomerService(context.Context, *empty.Empty, *ContractAddr) error
	GetAppVersion(context.Context, *VersionReq, *AppVersion) error
	GetUserRateRank(context.Context, *empty.Empty, *UserRateRankResp) error
	GetUserRateYearRank(context.Context, *empty.Empty, *UserRateRankResp) error
	PushProfitImageOss(context.Context, *PushImageReq, *ImageResp) error
}

func RegisterCommonHandler(s server.Server, hdlr CommonHandler, opts ...server.HandlerOption) error {
	type common interface {
		Carousel(ctx context.Context, in *empty.Empty, out *CarouselList) error
		CustomerService(ctx context.Context, in *empty.Empty, out *ContractAddr) error
		GetAppVersion(ctx context.Context, in *VersionReq, out *AppVersion) error
		GetUserRateRank(ctx context.Context, in *empty.Empty, out *UserRateRankResp) error
		GetUserRateYearRank(ctx context.Context, in *empty.Empty, out *UserRateRankResp) error
		PushProfitImageOss(ctx context.Context, in *PushImageReq, out *ImageResp) error
	}
	type Common struct {
		common
	}
	h := &commonHandler{hdlr}
	return s.Handle(s.NewHandler(&Common{h}, opts...))
}

type commonHandler struct {
	CommonHandler
}

func (h *commonHandler) Carousel(ctx context.Context, in *empty.Empty, out *CarouselList) error {
	return h.CommonHandler.Carousel(ctx, in, out)
}

func (h *commonHandler) CustomerService(ctx context.Context, in *empty.Empty, out *ContractAddr) error {
	return h.CommonHandler.CustomerService(ctx, in, out)
}

func (h *commonHandler) GetAppVersion(ctx context.Context, in *VersionReq, out *AppVersion) error {
	return h.CommonHandler.GetAppVersion(ctx, in, out)
}

func (h *commonHandler) GetUserRateRank(ctx context.Context, in *empty.Empty, out *UserRateRankResp) error {
	return h.CommonHandler.GetUserRateRank(ctx, in, out)
}

func (h *commonHandler) GetUserRateYearRank(ctx context.Context, in *empty.Empty, out *UserRateRankResp) error {
	return h.CommonHandler.GetUserRateYearRank(ctx, in, out)
}

func (h *commonHandler) PushProfitImageOss(ctx context.Context, in *PushImageReq, out *ImageResp) error {
	return h.CommonHandler.PushProfitImageOss(ctx, in, out)
}
