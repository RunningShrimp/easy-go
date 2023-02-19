package easygo

import (
	"context"
	"fmt"
	"github.com/RunningShrimp/easy-go/core"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type easyGoCtx struct {
	parentCtx context.Context
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

	DEBUG bool
}

func NewEasyGo(options ...Option) *EasyGo {

	easyGo := &EasyGo{
		serveHandler: core.DefaultEasyGoServeHTTP(),
		port:         "2357",
		name:         "EasyGo",
	}

	for _, opt := range options {
		opt(easyGo)
	}

	if easyGo.appConfigYamlFilePath != "" {
		// init config from file
	}

	// init log

	// listening single or kill

	return easyGo
}

func (easyGo *EasyGo) Run() {

	if easyGo.baseServer == nil {
		easyGo.baseServer = &http.Server{
			Addr:    easyGo.port,
			Handler: core.DefaultEasyGoServeHTTP(),
		}
	}

	err := easyGo.baseServer.ListenAndServe()
	if err != nil {
		core.Log.Fatal("Server run failed", zap.String("err", err.Error()))
	}
	core.Log.Info(fmt.Sprintf("[Name-%s-Port-%s] HTTP server is running.", easyGo.name, easyGo.port))
}
