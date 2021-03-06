package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestVariableNode(t *testing.T) {
  x := NewVariableNode(loc(0,0), "a")
  s := `{
  "ClassName": "ast.VariableNode",
  "Location": "[:0,0]",
  "Name": "a",
  "Entity": null
}`
  xt.AssertStringEqualsDiff(t, "VariableNode", xt.JSON(x), s)
}
