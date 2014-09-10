package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// AddressNode
type AddressNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewAddressNode(loc core.Location, expr core.IExprNode) AddressNode {
  if expr == nil { panic("expr is nil") }
  return AddressNode { "ast.AddressNode", loc, expr }
}

func (self AddressNode) String() string {
  return fmt.Sprintf("<ast.AddressNode location=%s expr=%s>", self.Location, self.Expr)
}

func (self AddressNode) IsExprNode() bool {
  return true
}

func (self AddressNode) GetLocation() core.Location {
  return self.Location
}
