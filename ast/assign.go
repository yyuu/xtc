package ast

import (
  "fmt"
)

// AssignNode
type assignNode struct {
  Location ILocation
  Lhs IExprNode
  Rhs IExprNode
}

func AssignNode(lhs IExprNode, rhs IExprNode) assignNode {
  return assignNode { lhs, rhs }
}

func (self assignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Lhs, self.Rhs)
}

func (self assignNode) IsExpr() bool {
  return true
}

// OpAssignNode
type opAssignNode struct {
  Location ILocation
  Operator string
  Lhs IExprNode
  Rhs IExprNode
}

func OpAssignNode(operator string, lhs IExprNode, rhs IExprNode) opAssignNode {
  return opAssignNode { operator, lhs, rhs }
}

func (self opAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.Lhs, self.Operator, self.Lhs, self.Rhs)
}

func (self opAssignNode) IsExpr() bool {
  return true
}