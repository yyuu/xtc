package ast

import (
  "testing"
)

/*
func TestAddressNode(t *testing.T) {
}
 */

func TestArefNode(t *testing.T) {
  x := ArefNode(VariableNode("a"), IntegerLiteralNode("12345"))
  assertEquals(t, x.String(), "(vector-ref a 12345)")
}

/*
func TestDereferenceNode(t *testing.T) {
}
 */

func TestFuncallNode(t *testing.T) {
  x := FuncallNode(VariableNode("a"), []INode { IntegerLiteralNode("12345"), IntegerLiteralNode("67890") })
  assertEquals(t, x.String(), "(a 12345 67890)")

  y := FuncallNode(VariableNode("b"), []INode { })
  assertEquals(t, y.String(), "(b)")
}

func TestMemberNode(t *testing.T) {
  x := MemberNode(VariableNode("a"), "b")
  assertEquals(t, x.String(), "(slot-ref a 'b)")
}

func TestPtrMemberNode(t *testing.T) {
  x := PtrMemberNode(VariableNode("a"), "b")
  assertEquals(t, x.String(), "(slot-ref a 'b)")
}

func TestVariableNode(t *testing.T) {
  x := VariableNode("a")
  assertEquals(t, x.String(), "a")
}
