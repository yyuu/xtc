package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type UndefinedVariable struct {
  ClassName string
  Private bool
  Name string
  TypeNode core.ITypeNode
}

func NewUndefinedVariable(t core.ITypeNode, name string) *UndefinedVariable {
  return &UndefinedVariable { "entity.UndefinedVariable", false, name, t }
}

func NewUndefinedVariables(xs...*UndefinedVariable) []*UndefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []*UndefinedVariable { }
  }
}

func (self UndefinedVariable) String() string {
  return fmt.Sprintf("<entity.UndefinedVariable Name=%s Private=%v TypeNode=%s>", self.Name, self.Private, self.TypeNode)
}

func (self UndefinedVariable) IsEntity() bool {
  return true
}

func (self UndefinedVariable) IsVariable() bool {
  return true
}

func (self UndefinedVariable) IsUndefinedVariable() bool {
  return true
}

func (self UndefinedVariable) IsDefined() bool {
  return false
}

func (self UndefinedVariable) IsConstant() bool {
  return false
}

func (self UndefinedVariable) IsPrivate() bool {
  return true
}

func (self UndefinedVariable) IsRefered() bool {
  return true // FIXME: count up references
}

func (self UndefinedVariable) GetName() string {
  return self.Name
}
