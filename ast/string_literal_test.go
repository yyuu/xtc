package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestStringLiteral1(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "\"hello, world\"")
  s := `{
  "ClassName": "ast.StringLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "char*",
    "Type": null
  },
  "Value": "\"hello, world\""
}`
  xt.AssertStringEqualsDiff(t, "StringLiteralNode", xt.JSON(x), s)
}
