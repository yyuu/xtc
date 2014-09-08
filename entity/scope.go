package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type VariableScope struct {
  Parent *VariableScope
  Children []*VariableScope
  Variables map[string]*duck.IEntity
}

func NewToplevelScope() *VariableScope {
  return &VariableScope { nil, []*VariableScope { }, make(map[string]*duck.IEntity) }
}

func NewLocalScope(parent *VariableScope) *VariableScope {
  return &VariableScope { parent, []*VariableScope { }, make(map[string]*duck.IEntity) }
}

func (self *VariableScope) IsVariableScope() bool {
  return true
}

func (self *VariableScope) IsToplevel() bool {
  return self.Parent == nil
}

func (self *VariableScope) GetParent() duck.IVariableScope {
  return self.Parent
}

func (self *VariableScope) AddChild(scope *VariableScope) {
  self.Children = append(self.Children, scope)
}

func (self *VariableScope) GetByName(name string) *duck.IEntity {
  return self.Variables[name]
}

func (self *VariableScope) DeclareEntity(entity duck.IEntity) {
  name := entity.GetName()
  e := self.GetByName(name)
  if e != nil {
    panic(fmt.Errorf("duplicated declaration: %s", name))
  } else {
    self.Variables[name] = &entity
  }
}

func (self *VariableScope) DefineEntity(entity duck.IEntity) {
  name := entity.GetName()
  e := self.GetByName(name)
  if e != nil && (*e).IsDefined() {
    panic(fmt.Errorf("duplicated definition: %s", name))
  } else {
    self.Variables[name] = &entity
  }
}

func (self *VariableScope) DefineVariable(v duck.IDefinedVariable) {
  name := v.GetName()
  if self.IsDefinedLocally(name) {
    panic(fmt.Errorf("duplicated variable: %s", name))
  }
  entity := v.(duck.IEntity)
  self.Variables[name] = &entity
}

func (self *VariableScope) GetToplevel() duck.IVariableScope {
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

type ConstantTable struct {
  Constants map[string]*duck.IEntity
}

func NewConstantTable() *ConstantTable {
  return &ConstantTable { make(map[string]*duck.IEntity) }
}

func (self *ConstantTable) IsConstantTable() bool {
  return true
}

func (self *ConstantTable) Intern(s string) *ConstantEntry {
  return &ConstantEntry { s }
}

type ConstantEntry struct {
  s string
}

func (self *ConstantEntry) IsConstantEntry() bool {
  return true
}
