package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestContinue(t *testing.T) {
  x := NewContinueNode(loc(0,0))
  s := `{
  "ClassName": "ast.ContinueNode",
  "Location": "[:0,0]"
}`
  xt.AssertStringEqualsDiff(t, "ContinueNode", xt.JSON(x), s)
}
