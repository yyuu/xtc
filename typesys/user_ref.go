package typesys

import (
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

func (self UserTypeRef) Key() string {
  return self.Name
}

func (self UserTypeRef) String() string {
  return self.Key()
}

func (self UserTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self UserTypeRef) IsTypeRef() bool {
  return true
}

func (self UserTypeRef) GetName() string {
  return self.Name
}
