%{
package parser

import (
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
         for i := range $1.nodes {
           fmt.Println($1.nodes[i])
         }
       }
       ;

block: '{' defvar_list stmts '}'
     {
       $$.node = ast.BlockNode($2.nodes, $3.nodes)
     }
     ;

defvar_list:
           ;

/*
defvars: storage type name ( '=' expr )? (',' name ( '=' expr )? )* ';'
       ;

type: typeref
    ;

typeref:
       ;

storage:
       | STATIC
       ;
 */

stmts:
     | stmts stmt
     {
       $$.nodes = append($1.nodes, $2.node)
     }
     ;

stmt: ';'
    | expr ';'
    | block
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
       | primary '(' args ')'
       {
         $$.node = ast.FuncallNode($1.node, $3.nodes)
       }
       ;

name: IDENTIFIER
    ;

args: expr
    {
      $$.nodes = []ast.INode { $1.node }
    }
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

func ParseExpr(s string) {
  yyParse(lexer("main.cb", s))
}
