package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestReturn(t *testing.T) {
  x := NewReturnNode(loc(0,0), NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.ReturnNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a",
    "Entity": null
  }
}`
  xt.AssertStringEqualsDiff(t, "VariableNode", xt.JSON(x), s)
}
