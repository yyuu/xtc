package ast

import (
  "bitbucket.org/yyuu/bs/core"
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

func (self ContinueNode) IsStmtNode() bool {
  return true
}

func (self ContinueNode) GetLocation() core.Location {
  return self.Location
}
