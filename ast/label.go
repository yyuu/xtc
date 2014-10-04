package ast

import (
  "bitbucket.org/yyuu/xtc/core"
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

func (self *LabelNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self LabelNode) GetLocation() core.Location {
  return self.Location
}

func (self *LabelNode) GetName() string {
  return self.Name
}

func (self *LabelNode) GetStmt() core.IStmtNode {
  return self.Stmt
}

func (self LabelNode) GetScope() core.IScope {
  panic("#GetScope called")
}
