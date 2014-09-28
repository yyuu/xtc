package parser

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

func asDefinedFunction(x core.IEntity) *entity.DefinedFunction {
  return x.(*entity.DefinedFunction)
}

func asUndefinedFunction(x core.IEntity) *entity.UndefinedFunction {
  return x.(*entity.UndefinedFunction)
}

func asDefinedVariable(x core.IEntity) *entity.DefinedVariable {
  return x.(*entity.DefinedVariable)
}

func asDefinedVariables(xs []core.IEntity) []*entity.DefinedVariable {
  ys := make([]*entity.DefinedVariable, len(xs))
  for i := range xs {
    ys[i] = asDefinedVariable(xs[i])
  }
  return ys
}

func asUndefinedVariable(x core.IEntity) *entity.UndefinedVariable {
  return x.(*entity.UndefinedVariable)
}

func asConstant(x core.IEntity) *entity.Constant {
  return x.(*entity.Constant)
}

func asParams(x core.IEntity) *entity.Params {
  return x.(*entity.Params)
}

func asParameter(x core.IEntity) *entity.Parameter {
  return x.(*entity.Parameter)
}

func parametersTypeRef(params *entity.Params) *typesys.ParamTypeRefs {
  paramDescs := params.GetParamDescs()
  ps := make([]core.ITypeRef, len(paramDescs))
  for i := range paramDescs {
    ps[i] = paramDescs[i].GetTypeNode().GetTypeRef()
  }
  return typesys.NewParamTypeRefs(params.GetLocation(), ps, params.IsVararg())
}
