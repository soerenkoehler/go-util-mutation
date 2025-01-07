package common

import (
	_ "embed"
	"encoding/json"
	"fmt"
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

	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		err = os.WriteFile(configFile, _defaultConfiguration, 0644)
	}

	if err != nil {
		return
	}

	data, err := os.ReadFile(configFile)

	if err != nil {
		return
	}

	return json.Unmarshal(data, &Config)
}

func GetConfigName() (string, error) {
	if len(os.Args) == 1 {
		return ConfigFileDefault, nil
	} else if len(os.Args) == 2 {
		return os.Args[1], nil
	}
	return "", fmt.Errorf("expected at most one argument for CONFIG-NAME")
}
