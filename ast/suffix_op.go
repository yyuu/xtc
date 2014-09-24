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
  amount int
  t core.IType
}

func NewSuffixOpNode(loc core.Location, operator string, expr core.IExprNode) *SuffixOpNode {
  if expr == nil { panic("expr is nil") }
  return &SuffixOpNode { "ast.SuffixOpNode", loc, operator, expr, 1, nil }
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

func (self *SuffixOpNode) SetExpr(expr core.IExprNode) {
  self.Expr = expr
}

func (self SuffixOpNode) GetAmount() int {
  return self.amount
}

func (self *SuffixOpNode) SetAmount(i int) {
  self.amount = i
}

func (self SuffixOpNode) GetOpType() core.IType {
  return self.t
}

func (self *SuffixOpNode) SetOpType(t core.IType) {
  self.t = t
}

func (self SuffixOpNode) GetType() core.IType {
  return self.Expr.GetType()
}

func (self *SuffixOpNode) SetType(t core.IType) {
  panic("SuffixOpNode#SetType called")
}

func (self SuffixOpNode) IsConstant() bool {
  return false
}

func (self SuffixOpNode) IsParameter() bool {
  return false
}

func (self SuffixOpNode) IsLvalue() bool {
  return false
}

func (self SuffixOpNode) IsAssignable() bool {
  return false
}

func (self SuffixOpNode) IsLoadable() bool {
  return false
}

func (self SuffixOpNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self SuffixOpNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
