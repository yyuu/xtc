package ast

import (
  "fmt"
)

type assignNode struct {
  Lhs IExprNode
  Rhs IExprNode
}

func AssignNode(lhs IExprNode, rhs IExprNode) assignNode {
  return assignNode { lhs, rhs }
}

func (self assignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Lhs, self.Rhs)
}

type opAssignNode struct {
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
