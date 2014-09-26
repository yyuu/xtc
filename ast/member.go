package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// MemberNode
type MemberNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Member string
  Type core.IType
}

func NewMemberNode(loc core.Location, expr core.IExprNode, member string) *MemberNode {
  if expr == nil { panic("expr is nil") }
  return &MemberNode { "ast.MemberNode", loc, expr, member, nil }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self *MemberNode) AsExprNode() core.IExprNode {
  return self
}

func (self MemberNode) GetLocation() core.Location {
  return self.Location
}

func (self MemberNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self MemberNode) GetMember() string {
  return self.Member
}

func (self MemberNode) GetType() core.IType {
  if self.Type == nil {
    self.Type = self.GetOrigType()
  }
  return self.Type
}

func (self *MemberNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self MemberNode) GetOrigType() core.IType {
  return self.GetBaseType().GetMemberType(self.Member)
}

func (self MemberNode) IsConstant() bool {
  return false
}

func (self MemberNode) IsParameter() bool {
  return false
}

func (self MemberNode) IsLvalue() bool {
  return true
}

func (self MemberNode) IsAssignable() bool {
  return true
}

func (self MemberNode) IsLoadable() bool {
  t := self.GetOrigType()
  return !t.IsArray() && !t.IsFunction()
}

func (self MemberNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self MemberNode) IsPointer() bool {
  return self.GetType().IsPointer()
}

func (self MemberNode) GetBaseType() core.ICompositeType {
  t := self.Expr.GetType()
  return t.(core.ICompositeType)
}

func (self MemberNode) GetOffset() int {
  return self.GetBaseType().GetMemberOffset(self.Member)
}
