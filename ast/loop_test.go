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
    LOC,
    ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "b"), []IExprNode { VariableNode(LOC, "a") })),
    BinaryOpNode(LOC, "<", VariableNode(LOC, "a"), IntegerLiteralNode(LOC, "100")),
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
    LOC,
    AssignNode(LOC, VariableNode(LOC, "i"), IntegerLiteralNode(LOC, "0")),
    BinaryOpNode(LOC, "<", VariableNode(LOC, "i"), IntegerLiteralNode(LOC, "100")),
    SuffixOpNode(LOC, "++", VariableNode(LOC, "i")),
    ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "f"), []IExprNode { VariableNode(LOC, "i") })),
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
    LOC,
    UnaryOpNode(LOC, "!", VariableNode(LOC, "eof")),
    ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "gets"), []IExprNode { })),
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
