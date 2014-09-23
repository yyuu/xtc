package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestFunctionTypeRefKeyString1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewFunctionTypeRef(
    NewIntTypeRef(loc),
    NewParamTypeRefs(loc,
      []core.ITypeRef {
        NewIntTypeRef(loc),
      },
      false,
    ),
  )
  xt.AssertEquals(t, "int f(int x)", x.Key(), "int(int)")
}

func TestFunctionTypeRefKeyString2(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewFunctionTypeRef(
    NewIntTypeRef(loc),
    NewParamTypeRefs(loc,
      []core.ITypeRef {
        NewIntTypeRef(loc),
        NewArrayTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)), -1),
      },
      false,
    ),
  )
  xt.AssertEquals(t, "int main(int argc, char*[] argv)", x.Key(), "int(int,char*[])")
}
