package parser

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

func parametersTypeRef(params *entity.Params) *typesys.ParamTypeRefs {
  paramDescs := params.GetParamDescs()
  ps := make([]core.ITypeRef, len(paramDescs))
  for i := range paramDescs {
    ps[i] = paramDescs[i].GetTypeNode().GetTypeRef()
  }
  return typesys.NewParamTypeRefs(params.GetLocation(), ps, params.IsVararg())
}
