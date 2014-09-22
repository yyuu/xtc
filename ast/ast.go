package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type AST struct {
  ClassName string
  Location core.Location
  Declaration *Declaration
  scope *entity.ToplevelScope
  constantTable *entity.ConstantTable
}

func NewAST(loc core.Location, declaration *Declaration) *AST {
  return &AST { "ast.AST", loc, declaration, nil, nil }
}

func (self AST) String() string {
  return fmt.Sprintf(";; %s\n%s", self.Location, self.Declaration)
}

func (self AST) GetLocation() core.Location {
  return self.Location
}

func (self AST) GetDeclaration() *Declaration {
  return self.Declaration
}

func (self AST) ListTypes() []core.ITypeDefinition {
  var result []core.ITypeDefinition
  decl := self.Declaration
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
  decl := self.Declaration
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

func (self AST) ListDeclarations() []core.IEntity {
  var result []core.IEntity
  decl := self.Declaration
  for i := range decl.Funcdecls {
    result = append(result, decl.Funcdecls[i])
  }
  for i := range decl.Vardecls {
    result = append(result, decl.Vardecls[i])
  }
  return result
}

func (self AST) ListDefinitions() []core.IEntity {
  var result []core.IEntity
  decl := self.Declaration
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

func (self AST) GetDefinedVariables() []*entity.DefinedVariable {
  return self.Declaration.Defvars
}

func (self *AST) SetDefinedVariables(xs []*entity.DefinedVariable) {
  self.Declaration.Defvars = xs
}

func (self AST) GetDefinedFunctions() []*entity.DefinedFunction {
  return self.Declaration.Defuns
}

func (self *AST) SetDefinedFunctions(xs []*entity.DefinedFunction) {
  self.Declaration.Defuns = xs
}

func (self AST) GetConstants() []*entity.Constant {
  return self.Declaration.Constants
}

func (self *AST) SetConstants(xs []*entity.Constant) {
  self.Declaration.Constants = xs
}

func (self AST) GetScope() *entity.ToplevelScope {
  return self.scope
}

func (self *AST) SetScope(scope *entity.ToplevelScope) {
  self.scope = scope
}

func (self AST) GetConstantTable() *entity.ConstantTable {
  return self.constantTable
}

func (self *AST) SetConstantTable(table *entity.ConstantTable) {
  self.constantTable = table
}
