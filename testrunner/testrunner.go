package testrunner

import (
	"os"
	"os/exec"
)

type TestRunner struct {
	process *exec.Cmd
}

func New() TestRunner {
	return TestRunner{process: exec.Command("go", "test", "./...")}
}

func (tr TestRunner) WithOutput() TestRunner {
	tr.process.Stdout = os.Stdout
	tr.process.Stderr = os.Stdout
	return tr
}

func (tr TestRunner) WithDir(dir string) TestRunner {
	tr.process.Dir = dir
	return tr
}

func (tr TestRunner) Run() error {
	return tr.process.Run()
}
