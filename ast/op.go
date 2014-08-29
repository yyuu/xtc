package ast

import (
  "fmt"
)

// BinaryOpNode
type binaryOpNode struct {
  Location ILocation
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
    case "<<": return fmt.Sprintf("(bitwise-arithmetic-left %s %s)", self.Left, self.Right)
    case ">>": return fmt.Sprintf("(bitwise-arithmetic-right %s %s)", self.Left, self.Right)
    case "%":  return fmt.Sprintf("(modulo %s %s)", self.Left, self.Right)
    default:   return fmt.Sprintf("(%s %s %s)", self.Operator, self.Left, self.Right)
  }
}

func (self binaryOpNode) IsExpr() bool {
  return true
}

// LogicalAndNode
type logicalAndNode struct {
  Location ILocation
  Left IExprNode
  Right IExprNode
}

func LogicalAndNode(left IExprNode, right IExprNode) logicalAndNode {
  return logicalAndNode { left, right }
}

func (self logicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

func (self logicalAndNode) IsExpr() bool {
  return true
}

// LogicalOrNode
type logicalOrNode struct {
  Location ILocation
  Left IExprNode
  Right IExprNode
}

func LogicalOrNode(left IExprNode, right IExprNode) logicalOrNode {
  return logicalOrNode { left, right }
}

func (self logicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self logicalOrNode) IsExpr() bool {
  return true
}

// PrefixOpNode
type prefixOpNode struct {
  Location ILocation
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

func (self prefixOpNode) IsExpr() bool {
  return true
}

// SuffixOpNode
type suffixOpNode struct {
  Location ILocation
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

func (self suffixOpNode) IsExpr() bool {
  return true
}

// UnaryOpNode
type unaryOpNode struct {
  Location ILocation
  Operator string
  Expr IExprNode
}

func UnaryOpNode(operator string, expr IExprNode) unaryOpNode {
  return unaryOpNode { operator, expr }
}

func (self unaryOpNode) String() string {
  switch self.Operator {
    case "!": return fmt.Sprintf("(not %s)", self.Expr)
    default:  return fmt.Sprintf("%s%s", self.Operator, self.Expr)
  }
}

func (self unaryOpNode) IsExpr() bool {
  return true
}
