package mutator

import "go/ast"

type mutationContext struct {
	dir  string
	file string
	root *ast.File
}

func New(dir string) mutationContext {
	return mutationContext{
		dir:  dir,
		file: "",
		root: nil,
	}
}
