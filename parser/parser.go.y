%{
package parser

import (
  "fmt"
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
%token LTEQ
%token EQEQ
%token GTEQ
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
      fmt.Println($1.node)
    }
    ;

expr: term '=' expr
    | term opassign_op expr
    | expr10
    ;

opassign_op: PLUSEQ
           | MINUSEQ
           | MULEQ
           | DIVEQ
           | MODEQ
           | ANDEQ
           | OREQ
           | XOREQ
           | LSHIFTEQ
           | RSHIFTEQ
           ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      ;

expr9: expr8
     | expr8 OROR expr8
     ;

expr8: expr7
     | expr7 ANDAND expr7
     ;

expr7: expr6
     | expr6 '>' expr6
     | expr6 '<' expr6
     | expr6 GTEQ expr6
     | expr6 LTEQ expr6
     | expr6 EQEQ expr6
     | expr6 NEQ expr6
     ;

expr6: expr5
     | expr5 '|' expr5
     ;

expr5: expr4
     | expr4 '^' expr4
     ;

expr4: expr3
     | expr3 '&' expr3
     ;

expr3: expr2
     | expr2 RSHIFT expr2
     | expr2 LSHIFT expr2
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
         // TODO: decode character literal
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
  panic(fmt.Errorf("%s: %s", self, s))
}

func ParseExpr(s string) {
  yyParse(NewLexer("main.cb", s))
}
