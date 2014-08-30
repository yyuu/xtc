package ast

import (
  "testing"
)

func TestBreak(t *testing.T) {
  x := NewBreakNode(loc(0,0))
  assertEquals(t, jsonString(x), "(break)")
}

func TestContinue(t *testing.T) {
  x := NewContinueNode(loc(0,0))
  assertEquals(t, jsonString(x), "(continue)")
}

func TestExprStmt(t *testing.T) {
  x := NewExprStmtNode(loc(0,0), NewBinaryOpNode(loc(0,0), "+", NewIntegerLiteralNode(loc(0,0), "1"), NewIntegerLiteralNode(loc(0,0), "1")))
  assertEquals(t, jsonString(x), "(+ 1 1)")
}

func TestGoto(t *testing.T) {
  x := NewGotoNode(loc(0,0), "a")
  assertEquals(t, jsonString(x), "(goto a)")
}

/*
func TestLabel(t *testing.T) {
}
 */

func TestReturn(t *testing.T) {
  x := NewReturnNode(loc(0,0), NewVariableNode(loc(0,0), "a"))
  assertEquals(t, jsonString(x), "a")
}
