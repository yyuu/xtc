package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

// DereferenceNode
type DereferenceNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewDereferenceNode(loc core.Location, expr core.IExprNode) DereferenceNode {
  if expr == nil { panic("expr is nil") }
  return DereferenceNode { "ast.DereferenceNode", loc, expr }
}

func (self DereferenceNode) String() string {
  panic("not implemented")
}

func (self DereferenceNode) IsExprNode() bool {
  return true
}

func (self DereferenceNode) GetLocation() core.Location {
  return self.Location
}
