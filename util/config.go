package util

import (
	_ "embed"
	"encoding/json"
	"os"
	"path"
)

const (
	WorkFolder        = ".go-util-mutation"
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

func (cfg *ConfigData) Load(filename string) error {
	configFile := path.Join(WorkFolder, filename)

	if _, err := os.Stat(configFile); err != nil {
		os.MkdirAll(WorkFolder, 0755)
		os.WriteFile(configFile, _defaultConfiguration, 0644)
	}

	data := []byte{}
	chain := &ChainContext{}

	return chain.Chain(func() {
		data, chain.Err = os.ReadFile(configFile)
	}).Chain(func() {
		chain.Err = json.Unmarshal(data, &Config)
	}).Err
}
