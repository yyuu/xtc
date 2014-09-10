package compiler

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
  "bitbucket.org/yyuu/bs/xt"
)

func setupLocalResolver(a *ast.AST) *LocalResolver {
  resolver := NewLocalResolver()
  toplevel := entity.NewToplevelScope()
  resolver.scopeStack = append(resolver.scopeStack, toplevel)

  declarations := a.ListDeclarations()
  for i := range declarations {
    toplevel.DeclareEntity(declarations[i])
  }

  definitions := a.ListDefinitions()
  for i := range definitions {
    toplevel.DefineEntity(definitions[i])
  }
  return resolver
}

func TestLocalResolverPushPopStacks(t *testing.T) {
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
  xt.AssertNotNil(t, "scope1 should contain foo", scope1.GetByName("foo"))
  xt.AssertNil(t, "scope1 should not contain bar", scope1.GetByName("bar"))

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
  xt.AssertNil(t, "scope2 should not contain foo", scope2.GetByName("foo"))
  xt.AssertNotNil(t, "scope2 should contain bar", scope2.GetByName("bar"))

  xt.AssertEquals(t, "popScope should decrease the stack", resolver.popScope(), scope2)
  xt.AssertEquals(t, "stack should be decreased", len(resolver.scopeStack), 2)
  xt.AssertNotNil(t, "scope1 should contain foo", scope1.GetByName("foo"))
  xt.AssertNil(t, "scope1 should not contain bar", scope1.GetByName("bar"))

  xt.AssertEquals(t, "popScope should decrease the stack", resolver.popScope(), scope1)
  xt.AssertEquals(t, "stack should be decreased", len(resolver.scopeStack), 1)
}

func TestLocalResolverWithEmptyDeclarations(t *testing.T) {
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
  resolver := setupLocalResolver(a)
  resolver.resolveGvarInitializers(a)
  xt.AssertEquals(t, "empty declarations should only have toplevel", len(resolver.scopeStack), 1)
  xt.AssertTrue(t, "empty declarations should have empty constants", resolver.constantTable.IsEmpty())
}

func TestLocalResolverWithGlobalVariables(t *testing.T) {
/*
  static int foo = 12345;
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(core.NewLocation("", 0, 0),
    ast.NewDeclarations(
      entity.NewDefinedVariables(
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntegerTypeRef(loc, "int")), "foo", ast.NewIntegerLiteralNode(loc, "12345")),
      ),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  resolver := setupLocalResolver(a)
  resolver.resolveGvarInitializers(a)
  xt.AssertNotNil(t, "foo should be resolved", resolver.currentScope().GetByName("foo"))
}

func TestLocalResolverWithGlobalVariables2(t *testing.T) {
/*
  static int foo = bar; // undefined
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(core.NewLocation("", 0, 0),
    ast.NewDeclarations(
      entity.NewDefinedVariables(
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntegerTypeRef(loc, "int")), "foo", ast.NewVariableNode(loc, "bar")),
      ),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  resolver := setupLocalResolver(a)
  defer func() {
    if r := recover(); r != nil {
      return
    }
  }()
  resolver.resolveGvarInitializers(a)
  t.Error("it should fail on reference for undefined variables")
  t.Fail()
}

func TestLocalResolverWithGlobalVariables3(t *testing.T) {
/*
  static int foo = 12345; // reference ==> 2
  static int bar = foo;
  static int baz = foo;
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(core.NewLocation("", 0, 0),
    ast.NewDeclarations(
      entity.NewDefinedVariables(
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntegerTypeRef(loc, "int")), "foo", ast.NewIntegerLiteralNode(loc, "12345")),
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntegerTypeRef(loc, "int")), "bar", ast.NewVariableNode(loc, "foo")),
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntegerTypeRef(loc, "int")), "baz", ast.NewVariableNode(loc, "foo")),
      ),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  resolver := setupLocalResolver(a)
  resolver.resolveGvarInitializers(a)

  foo := resolver.currentScope().GetByName("foo")
  foo_var, ok := foo.(*entity.DefinedVariable)
  xt.AssertNotNil(t, "foo should not be nil", foo)
  xt.AssertTrue(t, "foo should be an *entity.DefinedVariable", ok)
  xt.AssertEquals(t, "foo should be refered 2 times", foo_var.GetNumRefered(), 2)

  bar := resolver.currentScope().GetByName("bar")
  bar_var, ok := bar.(*entity.DefinedVariable)
  xt.AssertNotNil(t, "bar should not be nil", bar)
  xt.AssertTrue(t, "bar should be an *entity.DefinedVariable", ok)
  xt.AssertEquals(t, "bar should be refered 0 times", bar_var.GetNumRefered(), 0)

  baz := resolver.currentScope().GetByName("baz")
  baz_var, ok := baz.(*entity.DefinedVariable)
  xt.AssertNotNil(t, "baz should not be nil", baz)
  xt.AssertTrue(t, "baz should be an *entity.DefinedVariable", ok)
  xt.AssertEquals(t, "baz should be refered 0 times", baz_var.GetNumRefered(), 0)
}
