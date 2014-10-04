package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// AssignNode
type AssignNode struct {
  ClassName string
  Location core.Location
  LHS core.IExprNode
  RHS core.IExprNode
  Type core.IType
}

func NewAssignNode(loc core.Location, lhs core.IExprNode, rhs core.IExprNode) *AssignNode {
  if lhs == nil { panic("lhs is nil") }
  if rhs == nil { panic("rhs is nil") }
  return &AssignNode { "ast.AssignNode", loc, lhs, rhs, nil }
}

func (self AssignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.LHS, self.RHS)
}

func (self *AssignNode) AsExprNode() core.IExprNode {
  return self
}

func (self AssignNode) GetLocation() core.Location {
  return self.Location
}

func (self *AssignNode) GetLHS() core.IExprNode {
  return self.LHS
}

func (self *AssignNode) SetLHS(expr core.IExprNode) {
  self.LHS = expr
}

func (self *AssignNode) GetRHS() core.IExprNode {
  return self.RHS
}

func (self *AssignNode) SetRHS(expr core.IExprNode) {
  self.RHS = expr
}

func (self *AssignNode) GetType() core.IType {
  if self.Type == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.Type
}

func (self *AssignNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self *AssignNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self *AssignNode) IsConstant() bool {
  return false
}

func (self *AssignNode) IsParameter() bool {
  return false
}

func (self *AssignNode) IsLvalue() bool {
  return false
}

func (self *AssignNode) IsAssignable() bool {
  return false
}

func (self *AssignNode) IsLoadable() bool {
  return false
}

func (self *AssignNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *AssignNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
