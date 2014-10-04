package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestBreak(t *testing.T) {
  x := NewBreakNode(loc(0,0))
  s := `{
  "ClassName": "ast.BreakNode",
  "Location": "[:0,0]"
}`
  xt.AssertStringEqualsDiff(t, "BreakNode", xt.JSON(x), s)
}
