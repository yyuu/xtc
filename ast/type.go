package ast

// CastNode
type CastNode struct {
  Location Location
  Type ITypeNode
  Expr IExprNode
}

func NewCastNode(location Location, t ITypeNode, expr IExprNode) CastNode {
  return CastNode { location, t, expr }
}

func (self CastNode) String() string {
  panic("not implemented")
}

func (self CastNode) IsExpr() bool {
  return true
}

func (self CastNode) GetLocation() Location {
  return self.Location
}

// SizeofExprNode
type SizeofExprNode struct {
  Location Location
  Expr IExprNode
  Type ITypeNode
}

func NewSizeofExprNode(location Location, expr IExprNode, t ITypeNode) SizeofExprNode {
  return SizeofExprNode { location, expr, t }
}

func (self SizeofExprNode) String() string {
  panic("not implemented")
}

func (self SizeofExprNode) IsExpr() bool {
  return true
}

func (self SizeofExprNode) GetLocation() Location {
  return self.Location
}

// SizeofTypeNode
type SizeofTypeNode struct {
  Location Location
  Type ITypeNode
  Operand ITypeNode
}

func NewSizeofTypeNode(location Location, t ITypeNode, operand ITypeNode) SizeofTypeNode {
  return SizeofTypeNode { location, t, operand }
}

func (self SizeofTypeNode) String() string {
  panic("not implemented")
}

func (self SizeofTypeNode) IsExpr() bool {
  return true
}

func (self SizeofTypeNode) GetLocation() Location {
  return self.Location
}
