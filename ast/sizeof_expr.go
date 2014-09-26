package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// SizeofExprNode
type SizeofExprNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  TypeNode core.ITypeNode
  Type core.IType
}

func NewSizeofExprNode(loc core.Location, expr core.IExprNode, t core.ITypeRef) *SizeofExprNode {
  if expr == nil { panic("expr is nil") }
  if t == nil { panic("t is nil") }
  return &SizeofExprNode { "ast.SizeofExprNode", loc, expr, NewTypeNode(loc, t), nil }
}

func (self SizeofExprNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.Expr)
}

func (self *SizeofExprNode) AsExprNode() core.IExprNode {
  return self
}

func (self SizeofExprNode) GetLocation() core.Location {
  return self.Location
}

func (self *SizeofExprNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *SizeofExprNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *SizeofExprNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self *SizeofExprNode) GetType() core.IType {
  if self.Type == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.Type
}

func (self *SizeofExprNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self *SizeofExprNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self *SizeofExprNode) IsConstant() bool {
  return false
}

func (self *SizeofExprNode) IsParameter() bool {
  return false
}

func (self *SizeofExprNode) IsLvalue() bool {
  return false
}

func (self *SizeofExprNode) IsAssignable() bool {
  return false
}

func (self *SizeofExprNode) IsLoadable() bool {
  return false
}

func (self *SizeofExprNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *SizeofExprNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
