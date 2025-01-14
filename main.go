package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/soerenkoehler/go-util-mutation/common"
	"github.com/soerenkoehler/go-util-mutation/mutator"
	"github.com/soerenkoehler/go-util-mutation/util"
)

var CLI struct {
	Verbose    bool   `short:"v" help:"Output diagnostic info."`
	ConfigFile string `arg:"" optional:"true" default:"${defaultConfigFile}" help:"Optional configuration file name relative to ${workDir}. Will be created if it does not exist. Default value: ${defaultConfigFile}"`
}

func main() {
	kong.Parse(
		&CLI,
		kong.UsageOnError(),
		kong.Vars{
			"workDir":           common.WorkDir,
			"defaultConfigFile": common.DefaultConfigFile,
		})

	util.InitLogger(os.Stdout)
	if CLI.Verbose {
		util.SetLogLevel(util.LOG_DEBUG)
	} else {
		util.SetLogLevel(util.LOG_INFO)
	}

	err := common.InitWorkspace(CLI.ConfigFile)

	if err == nil {
		err = common.InitMutationDir()
	}

	if err == nil {
		err = mutator.MutateFiles()
	}

	if err != nil {
		util.Error("%v", err)
	}
}
