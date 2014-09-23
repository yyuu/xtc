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
  t core.IType
}

func NewPrefixOpNode(loc core.Location, operator string, expr core.IExprNode) *PrefixOpNode {
  if expr == nil { panic("expr is nil") }
  return &PrefixOpNode { "ast.PrefixOpNode", loc, operator, expr, nil }
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

func (self PrefixOpNode) GetOperator() string {
  return self.Operator
}

func (self PrefixOpNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self PrefixOpNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *PrefixOpNode) SetType(t core.IType) {
  self.t = t
}
