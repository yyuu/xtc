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
  t core.IType
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

func (self UnaryOpNode) IsExprNode() bool {
  return true
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

func (self UnaryOpNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *UnaryOpNode) SetType(t core.IType) {
  self.t = t
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
