package common

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

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
		if TestRunnerWithOutput().Run() != nil {
			return fmt.Errorf("tests failed on unmutated sources")
		}
	}

	return
}

func TestRunnerWithOutput() (proc *exec.Cmd) {
	proc = TestRunner()
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stdout
	return
}

func TestRunner() (proc *exec.Cmd) {
	proc = exec.Command("go", "test", "./...")
	proc.Dir = MutationDir
	return
}
