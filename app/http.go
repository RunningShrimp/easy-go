package app

import "context"

type IController interface {
	InitController(ctx context.Context)
}
type IRequest interface {
	InitRequest(ctx context.Context)
}

type BaseController struct {
	Ctx context.Context
}

type BaseRequest struct {
	Ctx context.Context
}

func (c *BaseController) InitController(ctx context.Context) {
	c.Ctx = ctx
}

func (r *BaseRequest) InitRequest(ctx context.Context) {
	r.Ctx = ctx
}
