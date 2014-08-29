package ast

import (
  "testing"
)

func TestBreak(t *testing.T) {
  x := BreakNode(LOC)
  assertEquals(t, x.String(), "(break)")
}

func TestContinue(t *testing.T) {
  x := ContinueNode(LOC)
  assertEquals(t, x.String(), "(continue)")
}

func TestExprStmt(t *testing.T) {
  x := ExprStmtNode(LOC, BinaryOpNode(LOC, "+", IntegerLiteralNode(LOC, "1"), IntegerLiteralNode(LOC, "1")))
  assertEquals(t, x.String(), "(+ 1 1)")
}

func TestGoto(t *testing.T) {
  x := GotoNode(LOC, "a")
  assertEquals(t, x.String(), "(goto a)")
}

/*
func TestLabel(t *testing.T) {
}
 */

func TestReturn(t *testing.T) {
  x := ReturnNode(LOC, VariableNode(LOC, "a"))
  assertEquals(t, x.String(), "a")
}
