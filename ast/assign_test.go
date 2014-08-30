package ast

import (
  "testing"
)

func TestAssignNode(t *testing.T) {
  x := NewAssignNode(LOC, NewVariableNode(LOC, "a"), NewStringLiteralNode(LOC, "\"xxx\""))
  assertEquals(t, x.String(), "(a \"xxx\")")
}

func TestOpAssignNode(t *testing.T) {
  x := NewOpAssignNode(LOC, "+", NewVariableNode(LOC, "a"), NewIntegerLiteralNode(LOC, "12345"))
  assertEquals(t, x.String(), "(a (+ a 12345))")
}
