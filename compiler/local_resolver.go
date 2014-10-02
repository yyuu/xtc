package compiler

import (
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
)

type LocalResolver struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
  scopeStack []bs_core.IScope
  constantTable *bs_entity.ConstantTable
}

func NewLocalResolver(errorHandler *bs_core.ErrorHandler, options *bs_core.Options) *LocalResolver {
  return &LocalResolver { errorHandler, options, []bs_core.IScope { }, bs_entity.NewConstantTable() }
}

func (self *LocalResolver) Resolve(ast *bs_ast.AST) {
  self.errorHandler.Debug("starting local resolver.")
  toplevel := bs_entity.NewToplevelScope()
  self.scopeStack = append(self.scopeStack, toplevel)

  declarations := ast.ListDeclarations()
  for i := range declarations {
    toplevel.DeclareEntity(declarations[i])
  }
  definitions := ast.ListDefinitions()
  for i := range definitions {
    toplevel.DefineEntity(definitions[i])
  }

  self.resolveGvarInitializers(ast)
  self.resolveConstantValues(ast)
  self.resolveFunctions(ast)

  toplevel.CheckReferences(self.errorHandler)

  ast.SetScope(toplevel)
  ast.SetConstantTable(self.constantTable)
  if self.errorHandler.ErrorOccured() {
    self.errorHandler.Fatalf("found %d error(s).", self.errorHandler.GetErrors())
  } else {
    self.errorHandler.Debug("finished local resolver.")
  }
}

func (self *LocalResolver) resolveGvarInitializers(ast *bs_ast.AST) {
  variables := ast.GetDefinedVariables()
  for i := range variables {
    gvar := variables[i]
    if gvar.HasInitializer() {
      bs_ast.VisitExprNode(self, gvar.GetInitializer())
    }
  }
}

func (self *LocalResolver) resolveConstantValues(ast *bs_ast.AST) {
  constants := ast.GetConstants()
  for i := range constants {
    constant := constants[i]
    bs_ast.VisitExprNode(self, constant.GetValue())
  }
}

func (self *LocalResolver) resolveFunctions(ast *bs_ast.AST) {
  functions := ast.GetDefinedFunctions()
  for i := range functions {
    function := functions[i]
    self.pushScope(function.ListParameters())
    bs_ast.VisitStmtNode(self, function.GetBody())
    function.SetScope(self.popScope().(*bs_entity.LocalScope))
  }
}

func (self *LocalResolver) currentScope() bs_core.IScope {
  if len(self.scopeStack) < 1 {
    self.errorHandler.Fatal("stack is empty")
  }
  return self.scopeStack[len(self.scopeStack)-1]
}

func (self *LocalResolver) pushScope(vars []*bs_entity.DefinedVariable) {
  scope := bs_entity.NewLocalScope(self.currentScope())
  for i := range vars {
    v := vars[i]
    if scope.IsDefinedLocally(v.GetName()) {
      self.errorHandler.Errorf("duplicated variable in scope: %s", v.GetName())
    }
    scope.DefineVariable(v)
  }
  self.scopeStack = append(self.scopeStack, scope)
}

func (self *LocalResolver) popScope() bs_core.IScope {
  if len(self.scopeStack) < 1 {
    self.errorHandler.Fatal("stack is empty")
  }
  scope := self.currentScope()
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

func (self *LocalResolver) VisitStmtNode(unknown bs_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.BlockNode: {
      self.pushScope(node.GetVariables())
      visitBlockNode(self, node)
      node.SetScope(self.popScope().(*bs_entity.LocalScope))
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *LocalResolver) VisitExprNode(unknown bs_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.StringLiteralNode: {
      ent := self.constantTable.Intern(node.GetValue())
      node.SetEntry(ent)
    }
    case *bs_ast.VariableNode: {
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

func (self *LocalResolver) VisitTypeDefinition(unknown bs_core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
