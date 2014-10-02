package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type LocalResolver struct {
  errorHandler *core.ErrorHandler
  options *core.Options
  scopeStack []core.IScope
  constantTable *entity.ConstantTable
}

func NewLocalResolver(errorHandler *core.ErrorHandler, options *core.Options) *LocalResolver {
  return &LocalResolver { errorHandler, options, []core.IScope { }, entity.NewConstantTable() }
}

func (self *LocalResolver) Resolve(a *ast.AST) {
  self.errorHandler.Debug("starting local resolver.")
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
  self.errorHandler.Debug("finished local resolver.")
  if self.errorHandler.ErrorOccured() {
    self.errorHandler.Fatal("local resolver failed")
  }
}

func (self *LocalResolver) resolveGvarInitializers(a *ast.AST) {
  variables := a.GetDefinedVariables()
  for i := range variables {
    gvar := variables[i]
    if gvar.HasInitializer() {
      ast.VisitExprNode(self, gvar.GetInitializer())
    }
  }
}

func (self *LocalResolver) resolveConstantValues(a *ast.AST) {
  constants := a.GetConstants()
  for i := range constants {
    constant := constants[i]
    ast.VisitExprNode(self, constant.GetValue())
  }
}

func (self *LocalResolver) resolveFunctions(a *ast.AST) {
  functions := a.GetDefinedFunctions()
  for i := range functions {
    function := functions[i]
    self.pushScope(function.ListParameters())
    ast.VisitStmtNode(self, function.GetBody())
    function.SetScope(self.popScope().(*entity.LocalScope))
  }
}

func (self *LocalResolver) currentScope() core.IScope {
  if len(self.scopeStack) < 1 {
    self.errorHandler.Fatal("stack is empty")
  }
  return self.scopeStack[len(self.scopeStack)-1]
}

func (self *LocalResolver) pushScope(vars []*entity.DefinedVariable) {
  scope := entity.NewLocalScope(self.currentScope())
  for i := range vars {
    v := vars[i]
    if scope.IsDefinedLocally(v.GetName()) {
      self.errorHandler.Errorf("duplicated variable in scope: %s", v.GetName())
    }
    scope.DefineVariable(v)
  }
  self.scopeStack = append(self.scopeStack, scope)
}

func (self *LocalResolver) popScope() core.IScope {
  if len(self.scopeStack) < 1 {
    self.errorHandler.Fatal("stack is empty")
  }
  scope := self.currentScope()
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

func (self *LocalResolver) VisitStmtNode(unknown core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *ast.BlockNode: {
      self.pushScope(node.GetVariables())
      visitBlockNode(self, node)
      node.SetScope(self.popScope().(*entity.LocalScope))
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *LocalResolver) VisitExprNode(unknown core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *ast.StringLiteralNode: {
      ent := self.constantTable.Intern(node.GetValue())
      node.SetEntry(ent)
    }
    case *ast.VariableNode: {
      ent := self.currentScope().GetByName(node.GetName())
      if ent == nil {
        self.errorHandler.Errorf("undefined: %s", node.GetName())
      }
      ent.Refered()
      node.SetEntity(ent)
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *LocalResolver) VisitTypeDefinition(unknown core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
