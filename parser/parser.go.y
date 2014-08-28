%{
package parser

import (
  "errors"
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
)
%}

%union {
  node ast.INode
  nodes []ast.INode
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
           lex.nodes = $1.nodes
         } else {
           panic("parser is broken")
         }
       }
       ;

block: '{' defvar_list stmts '}'
     {
       $$.node = ast.BlockNode($2.nodes, $3.nodes)
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
       $$.nodes = append($1.nodes, $2.node)
     }
     ;

stmt: ';'
    | expr ';'
    {
      $$.node = ast.ExprStmtNode($1.node)
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
         $$.node = ast.IfNode($3.node, $5.node, $7.node)
       }
       ;

while_stmt: WHILE '(' expr ')' stmt
          {
            $$.node = ast.WhileNode($3.node, $5.node)
          }
          ;

dowhile_stmt: DO stmt WHILE '(' expr ')' ';'
            {
              $$.node = ast.DoWhileNode($2.node, $5.node)
            }
            ;

for_stmt: FOR '(' expr ';' expr ';' expr ')' stmt
        {
          $$.node = ast.ForNode($3.node, $5.node, $7.node, $9.node)
        }
        ;

switch_stmt: SWITCH '(' expr ')' '{' case_clauses '}'
           {
             $$.node = ast.SwitchNode($3.node, $6.nodes)
           }
           ;

case_clauses:
            | case_clauses case_clause
            {
              $$.nodes = append($1.nodes, $2.node)
            }
            | case_clauses default_clause
            {
              $$.nodes = append($1.nodes, $2.node)
            }
            ;

case_clause: cases case_body
           {
             $$.node = ast.CaseNode($1.nodes, $2.node)
           }
           ;

cases:
     | cases CASE primary ':'
     {
       $$.nodes = append($1.nodes, $3.node)
     }
     ;

default_clause: DEFAULT ':' case_body
              {
                _default := []ast.INode { ast.StringLiteralNode("default") } // FIXME:
                $$.node = ast.CaseNode(_default, $3.node)
              }
              ;

case_body: stmt

goto_stmt: GOTO IDENTIFIER ';'
         {
           $$.node = ast.GotoNode($2.token.Literal)
         }
         ;


break_stmt: BREAK ';'
          {
            $$.node = ast.BreakNode()
          }
          ;

continue_stmt: CONTINUE ';'
             {
               $$.node = ast.ContinueNode()
             }
             ;

return_stmt: RETURN expr ';'
           {
             $$.node = ast.ReturnNode($2.node)
           }
           ;

expr: term '=' expr
    {
      $$.node = ast.AssignNode($1.node, $3.node)
    }
    | term PLUSEQ expr
    {
      $$.node = ast.OpAssignNode("+", $1.node, $3.node)
    }
    | term MINUSEQ expr
    {
      $$.node = ast.OpAssignNode("-", $1.node, $3.node)
    }
    | term MULEQ expr
    {
      $$.node = ast.OpAssignNode("*", $1.node, $3.node)
    }
    | term DIVEQ expr
    {
      $$.node = ast.OpAssignNode("/", $1.node, $3.node)
    }
    | term MODEQ expr
    {
      $$.node = ast.OpAssignNode("%", $1.node, $3.node)
    }
    | term ANDEQ expr
    {
      $$.node = ast.OpAssignNode("&", $1.node, $3.node)
    }
    | term OREQ expr
    {
      $$.node = ast.OpAssignNode("|", $1.node, $3.node)
    }
    | term XOREQ expr
    {
      $$.node = ast.OpAssignNode("^", $1.node, $3.node)
    }
    | term LSHIFTEQ expr
    {
      $$.node = ast.OpAssignNode("<<", $1.node, $3.node)
    }
    | term RSHIFTEQ expr
    {
      $$.node = ast.OpAssignNode(">>", $1.node, $3.node)
    }
    | expr10
    ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      {
        $$.node = ast.CondExprNode($1.node, $3.node, $5.node)
      }
      ;

expr9: expr8
     | expr9 OROR expr8
     {
       $$.node = ast.LogicalOrNode($1.node, $3.node)
     }
     ;

expr8: expr7
     | expr8 ANDAND expr7
     {
       $$.node = ast.LogicalAndNode($1.node, $3.node)
     }
     ;

expr7: expr6
     | expr7 '>' expr6
     {
       $$.node = ast.BinaryOpNode(">", $1.node, $3.node)
     }
     | expr7 '<' expr6
     {
       $$.node = ast.BinaryOpNode("<", $1.node, $3.node)
     }
     | expr7 GTEQ expr6
     {
       $$.node = ast.BinaryOpNode(">=", $1.node, $3.node)
     }
     | expr7 LTEQ expr6
     {
       $$.node = ast.BinaryOpNode("<=", $1.node, $3.node)
     }
     | expr7 EQEQ expr6
     {
       $$.node = ast.BinaryOpNode("==", $1.node, $3.node)
     }
     | expr7 NEQ expr6
     {
       $$.node = ast.BinaryOpNode("!=", $1.node, $3.node)
     }
     ;

expr6: expr5
     | expr6 '|' expr5
     {
       $$.node = ast.BinaryOpNode("|", $1.node, $3.node)
     }
     ;

expr5: expr4
     | expr5 '^' expr4
     {
       $$.node = ast.BinaryOpNode("^", $1.node, $3.node)
     }
     ;

expr4: expr3
     | expr4 '&' expr3
     {
       $$.node = ast.BinaryOpNode("&", $1.node, $3.node)
     }
     ;

expr3: expr2
     | expr3 RSHIFT expr2
     {
       $$.node = ast.BinaryOpNode(">>", $1.node, $3.node)
     }
     | expr3 LSHIFT expr2
     {
       $$.node = ast.BinaryOpNode("<<", $1.node, $3.node)
     }
     ;

expr2: expr1
     | expr2 '+' expr1
     {
       $$.node = ast.BinaryOpNode("+", $1.node, $3.node)
     }
     | expr2 '-' expr1
     {
       $$.node = ast.BinaryOpNode("-", $1.node, $3.node)
     }
     ;

expr1: term
     | expr1 '*' term
     {
       $$.node = ast.BinaryOpNode("*", $1.node, $3.node)
     }
     | expr1 '/' term
     {
       $$.node = ast.BinaryOpNode("/", $1.node, $3.node)
     }
     | expr1 '%' term
     {
       $$.node = ast.BinaryOpNode("%", $1.node, $3.node)
     }
     ;

term: unary
    ;

unary: PLUSPLUS unary
     {
       $$.node = ast.PrefixOpNode("++", $2.node)
     }
     | MINUSMINUS unary
     {
       $$.node = ast.PrefixOpNode("--", $2.node)
     }
     | '+' term
     {
       $$.node = ast.UnaryOpNode("+", $2.node)
     }
     | '-' term
     {
       $$.node = ast.UnaryOpNode("-", $2.node)
     }
     | '!' term
     {
       $$.node = ast.UnaryOpNode("!", $2.node)
     }
     | '~' term
     {
       $$.node = ast.UnaryOpNode("~", $2.node)
     }
     | postfix
     ;

postfix: primary
       | primary PLUSPLUS
       {
         $$.node = ast.SuffixOpNode("++", $1.node)
       }
       | primary MINUSMINUS
       {
         $$.node = ast.SuffixOpNode("--", $1.node)
       }
       | primary '(' ')'
       {
         $$.node = ast.FuncallNode($1.node, []ast.INode { })
       }
       | primary '(' args ')'
       {
         $$.node = ast.FuncallNode($1.node, $3.nodes)
       }
       ;

name: IDENTIFIER
    ;

args: expr
    | args ',' expr
    {
      $$.nodes = append($1.nodes, $3.node)
    }
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
       | IDENTIFIER
       {
         $$.node = ast.VariableNode($1.token.Literal)
       }
       | '(' expr ')'
       {
         $$.node = $2.node
       }
       ;

%%

const EOF = 0
const DEBUG = true

func (self *lex) Lex(lval *yySymType) int {
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

func (self *lex) Error(s string) {
  panic(fmt.Errorf("%s: %s", self, s))
}

func ParseExpr(s string) ([]ast.INode, error) {
  lex := lexer("main.c", s)
  if yyParse(lex) == 0 {
    return lex.nodes, nil
  } else {
    return nil, errors.New("parse error") // TODO: get error via lexer
  }
}
