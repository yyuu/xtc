package ast

import (
  "strconv"
)

// IntegerLiteralNode
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

func (self integerLiteralNode) IsExpr() bool {
  return true
}

// StringLiteralNode
type stringLiteralNode struct {
  Value string
}

func StringLiteralNode(literal string) stringLiteralNode {
  return stringLiteralNode { literal }
}

func (self stringLiteralNode) String() string {
  return self.Value
}

func (self stringLiteralNode) IsExpr() bool {
  return true
}
