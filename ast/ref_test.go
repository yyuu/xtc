package ast

import (
  "testing"
)

/*
func TestAddressNode(t *testing.T) {
}
 */

func TestArefNode(t *testing.T) {
  x := NewArefNode(LOC, NewVariableNode(LOC, "a"), NewIntegerLiteralNode(LOC, "12345"))
  assertEquals(t, x.String(), "(vector-ref a 12345)")
}

/*
func TestDereferenceNode(t *testing.T) {
}
 */

func TestFuncallNode(t *testing.T) {
  x := NewFuncallNode(LOC, NewVariableNode(LOC, "a"), []IExprNode { NewIntegerLiteralNode(LOC, "12345"), NewIntegerLiteralNode(LOC, "67890") })
  assertEquals(t, x.String(), "(a 12345 67890)")

  y := NewFuncallNode(LOC, NewVariableNode(LOC, "b"), []IExprNode { })
  assertEquals(t, y.String(), "(b)")
}

func TestMemberNode(t *testing.T) {
  x := NewMemberNode(LOC, NewVariableNode(LOC, "a"), "b")
  assertEquals(t, x.String(), "(slot-ref a 'b)")
}

func TestPtrMemberNode(t *testing.T) {
  x := NewPtrMemberNode(LOC, NewVariableNode(LOC, "a"), "b")
  assertEquals(t, x.String(), "(slot-ref a 'b)")
}

func TestVariableNode(t *testing.T) {
  x := NewVariableNode(LOC, "a")
  assertEquals(t, x.String(), "a")
}
