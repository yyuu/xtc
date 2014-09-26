package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type ToplevelScope struct {
  Children []*LocalScope
  Entities map[string]core.IEntity
  StaticLocalVariables []*DefinedVariable
}

func NewToplevelScope() *ToplevelScope {
  return &ToplevelScope { []*LocalScope { }, make(map[string]core.IEntity), []*DefinedVariable { } }
}

func (self *ToplevelScope) IsToplevel() bool {
  return true
}

func (self *ToplevelScope) GetToplevel() core.IScope {
  return self
}

func (self *ToplevelScope) GetParent() core.IScope {
  return nil
}

func (self *ToplevelScope) AddChild(scope *LocalScope) {
  self.Children = append(self.Children, scope)
}

func (self *ToplevelScope) GetByName(name string) core.IEntity {
  ent := self.Entities[name]
  if ent == nil {
    panic(fmt.Errorf("unresolved reference: %s", name))
  }
  return ent
}

func (self *ToplevelScope) DeclareEntity(entity core.IEntity) {
  name := entity.GetName()
  ent := self.Entities[name]
  if ent != nil {
    panic(fmt.Errorf("duplicated declaration: %s", name))
  } else {
    self.Entities[name] = entity
  }
}

func (self *ToplevelScope) DefineEntity(entity core.IEntity) {
  if ! self.IsToplevel() {
    panic("cannot define entity to non-toplevel scope")
  }
  name := entity.GetName()
  ent := self.Entities[name]
  if ent != nil && ent.IsDefined() {
    panic(fmt.Errorf("duplicated definition: %s", name))
  } else {
    self.Entities[name] = entity
  }
}

/*
func (self *ToplevelScope) AllGlobalVariables() {
}
 */

/*
func (self *ToplevelScope) DefinedGlobalScopeVariables() {
}
 */

/*
func (self *ToplevelScope) StaticLocalVariables() {
}
 */

func (self *ToplevelScope) CheckReferences(errorHandler *core.ErrorHandler) {
  for _, ent := range self.Entities {
    if ent.IsDefined() && ent.IsPrivate() && !ent.IsConstant() && !ent.IsRefered() {
      errorHandler.Warnf("unused variable: %s\n", ent.GetName())
    }
  }
  for i := range self.Children {
    scope := self.Children[i]
    scope.CheckReferences(errorHandler)
  }
}
