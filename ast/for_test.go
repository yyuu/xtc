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
      "Name": "i"
    },
    "RHS": {
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
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i"
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
      "Value": 100
    }
  },
  "Incr": {
    "ClassName": "ast.SuffixOpNode",
    "Location": "[:0,0]",
    "Operator": "++",
    "Expr": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i"
    }
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
        "Name": "f"
      },
      "Args": [
        {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "i"
        }
      ]
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "ForNode", xt.JSON(x), s)
}
