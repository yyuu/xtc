package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type Statistics struct {
  registerUsage map[core.IRegister]int
  insnUsage map[string]int
  symbolUsage map[core.ISymbol]int
}

func NewStatistics() *Statistics {
  registerUsage := make(map[core.IRegister]int)
  insnUsage := make(map[string]int)
  symbolUsage := make(map[core.ISymbol]int)
  return &Statistics { registerUsage, insnUsage, symbolUsage }
}

func CollectStatistics(assemblies []core.IAssembly) *Statistics {
  stats := NewStatistics()
  for i := range assemblies {
    assemblies[i].CollectStatistics(stats)
  }
  return stats
}

func (self *Statistics) AsStatistics() core.IStatistics {
  return self
}

func (self *Statistics) DoesRegisterUsed(reg core.IRegister) bool {
  return 0 < self.NumRegisterUsed(reg)
}

func (self *Statistics) NumRegisterUsed(reg core.IRegister) int {
  return self.registerUsage[reg]
}

func (self *Statistics) RegisterUsed(reg core.IRegister) {
  self.registerUsage[reg]++
}

func (self *Statistics) DoesSymbolUsed(sym core.ISymbol) bool {
  return 0 < self.NumSymbolUsed(sym)
}

func (self *Statistics) NumSymbolUsed(sym core.ISymbol) int {
  return self.symbolUsage[sym]
}

func (self *Statistics) SymbolUsed(sym core.ISymbol) {
  self.symbolUsage[sym]++
}

func (self *Statistics) DoesInstructionUsed(insn string) bool {
  return 0 < self.NumInstructionUsed(insn)
}

func (self *Statistics) NumInstructionUsed(insn string) int {
  return self.insnUsage[insn]
}

func (self *Statistics) InstructionUsed(insn string) {
  self.insnUsage[insn]++
}
