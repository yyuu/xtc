package compiler

import (
  "fmt"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
)

type LocalResolver struct {
  errorHandler *xtc_core.ErrorHandler
  options *xtc_core.Options
  scopeStack []xtc_core.IScope
  constantTable *xtc_entity.ConstantTable
}

func NewLocalResolver(errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options) *LocalResolver {
  return &LocalResolver { errorHandler, options, []xtc_core.IScope { }, xtc_entity.NewConstantTable() }
}

func (self *LocalResolver) Resolve(ast *xtc_ast.AST) (*xtc_ast.AST, error) {
  toplevel := xtc_entity.NewToplevelScope()
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
    return ast, fmt.Errorf("found %d error(s).", self.errorHandler.GetErrors())
  }
  return ast, nil
}

func (self *LocalResolver) resolveGvarInitializers(ast *xtc_ast.AST) {
  variables := ast.GetDefinedVariables()
  for i := range variables {
    gvar := variables[i]
    if gvar.HasInitializer() {
      xtc_ast.VisitExprNode(self, gvar.GetInitializer())
    }
  }
}

func (self *LocalResolver) resolveConstantValues(ast *xtc_ast.AST) {
  constants := ast.GetConstants()
  for i := range constants {
    constant := constants[i]
    xtc_ast.VisitExprNode(self, constant.GetValue())
  }
}

func (self *LocalResolver) resolveFunctions(ast *xtc_ast.AST) {
  functions := ast.GetDefinedFunctions()
  for i := range functions {
    function := functions[i]
    self.pushScope(function.ListParameters())
    xtc_ast.VisitStmtNode(self, function.GetBody())
    function.SetScope(self.popScope().(*xtc_entity.LocalScope))
  }
}

func (self *LocalResolver) currentScope() xtc_core.IScope {
  if len(self.scopeStack) < 1 {
    self.errorHandler.Fatal("stack is empty")
  }
  return self.scopeStack[len(self.scopeStack)-1]
}

func (self *LocalResolver) pushScope(vars []*xtc_entity.DefinedVariable) {
  scope := xtc_entity.NewLocalScope(self.currentScope())
  for i := range vars {
    v := vars[i]
    if scope.IsDefinedLocally(v.GetName()) {
      self.errorHandler.Errorf("duplicated variable in scope: %s", v.GetName())
    }
    scope.DefineVariable(v)
  }
  self.scopeStack = append(self.scopeStack, scope)
}

func (self *LocalResolver) popScope() xtc_core.IScope {
  if len(self.scopeStack) < 1 {
    self.errorHandler.Fatal("stack is empty")
  }
  scope := self.currentScope()
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

func (self *LocalResolver) VisitStmtNode(unknown xtc_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.BlockNode: {
      self.pushScope(node.GetVariables())
      visitBlockNode(self, node)
      node.SetScope(self.popScope().(*xtc_entity.LocalScope))
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *LocalResolver) VisitExprNode(unknown xtc_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.StringLiteralNode: {
      ent := self.constantTable.Intern(node.GetValue())
      node.SetEntry(ent)
    }
    case *xtc_ast.VariableNode: {
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

func (self *LocalResolver) VisitTypeDefinition(unknown xtc_core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
