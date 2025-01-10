package common

import (
	"encoding/json"
	"os"
	"path"

	"github.com/soerenkoehler/go-util-mutation/util"
)

type ConfigData struct {
	Copy       []string
	DontMutate []string
}

var Config *ConfigData

func (cfg *ConfigData) load(filename string) (err error) {
	configFile := path.Join(WorkDir, filename)

	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		util.Debug("Initialize configuration file %s", configFile)
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
