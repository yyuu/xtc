package ast

import (
  "fmt"
  "strings"
)

// AddressNode
type addressNode struct {
  Location ILocation
  Expr IExprNode
}

func AddressNode(expr IExprNode) addressNode {
  return addressNode { expr }
}

func (self addressNode) String() string {
  panic("not implemented")
}

func (self addressNode) IsExpr() bool {
  return true
}

// ArefNode
type arefNode struct {
  Location ILocation
  Expr IExprNode
  Index IExprNode
}

func ArefNode(expr IExprNode, index IExprNode) arefNode {
  return arefNode { expr, index }
}

func (self arefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self arefNode) IsExpr() bool {
  return true
}

// FuncallNode
type funcallNode struct {
  Location ILocation
  Expr IExprNode
  Args []IExprNode
}

func FuncallNode(expr IExprNode, args []IExprNode) funcallNode {
  return funcallNode { expr, args }
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

// MemberNode
type memberNode struct {
  Location ILocation
  Expr IExprNode
  Member string
}

func MemberNode(expr IExprNode, member string) memberNode {
  return memberNode { expr, member }
}

func (self memberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self memberNode) IsExpr() bool {
  return true
}

// PtrMemberNode
type ptrMemberNode struct {
  Location ILocation
  Expr IExprNode
  Member string
}

func PtrMemberNode(expr IExprNode, member string) ptrMemberNode {
  return ptrMemberNode { expr, member }
}

func (self ptrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self ptrMemberNode) IsExpr() bool {
  return true
}

// VariableNode
type variableNode struct {
  Location ILocation
  Name string
}

func VariableNode(name string) variableNode {
  return variableNode { name }
}

func (self variableNode) String() string {
  return self.Name
}

func (self variableNode) IsExpr() bool {
  return true
}
