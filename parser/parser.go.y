%{
package parser

import (
  "fmt"
  "os"
  "bitbucket.org/yyuu/bs/ast"
)
%}

%union {
  literal string
  node ast.INode
}

%token SPACES
%token BLOCK_COMMENT
%token LINE_COMMENT
%token IDENTIFIER
%token INTEGER
%token CHARACTER
%token STRING

/* keywords */
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

/* operators */
%token DOTDOTDOT
%token SHIFTLEFTEQ
%token SHIFTRIGHTEQ
%token NEQ
%token MODEQ
%token ANDAND
%token ANDEQ
%token MULEQ
%token PLUSPLUS
%token PLUSEQ
%token MINUSMINUS
%token MINUSEQ
%token MINUSGT
%token DIVEQ
%token SHIFTLEFT
%token LE
%token EQEQ
%token GE
%token SHIFTRIGHT
%token XOREQ
%token OREQ
%token OROR

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
      $$.node = ast.StringLiteralNode($1.literal)
    }
    | INTEGER
    {
      $$.node = ast.IntegerLiteralNode($1.literal)
    }
    ;

%%

const EOF = 0

func (self *Lexer) Lex(lval *yySymType) int {
  t := self.GetToken()
  if t == nil {
    return EOF
  } else {
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
