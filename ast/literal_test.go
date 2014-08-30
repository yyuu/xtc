package ast

import (
  "testing"
)

func TestDecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "12345")
  assertEquals(t, x.Value, 12345)
  assertEquals(t, jsonString(x), "12345")
}

/*
func TestOctalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0755")
  assertEquals(t, x.Value, 493)
}
 */

/*
func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := NewIntegerLiteralNode(loc(0,0), "0xDEADBEEF")
  assertEquals(t, x.Value, 3735928559)
}
 */

func TestStringLiteral1(t *testing.T) {
  x := NewStringLiteralNode(loc(0,0), "\"hello, world\"")
  assertEquals(t, x.Value, "\"hello, world\"")
  assertEquals(t, jsonString(x), "\"hello, world\"")
}
