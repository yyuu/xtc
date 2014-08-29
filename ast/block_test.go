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
    []IExprNode { },
    []IStmtNode {
      ExprStmtNode(FuncallNode(VariableNode("println"), []IExprNode { StringLiteralNode("\"hello, world\"") })),
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
    []IExprNode {
      AssignNode(VariableNode("n"), IntegerLiteralNode("12345")),
    },
    []IStmtNode {
      ExprStmtNode(FuncallNode(VariableNode("printf"), []IExprNode { StringLiteralNode("\"%d\""), VariableNode("n") })),
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
    []IExprNode {
      AssignNode(VariableNode("n"), IntegerLiteralNode("12345")),
      AssignNode(VariableNode("m"), IntegerLiteralNode("67890")),
    },
    []IStmtNode {
      ExprStmtNode(FuncallNode(VariableNode("printf"), []IExprNode { StringLiteralNode("\"%d\""), VariableNode("n") })),
      ExprStmtNode(FuncallNode(VariableNode("printf"), []IExprNode { StringLiteralNode("\"%d\""), VariableNode("m") })),
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
