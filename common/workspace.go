package common

import (
	"os"

	"github.com/soerenkoehler/go-util-mutation/util"
)

func InitWorkspace() (err error) {
	if _, err = os.Stat(WorkDir); os.IsNotExist(err) {
		err = os.MkdirAll(WorkDir, 0755)
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
	return
}
