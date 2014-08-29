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
  x := BlockNode(
    LOC,
    []IExprNode { },
    []IStmtNode {
      ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "println"), []IExprNode { StringLiteralNode(LOC, "\"hello, world\"") })),
    },
  )
  s := `
    (println "hello, world")
  `
  assertEquals(t, x.String(), trimSpace(s))
}

func TestBlock2(t *testing.T) {
/*
  {
    int n = 12345;
    printf("%d", n);
  }
 */
  x := BlockNode(
    LOC,
    []IExprNode {
      AssignNode(LOC, VariableNode(LOC, "n"), IntegerLiteralNode(LOC, "12345")),
    },
    []IStmtNode {
      ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "printf"), []IExprNode { StringLiteralNode(LOC, "\"%d\""), VariableNode(LOC, "n") })),
    },
  )
  s := `
    (let ((n 12345))
      (printf "%d" n))
  `
  assertEquals(t, x.String(), trimSpace(s))
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
  x := BlockNode(
    LOC,
    []IExprNode {
      AssignNode(LOC, VariableNode(LOC, "n"), IntegerLiteralNode(LOC, "12345")),
      AssignNode(LOC, VariableNode(LOC, "m"), IntegerLiteralNode(LOC, "67890")),
    },
    []IStmtNode {
      ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "printf"), []IExprNode { StringLiteralNode(LOC, "\"%d\""), VariableNode(LOC, "n") })),
      ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "printf"), []IExprNode { StringLiteralNode(LOC, "\"%d\""), VariableNode(LOC, "m") })),
    },
  )
  s := `
    (let* ((n 12345)
           (m 67890))
      (begin
        (printf "%d" n)
        (printf "%d" m)))
  `
  assertEquals(t, x.String(), trimSpace(s))
}
