package ast

import (
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

func (self *DereferenceNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self *DereferenceNode) GetType() core.IType {
  if self.Type == nil {
    self.Type = self.GetOrigType()
  }
  return self.Type
}

func (self *DereferenceNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self *DereferenceNode) GetOrigType() core.IType {
  return self.Expr.GetType().GetBaseType()
}

func (self *DereferenceNode) IsConstant() bool {
  return false
}

func (self *DereferenceNode) IsParameter() bool {
  return false
}

func (self *DereferenceNode) IsLvalue() bool {
  return true
}

func (self *DereferenceNode) IsAssignable() bool {
  return true
}

func (self *DereferenceNode) IsLoadable() bool {
  t := self.GetOrigType()
  return !t.IsArray() && !t.IsFunction()
}

func (self *DereferenceNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *DereferenceNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
