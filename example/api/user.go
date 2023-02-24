package api

import (
	"fmt"
	"github.com/RunningShrimp/easy-go/core/router"
)

type UserRequest struct {
	router.BaseRequest
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Address  string `json:"address"`
}

func UserInfo(request UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}

type UserGroup struct {
	router.RestFulGrouper
}

func (c *UserGroup) Name(request *UserRequest) string {
	return request.Username
}

func (c *UserGroup) Get(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
func (c *UserGroup) Post(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
func (c *UserGroup) Put(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
func (c *UserGroup) Delete(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
