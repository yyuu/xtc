%{
package parser

import (
  "errors"
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
)
%}

%union {
  expr ast.IExprNode
  exprs []ast.IExprNode
  stmt ast.IStmtNode
  stmts []ast.IStmtNode
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
       {
         if lex, ok := yylex.(*lex); ok {
           lex.ast = &ast.AST { $1.stmts }
         } else {
           panic("parser is broken")
         }
       }
       ;

block: '{' defvar_list stmts '}'
     {
       $$.stmt = ast.BlockNode($1.token.location, $2.exprs, $3.stmts)
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
       $$.stmts = append($1.stmts, $2.stmt)
     }
     ;

stmt: ';'
    | expr ';'
    {
      $$.stmt = ast.ExprStmtNode($1.token.location, $1.expr)
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
         $$.stmt = ast.IfNode($1.token.location, $3.expr, $5.stmt, $7.stmt)
       }
       ;

while_stmt: WHILE '(' expr ')' stmt
          {
            $$.stmt = ast.WhileNode($1.token.location, $3.expr, $5.stmt)
          }
          ;

dowhile_stmt: DO stmt WHILE '(' expr ')' ';'
            {
              $$.stmt = ast.DoWhileNode($1.token.location, $2.stmt, $5.expr)
            }
            ;

for_stmt: FOR '(' expr ';' expr ';' expr ')' stmt
        {
          $$.stmt = ast.ForNode($1.token.location, $3.expr, $5.expr, $7.expr, $9.stmt)
        }
        ;

switch_stmt: SWITCH '(' expr ')' '{' case_clauses '}'
           {
             $$.stmt = ast.SwitchNode($1.token.location, $3.expr, $6.stmts)
           }
           ;

case_clauses:
            | case_clauses case_clause
            {
              $$.stmts = append($1.stmts, $2.stmt)
            }
            | case_clauses default_clause
            {
              $$.stmts = append($1.stmts, $2.stmt)
            }
            ;

case_clause: cases case_body
           {
             $$.stmt = ast.CaseNode($1.token.location, $1.exprs, $2.stmt)
           }
           ;

cases:
     | cases CASE primary ':'
     {
       $$.exprs = append($1.exprs, $3.expr)
     }
     ;

default_clause: DEFAULT ':' case_body
              {
                $$.stmt = ast.CaseNode($1.token.location, []ast.IExprNode { }, $3.stmt)
              }
              ;

case_body: stmt

goto_stmt: GOTO IDENTIFIER ';'
         {
           $$.stmt = ast.GotoNode($1.token.location, $2.token.literal)
         }
         ;


break_stmt: BREAK ';'
          {
            $$.stmt = ast.BreakNode($1.token.location)
          }
          ;

continue_stmt: CONTINUE ';'
             {
               $$.stmt = ast.ContinueNode($1.token.location)
             }
             ;

return_stmt: RETURN expr ';'
           {
             $$.stmt = ast.ReturnNode($1.token.location, $2.expr)
           }
           ;

expr: term '=' expr
    {
      $$.expr = ast.AssignNode($1.token.location, $1.expr, $3.expr)
    }
    | term PLUSEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "+", $1.expr, $3.expr)
    }
    | term MINUSEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "-", $1.expr, $3.expr)
    }
    | term MULEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "*", $1.expr, $3.expr)
    }
    | term DIVEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "/", $1.expr, $3.expr)
    }
    | term MODEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "%", $1.expr, $3.expr)
    }
    | term ANDEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "&", $1.expr, $3.expr)
    }
    | term OREQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "|", $1.expr, $3.expr)
    }
    | term XOREQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "^", $1.expr, $3.expr)
    }
    | term LSHIFTEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, "<<", $1.expr, $3.expr)
    }
    | term RSHIFTEQ expr
    {
      $$.expr = ast.OpAssignNode($1.token.location, ">>", $1.expr, $3.expr)
    }
    | expr10
    ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      {
        $$.expr = ast.CondExprNode($1.token.location, $1.expr, $3.expr, $5.expr)
      }
      ;

expr9: expr8
     | expr9 OROR expr8
     {
       $$.expr = ast.LogicalOrNode($1.token.location, $1.expr, $3.expr)
     }
     ;

expr8: expr7
     | expr8 ANDAND expr7
     {
       $$.expr = ast.LogicalAndNode($1.token.location, $1.expr, $3.expr)
     }
     ;

expr7: expr6
     | expr7 '>' expr6
     {
       $$.expr = ast.BinaryOpNode($1.token.location, ">", $1.expr, $3.expr)
     }
     | expr7 '<' expr6
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "<", $1.expr, $3.expr)
     }
     | expr7 GTEQ expr6
     {
       $$.expr = ast.BinaryOpNode($1.token.location, ">=", $1.expr, $3.expr)
     }
     | expr7 LTEQ expr6
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "<=", $1.expr, $3.expr)
     }
     | expr7 EQEQ expr6
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "==", $1.expr, $3.expr)
     }
     | expr7 NEQ expr6
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "!=", $1.expr, $3.expr)
     }
     ;

expr6: expr5
     | expr6 '|' expr5
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "|", $1.expr, $3.expr)
     }
     ;

expr5: expr4
     | expr5 '^' expr4
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "^", $1.expr, $3.expr)
     }
     ;

expr4: expr3
     | expr4 '&' expr3
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "&", $1.expr, $3.expr)
     }
     ;

expr3: expr2
     | expr3 RSHIFT expr2
     {
       $$.expr = ast.BinaryOpNode($1.token.location, ">>", $1.expr, $3.expr)
     }
     | expr3 LSHIFT expr2
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "<<", $1.expr, $3.expr)
     }
     ;

expr2: expr1
     | expr2 '+' expr1
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "+", $1.expr, $3.expr)
     }
     | expr2 '-' expr1
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "-", $1.expr, $3.expr)
     }
     ;

expr1: term
     | expr1 '*' term
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "*", $1.expr, $3.expr)
     }
     | expr1 '/' term
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "/", $1.expr, $3.expr)
     }
     | expr1 '%' term
     {
       $$.expr = ast.BinaryOpNode($1.token.location, "%", $1.expr, $3.expr)
     }
     ;

term: unary
    ;

unary: PLUSPLUS unary
     {
       $$.expr = ast.PrefixOpNode($1.token.location, "++", $2.expr)
     }
     | MINUSMINUS unary
     {
       $$.expr = ast.PrefixOpNode($1.token.location, "--", $2.expr)
     }
     | '+' term
     {
       $$.expr = ast.UnaryOpNode($1.token.location, "+", $2.expr)
     }
     | '-' term
     {
       $$.expr = ast.UnaryOpNode($1.token.location, "-", $2.expr)
     }
     | '!' term
     {
       $$.expr = ast.UnaryOpNode($1.token.location, "!", $2.expr)
     }
     | '~' term
     {
       $$.expr = ast.UnaryOpNode($1.token.location, "~", $2.expr)
     }
     | postfix
     ;

postfix: primary
       | primary PLUSPLUS
       {
         $$.expr = ast.SuffixOpNode($1.token.location, "++", $1.expr)
       }
       | primary MINUSMINUS
       {
         $$.expr = ast.SuffixOpNode($1.token.location, "--", $1.expr)
       }
       | primary '(' ')'
       {
         $$.expr = ast.FuncallNode($1.token.location, $1.expr, []ast.IExprNode { })
       }
       | primary '(' args ')'
       {
         $$.expr = ast.FuncallNode($1.token.location, $1.expr, $3.exprs)
       }
       ;

name: IDENTIFIER
    ;

args: expr
    {
      $$.exprs = []ast.IExprNode { $1.expr }
    }
    | args ',' expr
    {
      $$.exprs = append($1.exprs, $3.expr)
    }
    ;

primary: INTEGER
       {
         $$.expr = ast.IntegerLiteralNode($1.token.location, $1.token.literal)
       }
       | CHARACTER
       {
         // TODO: decode character literal
         $$.expr = ast.IntegerLiteralNode($1.token.location, $1.token.literal)
       }
       | STRING
       {
         $$.expr = ast.StringLiteralNode($1.token.location, $1.token.literal)
       }
       | IDENTIFIER
       {
         $$.expr = ast.VariableNode($1.token.location, $1.token.literal)
       }
       | '(' expr ')'
       {
         $$.expr = $2.expr
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
    lval.token = *t
    return t.id
  }
}

func (self *lex) Error(s string) {
  self.error = errors.New(s)
  panic(fmt.Errorf("%s: %s", self, s))
}

func ParseExpr(s string) (*ast.AST, error) {
  lex := lexer("-", s)
  if yyParse(lex) == 0 {
    return lex.ast, nil // success
  } else {
    if lex.error == nil {
      panic("must not happen")
    }
    return nil, lex.error
  }
}
