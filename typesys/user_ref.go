package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// UserTypeRef
type UserTypeRef struct {
  ClassName string
  Location core.Location
  Name string
}

func NewUserTypeRef(loc core.Location, name string) *UserTypeRef {
  return &UserTypeRef { "typesys.UserTypeRef", loc, name }
}

func (self UserTypeRef) String() string {
  return fmt.Sprintf("<typesys.UserTypeRef Name=%s Location=%s>", self.Name, self.Location)
}

func (self UserTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self UserTypeRef) IsTypeRef() bool {
  return true
}
