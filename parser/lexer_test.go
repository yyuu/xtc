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
  assertTokenNull(t, lex.getNextToken())
}

func TestSpaces(t *testing.T) {
  lex := lexer("test.txt", "\tfoo\n\t\tbar\n\n")
//assertToken(t, lex.getNextToken(), SPACES, "\t")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getNextToken(), SPACES, "\n\t\t")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "bar")
//assertToken(t, lex.getNextToken(), SPACES, "\n\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestBlockComment1(t *testing.T) {
  lex := lexer("test.txt", "/* foo */\n")
//assertToken(t, lex.getNextToken(), BLOCK_COMMENT, "/* foo */")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestBlockComment2(t *testing.T) {
  lex := lexer("test.txt", "foo\n/* bar\n   baz\n */\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
//assertToken(t, lex.getNextToken(), BLOCK_COMMENT, "/* bar\n   baz\n */")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestLineComment(t *testing.T) {
  lex := lexer("test.txt", "foo\n// bar\nbaz\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
//assertToken(t, lex.getNextToken(), LINE_COMMENT, "// bar\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "baz")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestKeyword1(t *testing.T) {
  lex := lexer("test.txt", "unsigned int foo;\n")
  assertToken(t, lex.getNextToken(), UNSIGNED, "unsigned")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), INT, "int")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "foo")
  assertToken(t, lex.getNextToken(), ';', ";")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestKeyword2(t *testing.T) {
  lex := lexer("test.txt", "\n\nif ( foo ) {\n  bar;\n}\nbaz;\n")
//assertToken(t, lex.getNextToken(), SPACES, "\n\n")
  assertToken(t, lex.getNextToken(), IF, "if")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), '(', "(")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), ')', ")")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), '{', "{")
//assertToken(t, lex.getNextToken(), SPACES, "\n  ")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "bar")
  assertToken(t, lex.getNextToken(), ';', ";")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), '}', "}")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "baz")
  assertToken(t, lex.getNextToken(), ';', ";")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestIdentifier1(t *testing.T) {
  lex := lexer("test.txt", "foo\nbar\nbaz\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "bar")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "baz")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestIdentifier2(t *testing.T) {
  lex := lexer("test.txt", "f00\n64r\nb42\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "f00")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "64")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "r")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "b42")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestInteger1(t *testing.T) {
  lex := lexer("test.txt", "1\n23U\n456L\n7890UL\n")
  assertToken(t, lex.getNextToken(), INTEGER, "1")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "23U")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "456L")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "7890UL")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestInteger2(t *testing.T) {
  lex := lexer("test.txt", "0xf00\n0x64U\n0X642L\n0xc4febabe\n0XC0FFEEUL\n")
  assertToken(t, lex.getNextToken(), INTEGER, "0xf00")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "0x64U")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "0X642L")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "0xc4febabe")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "0XC0FFEEUL")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestInteger3(t *testing.T) {
  lex := lexer("test.txt", "0\n012U\n034L\n056UL\n")
  assertToken(t, lex.getNextToken(), INTEGER, "0")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "012U")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "034L")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertToken(t, lex.getNextToken(), INTEGER, "056UL")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}

func TestCharacter1(t *testing.T) {
  lex := lexer("test.txt", "{'f', 'o', 'o'}")
  assertToken(t, lex.getNextToken(), '{', "{")
  assertToken(t, lex.getNextToken(), CHARACTER, "'f'")
  assertToken(t, lex.getNextToken(), ',', ",")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), CHARACTER, "'o'")
  assertToken(t, lex.getNextToken(), ',', ",")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), CHARACTER, "'o'")
  assertToken(t, lex.getNextToken(), '}', "}")
  assertTokenNull(t, lex.getNextToken())
}

func TestCharacter2(t *testing.T) {
  lex := lexer("test.txt", "'\x20'")
  assertToken(t, lex.getNextToken(), CHARACTER, "'\x20'")
  assertTokenNull(t, lex.getNextToken())
}

func TestString1(t *testing.T) {
  lex := lexer("test.txt", "\"foo, bar, baz\"")
  assertToken(t, lex.getNextToken(), STRING, "\"foo, bar, baz\"")
  assertTokenNull(t, lex.getNextToken())
}

func TestOperator1(t *testing.T) {
  lex := lexer("test.txt", "+++....<<<<===&=&&")
  assertToken(t, lex.getNextToken(), PLUSPLUS, "++")
  assertToken(t, lex.getNextToken(), '+', "+")
  assertToken(t, lex.getNextToken(), DOTDOTDOT, "...")
  assertToken(t, lex.getNextToken(), '.', ".")
  assertToken(t, lex.getNextToken(), LSHIFT, "<<")
  assertToken(t, lex.getNextToken(), LSHIFTEQ, "<<=")
  assertToken(t, lex.getNextToken(), EQEQ, "==")
  assertToken(t, lex.getNextToken(), ANDEQ, "&=")
  assertToken(t, lex.getNextToken(), ANDAND, "&&")
  assertTokenNull(t, lex.getNextToken())
}

func TestOperator2(t *testing.T) {
  lex := lexer("test.txt", "foo ? bar : baz;\n")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "foo")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), '?', "?")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "bar")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), ':', ":")
//assertToken(t, lex.getNextToken(), SPACES, " ")
  assertToken(t, lex.getNextToken(), IDENTIFIER, "baz")
  assertToken(t, lex.getNextToken(), ';', ";")
//assertToken(t, lex.getNextToken(), SPACES, "\n")
  assertTokenNull(t, lex.getNextToken())
}
