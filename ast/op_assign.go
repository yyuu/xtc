package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// OpAssignNode
type OpAssignNode struct {
  ClassName string
  Location core.Location
  Operator string
  Lhs core.IExprNode
  Rhs core.IExprNode
  t core.IType
}

func NewOpAssignNode(loc core.Location, operator string, lhs core.IExprNode, rhs core.IExprNode) *OpAssignNode {
  if lhs == nil { panic("lhs is nil") }
  if rhs == nil { panic("rhs is nil") }
  return &OpAssignNode { "ast.OpAssignNode", loc, operator, lhs, rhs, nil }
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.Lhs, self.Operator, self.Lhs, self.Rhs)
}

func (self OpAssignNode) IsExprNode() bool {
  return true
}

func (self OpAssignNode) GetLocation() core.Location {
  return self.Location
}

func (self OpAssignNode) GetOperator() string {
  return self.Operator
}

func (self OpAssignNode) GetLhs() core.IExprNode {
  return self.Lhs
}

func (self *OpAssignNode) SetLhs(expr core.IExprNode) {
  self.Lhs = expr
}

func (self OpAssignNode) GetRhs() core.IExprNode {
  return self.Rhs
}

func (self *OpAssignNode) SetRhs(expr core.IExprNode) {
  self.Rhs = expr
}

func (self OpAssignNode) GetType() core.IType {
  if self.t == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.t
}

func (self *OpAssignNode) SetType(t core.IType) {
  self.t = t
}

func (self OpAssignNode) IsConstant() bool {
  return false
}

func (self OpAssignNode) IsParameter() bool {
  return false
}

func (self OpAssignNode) IsLvalue() bool {
  return false
}

func (self OpAssignNode) IsAssignable() bool {
  return false
}

func (self OpAssignNode) IsLoadable() bool {
  return false
}

func (self OpAssignNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self OpAssignNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
