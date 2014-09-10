package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// IntegerTypeRef
type IntegerTypeRef struct {
  ClassName string
  Location core.Location
  Name string
}

func NewIntegerTypeRef(loc core.Location, name string) IntegerTypeRef {
  return IntegerTypeRef { "typesys.IntegerTypeRef", loc, name }
}

func (self IntegerTypeRef) String() string {
  return fmt.Sprintf("<typesys.IntegerTypeRef Name=%s Location=%s>", self.Name, self.Location)
}

func (self IntegerTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self IntegerTypeRef) IsTypeRef() bool {
  return true
}
