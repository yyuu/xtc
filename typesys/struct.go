package typesys

import (
  "fmt"
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

func (self StructType) String() string {
  return fmt.Sprintf("struct %s", self.Name)
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

func (self StructType) IsCallable() bool {
  return false
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

//func (self StructType) GetMemberType(name string) core.IType {
//  slot := self.GetMember(name)
//  if slot != nil {
//    return slot.GetType()
//  } else {
//    return nil
//  }
//}

//func (self StructType) GetMemberTypes() []core.IType {
//  types := make([]core.IType, len(self.Members))
//  for i := range self.Members {
//    types[i] = self.Members[i].GetType()
//  }
//  return types
//}

func (self StructType) HasMember(name string) bool {
  slot := self.GetMember(name)
  return slot != nil
}
