package ast

import (
  "testing"
)

func TestDecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345")
  s := `{
  "Location": "[:0,0]",
  "Value": 12345
}`
  assertJsonEquals(t, x, s)
}

func TestOctalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0755")
  s := `{
  "Location": "[:0,0]",
  "Value": 493
}`
  assertJsonEquals(t, x, s)
}

func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0xDEADBEEF")
  s := `{
  "Location": "[:0,0]",
  "Value": 3735928559
}`
  assertJsonEquals(t, x, s)
}

func TestCharacterIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "'a'")
  s := `{
  "Location": "[:0,0]",
  "Value": 97
}`
  assertJsonEquals(t, x, s)
}

func TestStringLiteral1(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "\"hello, world\"")
  s := `{
  "Location": "[:0,0]",
  "Value": "\"hello, world\""
}`
  assertJsonEquals(t, x, s)
}
