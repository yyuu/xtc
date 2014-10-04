package ast

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestBinaryOp(t *testing.T) {
  x := NewBinaryOpNode(loc(0,0), "*", NewBinaryOpNode(loc(0,0), "%", NewVariableNode(loc(0,0), "a"), NewVariableNode(loc(0,0), "b")), NewVariableNode(loc(0,0), "c"))
  s := `{
  "ClassName": "ast.BinaryOpNode",
  "Location": "[:0,0]",
  "Operator": "*",
  "Left": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "%",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "a",
      "Entity": null
    },
    "Right": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "b",
      "Entity": null
    },
    "Type": null
  },
  "Right": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "c",
    "Entity": null
  },
  "Type": null
}`
  xt.AssertStringEqualsDiff(t, "BinaryOpNode", xt.JSON(x), s)
}
