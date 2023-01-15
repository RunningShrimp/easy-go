package config

type server struct {
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
}

type redis struct {
	Database int    `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type database struct {
	Type     string `yaml:"type"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type config struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
	Redis    redis    `yaml:"redis"`
}

type WrapperConfig struct {
	Config config `yaml:"app"`
}

type AppConfig struct {
	EnvCfgFilePath string
	Env            string
	WrapperConfig
}
