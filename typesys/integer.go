package typesys

import (
  "math"
  "bitbucket.org/yyuu/bs/core"
)

// IntegerType
type IntegerType struct {
  ClassName string
  IntegerSize int
  Signed bool
  Name string
}

func NewIntegerType(size int, isSigned bool, name string) *IntegerType {
  return &IntegerType { "typesys.IntegerType", size, isSigned, name }
}

func NewCharType(size int) *IntegerType {
  return NewIntegerType(size, true, "char")
}

func NewShortType(size int) *IntegerType {
  return NewIntegerType(size, true, "short")
}

func NewIntType(size int) *IntegerType {
  return NewIntegerType(size, true, "int")
}

func NewLongType(size int) *IntegerType {
  return NewIntegerType(size, true, "long")
}

func NewUnsignedCharType(size int) *IntegerType {
  return NewIntegerType(size, false, "unsigned char")
}

func NewUnsignedShortType(size int) *IntegerType {
  return NewIntegerType(size, false, "unsigned short")
}

func NewUnsignedIntType(size int) *IntegerType {
  return NewIntegerType(size, false, "unsigned int")
}

func NewUnsignedLongType(size int) *IntegerType {
  return NewIntegerType(size, false, "unsigned long")
}

func (self IntegerType) Key() string {
  return self.Name
}

func (self IntegerType) String() string {
  return self.Key()
}

func (self IntegerType) Size() int {
  return self.IntegerSize
}

func (self IntegerType) AllocSize() int {
  return self.Size()
}

func (self IntegerType) Alignment() int {
  return self.AllocSize()
}

func (self IntegerType) IsVoid() bool {
  return false
}

func (self IntegerType) IsInteger() bool {
  return true
}

func (self IntegerType) IsSigned() bool {
  return self.Signed
}

func (self IntegerType) IsPointer() bool {
  return false
}

func (self IntegerType) IsArray() bool {
  return false
}

func (self IntegerType) IsCompositeType() bool {
  return false
}

func (self IntegerType) IsStruct() bool {
  return false
}

func (self IntegerType) IsUnion() bool {
  return false
}

func (self IntegerType) IsUserType() bool {
  return false
}

func (self IntegerType) IsFunction() bool {
  return false
}

func (self IntegerType) IsCallable() bool {
  return false
}

func (self IntegerType) GetName() string {
  return self.Name
}

func (self IntegerType) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}

func (self IntegerType) MinValue() int64 {
  if self.Signed {
    return 0 - int64(math.Pow(2, float64(self.IntegerSize*8-1)))
  } else {
    return 0
  }
}

func (self IntegerType) MaxValue() int64 {
  if self.Signed {
    return int64(math.Pow(2, float64(self.IntegerSize*8-1))) - 1
  } else {
    return int64(math.Pow(2, float64(self.IntegerSize*8))) - 1
  }
}
