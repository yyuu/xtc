package ast

import (
  "fmt"
)

type INode interface {
  String() string
  GetLocation() ILocation
}

type IExprNode interface {
  INode
  IsExpr() bool
}

type IStmtNode interface {
  INode
  IsStmt() bool
}

type ITypeNode interface {
  INode
  IsType() bool
}

type ILocation interface {
  GetSourceName() string
  GetLineNumber() int
  GetLineOffset() int
}

type AST struct {
  Stmts []IStmtNode
}

func (self AST) String() string {
  s := ""
  for i := range self.Stmts {
    s += fmt.Sprintf("%s\n", self.Stmts[i])
  }
  return s
}
