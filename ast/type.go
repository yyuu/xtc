package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/typesys"
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
  return fmt.Sprintf("(%s %s)", self.Type, self.Expr)
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

func NewSizeofExprNode(location Location, expr IExprNode, t typesys.ITypeRef) SizeofExprNode {
  return SizeofExprNode { location, expr, NewTypeNode(location, t) }
}

func (self SizeofExprNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.Expr)
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

func NewSizeofTypeNode(location Location, operand ITypeNode, t typesys.ITypeRef) SizeofTypeNode {
  return SizeofTypeNode { location, operand, NewTypeNode(location, t) }
}

func (self SizeofTypeNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.Type)
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

// TypeNode
type TypeNode struct {
  Location Location
  TypeRef typesys.ITypeRef
}

func NewTypeNode(location Location, t typesys.ITypeRef) TypeNode {
  return TypeNode { location, t }
}

func (self TypeNode) String() string {
  return fmt.Sprintf("(type %s)", self.TypeRef)
}

func (self TypeNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    TypeRef typesys.ITypeRef
  }
  x.ClassName = "ast.TypeNode"
  x.Location = self.Location
  x.TypeRef = self.TypeRef
  return json.Marshal(x)
}

func (self TypeNode) IsType() bool {
  return true
}

func (self TypeNode) GetLocation() Location {
  return self.Location
}

// TypedefNode
type TypedefNode struct {
  Location Location
  Real typesys.ITypeRef
  Name string
}

func NewTypedefNode(location Location, real typesys.ITypeRef, name string) TypedefNode {
  return TypedefNode { location, real, name }
}

func (self TypedefNode) String() string {
  return fmt.Sprintf("(typedef %s %s)", self.Name, self.Real)
}

func (self TypedefNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Real typesys.ITypeRef
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

func (self TypedefNode) GetLocation() Location {
  return self.Location
}

// TypeDefinition
type TypeDefinition struct {
  Location Location
  Name string
  TypeNode ITypeNode
}

func NewTypeDefinition(loc Location, ref typesys.ITypeRef, name string) TypeDefinition {
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

func (self TypeDefinition) GetLocation() Location {
  return self.Location
}
