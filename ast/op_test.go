package ast

import (
  "testing"
)

func TestBinaryOp(t *testing.T) {
  x := NewBinaryOpNode(loc(0,0), "*", NewBinaryOpNode(loc(0,0), "%", NewVariableNode(loc(0,0), "a"), NewVariableNode(loc(0,0), "b")), NewVariableNode(loc(0,0), "c"))
  assertEquals(t, jsonString(x), "(* (modulo a b) c)")
}

func TestLogicalAndNode(t *testing.T) {
  x := NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "b"), NewVariableNode(loc(0,0), "c")))
  assertEquals(t, jsonString(x), "(and a (and b c))")
}

func TestLogicalOrNode(t *testing.T) {
  x := NewLogicalOrNode(loc(0,0), NewLogicalOrNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewVariableNode(loc(0,0), "b")), NewVariableNode(loc(0,0), "c"))
  assertEquals(t, jsonString(x), "(or (or a b) c)")
}

func TestPrefixOpNode(t *testing.T) {
  x := NewPrefixOpNode(loc(0,0), "--", NewVariableNode(loc(0,0), "a"))
  assertEquals(t, jsonString(x), "(- 1 a)")
}

func TestSuffixOpNode(t *testing.T) {
  x := NewSuffixOpNode(loc(0,0), "++", NewVariableNode(loc(0,0), "a"))
  assertEquals(t, jsonString(x), "(+ a 1)")
}

func TestUnaryOpNode(t *testing.T) {
  x := NewUnaryOpNode(loc(0,0), "-", NewIntegerLiteralNode(loc(0,0), "12345"))
  assertEquals(t, jsonString(x), "-12345")

  y := NewUnaryOpNode(loc(0,0), "!", NewVariableNode(loc(0,0), "a"))
  assertEquals(t, jsonString(y), "(not a)")
}
