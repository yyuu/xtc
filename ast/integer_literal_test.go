package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestDecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "int",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "DecimalIntegerLiteralNode", xt.JSON(x), s)
}

func TestCharacterIntegerLiteral(t *testing.T) {
  x := NewCharacterLiteralNode(loc(0,0), "97")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "char",
    "Type": null
  },
  "Value": 97
}`
  xt.AssertStringEqualsDiff(t, "CharacterIntegerLiteralNode", xt.JSON(x), s)
}

func TestUnsignedIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345U")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "unsigned int",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "unsigned literal", xt.JSON(x), s)
}

func TestLongIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345L")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "long",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "long literal", xt.JSON(x), s)
}

func TestUnsignedLongIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345UL")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "TypeNode": {
    "ClassName": "ast.TypeNode",
    "Location": "[:0,0]",
    "TypeRef": "unsigned long",
    "Type": null
  },
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "unsigned long literal", xt.JSON(x), s)
}
