%{
package parser

import (
  "errors"
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
)
%}

%union {
  _token token

  _expr ast.IExprNode
  _exprs []ast.IExprNode

  _stmt ast.IStmtNode
  _stmts []ast.IStmtNode
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
       {
         if lex, ok := yylex.(*lex); ok {
           lex.ast = &ast.AST { $1._stmts }
         } else {
           panic("parser is broken")
         }
       }
       ;

block: '{' defvar_list stmts '}'
     {
       $$._stmt = ast.NewBlockNode($1._token.location, $2._exprs, $3._stmts)
     }
     ;

defvar_list:
           | defvar_list defvars
           ;

defvars: storage type defvars_names ';'
       ;

defvars_names: name
             | name '=' expr
             | defvars_names ',' name
             | defvars_names ',' name '=' expr
             ;

storage: 
       | STATIC
       ;

type: typeref
    ;

typeref: typeref_base
       | typeref '[' ']'
       | typeref '[' INTEGER ']'
       | typeref '*'
       | typeref '(' param_typerefs ')'
       ;

param_typerefs: VOID
              | fixedparam_typerefs
              ;

fixedparam_typerefs: typeref
                   | fixedparam_typerefs ',' typeref
                   ;

typeref_base: VOID
            | CHAR
            | SHORT
            | INT
            | LONG
            | UNSIGNED CHAR
            | UNSIGNED SHORT
            | UNSIGNED INT
            | UNSIGNED LONG
            | STRUCT IDENTIFIER
            | UNION IDENTIFIER
            ;

stmts:
     | stmts stmt
     {
       $$._stmts = append($1._stmts, $2._stmt)
     }
     ;

stmt: ';'
    | expr ';'
    {
      $$._stmt = ast.NewExprStmtNode($1._token.location, $1._expr)
    }
    | block
    | if_stmt
    | while_stmt
    | dowhile_stmt
    | for_stmt
    | switch_stmt
    | break_stmt
    | continue_stmt
    | goto_stmt
    | return_stmt
    ;

if_stmt: IF '(' expr ')' stmt ELSE stmt
       {
         $$._stmt = ast.NewIfNode($1._token.location, $3._expr, $5._stmt, $7._stmt)
       }
       ;

while_stmt: WHILE '(' expr ')' stmt
          {
            $$._stmt = ast.NewWhileNode($1._token.location, $3._expr, $5._stmt)
          }
          ;

dowhile_stmt: DO stmt WHILE '(' expr ')' ';'
            {
              $$._stmt = ast.NewDoWhileNode($1._token.location, $2._stmt, $5._expr)
            }
            ;

for_stmt: FOR '(' expr ';' expr ';' expr ')' stmt
        {
          $$._stmt = ast.NewForNode($1._token.location, $3._expr, $5._expr, $7._expr, $9._stmt)
        }
        ;

switch_stmt: SWITCH '(' expr ')' '{' case_clauses '}'
           {
             $$._stmt = ast.NewSwitchNode($1._token.location, $3._expr, $6._stmts)
           }
           ;

case_clauses:
            | case_clauses case_clause
            {
              $$._stmts = append($1._stmts, $2._stmt)
            }
            | case_clauses default_clause
            {
              $$._stmts = append($1._stmts, $2._stmt)
            }
            ;

case_clause: cases case_body
           {
             $$._stmt = ast.NewCaseNode($1._token.location, $1._exprs, $2._stmt)
           }
           ;

cases:
     | cases CASE primary ':'
     {
       $$._exprs = append($1._exprs, $3._expr)
     }
     ;

default_clause: DEFAULT ':' case_body
              {
                $$._stmt = ast.NewCaseNode($1._token.location, []ast.IExprNode { }, $3._stmt)
              }
              ;

case_body: stmt

goto_stmt: GOTO IDENTIFIER ';'
         {
           $$._stmt = ast.NewGotoNode($1._token.location, $2._token.literal)
         }
         ;


break_stmt: BREAK ';'
          {
            $$._stmt = ast.NewBreakNode($1._token.location)
          }
          ;

continue_stmt: CONTINUE ';'
             {
               $$._stmt = ast.NewContinueNode($1._token.location)
             }
             ;

return_stmt: RETURN expr ';'
           {
             $$._stmt = ast.NewReturnNode($1._token.location, $2._expr)
           }
           ;

expr: term '=' expr
    {
      $$._expr = ast.NewAssignNode($1._token.location, $1._expr, $3._expr)
    }
    | term PLUSEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "+", $1._expr, $3._expr)
    }
    | term MINUSEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "-", $1._expr, $3._expr)
    }
    | term MULEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "*", $1._expr, $3._expr)
    }
    | term DIVEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "/", $1._expr, $3._expr)
    }
    | term MODEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "%", $1._expr, $3._expr)
    }
    | term ANDEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "&", $1._expr, $3._expr)
    }
    | term OREQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "|", $1._expr, $3._expr)
    }
    | term XOREQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "^", $1._expr, $3._expr)
    }
    | term LSHIFTEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, "<<", $1._expr, $3._expr)
    }
    | term RSHIFTEQ expr
    {
      $$._expr = ast.NewOpAssignNode($1._token.location, ">>", $1._expr, $3._expr)
    }
    | expr10
    ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      {
        $$._expr = ast.NewCondExprNode($1._token.location, $1._expr, $3._expr, $5._expr)
      }
      ;

expr9: expr8
     | expr9 OROR expr8
     {
       $$._expr = ast.NewLogicalOrNode($1._token.location, $1._expr, $3._expr)
     }
     ;

expr8: expr7
     | expr8 ANDAND expr7
     {
       $$._expr = ast.NewLogicalAndNode($1._token.location, $1._expr, $3._expr)
     }
     ;

expr7: expr6
     | expr7 '>' expr6
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, ">", $1._expr, $3._expr)
     }
     | expr7 '<' expr6
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "<", $1._expr, $3._expr)
     }
     | expr7 GTEQ expr6
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, ">=", $1._expr, $3._expr)
     }
     | expr7 LTEQ expr6
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "<=", $1._expr, $3._expr)
     }
     | expr7 EQEQ expr6
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "==", $1._expr, $3._expr)
     }
     | expr7 NEQ expr6
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "!=", $1._expr, $3._expr)
     }
     ;

expr6: expr5
     | expr6 '|' expr5
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "|", $1._expr, $3._expr)
     }
     ;

expr5: expr4
     | expr5 '^' expr4
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "^", $1._expr, $3._expr)
     }
     ;

expr4: expr3
     | expr4 '&' expr3
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "&", $1._expr, $3._expr)
     }
     ;

expr3: expr2
     | expr3 RSHIFT expr2
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, ">>", $1._expr, $3._expr)
     }
     | expr3 LSHIFT expr2
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "<<", $1._expr, $3._expr)
     }
     ;

expr2: expr1
     | expr2 '+' expr1
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "+", $1._expr, $3._expr)
     }
     | expr2 '-' expr1
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "-", $1._expr, $3._expr)
     }
     ;

expr1: term
     | expr1 '*' term
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "*", $1._expr, $3._expr)
     }
     | expr1 '/' term
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "/", $1._expr, $3._expr)
     }
     | expr1 '%' term
     {
       $$._expr = ast.NewBinaryOpNode($1._token.location, "%", $1._expr, $3._expr)
     }
     ;

term: unary
    ;

unary: PLUSPLUS unary
     {
       $$._expr = ast.NewPrefixOpNode($1._token.location, "++", $2._expr)
     }
     | MINUSMINUS unary
     {
       $$._expr = ast.NewPrefixOpNode($1._token.location, "--", $2._expr)
     }
     | '+' term
     {
       $$._expr = ast.NewUnaryOpNode($1._token.location, "+", $2._expr)
     }
     | '-' term
     {
       $$._expr = ast.NewUnaryOpNode($1._token.location, "-", $2._expr)
     }
     | '!' term
     {
       $$._expr = ast.NewUnaryOpNode($1._token.location, "!", $2._expr)
     }
     | '~' term
     {
       $$._expr = ast.NewUnaryOpNode($1._token.location, "~", $2._expr)
     }
     | postfix
     ;

postfix: primary
       | primary PLUSPLUS
       {
         $$._expr = ast.NewSuffixOpNode($1._token.location, "++", $1._expr)
       }
       | primary MINUSMINUS
       {
         $$._expr = ast.NewSuffixOpNode($1._token.location, "--", $1._expr)
       }
       | primary '(' ')'
       {
         $$._expr = ast.NewFuncallNode($1._token.location, $1._expr, []ast.IExprNode { })
       }
       | primary '(' args ')'
       {
         $$._expr = ast.NewFuncallNode($1._token.location, $1._expr, $3._exprs)
       }
       ;

name: IDENTIFIER
    ;

args: expr
    {
      $$._exprs = []ast.IExprNode { $1._expr }
    }
    | args ',' expr
    {
      $$._exprs = append($1._exprs, $3._expr)
    }
    ;

primary: INTEGER
       {
         $$._expr = ast.NewIntegerLiteralNode($1._token.location, $1._token.literal)
       }
       | CHARACTER
       {
         // TODO: decode character literal
         $$._expr = ast.NewIntegerLiteralNode($1._token.location, $1._token.literal)
       }
       | STRING
       {
         $$._expr = ast.NewStringLiteralNode($1._token.location, $1._token.literal)
       }
       | IDENTIFIER
       {
         $$._expr = ast.NewVariableNode($1._token.location, $1._token.literal)
       }
       | '(' expr ')'
       {
         $$._expr = $2._expr
       }
       ;

%%

const EOF = 0
var VERBOSE = false

func (self *lex) Lex(lval *yySymType) int {
  t := self.getToken()
  if t == nil {
    return EOF
  } else {
    if VERBOSE {
      fmt.Println("token:", t)
    }
    lval._token = *t
    return t.id
  }
}

func (self *lex) Error(s string) {
  self.error = errors.New(s)
  panic(fmt.Errorf("%s: %s", self, s))
}

func ParseExpr(s string) (*ast.AST, error) {
  lex := lexer("", s)
  if yyParse(lex) == 0 {
    return lex.ast, nil // success
  } else {
    if lex.error == nil {
      panic("must not happen")
    }
    return nil, lex.error
  }
}
