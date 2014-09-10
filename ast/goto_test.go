package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestGoto(t *testing.T) {
  x := NewGotoNode(loc(0,0), "a")
  s := `{
  "ClassName": "ast.GotoNode",
  "Location": "[:0,0]",
  "Target": "a"
}`
  xt.AssertStringEqualsDiff(t, "GotoNode", xt.JSON(x), s)
}
