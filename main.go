package main

import (
	_ "embed"
	"fmt"
	"go/ast"
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
	//go:embed default.yaml
	_defaultConfiguration []byte
)

func main() {
	var configFile string
	if len(os.Args) == 1 {
		configFile = DefaultConfig
	} else if len(os.Args) == 2 {
		configFile = os.Args[1]
	} else {
		panic("Usage: go run soerenkoehler.de/go-util-mutation <config-name>")
	}

	configFile = path.Join(ConfigFolder, configFile)

	if _, err := os.Stat(configFile); err != nil {
		os.MkdirAll(ConfigFolder, 0755)
		os.WriteFile(configFile, _defaultConfiguration, 0644)
	}

	fmt.Printf("%v\n", _defaultConfiguration)

	if config, err := os.ReadFile(configFile); err != nil {
		panic(err)
	} else {
		fmt.Printf("%v\n", config)
	}

	// for _, file := range createAST(".") {
	// 	format.Node(os.Stdout, token.NewFileSet(), file)
	// }
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
