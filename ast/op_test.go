package ast

import (
  "testing"
)

func TestBinaryOp(t *testing.T) {
  x := NewBinaryOpNode(loc(0,0), "*", NewBinaryOpNode(loc(0,0), "%", NewVariableNode(loc(0,0), "a"), NewVariableNode(loc(0,0), "b")), NewVariableNode(loc(0,0), "c"))
  s := `{
  "Location": "[:0,0]",
  "Operator": "*",
  "Left": {
    "Location": "[:0,0]",
    "Operator": "%",
    "Left": {
      "Location": "[:0,0]",
      "Name": "a"
    },
    "Right": {
      "Location": "[:0,0]",
      "Name": "b"
    }
  },
  "Right": {
    "Location": "[:0,0]",
    "Name": "c"
  }
}`
  assertJsonEquals(t, x, s)
}

func TestLogicalAndNode(t *testing.T) {
  x := NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "a"), NewLogicalAndNode(loc(0,0), NewVariableNode(loc(0,0), "b"), NewVariableNode(loc(0,0), "c")))
  s := `{
  "Location": "[:0,0]",
  "Left": {
    "Location": "[:0,0]",
    "Name": "a"
  },
  "Right": {
    "Location": "[:0,0]",
    "Left": {
      "Location": "[:0,0]",
      "Name": "b"
    },
    "Right": {
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
  "Location": "[:0,0]",
  "Left": {
    "Location": "[:0,0]",
    "Left": {
      "Location": "[:0,0]",
      "Name": "a"
    },
    "Right": {
      "Location": "[:0,0]",
      "Name": "b"
    }
  },
  "Right": {
    "Location": "[:0,0]",
    "Name": "c"
  }
}`
  assertJsonEquals(t, x, s)
}

func TestPrefixOpNode(t *testing.T) {
  x := NewPrefixOpNode(loc(0,0), "--", NewVariableNode(loc(0,0), "a"))
  s := `{
  "Location": "[:0,0]",
  "Operator": "--",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  assertJsonEquals(t, x, s)
}

func TestSuffixOpNode(t *testing.T) {
  x := NewSuffixOpNode(loc(0,0), "++", NewVariableNode(loc(0,0), "a"))
  s := `{
  "Location": "[:0,0]",
  "Operator": "++",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  assertJsonEquals(t, x, s)
}

func TestUnaryOpNode1(t *testing.T) {
  x := NewUnaryOpNode(loc(0,0), "-", NewIntegerLiteralNode(loc(0,0), "12345"))
  s := `{
  "Location": "[:0,0]",
  "Operator": "-",
  "Expr": {
    "Location": "[:0,0]",
    "Value": 12345
  }
}`
  assertJsonEquals(t, x, s)
}

func TestUnaryOpNode2(t *testing.T) {
  x := NewUnaryOpNode(loc(0,0), "!", NewVariableNode(loc(0,0), "a"))
  s := `{
  "Location": "[:0,0]",
  "Operator": "!",
  "Expr": {
    "Location": "[:0,0]",
    "Name": "a"
  }
}`
  assertJsonEquals(t, x, s)
}
