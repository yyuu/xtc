package ast

import (
  "testing"
)

/*
func TestAddressNode(t *testing.T) {
}
 */

func TestArefNode(t *testing.T) {
  x := NewArefNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "12345"))
  assertEquals(t, jsonString(x), "(vector-ref a 12345)")
}

/*
func TestDereferenceNode(t *testing.T) {
}
 */

func TestFuncallNode(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "a"), []IExprNode { NewIntegerLiteralNode(loc(0,0), "12345"), NewIntegerLiteralNode(loc(0,0), "67890") })
  assertEquals(t, jsonString(x), "(a 12345 67890)")

  y := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []IExprNode { })
  assertEquals(t, jsonString(y), "(b)")
}

func TestMemberNode(t *testing.T) {
  x := NewMemberNode(loc(0,0), NewVariableNode(loc(0,0), "a"), "b")
  assertEquals(t, jsonString(x), "(slot-ref a 'b)")
}

func TestPtrMemberNode(t *testing.T) {
  x := NewPtrMemberNode(loc(0,0), NewVariableNode(loc(0,0), "a"), "b")
  assertEquals(t, jsonString(x), "(slot-ref a 'b)")
}

func TestVariableNode(t *testing.T) {
  x := NewVariableNode(loc(0,0), "a")
  assertEquals(t, jsonString(x), "a")
}
