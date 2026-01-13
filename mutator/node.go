package mutator

import (
	"go/ast"
	"go/token"
)

func (ctx mutationContext) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.ImportSpec:
		return nil
	case *ast.BasicLit:
		ctx.mutateBasicLit(n)
		// case *ast.BinaryExpr:
		// 	return mutateBinaryExpr(n)
		// case *ast.UnaryExpr:
		// 	return mutateUnaryExpr(n)
		// case *ast.AssignStmt:
		// 	return mutateAssignStmt(n)
		// case *ast.CallExpr:
		// 	return mutateCallExpr(n)
		// case *ast.ReturnStmt:
		// 	return mutateReturnStmt(n)
	}
	return ctx
}

func (ctx mutationContext) mutateBasicLit(n *ast.BasicLit) {
	old := n.Value

	testMutatedBasicLit := func(value string) {
		if value != old {
			ctx.testMutation("mutate string", n.Pos())
		}
	}

	switch n.Kind {
	case token.STRING:
		testMutatedBasicLit(`""`)
		testMutatedBasicLit(`"mutated"`)
	case token.INT:
	}

	n.Value = old

	a := 00
	println(a)
}

// func mutateBinaryExpr(n *ast.BinaryExpr) ast.Visitor {
// 	panic("unimplemented")
// }

// func mutateUnaryExpr(n *ast.UnaryExpr) ast.Visitor {
// 	panic("unimplemented")
// }

// func mutateAssignStmt(n *ast.AssignStmt) ast.Visitor {
// 	panic("unimplemented")
// }

// func mutateCallExpr(n *ast.CallExpr) ast.Visitor {
// 	panic("unimplemented")
// }

// func mutateReturnStmt(n *ast.ReturnStmt) ast.Visitor {
// 	panic("unimplemented")
// }
