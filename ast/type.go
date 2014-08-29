package ast

// CastNode
type castNode struct {
  location Location
  Type ITypeNode
  Expr IExprNode
}

func CastNode(location Location, t ITypeNode, expr IExprNode) castNode {
  return castNode { location, t, expr }
}

func (self castNode) String() string {
  panic("not implemented")
}

func (self castNode) IsExpr() bool {
  return true
}

func (self castNode) GetLocation() Location {
  return self.location
}

// SizeofExprNode
type sizeofExprNode struct {
  location Location
  Expr IExprNode
  Type ITypeNode
}

func SizeofExprNode(location Location, expr IExprNode, t ITypeNode) sizeofExprNode {
  return sizeofExprNode { location, expr, t }
}

func (self sizeofExprNode) String() string {
  panic("not implemented")
}

func (self sizeofExprNode) IsExpr() bool {
  return true
}

func (self sizeofExprNode) GetLocation() Location {
  return self.location
}

// SizeofTypeNode
type sizeofTypeNode struct {
  location Location
  Type ITypeNode
  Operand ITypeNode
}

func SizeofTypeNode(location Location, t ITypeNode, operand ITypeNode) sizeofTypeNode {
  return sizeofTypeNode { location, t, operand }
}

func (self sizeofTypeNode) String() string {
  panic("not implemented")
}

func (self sizeofTypeNode) IsExpr() bool {
  return true
}

func (self sizeofTypeNode) GetLocation() Location {
  return self.location
}
