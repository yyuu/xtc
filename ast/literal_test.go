package ast

import (
  "testing"
)

func TestDecimalIntegerLiteral(t *testing.T) {
  x := IntegerLiteralNode(LOC, "12345")
  assertEquals(t, x.Value, 12345)
  assertEquals(t, x.String(), "12345")
}

/*
func TestOctalIntegerLiteral(t *testing.T) {
  x := IntegerLiteralNode(LOC, "0755")
  assertEquals(t, x.Value, 493)
}
 */

/*
func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := IntegerLiteralNode(LOC, "0xDEADBEEF")
  assertEquals(t, x.Value, 3735928559)
}
 */

func TestStringLiteral1(t *testing.T) {
  x := StringLiteralNode(LOC, "\"hello, world\"")
  assertEquals(t, x.Value, "\"hello, world\"")
  assertEquals(t, x.String(), "\"hello, world\"")
}
