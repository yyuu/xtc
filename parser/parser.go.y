%{
package parser

import (
  "fmt"
  "os"
  "strconv"
  "bitbucket.org/yyuu/bs/ast"
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

func (self *Lexer) Lex(lval *yySymType) int {
  for {
    t := self.GetToken()
    if t == nil {
      return EOF
    }
    lval.literal = t.Literal
    return t.Id
  }
}

func (self *Lexer) Error(s string) {
  fmt.Fprintf(os.Stderr, "%s: %s\n", *self, s)
  os.Exit(1)
}

func ParseExpr(s string) {
  yyParse(NewLexer("main.cb", s))
}
