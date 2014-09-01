package ast

import (
  "encoding/json"
)

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

func (self CastNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Type ITypeNode
    Expr IExprNode
  }
  x.ClassName = "ast.CastNode"
  x.Location = self.Location
  x.Type = self.Type
  x.Expr = self.Expr
  return json.Marshal(x)
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

func (self SizeofExprNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
    Type ITypeNode
  }
  x.ClassName = "ast.SizeofExprNode"
  x.Location = self.Location
  x.Expr = self.Expr
  x.Type = self.Type
  return json.Marshal(x)
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

func (self SizeofTypeNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Type ITypeNode
    Operand ITypeNode
  }
  x.ClassName = "ast.SizeofTypeNode"
  x.Location = self.Location
  x.Type = self.Type
  x.Operand = self.Operand
  return json.Marshal(x)
}

func (self SizeofTypeNode) IsExpr() bool {
  return true
}

func (self SizeofTypeNode) GetLocation() Location {
  return self.Location
}
