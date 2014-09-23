package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

// DereferenceNode
type DereferenceNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  t core.IType
}

func NewDereferenceNode(loc core.Location, expr core.IExprNode) *DereferenceNode {
  if expr == nil { panic("expr is nil") }
  return &DereferenceNode { "ast.DereferenceNode", loc, expr, nil }
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

func (self DereferenceNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self DereferenceNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *DereferenceNode) SetType(t core.IType) {
  self.t = t
}
