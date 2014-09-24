package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// AssignNode
type AssignNode struct {
  ClassName string
  Location core.Location
  Lhs core.IExprNode
  Rhs core.IExprNode
  t core.IType
}

func NewAssignNode(loc core.Location, lhs core.IExprNode, rhs core.IExprNode) *AssignNode {
  if lhs == nil { panic("lhs is nil") }
  if rhs == nil { panic("rhs is nil") }
  return &AssignNode { "ast.AssignNode", loc, lhs, rhs, nil }
}

func (self AssignNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Lhs, self.Rhs)
}

func (self AssignNode) IsExprNode() bool {
  return true
}

func (self AssignNode) GetLocation() core.Location {
  return self.Location
}

func (self AssignNode) GetLhs() core.IExprNode {
  return self.Lhs
}

func (self *AssignNode) SetLhs(expr core.IExprNode) {
  self.Lhs = expr
}

func (self AssignNode) GetRhs() core.IExprNode {
  return self.Rhs
}

func (self *AssignNode) SetRhs(expr core.IExprNode) {
  self.Rhs = expr
}

func (self AssignNode) GetType() core.IType {
  if self.t == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.t
}

func (self *AssignNode) SetType(t core.IType) {
  self.t = t
}

func (self AssignNode) IsConstant() bool {
  return false
}

func (self AssignNode) IsParameter() bool {
  return false
}

func (self AssignNode) IsLvalue() bool {
  return false
}

func (self AssignNode) IsAssignable() bool {
  return false
}

func (self AssignNode) IsLoadable() bool {
  return false
}

func (self AssignNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self AssignNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
