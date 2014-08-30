package ast

import (
  "testing"
)

/*
func TestAddressNode(t *testing.T) {
}
 */

func TestArefNode(t *testing.T) {
  x := NewArefNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "Location": "[:0,0]",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Index": {
    "Location": "[:0,0]",
    "Value": 12345
  }
}`
  assertJsonEquals(t, x, s)
}

/*
func TestDereferenceNode(t *testing.T) {
}
 */

func TestFuncallNode(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "a"), []IExprNode { NewIntegerLiteralNode(loc(0,0), "12345"), NewIntegerLiteralNode(loc(0,0), "67890") })
  s := `{
  "Location": "[:0,0]",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Args": [
    {
      "Location": "[:0,0]",
      "Value": 12345
    },
    {
      "Location": "[:0,0]",
      "Value": 67890
    }
  ]
}`
  assertJsonEquals(t, x, s)
}

func TestFuncallNode2(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []IExprNode { })
  s := `{
  "Location": "[:0,0]",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "b"
  },
  "Args": []
}`
  assertJsonEquals(t, x, s)
}

func TestMemberNode(t *testing.T) {
  x := NewMemberNode(loc(0,0), NewVariableNode(loc(0,0), "a"), "b")
  s := `{
  "Location": "[:0,0]",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Member": "b"
}`
  assertJsonEquals(t, x, s)
}

func TestPtrMemberNode(t *testing.T) {
  x := NewPtrMemberNode(loc(0,0), NewVariableNode(loc(0,0), "a"), "b")
  s := `{
  "Location": "[:0,0]",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Member": "b"
}`
  assertJsonEquals(t, x, s)
}

func TestVariableNode(t *testing.T) {
  x := NewVariableNode(loc(0,0), "a")
  s := `{
  "Location": "[:0,0]",
  "Name": "a"
}`
  assertJsonEquals(t, x, s)
}
