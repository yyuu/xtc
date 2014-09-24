package ir

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type IR struct {
  ClassName string
  Source core.Location
  Defvars []*entity.DefinedVariable
  Defuns []*entity.DefinedFunction
  Funcdecls []*entity.UndefinedFunction
  Scope *entity.ToplevelScope
  ConstantTable *entity.ConstantTable
}

func NewIR(source core.Location, defvars []*entity.DefinedVariable, defuns []*entity.DefinedFunction, funcdecls []*entity.UndefinedFunction, scope *entity.ToplevelScope, constantTable *entity.ConstantTable) *IR {
  return &IR { "ir.IR", source, defvars, defuns, funcdecls, scope, constantTable }
}
