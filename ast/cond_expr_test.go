package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestCondExpr(t *testing.T) {
/*
  (n < 2) ? 1 : (f(n-1)+f(n-2))
 */
  x := NewCondExprNode(
    loc(0,0),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")),
    NewIntegerLiteralNode(loc(0,0), "1"),
    NewBinaryOpNode(loc(0,0), "+",
                 NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []core.IExprNode { NewBinaryOpNode(loc(0,0), "-", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "1")) }),
                 NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []core.IExprNode { NewBinaryOpNode(loc(0,0), "-", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")) })))
  s := `{
  "ClassName": "ast.CondExprNode",
  "Location": "[:0,0]",
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "n",
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
      "Value": 2
    },
    "Type": null
  },
  "ThenExpr": {
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
  "ElseExpr": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "+",
    "Left": {
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
          "ClassName": "ast.BinaryOpNode",
          "Location": "[:0,0]",
          "Operator": "-",
          "Left": {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "n",
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
            "Value": 1
          },
          "Type": null
        }
      ]
    },
    "Right": {
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
          "ClassName": "ast.BinaryOpNode",
          "Location": "[:0,0]",
          "Operator": "-",
          "Left": {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "n",
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
            "Value": 2
          },
          "Type": null
        }
      ]
    },
    "Type": null
  },
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "CondExprNode", xt.JSON(x), s)
}
