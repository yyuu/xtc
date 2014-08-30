package ast

import (
  "testing"
)

func TestDoWhile(t *testing.T) {
/*
  do {
    b(a);
  } while (a < 100);
 */
  x := NewDoWhileNode(
    loc(0,0),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []IExprNode { NewVariableNode(loc(0,0), "a") })),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "100")),
  )
  s := `{
  "Location": "[:0,0]",
  "Body": {
    "Location": "[:0,0]",
    "Expr": {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Name": "b"
      },
      "Args": [
        {
          "Location": "[:0,0]",
          "Name": "a"
        }
      ]
    }
  },
  "Cond": {
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "Location": "[:0,0]",
      "Name": "a"
    },
    "Right": {
      "Location": "[:0,0]",
      "Value": 100
    }
  }
}`
  assertJsonEquals(t, x, s)
}

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
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []IExprNode { NewVariableNode(loc(0,0), "i") })),
  )
  s := `{
  "Location": "[:0,0]",
  "Init": {
    "Location": "[:0,0]",
    "Lhs": {
      "Location": "[:0,0]",
      "Name": "i"
    },
    "Rhs": {
      "Location": "[:0,0]",
      "Value": 0
    }
  },
  "Cond": {
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "Location": "[:0,0]",
      "Name": "i"
    },
    "Right": {
      "Location": "[:0,0]",
      "Value": 100
    }
  },
  "Incr": {
    "Location": "[:0,0]",
    "Operator": "++",
    "Expr": {
      "Location": "[:0,0]",
      "Name": "i"
    }
  },
  "Body": {
    "Location": "[:0,0]",
    "Expr": {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Name": "f"
      },
      "Args": [
        {
          "Location": "[:0,0]",
          "Name": "i"
        }
      ]
    }
  }
}`
  assertJsonEquals(t, x, s)
}

func TestWhile(t *testing.T) {
/*
  while (!eof) {
    gets();
  }
 */
  x := NewWhileNode(
    loc(0,0),
    NewUnaryOpNode(loc(0,0), "!", NewVariableNode(loc(0,0), "eof")),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "gets"), []IExprNode { })),
  )
  s := `{
  "Location": "[:0,0]",
  "Cond": {
    "Location": "[:0,0]",
    "Operator": "!",
    "Expr": {
      "Location": "[:0,0]",
      "Name": "eof"
    }
  },
  "Body": {
    "Location": "[:0,0]",
    "Expr": {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Name": "gets"
      },
      "Args": []
    }
  }
}`
  assertJsonEquals(t, x, s)
}
