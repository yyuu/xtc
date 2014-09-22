package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type LocalResolver struct {
  errorHandler *core.ErrorHandler
  scopeStack []core.IScope
  constantTable *entity.ConstantTable
}

func NewLocalResolver(errorHandler *core.ErrorHandler) *LocalResolver {
  return &LocalResolver { errorHandler, []core.IScope { }, entity.NewConstantTable() }
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

  toplevel.CheckReferences(self.errorHandler)

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
    function.SetScope(self.popScope().(*entity.LocalScope))
  }
}

func (self *LocalResolver) currentScope() core.IScope {
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

func (self *LocalResolver) popScope() core.IScope {
  if len(self.scopeStack) < 1 {
    panic("stack is empty")
  }
  scope := self.currentScope()
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

var Verbose = 0

func (self *LocalResolver) VisitNode(node core.INode) {
  switch typed := node.(type) {
    case *ast.BlockNode: {
      self.pushScope(typed.GetVariables())
      visitBlockNode(self, typed)
      typed.SetScope(self.popScope().(*entity.LocalScope))
    }
    case *ast.StringLiteralNode: {
      ent := self.constantTable.Intern(typed.GetValue())
      typed.SetEntry(ent)
    }
    case *ast.VariableNode: {
      ent := self.currentScope().GetByName(typed.GetName())
      if ent == nil {
        panic(fmt.Errorf("undefined: %s", typed.GetName()))
      }
      ent.Refered()
      typed.SetEntity(ent)
    }
    default: {
      visitNode(self, node)
    }
  }
}
