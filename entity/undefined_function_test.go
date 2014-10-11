package entity

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestUndefinedFunctionWithFunctionType(t *testing.T) {
/*
  extern int f();
 */
  loc := xtc_core.NewLocation("", 0, 0)
  f := NewUndefinedFunction(
    xtc_ast.NewTypeNode(loc,
      xtc_typesys.NewFunctionTypeRef(
        xtc_typesys.NewIntTypeRef(loc),
        xtc_typesys.NewParamTypeRefs(loc, []xtc_core.ITypeRef { }, false),
      ),
    ),
    "f",
    NewParams(loc, NewParameters(), false),
  )
  xt.AssertNotNil(t, "f should not be nil", f)
}

func TestUndefinedFunctionWithInvalidType(t *testing.T) {
/*
  extern long g();
 */
  defer func() {
    r := recover()
    if r == "" {
      t.Errorf("functions must be declared as function type")
    }
  }()
  loc := xtc_core.NewLocation("", 0, 0)
  NewUndefinedFunction(
    xtc_ast.NewTypeNode(loc, xtc_typesys.NewLongTypeRef(loc)),
    "g",
    NewParams(loc, NewParameters(), false),
  )
}
