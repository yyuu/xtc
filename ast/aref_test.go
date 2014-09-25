package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestArefNode(t *testing.T) {
  x := NewArefNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "ClassName": "ast.ArefNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a",
    "Entity": null
  },
  "Index": {
    "ClassName": "ast.IntegerLiteralNode",
    "Location": "[:0,0]",
    "TypeNode": {
      "ClassName": "ast.TypeNode",
      "Location": "[:0,0]",
      "TypeRef": "int",
      "Type": null
    },
    "Value": 12345
  },
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "ArefNode", xt.JSON(x), s)
}
