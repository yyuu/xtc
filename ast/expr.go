package ast

import (
  "fmt"
  "strconv"
  "strings"
)

type addressNode struct {
  Expr IExprNode
}

type arefNode struct {
  Expr IExprNode
  Index IExprNode
}

func (self arefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

type assignNode struct {
  Lhs IExprNode
  Rhs IExprNode
}

func AssignNode(lhs IExprNode, rhs IExprNode) assignNode {
  return assignNode { lhs, rhs }
}

func (self assignNode) String() string {
  return fmt.Sprintf("(define %s %s)", self.Lhs, self.Rhs)
}

type binaryOpNode struct {
  Operator string
  Left IExprNode
  Right IExprNode
}

func BinaryOpNode(operator string, left IExprNode, right IExprNode) binaryOpNode {
  return binaryOpNode { operator, left, right }
}

func (self binaryOpNode) String() string {
  switch self.Operator {
    case "&&": return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
    case "||": return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
    case "==": return fmt.Sprintf("(= %s %s)", self.Left, self.Right)
    case "!=": return fmt.Sprintf("(not (= %s %s))", self.Left, self.Right)
    default:   return fmt.Sprintf("(%s %s %s)", self.Operator, self.Left, self.Right)
  }
}

type castNode struct {
  Type ITypeNode
  Expr IExprNode
}

type condExprNode struct {
  Cond IExprNode
  ThenExpr IExprNode
  ElseExpr IExprNode
}

func CondExprNode(cond IExprNode, thenExpr IExprNode, elseExpr IExprNode) condExprNode {
  return condExprNode { cond, thenExpr, elseExpr }
}

func (self condExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

type dereferenceNode struct {
  Expr IExprNode
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

type logicalAndNode struct {
  Left IExprNode
  Right IExprNode
}

func LogicalAndNode(left IExprNode, right IExprNode) logicalAndNode {
  return logicalAndNode { left, right }
}

func (self logicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

type logicalOrNode struct {
  Left IExprNode
  Right IExprNode
}

func LogicalOrNode(left IExprNode, right IExprNode) logicalOrNode {
  return logicalOrNode { left, right }
}

func (self logicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

type memberNode struct {
  Expr IExprNode
  Member string
}

func (self memberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

type opAssignNode struct {
  Operator string
  Lhs IExprNode
  Rhs IExprNode
}

func OpAssignNode(operator string, lhs IExprNode, rhs IExprNode) opAssignNode {
  return opAssignNode { operator, lhs, rhs }
}

func (self opAssignNode) String() string {
  return fmt.Sprintf("(define %s (%s %s %s)", self.Lhs, self.Operator, self.Lhs, self.Rhs)
}

type prefixOpNode struct {
  Operator string
  Expr IExprNode
}

func PrefixOpNode(operator string, expr IExprNode) prefixOpNode {
  return prefixOpNode { operator, expr }
}

func (self prefixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

type ptrMemberNode struct {
  Expr IExprNode
  Member string
}

func (self ptrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

type sizeofExprNode struct {
  Expr IExprNode
  Type ITypeNode
}

type sizeofTypeNode struct {
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
  return self.Value
}

type suffixOpNode struct {
  Operator string
  Expr IExprNode
}

func SuffixOpNode(operator string, expr IExprNode) suffixOpNode {
  return suffixOpNode { operator, expr }
}

func (self suffixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

type unaryOpNode struct {
  Operator string
  Expr IExprNode
}

func UnaryOpNode(operator string, expr IExprNode) unaryOpNode {
  return unaryOpNode { operator, expr }
}

func (self unaryOpNode) String() string {
  switch self.Operator {
    case "!": return fmt.Sprintf("(not %s)", self.Expr)
    default:  return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
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
