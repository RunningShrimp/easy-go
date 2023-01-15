package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"runtime/debug"
	"shrimp_blog_sever/framework/config"
	"strings"
)

type Server struct {
	name string
	port string
}

func NewServer(name ...string) *Server {
	var serverName = config.RsConfig.Config.Server.Name
	if serverName == "" {
		if len(name) == 1 {
			serverName = name[0]
		} else if len(name) > 1 {
			serverName = strings.Join(name, "-")
		}
	}
	return &Server{
		name: serverName,
		port: config.ServerPort(),
	}
}

func (s *Server) Run() {
	fmt.Printf("%s server is  running\n", s.name)
	err := http.ListenAndServe(s.port, s)
	if err != nil {
		panic(err)
	}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}()

	// 1. 获取请求方法与url
	httpMethod := request.Method
	urlStr := request.URL.Path
	// 2. 根据请求方法和url获取handler
	if info, ok := routes[httpMethod][urlStr]; ok {
		data := make(map[string]any)
		bodyData := request.Body
		defer func(bodyData io.ReadCloser) {
			err := bodyData.Close()
			if err != nil {
				panic(err)
			}
		}(bodyData)

		bytes, err := io.ReadAll(bodyData)
		if err != nil {
			fmt.Println(err)
			return
		}

		//TODO:支持url编辑参数
		if bytes == nil {
			form := request.Form
			for k, v := range form {
				data[k] = v
			}
		} else {
			err = json.Unmarshal(bytes, &data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		s.handleRequest(writer, data, info)
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

}
func (s *Server) dataMapStruct(data map[string]any, argType reflect.Type) reflect.Value {
	val := reflect.New(argType)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < argType.NumField(); i++ {
		t := argType.Field(i)
		f := val.Field(i)
		tag := t.Tag.Get("json")
		if v, ok := data[tag]; ok {
			// 检查是否需要类型转换
			dataType := reflect.TypeOf(v)
			structType := f.Type()
			if structType == dataType {
				f.Set(reflect.ValueOf(v))
			} else {
				if dataType.ConvertibleTo(structType) {
					// 转换类型
					f.Set(reflect.ValueOf(v).Convert(structType))
				} else {
					panic(t.Name + " type mismatch")
				}
			}
		}
	}
	return val
}

func (s *Server) handleRequest(writer http.ResponseWriter, data map[string]any, info *handlerInfo) {
	if routes == nil {
		panic("请注册路由")
	}

	// 3. 获取请求参数
	argValues := make([]reflect.Value, 0)
	// 4. 将请求参数注入到handler参数中
	for _, e := range info.in {
		argValues = append(argValues, s.dataMapStruct(data, *e))
	}
	// 5. 执行handler
	resultArr := info.value.Call(argValues)
	// 6. 获取handler执行结果，返回response
	// for _, v := range resultArr {
	//	// TODO: 检查error
	//
	//}
	//TODO:目前只支持string并且只支持一个返回参数
	if len(resultArr) > 0 {

		_, err := fmt.Fprintf(writer, "%s", resultArr[0].String())
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}
