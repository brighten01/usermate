// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.20.3
// source: api/usermate/v1/usermate.proto

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

const OperationUserMateAddLevel = "/usermate.v1.UserMate/AddLevel"
const OperationUserMateAddServiceCategory = "/usermate.v1.UserMate/AddServiceCategory"
const OperationUserMateAddUserMate = "/usermate.v1.UserMate/AddUserMate"
const OperationUserMateCreateOrder = "/usermate.v1.UserMate/CreateOrder"
const OperationUserMateDeleteUserMate = "/usermate.v1.UserMate/DeleteUserMate"
const OperationUserMateListUserMate = "/usermate.v1.UserMate/ListUserMate"
const OperationUserMateOrderDetail = "/usermate.v1.UserMate/OrderDetail"
const OperationUserMateOrderList = "/usermate.v1.UserMate/OrderList"
const OperationUserMateSearchUserMate = "/usermate.v1.UserMate/SearchUserMate"
const OperationUserMateUpdateOrder = "/usermate.v1.UserMate/UpdateOrder"
const OperationUserMateUpdateUserMate = "/usermate.v1.UserMate/UpdateUserMate"
const OperationUserMateUserMateDetail = "/usermate.v1.UserMate/UserMateDetail"

type UserMateHTTPServer interface {
	AddLevel(context.Context, *LevelRequest) (*LevelResponse, error)
	AddServiceCategory(context.Context, *ServiceCategoryRequest) (*ServiceCategoryResponse, error)
	// AddUserMate user mate add
	AddUserMate(context.Context, *UserMateRequest) (*UserMateReply, error)
	// CreateOrdercreate order
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
	// DeleteUserMateuser mate delete
	DeleteUserMate(context.Context, *DeleteMateRequest) (*DeleteMateReply, error)
	// ListUserMateusermate list
	ListUserMate(context.Context, *ListMateRequest) (*ListMateResponse, error)
	// OrderDetailorder detail
	OrderDetail(context.Context, *OrderDetailRequest) (*OrderDetailResponse, error)
	// OrderListorder list
	OrderList(context.Context, *OrderListRequest) (*OrderListResponse, error)
	SearchUserMate(context.Context, *SearchUserMateRequest) (*SearchUserMateResponse, error)
	// UpdateOrderupdate order
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderReply, error)
	// UpdateUserMateuser mate update data
	UpdateUserMate(context.Context, *UserMateUpdateRequest) (*UserMateUpdateReply, error)
	// UserMateDetailuser mates detail show
	UserMateDetail(context.Context, *UserMateShowRequest) (*UserMateShowReply, error)
}

func RegisterUserMateHTTPServer(s *http.Server, srv UserMateHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/usermate/add", _UserMate_AddUserMate0_HTTP_Handler(srv))
	r.GET("/api/v1/mate/delete/{id}", _UserMate_DeleteUserMate0_HTTP_Handler(srv))
	r.GET("/api/v1/mate/detail/{id}", _UserMate_UserMateDetail0_HTTP_Handler(srv))
	r.POST("/api/v1/usermate/update", _UserMate_UpdateUserMate0_HTTP_Handler(srv))
	r.GET("/api/v1/usermate/list", _UserMate_ListUserMate0_HTTP_Handler(srv))
	r.GET("/api/v1/usermate/search/{name}", _UserMate_SearchUserMate0_HTTP_Handler(srv))
	r.POST("/api/v1/order/create", _UserMate_CreateOrder0_HTTP_Handler(srv))
	r.POST("/api/v1/order/update", _UserMate_UpdateOrder0_HTTP_Handler(srv))
	r.POST("/api/v1/order/list", _UserMate_OrderList0_HTTP_Handler(srv))
	r.POST("/api/v1/order/detail", _UserMate_OrderDetail0_HTTP_Handler(srv))
	r.POST("/api/v1/level/create", _UserMate_AddLevel0_HTTP_Handler(srv))
	r.POST("/api/v1/category/create", _UserMate_AddServiceCategory0_HTTP_Handler(srv))
}

func _UserMate_AddUserMate0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateAddUserMate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddUserMate(ctx, req.(*UserMateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMateReply)
		return ctx.Result(200, reply)
	}
}

func _UserMate_DeleteUserMate0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateDeleteUserMate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUserMate(ctx, req.(*DeleteMateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMateReply)
		return ctx.Result(200, reply)
	}
}

func _UserMate_UserMateDetail0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMateShowRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateUserMateDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserMateDetail(ctx, req.(*UserMateShowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMateShowReply)
		return ctx.Result(200, reply)
	}
}

func _UserMate_UpdateUserMate0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserMateUpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateUpdateUserMate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserMate(ctx, req.(*UserMateUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserMateUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _UserMate_ListUserMate0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateListUserMate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserMate(ctx, req.(*ListMateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMateResponse)
		return ctx.Result(200, reply)
	}
}

func _UserMate_SearchUserMate0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchUserMateRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateSearchUserMate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SearchUserMate(ctx, req.(*SearchUserMateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SearchUserMateResponse)
		return ctx.Result(200, reply)
	}
}

func _UserMate_CreateOrder0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateCreateOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrder(ctx, req.(*CreateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrderReply)
		return ctx.Result(200, reply)
	}
}

func _UserMate_UpdateOrder0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateUpdateOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOrder(ctx, req.(*UpdateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateOrderReply)
		return ctx.Result(200, reply)
	}
}

func _UserMate_OrderList0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateOrderList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderList(ctx, req.(*OrderListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderListResponse)
		return ctx.Result(200, reply)
	}
}

func _UserMate_OrderDetail0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OrderDetailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateOrderDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OrderDetail(ctx, req.(*OrderDetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OrderDetailResponse)
		return ctx.Result(200, reply)
	}
}

func _UserMate_AddLevel0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LevelRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateAddLevel)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddLevel(ctx, req.(*LevelRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LevelResponse)
		return ctx.Result(200, reply)
	}
}

func _UserMate_AddServiceCategory0_HTTP_Handler(srv UserMateHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ServiceCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserMateAddServiceCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddServiceCategory(ctx, req.(*ServiceCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ServiceCategoryResponse)
		return ctx.Result(200, reply)
	}
}

type UserMateHTTPClient interface {
	AddLevel(ctx context.Context, req *LevelRequest, opts ...http.CallOption) (rsp *LevelResponse, err error)
	AddServiceCategory(ctx context.Context, req *ServiceCategoryRequest, opts ...http.CallOption) (rsp *ServiceCategoryResponse, err error)
	AddUserMate(ctx context.Context, req *UserMateRequest, opts ...http.CallOption) (rsp *UserMateReply, err error)
	CreateOrder(ctx context.Context, req *CreateOrderRequest, opts ...http.CallOption) (rsp *CreateOrderReply, err error)
	DeleteUserMate(ctx context.Context, req *DeleteMateRequest, opts ...http.CallOption) (rsp *DeleteMateReply, err error)
	ListUserMate(ctx context.Context, req *ListMateRequest, opts ...http.CallOption) (rsp *ListMateResponse, err error)
	OrderDetail(ctx context.Context, req *OrderDetailRequest, opts ...http.CallOption) (rsp *OrderDetailResponse, err error)
	OrderList(ctx context.Context, req *OrderListRequest, opts ...http.CallOption) (rsp *OrderListResponse, err error)
	SearchUserMate(ctx context.Context, req *SearchUserMateRequest, opts ...http.CallOption) (rsp *SearchUserMateResponse, err error)
	UpdateOrder(ctx context.Context, req *UpdateOrderRequest, opts ...http.CallOption) (rsp *UpdateOrderReply, err error)
	UpdateUserMate(ctx context.Context, req *UserMateUpdateRequest, opts ...http.CallOption) (rsp *UserMateUpdateReply, err error)
	UserMateDetail(ctx context.Context, req *UserMateShowRequest, opts ...http.CallOption) (rsp *UserMateShowReply, err error)
}

type UserMateHTTPClientImpl struct {
	cc *http.Client
}

func NewUserMateHTTPClient(client *http.Client) UserMateHTTPClient {
	return &UserMateHTTPClientImpl{client}
}

func (c *UserMateHTTPClientImpl) AddLevel(ctx context.Context, in *LevelRequest, opts ...http.CallOption) (*LevelResponse, error) {
	var out LevelResponse
	pattern := "/api/v1/level/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateAddLevel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) AddServiceCategory(ctx context.Context, in *ServiceCategoryRequest, opts ...http.CallOption) (*ServiceCategoryResponse, error) {
	var out ServiceCategoryResponse
	pattern := "/api/v1/category/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateAddServiceCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) AddUserMate(ctx context.Context, in *UserMateRequest, opts ...http.CallOption) (*UserMateReply, error) {
	var out UserMateReply
	pattern := "/api/v1/usermate/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateAddUserMate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...http.CallOption) (*CreateOrderReply, error) {
	var out CreateOrderReply
	pattern := "/api/v1/order/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateCreateOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) DeleteUserMate(ctx context.Context, in *DeleteMateRequest, opts ...http.CallOption) (*DeleteMateReply, error) {
	var out DeleteMateReply
	pattern := "/api/v1/mate/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserMateDeleteUserMate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) ListUserMate(ctx context.Context, in *ListMateRequest, opts ...http.CallOption) (*ListMateResponse, error) {
	var out ListMateResponse
	pattern := "/api/v1/usermate/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserMateListUserMate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) OrderDetail(ctx context.Context, in *OrderDetailRequest, opts ...http.CallOption) (*OrderDetailResponse, error) {
	var out OrderDetailResponse
	pattern := "/api/v1/order/detail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateOrderDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) OrderList(ctx context.Context, in *OrderListRequest, opts ...http.CallOption) (*OrderListResponse, error) {
	var out OrderListResponse
	pattern := "/api/v1/order/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateOrderList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) SearchUserMate(ctx context.Context, in *SearchUserMateRequest, opts ...http.CallOption) (*SearchUserMateResponse, error) {
	var out SearchUserMateResponse
	pattern := "/api/v1/usermate/search/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserMateSearchUserMate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...http.CallOption) (*UpdateOrderReply, error) {
	var out UpdateOrderReply
	pattern := "/api/v1/order/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateUpdateOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) UpdateUserMate(ctx context.Context, in *UserMateUpdateRequest, opts ...http.CallOption) (*UserMateUpdateReply, error) {
	var out UserMateUpdateReply
	pattern := "/api/v1/usermate/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserMateUpdateUserMate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserMateHTTPClientImpl) UserMateDetail(ctx context.Context, in *UserMateShowRequest, opts ...http.CallOption) (*UserMateShowReply, error) {
	var out UserMateShowReply
	pattern := "/api/v1/mate/detail/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserMateUserMateDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
