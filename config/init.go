package config

var (
	RsConfig = &AppConfig{ //nolint:gochecknoglobals
		EnvCfgFilePath: "",
		Env:            "",
		WrapperConfig:  WrapperConfig{},
	}
)

type IInit interface {
	InitDBClient(wrapperConfig WrapperConfig)
}
