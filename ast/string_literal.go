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
}

func NewStringLiteralNode(loc core.Location, literal string) *StringLiteralNode {
  ref := typesys.NewPointerTypeRef(typesys.NewCharTypeRef(loc))
  t := NewTypeNode(loc, ref)
  return &StringLiteralNode { "ast.StringLiteralNode", loc, t, literal, nil }
}

func (self StringLiteralNode) String() string {
  return self.Value
}

func (self *StringLiteralNode) AsExprNode() core.IExprNode {
  return self
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
  return self.TypeNode.GetType()
}

func (self *StringLiteralNode) SetType(t core.IType) {
  panic("#SetType called")
}

func (self StringLiteralNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self StringLiteralNode) IsConstant() bool {
  return true
}

func (self StringLiteralNode) IsParameter() bool {
  return false
}

func (self StringLiteralNode) IsLvalue() bool {
  return false
}

func (self StringLiteralNode) IsAssignable() bool {
  return false
}

func (self StringLiteralNode) IsLoadable() bool {
  return false
}

func (self StringLiteralNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self StringLiteralNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
