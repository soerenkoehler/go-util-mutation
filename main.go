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

// func createAST(dir string) map[string]*ast.File {
// 	files := make(chan string)

// 	go func() {
// 		defer close(files)
// 		if err := fs.WalkDir(
// 			os.DirFS(dir),
// 			".",
// 			func(path string, entry fs.DirEntry, err error) error {
// 				if err != nil {
// 					return err
// 				}
// 				if !entry.IsDir() {
// 					files <- path
// 				}
// 				return nil
// 			}); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	result := map[string]*ast.File{}

// 	for file := range files {
// 		if path.Ext(file) == ".go" {
// 			if ast, err := parser.ParseFile(token.NewFileSet(), file, nil, 0); err != nil {
// 				panic(err)
// 			} else {
// 				result[file] = ast
// 			}
// 		}
// 	}

// 	return result
// }
