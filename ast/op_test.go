package ast

import (
  "testing"
)

func TestBinaryOp(t *testing.T) {
  x := NewBinaryOpNode(LOC, "*", NewBinaryOpNode(LOC, "%", NewVariableNode(LOC, "a"), NewVariableNode(LOC, "b")), NewVariableNode(LOC, "c"))
  assertEquals(t, x.String(), "(* (modulo a b) c)")
}

func TestLogicalAndNode(t *testing.T) {
  x := NewLogicalAndNode(LOC, NewVariableNode(LOC, "a"), NewLogicalAndNode(LOC, NewVariableNode(LOC, "b"), NewVariableNode(LOC, "c")))
  assertEquals(t, x.String(), "(and a (and b c))")
}

func TestLogicalOrNode(t *testing.T) {
  x := NewLogicalOrNode(LOC, NewLogicalOrNode(LOC, NewVariableNode(LOC, "a"), NewVariableNode(LOC, "b")), NewVariableNode(LOC, "c"))
  assertEquals(t, x.String(), "(or (or a b) c)")
}

func TestPrefixOpNode(t *testing.T) {
  x := NewPrefixOpNode(LOC, "--", NewVariableNode(LOC, "a"))
  assertEquals(t, x.String(), "(- 1 a)")
}

func TestSuffixOpNode(t *testing.T) {
  x := NewSuffixOpNode(LOC, "++", NewVariableNode(LOC, "a"))
  assertEquals(t, x.String(), "(+ a 1)")
}

func TestUnaryOpNode(t *testing.T) {
  x := NewUnaryOpNode(LOC, "-", NewIntegerLiteralNode(LOC, "12345"))
  assertEquals(t, x.String(), "-12345")

  y := NewUnaryOpNode(LOC, "!", NewVariableNode(LOC, "a"))
  assertEquals(t, y.String(), "(not a)")
}
