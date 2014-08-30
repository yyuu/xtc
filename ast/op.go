package ast

import (
  "fmt"
)

// BinaryOpNode
type binaryOpNode struct {
  location Location
  Operator string
  Left IExprNode
  Right IExprNode
}

func NewBinaryOpNode(location Location, operator string, left IExprNode, right IExprNode) binaryOpNode {
  return binaryOpNode { location, operator, left, right }
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

func (self binaryOpNode) GetLocation() Location {
  return self.location
}

// LogicalAndNode
type logicalAndNode struct {
  location Location
  Left IExprNode
  Right IExprNode
}

func NewLogicalAndNode(location Location, left IExprNode, right IExprNode) logicalAndNode {
  return logicalAndNode { location, left, right }
}

func (self logicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

func (self logicalAndNode) IsExpr() bool {
  return true
}

func (self logicalAndNode) GetLocation() Location {
  return self.location
}

// LogicalOrNode
type logicalOrNode struct {
  location Location
  Left IExprNode
  Right IExprNode
}

func NewLogicalOrNode(location Location, left IExprNode, right IExprNode) logicalOrNode {
  return logicalOrNode { location, left, right }
}

func (self logicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self logicalOrNode) IsExpr() bool {
  return true
}

func (self logicalOrNode) GetLocation() Location {
  return self.location
}

// PrefixOpNode
type prefixOpNode struct {
  location Location
  Operator string
  Expr IExprNode
}

func NewPrefixOpNode(location Location, operator string, expr IExprNode) prefixOpNode {
  return prefixOpNode { location, operator, expr }
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

func (self prefixOpNode) GetLocation() Location {
  return self.location
}

// SuffixOpNode
type suffixOpNode struct {
  location Location
  Operator string
  Expr IExprNode
}

func NewSuffixOpNode(location Location, operator string, expr IExprNode) suffixOpNode {
  return suffixOpNode { location, operator, expr }
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

func (self suffixOpNode) GetLocation() Location {
  return self.location
}

// UnaryOpNode
type unaryOpNode struct {
  location Location
  Operator string
  Expr IExprNode
}

func NewUnaryOpNode(location Location, operator string, expr IExprNode) unaryOpNode {
  return unaryOpNode { location, operator, expr }
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

func (self unaryOpNode) GetLocation() Location {
  return self.location
}
