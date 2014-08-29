package ast

import (
  "testing"
)

func TestBinaryOp(t *testing.T) {
  x := BinaryOpNode(LOC, "*", BinaryOpNode(LOC, "%", VariableNode(LOC, "a"), VariableNode(LOC, "b")), VariableNode(LOC, "c"))
  assertEquals(t, x.String(), "(* (modulo a b) c)")
}

func TestLogicalAndNode(t *testing.T) {
  x := LogicalAndNode(LOC, VariableNode(LOC, "a"), LogicalAndNode(LOC, VariableNode(LOC, "b"), VariableNode(LOC, "c")))
  assertEquals(t, x.String(), "(and a (and b c))")
}

func TestLogicalOrNode(t *testing.T) {
  x := LogicalOrNode(LOC, LogicalOrNode(LOC, VariableNode(LOC, "a"), VariableNode(LOC, "b")), VariableNode(LOC, "c"))
  assertEquals(t, x.String(), "(or (or a b) c)")
}

func TestPrefixOpNode(t *testing.T) {
  x := PrefixOpNode(LOC, "--", VariableNode(LOC, "a"))
  assertEquals(t, x.String(), "(- 1 a)")
}

func TestSuffixOpNode(t *testing.T) {
  x := SuffixOpNode(LOC, "++", VariableNode(LOC, "a"))
  assertEquals(t, x.String(), "(+ a 1)")
}

func TestUnaryOpNode(t *testing.T) {
  x := UnaryOpNode(LOC, "-", IntegerLiteralNode(LOC, "12345"))
  assertEquals(t, x.String(), "-12345")

  y := UnaryOpNode(LOC, "!", VariableNode(LOC, "a"))
  assertEquals(t, y.String(), "(not a)")
}
