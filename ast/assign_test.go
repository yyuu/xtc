package ast

import (
  "testing"
)

func TestAssignNode(t *testing.T) {
  x := AssignNode(LOC, VariableNode(LOC, "a"), StringLiteralNode(LOC, "\"xxx\""))
  assertEquals(t, x.String(), "(a \"xxx\")")
}

func TestOpAssignNode(t *testing.T) {
  x := OpAssignNode(LOC, "+", VariableNode(LOC, "a"), IntegerLiteralNode(LOC, "12345"))
  assertEquals(t, x.String(), "(a (+ a 12345))")
}
