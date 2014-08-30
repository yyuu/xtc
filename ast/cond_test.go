package ast

import (
  "testing"
)

func TestCondExpr(t *testing.T) {
/*
  (n < 2) ? 1 : (f(n-1)+f(n-2))
 */
  x := NewCondExprNode(
    LOC,
    NewBinaryOpNode(LOC, "<", NewVariableNode(LOC, "n"), NewIntegerLiteralNode(LOC, "2")),
    NewIntegerLiteralNode(LOC, "1"),
    NewBinaryOpNode(LOC, "+",
                 NewFuncallNode(LOC, NewVariableNode(LOC, "f"), []IExprNode { NewBinaryOpNode(LOC, "-", NewVariableNode(LOC, "n"), NewIntegerLiteralNode(LOC, "1")) }),
                 NewFuncallNode(LOC, NewVariableNode(LOC, "f"), []IExprNode { NewBinaryOpNode(LOC, "-", NewVariableNode(LOC, "n"), NewIntegerLiteralNode(LOC, "2")) })))
  assertEquals(t, x.String(), "(if (< n 2) 1 (+ (f (- n 1)) (f (- n 2))))")
}

func TestIf(t *testing.T) {
/*
  if (n % 2 == 0) {
    println("even");
  } else {
    println("odd");
  }
 */
  x := NewIfNode(
    LOC,
    NewBinaryOpNode(LOC, "==", NewBinaryOpNode(LOC, "%", NewVariableNode(LOC, "n"), NewIntegerLiteralNode(LOC, "2")), NewIntegerLiteralNode(LOC, "0")),
    NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "println"), []IExprNode { NewStringLiteralNode(LOC, "\"even\"") })),
    NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "println"), []IExprNode { NewStringLiteralNode(LOC, "\"odd\"") })),
  )
  s := `
    (if (= (modulo n 2) 0)
      (println "even")
      (println "odd"))
  `
  assertEquals(t, x.String(), trimSpace(s))
}

func TestSwitch(t *testing.T) {
  /*
  switch (n) {
    case 1: println("one");
    case 2: println("two");
    default: println("plentiful")
  }
   */
  x := NewSwitchNode(
    LOC,
    NewVariableNode(LOC, "n"),
    []IStmtNode {
      NewCaseNode(
        LOC,
        []IExprNode { NewIntegerLiteralNode(LOC, "1") },
        NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "println"), []IExprNode { NewStringLiteralNode(LOC, "\"one\"") })),
      ),
      NewCaseNode(
        LOC, 
        []IExprNode { NewIntegerLiteralNode(LOC, "2") },
        NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "println"), []IExprNode { NewStringLiteralNode(LOC, "\"two\"") })),
      ),
      NewCaseNode(
        LOC,
        []IExprNode { },
        NewExprStmtNode(LOC, NewFuncallNode(LOC, NewVariableNode(LOC, "println"), []IExprNode { NewStringLiteralNode(LOC, "\"plentiful\"") })),
      ),
    },
  )
  s := `
    (let ((switch-cond n))
      (cond
        ((= switch-cond 1) (println "one"))
        ((= switch-cond 2) (println "two"))
        (else (println "plentiful"))))
  `
  assertEquals(t, x.String(), trimSpace(s))
}
