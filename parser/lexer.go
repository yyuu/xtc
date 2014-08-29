package parser

import (
  "fmt"
  "regexp"
  "strings"
  "unicode/utf8"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/strscan"
)

type lex struct {
  scanner strscan.StringScanner
  Filename string
  LineNumber int
  LineOffset int
  ignoreSpaces bool
  ignoreComments bool
  nodes []ast.INode
  error error
}

func (self lex) String() string {
  source := fmt.Sprintf("%s...", self.scanner.Peek(16))
  return fmt.Sprintf("%s:%d: %q", self.Filename, self.LineNumber, source)
}

type token struct {
  Id int
  Literal string
  Filename string
  LineNumber int
  LineOffset int
}

func (self token) String() string {
  return fmt.Sprintf("#<token:%d %s:%d,%d %q>", self.Id, self.Filename, self.LineNumber, self.LineOffset, self.Literal)
}

func lexer(filename string, source string) *lex {
  return &lex {
    scanner: strscan.New(source),
    Filename: filename,
    LineNumber: 0,
    LineOffset: 0,
    ignoreSpaces: true,
    ignoreComments: true,
    nodes: nil,
    error: nil,
  }
}

type key struct {
  re string
  id int
}

func fixed_key(s string, n int) key {
  return key { regexp.QuoteMeta(s), n }
}

var keywords []key = []key {
  fixed_key("void",     VOID),
  fixed_key("char",     CHAR),
  fixed_key("short",    SHORT),
  fixed_key("int",      INT),
  fixed_key("long",     LONG),
  fixed_key("struct",   STRUCT),
  fixed_key("union",    UNION),
  fixed_key("enum",     ENUM),
  fixed_key("static",   STATIC),
  fixed_key("extern",   EXTERN),
  fixed_key("const",    CONST),
  fixed_key("signed",   SIGNED),
  fixed_key("unsigned", UNSIGNED),
  fixed_key("if",       IF),
  fixed_key("else",     ELSE),
  fixed_key("switch",   SWITCH),
  fixed_key("case",     CASE),
  fixed_key("default",  DEFAULT),
  fixed_key("while",    WHILE),
  fixed_key("do",       DO),
  fixed_key("for",      FOR),
  fixed_key("return",   RETURN),
  fixed_key("break",    BREAK),
  fixed_key("continue", CONTINUE),
  fixed_key("goto",     GOTO),
  fixed_key("typedef",  TYPEDEF),
  fixed_key("import",   IMPORT),
  fixed_key("sizeof",   SIZEOF),
}

var operators []key = []key {
  fixed_key("...",      DOTDOTDOT),
  fixed_key("<<=",      LSHIFTEQ),
  fixed_key(">>=",      RSHIFTEQ),
  fixed_key("!=",       NEQ),
  fixed_key("%=",       MODEQ),
  fixed_key("&&",       ANDAND),
  fixed_key("&=",       ANDEQ),
  fixed_key("*=",       MULEQ),
  fixed_key("++",       PLUSPLUS),
  fixed_key("+=",       PLUSEQ),
  fixed_key("--",       MINUSMINUS),
  fixed_key("-=",       MINUSEQ),
  fixed_key("->",       ARROW),
  fixed_key("/=",       DIVEQ),
  fixed_key("<<",       LSHIFT),
  fixed_key("<=",       LTEQ),
  fixed_key("==",       EQEQ),
  fixed_key(">=",       GTEQ),
  fixed_key(">>",       RSHIFT),
  fixed_key("^=",       XOREQ),
  fixed_key("|=",       OREQ),
  fixed_key("||",       OROR),
}

func (self *lex) GetToken() (t *token) {
  if self.scanner.IsEOS() {
    return nil
  }

  t = self.scanSpaces()
  if t != nil {
    if ! self.ignoreSpaces {
      return t
    } else {
      return self.GetToken()
    }
  }

  t = self.scanBlockComment()
  if t != nil {
    if ! self.ignoreComments {
      return t
    } else {
      return self.GetToken()
    }
  }
  t = self.scanLineComment()
  if t != nil {
    if ! self.ignoreComments {
      return t
    } else {
      return self.GetToken()
    }
  }

  t = self.scanKeyword()
  if t != nil {
    return t
  }

  t = self.scanIdentifier()
  if t != nil {
    return t
  }

  t = self.scanInteger()
  if t != nil {
    return t
  }

  t = self.scanCharacter()
  if t != nil {
    return t
  }

  t = self.scanString()
  if t != nil {
    return t
  }

  t = self.scanOperator()
  if t != nil {
    return t
  }

  panic(fmt.Errorf("lexer error: %s", self))
}

func (self *lex) consume(id int, literal string) (t *token) {
  t = &token {
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

func (self *lex) scanBlockComment() *token {
  s := self.scanner.Scan("/\\*")
  if s == "" {
    return nil
  }
  more := self.scanner.ScanUntil("\\*/")
  if more == "" {
    panic(fmt.Errorf("lexer error: %s", self))
  }
  return self.consume(BLOCK_COMMENT, s + more)
}

func (self *lex) scanLineComment() *token {
  s := self.scanner.Scan("//")
  if s == "" {
    return nil
  }
  more := self.scanner.ScanUntil("(\n|\r\n|\r)")
  if more == "" {
    panic(fmt.Errorf("lexer error: %s", self))
  }
  return self.consume(LINE_COMMENT, s + more)
}

func (self *lex) scanSpaces() *token {
  s := self.scanner.Scan("[ \t\n\r\f]+")
  if s == "" {
    return nil
  }
  return self.consume(SPACES, s)
}

func (self *lex) scanIdentifier() *token {
  s := self.scanner.Scan("[_A-Za-z][_0-9A-Za-z]*")
  if s == "" {
    return nil
  }
  return self.consume(IDENTIFIER, s)
}

func (self *lex) scanInteger() *token {
  s := self.scanner.Scan("([1-9][0-9]*U?L?|0[Xx][0-9A-Fa-f]+U?L?|0[0-7]*U?L?)")
  if s == "" {
    return nil
  }
  return self.consume(INTEGER, s)
}

func (self *lex) scanKeyword() *token {
  for i := range keywords {
    x := keywords[i]
    s := self.scanner.Scan(x.re)
    if s != "" {
      return self.consume(x.id, s)
    }
  }
  return nil
}

func (self *lex) scanCharacter() *token {
  s := self.scanner.Scan("'")
  if s == "" {
    return nil
  }
  // TODO: handle escape character properly
  more := self.scanner.ScanUntil("'")
  if more == "" {
    panic(fmt.Errorf("lexer error: %s", self))
  }
  return self.consume(CHARACTER, s + more)
}

func (self *lex) scanString() *token {
  s := self.scanner.Scan("\"")
  if s == "" {
    return nil
  }
  // TODO: handle escape character properly
  more := self.scanner.ScanUntil("\"")
  if more == "" {
    panic(fmt.Errorf("lexer error: %s", self))
  }
  return self.consume(STRING, s + more)
}

func (self *lex) scanOperator() *token {
  for i := range operators {
    x := operators[i]
    s := self.scanner.Scan(x.re)
    if s != "" {
      return self.consume(x.id, s)
    }
  }

  // use next rune as an operator if available
  s := self.scanner.Scan(".")
  if s != "" {
    r, _ := utf8.DecodeRuneInString(s)
    return self.consume(int(r), s)
  }
  return nil
}
