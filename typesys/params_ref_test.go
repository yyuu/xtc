package typesys

import (
  "testing"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestParamTypeRefsKeyString1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewParamTypeRefs(loc,
         []core.ITypeRef { },
         false,
       )
  xt.AssertEquals(t, "(empty params)", x.Key(), "")
}

func TestParamTypeRefsKeyString2(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewParamTypeRefs(loc,
         []core.ITypeRef {
           NewIntTypeRef(loc),
           NewArrayTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)), -1),
         },
         false,
       )
  xt.AssertEquals(t, "int argc, char*[] argv", x.Key(), "int,char*[]")
}
