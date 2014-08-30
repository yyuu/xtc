package ast

import (
  "fmt"
)

// BinaryOpNode
type BinaryOpNode struct {
  location Location
  Operator string
  Left IExprNode
  Right IExprNode
}

func NewBinaryOpNode(location Location, operator string, left IExprNode, right IExprNode) BinaryOpNode {
  return BinaryOpNode { location, operator, left, right }
}

func (self BinaryOpNode) String() string {
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

func (self BinaryOpNode) IsExpr() bool {
  return true
}

func (self BinaryOpNode) GetLocation() Location {
  return self.location
}

// LogicalAndNode
type LogicalAndNode struct {
  location Location
  Left IExprNode
  Right IExprNode
}

func NewLogicalAndNode(location Location, left IExprNode, right IExprNode) LogicalAndNode {
  return LogicalAndNode { location, left, right }
}

func (self LogicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

func (self LogicalAndNode) IsExpr() bool {
  return true
}

func (self LogicalAndNode) GetLocation() Location {
  return self.location
}

// LogicalOrNode
type LogicalOrNode struct {
  location Location
  Left IExprNode
  Right IExprNode
}

func NewLogicalOrNode(location Location, left IExprNode, right IExprNode) LogicalOrNode {
  return LogicalOrNode { location, left, right }
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self LogicalOrNode) IsExpr() bool {
  return true
}

func (self LogicalOrNode) GetLocation() Location {
  return self.location
}

// PrefixOpNode
type PrefixOpNode struct {
  location Location
  Operator string
  Expr IExprNode
}

func NewPrefixOpNode(location Location, operator string, expr IExprNode) PrefixOpNode {
  return PrefixOpNode { location, operator, expr }
}

func (self PrefixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self PrefixOpNode) IsExpr() bool {
  return true
}

func (self PrefixOpNode) GetLocation() Location {
  return self.location
}

// SuffixOpNode
type SuffixOpNode struct {
  location Location
  Operator string
  Expr IExprNode
}

func NewSuffixOpNode(location Location, operator string, expr IExprNode) SuffixOpNode {
  return SuffixOpNode { location, operator, expr }
}

func (self SuffixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self SuffixOpNode) IsExpr() bool {
  return true
}

func (self SuffixOpNode) GetLocation() Location {
  return self.location
}

// UnaryOpNode
type UnaryOpNode struct {
  location Location
  Operator string
  Expr IExprNode
}

func NewUnaryOpNode(location Location, operator string, expr IExprNode) UnaryOpNode {
  return UnaryOpNode { location, operator, expr }
}

func (self UnaryOpNode) String() string {
  switch self.Operator {
    case "!": return fmt.Sprintf("(not %s)", self.Expr)
    default:  return fmt.Sprintf("%s%s", self.Operator, self.Expr)
  }
}

func (self UnaryOpNode) IsExpr() bool {
  return true
}

func (self UnaryOpNode) GetLocation() Location {
  return self.location
}
