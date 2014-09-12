package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestTypeTableGetDefaultTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  xt.AssertEquals(t, "char is 1 byte", table.GetType(NewCharTypeRef(loc)).Size(), 1)
  xt.AssertEquals(t, "short is 2 bytes", table.GetType(NewShortTypeRef(loc)).Size(), 2)
  xt.AssertEquals(t, "int is 4 bytes", table.GetType(NewIntTypeRef(loc)).Size(), 4)
  xt.AssertEquals(t, "long is 4 bytes", table.GetType(NewLongTypeRef(loc)).Size(), 4)
}

func TestTypeTableGetPointerTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  ref1 := NewPointerTypeRef(NewCharTypeRef(loc))
  xt.AssertNotNil(t, "char* != nil", table.GetType(ref1))

  ref2 := NewPointerTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)))
  xt.AssertNotNil(t, "char** != nil", table.GetType(ref2))
}

func TestTypeTableGetArrayTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  ref1 := NewArrayTypeRef(NewIntTypeRef(loc), 255)
  xt.AssertNotNil(t, "int[255] != nil", table.GetType(ref1))

  ref2 := NewArrayTypeRef(NewArrayTypeRef(NewIntTypeRef(loc), 255), 255)
  xt.AssertNotNil(t, "int[255][255] != nil", table.GetType(ref2))
}

func TestTypeTableGetFunctionTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  ref1 := NewFunctionTypeRef(
    NewIntTypeRef(loc),
    NewParamTypeRefs(loc, []core.ITypeRef { NewIntTypeRef(loc) }, false),
  )
  xt.AssertNotNil(t, "int f(int x) != nil", table.GetType(ref1))
}

func TestTypeTableGetMixedTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  ref1 := NewFunctionTypeRef(
    NewIntTypeRef(loc),
    NewParamTypeRefs(loc,
      []core.ITypeRef {
        NewIntTypeRef(loc),
        NewArrayTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)), -1),
      },
      false,
    ),
  )
  xt.AssertNotNil(t, "int main(int argc, char*[] argv) != nil", table.GetType(ref1))
}
