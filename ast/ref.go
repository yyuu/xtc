package ast

import (
  "encoding/json"
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// AddressNode
type AddressNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
}

func NewAddressNode(location duck.ILocation, expr duck.IExprNode) AddressNode {
  return AddressNode { location, expr }
}

func (self AddressNode) String() string {
  panic("not implemented")
}

func (self AddressNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
  }
  x.ClassName = "ast.AddressNode"
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self AddressNode) IsExpr() bool {
  return true
}

func (self AddressNode) GetLocation() duck.ILocation {
  return self.Location
}

// ArefNode
type ArefNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
  Index duck.IExprNode
}

func NewArefNode(location duck.ILocation, expr duck.IExprNode, index duck.IExprNode) ArefNode {
  return ArefNode { location, expr, index }
}

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self ArefNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Index duck.IExprNode
  }
  x.ClassName = "ast.ArefNode"
  x.Location = self.Location
  x.Expr = self.Expr
  x.Index = self.Index
  return json.Marshal(x)
}

func (self ArefNode) IsExpr() bool {
  return true
}

func (self ArefNode) GetLocation() duck.ILocation {
  return self.Location
}

// DereferenceNode
type DereferenceNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
}

func NewDereferenceNode(location Location, expr duck.IExprNode) DereferenceNode {
  return DereferenceNode { location, expr }
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
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self DereferenceNode) IsExpr() bool {
  return true
}

func (self DereferenceNode) GetLocation() duck.ILocation {
  return self.Location
}

// FuncallNode
type FuncallNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
  Args []duck.IExprNode
}

func NewFuncallNode(location duck.ILocation, expr duck.IExprNode, args []duck.IExprNode) FuncallNode {
  return FuncallNode { location, expr, args }
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

func (self FuncallNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Args []duck.IExprNode
  }
  x.ClassName = "ast.FuncallNode"
  x.Location = self.Location
  x.Expr = self.Expr
  x.Args = self.Args
  return json.Marshal(x)
}

func (self FuncallNode) IsExpr() bool {
  return true
}

func (self FuncallNode) GetLocation() duck.ILocation {
  return self.Location
}

// MemberNode
type MemberNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
  Member string
}

func NewMemberNode(location duck.ILocation, expr duck.IExprNode, member string) MemberNode {
  return MemberNode { location, expr, member }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self MemberNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Member string
  }
  x.ClassName = "ast.MemberNode"
  x.Location = self.Location
  x.Expr = self.Expr
  x.Member = self.Member
  return json.Marshal(x)
}

func (self MemberNode) IsExpr() bool {
  return true
}

func (self MemberNode) GetLocation() duck.ILocation {
  return self.Location
}

// PtrMemberNode
type PtrMemberNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
  Member string
}

func NewPtrMemberNode(location duck.ILocation, expr duck.IExprNode, member string) PtrMemberNode {
  return PtrMemberNode { location, expr, member }
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self PtrMemberNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Member string
  }
  x.ClassName = "ast.PtrMemberNode"
  x.Location = self.Location
  x.Expr = self.Expr
  x.Member = self.Member
  return json.Marshal(x)
}

func (self PtrMemberNode) IsExpr() bool {
  return true
}

func (self PtrMemberNode) GetLocation() duck.ILocation {
  return self.Location
}

// VariableNode
type VariableNode struct {
  Location duck.ILocation
  Name string
}

func NewVariableNode(location duck.ILocation, name string) VariableNode {
  return VariableNode { location, name }
}

func (self VariableNode) String() string {
  return self.Name
}

func (self VariableNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
  }
  x.ClassName = "ast.VariableNode"
  x.Location = self.Location
  x.Name = self.Name
  return json.Marshal(x)
}

func (self VariableNode) IsExpr() bool {
  return true
}

func (self VariableNode) GetLocation() duck.ILocation {
  return self.Location
}
