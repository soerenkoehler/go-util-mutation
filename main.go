package main

import (
	"flag"
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
	util.SetLogLevel(util.LOG_DEBUG)

	flag.Usage = func() {
		fmt.Printf("Usage: go run soerenkoehler.de/go-util-mutation [CONFIG-NAME]\n")
	}
	flag.Parse()

	var err error
	var printUsage bool

	if flag.NArg() > 1 {
		err = fmt.Errorf("too many arguments")
		printUsage = true
	}

	if err == nil {
		err = common.InitWorkspace(flag.Arg(0))
	}

	if err == nil {
		err = common.InitMutationDir()
	}

	if err != nil {
		util.Error("%v", err)
	}

	if printUsage {
		flag.Usage()
		os.Exit(2)
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
