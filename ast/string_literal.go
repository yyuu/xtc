package ast

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

// StringLiteralNode
type StringLiteralNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Value string
  entry *entity.ConstantEntry
  t core.IType
}

func NewStringLiteralNode(loc core.Location, literal string) *StringLiteralNode {
  ref := typesys.NewPointerTypeRef(typesys.NewCharTypeRef(loc))
  t := NewTypeNode(loc, ref)
  return &StringLiteralNode { "ast.StringLiteralNode", loc, t, literal, nil, nil }
}

func (self StringLiteralNode) String() string {
  return self.Value
}

func (self StringLiteralNode) IsExprNode() bool {
  return true
}

func (self StringLiteralNode) GetLocation() core.Location {
  return self.Location
}

func (self StringLiteralNode) GetValue() string {
  return self.Value
}

func (self StringLiteralNode) GetEntry() *entity.ConstantEntry {
  return self.entry
}

func (self *StringLiteralNode) SetEntry(e *entity.ConstantEntry) {
  self.entry = e
}

func (self StringLiteralNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self StringLiteralNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self StringLiteralNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *StringLiteralNode) SetType(t core.IType) {
  self.t = t
}
