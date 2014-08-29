package ast

import (
  "testing"
)

func TestAssignNode(t *testing.T) {
  x := AssignNode(VariableNode("a"), StringLiteralNode("\"xxx\""))
  assertEquals(t, x.String(), "(a \"xxx\")")
}

func TestOpAssignNode(t *testing.T) {
  x := OpAssignNode("+", VariableNode("a"), IntegerLiteralNode("12345"))
  assertEquals(t, x.String(), "(a (+ a 12345))")
}
