package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// PrefixOpNode
type PrefixOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
}

func NewPrefixOpNode(loc core.Location, operator string, expr core.IExprNode) *PrefixOpNode {
  if expr == nil { panic("expr is nil") }
  return &PrefixOpNode { "ast.PrefixOpNode", loc, operator, expr }
}

func (self PrefixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self PrefixOpNode) IsExprNode() bool {
  return true
}

func (self PrefixOpNode) GetLocation() core.Location {
  return self.Location
}
