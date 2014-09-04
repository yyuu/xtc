package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/duck"
)

func TestBlock1(t *testing.T) {
/*
  {
    println("hello, world");
  }
 */
  x := NewBlockNode(
    loc(0,0),
    []duck.IExprNode { },
    []duck.IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []duck.IExprNode { NewStringLiteralNode(loc(0,0), "\"hello, world\"") })),
    },
  )
  s := `{
  "ClassName": "ast.BlockNode",
  "Location": "[:0,0]",
  "Variables": [],
  "Stmts": [
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "println"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
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
    []duck.IExprNode {
      NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "12345")),
    },
    []duck.IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []duck.IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "n") })),
    },
  )
  s := `{
  "ClassName": "ast.BlockNode",
  "Location": "[:0,0]",
  "Variables": [
    {
      "ClassName": "ast.AssignNode",
      "Location": "[:0,0]",
      "Lhs": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "n"
      },
      "Rhs": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "Value": 12345
      }
    }
  ],
  "Stmts": [
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "printf"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
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
    []duck.IExprNode {
      NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "12345")),
      NewAssignNode(loc(0,0), NewVariableNode(loc(0,0), "m"), NewIntegerLiteralNode(loc(0,0), "67890")),
    },
    []duck.IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []duck.IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "n") })),
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []duck.IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "m") })),
    },
  )
  s := `{
  "ClassName": "ast.BlockNode",
  "Location": "[:0,0]",
  "Variables": [
    {
      "ClassName": "ast.AssignNode",
      "Location": "[:0,0]",
      "Lhs": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "n"
      },
      "Rhs": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "Value": 12345
      }
    },
    {
      "ClassName": "ast.AssignNode",
      "Location": "[:0,0]",
      "Lhs": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "m"
      },
      "Rhs": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "Value": 67890
      }
    }
  ],
  "Stmts": [
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "printf"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "n"
          }
        ]
      }
    },
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "printf"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
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
