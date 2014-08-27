%{
package parser

import (
  "fmt"
  "os"
  "strconv"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/lexer"
)
%}

%union {
  node ast.INode
  literal string
}

%token SPACES
%token BLOCK_COMMENT
%token LINE_COMMENT
%token VOID
%token CHAR
%token SHORT
%token INT
%token LONG
%token STRUCT
%token UNION
%token ENUM
%token STATIC
%token EXTERN
%token CONST
%token SIGNED
%token UNSIGNED
%token IF
%token ELSE
%token SWITCH
%token CASE
%token DEFAULT
%token WHILE
%token DO
%token FOR
%token RETURN
%token BREAK
%token CONTINUE
%token GOTO
%token TYPEDEF
%token IMPORT
%token SIZEOF
%token IDENTIFIER
%token INTEGER
%token CHARACTER
%token STRING
%token OPERATOR

%%

program: stmts
       ;

stmts:
     | stmts stmt
     ;

stmt: expr
    {
      fmt.Println($1.node)
    }
    ;

expr: STRING
    {
      $$.node = ast.StringLiteralNode { $1.literal }
    }
    | INTEGER
    {
      i, _ := strconv.Atoi($1.literal)
      $$.node = ast.IntegerLiteralNode { i }
    }
    ;

%%

const EOF = 0

type lex struct {
  lex *lexer.Lexer
}

func (self *lex) Lex(lval *yySymType) int {
  for {
    t := self.lex.GetToken()
    if t == nil {
      return EOF
    }
    if id, ok := id2id[t.Id]; ok {
      lval.literal = t.Literal
      return id
    }
  }
}

func (self *lex) Error(s string) {
  fmt.Fprintf(os.Stderr, "parse error: %s at %s:%d,%d\n", s, self.lex.Filename, self.lex.LineNumber, self.lex.LineOffset)
}

var id2id map[int]int = map[int]int {
//lexer.SPACES: SPACES,
//lexer.BLOCK_COMMENT: BLOCK_COMMENT,
//lexer.LINE_COMMENT: LINE_COMMENT,
  lexer.VOID: VOID,
  lexer.CHAR: CHAR,
  lexer.SHORT: SHORT,
  lexer.INT: INT,
  lexer.LONG: LONG,
  lexer.STRUCT: STRUCT,
  lexer.UNION: UNION,
  lexer.ENUM: ENUM,
  lexer.STATIC: STATIC,
  lexer.EXTERN: EXTERN,
  lexer.CONST: CONST,
  lexer.SIGNED: SIGNED,
  lexer.UNSIGNED: UNSIGNED,
  lexer.IF: IF,
  lexer.ELSE: ELSE,
  lexer.SWITCH: SWITCH,
  lexer.CASE: CASE,
  lexer.DEFAULT: DEFAULT,
  lexer.WHILE: WHILE,
  lexer.DO: DO,
  lexer.FOR: FOR,
  lexer.RETURN: RETURN,
  lexer.BREAK: BREAK,
  lexer.CONTINUE: CONTINUE,
  lexer.GOTO: GOTO,
  lexer.TYPEDEF: TYPEDEF,
  lexer.IMPORT: IMPORT,
  lexer.SIZEOF: SIZEOF,
  lexer.IDENTIFIER: IDENTIFIER,
  lexer.INTEGER: INTEGER,
  lexer.CHARACTER: CHARACTER,
  lexer.STRING: STRING,
  lexer.OPERATOR: OPERATOR,
}

func ParseExpr(s string) {
  yyParse(&lex { lexer.NewLexer("-", s) })
}
