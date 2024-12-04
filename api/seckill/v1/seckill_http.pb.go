// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.29.0
// source: seckill/v1/seckill.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserXCreateUser = "/api.shortUrlX.v1.UserX/CreateUser"
const OperationUserXGetUser = "/api.shortUrlX.v1.UserX/GetUser"
const OperationUserXGetUserByName = "/api.shortUrlX.v1.UserX/GetUserByName"

type UserXHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	GetUserByName(context.Context, *GetUserByNameRequest) (*GetUserByNameReply, error)
}

func RegisterUserXHTTPServer(s *http.Server, srv UserXHTTPServer) {
	r := s.Route("/")
	r.POST("/create_user", _UserX_CreateUser0_HTTP_Handler(srv))
	r.GET("/get_user_info", _UserX_GetUser0_HTTP_Handler(srv))
	r.GET("/get_user_info_byname", _UserX_GetUserByName0_HTTP_Handler(srv))
}

func _UserX_CreateUser0_HTTP_Handler(srv UserXHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserXCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _UserX_GetUser0_HTTP_Handler(srv UserXHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserXGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _UserX_GetUserByName0_HTTP_Handler(srv UserXHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserByNameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserXGetUserByName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserByName(ctx, req.(*GetUserByNameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserByNameReply)
		return ctx.Result(200, reply)
	}
}

type UserXHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	GetUserByName(ctx context.Context, req *GetUserByNameRequest, opts ...http.CallOption) (rsp *GetUserByNameReply, err error)
}

type UserXHTTPClientImpl struct {
	cc *http.Client
}

func NewUserXHTTPClient(client *http.Client) UserXHTTPClient {
	return &UserXHTTPClientImpl{client}
}

func (c *UserXHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/create_user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserXCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserXHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/get_user_info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserXGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserXHTTPClientImpl) GetUserByName(ctx context.Context, in *GetUserByNameRequest, opts ...http.CallOption) (*GetUserByNameReply, error) {
	var out GetUserByNameReply
	pattern := "/get_user_info_byname"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserXGetUserByName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
