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
    LOC,
    []IExprNode { },
    []IStmtNode {
      NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "println"), []IExprNode { NewStringLiteralNode(LOC, "\"hello, world\"") })),
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
  x := NewBlockNode(
    LOC,
    []IExprNode {
      NewAssignNode(LOC, NewVariableNode(LOC, "n"), NewIntegerLiteralNode(LOC, "12345")),
    },
    []IStmtNode {
      NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "printf"), []IExprNode { NewStringLiteralNode(LOC, "\"%d\""), NewVariableNode(LOC, "n") })),
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
  x := NewBlockNode(
    LOC,
    []IExprNode {
      NewAssignNode(LOC, NewVariableNode(LOC, "n"), NewIntegerLiteralNode(LOC, "12345")),
      NewAssignNode(LOC, NewVariableNode(LOC, "m"), NewIntegerLiteralNode(LOC, "67890")),
    },
    []IStmtNode {
      NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "printf"), []IExprNode { NewStringLiteralNode(LOC, "\"%d\""), NewVariableNode(LOC, "n") })),
      NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "printf"), []IExprNode { NewStringLiteralNode(LOC, "\"%d\""), NewVariableNode(LOC, "m") })),
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
