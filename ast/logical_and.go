package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// LogicalAndNode
type LogicalAndNode struct {
  ClassName string
  Location core.Location
  Left core.IExprNode
  Right core.IExprNode
  Type core.IType
}

func NewLogicalAndNode(loc core.Location, left core.IExprNode, right core.IExprNode) *LogicalAndNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return &LogicalAndNode { "ast.LogicalAndNode", loc, left, right, nil }
}

func (self LogicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

func (self *LogicalAndNode) AsExprNode() core.IExprNode {
  return self
}

func (self LogicalAndNode) GetLocation() core.Location {
  return self.Location
}

func (self LogicalAndNode) GetOperator() string {
  return "&&"
}

func (self LogicalAndNode) GetLeft() core.IExprNode {
  return self.Left
}

func (self *LogicalAndNode) SetLeft(expr core.IExprNode) {
  self.Left = expr
}

func (self LogicalAndNode) GetRight() core.IExprNode {
  return self.Right
}

func (self *LogicalAndNode) SetRight(expr core.IExprNode) {
  self.Right = expr
}

func (self LogicalAndNode) GetType() core.IType {
  if self.Type == nil {
    self.Type = self.Left.GetType()
  }
  return self.Type
}

func (self *LogicalAndNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self LogicalAndNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self LogicalAndNode) IsConstant() bool {
  return false
}

func (self LogicalAndNode) IsParameter() bool {
  return false
}

func (self LogicalAndNode) IsLvalue() bool {
  return false
}

func (self LogicalAndNode) IsAssignable() bool {
  return false
}

func (self LogicalAndNode) IsLoadable() bool {
  return false
}

func (self LogicalAndNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self LogicalAndNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
