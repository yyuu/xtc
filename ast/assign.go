package ast

import (
  "fmt"
)

// AssignNode
type AssignNode struct {
  Location Location
  Lhs IExprNode
  Rhs IExprNode
}

func NewAssignNode(location Location, lhs IExprNode, rhs IExprNode) AssignNode {
  return AssignNode { location, lhs, rhs }
}

func (self AssignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Lhs, self.Rhs)
}

func (self AssignNode) IsExpr() bool {
  return true
}

func (self AssignNode) GetLocation() Location {
  return self.Location
}

// OpAssignNode
type OpAssignNode struct {
  Location Location
  Operator string
  Lhs IExprNode
  Rhs IExprNode
}

func NewOpAssignNode(location Location, operator string, lhs IExprNode, rhs IExprNode) OpAssignNode {
  return OpAssignNode { location, operator, lhs, rhs }
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.Lhs, self.Operator, self.Lhs, self.Rhs)
}

func (self OpAssignNode) IsExpr() bool {
  return true
}

func (self OpAssignNode) GetLocation() Location {
  return self.Location
}
