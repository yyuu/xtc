package ast

import (
  "fmt"
  "strings"
)

type addressNode struct {
  Expr IExprNode
}

func AddressNode(expr INode) addressNode {
  return addressNode { expr.(IExprNode) }
}

func (self addressNode) String() string {
  panic("not implemented")
}

type arefNode struct {
  Expr IExprNode
  Index IExprNode
}

func ArefNode(expr INode, index INode) arefNode {
  return arefNode { expr.(IExprNode), index.(IExprNode) }
}

func (self arefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

type funcallNode struct {
  Expr IExprNode
  Args []IExprNode
}

func FuncallNode(_expr INode, _args []INode) funcallNode {
  expr := _expr.(IExprNode)
  args := make([]IExprNode, len(_args))
  for i := range _args {
    args[i] = _args[i].(IExprNode)
  }
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

type memberNode struct {
  Expr IExprNode
  Member string
}

func MemberNode(expr INode, member string) memberNode {
  return memberNode { expr.(IExprNode), member }
}

func (self memberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

type ptrMemberNode struct {
  Expr IExprNode
  Member string
}

func PtrMemberNode(expr INode, member string) ptrMemberNode {
  return ptrMemberNode { expr.(IExprNode), member }
}

func (self ptrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

type variableNode struct {
  Name string
}

func VariableNode(name string) variableNode {
  return variableNode { name }
}

func (self variableNode) String() string {
  return self.Name
}
