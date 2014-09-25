package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestPrefixOpNode(t *testing.T) {
  x := NewPrefixOpNode(loc(0,0), "--", NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.PrefixOpNode",
  "Location": "[:0,0]",
  "Operator": "--",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "PrefixOpNode", xt.JSON(x), s)
}
