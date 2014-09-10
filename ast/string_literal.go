package ast

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

// StringLiteralNode
type StringLiteralNode struct {
  ClassName string
  Location core.Location
  Value string
  entry *entity.ConstantEntry
}

func NewStringLiteralNode(loc core.Location, literal string) *StringLiteralNode {
  return &StringLiteralNode { "ast.StringLiteralNode", loc, literal, nil }
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
