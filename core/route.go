package core

import (
	"net/http"
	"reflect"
)

var MRoutes *mappingRoutes

func init() {
	MRoutes = newMappingRoutes()
}

var nullHandlerFunc = handlerFunc{}

type handlerFunc struct {
	in    []*reflect.Type // 入参列表,按照
	out   []*reflect.Type // 出参列表
	value reflect.Value   // 处理方法
}

type EasyGoHttpRouter interface {
	DispatchHandlerByMethodAndUlr(method, urlPattern string) (handlerFunc, bool)
}

type mappingRoutes struct {
	routes map[string]map[string]*handlerFunc
}

func newMappingRoutes() *mappingRoutes {
	return &mappingRoutes{}
}

func (mr *mappingRoutes) DispatchHandlerByMethodAndUlr(method, urlPattern string) (handlerFunc, bool) {

	handlers, ok := mr.routes[method]
	if !ok {
		return nullHandlerFunc, false
	}
	hitHandler, ok := handlers[urlPattern]
	if !ok {
		return nullHandlerFunc, false
	}

	return *hitHandler, ok
}

// RestGroup 批量添加路由，添加GET，POST，PUT，DELETE方法，暂不支持从路由上解析参数
// TODO.md: 从路由解析参数
func (mr *mappingRoutes) RestGroup(patten string, controller IController) {
	methodValue := reflect.ValueOf(controller)
	getMethod := methodValue.MethodByName("Get")
	postMethod := methodValue.MethodByName("Post")
	putMethod := methodValue.MethodByName("Put")
	deleteMethod := methodValue.MethodByName("Delete")
	if getMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, getMethod, getMethod.Type())
	}
	if postMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, postMethod, postMethod.Type())
	}
	if putMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, putMethod, putMethod.Type())
	}
	if deleteMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, deleteMethod, deleteMethod.Type())
	}
}

// Get http-get
func (mr *mappingRoutes) Get(patten string, handler any) {
	mr.addRouter(http.MethodGet, patten, handler)
}

// Post http-post
func (mr *mappingRoutes) Post(patten string, handler any) {
	mr.addRouter(http.MethodPost, patten, handler)
}

// Put http-put
func (mr *mappingRoutes) Put(patten string, handler any) {
	mr.addRouter(http.MethodPut, patten, handler)
}

// Delete http-delete
func (mr *mappingRoutes) Delete(patten string, handler any) {
	mr.addRouter(http.MethodDelete, patten, handler)
}

func (mr *mappingRoutes) addRouter(method, patten string, handler any) {
	//todo map存在线程安全问题
	if mr.routes == nil {
		mr.routes = make(map[string]map[string]*handlerFunc, 4)
	}
	if _, ok := mr.routes[method]; !ok {
		mr.routes[method] = make(map[string]*handlerFunc)
	}

	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	mr.handlerRouter(method, patten, handlerValue, handlerType)
}

func (mr *mappingRoutes) handlerRouter(method, patten string, handlerValue reflect.Value, handlerType reflect.Type) {
	if handlerType.Kind() != reflect.Func {
		panic("请添加方法")
	}

	argInNum := handlerType.NumIn()
	argOutNum := handlerType.NumOut()
	info := &handlerFunc{
		in:    make([]*reflect.Type, argInNum),
		out:   make([]*reflect.Type, argOutNum),
		value: handlerValue,
	}
	for i := 0; i < argInNum; i++ {
		in := handlerType.In(i)

		info.in[i] = &in
	}
	for i := 0; i < argOutNum; i++ {
		out := handlerType.Out(i)
		info.out[i] = &out
	}
	mr.routes[method][patten] = info
}
