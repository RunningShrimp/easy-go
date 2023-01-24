package server

import "reflect"

type handlerInfo struct {
	in    []*reflect.Type // 入参列表,按照
	out   []*reflect.Type // 出参列表
	value reflect.Value   // 处理方法
}

var routes map[string]map[string]*handlerInfo
