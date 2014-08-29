package ast

import (
  "strconv"
)

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

type stringLiteralNode struct {
  Value string
}

func StringLiteralNode(literal string) stringLiteralNode {
  return stringLiteralNode { literal }
}

func (self stringLiteralNode) String() string {
  return self.Value
}
