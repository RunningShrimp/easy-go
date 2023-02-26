package router

import (
	"net/http"
	"reflect"
)

type mappingRouter struct {
	routes map[string]map[string]*EasyGoHandlerFunc
}

func NewMappingRouter() *mappingRouter {
	return &mappingRouter{}
}

func (mr *mappingRouter) FindHandlerByMethodURL(urlPattern, method string) (EasyGoHandlerFunc, bool, int) {

	handlers, ok := mr.routes[urlPattern]
	if !ok {
		return nullHandlerFunc, false, http.StatusNotFound
	}

	if _, ok := handlers[method]; !ok {
		return nullHandlerFunc, false, http.StatusMethodNotAllowed
	}

	return *handlers[method], true, http.StatusOK
}

// RestGroup 批量添加路由，添加GET，POST，PUT，DELETE方法，暂不支持从路由上解析参数
// TODO.md: 从路由解析参数
func (mr *mappingRouter) RestGroup(patten string, controller RestFulGrouper) {
	methodValue := reflect.ValueOf(controller)

	if getMethod := methodValue.MethodByName("Get"); getMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, getMethod, getMethod.Type())
	}
	if postMethod := methodValue.MethodByName("Post"); postMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, postMethod, postMethod.Type())
	}
	if putMethod := methodValue.MethodByName("Put"); putMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, putMethod, putMethod.Type())
	}
	if deleteMethod := methodValue.MethodByName("Delete"); deleteMethod.IsValid() {
		mr.handlerRouter(http.MethodGet, patten, deleteMethod, deleteMethod.Type())
	}
}

// Get http-get
func (mr *mappingRouter) Get(patten string, handler any) { //nolint:typecheck
	mr.addRouter(http.MethodGet, patten, handler)
}

// Post http-post
func (mr *mappingRouter) Post(patten string, handler any) { //nolint:typecheck
	mr.addRouter(http.MethodPost, patten, handler)
}

// Put http-put
func (mr *mappingRouter) Put(patten string, handler any) { //nolint:typecheck
	mr.addRouter(http.MethodPut, patten, handler)
}

// Delete http-delete
func (mr *mappingRouter) Delete(patten string, handler any) { //nolint:typecheck
	mr.addRouter(http.MethodDelete, patten, handler)
}

func (mr *mappingRouter) addRouter(method, patten string, handler any) { //nolint:typecheck
	// todo map存在线程安全问题
	if mr.routes == nil {
		mr.routes = make(map[string]map[string]*EasyGoHandlerFunc)
	}
	if _, ok := mr.routes[patten]; !ok {
		mr.routes[patten] = make(map[string]*EasyGoHandlerFunc)
	}

	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	mr.handlerRouter(method, patten, handlerValue, handlerType)
}

func (mr *mappingRouter) handlerRouter(method, patten string, handlerValue reflect.Value, handlerType reflect.Type) {
	if handlerType.Kind() != reflect.Func {
		panic("请添加方法")
	}

	argInNum := handlerType.NumIn()
	argOutNum := handlerType.NumOut()
	handelFunc := &EasyGoHandlerFunc{
		InParameter:  make([]*reflect.Type, argInNum),
		OutParameter: make([]*reflect.Type, argOutNum),
		HFunc:        handlerValue,
	}
	for i := 0; i < argInNum; i++ {
		in := handlerType.In(i)
		handelFunc.InParameter[i] = &in
	}

	for j := 0; j < argOutNum; j++ {
		out := handlerType.Out(j)
		handelFunc.OutParameter[j] = &out
	}

	mr.routes[patten][method] = handelFunc
}
