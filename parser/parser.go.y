%{
package parser

import (
  "fmt"
  "os"
  "bitbucket.org/yyuu/bs/ast"
)
%}

%union {
  node ast.INode
  token token
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
%token LSHIFTEQ
%token RSHIFTEQ
%token NEQ
%token MODEQ
%token ANDAND
%token ANDEQ
%token MULEQ
%token PLUSPLUS
%token PLUSEQ
%token MINUSMINUS
%token MINUSEQ
%token ARROW
%token DIVEQ
%token LSHIFT
%token LE
%token EQEQ
%token GE
%token RSHIFT
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
      fmt.Println($1.node.DumpString())
    }
    ;

expr: expr2
    {
      // FIXME:
      $$.node = $1.node
    }
    ;

expr2: expr1
     | expr1 '+' expr1
     | expr1 '-' expr1
     ;

expr1: term
     | term '*' term
     | term '/' term
     | term '%' term
     ;

term: unary
    ;

unary: PLUSPLUS unary
     | MINUSMINUS unary
     | '+' term
     | '-' term
     | '!' term
     | '~' term
     | '*' term
     | '&' term
     | postfix
     ;

postfix: primary PLUSPLUS
       | primary MINUSMINUS
       | primary '[' expr ']'
       | primary '.' name
       | primary ARROW name
       | primary '(' args ')'
       ;

name: IDENTIFIER
    ;

args: expr
    | args ',' expr
    ;

primary: INTEGER
       {
         $$.node = ast.IntegerLiteralNode($1.token.Literal)
       }
       | CHARACTER
       {
         $$.node = ast.IntegerLiteralNode($1.token.Literal)
       }
       | STRING
       {
         $$.node = ast.StringLiteralNode($1.token.Literal)
       }
       ;

%%

const EOF = 0
const DEBUG = true

func (self *Lexer) Lex(lval *yySymType) int {
  t := self.GetToken()
  if t == nil {
    return EOF
  } else {
    if DEBUG {
      fmt.Println("token:", t)
    }
    lval.token = *t
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
