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
		router.RestGroup("/user", &api.UserController{})
	}

	easygoApp.Run()

}
