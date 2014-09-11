package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type LocalResolver struct {
  scopeStack []*entity.VariableScope
  constantTable *entity.ConstantTable
}

func NewLocalResolver() *LocalResolver {
  return &LocalResolver { []*entity.VariableScope { }, entity.NewConstantTable() }
}

func (self *LocalResolver) Resolve(a *ast.AST) {
  toplevel := entity.NewToplevelScope()
  self.scopeStack = append(self.scopeStack, toplevel)

  declarations := a.ListDeclarations()
  for i := range declarations {
    toplevel.DeclareEntity(declarations[i])
  }
  definitions := a.ListDefinitions()
  for i := range definitions {
    toplevel.DefineEntity(definitions[i])
  }

  self.resolveGvarInitializers(a)
  self.resolveConstantValues(a)
  self.resolveFunctions(a)

  toplevel.CheckReferences()

  a.SetScope(toplevel)
  a.SetConstantTable(self.constantTable)
}

func (self *LocalResolver) resolveGvarInitializers(a *ast.AST) {
  variables := a.GetDefinedVariables()
  for i := range variables {
    gvar := variables[i]
    if gvar.HasInitializer() {
      ast.VisitNode(self, gvar.GetInitializer())
    }
  }
}

func (self *LocalResolver) resolveConstantValues(a *ast.AST) {
  constants := a.GetConstants()
  for i := range constants {
    constant := constants[i]
    ast.VisitNode(self, constant.GetValue())
  }
}

func (self *LocalResolver) resolveFunctions(a *ast.AST) {
  functions := a.GetDefinedFunctions()
  for i := range functions {
    function := functions[i]
    self.pushScope(function.ListParameters())
    ast.VisitNode(self, function.GetBody())
    function.SetScope(self.popScope())
  }
}

func (self *LocalResolver) currentScope() *entity.VariableScope {
  if len(self.scopeStack) < 1 {
    panic("stack is empty")
  }
  return self.scopeStack[len(self.scopeStack)-1]
}

func (self *LocalResolver) pushScope(vars []*entity.DefinedVariable) {
  scope := entity.NewLocalScope(self.currentScope())
  for i := range vars {
    v := vars[i]
    if scope.IsDefinedLocally(v.GetName()) {
      panic(fmt.Errorf("duplicated variable in scope: %s", v.GetName()))
    }
    scope.DefineVariable(v)
  }
  self.scopeStack = append(self.scopeStack, scope)
}

func (self *LocalResolver) popScope() *entity.VariableScope {
  if len(self.scopeStack) < 1 {
    panic("stack is empty")
  }
  scope := self.scopeStack[len(self.scopeStack)-1]
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

var Verbose = 0

func (self *LocalResolver) VisitNode(node core.INode) {
  switch typed := node.(type) {
    case *ast.BlockNode: {
      self.pushScope(typed.GetVariables())
      typed.SetScope(self.popScope())
    }
    case *ast.StringLiteralNode: {
      e := self.constantTable.Intern(typed.GetValue())
      typed.SetEntry(e)
    }
    case *ast.VariableNode: {
      e := self.currentScope().GetByName(typed.GetName())
      if e == nil {
        panic(fmt.Errorf("undefined: %s", typed.GetName()))
      }
      variable, ok := e.(*entity.DefinedVariable)
      if ! ok {
        panic(fmt.Errorf("not a variable: %s", typed.GetName()))
      }
      variable.Refered()
      typed.SetEntity(variable)
    }
  }
}
