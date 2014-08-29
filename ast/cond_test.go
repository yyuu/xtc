package ast

import (
  "testing"
)

func TestCondExpr(t *testing.T) {
/*
  (n < 2) ? 1 : (f(n-1)+f(n-2))
 */
  x := CondExprNode(
    LOC,
    BinaryOpNode(LOC, "<", VariableNode(LOC, "n"), IntegerLiteralNode(LOC, "2")),
    IntegerLiteralNode(LOC, "1"),
    BinaryOpNode(LOC, "+",
                 FuncallNode(LOC, VariableNode(LOC, "f"), []IExprNode { BinaryOpNode(LOC, "-", VariableNode(LOC, "n"), IntegerLiteralNode(LOC, "1")) }),
                 FuncallNode(LOC, VariableNode(LOC, "f"), []IExprNode { BinaryOpNode(LOC, "-", VariableNode(LOC, "n"), IntegerLiteralNode(LOC, "2")) })))
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
  x := IfNode(
    LOC,
    BinaryOpNode(LOC, "==", BinaryOpNode(LOC, "%", VariableNode(LOC, "n"), IntegerLiteralNode(LOC, "2")), IntegerLiteralNode(LOC, "0")),
    ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "println"), []IExprNode { StringLiteralNode(LOC, "\"even\"") })),
    ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "println"), []IExprNode { StringLiteralNode(LOC, "\"odd\"") })),
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
  x := SwitchNode(
    LOC,
    VariableNode(LOC, "n"),
    []IStmtNode {
      CaseNode(
        LOC,
        []IExprNode { IntegerLiteralNode(LOC, "1") },
        ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "println"), []IExprNode { StringLiteralNode(LOC, "\"one\"") })),
      ),
      CaseNode(
        LOC, 
        []IExprNode { IntegerLiteralNode(LOC, "2") },
        ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "println"), []IExprNode { StringLiteralNode(LOC, "\"two\"") })),
      ),
      CaseNode(
        LOC,
        []IExprNode { },
        ExprStmtNode(LOC, FuncallNode(LOC, VariableNode(LOC, "println"), []IExprNode { StringLiteralNode(LOC, "\"plentiful\"") })),
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
