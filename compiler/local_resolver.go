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
  xs := a.GetDefinedVariables()
  ys := make([]*entity.DefinedVariable, len(xs))
  for i := range xs {
    gvar := xs[i]
    if gvar.HasInitializer() {
      ast.VisitNode(self, gvar.GetInitializer())
    }
    ys[i] = gvar
  }
  a.SetDefinedVariables(ys)
}

func (self *LocalResolver) resolveConstantValues(a *ast.AST) {
  xs := a.GetConstants()
  ys := make([]*entity.Constant, len(xs))
  for i := range xs {
    constant := xs[i]
    ast.VisitNode(self, constant.GetValue())
    ys[i] = constant
  }
  a.SetConstants(ys)
}

func (self *LocalResolver) resolveFunctions(a *ast.AST) {
  xs := a.GetDefinedFunctions()
  ys := make([]*entity.DefinedFunction, len(xs))
  for i := range xs {
    function := xs[i]
    self.pushScope(function.ListParameters())
    ast.VisitNode(self, function.GetBody())
    function.SetScope(self.popScope())
    ys[i] = function
  }
  a.SetDefinedFunctions(ys)
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
