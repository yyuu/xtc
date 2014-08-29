package ast

import (
  "testing"
)

func TestBinaryOp(t *testing.T) {
  x := BinaryOpNode("*", BinaryOpNode("%", VariableNode("a"), VariableNode("b")), VariableNode("c"))
  assertEquals(t, x.String(), "(* (modulo a b) c)")
}

func TestCondExpr(t *testing.T) {
  x := CondExprNode(
    BinaryOpNode("<", VariableNode("n"), IntegerLiteralNode("2")),
    IntegerLiteralNode("1"),
    BinaryOpNode("+",
                 FuncallNode(VariableNode("f"), []INode { BinaryOpNode("-", VariableNode("n"), IntegerLiteralNode("1")) }),
                 FuncallNode(VariableNode("f"), []INode { BinaryOpNode("-", VariableNode("n"), IntegerLiteralNode("2")) })))
  assertEquals(t, x.String(), "(if (< n 2) 1 (+ (f (- n 1)) (f (- n 2))))")
}

func TestLogicalAndNode(t *testing.T) {
  x := LogicalAndNode(VariableNode("a"), LogicalAndNode(VariableNode("b"), VariableNode("c")))
  assertEquals(t, x.String(), "(and a (and b c))")
}

func TestLogicalOrNode(t *testing.T) {
  x := LogicalOrNode(LogicalOrNode(VariableNode("a"), VariableNode("b")), VariableNode("c"))
  assertEquals(t, x.String(), "(or (or a b) c)")
}

func TestPrefixOpNode(t *testing.T) {
  x := PrefixOpNode("--", VariableNode("a"))
  assertEquals(t, x.String(), "(- 1 a)")
}

func TestSuffixOpNode(t *testing.T) {
  x := SuffixOpNode("++", VariableNode("a"))
  assertEquals(t, x.String(), "(+ a 1)")
}

func TestUnaryOpNode(t *testing.T) {
  x := UnaryOpNode("-", IntegerLiteralNode("12345"))
  assertEquals(t, x.String(), "-12345")

  y := UnaryOpNode("!", VariableNode("a"))
  assertEquals(t, y.String(), "(not a)")
}
