package ast

import (
  "fmt"
)

// AssignNode
type assignNode struct {
  location ILocation
  Lhs IExprNode
  Rhs IExprNode
}

func AssignNode(location ILocation, lhs IExprNode, rhs IExprNode) assignNode {
  return assignNode { location, lhs, rhs }
}

func (self assignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Lhs, self.Rhs)
}

func (self assignNode) IsExpr() bool {
  return true
}

func (self assignNode) GetLocation() ILocation {
  return self.location
}

// OpAssignNode
type opAssignNode struct {
  location ILocation
  Operator string
  Lhs IExprNode
  Rhs IExprNode
}

func OpAssignNode(location ILocation, operator string, lhs IExprNode, rhs IExprNode) opAssignNode {
  return opAssignNode { location, operator, lhs, rhs }
}

func (self opAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.Lhs, self.Operator, self.Lhs, self.Rhs)
}

func (self opAssignNode) IsExpr() bool {
  return true
}

func (self opAssignNode) GetLocation() ILocation {
  return self.location
}
