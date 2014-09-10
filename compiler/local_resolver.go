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

  declarations := a.ListDeclaration()
  for i := range declarations {
    toplevel.DeclareEntity(declarations[i])
  }
  definitions := a.ListDefinition()
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
  ys := make([]core.IDefinedVariable, len(xs))
  for i := range xs {
    gvar := xs[i]
    if gvar.HasInitializer() {
      init := gvar.GetInitializer()
      ast.Visit(self, &init)
      gvar = gvar.SetInitializer(init)
    }
    ys[i] = gvar
  }
  a.SetDefinedVariables(ys)
}

func (self *LocalResolver) resolveConstantValues(a *ast.AST) {
  xs := a.GetConstants()
  ys := make([]core.IConstant, len(xs))
  for i := range xs {
    constant := xs[i]
    value := constant.GetValue()
    ast.Visit(self, &value)
    constant = constant.SetValue(value)
    ys[i] = constant
  }
  a.SetConstants(ys)
}

func (self *LocalResolver) resolveFunctions(a *ast.AST) {
  xs := a.GetDefinedFunctions()
  ys := make([]core.IDefinedFunction, len(xs))
  for i := range xs {
    function := xs[i]
    self.pushScope(function.ListParameters())
    body := function.GetBody()
    ast.Visit(self, &body)
    function = function.SetBody(body)
    function = function.SetScope(self.popScope())
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

func (self *LocalResolver) pushScope(vars []core.IDefinedVariable) {
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

func (self *LocalResolver) debugVisit(key string, unknown interface{}) {
  if Verbose < 1 {
    return
  }
  fmt.Println("-----", key, "-----")
  variables := self.currentScope().Variables
  for i := range variables {
    fmt.Println(*variables[i])
  }
  fmt.Println("-----", key, "-----")

  switch a := unknown.(type) {
    case *core.IExprNode: {
      switch b := (*a).(type) {
        case ast.IntegerLiteralNode: {
          fmt.Println("int:", b)
        }
        case ast.VariableNode: {
          fmt.Println("var:", b, ":", b.GetEntity())
        }
      }
    }
  }
}

func (self *LocalResolver) Visit(unknown interface{}) {
  self.debugVisit("BEGIN VISIT", unknown) // TODO: remove this
  switch typed := unknown.(type) {
    case *core.IStmtNode: {
      switch stmt := (*typed).(type) {
        case ast.BlockNode: {
          self.pushScope(stmt.GetVariables())
          stmt.SetScope(self.popScope())
          *typed = stmt
        }
      }
    }
    case *core.IExprNode: {
      switch expr := (*typed).(type) {
        case ast.VariableNode: {
          e := self.currentScope().GetByName(expr.GetName())
          if e == nil {
            panic(fmt.Errorf("undefined: %s", expr.GetName()))
          }
          variable, ok := (*e).(entity.DefinedVariable)
          if ! ok {
            panic(fmt.Errorf("not a variable: %s", expr.GetName()))
          }
          p := &variable
          p.Refered()
          *e = p
          expr.SetEntity(p)
          *typed = expr
        }
        case ast.StringLiteralNode: {
          e := self.constantTable.Intern(expr.GetValue())
          expr.SetEntry(e)
          *typed = expr
        }
      }
    }
  }
  self.debugVisit("END VISIT", unknown) // TODO: remove this
}
