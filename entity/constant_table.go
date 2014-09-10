package entity

import (
  "bitbucket.org/yyuu/bs/core"
)

type ConstantTable struct {
  Constants map[string]*core.IEntity
}

func NewConstantTable() *ConstantTable {
  return &ConstantTable { make(map[string]*core.IEntity) }
}

func (self *ConstantTable) IsConstantTable() bool {
  return true
}

func (self *ConstantTable) Intern(s string) *ConstantEntry {
  return NewConstantEntry(s)
}
