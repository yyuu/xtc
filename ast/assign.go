package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// AssignNode
type AssignNode struct {
  ClassName string
  Location duck.Location
  Lhs duck.IExprNode
  Rhs duck.IExprNode
}

func NewAssignNode(loc duck.Location, lhs duck.IExprNode, rhs duck.IExprNode) AssignNode {
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

func (self AssignNode) GetLocation() duck.Location {
  return self.Location
}

// OpAssignNode
type OpAssignNode struct {
  ClassName string
  Location duck.Location
  Operator string
  Lhs duck.IExprNode
  Rhs duck.IExprNode
}

func NewOpAssignNode(loc duck.Location, operator string, lhs duck.IExprNode, rhs duck.IExprNode) OpAssignNode {
  if lhs == nil { panic("lhs is nil") }
  if rhs == nil { panic("rhs is nil") }
  return OpAssignNode { "ast.OpAssignNode", loc, operator, lhs, rhs }
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.Lhs, self.Operator, self.Lhs, self.Rhs)
}

func (self OpAssignNode) IsExprNode() bool {
  return true
}

func (self OpAssignNode) GetLocation() duck.Location {
  return self.Location
}
