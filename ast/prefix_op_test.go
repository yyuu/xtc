package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
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
    "Name": "a",
    "Entity": null
  },
  "Amount": 1,
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "PrefixOpNode", xt.JSON(x), s)
}
