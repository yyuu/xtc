package ast

import (
  "testing"
)

func TestAssignNode(t *testing.T) {
  x := NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewStringLiteralNode(loc(0,0), "\"xxx\""))
  s := `{
  "Location": "[:0,0]",
  "Lhs": {
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Rhs": {
    "Location": "[:0,0]",
    "Value": "\"xxx\""
  }
}`
  assertJsonEquals(t, x, s)
}

func TestOpAssignNode(t *testing.T) {
  x := NewOpAssignNode(loc(0,0), "+", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "Location": "[:0,0]",
  "Operator": "+",
  "Lhs": {
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Rhs": {
    "Location": "[:0,0]",
    "Value": 12345
  }
}`
  assertJsonEquals(t, x, s)
}
