package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type TypeTable struct {
  charSize int
  shortSize int
  intSize int
  longSize int
  ptrSize int
  table map[core.ITypeRef]core.IType
}

func NewTypeTable(charSize, shortSize, intSize, longSize, ptrSize int) *TypeTable {
  loc := core.NewLocation("[builtin:typesys]", 0, 0)
  tt := TypeTable { charSize, shortSize, intSize, longSize, ptrSize, make(map[core.ITypeRef]core.IType) }
  tt.PutType(NewVoidTypeRef(loc),
             NewVoidType())
  tt.PutType(NewIntegerTypeRef(loc, "char"),
             NewIntegerType(charSize, true, "char"))
  tt.PutType(NewIntegerTypeRef(loc, "short"),
             NewIntegerType(shortSize, true, "short"))
  tt.PutType(NewIntegerTypeRef(loc, "int"),
             NewIntegerType(intSize, true, "int"))
  tt.PutType(NewIntegerTypeRef(loc, "long"),
             NewIntegerType(longSize, true, "long"))
  tt.PutType(NewIntegerTypeRef(loc, "unsigned char"),
             NewIntegerType(charSize, true, "unsigned char"))
  tt.PutType(NewIntegerTypeRef(loc, "unsigned short"),
             NewIntegerType(shortSize, true, "unsigned short"))
  tt.PutType(NewIntegerTypeRef(loc, "unsigned int"),
             NewIntegerType(intSize, true, "unsigned int"))
  tt.PutType(NewIntegerTypeRef(loc, "unsigned long"),
             NewIntegerType(longSize, true, "unsigned long"))

  return &tt
}

func NewTypeTableILP32() *TypeTable {
  return NewTypeTable(1, 2, 4, 4, 4)
}

func NewTypeTableILP64() *TypeTable {
  return NewTypeTable(1, 2, 8, 8, 8)
}

func NewTypeTableLP64() *TypeTable {
  return NewTypeTable(1, 2, 4, 8, 8)
}

func NewTypeTableLLP64() *TypeTable {
  return NewTypeTable(1, 2, 4, 4, 8)
}

func NewTypeTableFor(platform string) *TypeTable {
  switch platform {
    case "x86-linux": return NewTypeTableILP32()
    default: panic(fmt.Errorf("unknown platform: %s", platform))
  }
}

func (self *TypeTable) PutType(ref core.ITypeRef, t core.IType) {
  self.table[ref] = t
}

func (self TypeTable) GetType(ref core.ITypeRef) core.IType {
  return self.table[ref]
}

func (self TypeTable) GetCharSize() int {
  return self.charSize
}

func (self TypeTable) GetShortSize() int {
  return self.shortSize
}

func (self TypeTable) GetIntSize() int {
  return self.intSize
}

func (self TypeTable) GetLongSize() int {
  return self.longSize
}

func (self TypeTable) GetPointerSize() int {
  return self.ptrSize
}

func (self TypeTable) IsTypeTable() bool {
  return true
}

func (self TypeTable) String() string {
  return fmt.Sprintf("(let typetable ((charSize %d) (shortSize %d) (intSize %d) (longSize %d) (ptrSize %d)) '())", self.charSize, self.shortSize, self.intSize, self.longSize, self.ptrSize)
}

func (self TypeTable) IsDefined(ref core.ITypeRef) bool {
  _, ok := self.table[ref]
  return ok
}
