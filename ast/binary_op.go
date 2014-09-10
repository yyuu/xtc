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
