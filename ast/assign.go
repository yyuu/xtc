package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// AssignNode
type AssignNode struct {
  Location duck.ILocation
  Lhs duck.IExprNode
  Rhs duck.IExprNode
}

func NewAssignNode(location duck.ILocation, lhs duck.IExprNode, rhs duck.IExprNode) AssignNode {
  return AssignNode { location, lhs, rhs }
}

func (self AssignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Lhs, self.Rhs)
}

func (self AssignNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Lhs duck.IExprNode
    Rhs duck.IExprNode
  }
  x.ClassName = "ast.AssignNode"
  x.Location = self.Location
  x.Lhs = self.Lhs
  x.Rhs = self.Rhs
  return json.Marshal(x)
}

func (self AssignNode) IsExpr() bool {
  return true
}

func (self AssignNode) GetLocation() duck.ILocation {
  return self.Location
}

// OpAssignNode
type OpAssignNode struct {
  Location duck.ILocation
  Operator string
  Lhs duck.IExprNode
  Rhs duck.IExprNode
}

func NewOpAssignNode(location duck.ILocation, operator string, lhs duck.IExprNode, rhs duck.IExprNode) OpAssignNode {
  return OpAssignNode { location, operator, lhs, rhs }
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.Lhs, self.Operator, self.Lhs, self.Rhs)
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
  x.Location = self.Location
  x.Operator = self.Operator
  x.Lhs = self.Lhs
  x.Rhs = self.Rhs
  return json.Marshal(x)
}

func (self OpAssignNode) IsExpr() bool {
  return true
}

func (self OpAssignNode) GetLocation() duck.ILocation {
  return self.Location
}
