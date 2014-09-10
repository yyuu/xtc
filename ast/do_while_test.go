package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestDoWhile(t *testing.T) {
/*
  do {
    b(a);
  } while (a < 100);
 */
  x := NewDoWhileNode(
    loc(0,0),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []core.IExprNode { NewVariableNode(loc(0,0), "a") })),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "100")),
  )
  s := `{
  "ClassName": "ast.DoWhileNode",
  "Location": "[:0,0]",
  "Body": {
    "ClassName": "ast.ExprStmtNode",
    "Location": "[:0,0]",
    "Expr": {
      "ClassName": "ast.FuncallNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "b"
      },
      "Args": [
        {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "a"
        }
      ]
    }
  },
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "a"
    },
    "Right": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 100
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "DoWhileNode", xt.JSON(x), s)
}
