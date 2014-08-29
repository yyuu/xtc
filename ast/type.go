package ast

import (
  "fmt"
)

type castNode struct {
  Type ITypeNode
  Expr IExprNode
}

func CastNode(t INode, expr INode) castNode {
  return castNode { t.(ITypeNode), expr.(IExprNode) }
}

func (self castNode) String() string {
  panic("not implemented")
}

type condExprNode struct {
  Cond IExprNode
  ThenExpr IExprNode
  ElseExpr IExprNode
}

func CondExprNode(cond IExprNode, thenExpr IExprNode, elseExpr IExprNode) condExprNode {
  return condExprNode { cond, thenExpr, elseExpr }
}

func (self condExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

type sizeofExprNode struct {
  Expr IExprNode
  Type ITypeNode
}

func SizeofExprNode(expr INode, t INode) sizeofExprNode {
  return sizeofExprNode { expr.(IExprNode), t.(ITypeNode) }
}

func (self sizeofExprNode) String() string {
  panic("not implemented")
}

type sizeofTypeNode struct {
  Type ITypeNode
  Operand ITypeNode
}

func SizeofTypeNode(t INode, operand INode) sizeofTypeNode {
  return sizeofTypeNode { t.(ITypeNode), operand.(ITypeNode) }
}

func (self sizeofTypeNode) String() string {
  panic("not implemented")
}
