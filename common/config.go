package common

import (
	_ "embed"
	"encoding/json"
	"os"
	"path"
)

const (
	WorkDir           = ".go-util-mutation"
	MutationDir       = WorkDir + "/tmp"
	ConfigFileDefault = "default.json"
)

var (
	//go:embed defaultConfig.json
	_defaultConfiguration []byte
)

type ConfigData struct {
	Copy       []string
	DontMutate []string
}

var Config *ConfigData

func (cfg *ConfigData) Load(filename string) (err error) {
	if path.Ext(filename) == "" {
		filename += ".json"
	}
	configFile := path.Join(WorkDir, filename)

	_, err = os.Stat(WorkDir)

	if err != nil {
		os.MkdirAll(
			WorkDir,
			0755)
		os.WriteFile(
			path.Join(WorkDir, ConfigFileDefault),
			_defaultConfiguration,
			0644)
	}

	data, err := os.ReadFile(configFile)

	if err == nil {
		err = json.Unmarshal(data, &Config)
	}

	return
}
