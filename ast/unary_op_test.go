package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestUnaryOpNode1(t *testing.T) {
  x := NewUnaryOpNode(loc(0,0), "-", NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "ClassName": "ast.UnaryOpNode",
  "Location": "[:0,0]",
  "Operator": "-",
  "Expr": {
    "ClassName": "ast.IntegerLiteralNode",
    "Location": "[:0,0]",
    "Value": 12345
  }
}`
  xt.AssertStringEqualsDiff(t, "UnaryOpNode1", xt.JSON(x), s)
}

func TestUnaryOpNode2(t *testing.T) {
  x := NewUnaryOpNode(loc(0,0), "!", NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.UnaryOpNode",
  "Location": "[:0,0]",
  "Operator": "!",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  xt.AssertStringEqualsDiff(t, "UnaryOpNode2", xt.JSON(x), s)
}
