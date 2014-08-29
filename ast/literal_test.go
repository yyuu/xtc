package ast

import (
  "testing"
)

func assertEquals(t *testing.T, got interface{}, expected interface{}) {
  if got != expected {
    t.Errorf("not equals: expected %q, got %q", expected, got)
  }
}

func TestDecimalIntegerLiteral(t *testing.T) {
  x := IntegerLiteralNode("12345")
  assertEquals(t, x.Value, 12345)
  assertEquals(t, x.String(), "12345")
}

/*
func TestOctalIntegerLiteral(t *testing.T) {
  x := IntegerLiteralNode("0755")
  assertEquals(t, x.Value, 493)
}
 */

/*
func TestHexadecimalIntegerLiteral(t *testing.T) {
  x := IntegerLiteralNode("0xDEADBEEF")
  assertEquals(t, x.Value, 3735928559)
}
 */

func TestStringLiteral1(t *testing.T) {
  x := StringLiteralNode("\"hello, world\"")
  assertEquals(t, x.Value, "\"hello, world\"")
  assertEquals(t, x.String(), "\"hello, world\"")
}
