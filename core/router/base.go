package router

import (
	"context"
	"github.com/RunningShrimp/easy-go/core/log"
	"reflect"
)

var routerLog = log.Log

func init() {
	MRoutes = newMappingRouter()
}

type RouteRegister interface {
	Get(patten string, handler any)
	Post(patten string, handler any)
	Put(patten string, handler any)
	Delete(patten string, handler any)
	RestGroup(patten string, controller RestFulGrouper)
}

var nullHandlerFunc = EasyGoHandlerFunc{}

type EasyGoHandlerFunc struct {
	InParameter  []*reflect.Type // 入参列表,按照
	OutParameter []*reflect.Type // 出参列表
	HFunc        reflect.Value   // 处理方法
}

type EasyGoHttpRouter interface {
	FindHandlerByMethodUrl(method, urlPattern string) (EasyGoHandlerFunc, bool, int)
	RouteRegister
}

// EasyGoContext request -> mid with RsCtx -> handler(ctx, req, res)
type EasyGoContext struct {
	ctx context.Context
	env string // init
}

func (c EasyGoContext) GetEnv() string {
	return c.env
}

func (c EasyGoContext) Context() context.Context {
	return c.ctx
}

type RestFulGrouper interface {
	RegisterRouter(ctx *EasyGoContext)
}
type IRequest interface {
	InitRequest(ctx *EasyGoContext)
}

// RestFulGroup 基础的Controller，
// 所有的实现自动注册的rest风格的api都应该继承RestFulGrouper
type RestFulGroup struct {
	Ctx *EasyGoContext
}

// BaseRequest 基础的request参数，所有的请求处理handler的参数都应该继承BaseRequest
type BaseRequest struct {
	Ctx *EasyGoContext
}

func (c *RestFulGroup) RegisterRouter(ctx *EasyGoContext) {
	c.Ctx = ctx
}

func (r *BaseRequest) InitRequest(ctx *EasyGoContext) {
	r.Ctx = ctx
}
