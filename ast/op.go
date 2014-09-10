package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// BinaryOpNode
type BinaryOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Left core.IExprNode
  Right core.IExprNode
}

func NewBinaryOpNode(loc core.Location, operator string, left core.IExprNode, right core.IExprNode) BinaryOpNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return BinaryOpNode { "ast.BinaryOpNode", loc, operator, left, right }
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

func (self BinaryOpNode) IsExprNode() bool {
  return true
}

func (self BinaryOpNode) GetLocation() core.Location {
  return self.Location
}

// LogicalAndNode
type LogicalAndNode struct {
  ClassName string
  Location core.Location
  Left core.IExprNode
  Right core.IExprNode
}

func NewLogicalAndNode(loc core.Location, left core.IExprNode, right core.IExprNode) LogicalAndNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return LogicalAndNode { "ast.LogicalAndNode", loc, left, right }
}

func (self LogicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

func (self LogicalAndNode) IsExprNode() bool {
  return true
}

func (self LogicalAndNode) GetLocation() core.Location {
  return self.Location
}

// LogicalOrNode
type LogicalOrNode struct {
  ClassName string
  Location core.Location
  Left core.IExprNode
  Right core.IExprNode
}

func NewLogicalOrNode(loc core.Location, left core.IExprNode, right core.IExprNode) LogicalOrNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return LogicalOrNode { "ast.LogicalOrNode", loc, left, right }
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self LogicalOrNode) IsExprNode() bool {
  return true
}

func (self LogicalOrNode) GetLocation() core.Location {
  return self.Location
}

// PrefixOpNode
type PrefixOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
}

func NewPrefixOpNode(loc core.Location, operator string, expr core.IExprNode) PrefixOpNode {
  if expr == nil { panic("expr is nil") }
  return PrefixOpNode { "ast.PrefixOpNode", loc, operator, expr }
}

func (self PrefixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self PrefixOpNode) IsExprNode() bool {
  return true
}

func (self PrefixOpNode) GetLocation() core.Location {
  return self.Location
}

// SuffixOpNode
type SuffixOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
}

func NewSuffixOpNode(loc core.Location, operator string, expr core.IExprNode) SuffixOpNode {
  if expr == nil { panic("expr is nil") }
  return SuffixOpNode { "ast.SuffixOpNode", loc, operator, expr }
}

func (self SuffixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
  }
}

func (self SuffixOpNode) IsExprNode() bool {
  return true
}

func (self SuffixOpNode) GetLocation() core.Location {
  return self.Location
}

// UnaryOpNode
type UnaryOpNode struct {
  ClassName string
  Location core.Location
  Operator string
  Expr core.IExprNode
}

func NewUnaryOpNode(loc core.Location, operator string, expr core.IExprNode) UnaryOpNode {
  if expr == nil { panic("expr is nil") }
  return UnaryOpNode { "ast.UnaryOpNode", loc, operator, expr }
}

func (self UnaryOpNode) String() string {
  switch self.Operator {
    case "!": return fmt.Sprintf("(not %s)", self.Expr)
    default:  return fmt.Sprintf("%s%s", self.Operator, self.Expr)
  }
}

func (self UnaryOpNode) IsExprNode() bool {
  return true
}

func (self UnaryOpNode) GetLocation() core.Location {
  return self.Location
}
