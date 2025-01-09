package common

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/soerenkoehler/go-util-mutation/util"
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

func InitWorkspace(configFile string) (err error) {
	if _, err = os.Stat(WorkDir); os.IsNotExist(err) {
		util.Debug("Creating workspace directory %s", WorkDir)
		err = os.MkdirAll(WorkDir, 0755)
	}

	if err == nil {
		err = Config.load(configFile)
	}

	return
}

func InitMutationDir() (err error) {
	if err = os.RemoveAll(MutationDir); err == nil || os.IsNotExist(err) {
		for _, pattern := range Config.Copy {
			if err == nil {
				err = util.GlobCopy(".", MutationDir, pattern)
			}
		}
	}

	if err == nil {
		util.Debug("Running tests before mutation")
		if RunTests() != nil {
			return fmt.Errorf("tests failed on unmutated sources")
		}
	}

	return
}

func RunTests() (err error) {
	proc := exec.Command("go", "test", "./...")
	proc.Dir = MutationDir
	return proc.Run()
}

func (cfg *ConfigData) load(filename string) (err error) {
	if path.Ext(filename) == "" {
		filename += ".json"
	}
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
