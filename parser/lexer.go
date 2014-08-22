package parser

import (
  "fmt"
  "regexp"
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
  filename string
  scanner strscan.StringScanner
}

type Token struct {
  Id int
  Literal string
}

func (t *Token) ToString() string {
  name, ok := id2name[t.Id]
  if ok {
    return fmt.Sprintf("<Token:%s %q>", name, t.Literal)
  } else {
    return fmt.Sprintf("<Token:%d %q>", t.Id, t.Literal)
  }
}

func NewLexer(filename string, source string) *Lexer {
  return &Lexer {
    filename: filename,
    scanner: strscan.NewStringScanner(source),
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

func (self *Lexer) readBlockComment() *Token {
  s := self.scanner.Scan("/\\*")
  if s == "" {
    return nil
  }
  for {
    more := self.scanner.ScanUntil("\\*/")
    s += more
    if more != "" {
      break
    }
  }
  return &Token { BLOCK_COMMENT, s }
}

func (self *Lexer) readLineComment() *Token {
  s := self.scanner.Scan("//")
  if s == "" {
    return nil
  }
  for {
    more := self.scanner.ScanUntil("(\n|\r\n|\r)")
    s += more
    if more != "" {
      break
    }
  }
  return &Token { LINE_COMMENT, s }
}

func (self *Lexer) readSpaces() *Token {
  s := self.scanner.Scan("[ \t\n\r\f]+")
  if s == "" {
    return nil
  }
  return &Token { SPACES, s }
}

func (self *Lexer) readIdentifier() *Token {
  s := self.scanner.Scan("[_A-Za-z][_0-9A-Za-z]*")
  if s == "" {
    return nil
  }
  return &Token { IDENTIFIER, s }
}

func (self *Lexer) readInteger() *Token {
  s := self.scanner.Scan("([1-9][0-9]*U?L?|0[Xx][0-9A-Fa-f]+U?L?|0[0-7]*U?L?)")
  if s == "" {
    return nil
  }
  return &Token { INTEGER, s }
}

func (self *Lexer) readKeyword() *Token {
  for keyword, id := range keywords {
    s := self.scanner.Scan(regexp.QuoteMeta(keyword))
    if s != "" {
      return &Token { id, s }
    }
  }
  return nil
}

func (self *Lexer) readCharacter() *Token {
  s := self.scanner.Scan("'")
  if s == "" {
    return nil
  }
  for {
    // TODO: handle escape character properly
    more := self.scanner.ScanUntil("'")
    s += more
    if more != "" {
      break
    }
  }
  return &Token { CHARACTER, s }
}

func (self *Lexer) readString() *Token {
  s := self.scanner.Scan("\"")
  if s == "" {
    return nil
  }
  for {
    // TODO: handle escape character properly
    more := self.scanner.ScanUntil("\"")
    s += more
    if more != "" {
      break
    }
  }
  return &Token { STRING, s }
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
      return &Token { OPERATOR, s }
    }
  }
  return nil
}
