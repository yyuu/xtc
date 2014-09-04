package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// BinaryOpNode
type BinaryOpNode struct {
  Location duck.ILocation
  Operator string
  Left duck.IExprNode
  Right duck.IExprNode
}

func NewBinaryOpNode(location duck.ILocation, operator string, left duck.IExprNode, right duck.IExprNode) BinaryOpNode {
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

func (self BinaryOpNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Operator string
    Left duck.IExprNode
    Right duck.IExprNode
  }
  x.ClassName = "ast.BinaryOpNode"
  x.Location = self.Location
  x.Operator = self.Operator
  x.Left = self.Left
  x.Right = self.Right
  return json.Marshal(x)
}

func (self BinaryOpNode) IsExpr() bool {
  return true
}

func (self BinaryOpNode) GetLocation() duck.ILocation {
  return self.Location
}

// LogicalAndNode
type LogicalAndNode struct {
  Location duck.ILocation
  Left duck.IExprNode
  Right duck.IExprNode
}

func NewLogicalAndNode(location duck.ILocation, left duck.IExprNode, right duck.IExprNode) LogicalAndNode {
  return LogicalAndNode { location, left, right }
}

func (self LogicalAndNode) String() string {
  return fmt.Sprintf("(and %s %s)", self.Left, self.Right)
}

func (self LogicalAndNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Left duck.IExprNode
    Right duck.IExprNode
  }
  x.ClassName = "ast.LogicalAndNode"
  x.Location = self.Location
  x.Left = self.Left
  x.Right = self.Right
  return json.Marshal(x)
}

func (self LogicalAndNode) IsExpr() bool {
  return true
}

func (self LogicalAndNode) GetLocation() duck.ILocation {
  return self.Location
}

// LogicalOrNode
type LogicalOrNode struct {
  Location duck.ILocation
  Left duck.IExprNode
  Right duck.IExprNode
}

func NewLogicalOrNode(location duck.ILocation, left duck.IExprNode, right duck.IExprNode) LogicalOrNode {
  return LogicalOrNode { location, left, right }
}

func (self LogicalOrNode) String() string {
  return fmt.Sprintf("(or %s %s)", self.Left, self.Right)
}

func (self LogicalOrNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Left duck.IExprNode
    Right duck.IExprNode
  }
  x.ClassName = "ast.LogicalOrNode"
  x.Location = self.Location
  x.Left = self.Left
  x.Right = self.Right
  return json.Marshal(x)
}

func (self LogicalOrNode) IsExpr() bool {
  return true
}

func (self LogicalOrNode) GetLocation() duck.ILocation {
  return self.Location
}

// PrefixOpNode
type PrefixOpNode struct {
  Location duck.ILocation
  Operator string
  Expr duck.IExprNode
}

func NewPrefixOpNode(location duck.ILocation, operator string, expr duck.IExprNode) PrefixOpNode {
  return PrefixOpNode { location, operator, expr }
}

func (self PrefixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ 1 %s)", self.Expr)
    case "--": return fmt.Sprintf("(- 1 %s)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
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
  x.Location = self.Location
  x.Operator = self.Operator
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self PrefixOpNode) IsExpr() bool {
  return true
}

func (self PrefixOpNode) GetLocation() duck.ILocation {
  return self.Location
}

// SuffixOpNode
type SuffixOpNode struct {
  Location duck.ILocation
  Operator string
  Expr duck.IExprNode
}

func NewSuffixOpNode(location duck.ILocation, operator string, expr duck.IExprNode) SuffixOpNode {
  return SuffixOpNode { location, operator, expr }
}

func (self SuffixOpNode) String() string {
  switch self.Operator {
    case "++": return fmt.Sprintf("(+ %s 1)", self.Expr)
    case "--": return fmt.Sprintf("(- %s 1)", self.Expr)
    default:   return fmt.Sprintf("(%s %s)", self.Operator, self.Expr)
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
  x.Location = self.Location
  x.Operator = self.Operator
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self SuffixOpNode) IsExpr() bool {
  return true
}

func (self SuffixOpNode) GetLocation() duck.ILocation {
  return self.Location
}

// UnaryOpNode
type UnaryOpNode struct {
  Location duck.ILocation
  Operator string
  Expr duck.IExprNode
}

func NewUnaryOpNode(location duck.ILocation, operator string, expr duck.IExprNode) UnaryOpNode {
  return UnaryOpNode { location, operator, expr }
}

func (self UnaryOpNode) String() string {
  switch self.Operator {
    case "!": return fmt.Sprintf("(not %s)", self.Expr)
    default:  return fmt.Sprintf("%s%s", self.Operator, self.Expr)
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
  x.Location = self.Location
  x.Operator = self.Operator
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self UnaryOpNode) IsExpr() bool {
  return true
}

func (self UnaryOpNode) GetLocation() duck.ILocation {
  return self.Location
}
