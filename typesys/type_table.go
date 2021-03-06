package typesys

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/xtc/core"
)

type TypeTable struct {
  charSize int
  shortSize int
  intSize int
  longSize int
  pointerSize int
  typeTable map[string]core.IType
  refTable map[string]core.ITypeRef
}

func NewTypeTable(charSize, shortSize, intSize, longSize, ptrSize int) *TypeTable {
  loc := core.NewLocation("[builtin:typesys]", 0, 0)
  tt := TypeTable { charSize, shortSize, intSize, longSize, ptrSize, make(map[string]core.IType), make(map[string]core.ITypeRef) }
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

func NewTypeTableFor(platformId int) *TypeTable {
  switch platformId {
    case core.PLATFORM_X86_LINUX: {
      return NewTypeTableILP32()
    }
    default: {
      panic(fmt.Errorf("unknown platformId: %d", platformId))
    }
  }
}

func (self *TypeTable) PutType(ref core.ITypeRef, t core.IType) {
  self.typeTable[ref.Key()] = t
  self.refTable[ref.Key()] = ref
}

func (self TypeTable) GetType(ref core.ITypeRef) core.IType {
  t := self.typeTable[ref.Key()]
  if t == nil {
    switch typed := ref.(type) {
      case *UserTypeRef: {
        panic(fmt.Errorf("undefined type: %s", typed.GetName()))
      }
      case *PointerTypeRef: {
        t = NewPointerType(self.pointerSize, self.GetType(typed.GetBaseType()))
        self.PutType(typed, t)
      }
      case *ArrayTypeRef: {
        t = NewArrayType(self.GetType(typed.GetBaseType()), typed.GetLength(), self.pointerSize)
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
        panic(fmt.Errorf("unregistered type: %s", ref))
      }
    }
  }
  return t
}

func (self TypeTable) GetTypeRef(target core.IType) core.ITypeRef {
  for key, t := range self.typeTable {
    if t == target {
      return self.refTable[key]
    }
  }
  return nil
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
  return self.pointerSize
}

func (self TypeTable) IsTypeTable() bool {
  return true
}

func (self TypeTable) String() string {
  xs := make([]string, len(self.typeTable))
  for key, _ := range self.typeTable {
    xs = append(xs, fmt.Sprintf("%s", key))
  }
  return fmt.Sprintf("(%s)", strings.Join(xs, "\n"))
}

func (self TypeTable) IsDefined(ref core.ITypeRef) bool {
  _, ok := self.typeTable[ref.Key()]
  return ok
}

// array is really a pointer on parameters.
func (self TypeTable) GetParamType(ref core.ITypeRef) core.IType {
  t := self.GetType(ref)
  if t == nil {
    panic(fmt.Errorf("unknown parameter type: %s", ref))
  }
  if t.IsArray() {
    return NewPointerType(self.pointerSize, t.(*ArrayType).GetBaseType())
  } else {
    return t
  }
}

func (self TypeTable) NumTypes() int {
  return len(self.typeTable)
}

func (self TypeTable) GetTypes() []core.IType {
  ts := []core.IType { }
  for _, t := range self.typeTable {
    ts = append(ts, t)
  }
  return ts
}

func (self *TypeTable) SemanticCheck(errorHandler *core.ErrorHandler) {
  ts := self.GetTypes()
  for i := range ts {
    t := ts[i]
    if t.IsCompositeType() {
      self.checkCompositeVoidMembers(t.(core.ICompositeType), errorHandler)
      self.checkDuplicatedMembers(t.(core.ICompositeType), errorHandler)
    } else {
      if t.IsArray() {
        self.checkArrayVoidMembers(t.(*ArrayType), errorHandler)
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
      errorHandler.Fatal("struct/union cannot contain void")
    }
  }
}

func (self TypeTable) checkArrayVoidMembers(t *ArrayType, errorHandler *core.ErrorHandler) {
  if t.GetBaseType().IsVoid() {
    errorHandler.Fatal("array cannot contain void")
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
      errorHandler.Fatalf("%s has duplicated member: %s", t.GetName(), name)
    }
    seen[name] = slot
  }
}

func (self TypeTable) checkRecursiveDefinition(t core.IType, h *core.ErrorHandler) {
  self._checkRecursiveDefinition(t, make(map[core.IType]int), h)
}

const (
  checking = 1 << iota
  checked
)

func (self TypeTable) _checkRecursiveDefinition(t core.IType, marks map[core.IType]int, h *core.ErrorHandler) {
  switch {
    case marks[t] == checking: {
      h.Fatalf("recursive type definition: %s", t)
      return
    }
    case marks[t] == checked: {
      return
    }
    default: {
      marks[t] = checking
      ct, ok := t.(core.ICompositeType)
      if ok {
        members := ct.GetMembers()
        for i := range members {
          self._checkRecursiveDefinition(members[i].GetType(), marks, h)
        }
      } else {
        switch xt := t.(type) {
          case *ArrayType: self._checkRecursiveDefinition(xt.GetBaseType(), marks, h)
          case *UserType:  self._checkRecursiveDefinition(xt.GetRealType(), marks, h)
        }
      }
      marks[t] = checked
    }
  }
}

func (self TypeTable) VoidType() *VoidType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewVoidTypeRef(loc)
  return self.GetType(ref).(*VoidType)
}

func (self TypeTable) SignedChar() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewCharTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) SignedShort() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewShortTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) SignedInt() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewIntTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) SignedLong() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewLongTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) UnsignedChar() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewUnsignedCharTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) UnsignedShort() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewUnsignedShortTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) UnsignedInt() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewUnsignedIntTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) UnsignedLong() *IntegerType {
  loc := core.NewLocation("[typesys:builtin]", 0, 0)
  ref := NewUnsignedLongTypeRef(loc)
  return self.GetType(ref).(*IntegerType)
}

func (self TypeTable) PointerTo(baseType core.IType) *PointerType {
  return NewPointerType(self.pointerSize, baseType)
}

func (self TypeTable) PtrDiffType() core.IType {
  return self.GetType(self.PtrDiffTypeRef())
}

func (self TypeTable) PtrDiffTypeRef() core.ITypeRef {
  loc := core.NewLocation("[builtin:typesys]", 0, 0)
  return NewIntegerTypeRef(loc, self.PtrDiffTypeName())
}

func (self TypeTable) PtrDiffTypeName() string {
  if self.SignedLong().Size() == self.pointerSize {
    return "long"
  } else {
    if self.SignedInt().Size() == self.pointerSize {
      return "int"
    } else {
      if self.SignedShort().Size() == self.pointerSize {
        return "short"
      } else {
        panic("must not happen: interger.size != pointer.size")
      }
    }
  }
}

func (self TypeTable) SignedStackType() *IntegerType {
  return self.SignedLong()
}

func (self TypeTable) UnsignedStackType() *IntegerType {
  return self.UnsignedLong()
}
