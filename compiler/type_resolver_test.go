package compiler

import (
  "testing"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
  "bitbucket.org/yyuu/bs/xt"
)

func setupTypeResolver(a *bs_ast.AST, table *bs_typesys.TypeTable) (int, *TypeResolver) {
  resolver := NewTypeResolver(bs_core.NewErrorHandler(bs_core.LOG_WARN), bs_core.NewOptions("type_resolver_test.go"), table)
  numTypes := resolver.typeTable.NumTypes()
  resolver.defineTypes(a.ListTypes())
  return numTypes, resolver
}

func TestTypeResolverVisitTypeDefinitionsEmpty(t *testing.T) {
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  numTypes, resolver := setupTypeResolver(a, table)

  bs_ast.VisitTypeDefinitions(resolver, a.ListTypes())
  xt.AssertEquals(t, "empty declaration should not have new types", resolver.typeTable.NumTypes(), numTypes)
}

func TestTypeResolverWithStruct(t *testing.T) {
/*
  struct foo {
    int n
    int m
  };
 */
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(
        bs_ast.NewStructNode(
          loc,
          bs_typesys.NewStructTypeRef(loc, "foo"),
          "foo",
          []bs_core.ISlot {
            bs_ast.NewSlot(bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)), "n"),
            bs_ast.NewSlot(bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)), "m"),
          },
        ),
      ),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  numTypes, resolver := setupTypeResolver(a, table)

  bs_ast.VisitTypeDefinitions(resolver, a.ListTypes())
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
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(
        bs_ast.NewUnionNode(
          loc,
          bs_typesys.NewUnionTypeRef(loc, "foo"),
          "foo",
          []bs_core.ISlot {
            bs_ast.NewSlot(bs_ast.NewTypeNode(loc, bs_typesys.NewShortTypeRef(loc)), "n"),
          },
        ),
        bs_ast.NewUnionNode(
          loc,
          bs_typesys.NewUnionTypeRef(loc, "bar"),
          "bar",
          []bs_core.ISlot {
            bs_ast.NewSlot(bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)), "m"),
          },
        ),
      ),
      bs_ast.NewTypedefNodes(),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  numTypes, resolver := setupTypeResolver(a, table)

  bs_ast.VisitTypeDefinitions(resolver, a.ListTypes())
  xt.AssertEquals(t, "new union `foo' and `bar' should be declared", resolver.typeTable.NumTypes(), numTypes+2)
}

func TestTypeResolverWithTypedef(t *testing.T) {
/*
  typedef short foo;
  typedef int bar;
  typedef long baz;
 */
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(
        bs_ast.NewTypedefNode(loc, bs_typesys.NewShortTypeRef(loc), "foo"),
        bs_ast.NewTypedefNode(loc, bs_typesys.NewIntTypeRef(loc), "bar"),
        bs_ast.NewTypedefNode(loc, bs_typesys.NewLongTypeRef(loc), "baz"),
      ),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  numTypes, resolver := setupTypeResolver(a, table)

  bs_ast.VisitTypeDefinitions(resolver, a.ListTypes())
  xt.AssertEquals(t, "new typedef `foo', `bar' and `baz' should be declared", resolver.typeTable.NumTypes(), numTypes+3)
}

func TestTypeResolverVisitEntity(t *testing.T) {
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  _, resolver := setupTypeResolver(a, table)

  bs_entity.VisitEntities(resolver, a.ListEntities())
  assertTypeResolved(t, "visit entity", a)
}

func TestTypeResolverWithFunctionWithoutArguments(t *testing.T) {
/*
  void hello() {
    println("hello, world");
  }
 */
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc, bs_typesys.NewVoidTypeRef(loc)),
          "hello",
          bs_entity.NewParams(loc,
            bs_entity.NewParameters(),
            false,
          ),
          bs_ast.NewBlockNode(loc,
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewExprStmtNode(loc,
                bs_ast.NewFuncallNode(loc,
                  bs_ast.NewVariableNode(loc, "println"),
                  []bs_core.IExprNode {
                    bs_ast.NewStringLiteralNode(loc, "hello, world"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  _, resolver := setupTypeResolver(a, table)

  bs_entity.VisitEntities(resolver, a.ListEntities())
  assertTypeResolved(t, "function w/o args", a)
}

func TestTypeResolverWithFunctionWithArguments(t *testing.T) {
/*
  int main(int argc, char*[] argv) {
    println("hello, world");
  }
 */
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)),
          "main",
          bs_entity.NewParams(loc,
            bs_entity.NewParameters(
              bs_entity.NewParameter(bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)), "argc"),
              bs_entity.NewParameter(bs_ast.NewTypeNode(loc, bs_typesys.NewArrayTypeRef(bs_typesys.NewPointerTypeRef(bs_typesys.NewCharTypeRef(loc)), 0)), "argv"),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc,
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewExprStmtNode(loc,
                bs_ast.NewFuncallNode(loc,
                  bs_ast.NewVariableNode(loc, "println"),
                  []bs_core.IExprNode {
                    bs_ast.NewStringLiteralNode(loc, "hello, world"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  _, resolver := setupTypeResolver(a, table)

  bs_entity.VisitEntities(resolver, a.ListEntities())
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
  loc := bs_core.NewLocation("", 0, 0)
  a := bs_ast.NewAST(loc,
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)),
          "f",
          bs_entity.NewParams(loc,
            bs_entity.NewParameters(
              bs_entity.NewParameter(bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)), "x"),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc,
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewReturnNode(loc, bs_ast.NewVariableNode(loc, "x")),
            },
          ),
        ),
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)),
          "g",
          bs_entity.NewParams(loc,
            bs_entity.NewParameters(
              bs_entity.NewParameter(bs_ast.NewTypeNode(loc, bs_typesys.NewIntTypeRef(loc)), "x"),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc,
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewExprStmtNode(loc,
                bs_ast.NewFuncallNode(loc,
                  bs_ast.NewVariableNode(loc, "f"),
                  []bs_core.IExprNode {
                    bs_ast.NewBinaryOpNode(loc,
                      "*",
                      bs_ast.NewVariableNode(loc, "x"),
                      bs_ast.NewIntegerLiteralNode(loc, "2"),
                    ),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
    ),
  )
  table := bs_typesys.NewTypeTableFor(bs_core.PLATFORM_LINUX_X86)
  errorHandler := bs_core.NewErrorHandler(bs_core.LOG_WARN)
  options := bs_core.NewOptions("type_resolver_test.go")

  localResolver := NewLocalResolver(bs_core.NewErrorHandler(bs_core.LOG_WARN), options)
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
