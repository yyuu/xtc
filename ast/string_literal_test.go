package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestStringLiteral1(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "hello, world")
  s := `{
  "ClassName": "ast.StringLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "char*",
    "Type": null
  },
  "Value": "hello, world"
}`
  xt.AssertStringEqualsDiff(t, "string literal1", xt.JSON(x), s)
}

func TestStringLiteral2(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "foo\tbar\r\nbaz\r\n")
  s := `{
  "ClassName": "ast.StringLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "char*",
    "Type": null
  },
  "Value": "foo\tbar\r\nbaz\r\n"
}`
  xt.AssertStringEqualsDiff(t, "string literal2", xt.JSON(x), s)
}

func TestStringLiteral3(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "にほんご\n日本語")
  s := `{
  "ClassName": "ast.StringLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "char*",
    "Type": null
  },
  "Value": "にほんご\n日本語"
}`
  xt.AssertStringEqualsDiff(t, "string literal3", xt.JSON(x), s)
}
