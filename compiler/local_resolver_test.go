package compiler

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
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

func TestLocalResolverStacking(t *testing.T) {
  resolver := NewLocalResolver()
  loc := core.NewLocation("", 0, 0)
  toplevel := entity.NewToplevelScope()

  resolver.scopeStack = append(resolver.scopeStack, toplevel)
  xt.AssertEquals(t, "let there be toplevel", len(resolver.scopeStack), 1)
  xt.AssertEquals(t, "there is toplevel", resolver.currentScope(), toplevel)

  resolver.pushScope(
    entity.NewDefinedVariables(
      entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntegerTypeRef(loc, "int")), "foo", ast.NewIntegerLiteralNode(loc, "12345")),
    ),
  )
  scope1 := resolver.currentScope()
  xt.AssertEquals(t, "pushScope should increase the stack", len(resolver.scopeStack), 2)
  xt.AssertEquals(t, "stack should be increased", resolver.currentScope(), scope1)
  xt.AssertEquals(t, "scope1.GetParent should return toplevel", resolver.currentScope().GetParent(), toplevel)
  xt.AssertEquals(t, "scope1.GetToplevel should return toplevel", resolver.currentScope().GetToplevel(), toplevel)

  resolver.pushScope(
    entity.NewDefinedVariables(
      entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntegerTypeRef(loc, "int")), "bar", ast.NewIntegerLiteralNode(loc, "67890")),
    ),
  )
  scope2 := resolver.currentScope()
  xt.AssertEquals(t, "pushScope should increase the stack", len(resolver.scopeStack), 3)
  xt.AssertEquals(t, "stack should be increased", resolver.currentScope(), scope2)
  xt.AssertEquals(t, "scope2.GetParent should return scope1", resolver.currentScope().GetParent(), scope1)
  xt.AssertEquals(t, "scope2.GetToplevel should return toplevel", resolver.currentScope().GetToplevel(), toplevel)

  xt.AssertEquals(t, "popScope should decrease the stack", resolver.popScope(), scope2)
  xt.AssertEquals(t, "stack should be decreased", len(resolver.scopeStack), 2)

  xt.AssertEquals(t, "popScope should decrease the stack", resolver.popScope(), scope1)
  xt.AssertEquals(t, "stack should be decreased", len(resolver.scopeStack), 1)
}
