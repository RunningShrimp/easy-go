package log

import (
	"github.com/RunningShrimp/easy-go/config"
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	var err error

	switch config.RsConfig.Env {
	case "dev":
		Log, err = zap.NewDevelopment()
	case "prd":
		Log, err = zap.NewProduction()
	default:
		Log, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
}
