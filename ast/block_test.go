package ast

import (
  "testing"
)

func TestBlock1(t *testing.T) {
/*
  {
    println("hello, world");
  }
 */
  x := NewBlockNode(
    loc(0,0),
    []IExprNode { },
    []IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"hello, world\"") })),
    },
  )
  s := `{
  "Location": "[:0,0]",
  "Variables": [],
  "Stmts": [
    {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Expr": {
          "Location": "[:0,0]",
          "Name": "println"
        },
        "Args": [
          {
            "Location": "[:0,0]",
            "Value": "\"hello, world\""
          }
        ]
      }
    }
  ]
}`
  assertJsonEquals(t, x, s)
}

func TestBlock2(t *testing.T) {
/*
  {
    int n = 12345;
    printf("%d", n);
  }
 */
  x := NewBlockNode(
    loc(0,0),
    []IExprNode {
      NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "12345")),
    },
    []IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "n") })),
    },
  )
  s := `{
  "Location": "[:0,0]",
  "Variables": [
    {
      "Location": "[:0,0]",
      "Lhs": {
        "Location": "[:0,0]",
        "Name": "n"
      },
      "Rhs": {
        "Location": "[:0,0]",
        "Value": 12345
      }
    }
  ],
  "Stmts": [
    {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Expr": {
          "Location": "[:0,0]",
          "Name": "printf"
        },
        "Args": [
          {
            "Location": "[:0,0]",
            "Value": "\"%d\""
          },
          {
            "Location": "[:0,0]",
            "Name": "n"
          }
        ]
      }
    }
  ]
}`
  assertJsonEquals(t, x, s)
}

func TestBlock3(t *testing.T) {
/*
  {
    int n = 12345;
    int m = 67890;
    printf("%d", n);
    printf("%d", m);
  }
 */
  x := NewBlockNode(
    loc(0,0),
    []IExprNode {
      NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "12345")),
      NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "m"), NewIntegerLiteralNode(loc(0,0), "67890")),
    },
    []IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "n") })),
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "m") })),
    },
  )
  s := `{
  "Location": "[:0,0]",
  "Variables": [
    {
      "Location": "[:0,0]",
      "Lhs": {
        "Location": "[:0,0]",
        "Name": "n"
      },
      "Rhs": {
        "Location": "[:0,0]",
        "Value": 12345
      }
    },
    {
      "Location": "[:0,0]",
      "Lhs": {
        "Location": "[:0,0]",
        "Name": "m"
      },
      "Rhs": {
        "Location": "[:0,0]",
        "Value": 67890
      }
    }
  ],
  "Stmts": [
    {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Expr": {
          "Location": "[:0,0]",
          "Name": "printf"
        },
        "Args": [
          {
            "Location": "[:0,0]",
            "Value": "\"%d\""
          },
          {
            "Location": "[:0,0]",
            "Name": "n"
          }
        ]
      }
    },
    {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Expr": {
          "Location": "[:0,0]",
          "Name": "printf"
        },
        "Args": [
          {
            "Location": "[:0,0]",
            "Value": "\"%d\""
          },
          {
            "Location": "[:0,0]",
            "Name": "m"
          }
        ]
      }
    }
  ]
}`
  assertJsonEquals(t, x, s)
}
