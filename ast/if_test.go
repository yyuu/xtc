package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestIf(t *testing.T) {
/*
  if (n % 2 == 0) {
    println("even");
  } else {
    println("odd");
  }
 */
  x := NewIfNode(
    loc(0,0),
    NewBinaryOpNode(loc(0,0), "==", NewBinaryOpNode(loc(0,0), "%", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")), NewIntegerLiteralNode(loc(0,0), "0")),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"even\"") })),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"odd\"") })),
  )
  s := `{
  "ClassName": "ast.IfNode",
  "Location": "[:0,0]",
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "==",
    "Left": {
      "ClassName": "ast.BinaryOpNode",
      "Location": "[:0,0]",
      "Operator": "%",
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
    "Right": {
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
  "ThenBody": {
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
          "Value": "even"
        }
      ]
    }
  },
  "ElseBody": {
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
          "Value": "odd"
        }
      ]
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "if w/ else", xt.JSON(x), s)
}

func TestIfWithoutElse(t *testing.T) {
/*
  if (n % 2 == 0) {
    println("even");
  }
 */
  x := NewIfNode(
    loc(0,0),
    NewBinaryOpNode(loc(0,0), "==", NewBinaryOpNode(loc(0,0), "%", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")), NewIntegerLiteralNode(loc(0,0), "0")),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"even\"") })),
    nil,
  )
  s := `{
  "ClassName": "ast.IfNode",
  "Location": "[:0,0]",
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "==",
    "Left": {
      "ClassName": "ast.BinaryOpNode",
      "Location": "[:0,0]",
      "Operator": "%",
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
    "Right": {
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
  "ThenBody": {
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
          "Value": "even"
        }
      ]
    }
  },
  "ElseBody": null
}`
  xt.AssertStringEqualsDiff(t, "if w/o else", xt.JSON(x), s)
}
