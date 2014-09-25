package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestFor(t *testing.T) {
/*
  for (i=0; i<100; i++) {
    f(i);
  }
 */
  x := NewForNode(
    loc(0,0),
    NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "i"), NewIntegerLiteralNode(loc(0,0), "0")),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "i"), NewIntegerLiteralNode(loc(0,0), "100")),
    NewSuffixOpNode(loc(0,0), "++", NewVariableNode(loc(0,0), "i")),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []core.IExprNode { NewVariableNode(loc(0,0), "i") })),
  )
  s := `{
  "ClassName": "ast.ForNode",
  "Location": "[:0,0]",
  "Init": {
    "ClassName": "ast.AssignNode",
    "Location": "[:0,0]",
    "LHS": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i",
      "Entity": null
    },
    "RHS": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int",
        "Type": null
      },
      "Value": 0
    },
    "Type": null
  },
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i",
      "Entity": null
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
      "Value": 100
    },
    "Type": null
  },
  "Incr": {
    "ClassName": "ast.SuffixOpNode",
    "Location": "[:0,0]",
    "Operator": "++",
    "Expr": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i",
      "Entity": null
    },
    "Amount": 1,
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
        "Name": "f",
        "Entity": null
      },
      "Args": [
        {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "i",
          "Entity": null
        }
      ]
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "ForNode", xt.JSON(x), s)
}
