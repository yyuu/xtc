package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// AssignNode
type AssignNode struct {
  ClassName string
  Location core.Location
  Lhs core.IExprNode
  Rhs core.IExprNode
}

func NewAssignNode(loc core.Location, lhs core.IExprNode, rhs core.IExprNode) AssignNode {
  if lhs == nil { panic("lhs is nil") }
  if rhs == nil { panic("rhs is nil") }
  return AssignNode { "ast.AssignNode", loc, lhs, rhs }
}

func (self AssignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Lhs, self.Rhs)
}

func (self AssignNode) IsExprNode() bool {
  return true
}

func (self AssignNode) GetLocation() core.Location {
  return self.Location
}

func (self AssignNode) GetLhs() core.IExprNode {
  return self.Lhs
}

func (self AssignNode) GetRhs() core.IExprNode {
  return self.Rhs
}
