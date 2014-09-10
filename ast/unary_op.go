package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// UnaryOpNode
type UnaryOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
}

func NewUnaryOpNode(loc core.Location, operator string, expr core.IExprNode) *UnaryOpNode {
  if expr == nil { panic("expr is nil") }
  return &UnaryOpNode { "ast.UnaryOpNode", loc, operator, expr }
}

func (self UnaryOpNode) String() string {
  switch self.Operator {
    case "!": return fmt.Sprintf("(not %s)", self.Expr)
    default:  return fmt.Sprintf("%s%s", self.Operator, self.Expr)
  }
}

func (self UnaryOpNode) IsExprNode() bool {
  return true
}

func (self UnaryOpNode) GetLocation() core.Location {
  return self.Location
}
