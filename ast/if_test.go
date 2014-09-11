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
        "Name": "n"
      },
      "Right": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "TypeNode": {
          "ClassName": "ast.TypeNode",
          "Location": "[:0,0]",
          "TypeRef": {
            "ClassName": "typesys.IntegerTypeRef",
            "Location": "[:0,0]",
            "Name": "int"
          }
        },
        "Value": 2
      }
    },
    "Right": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": {
          "ClassName": "typesys.IntegerTypeRef",
          "Location": "[:0,0]",
          "Name": "int"
        }
      },
      "Value": 0
    }
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
        "Name": "println"
      },
      "Args": [
        {
          "ClassName": "ast.StringLiteralNode",
          "Location": "[:0,0]",
          "TypeNode": {
            "ClassName": "ast.TypeNode",
            "Location": "[:0,0]",
            "TypeRef": {
              "ClassName": "typesys.PointerTypeRef",
              "Location": "[:0,0]",
              "BaseType": {
                "ClassName": "typesys.IntegerTypeRef",
                "Location": "[:0,0]",
                "Name": "char"
              }
            }
          },
          "Value": "\"even\""
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
        "Name": "println"
      },
      "Args": [
        {
          "ClassName": "ast.StringLiteralNode",
          "Location": "[:0,0]",
          "TypeNode": {
            "ClassName": "ast.TypeNode",
            "Location": "[:0,0]",
            "TypeRef": {
              "ClassName": "typesys.PointerTypeRef",
              "Location": "[:0,0]",
              "BaseType": {
                "ClassName": "typesys.IntegerTypeRef",
                "Location": "[:0,0]",
                "Name": "char"
              }
            }
          },
          "Value": "\"odd\""
        }
      ]
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "IfNode", xt.JSON(x), s)
}
