package typesys

import (
  "fmt"
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

func (self UnionType) String() string {
  return fmt.Sprintf("<typesys.UnionType Name=%s Location=%s Members=%s>", self.Name, self.Location, self.Members)
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

//func (self UnionType) GetMemberType(name string) core.IType {
//  slot := self.GetMember(name)
//  if slot != nil {
//    return slot.GetType()
//  } else {
//    return nil
//  }
//}

//func (self UnionType) GetMemberTypes() []core.IType {
//  types := make([]core.IType, len(self.Members))
//  for i := range self.Members {
//    types[i] = self.Members[i].GetType()
//  }
//  return types
//}

func (self UnionType) HasMember(name string) bool {
  slot := self.GetMember(name)
  return slot != nil
}
