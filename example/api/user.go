package api

import (
	"fmt"
	"github.com/RunningShrimp/easy-go/core"
)

type UserRequest struct {
	core.BaseRequest
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Address  string `json:"address"`
}

func UserInfo(request UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}

type UserController struct {
	core.BaseController
}

func (c *UserController) Name(request *UserRequest) string {
	return request.Username
}

func (c *UserController) Get(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
func (c *UserController) Post(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
func (c *UserController) Put(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
func (c *UserController) Delete(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
