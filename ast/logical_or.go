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
}

func NewLogicalOrNode(loc core.Location, left core.IExprNode, right core.IExprNode) LogicalOrNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return LogicalOrNode { "ast.LogicalOrNode", loc, left, right }
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self LogicalOrNode) IsExprNode() bool {
  return true
}

func (self LogicalOrNode) GetLocation() core.Location {
  return self.Location
}
