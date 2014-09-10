package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

// LabelNode
type LabelNode struct {
  ClassName string
  Location core.Location
  Name string
  Stmt core.IStmtNode
}

func NewLabelNode(loc core.Location, name string, stmt core.IStmtNode) *LabelNode {
  if stmt == nil { panic("stmt is nil") }
  return &LabelNode { "ast.LabelNode", loc, name, stmt }
}

func (self LabelNode) String() string {
  panic("not implemented")
}

func (self LabelNode) IsStmtNode() bool {
  return true
}

func (self LabelNode) GetLocation() core.Location {
  return self.Location
}
