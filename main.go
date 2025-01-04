package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
)

const (
	ConfigFolder  = ".go-util-mutation"
	DefaultConfig = "default.yaml"
)

var (
	// embed: default.yaml
	_defaultConfiguration string
)

func main() {
	for _, file := range createAST(".") {
		format.Node(os.Stdout, token.NewFileSet(), file)
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
