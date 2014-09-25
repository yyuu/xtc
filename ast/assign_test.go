package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestAssignNode(t *testing.T) {
  x := NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewStringLiteralNode(loc(0,0), "\"xxx\""))
  s := `{
  "ClassName": "ast.AssignNode",
  "Location": "[:0,0]",
  "LHS": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "RHS": {
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
      },
      "Type": null
    },
    "Value": "\"xxx\""
  },
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "AssignNode", xt.JSON(x), s)
}
