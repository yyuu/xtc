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

func ExprStmtNode(expr INode) exprStmtNode {
  return exprStmtNode { expr.(IExprNode) }
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

func LabelNode(name string, stmt INode) labelNode {
  return labelNode { name, stmt.(IStmtNode) }
}

func (self labelNode) String() string {
  panic("not implemented")
}

type returnNode struct {
  Expr IExprNode
}

func ReturnNode(expr INode) returnNode {
  return returnNode { expr.(IExprNode) }
}

func (self returnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}
