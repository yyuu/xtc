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
