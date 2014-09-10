package typesys

import (
  "fmt"
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
  return fmt.Sprintf("<typesys.ParamTypeRefs Location=%s ParamDescs=%s Vararg=%v>", self.Location, self.ParamDescs, self.Vararg)
}

func (self ParamTypeRefs) GetLocation() core.Location {
  return self.Location
}

func (self ParamTypeRefs) IsTypeRef() bool {
  return true
}
