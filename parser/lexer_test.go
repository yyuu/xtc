package parser

import (
  "testing"
)

func assertToken(t *testing.T, tok *token, id int, literal string) {
  if tok == nil {
    t.Error("insufficient token")
  } else {
    if tok.Id != id {
      t.Errorf("invalid token id: %s: expected %d, got %d", tok, id, tok.Id)
    }
    if tok.Literal != literal {
      t.Errorf("invalid token literal: %s: expected %q, got %q", tok, literal, tok.Literal)
    }
  }
}

func assertTokenNull(t *testing.T, tok *token) {
  if tok != nil {
    t.Errorf("surplus token: %s", tok)
  }
}

func TestEmpty(t *testing.T) {
  lex := NewLexer("-", "")
  assertTokenNull(t, lex.GetToken())
}

func TestSpaces(t *testing.T) {
  lex := NewLexer("-", "\tfoo\n\t\tbar\n\n")
//assertToken(t, lex.GetToken(), SPACES, "\t")
  assertToken(t, lex.GetToken(), IDENTIFIER, "foo")
//assertToken(t, lex.GetToken(), SPACES, "\n\t\t")
  assertToken(t, lex.GetToken(), IDENTIFIER, "bar")
//assertToken(t, lex.GetToken(), SPACES, "\n\n")
  assertTokenNull(t, lex.GetToken())
}

func TestBlockComment1(t *testing.T) {
  lex := NewLexer("-", "/* foo */\n")
//assertToken(t, lex.GetToken(), BLOCK_COMMENT, "/* foo */")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestBlockComment2(t *testing.T) {
  lex := NewLexer("-", "foo\n/* bar\n   baz\n */\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "foo")
//assertToken(t, lex.GetToken(), SPACES, "\n")
//assertToken(t, lex.GetToken(), BLOCK_COMMENT, "/* bar\n   baz\n */")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestLineComment(t *testing.T) {
  lex := NewLexer("-", "foo\n// bar\nbaz\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "foo")
//assertToken(t, lex.GetToken(), SPACES, "\n")
//assertToken(t, lex.GetToken(), LINE_COMMENT, "// bar\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "baz")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestKeyword1(t *testing.T) {
  lex := NewLexer("-", "unsigned int foo;\n")
  assertToken(t, lex.GetToken(), UNSIGNED, "unsigned")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), INT, "int")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), IDENTIFIER, "foo")
  assertToken(t, lex.GetToken(), ';', ";")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestKeyword2(t *testing.T) {
  lex := NewLexer("-", "\n\nif ( foo ) {\n  bar;\n}\nbaz;\n")
//assertToken(t, lex.GetToken(), SPACES, "\n\n")
  assertToken(t, lex.GetToken(), IF, "if")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), '(', "(")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), IDENTIFIER, "foo")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), ')', ")")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), '{', "{")
//assertToken(t, lex.GetToken(), SPACES, "\n  ")
  assertToken(t, lex.GetToken(), IDENTIFIER, "bar")
  assertToken(t, lex.GetToken(), ';', ";")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), '}', "}")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "baz")
  assertToken(t, lex.GetToken(), ';', ";")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestIdentifier1(t *testing.T) {
  lex := NewLexer("-", "foo\nbar\nbaz\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "foo")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "bar")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "baz")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestIdentifier2(t *testing.T) {
  lex := NewLexer("-", "f00\n64r\nb42\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "f00")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "64")
  assertToken(t, lex.GetToken(), IDENTIFIER, "r")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "b42")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestInteger1(t *testing.T) {
  lex := NewLexer("-", "1\n23U\n456L\n7890UL\n")
  assertToken(t, lex.GetToken(), INTEGER, "1")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "23U")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "456L")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "7890UL")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestInteger2(t *testing.T) {
  lex := NewLexer("-", "0xf00\n0x64U\n0X642L\n0xc4febabe\n0XC0FFEEUL\n")
  assertToken(t, lex.GetToken(), INTEGER, "0xf00")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "0x64U")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "0X642L")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "0xc4febabe")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "0XC0FFEEUL")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestInteger3(t *testing.T) {
  lex := NewLexer("-", "0\n012U\n034L\n056UL\n")
  assertToken(t, lex.GetToken(), INTEGER, "0")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "012U")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "034L")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertToken(t, lex.GetToken(), INTEGER, "056UL")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}

func TestCharacter1(t *testing.T) {
  lex := NewLexer("-", "{'f', 'o', 'o'}")
  assertToken(t, lex.GetToken(), '{', "{")
  assertToken(t, lex.GetToken(), CHARACTER, "'f'")
  assertToken(t, lex.GetToken(), ',', ",")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), CHARACTER, "'o'")
  assertToken(t, lex.GetToken(), ',', ",")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), CHARACTER, "'o'")
  assertToken(t, lex.GetToken(), '}', "}")
  assertTokenNull(t, lex.GetToken())
}

func TestCharacter2(t *testing.T) {
  lex := NewLexer("-", "'\x20'")
  assertToken(t, lex.GetToken(), CHARACTER, "'\x20'")
  assertTokenNull(t, lex.GetToken())
}

func TestString1(t *testing.T) {
  lex := NewLexer("-", "\"foo, bar, baz\"")
  assertToken(t, lex.GetToken(), STRING, "\"foo, bar, baz\"")
  assertTokenNull(t, lex.GetToken())
}

func TestOperator1(t *testing.T) {
  lex := NewLexer("-", "+++....<<<<===&=&&")
  assertToken(t, lex.GetToken(), PLUSPLUS, "++")
  assertToken(t, lex.GetToken(), '+', "+")
  assertToken(t, lex.GetToken(), DOTDOTDOT, "...")
  assertToken(t, lex.GetToken(), '.', ".")
  assertToken(t, lex.GetToken(), SHIFTLEFT, "<<")
  assertToken(t, lex.GetToken(), SHIFTLEFTEQ, "<<=")
  assertToken(t, lex.GetToken(), EQEQ, "==")
  assertToken(t, lex.GetToken(), ANDEQ, "&=")
  assertToken(t, lex.GetToken(), ANDAND, "&&")
  assertTokenNull(t, lex.GetToken())
}

func TestOperator2(t *testing.T) {
  lex := NewLexer("-", "foo ? bar : baz;\n")
  assertToken(t, lex.GetToken(), IDENTIFIER, "foo")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), '?', "?")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), IDENTIFIER, "bar")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), ':', ":")
//assertToken(t, lex.GetToken(), SPACES, " ")
  assertToken(t, lex.GetToken(), IDENTIFIER, "baz")
  assertToken(t, lex.GetToken(), ';', ";")
//assertToken(t, lex.GetToken(), SPACES, "\n")
  assertTokenNull(t, lex.GetToken())
}
