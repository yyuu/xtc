package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

type INodeVisitor interface {
  VisitNode(core.INode) interface{}
}

func VisitNode(v INodeVisitor, node core.INode) interface{} {
  return v.VisitNode(node)
}

func VisitNodes(v INodeVisitor, nodes []core.INode) interface{} {
  var x interface{}
  for i := range nodes {
    x = VisitNode(v, nodes[i])
  }
  return x
}

func VisitStmt(v INodeVisitor, stmt core.IStmtNode) interface{} {
  return VisitNode(v, stmt.(core.INode))
}

func VisitStmts(v INodeVisitor, stmts []core.IStmtNode) interface{} {
  var x interface{}
  for i := range stmts {
    x = VisitStmt(v, stmts[i])
  }
  return x
}

func VisitExpr(v INodeVisitor, expr core.IExprNode) interface{} {
  return VisitNode(v, expr.(core.INode))
}

func VisitExprs(v INodeVisitor, exprs []core.IExprNode) interface{} {
  var x interface{}
  for i := range exprs {
    x = VisitExpr(v, exprs[i])
  }
  return x
}
