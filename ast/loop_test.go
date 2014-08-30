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
  x := NewDoWhileNode(
    LOC,
    NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "b"), []IExprNode { NewVariableNode(LOC, "a") })),
    NewBinaryOpNode(LOC, "<", NewVariableNode(LOC, "a"), NewIntegerLiteralNode(LOC, "100")),
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
  x := NewForNode(
    LOC,
    NewAssignNode(LOC, NewVariableNode(LOC, "i"), NewIntegerLiteralNode(LOC, "0")),
    NewBinaryOpNode(LOC, "<", NewVariableNode(LOC, "i"), NewIntegerLiteralNode(LOC, "100")),
    NewSuffixOpNode(LOC, "++", NewVariableNode(LOC, "i")),
    NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "f"), []IExprNode { NewVariableNode(LOC, "i") })),
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
  x := NewWhileNode(
    LOC,
    NewUnaryOpNode(LOC, "!", NewVariableNode(LOC, "eof")),
    NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "gets"), []IExprNode { })),
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
