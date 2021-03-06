package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
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
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int",
        "Type": null
      },
      "Value": 1
    },
    "Right": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int",
        "Type": null
      },
      "Value": 1
    },
    "Type": null
  }
}`
  xt.AssertStringEqualsDiff(t, "ExprStmtNode", xt.JSON(x), s)
}
