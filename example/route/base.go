package route

import (
	"github.com/RunningShrimp/easy-go/core"
	"github.com/RunningShrimp/easy-go/example/api"
)

func init() {
	core.Get("/user/info", api.UserInfo)
	core.Rest("/user", &api.UserController{})
}
