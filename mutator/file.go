package mutator

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/soerenkoehler/go-util-mutation/common"
	"github.com/soerenkoehler/go-util-mutation/testrunner"
	"github.com/soerenkoehler/go-util-mutation/util"
)

func (ctx mutationContext) MutateFiles() error {
	return doublestar.GlobWalk(
		os.DirFS(ctx.dir),
		"**/*.go",
		func(path string, d fs.DirEntry) (err error) {
			for _, pattern := range common.Config.DontMutate {
				var match bool
				if match, err = doublestar.Match(pattern, path); err != nil {
					return
				}
				if !match {
					err = ctx.mutateFile(path)
				}
			}
			return
		},
		doublestar.WithFilesOnly())
}

func (ctx mutationContext) mutateFile(file string) (err error) {
	util.Debug("Mutating %s", file)

	ctx.fileset = token.NewFileSet()
	ctx.root, err = parser.ParseFile(ctx.fileset, file, nil, 0)

	if err != nil {
		util.Fatal("parsing %v: %v", file, err)
	}

	ast.Walk(ctx, ctx.root)

	return
}

func (ctx mutationContext) testMutation(label string, pos token.Pos) (failed bool, err error) {
	file := ctx.fileset.File(pos)

	dst, err := os.Create(path.Join(ctx.dir, file.Name()))
	if err != nil {
		return
	}

	defer dst.Close()

	err = format.Node(dst, token.NewFileSet(), ctx.root)
	if err != nil {
		return
	}

	util.Debug("%s at %v", label, file.Position(pos))

	failed = testrunner.New().WithDir(common.MutationDir).Run() != nil

	return
}
