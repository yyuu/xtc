package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestBreak(t *testing.T) {
  x := NewBreakNode(loc(0,0))
  s := `{
  "ClassName": "ast.BreakNode",
  "Location": "[:0,0]"
}`
  xt.AssertStringEqualsDiff(t, "BreakNode", xt.JSON(x), s)
}

func TestContinue(t *testing.T) {
  x := NewContinueNode(loc(0,0))
  s := `{
  "ClassName": "ast.ContinueNode",
  "Location": "[:0,0]"
}`
  xt.AssertStringEqualsDiff(t, "ContinueNode", xt.JSON(x), s)
}

func TestExprStmt(t *testing.T) {
  x := NewExprStmtNode(loc(0,0), NewBinaryOpNode(loc(0,0), "+", NewIntegerLiteralNode(loc(0,0), "1"), NewIntegerLiteralNode(loc(0,0), "1")))
  s := `{
  "ClassName": "ast.ExprStmtNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "+",
    "Left": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 1
    },
    "Right": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 1
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "ExprStmtNode", xt.JSON(x), s)
}

func TestGoto(t *testing.T) {
  x := NewGotoNode(loc(0,0), "a")
  s := `{
  "ClassName": "ast.GotoNode",
  "Location": "[:0,0]",
  "Target": "a"
}`
  xt.AssertStringEqualsDiff(t, "GotoNode", xt.JSON(x), s)
}

/*
func TestLabel(t *testing.T) {
}
 */

func TestReturn(t *testing.T) {
  x := NewReturnNode(loc(0,0), NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.ReturnNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  xt.AssertStringEqualsDiff(t, "VariableNode", xt.JSON(x), s)
}
