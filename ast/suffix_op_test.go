package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestSuffixOpNode(t *testing.T) {
  x := NewSuffixOpNode(loc(0,0), "++", NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.SuffixOpNode",
  "Location": "[:0,0]",
  "Operator": "++",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a",
    "Entity": null
  },
  "Amount": 1,
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "SuffixOpNode", xt.JSON(x), s)
}
