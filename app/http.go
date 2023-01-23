package app

import "net/http"

type IController interface {
	InitController(ctx *EasyGoContext)
}
type IRequest interface {
	InitRequest(ctx *EasyGoContext)
}

// RequestMid 请求中间件，所有请求中间件都应该实现这个函数
type RequestMid func(*http.Request)

// ResponseMid 响应中间件，所有的响应中间都应该实现这个函数
type ResponseMid func(*http.Response)

// RequestHandler 请求处理的handler
type RequestHandler func(IRequest) string

// BaseController 基础的Controller，
// 所有的实现自动注册的rest风格的api都应该继承BaseController
type BaseController struct {
	Ctx *EasyGoContext
}

// BaseRequest 基础的request参数，所有的请求处理handler的参数都应该继承BaseRequest
type BaseRequest struct {
	Ctx *EasyGoContext
}

func (c *BaseController) InitController(ctx *EasyGoContext) {
	c.Ctx = ctx
}

func (r *BaseRequest) InitRequest(ctx *EasyGoContext) {
	r.Ctx = ctx
}
