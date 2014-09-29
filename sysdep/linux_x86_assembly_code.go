package sysdep

import (
  "fmt"
  bs_asm "bitbucket.org/yyuu/bs/asm"
  bs_core "bitbucket.org/yyuu/bs/core"
)

type LinuxX86AssemblyCode struct {
  NaturalType int
  StackWordSize int64
  LabelSymbols *bs_asm.SymbolTable
//virtualStack VirtualStack
  Assemblies []bs_core.IAssembly
  commentIndentLevel int
}

func NewLinuxX86AssemblyCode(naturalType int, stackWordSize int64, labelSymbols *bs_asm.SymbolTable) *LinuxX86AssemblyCode {
  assemblies := []bs_core.IAssembly { }
  return &LinuxX86AssemblyCode { naturalType, stackWordSize, labelSymbols, assemblies, 0 }
}

func (self *LinuxX86AssemblyCode) comment(str string) {
  self.Assemblies = append(self.Assemblies, bs_asm.NewComment(str, self.commentIndentLevel))
}

func (self *LinuxX86AssemblyCode) indentComment() {
  self.commentIndentLevel++
}

func (self *LinuxX86AssemblyCode) unindentComment() {
  self.commentIndentLevel--
}

func (self *LinuxX86AssemblyCode) label(label *bs_asm.Label) {
  self.Assemblies = append(self.Assemblies, label)
}

func (self *LinuxX86AssemblyCode) directive(direc string) {
  self.Assemblies = append(self.Assemblies, bs_asm.NewDirective(direc))
}

func (self *LinuxX86AssemblyCode) _file(name string) {
  self.directive(fmt.Sprintf(".file\t%s", name))
}

func (self *LinuxX86AssemblyCode) _text() {
  self.directive(fmt.Sprintf("\t.text"))
}

func (self *LinuxX86AssemblyCode) _data() {
  self.directive(fmt.Sprintf("\t.data"))
}

func (self *LinuxX86AssemblyCode) _section(name string) {
  self.directive(fmt.Sprintf("\t.section\t%s", name))
}

func (self *LinuxX86AssemblyCode) _section2(name, flags, t, group, linkage string) {
  self.directive(fmt.Sprintf("\t.section\t%s,%s,%s,%s,%s", name, flags, t, group, linkage))
}

func (self *LinuxX86AssemblyCode) _globl(sym bs_core.ISymbol) {
  self.directive(fmt.Sprintf(".globl %s", sym.GetName()))
}

func (self *LinuxX86AssemblyCode) _local(sym bs_core.ISymbol) {
  self.directive(fmt.Sprintf(".local %s", sym.GetName()))
}

func (self *LinuxX86AssemblyCode) _hidden(sym bs_core.ISymbol) {
  self.directive(fmt.Sprintf("\t.hidden\t%s", sym.GetName()))
}

func (self *LinuxX86AssemblyCode) _comm(sym bs_core.ISymbol, size, alignment int64) {
  self.directive(fmt.Sprintf("\t.comm\t%s,%d,%d", sym.GetName(), size, alignment))
}

func (self *LinuxX86AssemblyCode) _align(n int64) {
  self.directive(fmt.Sprintf("\t.align\t%d", n))
}

func (self *LinuxX86AssemblyCode) _type(sym bs_core.ISymbol, t string) {
  self.directive(fmt.Sprintf("\t.type\t%s,%s", sym.GetName(), t))
}

func (self *LinuxX86AssemblyCode) _size(sym bs_core.ISymbol, size string) {
  self.directive(fmt.Sprintf("\t.size\t%s,%s", sym.GetName(), size))
}

func (self *LinuxX86AssemblyCode) _byte(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".byte\t%s", val))
}

func (self *LinuxX86AssemblyCode) _value(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".value\t%s", val))
}

func (self *LinuxX86AssemblyCode) _long(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".long\t%s", val))
}

func (self *LinuxX86AssemblyCode) _quad(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".quad\t%s", val))
}

func (self *LinuxX86AssemblyCode) _string(str string) {
  self.directive(fmt.Sprintf("\t.string\t%s", str))
}
