package typesys

import (
  "fmt"
  "reflect"
  "bitbucket.org/yyuu/bs/core"
)

// StructType
type StructType struct {
  ClassName string
  Location core.Location
  Name string
  Members []core.ISlot
}

func NewStructType(name string, membs []core.ISlot, loc core.Location) *StructType {
  return &StructType { "typesys.StructType", loc, name, membs }
}

func (self StructType) Key() string {
  return fmt.Sprintf("struct %s", self.Name)
}

func (self StructType) String() string {
  return self.Key()
}

func (self StructType) Size() int {
  panic("StructType#Size called")
}

func (self StructType) AllocSize() int {
  panic("StructType#AllocSize called")
}

func (self StructType) Alignment() int {
  panic("StructType#Alignment called")
}

func (self StructType) IsSameType(other core.IType) bool {
  if ! other.IsStruct() {
    return false
  } else {
    return reflect.DeepEqual(self, *(other.(*StructType)))
  }
}

func (self StructType) IsVoid() bool {
  return false
}

func (self StructType) IsInteger() bool {
  return false
}

func (self StructType) IsSigned() bool {
  return false
}

func (self StructType) IsPointer() bool {
  return false
}

func (self StructType) IsArray() bool {
  return false
}

func (self StructType) IsCompositeType() bool {
  return false
}

func (self StructType) IsStruct() bool {
  return true
}

func (self StructType) IsUnion() bool {
  return false
}

func (self StructType) IsUserType() bool {
  return false
}

func (self StructType) IsFunction() bool {
  return false
}

func (self StructType) IsAllocatedArray() bool {
  return false
}

func (self StructType) IsIncompleteArray() bool {
  return false
}

func (self StructType) IsScalar() bool {
  return false
}

func (self StructType) IsCallable() bool {
  return false
}

func (self StructType) IsCompatible(target core.IType) bool {
  if !target.IsStruct() {
    return false
  }
  t, ok := target.(StructType)
  if ! ok {
    return false
  }
  ts1, ts2 := self.GetMemberTypes(), t.GetMemberTypes()
  if len(ts1) != len(ts2) {
    return false
  }
  for i := range ts1 {
    if ! ts1[i].IsCompatible(ts2[i]) {
      return false
    }
  }
  return true
}

func (self StructType) IsCastableTo(target core.IType) bool {
  if !target.IsStruct() {
    return false
  }
  t, ok := target.(StructType)
  if ! ok {
    return false
  }
  ts1, ts2 := self.GetMemberTypes(), t.GetMemberTypes()
  if len(ts1) != len(ts2) {
    return false
  }
  for i := range self.Members {
    if ! ts1[i].IsCompatible(ts2[i]) {
      return false
    }
  }
  return true
}

func (self StructType) GetName() string {
  return self.Name
}

func (self StructType) GetMember(name string) core.ISlot {
  for i := range self.Members {
    slot := self.Members[i]
    if slot.GetName() == name {
      return slot
    }
  }
  return nil
}

func (self StructType) GetMembers() []core.ISlot {
  return self.Members
}

func (self StructType) GetMemberType(name string) core.IType {
  slot := self.GetMember(name)
  if slot != nil {
    return slot.GetType()
  } else {
    return nil
  }
}

func (self StructType) GetMemberOffset(name string) int {
  slot := self.GetMember(name)
  if slot == nil {
    panic(fmt.Sprintf("no such member in `%s' in %s", name, self.Name))
  }
  return slot.GetOffset()
}

func (self StructType) GetMemberTypes() []core.IType {
  types := make([]core.IType, len(self.Members))
  for i := range self.Members {
    types[i] = self.Members[i].GetType()
  }
  return types
}

func (self StructType) HasMember(name string) bool {
  slot := self.GetMember(name)
  return slot != nil
}

func (self StructType) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}
