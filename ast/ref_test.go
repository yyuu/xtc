package ast

import (
  "testing"
)

/*
func TestAddressNode(t *testing.T) {
}
 */

func TestArefNode(t *testing.T) {
  x := ArefNode(LOC, VariableNode(LOC, "a"), IntegerLiteralNode(LOC, "12345"))
  assertEquals(t, x.String(), "(vector-ref a 12345)")
}

/*
func TestDereferenceNode(t *testing.T) {
}
 */

func TestFuncallNode(t *testing.T) {
  x := FuncallNode(LOC, VariableNode(LOC, "a"), []IExprNode { IntegerLiteralNode(LOC, "12345"), IntegerLiteralNode(LOC, "67890") })
  assertEquals(t, x.String(), "(a 12345 67890)")

  y := FuncallNode(LOC, VariableNode(LOC, "b"), []IExprNode { })
  assertEquals(t, y.String(), "(b)")
}

func TestMemberNode(t *testing.T) {
  x := MemberNode(LOC, VariableNode(LOC, "a"), "b")
  assertEquals(t, x.String(), "(slot-ref a 'b)")
}

func TestPtrMemberNode(t *testing.T) {
  x := PtrMemberNode(LOC, VariableNode(LOC, "a"), "b")
  assertEquals(t, x.String(), "(slot-ref a 'b)")
}

func TestVariableNode(t *testing.T) {
  x := VariableNode(LOC, "a")
  assertEquals(t, x.String(), "a")
}
