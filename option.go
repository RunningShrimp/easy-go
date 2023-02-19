package easygo

import (
	"context"
	"net/http"
	"time"
)

type Option func(*EasyGo)

func WithName(name string) Option {
	return func(eg *EasyGo) {
		eg.name = name
	}
}

func WithPort(port string) Option {
	return func(eg *EasyGo) {
		eg.port = port
	}
}

func WithConfigPath(yamlFilePath string) Option {
	return func(eg *EasyGo) {
		eg.appConfigYamlFilePath = yamlFilePath
	}
}

func WithTimeOut(t time.Duration) Option {
	return func(eg *EasyGo) {
		eg.timeOut = t
	}
}

func WithCtx(ctx context.Context) Option {
	return func(eg *EasyGo) {
		eg.ctx.parentCtx = ctx
	}
}

func WithMaxConn(maxConn int64) Option {
	return func(eg *EasyGo) {
		eg.maxConn = maxConn
	}
}

func WithHttpServerFunc(httpServerFunc *http.Server) Option {
	return func(eg *EasyGo) {
		eg.baseServer = httpServerFunc
	}
}
