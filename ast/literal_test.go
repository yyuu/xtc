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

/*
func TestOctalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0755")
  s := `
`
  assertJsonEquals(t, x.Value, s)
}
 */

/*
func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0xDEADBEEF")
  s := `
`
  assertJsonEquals(t, x.Value, s)
}
 */

func TestStringLiteral1(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "\"hello, world\"")
  s := `{
  "Location": "[:0,0]",
  "Value": "\"hello, world\""
}`
  assertJsonEquals(t, x, s)
}
