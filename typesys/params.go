package typesys

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// ParamTypes
type ParamTypes struct {
  location duck.ILocation
  paramDescs []duck.IType
  vararg bool
}

func NewParamTypes(loc duck.ILocation, paramDescs []duck.IType, vararg bool) ParamTypes {
  return ParamTypes { loc, paramDescs, vararg }
}

func (self ParamTypes) String() string {
  return fmt.Sprintf("<typesys.ParamTypes Location=%s ParamDescs=%s Vararg=%v>", self.location, self.paramDescs, self.vararg)
}

func (self ParamTypes) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    ParamDescs []duck.IType
    VarArg bool
  }
  x.ClassName = "typesys.ParamTypes"
  x.Location = self.location
  x.ParamDescs = self.paramDescs
  x.VarArg = self.vararg
  return json.Marshal(x)
}

func (self ParamTypes) Size() int {
  panic("ParamTypes#Size called")
}

func (self ParamTypes) AllocSize() int {
  panic("ParamTypes#AllocSize called")
}

func (self ParamTypes) Alignment() int {
  panic("ParamTypes#Alignment called")
}

func (self ParamTypes) IsVoid() bool {
  return false
}

func (self ParamTypes) IsInteger() bool {
  return false
}

func (self ParamTypes) IsSigned() bool {
  return false
}

func (self ParamTypes) IsPointer() bool {
  return false
}

func (self ParamTypes) IsArray() bool {
  return false
}

func (self ParamTypes) IsCompositeType() bool {
  return false
}

func (self ParamTypes) IsStruct() bool {
  return false
}

func (self ParamTypes) IsUnion() bool {
  return false
}

func (self ParamTypes) IsUserType() bool {
  return false
}

func (self ParamTypes) IsFunction() bool {
  return false
}

// ParamTypeRefs
type ParamTypeRefs struct {
  location duck.ILocation
  paramDescs []duck.ITypeRef
  vararg bool
}

func NewParamTypeRefs(loc duck.ILocation, paramDescs []duck.ITypeRef, vararg bool) ParamTypeRefs {
  return ParamTypeRefs { loc, paramDescs, vararg }
}

func (self ParamTypeRefs) String() string {
  return fmt.Sprintf("<typesys.ParamTypeRefs Location=%s ParamDescs=%s Vararg=%v>", self.location, self.paramDescs, self.vararg)
}

func (self ParamTypeRefs) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    ParamDescs []duck.ITypeRef
    Vararg bool
  }
  x.ClassName = "typesys.ParamTypeRefs"
  x.Location = self.location
  x.ParamDescs = self.paramDescs
  x.Vararg = self.vararg
  return json.Marshal(x)
}

func (self ParamTypeRefs) GetLocation() duck.ILocation {
  return self.location
}

func (self ParamTypeRefs) IsTypeRef() bool {
  return true
}
