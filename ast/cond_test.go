package ast

import (
  "testing"
)

func TestCondExpr(t *testing.T) {
/*
  (n < 2) ? 1 : (f(n-1)+f(n-2))
 */
  x := NewCondExprNode(
    loc(0,0),
    NewBinaryOpNode(loc(0,0), "<", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")),
    NewIntegerLiteralNode(loc(0,0), "1"),
    NewBinaryOpNode(loc(0,0), "+",
                 NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []IExprNode { NewBinaryOpNode(loc(0,0), "-", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "1")) }),
                 NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "f"), []IExprNode { NewBinaryOpNode(loc(0,0), "-", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")) })))
  assertEquals(t, jsonString(x), "(if (< n 2) 1 (+ (f (- n 1)) (f (- n 2))))")
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
    loc(0,0),
    NewBinaryOpNode(loc(0,0), "==", NewBinaryOpNode(loc(0,0), "%", NewVariableNode(loc(0,0), "n"), NewIntegerLiteralNode(loc(0,0), "2")), NewIntegerLiteralNode(loc(0,0), "0")),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"even\"") })),
    NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"odd\"") })),
  )
  s := `
    (if (= (modulo n 2) 0)
      (println "even")
      (println "odd"))
  `
  assertEquals(t, jsonString(x), trimSpace(s))
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
    loc(0,0),
    NewVariableNode(loc(0,0), "n"),
    []IStmtNode {
      NewCaseNode(
        loc(0,0),
        []IExprNode { NewIntegerLiteralNode(loc(0,0), "1") },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"one\"") })),
      ),
      NewCaseNode(
        loc(0,0), 
        []IExprNode { NewIntegerLiteralNode(loc(0,0), "2") },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"two\"") })),
      ),
      NewCaseNode(
        loc(0,0),
        []IExprNode { },
        NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []IExprNode { NewStringLiteralNode(loc(0,0), "\"plentiful\"") })),
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
  assertEquals(t, jsonString(x), trimSpace(s))
}
