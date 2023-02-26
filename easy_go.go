package easygo

import (
	"context"
	"fmt"
	"github.com/RunningShrimp/easy-go/core"
	"github.com/RunningShrimp/easy-go/core/router"
	"net/http"

	"github.com/RunningShrimp/easy-go/core/log"

	"go.uber.org/zap"
	"time"
)

//var log = log.Log

type easyGoCtx struct {
	parentCtx context.Context
}

type RouteRegister interface {
	Get(patten string, handler any)
	Post(patten string, handler any)
	Put(patten string, handler any)
	Delete(patten string, handler any)
	RestGroup(patten string, controller router.RestFulGroup)
}

// EasyGo 启动实例
type EasyGo struct {
	// 支持自定义 HTTP baseServer handler
	baseServer *http.Server

	serveHandler http.Handler

	// 有默认值
	name string
	port string

	ctx easyGoCtx

	timeOut time.Duration
	maxConn int64

	// 支持从配置文件读取配置，方便统一管理配置，但大部分都是代码里硬编码
	appConfigYamlFilePath string

	route router.EasyGoHttpRouter
}

func (g EasyGo) NewRouter() router.EasyGoHttpRouter {

	return g.route
}

func NewEasyGo(options ...Option) *EasyGo {

	easyGo := &EasyGo{
		serveHandler: core.DefaultEasyGoServeHTTP(),
		port:         "2357",
		name:         "EasyGo",
		route:        router.MRoutes,
	}

	for _, opt := range options {
		opt(easyGo)
	}

	if easyGo.appConfigYamlFilePath != "" {
		// init config from file
	}

	if easyGo.baseServer == nil {
		easyGo.baseServer = &http.Server{
			Addr:    ":" + easyGo.port,
			Handler: core.DefaultEasyGoServeHTTP(),
		}
	}

	// init log

	// listening single or kill

	return easyGo
}

func (g EasyGo) Run() {

	log.Log.Info(fmt.Sprintf("[Name-%s-Port-%s] HTTP server is running.", g.name, g.port))
	err := g.baseServer.ListenAndServe()
	if err != nil {
		log.Log.Fatal("Server run failed", zap.String("err", err.Error()))
	}
}
