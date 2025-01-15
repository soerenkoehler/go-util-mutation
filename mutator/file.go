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

	ctx.file = file
	ctx.root, err = parser.ParseFile(token.NewFileSet(), ctx.file, nil, 0)

	if err != nil {
		util.Fatal("parsing %v: %v", ctx.file, err)
	}

	ast.Walk(ctx, ctx.root)

	return
}

func (ctx mutationContext) outputFile() (err error) {
	dst, err := os.Create(path.Join(ctx.dir, ctx.file))

	if err != nil {
		return
	}
	defer dst.Close()

	return format.Node(dst, token.NewFileSet(), ctx.root)
}
