package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// DereferenceNode
type DereferenceNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Type core.IType
}

func NewDereferenceNode(loc core.Location, expr core.IExprNode) *DereferenceNode {
  if expr == nil { panic("expr is nil") }
  return &DereferenceNode { "ast.DereferenceNode", loc, expr, nil }
}

func (self DereferenceNode) String() string {
  panic("not implemented")
}

func (self *DereferenceNode) AsExprNode() core.IExprNode {
  return self
}

func (self DereferenceNode) GetLocation() core.Location {
  return self.Location
}

func (self DereferenceNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self DereferenceNode) GetType() core.IType {
  if self.Type == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.Type
}

func (self *DereferenceNode) SetType(t core.IType) {
  self.Type = t
}

func (self DereferenceNode) GetOrigType() core.IType {
  return self.GetType().GetBaseType()
}

func (self DereferenceNode) IsConstant() bool {
  return false
}

func (self DereferenceNode) IsParameter() bool {
  return false
}

func (self DereferenceNode) IsLvalue() bool {
  return true
}

func (self DereferenceNode) IsAssignable() bool {
  return true
}

func (self DereferenceNode) IsLoadable() bool {
  return false
}

func (self DereferenceNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self DereferenceNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
