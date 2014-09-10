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
}

func NewLogicalAndNode(loc core.Location, left core.IExprNode, right core.IExprNode) *LogicalAndNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return &LogicalAndNode { "ast.LogicalAndNode", loc, left, right }
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
