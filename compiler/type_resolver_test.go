package compiler

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
//"bitbucket.org/yyuu/bs/xt"
)

func setupTypeResolver(a *ast.AST, table *typesys.TypeTable) *TypeResolver {
  resolver := NewTypeResolver(core.NewErrorHandler(core.LOG_DEBUG), table)
  resolver.defineTypes(a.ListTypes())
  return resolver
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
      ast.NewStructNodes(
        // TODO
      ),
      ast.NewUnionNodes(
        // TODO
      ),
      ast.NewTypedefNodes(
        // TODO
      ),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  resolver := setupTypeResolver(a, table)

  types := a.ListTypes()
  for i := range types {
    ast.VisitNode(resolver, types[i])
  }
  // TODO: write tests
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
  resolver := setupTypeResolver(a, table)

  types := a.ListTypes()
  for i := range types {
    ast.VisitNode(resolver, types[i])
  }
  // TODO: add asserts
  t.SkipNow()
}

func TestTypeResolverWithUnion(t *testing.T) {
/*
  union foo {
    short n
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
            ast.NewSlot(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "n"),
            ast.NewSlot(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "m"),
          },
        ),
      ),
      ast.NewTypedefNodes(),
    ),
  )
  table := typesys.NewTypeTableFor("x86-linux")
  resolver := setupTypeResolver(a, table)

  types := a.ListTypes()
  for i := range types {
    ast.VisitNode(resolver, types[i])
  }
  // FIXME: add asserts
  t.SkipNow()
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
  resolver := setupTypeResolver(a, table)

  entities := a.ListEntities()
  for i := range entities {
    entity.VisitEntity(resolver, entities[i])
  }
  // TODO: add asserts
  t.SkipNow()
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
  resolver := setupTypeResolver(a, table)

  entities := a.ListEntities()
  for i := range entities {
    entity.VisitEntity(resolver, entities[i])
  }
  // TODO: add asserts
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
  resolver := setupTypeResolver(a, table)

  entities := a.ListEntities()
  for i := range entities {
    entity.VisitEntity(resolver, entities[i])
  }
  // TODO: add asserts
}
