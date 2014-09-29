package asm

type Directive struct {
  ClassName string
  Content string
}

func NewDirective(content string) *Directive {
  return &Directive { "asm.Directive", content }
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
