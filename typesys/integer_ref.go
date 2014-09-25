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

func NewIntegerTypeRef(loc core.Location, name string) *IntegerTypeRef {
  return &IntegerTypeRef { "typesys.IntegerTypeRef", loc, name }
}

func NewCharTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "char")
}

func NewShortTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "short")
}

func NewIntTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "int")
}

func NewLongTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "long")
}

func NewUnsignedCharTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "unsigned char")
}

func NewUnsignedShortTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "unsigned short")
}

func NewUnsignedIntTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "unsigned int")
}

func NewUnsignedLongTypeRef(loc core.Location) *IntegerTypeRef {
  return NewIntegerTypeRef(loc, "unsigned long")
}

func (self IntegerTypeRef) Key() string {
  return self.Name
}

func (self IntegerTypeRef) String() string {
  return self.Key()
}

func (self IntegerTypeRef) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self IntegerTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self IntegerTypeRef) IsTypeRef() bool {
  return true
}
