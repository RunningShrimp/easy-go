package core

import (
	"encoding/json"
	"fmt"
	"github.com/RunningShrimp/easy-go/core/log" //nolint:gci
	"github.com/RunningShrimp/easy-go/core/router"
	"io"
	"net/http"
	"reflect"
	"strconv" //nolint:gci
)

type EasyGoServeHTTP struct {
	router router.EasyGoHTTPRouter
}

func DefaultEasyGoServeHTTP() *EasyGoServeHTTP {
	return &EasyGoServeHTTP{
		router.NewMappingRouter(),
	}
}

// http 处理引擎，不对外暴露
func (s *EasyGoServeHTTP) ServeHTTP(writer http.ResponseWriter, request *http.Request) { // 1. 获取请求方法与url
	httpMethod := request.Method
	urlStr := request.URL.Path
	// 2. 根据请求方法和url获取handler
	handleFunc, ok, statusCode := s.router.FindHandlerByMethodURL(urlStr, httpMethod)
	if !ok {
		writer.WriteHeader(statusCode)
	}

	data := make(map[string]any) //nolint:nolintlint
	bodyData := request.Body
	defer func(bodyData io.ReadCloser) {
		err := bodyData.Close()
		if err != nil {
			log.Log.Error("EasyGoServeHTTP.ServeHTTP error")
		}
	}(bodyData)

	bytes, err := io.ReadAll(bodyData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO.md:支持url编辑参数
	if len(bytes) == 0 {
		for k, v := range request.URL.Query() { // 这里因为取得数据为字符串数组，只要长度为1则认为是字符串
			if len(v) == 1 {
				data[k] = v[0]
			}
		}

		for k, v := range request.Form { // 这里因为取得数据为字符串数组，只要长度为1则认为是字符串
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
	s.dispatchRequest(writer, data, &handleFunc)
}

//nolint:cyclop
func (s *EasyGoServeHTTP) dispatchRequest(
	writer http.ResponseWriter,
	data map[string]any,
	egFunc *router.EasyGoHandlerFunc,
) { //nolint:nolintlint
	if s.router == nil {
		panic("请注册路由")
	}

	argValues := make([]reflect.Value, 0)
	for _, e := range egFunc.InParameter {
		fmt.Println(*e)
		argValues = append(argValues, s.mapStruct(data, *e))
	}
	resultArr := egFunc.HFunc.Call(argValues)

	if len(resultArr) > 0 {
		val := resultArr[0]

		switch val.Kind() { //nolint:exhaustive
		case reflect.Slice:
			_, _ = fmt.Fprintf(writer, "%v", val.String())
			return
		case reflect.Bool:
			_, _ = fmt.Fprintf(writer, "%v", val.Bool())
			return
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			_, _ = fmt.Fprintf(writer, "%d", val.Int())
			return
		case reflect.Uint, reflect.Uint8, reflect.Uint64, reflect.Uint16, reflect.Uint32:
			_, _ = fmt.Fprintf(writer, "%d", val.Uint())
			return
		case reflect.Float32, reflect.Float64:
			_, _ = fmt.Fprintf(writer, "%f", val.Float())
			return
		case reflect.String:

			_, _ = fmt.Fprintf(writer, "%s", val.String())
			return
		case reflect.Struct:
			bytes, err := json.Marshal(val.Bytes())
			if err != nil {
				return
			}
			_, _ = fmt.Fprintf(writer, "%s", string(bytes))
			return
		default:
			_, err := writer.Write(val.Bytes())
			if err != nil {
				return
			}
			writer.WriteHeader(http.StatusOK)
		}
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

//nolint:funlen
func (s *EasyGoServeHTTP) mapStruct(
	data map[string]any,
	argType reflect.Type,
) reflect.Value { //nolint:nolintlint
	val := reflect.New(argType)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < argType.NumField(); i++ {
		t := argType.Field(i)
		f := val.Field(i)

		tag := t.Tag.Get("json")
		v, ok := data[tag]

		if !ok {
			return val
		}

		dataType := reflect.TypeOf(v)
		structType := f.Type()

		if structType == dataType {
			f.Set(reflect.ValueOf(v))
			return val
		}

		if dataType.ConvertibleTo(structType) {
			// 转换类型
			f.Set(reflect.ValueOf(v).Convert(structType))
			return val
		}

		switch structType.Kind() { //nolint:exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				// 这里只给提示便可以，不需要处理错误
				// TODO.md：未来这里需要优化
				log.Log.Info("数据格式错误")

				break
			}
			f.SetInt(v)
		case reflect.Float32, reflect.Float64:
			v, err := strconv.ParseFloat(v.(string), 64)
			if err != nil {
				log.Log.Info("数据格式错误")
				break
			}

			f.SetFloat(v)
		case reflect.Uint, reflect.Uint8, reflect.Uint64, reflect.Uint16, reflect.Uint32:
			v, err := strconv.ParseUint(v.(string), 10, 64)
			if err != nil {
				// 这里只给提示便可以，不需要处理错误
				// TODO.md：未来这里需要优化
				log.Log.Info("数据格式错误")

				break
			}
			f.SetUint(v)
		case reflect.Bool:
			v, err := strconv.ParseBool(v.(string))
			if err != nil {
				// 这里只给提示便可以，不需要处理错误
				// TODO.md：未来这里需要优化
				log.Log.Info("数据格式错误")

				break
			}
			f.SetBool(v)
		default:
			panic(t.Name + " type mismatch")
		}
	}
	return val
}
