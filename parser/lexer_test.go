package parser

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func assertToken(t *testing.T, lex *lexer, id int, literal string) {
  tok, err := lex.getNextToken()
  if err != nil { t.Error(err.Error()) }
  xt.AssertNotNil(t, "insufficient token", tok)
  xt.AssertEquals(t, "invalid token id", tok.id, id)
  xt.AssertEquals(t, "invalid token literal", tok.literal, literal)
}

func assertTokenNull(t *testing.T, lex *lexer) {
  tok, err := lex.getNextToken()
  if err != nil { t.Error(err.Error()) }
  xt.AssertNotNil(t, "surplus token", tok)
}

func testNewLexer(source string) *lexer {
  errorHandler := core.NewErrorHandler(core.LOG_WARN)
  options := core.NewOptions("lexer_test.go")
  loader := newLibraryLoader(errorHandler, options)
  return newLexer("test.txt", source, loader, errorHandler, options)
}

func TestEmpty(t *testing.T) {
  lex := testNewLexer("")
  assertTokenNull(t, lex)
}

func TestSpaces(t *testing.T) {
  lex := testNewLexer("\tfoo\n\t\tbar\n\n")
//assertToken(t, lex, SPACES, "\t")
  assertToken(t, lex, IDENTIFIER, "foo")
//assertToken(t, lex, SPACES, "\n\t\t")
  assertToken(t, lex, IDENTIFIER, "bar")
//assertToken(t, lex, SPACES, "\n\n")
  assertTokenNull(t, lex)
}

func TestBlockComment1(t *testing.T) {
  lex := testNewLexer("/* foo */\n")
//assertToken(t, lex, BLOCK_COMMENT, "/* foo */")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestBlockComment2(t *testing.T) {
  lex := testNewLexer("foo\n/* bar\n   baz\n */\n")
  assertToken(t, lex, IDENTIFIER, "foo")
//assertToken(t, lex, SPACES, "\n")
//assertToken(t, lex, BLOCK_COMMENT, "/* bar\n   baz\n */")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestLineComment(t *testing.T) {
  lex := testNewLexer("foo\n// bar\nbaz\n")
  assertToken(t, lex, IDENTIFIER, "foo")
//assertToken(t, lex, SPACES, "\n")
//assertToken(t, lex, LINE_COMMENT, "// bar\n")
  assertToken(t, lex, IDENTIFIER, "baz")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestKeyword1(t *testing.T) {
  lex := testNewLexer("unsigned int foo;\n")
  assertToken(t, lex, UNSIGNED, "unsigned")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, INT, "int")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, IDENTIFIER, "foo")
  assertToken(t, lex, ';', ";")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestKeyword2(t *testing.T) {
  lex := testNewLexer("\n\nif ( foo ) {\n  bar;\n}\nbaz;\n")
//assertToken(t, lex, SPACES, "\n\n")
  assertToken(t, lex, IF, "if")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, '(', "(")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, IDENTIFIER, "foo")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, ')', ")")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, '{', "{")
//assertToken(t, lex, SPACES, "\n  ")
  assertToken(t, lex, IDENTIFIER, "bar")
  assertToken(t, lex, ';', ";")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, '}', "}")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, IDENTIFIER, "baz")
  assertToken(t, lex, ';', ";")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestIdentifier1(t *testing.T) {
  lex := testNewLexer("foo\nbar\nbaz\n")
  assertToken(t, lex, IDENTIFIER, "foo")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, IDENTIFIER, "bar")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, IDENTIFIER, "baz")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestIdentifier2(t *testing.T) {
  lex := testNewLexer("f00\n64r\nb42\n")
  assertToken(t, lex, IDENTIFIER, "f00")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "64")
  assertToken(t, lex, IDENTIFIER, "r")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, IDENTIFIER, "b42")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestInteger1(t *testing.T) {
  lex := testNewLexer("1\n23U\n456L\n7890UL\n")
  assertToken(t, lex, INTEGER, "1")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "23U")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "456L")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "7890UL")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestInteger2(t *testing.T) {
  lex := testNewLexer("0xf00\n0x64U\n0X642L\n0xc4febabe\n0XC0FFEEUL\n")
  assertToken(t, lex, INTEGER, "0xf00")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "0x64U")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "0X642L")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "0xc4febabe")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "0XC0FFEEUL")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestInteger3(t *testing.T) {
  lex := testNewLexer("0\n012U\n034L\n056UL\n")
  assertToken(t, lex, INTEGER, "0")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "012U")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "034L")
//assertToken(t, lex, SPACES, "\n")
  assertToken(t, lex, INTEGER, "056UL")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestCharacter1(t *testing.T) {
  lex := testNewLexer(`{'f', 'o', 'o'}`)
  assertToken(t, lex, '{', "{")
  assertToken(t, lex, CHARACTER, "102")
  assertToken(t, lex, ',', ",")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, CHARACTER, "111")
  assertToken(t, lex, ',', ",")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, CHARACTER, "111")
  assertToken(t, lex, '}', "}")
  assertTokenNull(t, lex)
}

func TestCharacter2(t *testing.T) {
  lex := testNewLexer(`{'\n','\t'}`)
  assertToken(t, lex, '{', "{")
  assertToken(t, lex, CHARACTER, "10")
  assertToken(t, lex, ',', ",")
  assertToken(t, lex, CHARACTER, "9")
  assertToken(t, lex, '}', "}")
  assertTokenNull(t, lex)
}

func TestCharacter3(t *testing.T) {
  lex := testNewLexer(`{'\u0009','\\','\''}`)
  assertToken(t, lex, '{', "{")
  assertToken(t, lex, CHARACTER, "9")
  assertToken(t, lex, ',', ",")
  assertToken(t, lex, CHARACTER, "92")
  assertToken(t, lex, ',', ",")
  assertToken(t, lex, CHARACTER, "39")
  assertToken(t, lex, '}', "}")
  assertTokenNull(t, lex)
}

func TestString1(t *testing.T) {
  lex := testNewLexer(`"foo, bar, baz"`)
  assertToken(t, lex, STRING, `foo, bar, baz`)
  assertTokenNull(t, lex)
}

func TestString2(t *testing.T) {
  lex := testNewLexer(`"You say \"Yes\", I say \"No\""`)
  assertToken(t, lex, STRING, `You say "Yes", I say "No"`)
  assertTokenNull(t, lex)
}

func TestOperator1(t *testing.T) {
  lex := testNewLexer("+++....<<<<===&=&&")
  assertToken(t, lex, PLUSPLUS, "++")
  assertToken(t, lex, '+', "+")
  assertToken(t, lex, DOTDOTDOT, "...")
  assertToken(t, lex, '.', ".")
  assertToken(t, lex, LSHIFT, "<<")
  assertToken(t, lex, LSHIFTEQ, "<<=")
  assertToken(t, lex, EQEQ, "==")
  assertToken(t, lex, ANDEQ, "&=")
  assertToken(t, lex, ANDAND, "&&")
  assertTokenNull(t, lex)
}

func TestOperator2(t *testing.T) {
  lex := testNewLexer("foo ? bar : baz;\n")
  assertToken(t, lex, IDENTIFIER, "foo")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, '?', "?")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, IDENTIFIER, "bar")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, ':', ":")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, IDENTIFIER, "baz")
  assertToken(t, lex, ';', ";")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}

func TestIdentifierStartsWithKeyword(t *testing.T) {
  lex := testNewLexer(`format = "%d:%d"\n`)
  assertToken(t, lex, IDENTIFIER, "format")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, '=', "=")
//assertToken(t, lex, SPACES, " ")
  assertToken(t, lex, STRING, "%d:%d")
//assertToken(t, lex, SPACES, "\n")
  assertTokenNull(t, lex)
}
