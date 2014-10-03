package parser

import (
  "fmt"
  "regexp"
  "strings"
  "unicode/utf8"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_strscan "bitbucket.org/yyuu/bs/strscan"
)

type lexer struct {
  scanner *bs_strscan.StringScanner
  sourceName string
  lineNumber int
  lineOffset int
  eof bool
  knownTypedefs []string
  libraryLoader *libraryLoader
  ast *bs_ast.AST
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
}

func (self lexer) String() string {
  return fmt.Sprintf("%s %q", self.pos(), self.scanner.Peek(16))
}

func (self lexer) pos() bs_core.Location {
  return bs_core.NewLocation(self.sourceName, self.lineNumber, self.lineOffset)
}

func (self lexer) debugPos() bs_core.Location {
  // FIXME: inefficient
  s := self.scanner.String[0:self.scanner.Pos()]
  lineNumber := 1 + strings.Count(s, "\n")
  if 1 < lineNumber {
    lineOffset := len(s[1 + strings.LastIndex(s, "\n"):])
    return bs_core.NewLocation(self.sourceName, lineNumber, lineOffset)
  } else {
    return bs_core.NewLocation(self.sourceName, 1, len(s))
  }
}

func newLexer(filename string, source string, loader *libraryLoader, errorHandler *bs_core.ErrorHandler, options *bs_core.Options) *lexer {
  return &lexer {
    scanner: bs_strscan.New(source),
    sourceName: filename,
    lineNumber: 1,
    lineOffset: 1,
    eof: false,
    knownTypedefs: []string { },
    libraryLoader: loader,
    ast: nil,
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

var keywords = []key {
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

var operators = []key {
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
      return &token { EOF, "", self.pos() }, nil
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

  return nil, fmt.Errorf("%s lexer error", self.pos())
}

func (self *lexer) consume(id int, raw, literal string) (t *token) {
  t = &token {
    id: id,
    literal: literal,
    location: self.pos(),
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
  raw := self.scanner.Scan("/\\*")
  if raw == "" {
    return nil, nil
  }
  val := self.scanner.ScanUntil("\\*/")
  raw += val
  if val == "" {
    return nil, fmt.Errorf("%s lexer error", self.pos())
  }
  return self.consume(BLOCK_COMMENT, raw, raw), nil
}

func (self *lexer) scanLineComment() (*token, error) {
  raw := self.scanner.Scan("//")
  if raw == "" {
    return nil, nil
  }
  val := self.scanner.ScanUntil("(\n|\r\n|\r)")
  raw += val
  if val == "" {
    return nil, fmt.Errorf("%s lexer error", self.pos())
  }
  return self.consume(LINE_COMMENT, raw, raw), nil
}

func (self *lexer) scanSpaces() (*token, error) {
  raw := self.scanner.Scan("[ \t\n\r\f]+")
  if raw == "" {
    return nil, nil
  }
  return self.consume(SPACES, raw, raw), nil
}

func (self *lexer) scanIdentifier() (*token, error) {
  raw := self.scanner.Scan("[_A-Za-z][_0-9A-Za-z]*")
  if raw == "" {
    return nil, nil
  }
  for i := range self.knownTypedefs {
    if self.knownTypedefs[i] == raw {
      return self.consume(TYPENAME, raw, raw), nil
    }
  }
  return self.consume(IDENTIFIER, raw, raw), nil
}

func (self *lexer) scanInteger() (*token, error) {
  var raw string
  var val int64
  hex := self.scanner.Scan("0[Xx][0-9A-Fa-f]+")
  raw += hex
  if hex != "" {
    // hexadecimal
    _, err := fmt.Sscanf(hex[2:], "%x", &val)
    if err != nil {
      return nil, err
    }
  } else {
    oct := self.scanner.Scan("0[0-7]+")
    raw += oct
    if oct != "" {
      // octal
      _, err := fmt.Sscanf(oct[1:], "%o", &val)
      if err != nil {
        return nil, err
      }
    } else {
      dec := self.scanner.Scan("[0-9]+")
      raw += dec
      if dec != "" {
        // decimal
        _, err := fmt.Sscanf(dec, "%d", &val)
        if err != nil {
          return nil, err
        }
      } else {
        return nil, nil
      }
    }
  }
  ul := self.scanner.Scan("U?L?")
  raw += ul
  return self.consume(INTEGER, raw, fmt.Sprintf("%d%s", val, ul)), nil
}

func (self *lexer) scanKeyword() (*token, error) {
  for i := range keywords {
    x := keywords[i]
    raw := self.scanner.Scan(x.re)
    if raw != "" {
      return self.consume(x.id, raw, raw), nil
    }
  }
  return nil, nil
}

func (self *lexer) scanCharacter() (*token, error) {
  var raw string
  q1 := self.scanner.Scan("'")
  raw += q1
  if q1 == "" {
    return nil, nil
  }
  var val int
  r1 := self.scanner.Scan(".")
  raw += r1
  switch r1 {
    case "\\":
      r2 := self.scanner.Scan(".")
      raw += r2
      switch r2 {
        case "a":  val = '\a'
        case "b":  val = '\b'
        case "f":  val = '\f'
        case "n":  val = '\n'
        case "r":  val = '\r'
        case "t":  val = '\t'
        case "u": {
          hex := self.scanner.Scan("[0-9A-Fa-f]+")
          raw += hex
          if hex == "" {
            return nil, fmt.Errorf("%s invalid unicode code point", self.debugPos())
          }
          _, err := fmt.Sscanf(hex, "%x", &val)
          if err != nil {
            return nil, err
          }
        }
        case "v":  val = '\v'
        case "'":  val = '\''
        case "\\": val = '\\'
        default: {
          return nil, fmt.Errorf("%s unknown escape character: %q", self.debugPos(), r2)
        }
      }
    default: {
      r, _ := utf8.DecodeRuneInString(r1)
      val = int(r)
    }
  }
  q2 := self.scanner.Scan("'")
  raw += q2
  if q2 == "" {
    return nil, fmt.Errorf("%s invalid character literal", self.debugPos())
  }
  return self.consume(CHARACTER, raw, fmt.Sprintf("%d", val)), nil
}

func (self *lexer) scanString() (*token, error) {
  var raw string
  q1 := self.scanner.Scan("\"")
  raw += q1
  if q1 == "" {
    return nil, nil
  }
  var val string
  for {
    if self.scanner.IsEOS() {
      return nil, fmt.Errorf("%s EOL while scanning string literal", self.debugPos())
    }
    r1 := self.scanner.Scan(".")
    raw += r1
    switch r1 {
      case "\"": {
        return self.consume(STRING, raw, val), nil
      }
      case "\\": {
        r2 := self.scanner.Scan(".")
        raw += r2
        switch r2 {
          case "a":  val += "\a"
          case "b":  val += "\b"
          case "f":  val += "\f"
          case "n":  val += "\n"
          case "r":  val += "\r"
          case "t":  val += "\t"
          case "u": {
            hex := self.scanner.Scan("[0-9A-Fa-f]+")
            raw += hex
            if hex == "" {
              return nil, fmt.Errorf("%s invalid unicode code point", self.debugPos())
            }
            var num int
            _, err := fmt.Sscanf(hex, "%x", &num)
            if err != nil {
              return nil, err
            }
            val += string(rune(num))
          }
          case "v":  val += "\v"
          case "\"": val += "\""
          case "\\": val += "\\"
          default: {
            return nil, fmt.Errorf("%s unknown escape character: %q", self.debugPos(), r2)
          }
        }
      }
      default: {
        val += r1
      }
    }
  }
  return nil, fmt.Errorf("%s lexer error", self.pos())
}

func (self *lexer) scanOperator() (*token, error) {
  for i := range operators {
    x := operators[i]
    raw := self.scanner.Scan(x.re)
    if raw != "" {
      return self.consume(x.id, raw, raw), nil
    }
  }

  // use next rune as an operator if available
  raw := self.scanner.Scan(".")
  if raw != "" {
    val, _ := utf8.DecodeRuneInString(raw)
    return self.consume(int(val), raw, string(val)), nil
  }
  return nil, nil
}

func (self *lexer) loadLibrary(name string) *bs_ast.Declaration {
  return self.libraryLoader.loadLibrary(name)
}
