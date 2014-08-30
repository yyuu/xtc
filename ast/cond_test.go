package ast

import (
  "testing"
)

func TestCondExpr(t *testing.T) {
/*
  (n < 2) ? 1 : (f(n-1)+f(n-2))
 */
  x := NewCondExprNode(
    loc(0,0),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")),
    NewIntegerLiteralNode(loc(0,0), "1"),
    NewBinaryOpNode(loc(0,0), "+",
                 NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []IExprNode { NewBinaryOpNode(loc(0,0), "-", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "1")) }),
                 NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []IExprNode { NewBinaryOpNode(loc(0,0), "-", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")) })))
  s := `{
  "Location": "[:0,0]",
  "Cond": {
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "Location": "[:0,0]",
      "Name": "n"
    },
    "Right": {
      "Location": "[:0,0]",
      "Value": 2
    }
  },
  "ThenExpr": {
    "Location": "[:0,0]",
    "Value": 1
  },
  "ElseExpr": {
    "Location": "[:0,0]",
    "Operator": "+",
    "Left": {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Name": "f"
      },
      "Args": [
        {
          "Location": "[:0,0]",
          "Operator": "-",
          "Left": {
            "Location": "[:0,0]",
            "Name": "n"
          },
          "Right": {
            "Location": "[:0,0]",
            "Value": 1
          }
        }
      ]
    },
    "Right": {
      "Location": "[:0,0]",
      "Expr": {
        "Location": "[:0,0]",
        "Name": "f"
      },
      "Args": [
        {
          "Location": "[:0,0]",
          "Operator": "-",
          "Left": {
            "Location": "[:0,0]",
            "Name": "n"
          },
          "Right": {
            "Location": "[:0,0]",
            "Value": 2
          }
        }
      ]
    }
  }
}`
  assertJsonEquals(t, x, s)
}

func TestIf(t *testing.T) {
/*
  if (n % 2 == 0) {
    println("even");
  } else {
    println("odd");
  }
 */
  x := NewIfNode(
    loc(0,0),
    NewBinaryOpNode(loc(0,0), "==", NewBinaryOpNode(loc(0,0), "%", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")), NewIntegerLiteralNode(loc(0,0), "0")),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"even\"") })),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"odd\"") })),
  )
  s := `{
  "Location": "[:0,0]",
  "Cond": {
    "Location": "[:0,0]",
    "Operator": "==",
    "Left": {
      "Location": "[:0,0]",
      "Operator": "%",
      "Left": {
        "Location": "[:0,0]",
        "Name": "n"
      },
      "Right": {
        "Location": "[:0,0]",
        "Value": 2
      }
    },
    "Right": {
      "Location": "[:0,0]",
      "Value": 0
    }
  },
  "ThenBody": {
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
          "Value": "\"even\""
        }
      ]
    }
  },
  "ElseBody": {
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
          "Value": "\"odd\""
        }
      ]
    }
  }
}`
  assertJsonEquals(t, x, s)
}

func TestSwitch(t *testing.T) {
  /*
  switch (n) {
    case 1: println("one");
    case 2: println("two");
    default: println("plentiful")
  }
   */
  x := NewSwitchNode(
    loc(0,0),
    NewVariableNode(loc(0,0), "n"),
    []IStmtNode {
      NewCaseNode(
        loc(0,0),
        []IExprNode { NewIntegerLiteralNode(loc(0,0), "1") },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"one\"") })),
      ),
      NewCaseNode(
        loc(0,0), 
        []IExprNode { NewIntegerLiteralNode(loc(0,0), "2") },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"two\"") })),
      ),
      NewCaseNode(
        loc(0,0),
        []IExprNode { },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"plentiful\"") })),
      ),
    },
  )
  s := `{
  "Location": "[:0,0]",
  "Cond": {
    "Location": "[:0,0]",
    "Name": "n"
  },
  "Cases": [
    {
      "Location": "[:0,0]",
      "Values": [
        {
          "Location": "[:0,0]",
          "Value": 1
        }
      ],
      "Body": {
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
              "Value": "\"one\""
            }
          ]
        }
      }
    },
    {
      "Location": "[:0,0]",
      "Values": [
        {
          "Location": "[:0,0]",
          "Value": 2
        }
      ],
      "Body": {
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
              "Value": "\"two\""
            }
          ]
        }
      }
    },
    {
      "Location": "[:0,0]",
      "Values": [],
      "Body": {
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
              "Value": "\"plentiful\""
            }
          ]
        }
      }
    }
  ]
}`
  assertJsonEquals(t, x, s)
}
