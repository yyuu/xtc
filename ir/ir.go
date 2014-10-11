package ir

import (
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
)

type IR struct {
  ClassName string
  Location xtc_core.Location
  Defvars []*xtc_entity.DefinedVariable
  Defuns []*xtc_entity.DefinedFunction
  Funcdecls []*xtc_entity.UndefinedFunction
  scope *xtc_entity.ToplevelScope
  constantTable *xtc_entity.ConstantTable
}

func NewIR(loc xtc_core.Location, defvars []*xtc_entity.DefinedVariable, defuns []*xtc_entity.DefinedFunction, funcdecls []*xtc_entity.UndefinedFunction, scope *xtc_entity.ToplevelScope, constantTable *xtc_entity.ConstantTable) *IR {
  return &IR { "ir.IR", loc, defvars, defuns, funcdecls, scope, constantTable }
}

func (self *IR) GetLocation() xtc_core.Location {
  return self.Location
}

func (self *IR) GetFileName() string {
  return self.Location.GetSourceName()
}

func (self *IR) GetDefinedVariables() []*xtc_entity.DefinedVariable {
  return self.Defvars
}

func (self *IR) IsFunctionDefined() bool {
  return 0 < len(self.Defuns)
}

func (self *IR) GetDefinedFunctions() []*xtc_entity.DefinedFunction {
  return self.Defuns
}

func (self *IR) GetScope() *xtc_entity.ToplevelScope {
  return self.scope
}

func (self *IR) AllFunctions() []xtc_core.IFunction {
  fs := []xtc_core.IFunction { }
  for i := range self.Defuns {
    fs = append(fs, self.Defuns[i])
  }
  for i := range self.Funcdecls {
    fs = append(fs, self.Funcdecls[i])
  }
  return fs
}

func (self *IR) AllGlobalVariables() []xtc_core.IVariable {
  return self.scope.AllGlobalVariables()
}

func (self *IR) IsGlobalVariableDefined() bool {
  gvars := self.GetDefinedGlobalVariables()
  return 0 < len(gvars)
}

func (self *IR) GetDefinedGlobalVariables() []*xtc_entity.DefinedVariable {
  gvars := []*xtc_entity.DefinedVariable { }
  vs := self.scope.GetDefinedGlobalScopeVariables()
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

func (self *IR) GetDefinedCommonSymbols() []*xtc_entity.DefinedVariable {
  comms := []*xtc_entity.DefinedVariable { }
  vs := self.scope.GetDefinedGlobalScopeVariables()
  for i := range vs {
    if ! vs[i].HasInitializer() {
      comms = append(comms, vs[i])
    }
  }
  return comms
}

func (self *IR) IsStringLiteralDefined() bool {
  return ! self.constantTable.IsEmpty()
}

func (self *IR) GetConstantTable() *xtc_entity.ConstantTable {
  return self.constantTable
}
