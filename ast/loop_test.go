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
    loc(0,0),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "b"), []IExprNode { NewVariableNode(loc(0,0), "a") })),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "a"), NewIntegerLiteralNode(loc(0,0), "100")),
  )
  s := `
    (let do-while-loop ()
      (begin
        (b a)
        (if (< a 100)
            (do-while-loop))))
  `
  assertEquals(t, jsonString(x), trimSpace(s))
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
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []IExprNode { NewVariableNode(loc(0,0), "i") })),
  )
  s := `
    (let for-loop ((i 0))
      (if (< i 100)
        (begin
          (f i)
          (for-loop (+ i 1)))))
  `
  assertEquals(t, jsonString(x), trimSpace(s))
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
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "gets"), []IExprNode { })),
  )
  s := `
    (let while-loop ((while-cond (not eof)))
      (if while-cond
        (begin
          (gets)
          (while-loop (not eof)))))
  `
  assertEquals(t, jsonString(x), trimSpace(s))
}
