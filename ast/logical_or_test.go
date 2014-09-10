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
      "Name": "a"
    },
    "Right": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "b"
    }
  },
  "Right": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "c"
  }
}`
  xt.AssertStringEqualsDiff(t, "LogicalOrNode", xt.JSON(x), s)
}
