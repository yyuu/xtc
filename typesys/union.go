package typesys

import (
  "fmt"
  "reflect"
  "bitbucket.org/yyuu/bs/core"
)

// UnionType
type UnionType struct {
  ClassName string
  Location core.Location
  Name string
  Members []core.ISlot
}

func NewUnionType(name string, membs []core.ISlot, loc core.Location) *UnionType {
  return &UnionType { "typesys.UnionType", loc, name, membs }
}

func (self UnionType) Key() string {
  return fmt.Sprintf("union %s", self.Name)
}

func (self UnionType) String() string {
  return self.Key()
}

func (self UnionType) Size() int {
  panic("UnionType#Size called")
}

func (self UnionType) AllocSize() int {
  panic("UnionType#AllocSize called")
}

func (self UnionType) Alignment() int {
  panic("UnionType#Alignment called")
}

func (self UnionType) IsSameType(other core.IType) bool {
  if !other.IsUnion() {
    return false
  } else {
    return reflect.DeepEqual(self, *(other.(*UnionType)))
  }
}

func (self UnionType) IsVoid() bool {
  return false
}

func (self UnionType) IsInteger() bool {
  return false
}

func (self UnionType) IsSigned() bool {
  return false
}

func (self UnionType) IsPointer() bool {
  return false
}

func (self UnionType) IsArray() bool {
  return false
}

func (self UnionType) IsCompositeType() bool {
  return false
}

func (self UnionType) IsStruct() bool {
  return false
}

func (self UnionType) IsUnion() bool {
  return true
}

func (self UnionType) IsUserType() bool {
  return false
}

func (self UnionType) IsFunction() bool {
  return false
}

func (self UnionType) IsAllocatedArray() bool {
  return false
}

func (self UnionType) IsIncompleteArray() bool {
  return false
}

func (self UnionType) IsScalar() bool {
  return false
}

func (self UnionType) IsCallable() bool {
  return false
}

func (self UnionType) IsCompatible(target core.IType) bool {
  if !target.IsUnion() {
    return false
  }
  t, ok := target.(UnionType)
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

func (self UnionType) IsCastableTo(target core.IType) bool {
  if !target.IsUnion() {
    return false
  }
  t, ok := target.(UnionType)
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

func (self UnionType) GetName() string {
  return self.Name
}

func (self UnionType) GetMember(name string) core.ISlot {
  for i := range self.Members {
    slot := self.Members[i]
    if slot.GetName() == name {
      return slot
    }
  }
  return nil
}

func (self UnionType) GetMembers() []core.ISlot {
  return self.Members
}

func (self UnionType) GetMemberType(name string) core.IType {
  slot := self.GetMember(name)
  if slot != nil {
    return slot.GetType()
  } else {
    return nil
  }
}

func (self UnionType) GetMemberOffset(name string) int {
  slot := self.GetMember(name)
  if slot == nil {
    panic(fmt.Sprintf("no such member in `%s' in %s", name, self.Name))
  }
  return slot.GetOffset()
}

func (self UnionType) GetMemberTypes() []core.IType {
  types := make([]core.IType, len(self.Members))
  for i := range self.Members {
    types[i] = self.Members[i].GetType()
  }
  return types
}

func (self UnionType) HasMember(name string) bool {
  slot := self.GetMember(name)
  return slot != nil
}

func (self UnionType) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}
