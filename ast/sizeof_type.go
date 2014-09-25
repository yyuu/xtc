package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// SizeofTypeNode
type SizeofTypeNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  OperandTypeNode core.ITypeNode
  t core.IType
}

func NewSizeofTypeNode(loc core.Location, operand core.ITypeNode, ref core.ITypeRef) *SizeofTypeNode {
  if operand == nil { panic("operand is nil") }
  if ref == nil { panic("t is nil") }
  t := NewTypeNode(loc, ref)
  return &SizeofTypeNode { "ast.SizeofTypeNode", loc, operand, t, nil }
}

func (self SizeofTypeNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.TypeNode)
}

func (self *SizeofTypeNode) AsExprNode() core.IExprNode {
  return self
}

func (self SizeofTypeNode) GetLocation() core.Location {
  return self.Location
}

func (self SizeofTypeNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self SizeofTypeNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self SizeofTypeNode) GetOperandTypeNode() core.ITypeNode {
  return self.OperandTypeNode
}

func (self SizeofTypeNode) GetOperandTypeRef() core.ITypeRef {
  return self.OperandTypeNode.GetTypeRef()
}

func (self SizeofTypeNode) GetType() core.IType {
  if self.t == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.t
}

func (self *SizeofTypeNode) SetType(t core.IType) {
  self.t = t
}

func (self SizeofTypeNode) IsConstant() bool {
  return false
}

func (self SizeofTypeNode) IsParameter() bool {
  return false
}

func (self SizeofTypeNode) IsLvalue() bool {
  return false
}

func (self SizeofTypeNode) IsAssignable() bool {
  return false
}

func (self SizeofTypeNode) IsLoadable() bool {
  return false
}

func (self SizeofTypeNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self SizeofTypeNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
