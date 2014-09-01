package ast

import (
  "testing"
)

func TestDecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 12345
}`
  assertJsonEquals(t, x, s)
}

func TestOctalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0755")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 493
}`
  assertJsonEquals(t, x, s)
}

func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0xFFFF")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 65535
}`
  assertJsonEquals(t, x, s)
}

func TestCharacterIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "'a'")
  s := `{
  "ClassName": "ast.IntegerLiteralNode",
  "Location": "[:0,0]",
  "Value": 97
}`
  assertJsonEquals(t, x, s)
}

func TestStringLiteral1(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "\"hello, world\"")
  s := `{
  "ClassName": "ast.StringLiteralNode",
  "Location": "[:0,0]",
  "Value": "\"hello, world\""
}`
  assertJsonEquals(t, x, s)
}
