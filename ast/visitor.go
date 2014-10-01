package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

type INodeVisitor interface {
  VisitStmtNode(core.IStmtNode) interface{}
  VisitExprNode(core.IExprNode) interface{}
  VisitTypeDefinition(core.ITypeDefinition) interface{}
}

func VisitStmtNode(v INodeVisitor, stmt core.IStmtNode) interface{} {
  return v.VisitStmtNode(stmt)
}

func VisitStmtNodes(v INodeVisitor, stmts []core.IStmtNode) interface{} {
  var x interface{}
  for i := range stmts {
    x = VisitStmtNode(v, stmts[i])
  }
  return x
}

func VisitExprNode(v INodeVisitor, expr core.IExprNode) interface{} {
  return v.VisitExprNode(expr)
}

func VisitExprNodes(v INodeVisitor, exprs []core.IExprNode) interface{} {
  var x interface{}
  for i := range exprs {
    x = VisitExprNode(v, exprs[i])
  }
  return x
}

func VisitTypeDefinition(v INodeVisitor, typedef core.ITypeDefinition) interface{} {
  return v.VisitTypeDefinition(typedef)
}

func VisitTypeDefinitions(v INodeVisitor, typedefs []core.ITypeDefinition) interface{} {
  var x interface{}
  for i := range typedefs {
    x = VisitTypeDefinition(v, typedefs[i])
  }
  return x
}
