package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// AddressNode
type AddressNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
}

func NewAddressNode(loc duck.Location, expr duck.IExprNode) AddressNode {
  if expr == nil { panic("expr is nil") }
  return AddressNode { "ast.AddressNode", loc, expr }
}

func (self AddressNode) String() string {
  return fmt.Sprintf("<ast.AddressNode location=%s expr=%s>", self.Location, self.Expr)
}

func (self AddressNode) IsExprNode() bool {
  return true
}

func (self AddressNode) GetLocation() duck.Location {
  return self.Location
}

// ArefNode
type ArefNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
  Index duck.IExprNode
}

func NewArefNode(loc duck.Location, expr duck.IExprNode, index duck.IExprNode) ArefNode {
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

func (self ArefNode) GetLocation() duck.Location {
  return self.Location
}

// DereferenceNode
type DereferenceNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
}

func NewDereferenceNode(loc duck.Location, expr duck.IExprNode) DereferenceNode {
  if expr == nil { panic("expr is nil") }
  return DereferenceNode { "ast.DereferenceNode", loc, expr }
}

func (self DereferenceNode) String() string {
  panic("not implemented")
}

func (self DereferenceNode) IsExprNode() bool {
  return true
}

func (self DereferenceNode) GetLocation() duck.Location {
  return self.Location
}

// FuncallNode
type FuncallNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
  Args []duck.IExprNode
}

func NewFuncallNode(loc duck.Location, expr duck.IExprNode, args []duck.IExprNode) FuncallNode {
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

func (self FuncallNode) GetLocation() duck.Location {
  return self.Location
}

// MemberNode
type MemberNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
  Member string
}

func NewMemberNode(loc duck.Location, expr duck.IExprNode, member string) MemberNode {
  if expr == nil { panic("expr is nil") }
  return MemberNode { "ast.MemberNode", loc, expr, member }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self MemberNode) IsExprNode() bool {
  return true
}

func (self MemberNode) GetLocation() duck.Location {
  return self.Location
}

// PtrMemberNode
type PtrMemberNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
  Member string
}

func NewPtrMemberNode(loc duck.Location, expr duck.IExprNode, member string) PtrMemberNode {
  if expr == nil { panic("expr is nil") }
  return PtrMemberNode { "ast.PtrMemberNode", loc, expr, member }
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self PtrMemberNode) IsExprNode() bool {
  return true
}

func (self PtrMemberNode) GetLocation() duck.Location {
  return self.Location
}

// VariableNode
type VariableNode struct {
  ClassName string
  Location duck.Location
  Name string
  entity duck.IEntity
}

func NewVariableNode(loc duck.Location, name string) VariableNode {
  return VariableNode { "ast.VariableNode", loc, name, nil }
}

func (self VariableNode) String() string {
  return self.Name
}

func (self VariableNode) IsExprNode() bool {
  return true
}

func (self VariableNode) GetLocation() duck.Location {
  return self.Location
}

func (self VariableNode) GetName() string {
  return self.Name
}

func (self *VariableNode) SetEntity(ent duck.IEntity) {
  self.entity = ent
}

func (self VariableNode) GetEntity() duck.IEntity {
  return self.entity
}
