package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type AST struct {
  ClassName string
  Location duck.Location
  Declarations Declarations
  scope duck.IVariableScope
  constantTable duck.IConstantTable
}

func NewAST(loc duck.Location, declarations Declarations) AST {
  return AST { "ast.AST", loc, declarations, nil, nil }
}

func (self AST) String() string {
  return fmt.Sprintf(";; %s\n%s", self.Location, self.Declarations)
}

func (self AST) GetLocation() duck.Location {
  return self.Location
}

func (self AST) GetDeclarations() Declarations {
  return self.Declarations
}

func (self AST) ListTypes() []duck.ITypeDefinition {
  var result []duck.ITypeDefinition
  decl := self.Declarations
  for i := range decl.Defstructs {
    result = append(result, decl.Defstructs[i])
  }
  for i := range decl.Defunions {
    result = append(result, decl.Defunions[i])
  }
  for i := range decl.Typedefs {
    result = append(result, decl.Typedefs[i])
  }
  return result
}

func (self AST) ListEntities() []duck.IEntity {
  var result []duck.IEntity
  decl := self.Declarations
  for i := range decl.Funcdecls {
    result = append(result, decl.Funcdecls[i])
  }
  for i := range decl.Vardecls {
    result = append(result, decl.Vardecls[i])
  }
  for i := range decl.Defvars {
    result = append(result, decl.Defvars[i])
  }
  for i := range decl.Defuns {
    result = append(result, decl.Defuns[i])
  }
  for i := range decl.Constants {
    result = append(result, decl.Constants[i])
  }
  return result
}

func (self AST) ListDeclaration() []duck.IEntity {
  var result []duck.IEntity
  decl := self.Declarations
  for i := range decl.Funcdecls {
    result = append(result, decl.Funcdecls[i])
  }
  for i := range decl.Vardecls {
    result = append(result, decl.Vardecls[i])
  }
  return result
}

func (self AST) ListDefinition() []duck.IEntity {
  var result []duck.IEntity
  decl := self.Declarations
  for i := range decl.Defvars {
    result = append(result, decl.Defvars[i])
  }
  for i := range decl.Defuns {
    result = append(result, decl.Defuns[i])
  }
  for i := range decl.Constants {
    result = append(result, decl.Constants[i])
  }
  return result
}

func (self AST) GetDefinedVariables() []duck.IDefinedVariable {
  return self.Declarations.Defvars
}

func (self *AST) SetDefinedVariables(xs []duck.IDefinedVariable) {
  self.Declarations.Defvars = xs
}

func (self AST) GetDefinedFunctions() []duck.IDefinedFunction {
  return self.Declarations.Defuns
}

func (self *AST) SetDefinedFunctions(xs []duck.IDefinedFunction) {
  self.Declarations.Defuns = xs
}

func (self AST) GetConstants() []duck.IConstant {
  return self.Declarations.Constants
}

func (self *AST) SetConstants(xs []duck.IConstant) {
  self.Declarations.Constants = xs
}

func (self AST) GetScope() duck.IVariableScope {
  return self.scope
}

func (self *AST) SetScope(scope duck.IVariableScope) {
  self.scope = scope
}

func (self AST) GetConstantTable() duck.IConstantTable {
  return self.constantTable
}

func (self *AST) SetConstantTable(table duck.IConstantTable) {
  self.constantTable = table
}
