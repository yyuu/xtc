package ast

// CastNode
type castNode struct {
  Type ITypeNode
  Expr IExprNode
}

func CastNode(t ITypeNode, expr IExprNode) castNode {
  return castNode { t, expr }
}

func (self castNode) String() string {
  panic("not implemented")
}

func (self castNode) IsExpr() bool {
  return true
}

// SizeofExprNode
type sizeofExprNode struct {
  Expr IExprNode
  Type ITypeNode
}

func SizeofExprNode(expr IExprNode, t ITypeNode) sizeofExprNode {
  return sizeofExprNode { expr, t }
}

func (self sizeofExprNode) String() string {
  panic("not implemented")
}

func (self sizeofExprNode) IsExpr() bool {
  return true
}

// SizeofTypeNode
type sizeofTypeNode struct {
  Type ITypeNode
  Operand ITypeNode
}

func SizeofTypeNode(t ITypeNode, operand ITypeNode) sizeofTypeNode {
  return sizeofTypeNode { t, operand }
}

func (self sizeofTypeNode) String() string {
  panic("not implemented")
}

func (self sizeofTypeNode) IsExpr() bool {
  return true
}
