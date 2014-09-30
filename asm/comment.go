package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type Comment struct {
  ClassName string
  String string
  IndentLevel int
}

func NewComment(s string, indentLevel int) *Comment {
  return &Comment { "asm.Comment", s, indentLevel }
}

func (self *Comment) AsAssembly() core.IAssembly {
  return self
}

func (self Comment) IsInstruction() bool {
  return false
}

func (self Comment) IsLabel() bool {
  return false
}

func (self Comment) IsDirective() bool {
  return false
}

func (self Comment) IsComment() bool {
  return true
}

func (self *Comment) CollectStatistics(stats core.IStatistics) {
  // does nothing by default
}
