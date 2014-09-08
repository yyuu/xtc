package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// CastNode
type CastNode struct {
  location duck.ILocation
  typeNode duck.ITypeNode
  expr duck.IExprNode
}

func NewCastNode(loc duck.ILocation, t duck.ITypeNode, expr duck.IExprNode) CastNode {
  if loc == nil { panic("location is nil") }
  if t == nil { panic("t is nil") }
  if expr == nil { panic("expr is nil") }
  return CastNode { loc, t, expr }
}

func (self CastNode) String() string {
  return fmt.Sprintf("(%s %s)", self.typeNode, self.expr)
}

func (self CastNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    TypeNode duck.ITypeNode
    Expr duck.IExprNode
  }
  x.ClassName = "ast.CastNode"
  x.Location = self.location
  x.TypeNode = self.typeNode
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self CastNode) IsExprNode() bool {
  return true
}

func (self CastNode) GetLocation() duck.ILocation {
  return self.location
}

// SizeofExprNode
type SizeofExprNode struct {
  location duck.ILocation
  expr duck.IExprNode
  typeNode duck.ITypeNode
}

func NewSizeofExprNode(loc duck.ILocation, expr duck.IExprNode, t duck.ITypeRef) SizeofExprNode {
  if loc == nil { panic("location is nil") }
  if expr == nil { panic("expr is nil") }
  if t == nil { panic("t is nil") }
  return SizeofExprNode { loc, expr, NewTypeNode(loc, t) }
}

func (self SizeofExprNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.expr)
}

func (self SizeofExprNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
    TypeNode duck.ITypeNode
  }
  x.ClassName = "ast.SizeofExprNode"
  x.Location = self.location
  x.Expr = self.expr
  x.TypeNode = self.typeNode
  return json.Marshal(x)
}

func (self SizeofExprNode) IsExprNode() bool {
  return true
}

func (self SizeofExprNode) GetLocation() duck.ILocation {
  return self.location
}

// SizeofTypeNode
type SizeofTypeNode struct {
  location duck.ILocation
  typeNode duck.ITypeNode
  operand duck.ITypeNode
}

func NewSizeofTypeNode(loc duck.ILocation, operand duck.ITypeNode, t duck.ITypeRef) SizeofTypeNode {
  if loc == nil { panic("location is nil") }
  if operand == nil { panic("operand is nil") }
  if t == nil { panic("t is nil") }
  return SizeofTypeNode { loc, operand, NewTypeNode(loc, t) }
}

func (self SizeofTypeNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.typeNode)
}

func (self SizeofTypeNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    TypeNode duck.ITypeNode
    Operand duck.ITypeNode
  }
  x.ClassName = "ast.SizeofTypeNode"
  x.Location = self.location
  x.TypeNode = self.typeNode
  x.Operand = self.operand
  return json.Marshal(x)
}

func (self SizeofTypeNode) IsExprNode() bool {
  return true
}

func (self SizeofTypeNode) GetLocation() duck.ILocation {
  return self.location
}

// TypeNode
type TypeNode struct {
  location duck.ILocation
  typeRef duck.ITypeRef
}

func NewTypeNode(loc duck.ILocation, t duck.ITypeRef) TypeNode {
  if loc == nil { panic("location is nil") }
  if t == nil { panic("t is nil") }
  return TypeNode { loc, t }
}

func (self TypeNode) String() string {
  return fmt.Sprintf("(type %s)", self.typeRef)
}

func (self TypeNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    TypeRef duck.ITypeRef
  }
  x.ClassName = "ast.TypeNode"
  x.Location = self.location
  x.TypeRef = self.typeRef
  return json.Marshal(x)
}

func (self TypeNode) GetTypeRef() duck.ITypeRef {
  return self.typeRef
}

func (self TypeNode) IsTypeNode() bool {
  return true
}

func (self TypeNode) GetLocation() duck.ILocation {
  return self.location
}

// TypedefNode
type TypedefNode struct {
  location duck.ILocation
  real duck.ITypeRef
  name string
}

func NewTypedefNode(loc duck.ILocation, real duck.ITypeRef, name string) TypedefNode {
  if loc == nil { panic("location is nil") }
  if real == nil { panic("real is nil") }
  return TypedefNode { loc, real, name }
}

func (self TypedefNode) String() string {
  return fmt.Sprintf("(typedef %s %s)", self.name, self.real)
}

func (self TypedefNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Real duck.ITypeRef
    Name string
  }
  x.ClassName = "ast.TypedefNode"
  x.Location = self.location
  x.Real = self.real
  x.Name = self.name
  return json.Marshal(x)
}

func (self TypedefNode) IsTypeDefinition() bool {
  return true
}

func (self TypedefNode) GetLocation() duck.ILocation {
  return self.location
}

// TypeDefinition
type TypeDefinition struct {
  location duck.ILocation
  name string
  typeNode duck.ITypeNode
}

func NewTypeDefinition(loc duck.ILocation, ref duck.ITypeRef, name string) TypeDefinition {
  if loc == nil { panic("location is nil") }
  if ref == nil { panic("ref is nil") }
  return TypeDefinition { loc, name, NewTypeNode(loc, ref) }
}

func (self TypeDefinition) String() string {
  return fmt.Sprintf("<ast.TypeDefinition Name=%s Location=%s TypeNode=%s>", self.name, self.location, self.typeNode)
}

func (self TypeDefinition) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
    TypeNode duck.ITypeNode
  }
  x.ClassName = "ast.TypeDefinition"
  x.Location = self.location
  x.Name = self.name
  x.TypeNode = self.typeNode
  return json.Marshal(x)
}

func (self TypeDefinition) IsTypeDefinition() bool {
  return true
}

func (self TypeDefinition) GetLocation() duck.ILocation {
  return self.location
}
