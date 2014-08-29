package ast

// CastNode
type castNode struct {
  location ILocation
  Type ITypeNode
  Expr IExprNode
}

func CastNode(location ILocation, t ITypeNode, expr IExprNode) castNode {
  return castNode { location, t, expr }
}

func (self castNode) String() string {
  panic("not implemented")
}

func (self castNode) IsExpr() bool {
  return true
}

func (self castNode) Location() ILocation {
  return self.location
}

// SizeofExprNode
type sizeofExprNode struct {
  location ILocation
  Expr IExprNode
  Type ITypeNode
}

func SizeofExprNode(location ILocation, expr IExprNode, t ITypeNode) sizeofExprNode {
  return sizeofExprNode { location, expr, t }
}

func (self sizeofExprNode) String() string {
  panic("not implemented")
}

func (self sizeofExprNode) IsExpr() bool {
  return true
}

func (self sizeofExprNode) Location() ILocation {
  return self.location
}

// SizeofTypeNode
type sizeofTypeNode struct {
  location ILocation
  Type ITypeNode
  Operand ITypeNode
}

func SizeofTypeNode(location ILocation, t ITypeNode, operand ITypeNode) sizeofTypeNode {
  return sizeofTypeNode { location, t, operand }
}

func (self sizeofTypeNode) String() string {
  panic("not implemented")
}

func (self sizeofTypeNode) IsExpr() bool {
  return true
}

func (self sizeofTypeNode) Location() ILocation {
  return self.location
}
