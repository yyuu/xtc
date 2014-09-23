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

func (self SizeofTypeNode) IsExprNode() bool {
  return true
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
    panic("type is nil")
  }
  return self.t
}

func (self *SizeofTypeNode) SetType(t core.IType) {
  self.t = t
}
