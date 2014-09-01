package ast

import (
  "testing"
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
  assertJsonEquals(t, x, s)
}

func TestLogicalAndNode(t *testing.T) {
  x := NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "b"), NewVariableNode(loc(0,0), "c")))
  s := `{
  "ClassName": "ast.LogicalAndNode",
  "Location": "[:0,0]",
  "Left": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Right": {
    "ClassName": "ast.LogicalAndNode",
    "Location": "[:0,0]",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "b"
    },
    "Right": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "c"
    }
  }
}`
  assertJsonEquals(t, x, s)
}

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
  assertJsonEquals(t, x, s)
}

func TestPrefixOpNode(t *testing.T) {
  x := NewPrefixOpNode(loc(0,0), "--", NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.PrefixOpNode",
  "Location": "[:0,0]",
  "Operator": "--",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  assertJsonEquals(t, x, s)
}

func TestSuffixOpNode(t *testing.T) {
  x := NewSuffixOpNode(loc(0,0), "++", NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.SuffixOpNode",
  "Location": "[:0,0]",
  "Operator": "++",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  assertJsonEquals(t, x, s)
}

func TestUnaryOpNode1(t *testing.T) {
  x := NewUnaryOpNode(loc(0,0), "-", NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "ClassName": "ast.UnaryOpNode",
  "Location": "[:0,0]",
  "Operator": "-",
  "Expr": {
    "ClassName": "ast.IntegerLiteralNode",
    "Location": "[:0,0]",
    "Value": 12345
  }
}`
  assertJsonEquals(t, x, s)
}

func TestUnaryOpNode2(t *testing.T) {
  x := NewUnaryOpNode(loc(0,0), "!", NewVariableNode(loc(0,0), "a"))
  s := `{
  "ClassName": "ast.UnaryOpNode",
  "Location": "[:0,0]",
  "Operator": "!",
  "Expr": {
    "ClassName": "ast.VariableNode",
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  assertJsonEquals(t, x, s)
}
