package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
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
  xt.AssertStringEqualsDiff(t, "ArefNode", xt.JSON(x), s)
}

/*
func TestDereferenceNode(t *testing.T) {
}
 */

func TestFuncallNode(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "a"), []core.IExprNode { NewIntegerLiteralNode(loc(0,0), "12345"), NewIntegerLiteralNode(loc(0,0), "67890") })
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
  xt.AssertStringEqualsDiff(t, "FuncallNode1", xt.JSON(x), s)
}

func TestFuncallNode2(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []core.IExprNode { })
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
  xt.AssertStringEqualsDiff(t, "FuncallNode2", xt.JSON(x), s)
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
  xt.AssertStringEqualsDiff(t, "MemberNode", xt.JSON(x), s)
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
  xt.AssertStringEqualsDiff(t, "PtrMemberNode", xt.JSON(x), s)
}

func TestVariableNode(t *testing.T) {
  x := NewVariableNode(loc(0,0), "a")
  s := `{
  "ClassName": "ast.VariableNode",
  "Location": "[:0,0]",
  "Name": "a"
}`
  xt.AssertStringEqualsDiff(t, "VariableNode", xt.JSON(x), s)
}
