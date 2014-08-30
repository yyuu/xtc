package ast

import (
  "strconv"
)

// IntegerLiteralNode
type integerLiteralNode struct {
  location Location
  Value int
}

func NewIntegerLiteralNode(location Location, literal string) integerLiteralNode {
  value, err := strconv.Atoi(literal)
  if err != nil { panic(err) }
  return integerLiteralNode { location, value }
}

func (self integerLiteralNode) String() string {
  return strconv.Itoa(self.Value)
}

func (self integerLiteralNode) IsExpr() bool {
  return true
}

func (self integerLiteralNode) GetLocation() Location {
  return self.location
}

// StringLiteralNode
type stringLiteralNode struct {
  location Location
  Value string
}

func NewStringLiteralNode(location Location, literal string) stringLiteralNode {
  return stringLiteralNode { location, literal }
}

func (self stringLiteralNode) String() string {
  return self.Value
}

func (self stringLiteralNode) IsExpr() bool {
  return true
}

func (self stringLiteralNode) GetLocation() Location {
  return self.location
}
