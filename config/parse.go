package config

import (
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func getEnvCfgFilePath(entranceFile string) string {
	resourceDir, resourceFile := getEntranceDirectory(entranceFile)

	envType := getEnvType(resourceFile)

	err := filepath.WalkDir(
		resourceDir,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			fileInfo, err := d.Info()
			if err != nil {
				return err
			}

			// todo: 多个文件怎么办？
			if strings.Contains(fileInfo.Name(), envType) {
				RsConfig.EnvCfgFilePath = path
			}

			return err
		},
	)
	if err != nil {
		panic(err)
	}

	return RsConfig.EnvCfgFilePath
}
func getEnvType(appFile string) string {
	type (
		appInfo struct {
			Profile struct {
				Active string `yaml:"active"`
			} `yaml:"profile"`
		}
	)

	pro := appInfo{}

	content, err := os.ReadFile(appFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &pro)
	if err != nil {
		panic(err)
	}

	RsConfig.Env = pro.Profile.Active

	return pro.Profile.Active
}

func getEntranceDirectory(configPathPrefix string) (string, string) {
	appEntranceDirectory, err := filepath.Abs(configPathPrefix)
	if err != nil {
		panic(err)
	}

	return appEntranceDirectory, filepath.Join(appEntranceDirectory, "resource", "app.yaml")
}
