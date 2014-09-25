package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// LogicalOrNode
type LogicalOrNode struct {
  ClassName string
  Location core.Location
  Left core.IExprNode
  Right core.IExprNode
  Type core.IType
}

func NewLogicalOrNode(loc core.Location, left core.IExprNode, right core.IExprNode) *LogicalOrNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return &LogicalOrNode { "ast.LogicalOrNode", loc, left, right, nil }
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self *LogicalOrNode) AsExprNode() core.IExprNode {
  return self
}

func (self LogicalOrNode) GetLocation() core.Location {
  return self.Location
}

func (self LogicalOrNode) GetOperator() string {
  return "||"
}

func (self LogicalOrNode) GetLeft() core.IExprNode {
  return self.Left
}

func (self *LogicalOrNode) SetLeft(expr core.IExprNode) {
  self.Left = expr
}

func (self LogicalOrNode) GetRight() core.IExprNode {
  return self.Right
}

func (self *LogicalOrNode) SetRight(expr core.IExprNode) {
  self.Right = expr
}

func (self LogicalOrNode) GetType() core.IType {
  if self.Type == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.Type
}

func (self *LogicalOrNode) SetType(t core.IType) {
  self.Type = t
}

func (self LogicalOrNode) IsConstant() bool {
  return false
}

func (self LogicalOrNode) IsParameter() bool {
  return false
}

func (self LogicalOrNode) IsLvalue() bool {
  return false
}

func (self LogicalOrNode) IsAssignable() bool {
  return false
}

func (self LogicalOrNode) IsLoadable() bool {
  return false
}

func (self LogicalOrNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self LogicalOrNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
