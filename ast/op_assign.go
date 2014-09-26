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
  LHS core.IExprNode
  RHS core.IExprNode
  Type core.IType
}

func NewOpAssignNode(loc core.Location, operator string, lhs core.IExprNode, rhs core.IExprNode) *OpAssignNode {
  if lhs == nil { panic("lhs is nil") }
  if rhs == nil { panic("rhs is nil") }
  return &OpAssignNode { "ast.OpAssignNode", loc, operator, lhs, rhs, nil }
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(%s (%s %s %s))", self.LHS, self.Operator, self.LHS, self.RHS)
}

func (self *OpAssignNode) AsExprNode() core.IExprNode {
  return self
}

func (self OpAssignNode) GetLocation() core.Location {
  return self.Location
}

func (self *OpAssignNode) GetOperator() string {
  return self.Operator
}

func (self *OpAssignNode) GetLHS() core.IExprNode {
  return self.LHS
}

func (self *OpAssignNode) SetLHS(expr core.IExprNode) {
  self.LHS = expr
}

func (self *OpAssignNode) GetRHS() core.IExprNode {
  return self.RHS
}

func (self *OpAssignNode) SetRHS(expr core.IExprNode) {
  self.RHS = expr
}

func (self *OpAssignNode) GetType() core.IType {
  if self.Type == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.Type
}

func (self *OpAssignNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self *OpAssignNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self *OpAssignNode) IsConstant() bool {
  return false
}

func (self *OpAssignNode) IsParameter() bool {
  return false
}

func (self *OpAssignNode) IsLvalue() bool {
  return false
}

func (self *OpAssignNode) IsAssignable() bool {
  return false
}

func (self *OpAssignNode) IsLoadable() bool {
  return false
}

func (self *OpAssignNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *OpAssignNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
