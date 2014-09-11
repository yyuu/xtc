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
}

func NewSizeofTypeNode(loc core.Location, operand core.ITypeNode, ref core.ITypeRef) *SizeofTypeNode {
  if operand == nil { panic("operand is nil") }
  if ref == nil { panic("t is nil") }
  t := NewTypeNode(loc, ref)
  return &SizeofTypeNode { "ast.SizeofTypeNode", loc, operand, t }
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
