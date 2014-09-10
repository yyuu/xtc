package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

// AddressNode
type AddressNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewAddressNode(loc core.Location, expr core.IExprNode) AddressNode {
  if expr == nil { panic("expr is nil") }
  return AddressNode { "ast.AddressNode", loc, expr }
}

func (self AddressNode) String() string {
  return fmt.Sprintf("<ast.AddressNode location=%s expr=%s>", self.Location, self.Expr)
}

func (self AddressNode) IsExprNode() bool {
  return true
}

func (self AddressNode) GetLocation() core.Location {
  return self.Location
}

// ArefNode
type ArefNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Index core.IExprNode
}

func NewArefNode(loc core.Location, expr core.IExprNode, index core.IExprNode) ArefNode {
  if expr == nil { panic("expr is nil") }
  if index == nil { panic("index is nil") }
  return ArefNode { "ast.ArefNode", loc, expr, index }
}

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self ArefNode) IsExprNode() bool {
  return true
}

func (self ArefNode) GetLocation() core.Location {
  return self.Location
}

// DereferenceNode
type DereferenceNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewDereferenceNode(loc core.Location, expr core.IExprNode) DereferenceNode {
  if expr == nil { panic("expr is nil") }
  return DereferenceNode { "ast.DereferenceNode", loc, expr }
}

func (self DereferenceNode) String() string {
  panic("not implemented")
}

func (self DereferenceNode) IsExprNode() bool {
  return true
}

func (self DereferenceNode) GetLocation() core.Location {
  return self.Location
}

// FuncallNode
type FuncallNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Args []core.IExprNode
}

func NewFuncallNode(loc core.Location, expr core.IExprNode, args []core.IExprNode) FuncallNode {
  if expr == nil { panic("expr is nil") }
  return FuncallNode { "ast.FuncallNode", loc, expr, args }
}

func (self FuncallNode) String() string {
  sArgs := make([]string, len(self.Args))
  for i := range self.Args {
    sArgs[i] = fmt.Sprintf("%s", self.Args[i])
  }
  if len(sArgs) == 0 {
    return fmt.Sprintf("(%s)", self.Expr)
  } else {
    return fmt.Sprintf("(%s %s)", self.Expr, strings.Join(sArgs, " "))
  }
}

func (self FuncallNode) IsExprNode() bool {
  return true
}

func (self FuncallNode) GetLocation() core.Location {
  return self.Location
}

// MemberNode
type MemberNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Member string
}

func NewMemberNode(loc core.Location, expr core.IExprNode, member string) MemberNode {
  if expr == nil { panic("expr is nil") }
  return MemberNode { "ast.MemberNode", loc, expr, member }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self MemberNode) IsExprNode() bool {
  return true
}

func (self MemberNode) GetLocation() core.Location {
  return self.Location
}

// PtrMemberNode
type PtrMemberNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Member string
}

func NewPtrMemberNode(loc core.Location, expr core.IExprNode, member string) PtrMemberNode {
  if expr == nil { panic("expr is nil") }
  return PtrMemberNode { "ast.PtrMemberNode", loc, expr, member }
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

// VariableNode
type VariableNode struct {
  ClassName string
  Location core.Location
  Name string
  entity core.IEntity
}

func NewVariableNode(loc core.Location, name string) VariableNode {
  return VariableNode { "ast.VariableNode", loc, name, nil }
}

func (self VariableNode) String() string {
  return self.Name
}

func (self VariableNode) IsExprNode() bool {
  return true
}

func (self VariableNode) GetLocation() core.Location {
  return self.Location
}

func (self VariableNode) GetName() string {
  return self.Name
}

func (self *VariableNode) SetEntity(ent core.IEntity) {
  self.entity = ent
}

func (self VariableNode) GetEntity() core.IEntity {
  return self.entity
}
