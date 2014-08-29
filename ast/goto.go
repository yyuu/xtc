package ast

import (
  "fmt"
)

type breakNode struct {
}

func BreakNode() breakNode {
  return breakNode { }
}

func (self breakNode) String() string {
  return "(break)"
}

type continueNode struct {
}

func ContinueNode() continueNode {
  return continueNode { }
}

func (self continueNode) String() string {
  return "(continue)"
}

type exprStmtNode struct {
  Expr IExprNode
}

func ExprStmtNode(expr IExprNode) exprStmtNode {
  return exprStmtNode { expr }
}

func (self exprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

type gotoNode struct {
  Target string
}

func GotoNode(target string) gotoNode {
  return gotoNode { target }
}

func (self gotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

type labelNode struct {
  Name string
  Stmt IStmtNode
}

func LabelNode(name string, stmt IStmtNode) labelNode {
  return labelNode { name, stmt }
}

func (self labelNode) String() string {
  panic("not implemented")
}

type returnNode struct {
  Expr IExprNode
}

func ReturnNode(expr IExprNode) returnNode {
  return returnNode { expr }
}

func (self returnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}
