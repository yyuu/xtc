package compiler

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/xt"
)

func TestLocalResolverWithEmptyDeclarations(t *testing.T) {
  resolver := NewLocalResolver()
  a := ast.NewAST(core.NewLocation("", 0, 0),
    ast.NewDeclarations(
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
  resolver.resolveGvarInitializers(a)
  xt.AssertEquals(t, "empty declarations should have empty stack", len(resolver.scopeStack), 0)
  xt.AssertTrue(t, "empty declarations should have empty constants", resolver.constantTable.IsEmpty())
}
