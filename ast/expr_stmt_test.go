package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

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
