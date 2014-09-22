package typesys

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

type TypeTable struct {
  charSize int
  shortSize int
  intSize int
  longSize int
  ptrSize int
  table map[string]core.IType
}

func NewTypeTable(charSize, shortSize, intSize, longSize, ptrSize int) *TypeTable {
  loc := core.NewLocation("[builtin:typesys]", 0, 0)
  tt := TypeTable { charSize, shortSize, intSize, longSize, ptrSize, make(map[string]core.IType) }
  tt.PutType(NewVoidTypeRef(loc), NewVoidType())
  tt.PutType(NewCharTypeRef(loc), NewCharType(charSize))
  tt.PutType(NewShortTypeRef(loc), NewShortType(shortSize))
  tt.PutType(NewIntTypeRef(loc), NewIntType(intSize))
  tt.PutType(NewLongTypeRef(loc), NewLongType(longSize))
  tt.PutType(NewUnsignedCharTypeRef(loc), NewUnsignedCharType(charSize))
  tt.PutType(NewUnsignedShortTypeRef(loc), NewUnsignedShortType(shortSize))
  tt.PutType(NewUnsignedIntTypeRef(loc), NewUnsignedIntType(intSize))
  tt.PutType(NewUnsignedLongTypeRef(loc), NewUnsignedLongType(longSize))
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
  self.table[ref.String()] = t
}

func (self TypeTable) GetType(ref core.ITypeRef) core.IType {
  t := self.table[ref.String()]
  if t == nil {
    switch typed := ref.(type) {
      case *UserTypeRef: {
        panic(fmt.Errorf("undefined type: %s", typed.GetName()))
      }
      case *PointerTypeRef: {
        t = NewPointerType(self.ptrSize, self.GetType(typed.GetBaseType()))
        self.PutType(typed, t)
      }
      case *ArrayTypeRef: {
        t = NewArrayType(self.GetType(typed.GetBaseType()), typed.GetLength(), self.ptrSize)
        self.PutType(typed, t)
      }
      case *FunctionTypeRef: {
        params := typed.GetParams()
        paramRefs := params.GetParamDescs()
        paramTypes := make([]core.IType, len(paramRefs))
        for i := range paramRefs {
          paramTypes[i] = self.GetParamType(paramRefs[i])
        }
        t = NewFunctionType(
          self.GetType(typed.GetReturnType()),
          NewParamTypes(typed.GetLocation(), paramTypes, params.IsVararg()),
        )
        self.PutType(typed, t)
      }
      default: {
        panic(fmt.Errorf("unregistered type: %s", ref.String()))
      }
    }
  }
  return t
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
  xs := make([]string, len(self.table))
  for key, _ := range self.table {
    xs = append(xs, fmt.Sprintf("%s", key))
  }
  return fmt.Sprintf("(%s)", strings.Join(xs, "\n"))
}

func (self TypeTable) IsDefined(ref core.ITypeRef) bool {
  _, ok := self.table[ref.String()]
  return ok
}

// array is really a pointer on parameters.
func (self TypeTable) GetParamType(ref core.ITypeRef) core.IType {
  t := self.GetType(ref)
  if t == nil {
    panic(fmt.Errorf("unknown parameter type: %s", ref))
  }
  if t.IsArray() {
    return NewPointerType(self.ptrSize, t.(*ArrayType).GetBaseType())
  } else {
    return t
  }
}

func (self TypeTable) NumTypes() int {
  return len(self.table)
}

func (self TypeTable) GetTypes() []core.IType {
  ts := []core.IType { }
  for _, t := range self.table {
    ts = append(ts, t)
  }
  return ts
}

func (self *TypeTable) SemanticCheck(errorHandler *core.ErrorHandler) {
  ts := self.GetTypes()
  for i := range ts {
    t := ts[i]
    if t.IsCompositeType() {
      ct, ok := t.(core.ICompositeType)
      if ! ok {
        errorHandler.Panicln("not a composite type")
      }
      self.checkCompositeVoidMembers(ct, errorHandler)
      self.checkDuplicatedMembers(ct, errorHandler)
    } else {
      if t.IsArray() {
        at, ok := t.(*ArrayType)
        if ! ok {
          errorHandler.Panicln("not an array type")
        }
        self.checkArrayVoidMembers(at, errorHandler)
      }
    }
    self.checkRecursiveDefinition(t, errorHandler)
  }
}

func (self TypeTable) checkCompositeVoidMembers(t core.ICompositeType, errorHandler *core.ErrorHandler) {
  members := t.GetMembers()
  for i := range members {
    slot := members[i]
    if slot.GetType().IsVoid() {
      errorHandler.Errorln("struct/union cannot contain void")
    }
  }
}

func (self TypeTable) checkArrayVoidMembers(t *ArrayType, errorHandler *core.ErrorHandler) {
  if t.GetBaseType().IsVoid() {
    errorHandler.Errorln("array cannot contain void")
  }
}

func (self TypeTable) checkDuplicatedMembers(t core.ICompositeType, errorHandler *core.ErrorHandler) {
  seen := make(map[string]core.ISlot)
  members := t.GetMembers()
  for i := range members {
    slot := members[i]
    name := slot.GetName()
    _, found := seen[name]
    if found {
      errorHandler.Errorf("%s has duplicated member: %s", t.GetName(), name)
    }
    seen[name] = slot
  }
}

func (self TypeTable) checkRecursiveDefinition(t core.IType, errorHandler *core.ErrorHandler) {
  // not implemented yet
}
