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

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

type AssignNode struct {
  Lhs IExprNode
  Rhs IExprNode
}

func (self AssignNode) String() string {
  return fmt.Sprintf("(define %s %s)", self.Lhs, self.Rhs)
}

type BinaryOpNode struct {
  Operator string
  Left IExprNode
  Right IExprNode
}

func (self BinaryOpNode) String() string {
  switch self.Operator {
    case "&&": return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
    case "||": return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
    case "==": return fmt.Sprintf("(= %s %s)", self.Left, self.Right)
    case "!=": return fmt.Sprintf("(not (= %s %s))", self.Left, self.Right)
    default:   return fmt.Sprintf("(%s %s %s)", self.Operator, self.Left, self.Right)
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

func (self CondExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

type DereferenceNode struct {
  Expr IExprNode
}

type FuncallNode struct {
  Expr IExprNode
  Args []IExprNode
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

type integerLiteralNode struct {
  Value int
}

func IntegerLiteralNode(literal string) integerLiteralNode {
  value, err := strconv.Atoi(literal)
  if err != nil { panic(err) }
  return integerLiteralNode { value }
}

func (self integerLiteralNode) String() string {
  return strconv.Itoa(self.Value)
}

type LogicalAndNode struct {
  Left IExprNode
  Right IExprNode
}

func (self LogicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

type LogicalOrNode struct {
  Left IExprNode
  Right IExprNode
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

type MemberNode struct {
  Expr IExprNode
  Member string
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

type OpAssignNode struct {
  Lhs IExprNode
  Rhs IExprNode
}

func (self OpAssignNode) String() string {
  return fmt.Sprintf("(define %s %s)", self.Lhs, self.Rhs)
}

type PrefixOpNode struct {
  Operator string
  Expr IExprNode
}

func (self PrefixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

type PtrMemberNode struct {
  Expr IExprNode
  Member string
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

type SizeofExprNode struct {
  Expr IExprNode
  Type ITypeNode
}

type SizeofTypeNode struct {
  Type ITypeNode
  Operand ITypeNode
}

type stringLiteralNode struct {
  Value string
}

func StringLiteralNode(literal string) stringLiteralNode {
  return stringLiteralNode { literal }
}

func (self stringLiteralNode) String() string {
  return fmt.Sprintf("%q", self.Value)
}

type SuffixOpNode struct {
  Operator string
  Expr IExprNode
}

func (self SuffixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

type VariableNode struct {
  Name string
}

func (self VariableNode) String() string {
  return self.Name
}
