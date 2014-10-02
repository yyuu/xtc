package parser

import (
  "fmt"
  "regexp"
  "strings"
  "unicode/utf8"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/strscan"
)

type lexer struct {
  scanner *strscan.StringScanner
  sourceName string
  lineNumber int
  lineOffset int
  eof bool
  knownTypedefs []string
  libraryLoader *libraryLoader
  ast *ast.AST
  error error
  errorHandler *core.ErrorHandler
  options *core.Options
}

func (self lexer) String() string {
  location := core.NewLocation(self.sourceName, self.lineNumber, self.lineOffset)
  source := fmt.Sprintf("%s", self.scanner.Peek(16))
  return fmt.Sprintf("%s: %q", location, source)
}

func newLexer(filename string, source string, loader *libraryLoader, errorHandler *core.ErrorHandler, options *core.Options) *lexer {
  return &lexer {
    scanner: strscan.New(source),
    sourceName: filename,
    lineNumber: 1,
    lineOffset: 1,
    eof: false,
    knownTypedefs: []string { },
    libraryLoader: loader,
    ast: nil,
    error: nil,
    errorHandler: errorHandler,
    options: options,
  }
}

type key struct {
  re string
  id int
}

func fixed_word(s string, n int) key {
  return key { regexp.QuoteMeta(s) + "\\b", n }
}

func fixed_sign(s string, n int) key {
  return key { regexp.QuoteMeta(s), n }
}

var keywords []key = []key {
  fixed_word("void",     VOID),
  fixed_word("char",     CHAR),
  fixed_word("short",    SHORT),
  fixed_word("int",      INT),
  fixed_word("long",     LONG),
  fixed_word("struct",   STRUCT),
  fixed_word("union",    UNION),
  fixed_word("enum",     ENUM),
  fixed_word("static",   STATIC),
  fixed_word("extern",   EXTERN),
  fixed_word("const",    CONST),
  fixed_word("signed",   SIGNED),
  fixed_word("unsigned", UNSIGNED),
  fixed_word("if",       IF),
  fixed_word("else",     ELSE),
  fixed_word("switch",   SWITCH),
  fixed_word("case",     CASE),
  fixed_word("default",  DEFAULT),
  fixed_word("while",    WHILE),
  fixed_word("do",       DO),
  fixed_word("for",      FOR),
  fixed_word("return",   RETURN),
  fixed_word("break",    BREAK),
  fixed_word("continue", CONTINUE),
  fixed_word("goto",     GOTO),
  fixed_word("typedef",  TYPEDEF),
  fixed_word("import",   IMPORT),
  fixed_word("sizeof",   SIZEOF),
}

var operators []key = []key {
  fixed_sign("...",      DOTDOTDOT),
  fixed_sign("<<=",      LSHIFTEQ),
  fixed_sign(">>=",      RSHIFTEQ),
  fixed_sign("!=",       NEQ),
  fixed_sign("%=",       MODEQ),
  fixed_sign("&&",       ANDAND),
  fixed_sign("&=",       ANDEQ),
  fixed_sign("*=",       MULEQ),
  fixed_sign("++",       PLUSPLUS),
  fixed_sign("+=",       PLUSEQ),
  fixed_sign("--",       MINUSMINUS),
  fixed_sign("-=",       MINUSEQ),
  fixed_sign("->",       ARROW),
  fixed_sign("/=",       DIVEQ),
  fixed_sign("<<",       LSHIFT),
  fixed_sign("<=",       LTEQ),
  fixed_sign("==",       EQEQ),
  fixed_sign(">=",       GTEQ),
  fixed_sign(">>",       RSHIFT),
  fixed_sign("^=",       XOREQ),
  fixed_sign("|=",       OREQ),
  fixed_sign("||",       OROR),
}

func (self *lexer) getNextToken() (t *token, err error) {
  if self.scanner.IsEOS() {
    if ! self.eof {
      self.eof = true
      return &token { EOF, "", core.NewLocation(self.sourceName, self.lineNumber, self.lineOffset) }, nil
    }
    return nil, nil
  }

  t, err = self.scanSpaces()
  if err != nil { return nil, err }
  // ignore spaces
  if t != nil {
    return self.getNextToken()
  }

  t, err = self.scanBlockComment()
  if err != nil { return nil, err }
  // ignore comments
  if t != nil {
    return self.getNextToken()
  }

  t, err = self.scanLineComment()
  if err != nil { return nil, err }
  // ignore comments
  if t != nil {
    return self.getNextToken()
  }

  t, err = self.scanKeyword()
  if err != nil { return nil, err }
  if t != nil { return t, nil }

  t, err = self.scanIdentifier()
  if err != nil { return nil, err }
  if t != nil { return t, nil }

  t, err = self.scanInteger()
  if err != nil { return nil, err }
  if t != nil { return t, nil }

  t, err = self.scanCharacter()
  if err != nil { return nil, err }
  if t != nil { return t, nil }

  t, err = self.scanString()
  if err != nil { return nil, err }
  if t != nil { return t, nil }

  t, err = self.scanOperator()
  if err != nil { return nil, err }
  if t != nil { return t, nil }

  return nil, fmt.Errorf("lexer error: %s", self)
}

func (self *lexer) consume(id int, raw, literal string) (t *token) {
  t = &token {
    id: id,
    literal: literal,
    location: core.NewLocation(self.sourceName, self.lineNumber, self.lineOffset),
  }

  self.lineNumber += strings.Count(raw, "\n")
  i := strings.LastIndex(raw, "\n")
  if i < 0 {
    self.lineOffset += len(raw)
  } else {
    self.lineOffset = len(raw[i:])
  }

  return t
}

func (self *lexer) scanBlockComment() (*token, error) {
  s := self.scanner.Scan("/\\*")
  if s == "" {
    return nil, nil
  }
  more := self.scanner.ScanUntil("\\*/")
  if more == "" {
    return nil, fmt.Errorf("lexer error: %s", self)
  }
  return self.consume(BLOCK_COMMENT, s+more, s+more), nil
}

func (self *lexer) scanLineComment() (*token, error) {
  s := self.scanner.Scan("//")
  if s == "" {
    return nil, nil
  }
  more := self.scanner.ScanUntil("(\n|\r\n|\r)")
  if more == "" {
    return nil, fmt.Errorf("lexer error: %s", self)
  }
  return self.consume(LINE_COMMENT, s+more, s+more), nil
}

func (self *lexer) scanSpaces() (*token, error) {
  s := self.scanner.Scan("[ \t\n\r\f]+")
  if s == "" {
    return nil, nil
  }
  return self.consume(SPACES, s, s), nil
}

func (self *lexer) scanIdentifier() (*token, error) {
  s := self.scanner.Scan("[_A-Za-z][_0-9A-Za-z]*")
  if s == "" {
    return nil, nil
  }
  for i := range self.knownTypedefs {
    if self.knownTypedefs[i] == s {
      return self.consume(TYPENAME, s, s), nil
    }
  }
  return self.consume(IDENTIFIER, s, s), nil
}

func (self *lexer) scanInteger() (*token, error) {
  s := self.scanner.Scan("([1-9][0-9]*U?L?|0[Xx][0-9A-Fa-f]+U?L?|0[0-7]*U?L?)")
  if s == "" {
    return nil, nil
  }
  return self.consume(INTEGER, s, s), nil
}

func (self *lexer) scanKeyword() (*token, error) {
  for i := range keywords {
    x := keywords[i]
    s := self.scanner.Scan(x.re)
    if s != "" {
      return self.consume(x.id, s, s), nil
    }
  }
  return nil, nil
}

func (self *lexer) scanCharacter() (*token, error) {
  s := self.scanner.Scan("'")
  if s == "" {
    return nil, nil
  }
  e := self.scanner.Scan("\\\\([abfnrtv]|u[0-9]+|'|\\\\)'")
  if e != "" {
    return self.consume(CHARACTER, s+e, s+e), nil
  }
  r := self.scanner.Scan(".'")
  if r != "" {
    return self.consume(CHARACTER, s+r, s+r), nil
  }
  return nil, fmt.Errorf("lexer error: %s", self)
}

func (self *lexer) scanString() (*token, error) {
  s := self.scanner.Scan("\"")
  if s == "" {
    return nil, nil
  }
  var more string
  for {
    if self.scanner.IsEOS() {
      return nil, fmt.Errorf("EOL while scanning string literal")
    }
    r := self.scanner.Scan(".")
    switch r {
      case "\"": {
        return self.consume(STRING, s+more+r, more), nil
      }
      case "\\": {
        e := self.scanner.Scan(".")
        switch e {
          case "a":  more += "\a"
          case "b":  more += "\b"
          case "f":  more += "\f"
          case "n":  more += "\n"
          case "r":  more += "\r"
          case "t":  more += "\t"
          case "v":  more += "\v"
          case "\"": more += "\""
          case "\\": more += "\\"
          default: {
            return nil, fmt.Errorf("unknown escape character: %q", e)
          }
        }
      }
      default: {
        more += r
      }
    }
  }
  return nil, fmt.Errorf("lexer error: %s", self)
}

func (self *lexer) scanOperator() (*token, error) {
  for i := range operators {
    x := operators[i]
    s := self.scanner.Scan(x.re)
    if s != "" {
      return self.consume(x.id, s, s), nil
    }
  }

  // use next rune as an operator if available
  s := self.scanner.Scan(".")
  if s != "" {
    r, _ := utf8.DecodeRuneInString(s)
    return self.consume(int(r), s, s), nil
  }
  return nil, nil
}
