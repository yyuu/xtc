package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/xt"
)

func TestDoWhile(t *testing.T) {
/*
  do {
    b(a);
  } while (a < 100);
 */
  x := NewDoWhileNode(
    loc(0,0),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []duck.IExprNode { NewVariableNode(loc(0,0), "a") })),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "100")),
  )
  s := `{
  "ClassName": "ast.DoWhileNode",
  "Location": "[:0,0]",
  "Body": {
    "ClassName": "ast.ExprStmtNode",
    "Location": "[:0,0]",
    "Expr": {
      "ClassName": "ast.FuncallNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "b"
      },
      "Args": [
        {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "a"
        }
      ]
    }
  },
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "a"
    },
    "Right": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 100
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "DoWhileNode", xt.JSON(x), s)
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
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []duck.IExprNode { NewVariableNode(loc(0,0), "i") })),
  )
  s := `{
  "ClassName": "ast.ForNode",
  "Location": "[:0,0]",
  "Init": {
    "ClassName": "ast.AssignNode",
    "Location": "[:0,0]",
    "Lhs": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i"
    },
    "Rhs": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 0
    }
  },
  "Cond": {
    "ClassName": "ast.BinaryOpNode",
    "Location": "[:0,0]",
    "Operator": "\u003c",
    "Left": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i"
    },
    "Right": {
      "ClassName": "ast.IntegerLiteralNode",
      "Location": "[:0,0]",
      "Value": 100
    }
  },
  "Incr": {
    "ClassName": "ast.SuffixOpNode",
    "Location": "[:0,0]",
    "Operator": "++",
    "Expr": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "i"
    }
  },
  "Body": {
    "ClassName": "ast.ExprStmtNode",
    "Location": "[:0,0]",
    "Expr": {
      "ClassName": "ast.FuncallNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "f"
      },
      "Args": [
        {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "i"
        }
      ]
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "ForNode", xt.JSON(x), s)
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
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "gets"), []duck.IExprNode { })),
  )
  s := `{
  "ClassName": "ast.WhileNode",
  "Location": "[:0,0]",
  "Cond": {
    "ClassName": "ast.UnaryOpNode",
    "Location": "[:0,0]",
    "Operator": "!",
    "Expr": {
      "ClassName": "ast.VariableNode",
      "Location": "[:0,0]",
      "Name": "eof"
    }
  },
  "Body": {
    "ClassName": "ast.ExprStmtNode",
    "Location": "[:0,0]",
    "Expr": {
      "ClassName": "ast.FuncallNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.VariableNode",
        "Location": "[:0,0]",
        "Name": "gets"
      },
      "Args": []
    }
  }
}`
  xt.AssertStringEqualsDiff(t, "WhileNode", xt.JSON(x), s)
}
