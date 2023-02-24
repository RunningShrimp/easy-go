package middleware

import "net/http"

// Middleware  请求中间件，所有请求中间件都应该实现这个函数
type Middleware func(*http.Request)
