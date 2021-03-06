package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// PrefixOpNode
type PrefixOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
  Amount int
  Type core.IType
}

func NewPrefixOpNode(loc core.Location, operator string, expr core.IExprNode) *PrefixOpNode {
  if expr == nil { panic("expr is nil") }
  return &PrefixOpNode { "ast.PrefixOpNode", loc, operator, expr, 1, nil }
}

func (self PrefixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self *PrefixOpNode) AsExprNode() core.IExprNode {
  return self
}

func (self PrefixOpNode) GetLocation() core.Location {
  return self.Location
}

func (self *PrefixOpNode) GetOperator() string {
  return self.Operator
}

func (self *PrefixOpNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self *PrefixOpNode) SetExpr(expr core.IExprNode) {
  self.Expr = expr
}

func (self *PrefixOpNode) GetAmount() int {
  return self.Amount
}

func (self *PrefixOpNode) SetAmount(i int) {
  self.Amount = i
}

func (self *PrefixOpNode) GetOpType() core.IType {
  return self.Type
}

func (self *PrefixOpNode) SetOpType(t core.IType) {
  self.Type = t
}

func (self *PrefixOpNode) GetType() core.IType {
  return self.Expr.GetType()
}

func (self *PrefixOpNode) SetType(t core.IType) {
  panic("#SetType called")
}

func (self *PrefixOpNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self *PrefixOpNode) IsConstant() bool {
  return false
}

func (self *PrefixOpNode) IsParameter() bool {
  return false
}

func (self *PrefixOpNode) IsLvalue() bool {
  return false
}

func (self *PrefixOpNode) IsAssignable() bool {
  return false
}

func (self *PrefixOpNode) IsLoadable() bool {
  return false
}

func (self *PrefixOpNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *PrefixOpNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
