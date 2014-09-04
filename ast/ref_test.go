package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/duck"
)

/*
func TestAddressNode(t *testing.T) {
}
 */

func TestArefNode(t *testing.T) {
  x := NewArefNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "ClassName": "ast.ArefNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Index": {
    "ClassName": "ast.IntegerLiteralNode",
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
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "a"), []duck.IExprNode { NewIntegerLiteralNode(loc(0,0), "12345"), NewIntegerLiteralNode(loc(0,0), "67890") })
  s := `{
  "ClassName": "ast.FuncallNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Args": [
    {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 12345
    },
    {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 67890
    }
  ]
}`
  assertJsonEquals(t, x, s)
}

func TestFuncallNode2(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []duck.IExprNode { })
  s := `{
  "ClassName": "ast.FuncallNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
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
  "ClassName": "ast.MemberNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
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
  "ClassName": "ast.PtrMemberNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
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
  "ClassName": "ast.VariableNode",
  "Location": "[:0,0]",
  "Name": "a"
}`
  assertJsonEquals(t, x, s)
}
