package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// BinaryOpNode
type BinaryOpNode struct {
  location duck.ILocation
  operator string
  left duck.IExprNode
  right duck.IExprNode
}

func NewBinaryOpNode(loc duck.ILocation, operator string, left duck.IExprNode, right duck.IExprNode) BinaryOpNode {
  return BinaryOpNode { loc, operator, left, right }
}

func (self BinaryOpNode) String() string {
  switch self.operator {
    case "&&": return fmt.Sprintf("(and %s %s)", self.left, self.right)
    case "||": return fmt.Sprintf("(or %s %s)", self.left, self.right)
    case "==": return fmt.Sprintf("(= %s %s)", self.left, self.right)
    case "!=": return fmt.Sprintf("(not (= %s %s))", self.left, self.right)
    case "<<": return fmt.Sprintf("(bitwise-arithmetic-left %s %s)", self.left, self.right)
    case ">>": return fmt.Sprintf("(bitwise-arithmetic-right %s %s)", self.left, self.right)
    case "%":  return fmt.Sprintf("(modulo %s %s)", self.left, self.right)
    default:   return fmt.Sprintf("(%s %s %s)", self.operator, self.left, self.right)
  }
}

func (self BinaryOpNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Operator string
    Left duck.IExprNode
    Right duck.IExprNode
  }
  x.ClassName = "ast.BinaryOpNode"
  x.Location = self.location
  x.Operator = self.operator
  x.Left = self.left
  x.Right = self.right
  return json.Marshal(x)
}

func (self BinaryOpNode) IsExpr() bool {
  return true
}

func (self BinaryOpNode) GetLocation() duck.ILocation {
  return self.location
}

// LogicalAndNode
type LogicalAndNode struct {
  location duck.ILocation
  left duck.IExprNode
  right duck.IExprNode
}

func NewLogicalAndNode(loc duck.ILocation, left duck.IExprNode, right duck.IExprNode) LogicalAndNode {
  return LogicalAndNode { loc, left, right }
}

func (self LogicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.left, self.right)
}

func (self LogicalAndNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Left duck.IExprNode
    Right duck.IExprNode
  }
  x.ClassName = "ast.LogicalAndNode"
  x.Location = self.location
  x.Left = self.left
  x.Right = self.right
  return json.Marshal(x)
}

func (self LogicalAndNode) IsExpr() bool {
  return true
}

func (self LogicalAndNode) GetLocation() duck.ILocation {
  return self.location
}

// LogicalOrNode
type LogicalOrNode struct {
  location duck.ILocation
  left duck.IExprNode
  right duck.IExprNode
}

func NewLogicalOrNode(loc duck.ILocation, left duck.IExprNode, right duck.IExprNode) LogicalOrNode {
  return LogicalOrNode { loc, left, right }
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.left, self.right)
}

func (self LogicalOrNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Left duck.IExprNode
    Right duck.IExprNode
  }
  x.ClassName = "ast.LogicalOrNode"
  x.Location = self.location
  x.Left = self.left
  x.Right = self.right
  return json.Marshal(x)
}

func (self LogicalOrNode) IsExpr() bool {
  return true
}

func (self LogicalOrNode) GetLocation() duck.ILocation {
  return self.location
}

// PrefixOpNode
type PrefixOpNode struct {
  location duck.ILocation
  operator string
  expr duck.IExprNode
}

func NewPrefixOpNode(loc duck.ILocation, operator string, expr duck.IExprNode) PrefixOpNode {
  return PrefixOpNode { loc, operator, expr }
}

func (self PrefixOpNode) String() string {
  switch self.operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.expr)
    default:   return fmt.Sprintf("(%s %s)", self.operator, self.expr)
  }
}

func (self PrefixOpNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Operator string
    Expr duck.IExprNode
  }
  x.ClassName = "ast.PrefixOpNode"
  x.Location = self.location
  x.Operator = self.operator
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self PrefixOpNode) IsExpr() bool {
  return true
}

func (self PrefixOpNode) GetLocation() duck.ILocation {
  return self.location
}

// SuffixOpNode
type SuffixOpNode struct {
  location duck.ILocation
  operator string
  expr duck.IExprNode
}

func NewSuffixOpNode(loc duck.ILocation, operator string, expr duck.IExprNode) SuffixOpNode {
  return SuffixOpNode { loc, operator, expr }
}

func (self SuffixOpNode) String() string {
  switch self.operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.expr)
    default:   return fmt.Sprintf("(%s %s)", self.operator, self.expr)
  }
}

func (self SuffixOpNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Operator string
    Expr duck.IExprNode
  }
  x.ClassName = "ast.SuffixOpNode"
  x.Location = self.location
  x.Operator = self.operator
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self SuffixOpNode) IsExpr() bool {
  return true
}

func (self SuffixOpNode) GetLocation() duck.ILocation {
  return self.location
}

// UnaryOpNode
type UnaryOpNode struct {
  location duck.ILocation
  operator string
  expr duck.IExprNode
}

func NewUnaryOpNode(loc duck.ILocation, operator string, expr duck.IExprNode) UnaryOpNode {
  return UnaryOpNode { loc, operator, expr }
}

func (self UnaryOpNode) String() string {
  switch self.operator {
    case "!": return fmt.Sprintf("(not %s)", self.expr)
    default:  return fmt.Sprintf("%s%s", self.operator, self.expr)
  }
}

func (self UnaryOpNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Operator string
    Expr duck.IExprNode
  }
  x.ClassName = "ast.UnaryOpNode"
  x.Location = self.location
  x.Operator = self.operator
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self UnaryOpNode) IsExpr() bool {
  return true
}

func (self UnaryOpNode) GetLocation() duck.ILocation {
  return self.location
}
