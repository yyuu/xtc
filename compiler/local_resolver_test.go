package compiler

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func setupLocalResolver(a *xtc_ast.AST) *LocalResolver {
  resolver := NewLocalResolver(xtc_core.NewErrorHandler(xtc_core.LOG_WARN), xtc_core.NewOptions("local_resolver_test.go"))
  toplevel := xtc_entity.NewToplevelScope()
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
  resolver := NewLocalResolver(xtc_core.NewErrorHandler(xtc_core.LOG_WARN), xtc_core.NewOptions("local_resolver_test.go"))
  loc := xtc_core.NewLocation("", 0, 0)
  toplevel := xtc_entity.NewToplevelScope()

  resolver.scopeStack = append(resolver.scopeStack, toplevel)
  xt.AssertEquals(t, "let there be toplevel", len(resolver.scopeStack), 1)
  xt.AssertEquals(t, "there is toplevel", resolver.currentScope(), toplevel)

  resolver.pushScope(
    xtc_entity.NewDefinedVariables(
      xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "foo", xtc_ast.NewIntegerLiteralNode(loc, "12345")),
    ),
  )
  scope1 := resolver.currentScope()
  xt.AssertEquals(t, "pushScope should increase the stack", len(resolver.scopeStack), 2)
  xt.AssertEquals(t, "stack should be increased", resolver.currentScope(), scope1)
  xt.AssertEquals(t, "scope1.GetParent should return toplevel", resolver.currentScope().GetParent(), toplevel)
  xt.AssertEquals(t, "scope1.GetToplevel should return toplevel", resolver.currentScope().GetToplevel(), toplevel)
  xt.AssertNotNil(t, "scope1 should contain foo", scope1.GetByName("foo"))

  resolver.pushScope(
    xtc_entity.NewDefinedVariables(
      xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "bar", xtc_ast.NewIntegerLiteralNode(loc, "67890")),
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
  resolver := setupLocalResolver(a)
  resolver.resolveGvarInitializers(a)
  xt.AssertEquals(t, "empty declaration should only have toplevel", len(resolver.scopeStack), 1)
  xt.AssertTrue(t, "empty declaration should have empty constants", resolver.constantTable.IsEmpty())
}

func TestLocalResolverWithGlobalVariables(t *testing.T) {
/*
  static int foo = 12345;
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(
        xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "foo", xtc_ast.NewIntegerLiteralNode(loc, "12345")),
      ),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
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
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(
        xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "foo", xtc_ast.NewVariableNode(loc, "bar")),
      ),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
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
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(
        xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "foo", xtc_ast.NewIntegerLiteralNode(loc, "12345")),
        xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "bar", xtc_ast.NewVariableNode(loc, "foo")),
        xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "baz", xtc_ast.NewVariableNode(loc, "foo")),
      ),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  resolver := setupLocalResolver(a)
  resolver.resolveGvarInitializers(a)

  foo := resolver.currentScope().GetByName("foo")
  foo_var, ok := foo.(*xtc_entity.DefinedVariable)
  xt.AssertNotNil(t, "foo should not be nil", foo)
  xt.AssertTrue(t, "foo should be an *entity.DefinedVariable", ok)
  xt.AssertEquals(t, "foo should be refered 2 times", foo_var.GetNumRefered(), 2)

  bar := resolver.currentScope().GetByName("bar")
  bar_var, ok := bar.(*xtc_entity.DefinedVariable)
  xt.AssertNotNil(t, "bar should not be nil", bar)
  xt.AssertTrue(t, "bar should be an *entity.DefinedVariable", ok)
  xt.AssertEquals(t, "bar should be refered 0 times", bar_var.GetNumRefered(), 0)

  baz := resolver.currentScope().GetByName("baz")
  baz_var, ok := baz.(*xtc_entity.DefinedVariable)
  xt.AssertNotNil(t, "baz should not be nil", baz)
  xt.AssertTrue(t, "baz should be an *entity.DefinedVariable", ok)
  xt.AssertEquals(t, "baz should be refered 0 times", baz_var.GetNumRefered(), 0)
}

func TestLocalResolverWithConstants1(t *testing.T) {
/*
  const char[] foo = "foo";
  const char[] bar = "bar";
 */
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(
        xtc_entity.NewConstant(
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewArrayTypeRef(xtc_typesys.NewCharTypeRef(loc), len("foo"))),
          "foo",
          xtc_ast.NewStringLiteralNode(loc, "foo"),
        ),
        xtc_entity.NewConstant(
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewArrayTypeRef(xtc_typesys.NewCharTypeRef(loc), len("bar"))),
          "bar",
          xtc_ast.NewStringLiteralNode(loc, "bar"),
        ),
      ),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  resolver := setupLocalResolver(a)
  resolver.resolveConstantValues(a)

  constants := resolver.constantTable.GetEntries()
  xt.AssertEquals(t, "there should be 2 constants", len(constants), 2)
  if constants[0].GetValue() != "foo" {
    xt.AssertEquals(t, "rest of constant must be \"foo\"", constants[1].GetValue(), "foo")
  } else {
    xt.AssertEquals(t, "rest of constant must be \"bar\"", constants[1].GetValue(), "bar")
  }

  nodes := a.GetConstants()
  xt.AssertEquals(t, "there should be 2 constant nodes", len(nodes), 2)
  for i := range nodes {
    s, ok := nodes[i].GetValue().(*xtc_ast.StringLiteralNode)
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
  loc := xtc_core.NewLocation("", 0, 0)
  a := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          true,
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
          "foo",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "n"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc,
                xtc_ast.NewBinaryOpNode(loc,
                  "+",
                  xtc_ast.NewIntegerLiteralNode(loc, "12345"),
                  xtc_ast.NewVariableNode(loc, "n"),
                ),
              ),
            },
          ),
        ),
        xtc_entity.NewDefinedFunction(
          true,
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
          "bar",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "m"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc,
                xtc_ast.NewBinaryOpNode(loc,
                  "+",
                  xtc_ast.NewIntegerLiteralNode(loc, "67890"),
                  xtc_ast.NewVariableNode(loc, "m"),
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
  resolver := setupLocalResolver(a)
  resolver.resolveFunctions(a)

  functions := a.GetDefinedFunctions()
  xt.AssertEquals(t, "there should be 2 functions", len(functions), 2)
  for i := range functions {
    function := functions[i]
    xt.AssertNotNil(t, "defined functions should have its own scope", function.GetScope())
  }
}
