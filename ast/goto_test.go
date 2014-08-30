package ast

import (
  "testing"
)

func TestBreak(t *testing.T) {
  x := NewBreakNode(LOC)
  assertEquals(t, x.String(), "(break)")
}

func TestContinue(t *testing.T) {
  x := NewContinueNode(LOC)
  assertEquals(t, x.String(), "(continue)")
}

func TestExprStmt(t *testing.T) {
  x := NewExprStmtNode(LOC, NewBinaryOpNode(LOC, "+", NewIntegerLiteralNode(LOC, "1"), NewIntegerLiteralNode(LOC, "1")))
  assertEquals(t, x.String(), "(+ 1 1)")
}

func TestGoto(t *testing.T) {
  x := NewGotoNode(LOC, "a")
  assertEquals(t, x.String(), "(goto a)")
}

/*
func TestLabel(t *testing.T) {
}
 */

func TestReturn(t *testing.T) {
  x := NewReturnNode(LOC, NewVariableNode(LOC, "a"))
  assertEquals(t, x.String(), "a")
}
