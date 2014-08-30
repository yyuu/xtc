package ast

import (
  "testing"
)

func TestBreak(t *testing.T) {
  x := NewBreakNode(loc(0,0))
  s := `{
  "Location": "[:0,0]"
}`
  assertJsonEquals(t, x, s)
}

func TestContinue(t *testing.T) {
  x := NewContinueNode(loc(0,0))
  s := `{
  "Location": "[:0,0]"
}`
  assertJsonEquals(t, x, s)
}

func TestExprStmt(t *testing.T) {
  x := NewExprStmtNode(loc(0,0), NewBinaryOpNode(loc(0,0), "+", NewIntegerLiteralNode(loc(0,0), "1"), NewIntegerLiteralNode(loc(0,0), "1")))
  s := `{
  "Location": "[:0,0]",
  "Expr": {
    "Location": "[:0,0]",
    "Operator": "+",
    "Left": {
      "Location": "[:0,0]",
      "Value": 1
    },
    "Right": {
      "Location": "[:0,0]",
      "Value": 1
    }
  }
}`
  assertJsonEquals(t, x, s)
}

func TestGoto(t *testing.T) {
  x := NewGotoNode(loc(0,0), "a")
  s := `{
  "Location": "[:0,0]",
  "Target": "a"
}`
  assertJsonEquals(t, x, s)
}

/*
func TestLabel(t *testing.T) {
}
 */

func TestReturn(t *testing.T) {
  x := NewReturnNode(loc(0,0), NewVariableNode(loc(0,0), "a"))
  s := `{
  "Location": "[:0,0]",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  assertJsonEquals(t, x, s)
}
