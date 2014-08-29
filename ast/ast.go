package ast

import (
  "fmt"
  "strings"
)

type INode interface {
  String() string
  GetLocation() Location
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

type Location struct {
  SourceName string
  LineNumber int
  LineOffset int
}

type AST struct {
  Stmts []IStmtNode
}

func (self AST) String() string {
  xs := make([]string, len(self.Stmts))
  for i := range self.Stmts {
    stmt := self.Stmts[i]
    location := stmt.GetLocation()
    xs[i] = fmt.Sprintf(";; %s:%d,%d\n%s", location.SourceName, location.LineNumber+1, location.LineOffset+1, stmt)
  }
  return strings.Join(xs, "\n")
}
