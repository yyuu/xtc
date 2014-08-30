package ast

import (
  "strconv"
)

// IntegerLiteralNode
type IntegerLiteralNode struct {
  Location Location
  Value int
}

func NewIntegerLiteralNode(location Location, literal string) IntegerLiteralNode {
  value, err := strconv.Atoi(literal)
  if err != nil { panic(err) }
  return IntegerLiteralNode { location, value }
}

func (self IntegerLiteralNode) String() string {
  return strconv.Itoa(self.Value)
}

func (self IntegerLiteralNode) IsExpr() bool {
  return true
}

func (self IntegerLiteralNode) GetLocation() Location {
  return self.Location
}

// StringLiteralNode
type StringLiteralNode struct {
  Location Location
  Value string
}

func NewStringLiteralNode(location Location, literal string) StringLiteralNode {
  return StringLiteralNode { location, literal }
}

func (self StringLiteralNode) String() string {
  return self.Value
}

func (self StringLiteralNode) IsExpr() bool {
  return true
}

func (self StringLiteralNode) GetLocation() Location {
  return self.Location
}
