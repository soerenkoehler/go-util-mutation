package mutator

import (
	"go/ast"
	"go/parser"
	"go/token"
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
					err = mutateFile(path)
				}
			}
			return nil
		},
		doublestar.WithFilesOnly())
}

func mutateFile(file string) (err error) {
	util.Debug("Mutating %s", file)

	root, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
	if err != nil {
		util.Fatal("parsing %v: %v", file, err)
	}

	ast.Walk(mutator{}, root)

	return
}

type mutator struct{}

func (m mutator) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.BinaryExpr:
		return mutateBinaryExpr(n)
	case *ast.UnaryExpr:
		return mutateUnaryExpr(n)
	case *ast.AssignStmt:
		return mutateAssignStmt(n)
	case *ast.CallExpr:
		return mutateCallExpr(n)
	case *ast.ReturnStmt:
		return mutateReturnStmt(n)
	}
	return nil
}

func mutateReturnStmt(n *ast.ReturnStmt) ast.Visitor {
	panic("unimplemented")
}

func mutateCallExpr(n *ast.CallExpr) ast.Visitor {
	panic("unimplemented")
}

func mutateAssignStmt(n *ast.AssignStmt) ast.Visitor {
	panic("unimplemented")
}

func mutateUnaryExpr(n *ast.UnaryExpr) ast.Visitor {
	panic("unimplemented")
}

func mutateBinaryExpr(n *ast.BinaryExpr) ast.Visitor {
	panic("unimplemented")
}
