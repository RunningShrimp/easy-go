package core

type IController interface {
	InitController(ctx *EasyGoContext)
}
type IRequest interface {
	InitRequest(ctx *EasyGoContext)
}

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
