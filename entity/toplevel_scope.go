package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
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

func (self *ToplevelScope) AddChild(scope core.IScope) {
  self.Children = append(self.Children, scope.(*LocalScope))
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

func (self *ToplevelScope) AllGlobalVariables() []core.IVariable {
  variables := []core.IVariable { }
  for _, v := range self.Entities {
    if v.IsVariable() {
      variables = append(variables, v)
    }
  }
  staticLocalVariables := self.GetStaticLocalVariables()
  for i := range staticLocalVariables {
    var v core.IVariable = staticLocalVariables[i]
    variables = append(variables, v)
  }
  return variables
}

func (self *ToplevelScope) GetDefinedGlobalScopeVariables() []*DefinedVariable {
  result := []*DefinedVariable { }
  for _, ent := range self.Entities {
    v, ok := ent.(*DefinedVariable)
    if ok {
      result = append(result, v)
    }
  }
  result = append(result, self.GetStaticLocalVariables()...)
  return result
}

func (self *ToplevelScope) GetStaticLocalVariables() []*DefinedVariable {
  if len(self.StaticLocalVariables) < 1 {
    for i := range self.Children {
      s := self.Children[i]
      self.StaticLocalVariables = append(self.StaticLocalVariables, s.GetStaticLocalVariables()...)
    }
    seqTable := make(map[string]int)
    vars := self.StaticLocalVariables
    for i := range vars {
      seq, ok := seqTable[vars[i].GetName()]
      if ! ok {
        vars[i].SetSequence(0)
        seqTable[vars[i].GetName()] = 1
      } else {
        vars[i].SetSequence(seq)
        seqTable[vars[i].GetName()] = seq + 1
      }
    }
  }
  return self.StaticLocalVariables
}

func (self *ToplevelScope) CheckReferences(errorHandler *core.ErrorHandler) {
  for _, ent := range self.Entities {
    if ent.IsDefined() && ent.IsPrivate() && !ent.IsConstant() && !ent.IsRefered() {
      errorHandler.Warnf("unused variable: %s", ent.GetName())
    }
  }
  for i := range self.Children {
    scope := self.Children[i]
    scope.CheckReferences(errorHandler)
  }
}
