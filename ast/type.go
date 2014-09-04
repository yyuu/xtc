package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// CastNode
type CastNode struct {
  Location duck.ILocation
  Type duck.ITypeNode
  Expr duck.IExprNode
}

func NewCastNode(location duck.ILocation, t duck.ITypeNode, expr duck.IExprNode) CastNode {
  return CastNode { location, t, expr }
}

func (self CastNode) String() string {
  return fmt.Sprintf("(%s %s)", self.Type, self.Expr)
}

func (self CastNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Type duck.ITypeNode
    Expr duck.IExprNode
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

func (self CastNode) GetLocation() duck.ILocation {
  return self.Location
}

// SizeofExprNode
type SizeofExprNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
  Type duck.ITypeNode
}

func NewSizeofExprNode(location duck.ILocation, expr duck.IExprNode, t duck.ITypeRef) SizeofExprNode {
  return SizeofExprNode { location, expr, NewTypeNode(location, t) }
}

func (self SizeofExprNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.Expr)
}

func (self SizeofExprNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    Type duck.ITypeNode
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

func (self SizeofExprNode) GetLocation() duck.ILocation {
  return self.Location
}

// SizeofTypeNode
type SizeofTypeNode struct {
  Location duck.ILocation
  Type duck.ITypeNode
  Operand duck.ITypeNode
}

func NewSizeofTypeNode(location duck.ILocation, operand duck.ITypeNode, t duck.ITypeRef) SizeofTypeNode {
  return SizeofTypeNode { location, operand, NewTypeNode(location, t) }
}

func (self SizeofTypeNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.Type)
}

func (self SizeofTypeNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Type duck.ITypeNode
    Operand duck.ITypeNode
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

func (self SizeofTypeNode) GetLocation() duck.ILocation {
  return self.Location
}

// TypeNode
type TypeNode struct {
  Location duck.ILocation
  TypeRef duck.ITypeRef
}

func NewTypeNode(location duck.ILocation, t duck.ITypeRef) TypeNode {
  return TypeNode { location, t }
}

func (self TypeNode) String() string {
  return fmt.Sprintf("(type %s)", self.TypeRef)
}

func (self TypeNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    TypeRef duck.ITypeRef
  }
  x.ClassName = "ast.TypeNode"
  x.Location = self.Location
  x.TypeRef = self.TypeRef
  return json.Marshal(x)
}

func (self TypeNode) GetTypeRef() duck.ITypeRef {
  return self.TypeRef
}

func (self TypeNode) IsType() bool {
  return true
}

func (self TypeNode) GetLocation() duck.ILocation {
  return self.Location
}

// TypedefNode
type TypedefNode struct {
  Location duck.ILocation
  Real duck.ITypeRef
  Name string
}

func NewTypedefNode(location duck.ILocation, real duck.ITypeRef, name string) TypedefNode {
  return TypedefNode { location, real, name }
}

func (self TypedefNode) String() string {
  return fmt.Sprintf("(typedef %s %s)", self.Name, self.Real)
}

func (self TypedefNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Real duck.ITypeRef
    Name string
  }
  x.ClassName = "ast.TypedefNode"
  x.Location = self.Location
  x.Real = self.Real
  x.Name = self.Name
  return json.Marshal(x)
}

func (self TypedefNode) IsTypeDefinition() bool {
  return true
}

func (self TypedefNode) GetLocation() duck.ILocation {
  return self.Location
}

// TypeDefinition
type TypeDefinition struct {
  Location duck.ILocation
  Name string
  TypeNode duck.ITypeNode
}

func NewTypeDefinition(loc duck.ILocation, ref duck.ITypeRef, name string) TypeDefinition {
  return TypeDefinition { loc, name, NewTypeNode(loc, ref) }
}

func (self TypeDefinition) String() string {
  panic("not implemented")
}

func (self TypeDefinition) MarshalJSON() ([]byte, error) {
  panic("not implemented")
}

func (self TypeDefinition) IsTypeDefinition() bool {
  return true
}

func (self TypeDefinition) GetLocation() duck.ILocation {
  return self.Location
}
