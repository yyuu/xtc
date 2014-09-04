package parser

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
)

func asExpr(x duck.INode) duck.IExprNode {
  return x.(duck.IExprNode)
}

func asExprs(xs []duck.INode) []duck.IExprNode {
  ys := make([]duck.IExprNode, len(xs))
  for i := range xs {
    ys[i] = asExpr(xs[i])
  }
  return ys
}

func asStmt(x duck.INode) duck.IStmtNode {
  return x.(duck.IStmtNode)
}

func asStmts(xs []duck.INode) []duck.IStmtNode {
  ys := make([]duck.IStmtNode, len(xs))
  for i := range xs {
    ys[i] = asStmt(xs[i])
  }
  return ys
}

func asType(x duck.INode) duck.ITypeNode {
  return x.(duck.ITypeNode)
}

func asTypes(xs []duck.INode) []duck.ITypeNode {
  ys := make([]duck.ITypeNode, len(xs))
  for i := range xs {
    ys[i] = asType(xs[i])
  }
  return ys
}

func asSlot(x duck.INode) ast.Slot {
  return x.(ast.Slot)
}

func asSlots(xs []duck.INode) []ast.Slot {
  ys := make([]ast.Slot, len(xs))
  for i := range xs {
    ys[i] = asSlot(xs[i])
  }
  return ys
}

func asTypeDefinition(x duck.INode) duck.ITypeDefinition {
  return x.(duck.ITypeDefinition)
}

func asTypeDefinitions(xs []duck.INode) []duck.ITypeDefinition {
  ys := make([]duck.ITypeDefinition, len(xs))
  for i := range xs {
    ys[i] = asTypeDefinition(xs[i])
  }
  return ys
}

func asDeclarations(x duck.INode) ast.Declarations {
  return x.(ast.Declarations)
}
