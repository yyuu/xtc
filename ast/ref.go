package ast

import (
  "encoding/json"
  "fmt"
  "strings"
)

// AddressNode
type AddressNode struct {
  Location Location
  Expr IExprNode
}

func NewAddressNode(location Location, expr IExprNode) AddressNode {
  return AddressNode { location, expr }
}

func (self AddressNode) String() string {
  panic("not implemented")
}

func (self AddressNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
  }
  x.ClassName = "ast.AddressNode"
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self AddressNode) IsExpr() bool {
  return true
}

func (self AddressNode) GetLocation() Location {
  return self.Location
}

// ArefNode
type ArefNode struct {
  Location Location
  Expr IExprNode
  Index IExprNode
}

func NewArefNode(location Location, expr IExprNode, index IExprNode) ArefNode {
  return ArefNode { location, expr, index }
}

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self ArefNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
    Index IExprNode
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

func (self ArefNode) GetLocation() Location {
  return self.Location
}

// DereferenceNode
type DereferenceNode struct {
  Location Location
  Expr IExprNode
}

func NewDereferenceNode(location Location, expr IExprNode) DereferenceNode {
  return DereferenceNode { location, expr }
}

func (self DereferenceNode) String() string {
  panic("not implemented")
}

func (self DereferenceNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
  }
  x.ClassName = "ast.DereferenceNode"
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self DereferenceNode) IsExpr() bool {
  return true
}

func (self DereferenceNode) GetLocation() Location {
  return self.Location
}

// FuncallNode
type FuncallNode struct {
  Location Location
  Expr IExprNode
  Args []IExprNode
}

func NewFuncallNode(location Location, expr IExprNode, args []IExprNode) FuncallNode {
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
    Location Location
    Expr IExprNode
    Args []IExprNode
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

func (self FuncallNode) GetLocation() Location {
  return self.Location
}

// MemberNode
type MemberNode struct {
  Location Location
  Expr IExprNode
  Member string
}

func NewMemberNode(location Location, expr IExprNode, member string) MemberNode {
  return MemberNode { location, expr, member }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self MemberNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
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

func (self MemberNode) GetLocation() Location {
  return self.Location
}

// PtrMemberNode
type PtrMemberNode struct {
  Location Location
  Expr IExprNode
  Member string
}

func NewPtrMemberNode(location Location, expr IExprNode, member string) PtrMemberNode {
  return PtrMemberNode { location, expr, member }
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self PtrMemberNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
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

func (self PtrMemberNode) GetLocation() Location {
  return self.Location
}

// VariableNode
type VariableNode struct {
  Location Location
  Name string
}

func NewVariableNode(location Location, name string) VariableNode {
  return VariableNode { location, name }
}

func (self VariableNode) String() string {
  return self.Name
}

func (self VariableNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
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

func (self VariableNode) GetLocation() Location {
  return self.Location
}
