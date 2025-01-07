package common

import (
	"os"
	"path"

	"github.com/soerenkoehler/go-util-mutation/util"
)

func InitWorkspace() (err error) {
	if _, err = os.Stat(WorkDir); os.IsNotExist(err) {
		err = os.MkdirAll(WorkDir, 0755)
	}
	return
}

func InitMutationDir() (err error) {
	if err = os.RemoveAll(MutationDir); err == nil {
		for file := range util.ReadDir(".") {
			if path.Ext(file) == ".go" && err == nil {
				err = util.CopyFile(".", MutationDir, file)
			}
		}
	}
	return
}
