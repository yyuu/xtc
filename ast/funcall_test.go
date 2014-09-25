package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestFuncallNode(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "a"), []core.IExprNode { NewIntegerLiteralNode(loc(0,0), "12345"), NewIntegerLiteralNode(loc(0,0), "67890") })
  s := `{
  "ClassName": "ast.FuncallNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a",
    "Entity": null
  },
  "Args": [
    {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": {
          "ClassName": "typesys.IntegerTypeRef",
          "Location": "[:0,0]",
          "Name": "int"
        },
        "Type": null
      },
      "Value": 12345
    },
    {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": {
          "ClassName": "typesys.IntegerTypeRef",
          "Location": "[:0,0]",
          "Name": "int"
        },
        "Type": null
      },
      "Value": 67890
    }
  ]
}`
  xt.AssertStringEqualsDiff(t, "FuncallNode1", xt.JSON(x), s)
}

func TestFuncallNode2(t *testing.T) {
  x := NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []core.IExprNode { })
  s := `{
  "ClassName": "ast.FuncallNode",
  "Location": "[:0,0]",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "b",
    "Entity": null
  },
  "Args": []
}`
  xt.AssertStringEqualsDiff(t, "FuncallNode2", xt.JSON(x), s)
}
