package ast

import (
  "testing"
)

func TestBreak(t *testing.T) {
  x := BreakNode()
  assertEquals(t, x.String(), "(break)")
}

func TestContinue(t *testing.T) {
  x := ContinueNode()
  assertEquals(t, x.String(), "(continue)")
}

func TestExprStmt(t *testing.T) {
  x := ExprStmtNode(BinaryOpNode("+", IntegerLiteralNode("1"), IntegerLiteralNode("1")))
  assertEquals(t, x.String(), "(+ 1 1)")
}

func TestGoto(t *testing.T) {
  x := GotoNode("a")
  assertEquals(t, x.String(), "(goto a)")
}

/*
func TestLabel(t *testing.T) {
}
 */

func TestReturn(t *testing.T) {
  x := ReturnNode(VariableNode("a"))
  assertEquals(t, x.String(), "a")
}
