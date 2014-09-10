package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

// StringLiteralNode
type StringLiteralNode struct {
  ClassName string
  Location core.Location
  Value string
  entry core.IConstantEntry
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

func (self StringLiteralNode) GetEntry() core.IConstantEntry {
  return self.entry
}

func (self *StringLiteralNode) SetEntry(e core.IConstantEntry) {
  self.entry = e
}
