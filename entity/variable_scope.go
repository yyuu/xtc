package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type VariableScope struct {
  Parent *VariableScope
  Children []*VariableScope
  Variables map[string]core.IEntity
}

func NewToplevelScope() *VariableScope {
  return &VariableScope { nil, []*VariableScope { }, make(map[string]core.IEntity) }
}

func NewLocalScope(parent *VariableScope) *VariableScope {
  return &VariableScope { parent, []*VariableScope { }, make(map[string]core.IEntity) }
}

func (self *VariableScope) IsVariableScope() bool {
  return true
}

func (self *VariableScope) IsToplevel() bool {
  return self.Parent == nil
}

func (self *VariableScope) GetParent() *VariableScope {
  return self.Parent
}

func (self *VariableScope) AddChild(scope *VariableScope) {
  self.Children = append(self.Children, scope)
}

func (self *VariableScope) GetByName(name string) core.IEntity {
  ent := self.Variables[name]
  if ent == nil && !self.IsToplevel() {
    return self.Parent.GetByName(name)
  } else {
    return ent
  }
}

func (self *VariableScope) DeclareEntity(entity core.IEntity) {
  name := entity.GetName()
  e := self.GetByName(name)
  if e != nil {
    panic(fmt.Errorf("duplicated declaration: %s", name))
  } else {
    self.Variables[name] = entity
  }
}

func (self *VariableScope) DefineEntity(entity core.IEntity) {
  name := entity.GetName()
  e := self.GetByName(name)
  if e != nil && e.IsDefined() {
    panic(fmt.Errorf("duplicated definition: %s", name))
  } else {
    self.Variables[name] = entity
  }
}

func (self *VariableScope) DefineVariable(v *DefinedVariable) {
  name := v.GetName()
  if self.IsDefinedLocally(name) {
    panic(fmt.Errorf("duplicated variable: %s", name))
  }
  var entity core.IEntity = v
  self.Variables[name] = entity
}

func (self *VariableScope) GetToplevel() *VariableScope {
  if self.Parent == nil {
    return self
  } else {
    return self.Parent.GetToplevel()
  }
}

func (self *VariableScope) IsDefinedLocally(name string) bool {
  _, ok := self.Variables[name]
  return ok
}

/*
func (self *VariableScope) AllGlobalVariables() {
}
 */

/*
func (self *VariableScope) DefinedGlobalScopeVariables() {
}
 */

/*
func (self *VariableScope) StaticLocalVariables() {
}
 */

func (self *VariableScope) CheckReferences() {
  for _, ent := range self.Variables {
    if ent.IsDefined() && ent.IsPrivate() && !ent.IsConstant() && !ent.IsRefered() {
      panic(fmt.Errorf("unused variable: %s", ent.GetName()))
    }
  }
  for i := range self.Children {
    a := self.Children[i]
    for j := range a.Children {
      b := a.Children[j]
      b.CheckReferences()
    }
  }
}
