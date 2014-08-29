package ast

import (
  "strconv"
)

// IntegerLiteralNode
type integerLiteralNode struct {
  location ILocation
  Value int
}

func IntegerLiteralNode(location ILocation, literal string) integerLiteralNode {
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

func (self integerLiteralNode) Location() ILocation {
  return self.location
}

// StringLiteralNode
type stringLiteralNode struct {
  location ILocation
  Value string
}

func StringLiteralNode(location ILocation, literal string) stringLiteralNode {
  return stringLiteralNode { location, literal }
}

func (self stringLiteralNode) String() string {
  return self.Value
}

func (self stringLiteralNode) IsExpr() bool {
  return true
}

func (self stringLiteralNode) Location() ILocation {
  return self.location
}
