package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// OpAssignNode
type OpAssignNode struct {
  ClassName string
  Location core.Location
  Operator string
  Lhs core.IExprNode
  Rhs core.IExprNode
}

func NewOpAssignNode(loc core.Location, operator string, lhs core.IExprNode, rhs core.IExprNode) *OpAssignNode {
  if lhs == nil { panic("lhs is nil") }
  if rhs == nil { panic("rhs is nil") }
  return &OpAssignNode { "ast.OpAssignNode", loc, operator, lhs, rhs }
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.Lhs, self.Operator, self.Lhs, self.Rhs)
}

func (self OpAssignNode) IsExprNode() bool {
  return true
}

func (self OpAssignNode) GetLocation() core.Location {
  return self.Location
}