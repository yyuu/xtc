package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type ToplevelScope struct {
  Entities map[string]duck.IEntity
  Children []LocalScope
}

func NewToplevelScope() ToplevelScope {
  return ToplevelScope { make(map[string]duck.IEntity), []LocalScope { } }
}

func (self ToplevelScope) IsToplevel() bool {
  return true
}

func (self ToplevelScope) GetToplevel() duck.IScope {
  return self
}

func (self ToplevelScope) GetParent() duck.IScope {
  return self
}

func (self *ToplevelScope) AddChild(s LocalScope) {
  self.Children = append(self.Children, s)
}

func (self ToplevelScope) Get(name string) duck.IEntity {
  return self.Entities[name]
}

func (self ToplevelScope) DeclareEntity(entity duck.IEntity) {
  name := entity.GetName()
  e := self.Get(name)
  if e != nil {
    panic(fmt.Sprintf("semantic exception: duplicated declaration: %s", name))
  } else {
    self.Entities[name] = entity
  }
}

func (self ToplevelScope) DefineEntity(entity duck.IEntity) {
  name := entity.GetName()
  e := self.Get(name)
  if e != nil && e.IsDefined() {
    panic(fmt.Sprintf("semantic exception: duplicated definition: %s", name))
  } else {
    self.Entities[name] = entity
  }
}

/*
func (self ToplevelScope) AllGlobalVariables() {
}
 */

/*
func (self ToplevelScope) DefinedGlobalScopeVariables() {
}
 */

/*
func (self ToplevelScope) StaticLocalVariables() {
}
 */

/*
func (self ToplevelScope) CheckReferences() {
}
 */

type LocalScope struct {
  Parent duck.IScope
  Variables map[string]duck.IEntity
}

func NewLocalScope(parent duck.IScope) LocalScope {
  return LocalScope { parent, make(map[string]duck.IEntity) }
}

func (self LocalScope) IsToplevel() bool {
  return false
}

func (self LocalScope) GetToplevel() duck.IScope {
  if self.Parent.IsToplevel() {
    return self.Parent
  } else {
    return self.Parent.GetToplevel()
  }
}

func (self LocalScope) GetParent() duck.IScope {
  return self
}
