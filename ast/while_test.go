package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestWhile(t *testing.T) {
/*
  while (!eof) {
    gets();
  }
 */
  x := NewWhileNode(
    loc(0,0),
    NewUnaryOpNode(loc(0,0), "!", NewVariableNode(loc(0,0), "eof")),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "gets"), []core.IExprNode { })),
  )
  s := `{
  "ClassName": "ast.WhileNode",
  "Location": "[:0,0]",
  "Cond": {
    "ClassName": "ast.UnaryOpNode",
    "Location": "[:0,0]",
    "Operator": "!",
    "Expr": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "eof",
      "Entity": null
    },
    "Type": null
  },
  "Body": {
    "ClassName": "ast.ExprStmtNode",
    "Location": "[:0,0]",
    "Expr": {
      "ClassName": "ast.FuncallNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "gets",
        "Entity": null
      },
      "Args": []
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "WhileNode", xt.JSON(x), s)
}
