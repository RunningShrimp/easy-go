package server

import (
	"net/http"
	"reflect"
	"shrimp_blog_sever/framework/app"
)

// Rest 批量添加路由，添加GET，POST，PUT，DELETE方法，暂不支持从路由上解析参数
// TODO: 从路由解析参数
func Rest(patten string, controller app.IController) {
	methodValue := reflect.ValueOf(controller)
	getMethod := methodValue.MethodByName("Get")
	postMethod := methodValue.MethodByName("Post")
	putMethod := methodValue.MethodByName("Put")
	deleteMethod := methodValue.MethodByName("Delete")
	if getMethod.IsValid() {
		handlerRouter(http.MethodGet, patten, getMethod, getMethod.Type())
	}
	if postMethod.IsValid() {
		handlerRouter(http.MethodGet, patten, postMethod, postMethod.Type())
	}
	if putMethod.IsValid() {
		handlerRouter(http.MethodGet, patten, putMethod, putMethod.Type())
	}
	if deleteMethod.IsValid() {
		handlerRouter(http.MethodGet, patten, deleteMethod, deleteMethod.Type())
	}
}

// Get http-get
func Get(patten string, handler any) {
	addRouter(http.MethodGet, patten, handler)
}

// Post http-post
func Post(patten string, handler any) {
	addRouter(http.MethodPost, patten, handler)
}

// Put http-put
func Put(patten string, handler any) {
	addRouter(http.MethodPut, patten, handler)
}

// Delete http-delete
func Delete(patten string, handler any) {
	addRouter(http.MethodDelete, patten, handler)
}

func addRouter(method, patten string, handler any) {
	if routes == nil {
		routes = make(map[string]map[string]*handlerInfo, 4)
	}
	if _, ok := routes[method]; !ok {
		routes[method] = make(map[string]*handlerInfo)
	}

	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	handlerRouter(method, patten, handlerValue, handlerType)
}

func handlerRouter(method, patten string, handlerValue reflect.Value, handlerType reflect.Type) {
	if handlerType.Kind() != reflect.Func {
		panic("请添加方法")
	}

	argInNum := handlerType.NumIn()
	argOutNum := handlerType.NumOut()
	info := &handlerInfo{
		in:    make([]*reflect.Type, argInNum),
		out:   make([]*reflect.Type, argOutNum),
		value: handlerValue,
	}
	//TODO:获取入参名称
	for i := 0; i < argInNum; i++ {
		in := handlerType.In(i)

		info.in[i] = &in
	}
	//获取出参名称
	for i := 0; i < argOutNum; i++ {
		out := handlerType.Out(i)
		info.out[i] = &out
	}
	routes[method][patten] = info
}
