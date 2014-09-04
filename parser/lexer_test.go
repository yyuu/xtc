package parser

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func assertToken(t *testing.T, tok *token, id int, literal string) {
  xt.AssertNotNil(t, "insufficient token", tok)
  xt.AssertEquals(t, "invalid token id", tok.id, id)
  xt.AssertEquals(t, "invalid token literal", tok.literal, literal)
}

func assertTokenNull(t *testing.T, tok *token) {
  xt.AssertNotNil(t, "surplus token", tok)
}

func TestEmpty(t *testing.T) {
  lex := lexer("test.txt", "")
  assertTokenNull(t, lex.getToken())
}

func TestSpaces(t *testing.T) {
  lex := lexer("test.txt", "\tfoo\n\t\tbar\n\n")
//assertToken(t, lex.getToken(), SPACES, "\t")
  assertToken(t, lex.getToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getToken(), SPACES, "\n\t\t")
  assertToken(t, lex.getToken(), IDENTIFIER, "bar")
//assertToken(t, lex.getToken(), SPACES, "\n\n")
  assertTokenNull(t, lex.getToken())
}

func TestBlockComment1(t *testing.T) {
  lex := lexer("test.txt", "/* foo */\n")
//assertToken(t, lex.getToken(), BLOCK_COMMENT, "/* foo */")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestBlockComment2(t *testing.T) {
  lex := lexer("test.txt", "foo\n/* bar\n   baz\n */\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getToken(), SPACES, "\n")
//assertToken(t, lex.getToken(), BLOCK_COMMENT, "/* bar\n   baz\n */")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestLineComment(t *testing.T) {
  lex := lexer("test.txt", "foo\n// bar\nbaz\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getToken(), SPACES, "\n")
//assertToken(t, lex.getToken(), LINE_COMMENT, "// bar\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "baz")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestKeyword1(t *testing.T) {
  lex := lexer("test.txt", "unsigned int foo;\n")
  assertToken(t, lex.getToken(), UNSIGNED, "unsigned")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), INT, "int")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), IDENTIFIER, "foo")
  assertToken(t, lex.getToken(), ';', ";")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestKeyword2(t *testing.T) {
  lex := lexer("test.txt", "\n\nif ( foo ) {\n  bar;\n}\nbaz;\n")
//assertToken(t, lex.getToken(), SPACES, "\n\n")
  assertToken(t, lex.getToken(), IF, "if")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), '(', "(")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), ')', ")")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), '{', "{")
//assertToken(t, lex.getToken(), SPACES, "\n  ")
  assertToken(t, lex.getToken(), IDENTIFIER, "bar")
  assertToken(t, lex.getToken(), ';', ";")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), '}', "}")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "baz")
  assertToken(t, lex.getToken(), ';', ";")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestIdentifier1(t *testing.T) {
  lex := lexer("test.txt", "foo\nbar\nbaz\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "bar")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "baz")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestIdentifier2(t *testing.T) {
  lex := lexer("test.txt", "f00\n64r\nb42\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "f00")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "64")
  assertToken(t, lex.getToken(), IDENTIFIER, "r")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "b42")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestInteger1(t *testing.T) {
  lex := lexer("test.txt", "1\n23U\n456L\n7890UL\n")
  assertToken(t, lex.getToken(), INTEGER, "1")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "23U")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "456L")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "7890UL")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestInteger2(t *testing.T) {
  lex := lexer("test.txt", "0xf00\n0x64U\n0X642L\n0xc4febabe\n0XC0FFEEUL\n")
  assertToken(t, lex.getToken(), INTEGER, "0xf00")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "0x64U")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "0X642L")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "0xc4febabe")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "0XC0FFEEUL")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestInteger3(t *testing.T) {
  lex := lexer("test.txt", "0\n012U\n034L\n056UL\n")
  assertToken(t, lex.getToken(), INTEGER, "0")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "012U")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "034L")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertToken(t, lex.getToken(), INTEGER, "056UL")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}

func TestCharacter1(t *testing.T) {
  lex := lexer("test.txt", "{'f', 'o', 'o'}")
  assertToken(t, lex.getToken(), '{', "{")
  assertToken(t, lex.getToken(), CHARACTER, "'f'")
  assertToken(t, lex.getToken(), ',', ",")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), CHARACTER, "'o'")
  assertToken(t, lex.getToken(), ',', ",")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), CHARACTER, "'o'")
  assertToken(t, lex.getToken(), '}', "}")
  assertTokenNull(t, lex.getToken())
}

func TestCharacter2(t *testing.T) {
  lex := lexer("test.txt", "'\x20'")
  assertToken(t, lex.getToken(), CHARACTER, "'\x20'")
  assertTokenNull(t, lex.getToken())
}

func TestString1(t *testing.T) {
  lex := lexer("test.txt", "\"foo, bar, baz\"")
  assertToken(t, lex.getToken(), STRING, "\"foo, bar, baz\"")
  assertTokenNull(t, lex.getToken())
}

func TestOperator1(t *testing.T) {
  lex := lexer("test.txt", "+++....<<<<===&=&&")
  assertToken(t, lex.getToken(), PLUSPLUS, "++")
  assertToken(t, lex.getToken(), '+', "+")
  assertToken(t, lex.getToken(), DOTDOTDOT, "...")
  assertToken(t, lex.getToken(), '.', ".")
  assertToken(t, lex.getToken(), LSHIFT, "<<")
  assertToken(t, lex.getToken(), LSHIFTEQ, "<<=")
  assertToken(t, lex.getToken(), EQEQ, "==")
  assertToken(t, lex.getToken(), ANDEQ, "&=")
  assertToken(t, lex.getToken(), ANDAND, "&&")
  assertTokenNull(t, lex.getToken())
}

func TestOperator2(t *testing.T) {
  lex := lexer("test.txt", "foo ? bar : baz;\n")
  assertToken(t, lex.getToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), '?', "?")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), IDENTIFIER, "bar")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), ':', ":")
//assertToken(t, lex.getToken(), SPACES, " ")
  assertToken(t, lex.getToken(), IDENTIFIER, "baz")
  assertToken(t, lex.getToken(), ';', ";")
//assertToken(t, lex.getToken(), SPACES, "\n")
  assertTokenNull(t, lex.getToken())
}
