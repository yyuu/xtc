package parser

import (
  "bitbucket.org/yyuu/bs/ast"
)

func asExpr(x ast.INode) ast.IExprNode {
  return x.(ast.IExprNode)
}

func asExprs(xs []ast.INode) []ast.IExprNode {
  ys := make([]ast.IExprNode, len(xs))
  for i := range xs {
    ys[i] = asExpr(xs[i])
  }
  return ys
}

func asStmt(x ast.INode) ast.IStmtNode {
  return x.(ast.IStmtNode)
}

func asStmts(xs []ast.INode) []ast.IStmtNode {
  ys := make([]ast.IStmtNode, len(xs))
  for i := range xs {
    ys[i] = asStmt(xs[i])
  }
  return ys
}

func asType(x ast.INode) ast.ITypeNode {
  return x.(ast.ITypeNode)
}

func asTypes(xs []ast.INode) []ast.ITypeNode {
  ys := make([]ast.ITypeNode, len(xs))
  for i := range xs {
    ys[i] = asType(xs[i])
  }
  return ys
}

func asSlot(x ast.INode) ast.Slot {
  return x.(ast.Slot)
}

func asSlots(xs []ast.INode) []ast.Slot {
  ys := make([]ast.Slot, len(xs))
  for i := range xs {
    ys[i] = asSlot(xs[i])
  }
  return ys
}

func asTypeDefinition(x ast.INode) ast.ITypeDefinition {
  return x.(ast.ITypeDefinition)
}

func asTypeDefinitions(xs []ast.INode) []ast.ITypeDefinition {
  ys := make([]ast.ITypeDefinition, len(xs))
  for i := range xs {
    ys[i] = asTypeDefinition(xs[i])
  }
  return ys
}

