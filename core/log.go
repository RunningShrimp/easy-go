package core

import (
	"github.com/RunningShrimp/easy-go/config"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func init() {
	var err error

	switch config.RsConfig.Env {
	case "dev":
		Logger, err = zap.NewDevelopment()
	case "prd":
		Logger, err = zap.NewProduction()
	default:
		Logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
}
