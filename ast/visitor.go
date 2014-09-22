package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

type INodeVisitor interface {
  VisitNode(core.INode)
}

func VisitNode(v INodeVisitor, node core.INode) {
  v.VisitNode(node)
}

func VisitNodes(v INodeVisitor, nodes []core.INode) {
  for i := range nodes {
    VisitNode(v, nodes[i])
  }
}

func VisitStmt(v INodeVisitor, stmt core.IStmtNode) {
  VisitNode(v, stmt.(core.INode))
}

func VisitStmts(v INodeVisitor, stmts []core.IStmtNode) {
  for i := range stmts {
    VisitStmt(v, stmts[i])
  }
}

func VisitExpr(v INodeVisitor, expr core.IExprNode) {
  VisitNode(v, expr.(core.INode))
}

func VisitExprs(v INodeVisitor, exprs []core.IExprNode) {
  for i := range exprs {
    VisitExpr(v, exprs[i])
  }
}
