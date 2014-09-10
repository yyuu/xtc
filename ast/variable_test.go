package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestVariableNode(t *testing.T) {
  x := NewVariableNode(loc(0,0), "a")
  s := `{
  "ClassName": "ast.VariableNode",
  "Location": "[:0,0]",
  "Name": "a"
}`
  xt.AssertStringEqualsDiff(t, "VariableNode", xt.JSON(x), s)
}
