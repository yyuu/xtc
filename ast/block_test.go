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
    []INode { },
    []INode {
      ExprStmtNode(FuncallNode(VariableNode("println"), []INode { StringLiteralNode("\"hello, world\"") })),
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
    []INode {
      AssignNode(VariableNode("n"), IntegerLiteralNode("12345")),
    },
    []INode {
      ExprStmtNode(FuncallNode(VariableNode("printf"), []INode { StringLiteralNode("\"%d\""), VariableNode("n") })),
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
    []INode {
      AssignNode(VariableNode("n"), IntegerLiteralNode("12345")),
      AssignNode(VariableNode("m"), IntegerLiteralNode("67890")),
    },
    []INode {
      ExprStmtNode(FuncallNode(VariableNode("printf"), []INode { StringLiteralNode("\"%d\""), VariableNode("n") })),
      ExprStmtNode(FuncallNode(VariableNode("printf"), []INode { StringLiteralNode("\"%d\""), VariableNode("m") })),
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
