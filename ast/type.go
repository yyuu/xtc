package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/typesys"
)

// CastNode
type CastNode struct {
  ClassName string
  Location duck.Location
  TypeNode duck.ITypeNode
  Expr duck.IExprNode
}

func NewCastNode(loc duck.Location, t duck.ITypeNode, expr duck.IExprNode) CastNode {
  if t == nil { panic("t is nil") }
  if expr == nil { panic("expr is nil") }
  return CastNode { "ast.CastNode", loc, t, expr }
}

func (self CastNode) String() string {
  return fmt.Sprintf("(%s %s)", self.TypeNode, self.Expr)
}

func (self CastNode) IsExprNode() bool {
  return true
}

func (self CastNode) GetLocation() duck.Location {
  return self.Location
}

// SizeofExprNode
type SizeofExprNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
  TypeNode duck.ITypeNode
}

func NewSizeofExprNode(loc duck.Location, expr duck.IExprNode, t duck.ITypeRef) SizeofExprNode {
  if expr == nil { panic("expr is nil") }
  if t == nil { panic("t is nil") }
  return SizeofExprNode { "ast.SizeofExprNode", loc, expr, NewTypeNode(loc, t) }
}

func (self SizeofExprNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.Expr)
}

func (self SizeofExprNode) IsExprNode() bool {
  return true
}

func (self SizeofExprNode) GetLocation() duck.Location {
  return self.Location
}

// SizeofTypeNode
type SizeofTypeNode struct {
  ClassName string
  Location duck.Location
  TypeNode duck.ITypeNode
  Operand duck.ITypeNode
}

func NewSizeofTypeNode(loc duck.Location, operand duck.ITypeNode, t duck.ITypeRef) SizeofTypeNode {
  if operand == nil { panic("operand is nil") }
  if t == nil { panic("t is nil") }
  return SizeofTypeNode { "ast.SizeofTypeNode", loc, operand, NewTypeNode(loc, t) }
}

func (self SizeofTypeNode) String() string {
  return fmt.Sprintf("(sizeof %s)", self.TypeNode)
}

func (self SizeofTypeNode) IsExprNode() bool {
  return true
}

func (self SizeofTypeNode) GetLocation() duck.Location {
  return self.Location
}

// TypeNode
type TypeNode struct {
  ClassName string
  Location duck.Location
  TypeRef duck.ITypeRef
}

func NewTypeNode(loc duck.Location, t duck.ITypeRef) TypeNode {
  if t == nil { panic("t is nil") }
  return TypeNode { "ast.TypeNode", loc, t }
}

func (self TypeNode) String() string {
  return fmt.Sprintf("(type %s)", self.TypeRef)
}

func (self TypeNode) GetTypeRef() duck.ITypeRef {
  return self.TypeRef
}

func (self TypeNode) IsTypeNode() bool {
  return true
}

func (self TypeNode) GetLocation() duck.Location {
  return self.Location
}

// TypedefNode
type TypedefNode struct {
  ClassName string
  Location duck.Location
  NewType duck.ITypeRef
  Real duck.ITypeRef
  Name string
}

func NewTypedefNode(loc duck.Location, real duck.ITypeRef, name string) TypedefNode {
  if real == nil { panic("real is nil") }
  newType := real
  return TypedefNode { "ast.TypedefNode", loc, newType, real, name }
}

func (self TypedefNode) String() string {
  return fmt.Sprintf("(typedef %s %s)", self.Name, self.Real)
}

func (self TypedefNode) IsTypeDefinition() bool {
  return true
}

func (self TypedefNode) GetLocation() duck.Location {
  return self.Location
}

func (self TypedefNode) GetTypeRef() duck.ITypeRef {
  return self.NewType
}

func (self TypedefNode) DefiningType() duck.IType {
  realTypeNode := NewTypeNode(self.Location, self.Real)
  return typesys.NewUserType(self.Name, realTypeNode, self.Location)
}

// TypeDefinition
type TypeDefinition struct {
  ClassName string
  Location duck.Location
  Name string
  TypeNode duck.ITypeNode
}

func NewTypeDefinition(loc duck.Location, ref duck.ITypeRef, name string) TypeDefinition {
  if ref == nil { panic("ref is nil") }
  return TypeDefinition { "ast.TypeDefinition", loc, name, NewTypeNode(loc, ref) }
}

func (self TypeDefinition) String() string {
  return fmt.Sprintf("<ast.TypeDefinition Name=%s Location=%s TypeNode=%s>", self.Name, self.Location, self.TypeNode)
}

func (self TypeDefinition) IsTypeDefinition() bool {
  return true
}

func (self TypeDefinition) GetLocation() duck.Location {
  return self.Location
}
