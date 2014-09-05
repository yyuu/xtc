package ast

import (
  "encoding/json"
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// AddressNode
type AddressNode struct {
  location duck.ILocation
  expr duck.IExprNode
}

func NewAddressNode(loc duck.ILocation, expr duck.IExprNode) AddressNode {
  return AddressNode { loc, expr }
}

func (self AddressNode) String() string {
  return fmt.Sprintf("<ast.AddressNode location=%s expr=%s>", self.location, self.expr)
}

func (self AddressNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
  }
  x.ClassName = "ast.AddressNode"
  x.Location = self.location
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self AddressNode) IsExpr() bool {
  return true
}

func (self AddressNode) GetLocation() duck.ILocation {
  return self.location
}

// ArefNode
type ArefNode struct {
  location duck.ILocation
  expr duck.IExprNode
  index duck.IExprNode
}

func NewArefNode(loc duck.ILocation, expr duck.IExprNode, index duck.IExprNode) ArefNode {
  return ArefNode { loc, expr, index }
}

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.expr, self.index)
}

func (self ArefNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Index duck.IExprNode
  }
  x.ClassName = "ast.ArefNode"
  x.Location = self.location
  x.Expr = self.expr
  x.Index = self.index
  return json.Marshal(x)
}

func (self ArefNode) IsExpr() bool {
  return true
}

func (self ArefNode) GetLocation() duck.ILocation {
  return self.location
}

// DereferenceNode
type DereferenceNode struct {
  location duck.ILocation
  expr duck.IExprNode
}

func NewDereferenceNode(loc Location, expr duck.IExprNode) DereferenceNode {
  return DereferenceNode { loc, expr }
}

func (self DereferenceNode) String() string {
  panic("not implemented")
}

func (self DereferenceNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
  }
  x.ClassName = "ast.DereferenceNode"
  x.Location = self.location
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self DereferenceNode) IsExpr() bool {
  return true
}

func (self DereferenceNode) GetLocation() duck.ILocation {
  return self.location
}

// FuncallNode
type FuncallNode struct {
  location duck.ILocation
  expr duck.IExprNode
  args []duck.IExprNode
}

func NewFuncallNode(loc duck.ILocation, expr duck.IExprNode, args []duck.IExprNode) FuncallNode {
  return FuncallNode { loc, expr, args }
}

func (self FuncallNode) String() string {
  sArgs := make([]string, len(self.args))
  for i := range self.args {
    sArgs[i] = fmt.Sprintf("%s", self.args[i])
  }
  if len(sArgs) == 0 {
    return fmt.Sprintf("(%s)", self.expr)
  } else {
    return fmt.Sprintf("(%s %s)", self.expr, strings.Join(sArgs, " "))
  }
}

func (self FuncallNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Args []duck.IExprNode
  }
  x.ClassName = "ast.FuncallNode"
  x.Location = self.location
  x.Expr = self.expr
  x.Args = self.args
  return json.Marshal(x)
}

func (self FuncallNode) IsExpr() bool {
  return true
}

func (self FuncallNode) GetLocation() duck.ILocation {
  return self.location
}

// MemberNode
type MemberNode struct {
  location duck.ILocation
  expr duck.IExprNode
  member string
}

func NewMemberNode(loc duck.ILocation, expr duck.IExprNode, member string) MemberNode {
  return MemberNode { loc, expr, member }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.expr, self.member)
}

func (self MemberNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Member string
  }
  x.ClassName = "ast.MemberNode"
  x.Location = self.location
  x.Expr = self.expr
  x.Member = self.member
  return json.Marshal(x)
}

func (self MemberNode) IsExpr() bool {
  return true
}

func (self MemberNode) GetLocation() duck.ILocation {
  return self.location
}

// PtrMemberNode
type PtrMemberNode struct {
  location duck.ILocation
  expr duck.IExprNode
  member string
}

func NewPtrMemberNode(loc duck.ILocation, expr duck.IExprNode, member string) PtrMemberNode {
  return PtrMemberNode { loc, expr, member }
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.expr, self.member)
}

func (self PtrMemberNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Member string
  }
  x.ClassName = "ast.PtrMemberNode"
  x.Location = self.location
  x.Expr = self.expr
  x.Member = self.member
  return json.Marshal(x)
}

func (self PtrMemberNode) IsExpr() bool {
  return true
}

func (self PtrMemberNode) GetLocation() duck.ILocation {
  return self.location
}

// VariableNode
type VariableNode struct {
  location duck.ILocation
  name string
}

func NewVariableNode(loc duck.ILocation, name string) VariableNode {
  return VariableNode { loc, name }
}

func (self VariableNode) String() string {
  return self.name
}

func (self VariableNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
  }
  x.ClassName = "ast.VariableNode"
  x.Location = self.location
  x.Name = self.name
  return json.Marshal(x)
}

func (self VariableNode) IsExpr() bool {
  return true
}

func (self VariableNode) GetLocation() duck.ILocation {
  return self.location
}
