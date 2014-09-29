package ir

import (
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
)

type IR struct {
  ClassName string
  Location bs_core.Location
  Defvars []*bs_entity.DefinedVariable
  Defuns []*bs_entity.DefinedFunction
  Funcdecls []*bs_entity.UndefinedFunction
  Scope *bs_entity.ToplevelScope
  ConstantTable *bs_entity.ConstantTable
}

func NewIR(loc bs_core.Location, defvars []*bs_entity.DefinedVariable, defuns []*bs_entity.DefinedFunction, funcdecls []*bs_entity.UndefinedFunction, scope *bs_entity.ToplevelScope, constantTable *bs_entity.ConstantTable) *IR {
  return &IR { "ir.IR", loc, defvars, defuns, funcdecls, scope, constantTable }
}

func (self *IR) GetLocation() bs_core.Location {
  return self.Location
}

func (self *IR) GetFileName() string {
  return self.Location.GetSourceName()
}

func (self *IR) GetDefinedVariables() []*bs_entity.DefinedVariable {
  return self.Defvars
}

func (self *IR) IsFunctionDefined() bool {
  return 0 < len(self.Defuns)
}

func (self *IR) GetDefinedFunctions() []*bs_entity.DefinedFunction {
  return self.Defuns
}

func (self *IR) GetScope() *bs_entity.ToplevelScope {
  return self.Scope
}

func (self *IR) AllFunctions() []bs_core.IFunction {
  fs := []bs_core.IFunction { }
  for i := range self.Defuns {
    fs = append(fs, self.Defuns[i])
  }
  for i := range self.Funcdecls {
    fs = append(fs, self.Funcdecls[i])
  }
  return fs
}

func (self *IR) AllGlobalVariables() []bs_core.IVariable {
  return self.Scope.AllGlobalVariables()
}

func (self *IR) IsGlobalVariableDefined() bool {
  gvars := self.GetDefinedGlobalVariables()
  return 0 < len(gvars)
}

func (self *IR) GetDefinedGlobalVariables() []*bs_entity.DefinedVariable {
  gvars := []*bs_entity.DefinedVariable { }
  vs := self.Scope.GetDefinedGlobalScopeVariables()
  for i := range vs {
    if vs[i].HasInitializer() {
      gvars = append(gvars, vs[i])
    }
  }
  return gvars
}

func (self *IR) IsCommonSymbolDefined() bool {
  comms := self.GetDefinedCommonSymbols()
  return 0 < len(comms)
}

func (self *IR) GetDefinedCommonSymbols() []*bs_entity.DefinedVariable {
  comms := []*bs_entity.DefinedVariable { }
  vs := self.Scope.GetDefinedGlobalScopeVariables()
  for i := range vs {
    if ! vs[i].HasInitializer() {
      comms = append(comms, vs[i])
    }
  }
  return comms
}

func (self *IR) IsStringLiteralDefined() bool {
  return ! self.ConstantTable.IsEmpty()
}

func (self *IR) GetConstantTable() *bs_entity.ConstantTable {
  return self.ConstantTable
}
