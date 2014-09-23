package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// PtrMemberNode
type PtrMemberNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Member string
  t core.IType
}

func NewPtrMemberNode(loc core.Location, expr core.IExprNode, member string) *PtrMemberNode {
  if expr == nil { panic("expr is nil") }
  return &PtrMemberNode { "ast.PtrMemberNode", loc, expr, member, nil }
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self PtrMemberNode) IsExprNode() bool {
  return true
}

func (self PtrMemberNode) GetLocation() core.Location {
  return self.Location
}

func (self PtrMemberNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self PtrMemberNode) GetMember() string {
  return self.Member
}

func (self PtrMemberNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *PtrMemberNode) SetType(t core.IType) {
  self.t = t
}

func (self PtrMemberNode) IsConstant() bool {
  return false
}

func (self PtrMemberNode) IsParameter() bool {
  return false
}

func (self PtrMemberNode) IsLvalue() bool {
  return true
}

func (self PtrMemberNode) IsAssignable() bool {
  return true
}

func (self PtrMemberNode) IsLoadable() bool {
  return false
}

func (self PtrMemberNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self PtrMemberNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
