package ir

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type IR struct {
  source core.Location
  defvars []*entity.DefinedVariable
  defuns []*entity.DefinedFunction
  funcdecls []*entity.UndefinedFunction
  scope *entity.ToplevelScope
  constantTable *entity.ConstantTable
}

func NewIR(source core.Location, defvars []*entity.DefinedVariable, defuns []*entity.DefinedFunction, funcdecls []*entity.UndefinedFunction, scope *entity.ToplevelScope, constantTable *entity.ConstantTable) *IR {
  return &IR { source, defvars, defuns, funcdecls, scope, constantTable }
}
