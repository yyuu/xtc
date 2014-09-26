package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestSwitch(t *testing.T) {
  /*
  switch (n) {
    case 1: println("one");
    case 2: println("two");
    default: println("plentiful")
  }
   */
  x := NewSwitchNode(
    loc(0,0),
    NewVariableNode(loc(0,0), "n"),
    []core.IStmtNode {
      NewCaseNode(
        loc(0,0),
        []core.IExprNode { NewIntegerLiteralNode(loc(0,0), "1") },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"one\"") })),
      ),
      NewCaseNode(
        loc(0,0), 
        []core.IExprNode { NewIntegerLiteralNode(loc(0,0), "2") },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"two\"") })),
      ),
      NewCaseNode(
        loc(0,0),
        []core.IExprNode { },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"plentiful\"") })),
      ),
    },
  )
  s := `{
  "ClassName": "ast.SwitchNode",
  "Location": "[:0,0]",
  "Cond": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "n",
    "Entity": null
  },
  "Cases": [
    {
      "ClassName": "ast.CaseNode",
      "Location": "[:0,0]",
      "Label": null,
      "Values": [
        {
          "ClassName": "ast.IntegerLiteralNode",
          "Location": "[:0,0]",
          "TypeNode": {
            "ClassName": "ast.TypeNode",
            "Location": "[:0,0]",
            "TypeRef": "int",
            "Type": null
          },
          "Value": 1
        }
      ],
      "Body": {
        "ClassName": "ast.ExprStmtNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.FuncallNode",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "println",
            "Entity": null
          },
          "Args": [
            {
              "ClassName": "ast.StringLiteralNode",
              "Location": "[:0,0]",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "char*",
                "Type": null
              },
              "Value": "\"one\""
            }
          ]
        }
      }
    },
    {
      "ClassName": "ast.CaseNode",
      "Location": "[:0,0]",
      "Label": null,
      "Values": [
        {
          "ClassName": "ast.IntegerLiteralNode",
          "Location": "[:0,0]",
          "TypeNode": {
            "ClassName": "ast.TypeNode",
            "Location": "[:0,0]",
            "TypeRef": "int",
            "Type": null
          },
          "Value": 2
        }
      ],
      "Body": {
        "ClassName": "ast.ExprStmtNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.FuncallNode",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "println",
            "Entity": null
          },
          "Args": [
            {
              "ClassName": "ast.StringLiteralNode",
              "Location": "[:0,0]",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "char*",
                "Type": null
              },
              "Value": "\"two\""
            }
          ]
        }
      }
    },
    {
      "ClassName": "ast.CaseNode",
      "Location": "[:0,0]",
      "Label": null,
      "Values": [],
      "Body": {
        "ClassName": "ast.ExprStmtNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.FuncallNode",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "println",
            "Entity": null
          },
          "Args": [
            {
              "ClassName": "ast.StringLiteralNode",
              "Location": "[:0,0]",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "char*",
                "Type": null
              },
              "Value": "\"plentiful\""
            }
          ]
        }
      }
    }
  ]
}`
  xt.AssertStringEqualsDiff(t, "SwitchNode", xt.JSON(x), s)
}
