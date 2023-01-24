package main

import (
	_ "github.com/RunningShrimp/easy-go/example/route"
	"github.com/RunningShrimp/easy-go/server"
)

func main() {
	server.NewServer("./example").Run()
}
