package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

type LocalScope struct {
  Parent core.IScope
  Children []*LocalScope
  Variables map[string]*DefinedVariable
}

func NewLocalScope(parent core.IScope) *LocalScope {
  return &LocalScope { parent, []*LocalScope { }, make(map[string]*DefinedVariable) }
}

func (self *LocalScope) IsToplevel() bool {
  return false
}

func (self *LocalScope) GetToplevel() core.IScope {
  return self.Parent.GetToplevel()
}

func (self *LocalScope) GetParent() core.IScope {
  return self.Parent
}

func (self *LocalScope) GetChildren() []*LocalScope {
  return self.Children
}

func (self *LocalScope) AddChild(scope *LocalScope) {
  self.Children = append(self.Children, scope)
}

func (self *LocalScope) GetByName(name string) core.IEntity {
  ent := self.Variables[name]
  if ent != nil {
    return ent
  } else {
    return self.Parent.GetByName(name)
  }
}

func (self *LocalScope) DefineVariable(v *DefinedVariable) error {
  name := v.GetName()
  if self.IsDefinedLocally(name) {
    return fmt.Errorf("duplicated variable: %s", name)
  }
  self.Variables[name] = v
  return nil
}

func (self *LocalScope) IsDefinedLocally(name string) bool {
  _, ok := self.Variables[name]
  return ok
}

func (self *LocalScope) CheckReferences(errorHandler *core.ErrorHandler) {
  for _, ent := range self.Variables {
    if !ent.IsRefered() {
      errorHandler.Warnf("unused variable: %s", ent.GetName())
    }
  }
  for i := range self.Children {
    scope := self.Children[i]
    scope.CheckReferences(errorHandler)
  }
}

func (self *LocalScope) AllocateTmp(typeNode core.ITypeNode) *DefinedVariable {
  v := temporaryDefinedVariable(typeNode)
  err := self.DefineVariable(v)
  if err != nil { panic(err) }
  return v
}

/**
 * Returns all local variables in this scope.
 * The result DOES includes all nested local variables.
 * while it does NOT include static local variables.
 */
func (self *LocalScope) AllLocalVariables() []*DefinedVariable {
  result := []*DefinedVariable { }
  scopes := self.allLocalScopes()
  for i := range scopes {
    result = append(result, scopes[i].GetLocalVariables()...)
  }
  return result
}

/**
 * Returns local variables defined in this scope.
 * Does not includes children&s local variables.
 * Does NOT include static local variables.
 */
func (self *LocalScope) GetLocalVariables() []*DefinedVariable {
  result := []*DefinedVariable { }
  for _, v := range self.Variables {
    if ! v.IsPrivate() {
      result = append(result, v)
    }
  }
  return result
}

func (self *LocalScope) GetStaticLocalVariables() []*DefinedVariable {
  result := []*DefinedVariable { }
  scopes := self.allLocalScopes()
  for i := range scopes {
    for _, v := range scopes[i].Variables {
      if v.IsPrivate() {
        result = append(result, v)
      }
    }
  }
  return result
}

func (self *LocalScope) allLocalScopes() []*LocalScope {
  result := []*LocalScope { self }
  for i := range self.Children {
    result = append(result, self.Children[i].allLocalScopes()...)
  }
  return result
}
