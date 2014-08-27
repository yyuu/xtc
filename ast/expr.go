package ast

import (
  "fmt"
  "strconv"
  "strings"
)

type AddressNode struct {
  Expr IExprNode
}

type ArefNode struct {
  Expr IExprNode
  Index IExprNode
}

func (self ArefNode) DumpString() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr.DumpString(), self.Index.DumpString())
}

type AssignNode struct {
  Lhs IExprNode
  Rhs IExprNode
}

func (self AssignNode) DumpString() string {
  return fmt.Sprintf("(define %s %s)", self.Lhs.DumpString(), self.Rhs.DumpString())
}

type BinaryOpNode struct {
  Operator string
  Left IExprNode
  Right IExprNode
}

func (self BinaryOpNode) DumpString() string {
  switch self.Operator {
    case "&&": return fmt.Sprintf("(and %s %s)", self.Left.DumpString(), self.Right.DumpString())
    case "||": return fmt.Sprintf("(or %s %s)", self.Left.DumpString(), self.Right.DumpString())
    case "==": return fmt.Sprintf("(= %s %s)", self.Left.DumpString(), self.Right.DumpString())
    case "!=": return fmt.Sprintf("(not (= %s %s))", self.Left.DumpString(), self.Right.DumpString())
    default:   return fmt.Sprintf("(%s %s %s)", self.Operator, self.Left.DumpString(), self.Right.DumpString())
  }
}

type CastNode struct {
  Type ITypeNode
  Expr IExprNode
}

type CondExprNode struct {
  Cond IExprNode
  ThenExpr IExprNode
  ElseExpr IExprNode
}

func (self CondExprNode) DumpString() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond.DumpString(), self.ThenExpr.DumpString(), self.ElseExpr.DumpString())
}

type DereferenceNode struct {
  Expr IExprNode
}

type FuncallNode struct {
  Expr IExprNode
  Args []IExprNode
}

func (self FuncallNode) DumpString() string {
  sArgs := make([]string, len(self.Args))
  for i := range self.Args {
    sArgs[i] = self.Args[i].DumpString()
  }
  if len(sArgs) == 0 {
    return fmt.Sprintf("(%s)", self.Expr.DumpString())
  } else {
    return fmt.Sprintf("(%s %s)", self.Expr.DumpString(), strings.Join(sArgs, " "))
  }
}

type IntegerLiteralNode struct {
  Value int
}

func (self IntegerLiteralNode) DumpString() string {
  return strconv.Itoa(self.Value)
}

type LogicalAndNode struct {
  Left IExprNode
  Right IExprNode
}

func (self LogicalAndNode) DumpString() string {
  return fmt.Sprintf("(and %s %s)", self.Left.DumpString(), self.Right.DumpString())
}

type LogicalOrNode struct {
  Left IExprNode
  Right IExprNode
}

func (self LogicalOrNode) DumpString() string {
  return fmt.Sprintf("(or %s %s)", self.Left.DumpString(), self.Right.DumpString())
}

type MemberNode struct {
  Expr IExprNode
  Member string
}

func (self MemberNode) DumpString() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr.DumpString(), self.Member)
}

type OpAssignNode struct {
  Lhs IExprNode
  Rhs IExprNode
}

func (self OpAssignNode) DumpString() string {
  return fmt.Sprintf("(define %s %s)", self.Lhs.DumpString(), self.Rhs.DumpString())
}

type PrefixOpNode struct {
  Operator string
  Expr IExprNode
}

func (self PrefixOpNode) DumpString() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr.DumpString())
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr.DumpString())
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr.DumpString())
  }
}

type PtrMemberNode struct {
  Expr IExprNode
  Member string
}

func (self PtrMemberNode) DumpString() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr.DumpString(), self.Member)
}

type SizeofExprNode struct {
  Expr IExprNode
  Type ITypeNode
}

type SizeofTypeNode struct {
  Type ITypeNode
  Operand ITypeNode
}

type StringLiteralNode struct {
  Value string
}

func (self StringLiteralNode) DumpString() string {
  return fmt.Sprintf("%q", self.Value)
}

type SuffixOpNode struct {
  Operator string
  Expr IExprNode
}

func (self SuffixOpNode) DumpString() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr.DumpString())
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr.DumpString())
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr.DumpString())
  }
}

type VariableNode struct {
  Name string
}

func (self VariableNode) DumpString() string {
  return self.Name
}
