package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestOpAssignNode(t *testing.T) {
  x := NewOpAssignNode(loc(0,0), "+", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "ClassName": "ast.OpAssignNode",
  "Location": "[:0,0]",
  "Operator": "+",
  "LHS": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a",
    "Entity": null
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
      },
      "Type": null
    },
    "Value": 12345
  },
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "OpAssignNode", xt.JSON(x), s)
}
