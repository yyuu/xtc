package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// SuffixOpNode
type SuffixOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
  t core.IType
}

func NewSuffixOpNode(loc core.Location, operator string, expr core.IExprNode) *SuffixOpNode {
  if expr == nil { panic("expr is nil") }
  return &SuffixOpNode { "ast.SuffixOpNode", loc, operator, expr, nil }
}

func (self SuffixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self SuffixOpNode) IsExprNode() bool {
  return true
}

func (self SuffixOpNode) GetLocation() core.Location {
  return self.Location
}

func (self SuffixOpNode) GetOperator() string {
  return self.Operator
}

func (self SuffixOpNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self SuffixOpNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *SuffixOpNode) SetType(t core.IType) {
  self.t = t
}
