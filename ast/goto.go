package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
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

func (self *GotoNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self GotoNode) GetLocation() core.Location {
  return self.Location
}

func (self *GotoNode) GetTarget() string {
  return self.Target
}

func (self GotoNode) GetScope() core.IScope {
  panic("#GetScope called")
}
