package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// CastNode
type CastNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Expr core.IExprNode
}

func NewCastNode(loc core.Location, t core.ITypeNode, expr core.IExprNode) CastNode {
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

func (self CastNode) GetLocation() core.Location {
  return self.Location
}

// SizeofExprNode
type SizeofExprNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  TypeNode core.ITypeNode
}

func NewSizeofExprNode(loc core.Location, expr core.IExprNode, t core.ITypeRef) SizeofExprNode {
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

func (self SizeofExprNode) GetLocation() core.Location {
  return self.Location
}

// SizeofTypeNode
type SizeofTypeNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Operand core.ITypeNode
}

func NewSizeofTypeNode(loc core.Location, operand core.ITypeNode, t core.ITypeRef) SizeofTypeNode {
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

func (self SizeofTypeNode) GetLocation() core.Location {
  return self.Location
}

// TypeNode
type TypeNode struct {
  ClassName string
  Location core.Location
  TypeRef core.ITypeRef
}

func NewTypeNode(loc core.Location, t core.ITypeRef) TypeNode {
  if t == nil { panic("t is nil") }
  return TypeNode { "ast.TypeNode", loc, t }
}

func (self TypeNode) String() string {
  return fmt.Sprintf("(type %s)", self.TypeRef)
}

func (self TypeNode) GetTypeRef() core.ITypeRef {
  return self.TypeRef
}

func (self TypeNode) IsTypeNode() bool {
  return true
}

func (self TypeNode) GetLocation() core.Location {
  return self.Location
}

// TypedefNode
type TypedefNode struct {
  ClassName string
  Location core.Location
  NewType core.ITypeRef
  Real core.ITypeRef
  Name string
}

func NewTypedefNode(loc core.Location, real core.ITypeRef, name string) TypedefNode {
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

func (self TypedefNode) GetLocation() core.Location {
  return self.Location
}

func (self TypedefNode) GetTypeRef() core.ITypeRef {
  return self.NewType
}

func (self TypedefNode) DefiningType() core.IType {
  realTypeNode := NewTypeNode(self.Location, self.Real)
  return typesys.NewUserType(self.Name, realTypeNode, self.Location)
}

// TypeDefinition
type TypeDefinition struct {
  ClassName string
  Location core.Location
  Name string
  TypeNode core.ITypeNode
}

func NewTypeDefinition(loc core.Location, ref core.ITypeRef, name string) TypeDefinition {
  if ref == nil { panic("ref is nil") }
  return TypeDefinition { "ast.TypeDefinition", loc, name, NewTypeNode(loc, ref) }
}

func (self TypeDefinition) String() string {
  return fmt.Sprintf("<ast.TypeDefinition Name=%s Location=%s TypeNode=%s>", self.Name, self.Location, self.TypeNode)
}

func (self TypeDefinition) IsTypeDefinition() bool {
  return true
}

func (self TypeDefinition) GetLocation() core.Location {
  return self.Location
}
