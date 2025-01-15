package mutator

import "go/ast"

type nodeMutator struct{}

func (m nodeMutator) Visit(node ast.Node) ast.Visitor {
	return mutateNode(node)
}

func mutateNode(node ast.Node) ast.Visitor {
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
