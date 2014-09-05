package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// AssignNode
type AssignNode struct {
  location duck.ILocation
  lhs duck.IExprNode
  rhs duck.IExprNode
}

func NewAssignNode(loc duck.ILocation, lhs duck.IExprNode, rhs duck.IExprNode) AssignNode {
  return AssignNode { loc, lhs, rhs }
}

func (self AssignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.lhs, self.rhs)
}

func (self AssignNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Lhs duck.IExprNode
    Rhs duck.IExprNode
  }
  x.ClassName = "ast.AssignNode"
  x.Location = self.location
  x.Lhs = self.lhs
  x.Rhs = self.rhs
  return json.Marshal(x)
}

func (self AssignNode) IsExpr() bool {
  return true
}

func (self AssignNode) GetLocation() duck.ILocation {
  return self.location
}

// OpAssignNode
type OpAssignNode struct {
  location duck.ILocation
  operator string
  lhs duck.IExprNode
  rhs duck.IExprNode
}

func NewOpAssignNode(loc duck.ILocation, operator string, lhs duck.IExprNode, rhs duck.IExprNode) OpAssignNode {
  return OpAssignNode { loc, operator, lhs, rhs }
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.lhs, self.operator, self.lhs, self.rhs)
}

func (self OpAssignNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Operator string
    Lhs duck.IExprNode
    Rhs duck.IExprNode
  }
  x.ClassName = "ast.OpAssignNode"
  x.Location = self.location
  x.Operator = self.operator
  x.Lhs = self.lhs
  x.Rhs = self.rhs
  return json.Marshal(x)
}

func (self OpAssignNode) IsExpr() bool {
  return true
}

func (self OpAssignNode) GetLocation() duck.ILocation {
  return self.location
}
