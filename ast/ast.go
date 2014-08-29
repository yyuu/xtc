package ast

import (
  "fmt"
  "strings"
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
  xs := make([]string, len(self.Stmts))
  for i := range self.Stmts {
    stmt := self.Stmts[i]
    location := stmt.GetLocation()
    xs[i] = fmt.Sprintf(";; %s:%d,%d\n%s", location.GetSourceName(), location.GetLineNumber()+1, location.GetLineOffset()+1, stmt)
  }
  return strings.Join(xs, "\n")
}
