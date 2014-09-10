package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type AST struct {
  ClassName string
  Location core.Location
  Declarations Declarations
  scope core.IVariableScope
  constantTable core.IConstantTable
}

func NewAST(loc core.Location, declarations Declarations) AST {
  return AST { "ast.AST", loc, declarations, nil, nil }
}

func (self AST) String() string {
  return fmt.Sprintf(";; %s\n%s", self.Location, self.Declarations)
}

func (self AST) GetLocation() core.Location {
  return self.Location
}

func (self AST) GetDeclarations() Declarations {
  return self.Declarations
}

func (self AST) ListTypes() []core.ITypeDefinition {
  var result []core.ITypeDefinition
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

func (self AST) ListEntities() []core.IEntity {
  var result []core.IEntity
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

func (self AST) ListDeclaration() []core.IEntity {
  var result []core.IEntity
  decl := self.Declarations
  for i := range decl.Funcdecls {
    result = append(result, decl.Funcdecls[i])
  }
  for i := range decl.Vardecls {
    result = append(result, decl.Vardecls[i])
  }
  return result
}

func (self AST) ListDefinition() []core.IEntity {
  var result []core.IEntity
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

func (self AST) GetDefinedVariables() []core.IDefinedVariable {
  return self.Declarations.Defvars
}

func (self *AST) SetDefinedVariables(xs []core.IDefinedVariable) {
  self.Declarations.Defvars = xs
}

func (self AST) GetDefinedFunctions() []core.IDefinedFunction {
  return self.Declarations.Defuns
}

func (self *AST) SetDefinedFunctions(xs []core.IDefinedFunction) {
  self.Declarations.Defuns = xs
}

func (self AST) GetConstants() []core.IConstant {
  return self.Declarations.Constants
}

func (self *AST) SetConstants(xs []core.IConstant) {
  self.Declarations.Constants = xs
}

func (self AST) GetScope() core.IVariableScope {
  return self.scope
}

func (self *AST) SetScope(scope core.IVariableScope) {
  self.scope = scope
}

func (self AST) GetConstantTable() core.IConstantTable {
  return self.constantTable
}

func (self *AST) SetConstantTable(table core.IConstantTable) {
  self.constantTable = table
}
