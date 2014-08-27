package parser

import (
  "fmt"
  "regexp"
  "strings"
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

type Token struct {
  Id int
  Literal string
  Filename string
  LineNumber int
  LineOffset int
}

func (self Token) String() string {
  t := "UNKNOWN"
  switch self.Id {
    case SPACES:        t = "SPACES"
    case BLOCK_COMMENT: t = "BLOCK_COMMENT"
    case LINE_COMMENT:  t = "LINE_COMMENT"
    case VOID:          t = "VOID"
    case CHAR:          t = "CHAR"
    case SHORT:         t = "SHORT"
    case INT:           t = "INT"
    case LONG:          t = "LONG"
    case STRUCT:        t = "STRUCT"
    case UNION:         t = "UNION"
    case ENUM:          t = "ENUM"
    case STATIC:        t = "STATIC"
    case EXTERN:        t = "EXTERN"
    case CONST:         t = "CONST"
    case SIGNED:        t = "SIGNED"
    case UNSIGNED:      t = "UNSIGNED"
    case IF:            t = "IF"
    case ELSE:          t = "ELSE"
    case SWITCH:        t = "SWITCH"
    case CASE:          t = "CASE"
    case DEFAULT:       t = "DEFAULT"
    case WHILE:         t = "WHILE"
    case DO:            t = "DO"
    case FOR:           t = "FOR"
    case RETURN:        t = "RETURN"
    case BREAK:         t = "BREAK"
    case CONTINUE:      t = "CONTINUE"
    case GOTO:          t = "GOTO"
    case TYPEDEF:       t = "TYPEDEF"
    case IMPORT:        t = "IMPORT"
    case SIZEOF:        t = "SIZEOF"
    case IDENTIFIER:    t = "IDENTIFIER"
    case INTEGER:       t = "INTEGER"
    case CHARACTER:     t = "CHARACTER"
    case STRING:        t = "STRING"
    case OPERATOR:      t = "OPERATOR"
  }
  return fmt.Sprintf("#<Token:%s %s:%d,%d %q>", t, self.Filename, self.LineNumber, self.LineOffset, self.Literal)

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
    if ! self.ignoreSpaces {
      return t
    } else {
      t = nil // ignore token
    }
  }

  t = self.readBlockComment()
  if t != nil {
    if ! self.ignoreComments {
      return t
    } else {
      t = nil // ignore token
    }
  }
  t = self.readLineComment()
  if t != nil {
    if ! self.ignoreComments {
      return t
    } else {
      t = nil // ignore token
    }
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

  return self.GetToken()
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
    panic(fmt.Errorf("lexer error: %s", self))
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
    panic(fmt.Errorf("lexer error: %s", self))
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
    panic(fmt.Errorf("lexer error: %s", self))
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
    panic(fmt.Errorf("lexer error: %s", self))
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
