package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestDecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 12345
}`
  xt.AssertStringEqualsDiff(t, "DecimalIntegerLiteralNode", xt.JSON(x), s)
}

func TestOctalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0755")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 493
}`
  xt.AssertStringEqualsDiff(t, "OctalIntegerLiteralNode", xt.JSON(x), s)
}

func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0xFFFF")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 65535
}`
  xt.AssertStringEqualsDiff(t, "HexadecimalIntegerLiteralNode", xt.JSON(x), s)
}

func TestCharacterIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "'a'")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 97
}`
  xt.AssertStringEqualsDiff(t, "CharacterIntegerLiteralNode", xt.JSON(x), s)
}
