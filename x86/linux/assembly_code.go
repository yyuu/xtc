package x86_linux

import (
  "fmt"
  "strings"
  xtc_asm "bitbucket.org/yyuu/xtc/asm"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/x86"
)

type AssemblyCode struct {
  NaturalType int
  LabelSymbols *xtc_asm.SymbolTable
  virtualStack *x86.VirtualStack
  Assemblies []xtc_core.IAssembly
  CommentIndentLevel int
  Statistics *xtc_asm.Statistics
}

func NewAssemblyCode(naturalType int, labelSymbols *xtc_asm.SymbolTable) *AssemblyCode {
  return &AssemblyCode {
    NaturalType: naturalType,
    LabelSymbols: labelSymbols,
    virtualStack: x86.NewVirtualStack(naturalType),
    Assemblies: []xtc_core.IAssembly { },
    CommentIndentLevel: 0,
    Statistics: nil,
  }
}

func (self *AssemblyCode) GetAssemblies() []xtc_core.IAssembly {
  return self.Assemblies
}

func (self *AssemblyCode) addAll(assemblies []xtc_core.IAssembly) {
  self.Assemblies = append(self.Assemblies, assemblies...)
}

func (self *AssemblyCode) ToSource() string {
  sources := make([]string, len(self.Assemblies))
  for i := range self.Assemblies {
    sources[i] = self.Assemblies[i].ToSource(self.LabelSymbols)
  }
  return strings.Join(sources, "\n") + "\n"
}

func (self *AssemblyCode) GetStatistics() *xtc_asm.Statistics {
  if self.Statistics == nil {
    self.Statistics = xtc_asm.CollectStatistics(self.Assemblies)
  }
  return self.Statistics
}

func (self *AssemblyCode) doesUses(reg *x86.Register) bool {
  if reg == nil { panic("reg is nil") }
  return self.GetStatistics().DoesRegisterUsed(reg)
}

func (self *AssemblyCode) comment(str string) {
  self.Assemblies = append(self.Assemblies, xtc_asm.NewComment(str, self.CommentIndentLevel))
}

func (self *AssemblyCode) indentComment() {
  self.CommentIndentLevel++
}

func (self *AssemblyCode) unindentComment() {
  self.CommentIndentLevel--
}

func (self *AssemblyCode) label1(label *xtc_asm.Label) {
  if label == nil { panic("label is nil") }
  self.Assemblies = append(self.Assemblies, label)
}

func (self *AssemblyCode) label2(sym xtc_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.label1(xtc_asm.NewLabel(sym))
}

func (self *AssemblyCode) reduceLabels() {
  stats := self.GetStatistics()
  result := []xtc_core.IAssembly { }
  for i := range self.Assemblies {
    asm := self.Assemblies[i]
    if asm.IsLabel() && !stats.DoesSymbolUsed(asm.(*xtc_asm.Label).GetSymbol()) {
      // noop
    } else {
      result = append(result, asm)
    }
  }
  self.Assemblies = result
}

func (self *AssemblyCode) directive(direc string) {
  self.Assemblies = append(self.Assemblies, xtc_asm.NewDirective(direc))
}

func (self *AssemblyCode) insn1(op string) {
  instruction := xtc_asm.NewInstruction(op, "", []xtc_core.IOperand { }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn2(op string, a xtc_core.IOperand) {
  if a == nil { panic("a is nil") }
  instruction := xtc_asm.NewInstruction(op, "", []xtc_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn3(op string, suffix string, a xtc_core.IOperand) {
  if a == nil { panic("a is nil") }
  instruction := xtc_asm.NewInstruction(op, suffix, []xtc_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn4(t int, op string, a xtc_core.IOperand) {
  if a == nil { panic("a is nil") }
  instruction := xtc_asm.NewInstruction(op, self.typeSuffix(t), []xtc_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn5(op string, suffix string, a xtc_core.IOperand, b xtc_core.IOperand) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  instruction := xtc_asm.NewInstruction(op, suffix, []xtc_core.IOperand { a, b }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn6(t int, op string, a xtc_core.IOperand, b xtc_core.IOperand) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  instruction := xtc_asm.NewInstruction(op, self.typeSuffix(t), []xtc_core.IOperand { a, b }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) typeSuffix2(t1, t2 int) string {
  return self.typeSuffix(t1) + self.typeSuffix(t2)
}

func (self *AssemblyCode) typeSuffix(t int) string {
  switch t {
    case xtc_asm.TYPE_INT8:  return "b"
    case xtc_asm.TYPE_INT16: return "w"
    case xtc_asm.TYPE_INT32: return "l"
    case xtc_asm.TYPE_INT64: return "q"
    default: {
      panic(fmt.Errorf("unknown register type: %d", t))
    }
  }
}

func (self *AssemblyCode) _file(name string) {
  self.directive(fmt.Sprintf("\t.file\t%q", name))
}

func (self *AssemblyCode) _text() {
  self.directive(fmt.Sprintf("\t.text"))
}

func (self *AssemblyCode) _data() {
  self.directive(fmt.Sprintf("\t.data"))
}

func (self *AssemblyCode) _section(name string) {
  self.directive(fmt.Sprintf("\t.section\t%s", name))
}

func (self *AssemblyCode) _section2(name, flags, t, group, linkage string) {
  self.directive(fmt.Sprintf("\t.section\t%s,%s,%s,%s,%s", name, flags, t, group, linkage))
}

func (self *AssemblyCode) _globl(sym xtc_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.globl %s", sym.GetName()))
}

func (self *AssemblyCode) _local(sym xtc_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf(".local %s", sym.GetName()))
}

func (self *AssemblyCode) _hidden(sym xtc_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.hidden\t%s", sym.GetName()))
}

func (self *AssemblyCode) _comm(sym xtc_core.ISymbol, size, alignment int64) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.comm\t%s, %d, %d", sym.GetName(), size, alignment))
}

func (self *AssemblyCode) _align(n int64) {
  self.directive(fmt.Sprintf("\t.align\t%d", n))
}

func (self *AssemblyCode) _type(sym xtc_core.ISymbol, t string) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.type\t%s, %s", sym.GetName(), t))
}

func (self *AssemblyCode) _size(sym xtc_core.ISymbol, size string) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.size\t%s, %s", sym.GetName(), size))
}

func (self *AssemblyCode) _byte(val xtc_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".byte\t%s", val))
}

func (self *AssemblyCode) _value(val xtc_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".value\t%s", val))
}

func (self *AssemblyCode) _long(val xtc_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".long\t%s", val))
}

func (self *AssemblyCode) _quad(val xtc_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".quad\t%s", val))
}

func (self *AssemblyCode) _string(str string) {
  self.directive(fmt.Sprintf("\t.string\t%q", str))
}

func (self *AssemblyCode) virtualPush(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.virtualStack.Extend(STACK_WORD_SIZE)
  self.mov3(reg, self.virtualStack.Top())
}

func (self *AssemblyCode) virtualPop(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.mov2(self.virtualStack.Top(), reg)
  self.virtualStack.Rewind(STACK_WORD_SIZE)
}

func (self *AssemblyCode) jmp(label *xtc_asm.Label) {
  if label == nil { panic("label is nil") }
  self.insn2("jmp", xtc_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *AssemblyCode) jnz(label *xtc_asm.Label) {
  if label == nil { panic("label is nil") }
  self.insn2("jnz", xtc_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *AssemblyCode) je(label *xtc_asm.Label) {
  if label == nil { panic("label is nil") }
  self.insn2("je", xtc_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *AssemblyCode) cmp(a xtc_core.IOperand, b *x86.Register) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  self.insn6(b.GetTypeId(), "cmp", a, b)
}

func (self *AssemblyCode) sete(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("sete", reg)
}

func (self *AssemblyCode) setne(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setne", reg)
}

func (self *AssemblyCode) seta(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("seta", reg)
}

func (self *AssemblyCode) setae(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setae", reg)
}

func (self *AssemblyCode) setb(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setb", reg)
}

func (self *AssemblyCode) setbe(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setbe", reg)
}

func (self *AssemblyCode) setg(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setg", reg)
}

func (self *AssemblyCode) setge(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setge", reg)
}

func (self *AssemblyCode) setl(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setl", reg)
}

func (self *AssemblyCode) setle(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setle", reg)
}

func (self *AssemblyCode) test(a *x86.Register, b *x86.Register) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  self.insn6(b.GetTypeId(), "test", a, b)
}

func (self *AssemblyCode) push(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn3("push", self.typeSuffix(self.NaturalType), reg)
}

func (self *AssemblyCode) pop(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn3("pop", self.typeSuffix(self.NaturalType), reg)
}

// call function by relative address
func (self *AssemblyCode) call(sym xtc_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.insn2("call", xtc_asm.NewDirectMemoryReference(sym))
}

// call function by absolute address
func (self *AssemblyCode) callAbsolute(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("call", xtc_asm.NewAbsoluteAddress(reg))
}

func (self *AssemblyCode) ret() {
  self.insn1("ret")
}

func (self *AssemblyCode) mov1(src *x86.Register, dest *x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(self.NaturalType, "mov", src, dest)
}

// load
func (self *AssemblyCode) mov2(src xtc_core.IOperand, dest *x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(dest.GetTypeId(), "mov", src, dest)
}

// save
func (self *AssemblyCode) mov3(src *x86.Register, dest xtc_core.IOperand) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(src.GetTypeId(), "mov", src, dest)
}

// for stack access
func (self *AssemblyCode) relocatableMov(src xtc_core.IOperand, dest xtc_core.IOperand) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  instruction := xtc_asm.NewInstruction("mov", self.typeSuffix(self.NaturalType), []xtc_core.IOperand { src, dest }, true)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) movsx(src *x86.Register, dest *x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn5("movs", self.typeSuffix2(src.GetTypeId(), dest.GetTypeId()), src, dest)
}

func (self *AssemblyCode) movzx(src *x86.Register, dest *x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn5("movz", self.typeSuffix2(src.GetTypeId(), dest.GetTypeId()), src, dest)
}

func (self *AssemblyCode) movzb(src *x86.Register, dest *x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn5("movz", "b"+self.typeSuffix(dest.GetTypeId()), src, dest)
}

func (self *AssemblyCode) lea(src xtc_core.IOperand, dest *x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(self.NaturalType, "lea", src, dest)
}

func (self *AssemblyCode) neg(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn4(reg.GetTypeId(), "neg", reg)
}

func (self *AssemblyCode) add(diff xtc_core.IOperand, base *x86.Register) {
  if diff == nil { panic("diff is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "add", diff, base)
}

func (self *AssemblyCode) sub(diff xtc_core.IOperand, base *x86.Register) {
  if diff == nil { panic("diff is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "sub", diff, base)
}

func (self *AssemblyCode) imul(m xtc_core.IOperand, base *x86.Register) {
  if m == nil { panic("m is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "imul", m, base)
}

func (self *AssemblyCode) cltd() {
  self.insn1("cltd")
}

func (self *AssemblyCode) div(base *x86.Register) {
  if base == nil { panic("base is nil") }
  self.insn4(base.GetTypeId(), "div", base)
}

func (self *AssemblyCode) idiv(base *x86.Register) {
  if base == nil { panic("base is nil") }
  self.insn4(base.GetTypeId(), "idiv", base)
}

func (self *AssemblyCode) not(reg *x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn4(reg.GetTypeId(), "not", reg)
}

func (self *AssemblyCode) and(bits xtc_core.IOperand, base *x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "and", bits, base)
}

func (self *AssemblyCode) or(bits xtc_core.IOperand, base *x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "or", bits, base)
}

func (self *AssemblyCode) xor(bits xtc_core.IOperand, base *x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "xor", bits, base)
}

func (self *AssemblyCode) sar(bits *x86.Register, base *x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "sar", bits, base)
}

func (self *AssemblyCode) sal(bits *x86.Register, base *x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "sal", bits, base)
}

func (self *AssemblyCode) shr(bits *x86.Register, base *x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "shr", bits, base)
}
