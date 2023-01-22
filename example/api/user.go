package api

import (
	"fmt"
	"github.com/RunningShrimp/easy-go/app"
)

type UserRequest struct {
	app.BaseRequest
	Username string `json:"username"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
}

func UserInfo(request *UserRequest) string {
	return fmt.Sprintf("你的名称：%s，你的年龄：%d，你住在：%s", request.Username, request.Age, request.Address)
}
