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
  t core.IType
}

func NewBinaryOpNode(loc core.Location, operator string, left core.IExprNode, right core.IExprNode) *BinaryOpNode {
  if left == nil { panic("left is nil") }
  if right == nil { panic("right is nil") }
  return &BinaryOpNode { "ast.BinaryOpNode", loc, operator, left, right, nil }
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

func (self BinaryOpNode) GetOperator() string {
  return self.Operator
}

func (self BinaryOpNode) GetLeft() core.IExprNode {
  return self.Left
}

func (self *BinaryOpNode) SetLeft(expr core.IExprNode) {
  self.Left = expr
}

func (self BinaryOpNode) GetRight() core.IExprNode {
  return self.Right
}

func (self *BinaryOpNode) SetRight(expr core.IExprNode) {
  self.Right = expr
}

func (self BinaryOpNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *BinaryOpNode) SetType(t core.IType) {
  self.t = t
}

func (self BinaryOpNode) IsConstant() bool {
  return false
}

func (self BinaryOpNode) IsParameter() bool {
  return false
}

func (self BinaryOpNode) IsLvalue() bool {
  return false
}

func (self BinaryOpNode) IsAssignable() bool {
  return false
}

func (self BinaryOpNode) IsLoadable() bool {
  return false
}

func (self BinaryOpNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self BinaryOpNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
