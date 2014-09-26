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
  Type core.IType
}

func NewUnaryOpNode(loc core.Location, operator string, expr core.IExprNode) *UnaryOpNode {
  if expr == nil { panic("expr is nil") }
  return &UnaryOpNode { "ast.UnaryOpNode", loc, operator, expr, nil }
}

func (self UnaryOpNode) String() string {
  switch self.Operator {
    case "!": return fmt.Sprintf("(not %s)", self.Expr)
    default:  return fmt.Sprintf("%s%s", self.Operator, self.Expr)
  }
}

func (self *UnaryOpNode) AsExprNode() core.IExprNode {
  return self
}

func (self UnaryOpNode) GetLocation() core.Location {
  return self.Location
}

func (self UnaryOpNode) GetOperator() string {
  return self.Operator
}

func (self UnaryOpNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self UnaryOpNode) GetOpType() core.IType {
  return self.Type
}

func (self *UnaryOpNode) SetOpType(t core.IType) {
  self.Type = t
}

func (self UnaryOpNode) GetType() core.IType {
  return self.Expr.GetType()
}

func (self *UnaryOpNode) SetType(t core.IType) {
  panic("#SetType called")
}

func (self UnaryOpNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self UnaryOpNode) IsConstant() bool {
  return false
}

func (self UnaryOpNode) IsParameter() bool {
  return false
}

func (self UnaryOpNode) IsLvalue() bool {
  return false
}

func (self UnaryOpNode) IsAssignable() bool {
  return false
}

func (self UnaryOpNode) IsLoadable() bool {
  return false
}

func (self UnaryOpNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self UnaryOpNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
