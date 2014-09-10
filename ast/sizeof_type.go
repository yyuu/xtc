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
  Operand core.ITypeNode
}

func NewSizeofTypeNode(loc core.Location, operand core.ITypeNode, t core.ITypeRef) *SizeofTypeNode {
  if operand == nil { panic("operand is nil") }
  if t == nil { panic("t is nil") }
  return &SizeofTypeNode { "ast.SizeofTypeNode", loc, operand, NewTypeNode(loc, t) }
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
