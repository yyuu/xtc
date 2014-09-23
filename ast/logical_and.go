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
  t core.IType
}

func NewLogicalAndNode(loc core.Location, left core.IExprNode, right core.IExprNode) *LogicalAndNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return &LogicalAndNode { "ast.LogicalAndNode", loc, left, right, nil }
}

func (self LogicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

func (self LogicalAndNode) IsExprNode() bool {
  return true
}

func (self LogicalAndNode) GetLocation() core.Location {
  return self.Location
}

func (self LogicalAndNode) GetLeft() core.IExprNode {
  return self.Left
}

func (self LogicalAndNode) GetRight() core.IExprNode {
  return self.Right
}

func (self LogicalAndNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *LogicalAndNode) SetType(t core.IType) {
  self.t = t
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
