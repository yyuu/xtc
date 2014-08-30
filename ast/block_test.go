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
  s := `
    (println "hello, world")
  `
  assertEquals(t, jsonString(x), trimSpace(s))
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
  s := `
    (let ((n 12345))
      (printf "%d" n))
  `
  assertEquals(t, jsonString(x), trimSpace(s))
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
  s := `
    (let* ((n 12345)
           (m 67890))
      (begin
        (printf "%d" n)
        (printf "%d" m)))
  `
  assertEquals(t, jsonString(x), trimSpace(s))
}
