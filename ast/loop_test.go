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
  x := DoWhileNode(
    ExprStmtNode(FuncallNode(VariableNode("b"), []IExprNode { VariableNode("a") })),
    BinaryOpNode("<", VariableNode("a"), IntegerLiteralNode("100")),
  )
  s := `
    (let do-while-loop ()
      (begin
        (b a)
        (if (< a 100)
            (do-while-loop))))
  `
  assertEquals(t, x.String(), trimSpace(s))
}

func TestFor(t *testing.T) {
/*
  for (i=0; i<100; i++) {
    f(i);
  }
 */
  x := ForNode(
    AssignNode(VariableNode("i"), IntegerLiteralNode("0")),
    BinaryOpNode("<", VariableNode("i"), IntegerLiteralNode("100")),
    SuffixOpNode("++", VariableNode("i")),
    ExprStmtNode(FuncallNode(VariableNode("f"), []IExprNode { VariableNode("i") })),
  )
  s := `
    (let for-loop ((i 0))
      (if (< i 100)
        (begin
          (f i)
          (for-loop (+ i 1)))))
  `
  assertEquals(t, x.String(), trimSpace(s))
}

func TestWhile(t *testing.T) {
/*
  while (!eof) {
    gets();
  }
 */
  x := WhileNode(
    UnaryOpNode("!", VariableNode("eof")),
    ExprStmtNode(FuncallNode(VariableNode("gets"), []IExprNode { })),
  )
  s := `
    (let while-loop ((while-cond (not eof)))
      (if while-cond
        (begin
          (gets)
          (while-loop (not eof)))))
  `
  assertEquals(t, x.String(), trimSpace(s))
}
