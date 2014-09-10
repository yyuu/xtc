package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// GotoNode
type GotoNode struct {
  ClassName string
  Location core.Location
  Target string
}

func NewGotoNode(loc core.Location, target string) *GotoNode {
  return &GotoNode { "ast.GotoNode", loc, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self GotoNode) IsStmtNode() bool {
  return true
}

func (self GotoNode) GetLocation() core.Location {
  return self.Location
}
