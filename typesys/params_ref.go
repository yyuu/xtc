package typesys

import (
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

// ParamTypeRefs
type ParamTypeRefs struct {
  ClassName string
  Location core.Location
  ParamDescs []core.ITypeRef
  Vararg bool
}

func NewParamTypeRefs(loc core.Location, paramDescs []core.ITypeRef, vararg bool) *ParamTypeRefs {
  return &ParamTypeRefs { "typesys.ParamTypeRefs", loc, paramDescs, vararg }
}

func (self ParamTypeRefs) String() string {
  params := make([]string, len(self.ParamDescs))
  for i := range self.ParamDescs {
    params[i] = self.ParamDescs[i].String()
  }
  return strings.Join(params, ",")
}

func (self ParamTypeRefs) GetLocation() core.Location {
  return self.Location
}

func (self ParamTypeRefs) IsTypeRef() bool {
  return true
}
