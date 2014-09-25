package asm

type Comment struct {
  ClassName string
  String string
  IndentLevel int
}

func NewComment(s string) Comment {
  return Comment { "asm.Comment", s, 0 }
}

func (self Comment) IsComment() bool {
  return true
}
