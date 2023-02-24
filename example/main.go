package main

import (
	easygo "github.com/RunningShrimp/easy-go"
	"github.com/RunningShrimp/easy-go/example/api"
)

func main() {
	easygoApp := easygo.NewEasyGo()

	{
		router := easygoApp.NewRouter()
		router.Get("/user/info", api.UserInfo)
		//todo: 下期重构
		//router.RestGroup("/user", &api.UserGroup.RegisterRouter)
	}

	easygoApp.Run()

}
