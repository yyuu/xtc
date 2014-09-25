package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestLogicalOrNode(t *testing.T) {
  x := NewLogicalOrNode(loc(0,0), NewLogicalOrNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewVariableNode(loc(0,0), "b")), NewVariableNode(loc(0,0), "c"))
  s := `{
  "ClassName": "ast.LogicalOrNode",
  "Location": "[:0,0]",
  "Left": {
    "ClassName": "ast.LogicalOrNode",
    "Location": "[:0,0]",
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
  xt.AssertStringEqualsDiff(t, "LogicalOrNode", xt.JSON(x), s)
}
