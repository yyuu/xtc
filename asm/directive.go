package asm

import (
  "bitbucket.org/yyuu/xtc/core"
)

type Directive struct {
  ClassName string
  Content string
}

func NewDirective(content string) *Directive {
  return &Directive { "asm.Directive", content }
}

func (self *Directive) AsAssembly() core.IAssembly {
  return self
}

func (self Directive) IsInstruction() bool {
  return false
}

func (self Directive) IsLabel() bool {
  return false
}

func (self Directive) IsDirective() bool {
  return true
}

func (self Directive) IsComment() bool {
  return false
}

func (self *Directive) CollectStatistics(stats core.IStatistics) {
  // does nothing by default
}

func (self *Directive) ToSource(table core.ISymbolTable) string {
  return self.Content
}
