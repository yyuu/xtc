package compiler

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func setupTypeResolver(a *xtc_ast.AST, table *xtc_typesys.TypeTable) (int, *TypeResolver) {
  resolver := NewTypeResolver(xtc_core.NewErrorHandler(xtc_core.LOG_WARN), xtc_core.NewOptions("type_resolver_test.go"), table)
  numTypes := resolver.typeTable.NumTypes()
  resolver.defineTypes(a.ListTypes())
  return numTypes, resolver
}

func TestTypeResolverVisitTypeDefinitionsEmpty(t *testing.T) {
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  numTypes, resolver := setupTypeResolver(a, table)

  xtc_ast.VisitTypeDefinitions(resolver, a.ListTypes())
  xt.AssertEquals(t, "empty declaration should not have new types", resolver.typeTable.NumTypes(), numTypes)
}

func TestTypeResolverWithStruct(t *testing.T) {
/*
  struct foo {
    int n
    int m
  };
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(
        xtc_ast.NewStructNode(
          loc,
          xtc_typesys.NewStructTypeRef(loc, "foo"),
          "foo",
          []xtc_core.ISlot {
            xtc_ast.NewSlot(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "n"),
            xtc_ast.NewSlot(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "m"),
          },
        ),
      ),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  numTypes, resolver := setupTypeResolver(a, table)

  xtc_ast.VisitTypeDefinitions(resolver, a.ListTypes())
  xt.AssertEquals(t, "new struct `foo' should be declared", resolver.typeTable.NumTypes(), numTypes+1)
}

func TestTypeResolverWithUnion(t *testing.T) {
/*
  union foo {
    short n
  };
  union bar {
    int m
  };
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(
        xtc_ast.NewUnionNode(
          loc,
          xtc_typesys.NewUnionTypeRef(loc, "foo"),
          "foo",
          []xtc_core.ISlot {
            xtc_ast.NewSlot(xtc_ast.NewTypeNode(loc, xtc_typesys.NewShortTypeRef(loc)), "n"),
          },
        ),
        xtc_ast.NewUnionNode(
          loc,
          xtc_typesys.NewUnionTypeRef(loc, "bar"),
          "bar",
          []xtc_core.ISlot {
            xtc_ast.NewSlot(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "m"),
          },
        ),
      ),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  numTypes, resolver := setupTypeResolver(a, table)

  xtc_ast.VisitTypeDefinitions(resolver, a.ListTypes())
  xt.AssertEquals(t, "new union `foo' and `bar' should be declared", resolver.typeTable.NumTypes(), numTypes+2)
}

func TestTypeResolverWithTypedef(t *testing.T) {
/*
  typedef short foo;
  typedef int bar;
  typedef long baz;
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(
        xtc_ast.NewTypedefNode(loc, xtc_typesys.NewShortTypeRef(loc), "foo"),
        xtc_ast.NewTypedefNode(loc, xtc_typesys.NewIntTypeRef(loc), "bar"),
        xtc_ast.NewTypedefNode(loc, xtc_typesys.NewLongTypeRef(loc), "baz"),
      ),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  numTypes, resolver := setupTypeResolver(a, table)

  xtc_ast.VisitTypeDefinitions(resolver, a.ListTypes())
  xt.AssertEquals(t, "new typedef `foo', `bar' and `baz' should be declared", resolver.typeTable.NumTypes(), numTypes+3)
}

func TestTypeResolverVisitEntity(t *testing.T) {
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  _, resolver := setupTypeResolver(a, table)

  xtc_entity.VisitEntities(resolver, a.ListEntities())
  assertTypeResolved(t, "visit entity", a)
}

func TestTypeResolverWithFunctionWithoutArguments(t *testing.T) {
/*
  void hello() {
    println("hello, world");
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc, []xtc_core.ITypeRef { }, false),
            ),
          ),
          "hello",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewExprStmtNode(loc,
                xtc_ast.NewFuncallNode(loc,
                  xtc_ast.NewVariableNode(loc, "println"),
                  []xtc_core.IExprNode {
                    xtc_ast.NewStringLiteralNode(loc, "hello, world"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  _, resolver := setupTypeResolver(a, table)

  xtc_entity.VisitEntities(resolver, a.ListEntities())
  assertTypeResolved(t, "function w/o args", a)
}

func TestTypeResolverWithFunctionWithArguments(t *testing.T) {
/*
  int main(int argc, char*[] argv) {
    println("hello, world");
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                  xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0),
                },
                false,
              ),
            ),
          ),
          "main",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "argc"),
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0)), "argv"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewExprStmtNode(loc,
                xtc_ast.NewFuncallNode(loc,
                  xtc_ast.NewVariableNode(loc, "println"),
                  []xtc_core.IExprNode {
                    xtc_ast.NewStringLiteralNode(loc, "hello, world"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  _, resolver := setupTypeResolver(a, table)

  xtc_entity.VisitEntities(resolver, a.ListEntities())
  // TODO: add asserts
  assertTypeResolved(t, "function w/ args", a)
}

func TestTypeResolverWithFunctionArguments(t *testing.T) {
/*
  int f(int x) {
    return x;
  }
  int g(int x) {
    return f(x*2);
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "x"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, xtc_ast.NewVariableNode(loc, "x")),
            },
          ),
        ),
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "g",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "x"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewExprStmtNode(loc,
                xtc_ast.NewFuncallNode(loc,
                  xtc_ast.NewVariableNode(loc, "f"),
                  []xtc_core.IExprNode {
                    xtc_ast.NewBinaryOpNode(loc,
                      "*",
                      xtc_ast.NewVariableNode(loc, "x"),
                      xtc_ast.NewIntegerLiteralNode(loc, "2"),
                    ),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  errorHandler := xtc_core.NewErrorHandler(xtc_core.LOG_WARN)
  options := xtc_core.NewOptions("type_resolver_test.go")

  localResolver := NewLocalResolver(xtc_core.NewErrorHandler(xtc_core.LOG_WARN), options)
  localResolver.Resolve(a)
  typeResolver := NewTypeResolver(errorHandler, options, table)
  typeResolver.Resolve(a)

  assertTypeResolved(t, "function args", a)
  defer func() {
    if r := recover(); r != nil {
      t.Error(r)
      t.Error("one (or more) of type has not been resolved")
      t.Error(xt.JSON(a))
      t.Fail()
    }
  }()
}

func TestTypeResolverNestedBlocks(t *testing.T) {
/*
  int main(int argc, char*[] argv) {
    int i;
    char* arg;
    for (i=0; i<argc; i++) {
      arg = argv[i];
      printf("argv[%d]=%s\n", i, arg);
    }
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                  xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0),
                },
                false,
              ),
            ),
          ),
          "main",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "argc"),
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0)), "argv"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(
              xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "i", nil),
              xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc))), "arg", nil),
            ),
            []xtc_core.IStmtNode {
              xtc_ast.NewForNode(loc,
                xtc_ast.NewAssignNode(loc, xtc_ast.NewVariableNode(loc, "i"), xtc_ast.NewIntegerLiteralNode(loc, "0")),
                xtc_ast.NewBinaryOpNode(loc, "<", xtc_ast.NewVariableNode(loc, "i"), xtc_ast.NewVariableNode(loc, "n")),
                xtc_ast.NewSuffixOpNode(loc, "++", xtc_ast.NewVariableNode(loc, "i")),
                xtc_ast.NewBlockNode(loc,
                  xtc_entity.NewDefinedVariables(),
                  []xtc_core.IStmtNode {
                    xtc_ast.NewExprStmtNode(loc,
                      xtc_ast.NewAssignNode(loc,
                        xtc_ast.NewVariableNode(loc, "arg"),
                        xtc_ast.NewArefNode(loc, xtc_ast.NewVariableNode(loc, "argv"), xtc_ast.NewVariableNode(loc, "i")),
                      ),
                    ),
                    xtc_ast.NewExprStmtNode(loc,
                      xtc_ast.NewFuncallNode(loc,
                        xtc_ast.NewVariableNode(loc, "printf"),
                        []xtc_core.IExprNode {
                          xtc_ast.NewVariableNode(loc, "argv[%d]=%s\n"),
                          xtc_ast.NewVariableNode(loc, "i"),
                          xtc_ast.NewVariableNode(loc, "arg"),
                        },
                      ),
                    ),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  _, resolver := setupTypeResolver(a, table)

  xtc_entity.VisitEntities(resolver, a.ListEntities())
  assertTypeResolved(t, "nested blocks", a)
}
