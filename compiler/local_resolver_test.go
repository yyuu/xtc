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
  resolver := NewLocalResolver(core.NewErrorHandler(core.LOG_WARN), core.NewOptions("local_resolver_test.go"))
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
  resolver := NewLocalResolver(core.NewErrorHandler(core.LOG_WARN), core.NewOptions("local_resolver_test.go"))
  loc := core.NewLocation("", 0, 0)
  toplevel := entity.NewToplevelScope()

  resolver.scopeStack = append(resolver.scopeStack, toplevel)
  xt.AssertEquals(t, "let there be toplevel", len(resolver.scopeStack), 1)
  xt.AssertEquals(t, "there is toplevel", resolver.currentScope(), toplevel)

  resolver.pushScope(
    entity.NewDefinedVariables(
      entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "foo", ast.NewIntegerLiteralNode(loc, "12345")),
    ),
  )
  scope1 := resolver.currentScope()
  xt.AssertEquals(t, "pushScope should increase the stack", len(resolver.scopeStack), 2)
  xt.AssertEquals(t, "stack should be increased", resolver.currentScope(), scope1)
  xt.AssertEquals(t, "scope1.GetParent should return toplevel", resolver.currentScope().GetParent(), toplevel)
  xt.AssertEquals(t, "scope1.GetToplevel should return toplevel", resolver.currentScope().GetToplevel(), toplevel)
  xt.AssertNotNil(t, "scope1 should contain foo", scope1.GetByName("foo"))

  resolver.pushScope(
    entity.NewDefinedVariables(
      entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "bar", ast.NewIntegerLiteralNode(loc, "67890")),
    ),
  )
  scope2 := resolver.currentScope()
  xt.AssertEquals(t, "pushScope should increase the stack", len(resolver.scopeStack), 3)
  xt.AssertEquals(t, "stack should be increased", resolver.currentScope(), scope2)
  xt.AssertEquals(t, "scope2.GetParent should return scope1", resolver.currentScope().GetParent(), scope1)
  xt.AssertEquals(t, "scope2.GetToplevel should return toplevel", resolver.currentScope().GetToplevel(), toplevel)
  xt.AssertNotNil(t, "scope2 should contain foo", scope2.GetByName("foo"))
  xt.AssertNotNil(t, "scope2 should contain bar", scope2.GetByName("bar"))

  xt.AssertEquals(t, "popScope should decrease the stack", resolver.popScope(), scope2)
  xt.AssertEquals(t, "stack should be decreased", len(resolver.scopeStack), 2)
  xt.AssertNotNil(t, "scope1 should contain foo", scope1.GetByName("foo"))

  xt.AssertEquals(t, "popScope should decrease the stack", resolver.popScope(), scope1)
  xt.AssertEquals(t, "stack should be decreased", len(resolver.scopeStack), 1)
}

func TestLocalResolverWithEmptyDeclaration(t *testing.T) {
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
  resolver := setupLocalResolver(a)
  resolver.resolveGvarInitializers(a)
  xt.AssertEquals(t, "empty declaration should only have toplevel", len(resolver.scopeStack), 1)
  xt.AssertTrue(t, "empty declaration should have empty constants", resolver.constantTable.IsEmpty())
}

func TestLocalResolverWithGlobalVariables(t *testing.T) {
/*
  static int foo = 12345;
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "foo", ast.NewIntegerLiteralNode(loc, "12345")),
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
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "foo", ast.NewVariableNode(loc, "bar")),
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
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "foo", ast.NewIntegerLiteralNode(loc, "12345")),
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "bar", ast.NewVariableNode(loc, "foo")),
        entity.NewDefinedVariable(false, ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "baz", ast.NewVariableNode(loc, "foo")),
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

func TestLocalResolverWithConstants1(t *testing.T) {
/*
  const char[] foo = "foo";
  const char[] bar = "bar";
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(
        entity.NewConstant(
          ast.NewTypeNode(loc, typesys.NewArrayTypeRef(typesys.NewCharTypeRef(loc), len("foo"))),
          "foo",
          ast.NewStringLiteralNode(loc, "\"foo\""),
        ),
        entity.NewConstant(
          ast.NewTypeNode(loc, typesys.NewArrayTypeRef(typesys.NewCharTypeRef(loc), len("bar"))),
          "bar",
          ast.NewStringLiteralNode(loc, "\"bar\""),
        ),
      ),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  resolver := setupLocalResolver(a)
  resolver.resolveConstantValues(a)

  constants := resolver.constantTable.GetEntries()
  xt.AssertEquals(t, "there should be 2 constants", len(constants), 2)
  if constants[0].GetValue() != "\"foo\"" {
    xt.AssertEquals(t, "rest of constant must be \"foo\"", constants[1].GetValue(), "\"foo\"")
  } else {
    xt.AssertEquals(t, "rest of constant must be \"bar\"", constants[1].GetValue(), "\"bar\"")
  }

  nodes := a.GetConstants()
  xt.AssertEquals(t, "there should be 2 constant nodes", len(nodes), 2)
  for i := range nodes {
    s, ok := nodes[i].GetValue().(*ast.StringLiteralNode)
    if ! ok {
      t.Errorf("there should be only string constants: %v", nodes[i])
    }
    xt.AssertNotNil(t, "string literal should have its constant entry", s.GetEntry())
  }
}

func TestLocalResolverWithFunctions1(t *testing.T) {
/*
   int foo(int n) {
     return 12345 + n;
   }

   int bar(int m) {
      return 67890 + m;
   }
 */
  loc := core.NewLocation("", 0, 0)
  a := ast.NewAST(loc,
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)),
          "foo",
          entity.NewParams(loc,
            entity.NewParameters(
              entity.NewParameter(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "n"),
            ),
            false,
          ),
          ast.NewBlockNode(loc,
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewReturnNode(loc,
                ast.NewBinaryOpNode(loc,
                  "+",
                  ast.NewIntegerLiteralNode(loc, "12345"),
                  ast.NewVariableNode(loc, "n"),
                ),
              ),
            },
          ),
        ),
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)),
          "bar",
          entity.NewParams(loc,
            entity.NewParameters(
              entity.NewParameter(ast.NewTypeNode(loc, typesys.NewIntTypeRef(loc)), "m"),
            ),
            false,
          ),
          ast.NewBlockNode(loc,
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewReturnNode(loc,
                ast.NewBinaryOpNode(loc,
                  "+",
                  ast.NewIntegerLiteralNode(loc, "67890"),
                  ast.NewVariableNode(loc, "m"),
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
  resolver := setupLocalResolver(a)
  resolver.resolveFunctions(a)

  functions := a.GetDefinedFunctions()
  xt.AssertEquals(t, "there should be 2 functions", len(functions), 2)
  for i := range functions {
    function := functions[i]
    xt.AssertNotNil(t, "defined functions should have its own scope", function.GetScope())
  }
}
