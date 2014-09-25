package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestPtrMemberNode(t *testing.T) {
  x := NewPtrMemberNode(loc(0,0), NewVariableNode(loc(0,0), "a"), "b")
  s := `{
  "ClassName": "ast.PtrMemberNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a",
    "Entity": null
  },
  "Member": "b",
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "PtrMemberNode", xt.JSON(x), s)
}
