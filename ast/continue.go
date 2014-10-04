package ast

import (
  "bitbucket.org/yyuu/xtc/core"
)

// ContinueNode
type ContinueNode struct {
  ClassName string
  Location core.Location
}

func NewContinueNode(loc core.Location) *ContinueNode {
  return &ContinueNode { "ast.ContinueNode", loc }
}

func (self ContinueNode) String() string {
  return "(continue)"
}

func (self *ContinueNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self ContinueNode) GetLocation() core.Location {
  return self.Location
}

func (self ContinueNode) GetScope() core.IScope {
  panic("#GetScope called")
}
