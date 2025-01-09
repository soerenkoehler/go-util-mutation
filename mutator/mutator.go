package mutator

import (
	"io/fs"
	"os"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/soerenkoehler/go-util-mutation/common"
	"github.com/soerenkoehler/go-util-mutation/util"
)

func MutateFiles() error {
	return doublestar.GlobWalk(
		os.DirFS(common.MutationDir),
		"**/*.go",
		func(path string, d fs.DirEntry) (err error) {
			for _, pattern := range common.Config.DontMutate {
				match, err := doublestar.Match(pattern, path)
				if err != nil {
					return err
				}
				if !match {
					MutateFile(path)
				}
			}
			return nil
		},
		doublestar.WithFilesOnly())
}

func MutateFile(file string) {
	util.Debug("Mutating %s", file)
}
