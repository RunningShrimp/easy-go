package server

import (
	"encoding/json"
	"fmt"
	"github.com/RunningShrimp/easy-go/config"
	"github.com/RunningShrimp/easy-go/logs"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Server struct {
	name string
	port string
}

func NewServer(path string, name ...string) *Server {
	var serverName = config.RsConfig.Config.Server.Name
	if serverName == "" {
		if len(name) == 1 {
			serverName = name[0]
		} else if len(name) > 1 {
			serverName = strings.Join(name, "-")
		}
	}
	config.Init(path)
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
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(r)
	//		writer.WriteHeader(http.StatusInternalServerError)
	//	}
	//}()

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
		if len(bytes) == 0 || bytes == nil {
			form := request.Form
			if form == nil {
				form = request.URL.Query()
			}
			for k, v := range form { // 这里因为取得数据为字符串数组，只要长度为1则认为是字符串
				if len(v) == 1 {
					data[k] = v[0]
				}
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
			fmt.Println(dataType)
			structType := f.Type()
			fmt.Println(structType)
			if structType == dataType {
				f.Set(reflect.ValueOf(v))
			} else {
				if dataType.ConvertibleTo(structType) {
					// 转换类型
					f.Set(reflect.ValueOf(v).Convert(structType))
				} else {
					switch structType.Kind() {
					case reflect.Int:
					case reflect.Int8:
					case reflect.Int16:
					case reflect.Int32:
					case reflect.Int64:
						v, err := strconv.ParseInt(v.(string), 10, 64)
						if err != nil {
							// 这里只给提示便可以，不需要处理错误
							//TODO：未来这里需要优化
							logs.Logger.Info("数据格式错误")
							break
						}
						f.SetInt(v)
						break
					case reflect.Float32:
					case reflect.Float64:
						v, err := strconv.ParseFloat(v.(string), 64)
						if err != nil {
							logs.Logger.Info("数据格式错误")
							break
						}

						f.SetFloat(v)
						break
					case reflect.Uint:
					case reflect.Uint8:
					case reflect.Uint16:
					case reflect.Uint32:
					case reflect.Uint64:
						v, err := strconv.ParseUint(v.(string), 10, 64)
						if err != nil {
							// 这里只给提示便可以，不需要处理错误
							//TODO：未来这里需要优化
							logs.Logger.Info("数据格式错误")
							break
						}
						f.SetUint(v)
						break
					case reflect.Bool:
						v, err := strconv.ParseBool(v.(string))
						if err != nil {
							// 这里只给提示便可以，不需要处理错误
							//TODO：未来这里需要优化
							logs.Logger.Info("数据格式错误")
							break
						}
						f.SetBool(v)
						break
					default:

						panic(t.Name + " type mismatch")
					}
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
		fmt.Println(*e)
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
			return
		}
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}
