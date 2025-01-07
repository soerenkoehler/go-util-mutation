package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"

	"github.com/soerenkoehler/go-util-mutation/common"
	"github.com/soerenkoehler/go-util-mutation/util"
)

func main() {
	util.InitLogger(os.Stdout)
	util.SetLogLevel(util.LOG_INFO)

	var configFile string

	configFile, err := common.GetConfigName()
	if configFile == "" {
		fmt.Println("Usage: go run soerenkoehler.de/go-util-mutation [CONFIG-NAME]")
	}

	if err == nil {
		err = common.InitWorkspace()
	}

	if err == nil {
		err = common.Config.Load(configFile)
	}

	if err == nil {
		err = common.InitMutationDir()
	}

	if err != nil {
		util.Error("%v\n", err)
	}
}

func createAST(dir string) map[string]*ast.File {
	files := make(chan string)

	go func() {
		defer close(files)
		if err := fs.WalkDir(
			os.DirFS(dir),
			".",
			func(path string, entry fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !entry.IsDir() {
					files <- path
				}
				return nil
			}); err != nil {
			panic(err)
		}
	}()

	result := map[string]*ast.File{}

	for file := range files {
		if path.Ext(file) == ".go" {
			if ast, err := parser.ParseFile(token.NewFileSet(), file, nil, 0); err != nil {
				panic(err)
			} else {
				result[file] = ast
			}
		}
	}

	return result
}
