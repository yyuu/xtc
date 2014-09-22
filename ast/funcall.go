package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

// FuncallNode
type FuncallNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Args []core.IExprNode
}

func NewFuncallNode(loc core.Location, expr core.IExprNode, args []core.IExprNode) *FuncallNode {
  if expr == nil { panic("expr is nil") }
  return &FuncallNode { "ast.FuncallNode", loc, expr, args }
}

func (self FuncallNode) String() string {
  sArgs := make([]string, len(self.Args))
  for i := range self.Args {
    sArgs[i] = fmt.Sprintf("%s", self.Args[i])
  }
  if len(sArgs) == 0 {
    return fmt.Sprintf("(%s)", self.Expr)
  } else {
    return fmt.Sprintf("(%s %s)", self.Expr, strings.Join(sArgs, " "))
  }
}

func (self FuncallNode) IsExprNode() bool {
  return true
}

func (self FuncallNode) GetLocation() core.Location {
  return self.Location
}

func (self FuncallNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self FuncallNode) GetArgs() []core.IExprNode {
  return self.Args
}
