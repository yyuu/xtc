package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestAssignNode(t *testing.T) {
  x := NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewStringLiteralNode(loc(0,0), "\"xxx\""))
  s := `{
  "ClassName": "ast.AssignNode",
  "Location": "[:0,0]",
  "Lhs": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Rhs": {
    "ClassName": "ast.StringLiteralNode",
    "Location": "[:0,0]",
    "Value": "\"xxx\""
  }
}`
  xt.AssertStringEqualsDiff(t, "AssignNode", xt.JSON(x), s)
}

func TestOpAssignNode(t *testing.T) {
  x := NewOpAssignNode(loc(0,0), "+", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "ClassName": "ast.OpAssignNode",
  "Location": "[:0,0]",
  "Operator": "+",
  "Lhs": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Rhs": {
    "ClassName": "ast.IntegerLiteralNode",
    "Location": "[:0,0]",
    "Value": 12345
  }
}`
  xt.AssertStringEqualsDiff(t, "OpAssignNode", xt.JSON(x), s)
}
