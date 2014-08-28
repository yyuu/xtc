package parser

import (
  "fmt"
  "regexp"
  "strings"
  "unicode/utf8"
  "bitbucket.org/yyuu/bs/strscan"
)

type Lexer struct {
  scanner strscan.StringScanner
  Filename string
  LineNumber int
  LineOffset int
  ignoreSpaces bool
  ignoreComments bool
}

func (self Lexer) String() string {
  source := fmt.Sprintf("%s...", self.scanner.Peek(32))
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

func NewLexer(filename string, source string) *Lexer {
  return &Lexer {
    scanner: strscan.NewStringScanner(source),
    Filename: filename,
    LineNumber: 0,
    LineOffset: 0,
    ignoreSpaces: true,
    ignoreComments: true,
  }
}

var keywords map[string]int = map[string]int {
  regexp.QuoteMeta("void"):     VOID,
  regexp.QuoteMeta("char"):     CHAR,
  regexp.QuoteMeta("short"):    SHORT,
  regexp.QuoteMeta("int"):      INT,
  regexp.QuoteMeta("long"):     LONG,
  regexp.QuoteMeta("struct"):   STRUCT,
  regexp.QuoteMeta("union"):    UNION,
  regexp.QuoteMeta("enum"):     ENUM,
  regexp.QuoteMeta("static"):   STATIC,
  regexp.QuoteMeta("extern"):   EXTERN,
  regexp.QuoteMeta("const"):    CONST,
  regexp.QuoteMeta("signed"):   SIGNED,
  regexp.QuoteMeta("unsigned"): UNSIGNED,
  regexp.QuoteMeta("if"):       IF,
  regexp.QuoteMeta("else"):     ELSE,
  regexp.QuoteMeta("switch"):   SWITCH,
  regexp.QuoteMeta("case"):     CASE,
  regexp.QuoteMeta("default"):  DEFAULT,
  regexp.QuoteMeta("while"):    WHILE,
  regexp.QuoteMeta("do"):       DO,
  regexp.QuoteMeta("for"):      FOR,
  regexp.QuoteMeta("return"):   RETURN,
  regexp.QuoteMeta("break"):    BREAK,
  regexp.QuoteMeta("continue"): CONTINUE,
  regexp.QuoteMeta("goto"):     GOTO,
  regexp.QuoteMeta("typedef"):  TYPEDEF,
  regexp.QuoteMeta("import"):   IMPORT,
  regexp.QuoteMeta("sizeof"):   SIZEOF,
}

var operators map[string]int = map[string]int {
  regexp.QuoteMeta("..."):      DOTDOTDOT,
  regexp.QuoteMeta("<<="):      SHIFTLEFTEQ,
  regexp.QuoteMeta(">>="):      SHIFTRIGHTEQ,
  regexp.QuoteMeta("!="):       NEQ,
  regexp.QuoteMeta("%="):       MODEQ,
  regexp.QuoteMeta("&&"):       ANDAND,
  regexp.QuoteMeta("&="):       ANDEQ,
  regexp.QuoteMeta("*="):       MULEQ,
  regexp.QuoteMeta("++"):       PLUSPLUS,
  regexp.QuoteMeta("+="):       PLUSEQ,
  regexp.QuoteMeta("--"):       MINUSMINUS,
  regexp.QuoteMeta("-="):       MINUSEQ,
  regexp.QuoteMeta("->"):       MINUSGT,
  regexp.QuoteMeta("/="):       DIVEQ,
  regexp.QuoteMeta("<<"):       SHIFTLEFT,
  regexp.QuoteMeta("<="):       LE,
  regexp.QuoteMeta("=="):       EQEQ,
  regexp.QuoteMeta(">="):       GE,
  regexp.QuoteMeta(">>"):       SHIFTRIGHT,
  regexp.QuoteMeta("^="):       XOREQ,
  regexp.QuoteMeta("|="):       OREQ,
  regexp.QuoteMeta("||"):       OROR,
}

func (self *Lexer) GetToken() (t *token) {
  if self.scanner.IsEOS() {
    return nil
  }

  t = self.scanSpaces()
  if t != nil {
    if ! self.ignoreSpaces {
      return t
    } else {
      t = nil // ignore token
    }
  }

  t = self.scanBlockComment()
  if t != nil {
    if ! self.ignoreComments {
      return t
    } else {
      t = nil // ignore token
    }
  }
  t = self.scanLineComment()
  if t != nil {
    if ! self.ignoreComments {
      return t
    } else {
      t = nil // ignore token
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

  return self.GetToken()
}

func (self *Lexer) consume(id int, literal string) (t *token) {
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

func (self *Lexer) scanBlockComment() *token {
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

func (self *Lexer) scanLineComment() *token {
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

func (self *Lexer) scanSpaces() *token {
  s := self.scanner.Scan("[ \t\n\r\f]+")
  if s == "" {
    return nil
  }
  return self.consume(SPACES, s)
}

func (self *Lexer) scanIdentifier() *token {
  s := self.scanner.Scan("[_A-Za-z][_0-9A-Za-z]*")
  if s == "" {
    return nil
  }
  return self.consume(IDENTIFIER, s)
}

func (self *Lexer) scanInteger() *token {
  s := self.scanner.Scan("([1-9][0-9]*U?L?|0[Xx][0-9A-Fa-f]+U?L?|0[0-7]*U?L?)")
  if s == "" {
    return nil
  }
  return self.consume(INTEGER, s)
}

func (self *Lexer) scanKeyword() *token {
  for r, id := range keywords {
    s := self.scanner.Scan(r)
    if s != "" {
      return self.consume(id, s)
    }
  }
  return nil
}

func (self *Lexer) scanCharacter() *token {
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

func (self *Lexer) scanString() *token {
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

func (self *Lexer) scanOperator() *token {
  for r, id := range operators {
    s := self.scanner.Scan(r)
    if s != "" {
      return self.consume(id, s)
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
