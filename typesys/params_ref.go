package typesys

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/xtc/core"
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

func (self ParamTypeRefs) Key() string {
  params := make([]string, len(self.ParamDescs))
  for i := range self.ParamDescs {
    params[i] = self.ParamDescs[i].String()
  }
  return strings.Join(params, ",")
}

func (self ParamTypeRefs) String() string {
  return self.Key()
}

func (self ParamTypeRefs) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self ParamTypeRefs) GetLocation() core.Location {
  return self.Location
}

func (self ParamTypeRefs) IsTypeRef() bool {
  return true
}

func (self ParamTypeRefs) GetParamDescs() []core.ITypeRef {
  return self.ParamDescs
}

func (self ParamTypeRefs) IsVararg() bool {
  return self.Vararg
}
