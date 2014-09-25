package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

// BreakNode
type BreakNode struct {
  ClassName string
  Location core.Location
}

func NewBreakNode(loc core.Location) *BreakNode {
  return &BreakNode { "ast.BreakNode", loc }
}

func (self BreakNode) String() string {
  return "(break)"
}

func (self *BreakNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self BreakNode) GetLocation() core.Location {
  return self.Location
}
