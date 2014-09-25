package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// UnionTypeRef
type UnionTypeRef struct {
  ClassName string
  Location core.Location
  Name string
}

func NewUnionTypeRef(loc core.Location, name string) *UnionTypeRef {
  return &UnionTypeRef { "typesys.UnionTypeRef", loc, name }
}

func (self UnionTypeRef) Key() string {
  return fmt.Sprintf("union %s", self.Name)
}

func (self UnionTypeRef) String() string {
  return self.Key()
}

func (self UnionTypeRef) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self UnionTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self UnionTypeRef) IsTypeRef() bool {
  return true
}
