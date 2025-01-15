package mutator

import (
	"go/ast"
	"go/token"
)

type mutationContext struct {
	dir     string
	fileset *token.FileSet
	root    *ast.File
}

func New(dir string) mutationContext {
	return mutationContext{
		dir:     dir,
		fileset: nil,
		root:    nil,
	}
}
