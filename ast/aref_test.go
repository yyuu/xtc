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
    "Name": "a"
  },
  "Index": {
    "ClassName": "ast.IntegerLiteralNode",
    "Location": "[:0,0]",
    "TypeNode": {
      "ClassName": "ast.TypeNode",
      "Location": "[:0,0]",
      "TypeRef": {
        "ClassName": "typesys.IntegerTypeRef",
        "Location": "[:0,0]",
        "Name": "int"
      }
    },
    "Value": 12345
  }
}`
  xt.AssertStringEqualsDiff(t, "ArefNode", xt.JSON(x), s)
}
