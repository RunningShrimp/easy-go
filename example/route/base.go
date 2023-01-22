package route

import (
	"github.com/RunningShrimp/easy-go/example/api"
	"github.com/RunningShrimp/easy-go/server"
)

func init() {
	server.Get("/user/info", &api.UserInfo)
}
