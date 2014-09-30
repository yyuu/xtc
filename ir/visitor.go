package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type IIRVisitor interface {
  VisitStmt(core.IStmt) interface{}
  VisitExpr(core.IExpr) interface{}
}

func VisitStmt(v IIRVisitor, stmt core.IStmt) interface{} {
  return v.VisitStmt(stmt)
}

func VisitExpr(v IIRVisitor, expr core.IExpr) interface{} {
  return v.VisitExpr(expr)
}
