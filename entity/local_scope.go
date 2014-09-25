package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
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

func (self LocalScope) GetParent() core.IScope {
  return self.Parent
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

func (self *LocalScope) DefineVariable(v *DefinedVariable) {
  name := v.GetName()
  if self.IsDefinedLocally(name) {
    panic(fmt.Errorf("duplicated variable: %s", name))
  }
  self.Variables[name] = v
}

func (self *LocalScope) IsDefinedLocally(name string) bool {
  _, ok := self.Variables[name]
  return ok
}

func (self *LocalScope) CheckReferences(errorHandler *core.ErrorHandler) {
  for _, ent := range self.Variables {
    if !ent.IsRefered() {
      errorHandler.Warnf("unused variable: %s\n", ent.GetName())
    }
  }
  for i := range self.Children {
    scope := self.Children[i]
    scope.CheckReferences(errorHandler)
  }
}

func (self *LocalScope) AllocateTmp(typeNode core.ITypeNode) *DefinedVariable {
  v := temporaryDefinedVariable(typeNode)
  self.DefineVariable(v)
  return v
}
