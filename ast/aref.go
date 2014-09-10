package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// ArefNode
type ArefNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Index core.IExprNode
}

func NewArefNode(loc core.Location, expr core.IExprNode, index core.IExprNode) ArefNode {
  if expr == nil { panic("expr is nil") }
  if index == nil { panic("index is nil") }
  return ArefNode { "ast.ArefNode", loc, expr, index }
}

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self ArefNode) IsExprNode() bool {
  return true
}

func (self ArefNode) GetLocation() core.Location {
  return self.Location
}
