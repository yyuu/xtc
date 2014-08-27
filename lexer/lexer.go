package lexer

import (
  "fmt"
  "regexp"
  "strings"
  "bitbucket.org/yyuu/bs/strscan"
)

const (
  SPACES = iota
  BLOCK_COMMENT
  LINE_COMMENT
  VOID
  CHAR
  SHORT
  INT
  LONG
  STRUCT
  UNION
  ENUM
  STATIC
  EXTERN
  CONST
  SIGNED
  UNSIGNED
  IF
  ELSE
  SWITCH
  CASE
  DEFAULT
  WHILE
  DO
  FOR
  RETURN
  BREAK
  CONTINUE
  GOTO
  TYPEDEF
  IMPORT
  SIZEOF
  IDENTIFIER
  INTEGER
  CHARACTER
  STRING
  OPERATOR
)

var id2name map[int]string = map[int]string {
  SPACES: "SPACES",
  BLOCK_COMMENT: "BLOCK_COMMENT",
  LINE_COMMENT: "LINE_COMMENT",
  VOID: "VOID",
  CHAR: "CHAR",
  SHORT: "",
  INT: "INT",
  LONG: "LONG",
  STRUCT: "STRUCT",
  UNION: "UNION",
  ENUM: "ENUM",
  STATIC: "STATIC",
  EXTERN: "EXTERN",
  CONST: "CONST",
  SIGNED: "SIGNED",
  UNSIGNED: "UNSIGNED",
  IF: "IF",
  ELSE: "ELSE",
  SWITCH: "SWITCH",
  CASE: "CASE",
  DEFAULT: "DEFAULT",
  WHILE: "WHILE",
  DO: "DO",
  FOR: "FOR",
  RETURN: "RETURN",
  BREAK: "BREAK",
  CONTINUE: "CONTINUE",
  GOTO: "GOTO",
  TYPEDEF: "TYPEDEF",
  IMPORT: "IMPORT",
  SIZEOF: "SIZEOF",
  IDENTIFIER: "IDENTIFIER",
  INTEGER: "INTEGER",
  CHARACTER: "CHARACTER",
  STRING: "STRING",
  OPERATOR: "OPERATOR",
}

type Lexer struct {
  scanner strscan.StringScanner
  Filename string
  LineNumber int
  LineOffset int
}

type Token struct {
  Id int
  Literal string
  Filename string
  LineNumber int
  LineOffset int
}

func (t *Token) ToString() string {
  name, ok := id2name[t.Id]
  if ok {
    return fmt.Sprintf("<Token:%s (%s:%d,%d) %q>", name, t.Filename, t.LineNumber, t.LineOffset, t.Literal)
  } else {
    return fmt.Sprintf("<Token:%d (%s:%d,%d) %q>", t.Id, t.Filename, t.LineNumber, t.LineOffset, t.Literal)
  }
}

func NewLexer(filename string, source string) *Lexer {
  return &Lexer {
    scanner: strscan.NewStringScanner(source),
    Filename: filename,
    LineNumber: 0,
    LineOffset: 0,
  }
}

var keywords map[string]int = map[string]int {
  "void": VOID,
  "char": CHAR,
  "short": SHORT,
  "int": INT,
  "long": LONG,
  "struct": STRUCT,
  "union": UNION,
  "enum": ENUM,
  "static": STATIC,
  "extern": EXTERN,
  "const": CONST,
  "signed": SIGNED,
  "unsigned": UNSIGNED,
  "if": IF,
  "else": ELSE,
  "switch": SWITCH,
  "case": CASE,
  "default": DEFAULT,
  "while": WHILE,
  "do": DO,
  "for": FOR,
  "return": RETURN,
  "break": BREAK,
  "continue": CONTINUE,
  "goto": GOTO,
  "typedef": TYPEDEF,
  "import": IMPORT,
  "sizeof": SIZEOF,
}

func (self *Lexer) GetToken() (t *Token) {
  if self.scanner.IsEOS() {
    return nil
  }

  t = self.readSpaces()
  if t != nil {
    return t
  }

  t = self.readBlockComment()
  if t != nil {
    return t
  }
  t = self.readLineComment()
  if t != nil {
    return t
  }

  t = self.readKeyword()
  if t != nil {
    return t
  }

  t = self.readIdentifier()
  if t != nil {
    return t
  }

  t = self.readInteger()
  if t != nil {
    return t
  }

  t = self.readCharacter()
  if t != nil {
    return t
  }

  t = self.readString()
  if t != nil {
    return t
  }

  t = self.readOperator()
  if t != nil {
    return t
  }

  panic("lexer error")
}

func (self *Lexer) newToken(id int, literal string) (t *Token) {
  t = &Token {
    Id: id,
    Literal: literal,
    Filename: self.Filename,
    LineNumber: self.LineNumber,
    LineOffset: self.LineOffset,
  }

  self.LineNumber += strings.Count(literal, "\n")
  i := strings.LastIndex(literal, "\n")
  if i < 0 {
    self.LineOffset += len(literal)
  } else {
    self.LineOffset = len(literal[i:])
  }

  return t
}

func (self *Lexer) readBlockComment() *Token {
  s := self.scanner.Scan("/\\*")
  if s == "" {
    return nil
  }
  more := self.scanner.ScanUntil("\\*/")
  if more == "" {
    panic("lexer error")
  }
  return self.newToken(BLOCK_COMMENT, s + more)
}

func (self *Lexer) readLineComment() *Token {
  s := self.scanner.Scan("//")
  if s == "" {
    return nil
  }
  more := self.scanner.ScanUntil("(\n|\r\n|\r)")
  if more == "" {
    panic("lexer error")
  }
  return self.newToken(LINE_COMMENT, s + more)
}

func (self *Lexer) readSpaces() *Token {
  s := self.scanner.Scan("[ \t\n\r\f]+")
  if s == "" {
    return nil
  }
  return self.newToken(SPACES, s)
}

func (self *Lexer) readIdentifier() *Token {
  s := self.scanner.Scan("[_A-Za-z][_0-9A-Za-z]*")
  if s == "" {
    return nil
  }
  return self.newToken(IDENTIFIER, s)
}

func (self *Lexer) readInteger() *Token {
  s := self.scanner.Scan("([1-9][0-9]*U?L?|0[Xx][0-9A-Fa-f]+U?L?|0[0-7]*U?L?)")
  if s == "" {
    return nil
  }
  return self.newToken(INTEGER, s)
}

func (self *Lexer) readKeyword() *Token {
  for keyword, id := range keywords {
    s := self.scanner.Scan(regexp.QuoteMeta(keyword))
    if s != "" {
      return self.newToken(id, s)
    }
  }
  return nil
}

func (self *Lexer) readCharacter() *Token {
  s := self.scanner.Scan("'")
  if s == "" {
    return nil
  }
  // TODO: handle escape character properly
  more := self.scanner.ScanUntil("'")
  if more == "" {
    panic("lexer error")
  }
  return self.newToken(CHARACTER, s + more)
}

func (self *Lexer) readString() *Token {
  s := self.scanner.Scan("\"")
  if s == "" {
    return nil
  }
  // TODO: handle escape character properly
  more := self.scanner.ScanUntil("\"")
  if more == "" {
    panic("lexer error")
  }
  return self.newToken(STRING, s + more)
}

func (self *Lexer) readOperator() *Token {
  operators := []string {
    "<<=", ">>=", "...",
    "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=",
    "||", "&&", ">=", "<=", "==", "!=", "++", "--", ">>", "<<", "->",
    "=", ">", "<", ":", ";", "?", "{", "}", "(", ")", "+", "-", "!", "~",
    "*", "&", "|", "^", "&", "+", "-", "*", "/", "%", "[", "]", ".", ",",
  }
  for i := range operators {
    s := self.scanner.Scan(regexp.QuoteMeta(operators[i]))
    if s != "" {
      return self.newToken(OPERATOR, s)
    }
  }
  return nil
}
