package compiler

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
  "bitbucket.org/yyuu/bs/xt"
)

func setupTypeResolver(a *ast.AST, table *typesys.TypeTable) (int, *TypeResolver) {
  resolver := NewTypeResolver(core.NewErrorHandler(core.LOG_DEBUG), table)
  numTypes := resolver.typeTable.NumTypes()
  resolver.defineTypes(a.ListTypes())
  return numTypes, resolver
}

func TestTypeResolverVisitNodeEmpty(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  numTypes, resolver := setupTypeResolver(a, table)

  types := a.ListTypes()
  for i := range types {
    ast.VisitNode(resolver, types[i])
  }
  xt.AssertEquals(t, "empty declaration should not have new types", resolver.typeTable.NumTypes(), numTypes)
}

func TestTypeResolverWithStruct(t *testing.T) {
/*
  struct foo {
    int n
    int m
  };
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(
        ast.NewStructNode(
          loc,
          typesys.NewStructTypeRef(loc, "foo"),
          "foo",
          []core.ISlot {
            ast.NewSlot(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "n"),
            ast.NewSlot(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "m"),
          },
        ),
      ),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  numTypes, resolver := setupTypeResolver(a, table)

  types := a.ListTypes()
  for i := range types {
    ast.VisitNode(resolver, types[i])
  }
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
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(
        ast.NewUnionNode(
          loc,
          typesys.NewUnionTypeRef(loc, "foo"),
          "foo",
          []core.ISlot {
            ast.NewSlot(ast.NewTypeNode(loc, typesys.NewShortTypeRef(loc)), "n"),
          },
        ),
        ast.NewUnionNode(
          loc,
          typesys.NewUnionTypeRef(loc, "bar"),
          "bar",
          []core.ISlot {
            ast.NewSlot(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "m"),
          },
        ),
      ),
      ast.NewTypedefNodes(),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  numTypes, resolver := setupTypeResolver(a, table)

  types := a.ListTypes()
  for i := range types {
    ast.VisitNode(resolver, types[i])
  }
  xt.AssertEquals(t, "new union `foo' and `bar' should be declared", resolver.typeTable.NumTypes(), numTypes+2)
}

func TestTypeResolverWithTypedef(t *testing.T) {
/*
  typedef short foo;
  typedef int bar;
  typedef long baz;
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(
        ast.NewTypedefNode(loc, typesys.NewShortTypeRef(loc), "foo"),
        ast.NewTypedefNode(loc, typesys.NewIntTypeRef(loc), "bar"),
        ast.NewTypedefNode(loc, typesys.NewLongTypeRef(loc), "baz"),
      ),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  numTypes, resolver := setupTypeResolver(a, table)

  types := a.ListTypes()
  for i := range types {
    ast.VisitNode(resolver, types[i])
  }
  xt.AssertEquals(t, "new typedef `foo', `bar' and `baz' should be declared", resolver.typeTable.NumTypes(), numTypes+3)
}

func TestTypeResolverVisitEntity(t *testing.T) {
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  _, resolver := setupTypeResolver(a, table)

  entities := a.ListEntities()
  for i := range entities {
    entity.VisitEntity(resolver, entities[i])
  }
  assertTypeResolved(t, "visit entity", a)
}

func TestTypeResolverWithFunctionWithoutArguments(t *testing.T) {
/*
  void hello() {
    println("hello, world");
  }
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          false,
          ast.NewTypeNode(loc, typesys.NewVoidTypeRef(loc)),
          "hello",
          entity.NewParams(loc,
            entity.NewParameters(),
            false,
          ),
          ast.NewBlockNode(loc,
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewExprStmtNode(loc,
                ast.NewFuncallNode(loc,
                  ast.NewVariableNode(loc, "println"),
                  []core.IExprNode {
                    ast.NewStringLiteralNode(loc, "\"hello, world\""),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  _, resolver := setupTypeResolver(a, table)

  entities := a.ListEntities()
  for i := range entities {
    entity.VisitEntity(resolver, entities[i])
  }
  assertTypeResolved(t, "function w/o args", a)
}

func TestTypeResolverWithFunctionWithArguments(t *testing.T) {
/*
  int main(int argc, char*[] argv) {
    println("hello, world");
  }
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          false,
          ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)),
          "main",
          entity.NewParams(loc,
            entity.NewParameters(
              entity.NewParameter(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "argc"),
              entity.NewParameter(ast.NewTypeNode(loc, typesys.NewArrayTypeRef(typesys.NewPointerTypeRef(typesys.NewCharTypeRef(loc)), 0)), "argv"),
            ),
            false,
          ),
          ast.NewBlockNode(loc,
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewExprStmtNode(loc,
                ast.NewFuncallNode(loc,
                  ast.NewVariableNode(loc, "println"),
                  []core.IExprNode {
                    ast.NewStringLiteralNode(loc, "\"hello, world\""),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  _, resolver := setupTypeResolver(a, table)

  entities := a.ListEntities()
  for i := range entities {
    entity.VisitEntity(resolver, entities[i])
  }
  // TODO: add asserts
  assertTypeResolved(t, "function w/ args", a)
}
