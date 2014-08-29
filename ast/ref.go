package ast

import (
  "fmt"
  "strings"
)

// AddressNode
type addressNode struct {
  location Location
  Expr IExprNode
}

func AddressNode(location Location, expr IExprNode) addressNode {
  return addressNode { location, expr }
}

func (self addressNode) String() string {
  panic("not implemented")
}

func (self addressNode) IsExpr() bool {
  return true
}

func (self addressNode) GetLocation() Location {
  return self.location
}

// ArefNode
type arefNode struct {
  location Location
  Expr IExprNode
  Index IExprNode
}

func ArefNode(location Location, expr IExprNode, index IExprNode) arefNode {
  return arefNode { location, expr, index }
}

func (self arefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self arefNode) IsExpr() bool {
  return true
}

func (self arefNode) GetLocation() Location {
  return self.location
}

// FuncallNode
type funcallNode struct {
  location Location
  Expr IExprNode
  Args []IExprNode
}

func FuncallNode(location Location, expr IExprNode, args []IExprNode) funcallNode {
  return funcallNode { location, expr, args }
}

func (self funcallNode) String() string {
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

func (self funcallNode) IsExpr() bool {
  return true
}

func (self funcallNode) GetLocation() Location {
  return self.location
}

// MemberNode
type memberNode struct {
  location Location
  Expr IExprNode
  Member string
}

func MemberNode(location Location, expr IExprNode, member string) memberNode {
  return memberNode { location, expr, member }
}

func (self memberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self memberNode) IsExpr() bool {
  return true
}

func (self memberNode) GetLocation() Location {
  return self.location
}

// PtrMemberNode
type ptrMemberNode struct {
  location Location
  Expr IExprNode
  Member string
}

func PtrMemberNode(location Location, expr IExprNode, member string) ptrMemberNode {
  return ptrMemberNode { location, expr, member }
}

func (self ptrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self ptrMemberNode) IsExpr() bool {
  return true
}

func (self ptrMemberNode) GetLocation() Location {
  return self.location
}

// VariableNode
type variableNode struct {
  location Location
  Name string
}

func VariableNode(location Location, name string) variableNode {
  return variableNode { location, name }
}

func (self variableNode) String() string {
  return self.Name
}

func (self variableNode) IsExpr() bool {
  return true
}

func (self variableNode) GetLocation() Location {
  return self.location
}
