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
