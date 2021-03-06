package typesys

import (
  "testing"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestTypeTableGetDefaultTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  xt.AssertEquals(t, "char is 1 byte", table.GetType(NewCharTypeRef(loc)).Size(), 1)
  xt.AssertEquals(t, "short is 2 bytes", table.GetType(NewShortTypeRef(loc)).Size(), 2)
  xt.AssertEquals(t, "int is 4 bytes", table.GetType(NewIntTypeRef(loc)).Size(), 4)
  xt.AssertEquals(t, "long is 4 bytes", table.GetType(NewLongTypeRef(loc)).Size(), 4)
}

func TestTypeTableGetTypesMultipleTimes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  xt.AssertEquals(t, "get char multipe times", table.GetType(NewCharTypeRef(loc)), table.GetType(NewCharTypeRef(loc)))
  xt.AssertEquals(t, "get short multipe times", table.GetType(NewShortTypeRef(loc)), table.GetType(NewShortTypeRef(loc)))
  xt.AssertEquals(t, "get int multipe times", table.GetType(NewIntTypeRef(loc)), table.GetType(NewIntTypeRef(loc)))
  xt.AssertEquals(t, "get long multipe times", table.GetType(NewLongTypeRef(loc)), table.GetType(NewLongTypeRef(loc)))
}

func TestTypeTableGetTypesWithDifferentLocations(t *testing.T) {
  loc1 := core.NewLocation("", 1, 2)
  loc2 := core.NewLocation("", 3, 4)
  table := NewTypeTableILP32()
  xt.AssertEquals(t, "get char with different locations", table.GetType(NewCharTypeRef(loc1)), table.GetType(NewCharTypeRef(loc2)))
  xt.AssertEquals(t, "get short with different locations", table.GetType(NewShortTypeRef(loc1)), table.GetType(NewShortTypeRef(loc2)))
  xt.AssertEquals(t, "get int with different locations", table.GetType(NewIntTypeRef(loc1)), table.GetType(NewIntTypeRef(loc2)))
  xt.AssertEquals(t, "get long with different locations", table.GetType(NewLongTypeRef(loc1)), table.GetType(NewLongTypeRef(loc2)))
}

func TestTypeTableGetPointerTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  ref1 := NewPointerTypeRef(NewCharTypeRef(loc))
  xt.AssertNotNil(t, "char* != nil", table.GetType(ref1))
  xt.AssertEquals(t, "char*", table.GetType(ref1).(*PointerType).GetBaseType().(*IntegerType).GetName(), "char")

  ref2 := NewPointerTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)))
  xt.AssertNotNil(t, "char** != nil", table.GetType(ref2))
  xt.AssertEquals(t, "char**", table.GetType(ref2).(*PointerType).GetBaseType().(*PointerType).GetBaseType().(*IntegerType).GetName(), "char")
}

func TestTypeTableGetArrayTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  ref1 := NewArrayTypeRef(NewIntTypeRef(loc), 255)
  xt.AssertNotNil(t, "int[255] != nil", table.GetType(ref1))
  xt.AssertEquals(t, "int[255]", table.GetType(ref1).(*ArrayType).GetBaseType().(*IntegerType).GetName(), "int")

  ref2 := NewArrayTypeRef(NewArrayTypeRef(NewIntTypeRef(loc), 255), 255)
  xt.AssertNotNil(t, "int[255][255] != nil", table.GetType(ref2))
  xt.AssertEquals(t, "int[255][255]", table.GetType(ref2).(*ArrayType).GetBaseType().(*ArrayType).GetBaseType().(*IntegerType).GetName(), "int")
}

func TestTypeTableGetFunctionTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  ref1 := NewFunctionTypeRef(
    NewIntTypeRef(loc),
    NewParamTypeRefs(loc, []core.ITypeRef { NewIntTypeRef(loc) }, false),
  )
  xt.AssertNotNil(t, "int f(int x) != nil", table.GetType(ref1))
  xt.AssertEquals(t, "int f(int x)", table.GetType(ref1).(*FunctionType).GetReturnType().(*IntegerType).GetName(), "int")
  xt.AssertEquals(t, "int f(int x)", len(table.GetType(ref1).(*FunctionType).GetParamTypes().GetParamDescs()), 1)
  xt.AssertEquals(t, "int f(int x)", table.GetType(ref1).(*FunctionType).GetParamTypes().String(), "int")
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
  xt.AssertEquals(t, "int main(int argc, char*[] argv)", table.GetType(ref1).(*FunctionType).GetReturnType().(*IntegerType).GetName(), "int")
  xt.AssertEquals(t, "int main(int argc, char*[] argv)", len(table.GetType(ref1).(*FunctionType).GetParamTypes().GetParamDescs()), 2)
  xt.AssertEquals(t, "int main(int argc, char*[] argv)", table.GetType(ref1).(*FunctionType).GetParamTypes().String(), "int,char**")
}

func TestTypeTableGetUserTypes(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  table := NewTypeTableILP32()
  t1 := NewUserType("foo", newTypeNode(loc, NewIntTypeRef(loc)), loc)
  ref1 := NewUserTypeRef(loc, "foo")
  table.PutType(ref1, t1)
  xt.AssertNotNil(t, "typedef foo int", table.GetType(ref1))
}
