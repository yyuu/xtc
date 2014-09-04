%{
package parser

import (
  "errors"
  "fmt"
  "strconv"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)
%}

%union {
  _token token

  _node duck.INode
  _nodes []duck.INode

  _entity duck.IEntity

  _typeref duck.ITypeRef
  _typerefs []duck.ITypeRef
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

compilation_unit:
                | import_stmts top_defs
                {
                  if lex, ok := yylex.(*lex); ok {
//                  ast := ast.NewAST($1._token.location, asDeclarations($2._node))
                    ast := ast.NewAST($2._token.location, asDeclarations($2._node))
                    lex.ast = &ast
                  } else {
                    panic("parser is broken")
                  }
                }
                ;

declaration_file: import_stmts
                {
//
                }
                | declaration_file funcdecl
                {
                  $$._node = asDeclarations($1._node).AddFuncdecl(asUndefinedFunction($2._entity))
                }
                | declaration_file vardecl
                {
                  $$._node = asDeclarations($1._node).AddVardecl(asUndefinedVariable($2._entity))
                }
                | declaration_file defconst
                {
                  $$._node = asDeclarations($1._node).AddDefconst(asConstant($2._entity))
                }
                | declaration_file defstruct
                {
                  $$._node = asDeclarations($1._node).AddDefstruct(asStructNode($2._node))
                }
                | declaration_file defunion
                {
                  $$._node = asDeclarations($1._node).AddDefunion(asUnionNode($2._node))
                }
                | declaration_file typedef
                {
                  $$._node = asDeclarations($1._node).AddTypedef(asTypedefNode($2._node))
                }
                ;

import_stmts:
            | import_stmts import_stmt
            {
              $$._nodes = append($1._nodes, $2._node)
            }
            ;

import_stmt: IMPORT import_name ';'
           {
//           $$._nodes, err := loader.LoadLibrary($2._tokens.literal)
//           if err != nil {
//             panic(err)
//           }
           }
           ;

import_name: name
           | import_name '.' name
           {
             $$._token.literal = fmt.Sprintf("%s.%s", $1._token.literal, $3._token.literal)
           }
           ;

top_defs:
        {
          $$._node = ast.NewDeclarations()
        }
        | top_defs defun
        {
          $$._node = asDeclarations($1._node).AddDefun(asDefinedFunction($2._entity))
        }
        | top_defs defvars
        {
          $$._node = asDeclarations($1._node).AddDefvar(asDefinedVariable($2._entity))
        }
        | top_defs defconst
        {
          $$._node = asDeclarations($1._node).AddDefconst(asConstant($2._entity))
        }
        | top_defs defstruct
        {
          $$._node = asDeclarations($1._node).AddDefstruct(asStructNode($2._node))
        }
        | top_defs defunion
        {
          $$._node = asDeclarations($1._node).AddDefunion(asUnionNode($2._node))
        }
        | top_defs typedef
        {
          $$._node = asDeclarations($1._node).AddTypedef(asTypedefNode($2._node))
        }
        ;

defvars: storage type name '=' expr ';'
       {
         priv := $1._token.literal == "storage"
         $$._entity = entity.NewDefinedVariable(priv, asType($2._node), $3._token.literal, asExpr($5._node))
       }
       ;

defconst: CONST type name '=' expr ';'
        {
          $$._entity = entity.NewConstant(asType($2._node), $3._token.literal, asExpr($5._node))
        }
        ;

defun: storage typeref name '(' params ')' block
     {
       priv := $1._token.literal == "storage"
       ps := asParams($5._entity)
       t := typesys.NewFunctionTypeRef($2._typeref, ps.ParametersTypeRef())
       $$._entity = entity.NewDefinedFunction(priv, ast.NewTypeNode($1._token.location, t), $3._token.literal, ps, asStmt($7._node))
     }
     ;

storage:
       | STATIC
       ;

params: VOID
      {
        $$._entity = entity.NewParams($1._token.location, []entity.Parameter { })
      }
      | fixedparams
      {
        $$._entity = entity.NewParams($1._token.location, asParams($1._entity).ParamDescs)
      }
      | fixedparams ',' DOTDOTDOT
      {
        $$._entity = entity.NewParams($1._token.location, asParams($1._entity).ParamDescs)
//      $$._entity.AcceptVarArgs()
      }
      ;

fixedparams: param
           {
             $$._entity = entity.NewParams($1._token.location, []entity.Parameter { asParameter($1._entity) })
           }
           | fixedparams ',' param
           {
             $$._entity = entity.NewParams($1._token.location, append(asParams($1._entity).ParamDescs, asParameter($3._entity)))
           }
           ;

param: type name
     {
       $$._entity = entity.NewParameter(asType($1._node), $2._token.literal)
     }
     ;

block: '{' defvar_list stmts '}'
     {
       $$._node = ast.NewBlockNode($1._token.location, asExprs($2._nodes), asStmts($3._nodes))
     }
     ;

defvar_list:
           | defvar_list defvars
           {
             $$._nodes = append($1._nodes, $2._node)
           }
           ;

defstruct: STRUCT name member_list ';'
         {
           $$._node = ast.NewStructNode($1._token.location, typesys.NewStructTypeRef($1._token.location, $2._token.literal), $2._token.literal, asSlots($3._nodes))
         }
         ;

defunion: UNION name member_list ';'
        {
          $$._node = ast.NewUnionNode($1._token.location, typesys.NewUnionTypeRef($1._token.location, $2._token.literal), $2._token.literal, asSlots($3._nodes))
        }
        ;

member_list: '{' member_list_body '}'
           {
             $$._nodes = $2._nodes
           }
           ;

member_list_body: slot ';'
                {
                  $$._nodes = []duck.INode { $1._node }
                }
                | member_list_body slot ';'
                {
                  $$._nodes = append($1._nodes, $2._node)
                }
                ;

slot: type name
    {
      $$._node = ast.NewSlot(asType($1._node), $2._token.literal)
    }
    ;

funcdecl: EXTERN typeref name '(' params ')' ';'
        {
          ps := asParams($5._entity)
          ref := typesys.NewFunctionTypeRef($2._typeref, ps.ParametersTypeRef())
          $$._entity = entity.NewUndefinedFunction(ast.NewTypeNode($1._token.location, ref), $3._token.literal, ps)
        }
        ;

vardecl: EXTERN type name ';'
       {
         $$._entity = entity.NewUndefinedVariable(asType($2._node), $3._token.literal)
       }
       ;

type: typeref
    {
      $$._node = ast.NewTypeNode($1._token.location, $1._typeref)
    }
    ;

typeref: typeref_base
       | typeref '[' ']'
       {
         $$._typeref = typesys.NewArrayTypeRef($1._typeref, 0)
       }
       | typeref '[' INTEGER ']'
       {
         n, _ := strconv.Atoi($3._token.literal)
         $$._typeref = typesys.NewArrayTypeRef($1._typeref, n)
       }
       | typeref '*'
       {
         $$._typeref = typesys.NewPointerTypeRef($1._typeref)
       }
       | typeref '(' param_typerefs ')'
       {
         $$._typeref = typesys.NewFunctionTypeRef($1._typeref, $3._typeref)
       }
       ;

param_typerefs: VOID
              {
                $$._typeref = typesys.NewParamTypeRefs($1._token.location, []duck.ITypeRef { }, false)
              }
              | fixedparam_typerefs
              {
//              $1._typerefs.AcceptVArgs()
                $$._typerefs = $1._typerefs
              }
              ;

fixedparam_typerefs: typeref
                   {
                     $$._typerefs = []duck.ITypeRef { $1._typeref }
                   }
                   | fixedparam_typerefs ',' typeref
                   {
                     $$._typerefs = append($1._typerefs, $3._typeref)
                   }
                   ;

typeref_base: VOID
            {
              $$._typeref = typesys.NewVoidTypeRef($1._token.location)
            }
            | CHAR
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "char")
            }
            | SHORT
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "short")
            }
            | INT
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "int")
            }
            | LONG
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "long")
            }
            | UNSIGNED CHAR
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "unsigned char")
            }
            | UNSIGNED SHORT
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "unsigned short")
            }
            | UNSIGNED INT
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "unsigned int")
            }
            | UNSIGNED LONG
            {
              $$._typeref = typesys.NewIntegerTypeRef($1._token.location, "unsigned long")
            }
            | STRUCT IDENTIFIER
            {
              $$._typeref = typesys.NewStructTypeRef($1._token.location, $2._token.literal)
            }
            | UNION IDENTIFIER
            {
              $$._typeref = typesys.NewUnionTypeRef($1._token.location, $2._token.literal)
            }
            ;

typedef: TYPEDEF typeref IDENTIFIER ';'
       {
         $$._node = ast.NewTypedefNode($1._token.location, $2._typeref, $3._token.literal)
       }
       ;

stmts:
     | stmts stmt
     {
       $$._nodes = append($1._nodes, $2._node)
     }
     ;

stmt: ';'
    | labeled_stmt
    | expr ';'
    {
      $$._node = ast.NewExprStmtNode($1._token.location, asExpr($1._node))
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

labeled_stmt: IDENTIFIER ':' stmt
            {
              $$._node = ast.NewLabelNode($1._token.location, $1._token.literal, asStmt($3._node))
            }
            ;

if_stmt: IF '(' expr ')' stmt ELSE stmt
       {
         $$._node = ast.NewIfNode($1._token.location, asExpr($3._node), asStmt($5._node), asStmt($7._node))
       }
       ;

while_stmt: WHILE '(' expr ')' stmt
          {
            $$._node = ast.NewWhileNode($1._token.location, asExpr($3._node), asStmt($5._node))
          }
          ;

dowhile_stmt: DO stmt WHILE '(' expr ')' ';'
            {
              $$._node = ast.NewDoWhileNode($1._token.location, asStmt($2._node), asExpr($5._node))
            }
            ;

for_stmt: FOR '(' expr ';' expr ';' expr ')' stmt
        {
          $$._node = ast.NewForNode($1._token.location, asExpr($3._node), asExpr($5._node), asExpr($7._node), asStmt($9._node))
        }
        ;

switch_stmt: SWITCH '(' expr ')' '{' case_clauses '}'
           {
             $$._node = ast.NewSwitchNode($1._token.location, asExpr($3._node), asStmts($6._nodes))
           }
           ;

case_clauses:
            | case_clauses case_clause
            {
              $$._nodes = append($1._nodes, $2._node)
            }
            | case_clauses default_clause
            {
              $$._nodes = append($1._nodes, $2._node)
            }
            ;

case_clause: cases case_body
           {
             $$._node = ast.NewCaseNode($1._token.location, asExprs($1._nodes), asStmt($2._node))
           }
           ;

cases:
     | cases CASE primary ':'
     {
       $$._nodes = append($1._nodes, $3._node)
     }
     ;

default_clause: DEFAULT ':' case_body
              {
                $$._node = ast.NewCaseNode($1._token.location, []duck.IExprNode { }, asStmt($3._node))
              }
              ;

case_body: stmt

goto_stmt: GOTO IDENTIFIER ';'
         {
           $$._node = ast.NewGotoNode($1._token.location, $2._token.literal)
         }
         ;


break_stmt: BREAK ';'
          {
            $$._node = ast.NewBreakNode($1._token.location)
          }
          ;

continue_stmt: CONTINUE ';'
             {
               $$._node = ast.NewContinueNode($1._token.location)
             }
             ;

return_stmt: RETURN expr ';'
           {
             $$._node = ast.NewReturnNode($1._token.location, asExpr($2._node))
           }
           ;

expr: term '=' expr
    {
      $$._node = ast.NewAssignNode($1._token.location, asExpr($1._node), asExpr($3._node))
    }
    | term PLUSEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "+", asExpr($1._node), asExpr($3._node))
    }
    | term MINUSEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "-", asExpr($1._node), asExpr($3._node))
    }
    | term MULEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "*", asExpr($1._node), asExpr($3._node))
    }
    | term DIVEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "/", asExpr($1._node), asExpr($3._node))
    }
    | term MODEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "%", asExpr($1._node), asExpr($3._node))
    }
    | term ANDEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "&", asExpr($1._node), asExpr($3._node))
    }
    | term OREQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "|", asExpr($1._node), asExpr($3._node))
    }
    | term XOREQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "^", asExpr($1._node), asExpr($3._node))
    }
    | term LSHIFTEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "<<", asExpr($1._node), asExpr($3._node))
    }
    | term RSHIFTEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, ">>", asExpr($1._node), asExpr($3._node))
    }
    | expr10
    ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      {
        $$._node = ast.NewCondExprNode($1._token.location, asExpr($1._node), asExpr($3._node), asExpr($5._node))
      }
      ;

expr9: expr8
     | expr9 OROR expr8
     {
       $$._node = ast.NewLogicalOrNode($1._token.location, asExpr($1._node), asExpr($3._node))
     }
     ;

expr8: expr7
     | expr8 ANDAND expr7
     {
       $$._node = ast.NewLogicalAndNode($1._token.location, asExpr($1._node), asExpr($3._node))
     }
     ;

expr7: expr6
     | expr7 '>' expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, ">", asExpr($1._node), asExpr($3._node))
     }
     | expr7 '<' expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "<", asExpr($1._node), asExpr($3._node))
     }
     | expr7 GTEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, ">=", asExpr($1._node), asExpr($3._node))
     }
     | expr7 LTEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "<=", asExpr($1._node), asExpr($3._node))
     }
     | expr7 EQEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "==", asExpr($1._node), asExpr($3._node))
     }
     | expr7 NEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "!=", asExpr($1._node), asExpr($3._node))
     }
     ;

expr6: expr5
     | expr6 '|' expr5
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "|", asExpr($1._node), asExpr($3._node))
     }
     ;

expr5: expr4
     | expr5 '^' expr4
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "^", asExpr($1._node), asExpr($3._node))
     }
     ;

expr4: expr3
     | expr4 '&' expr3
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "&", asExpr($1._node), asExpr($3._node))
     }
     ;

expr3: expr2
     | expr3 RSHIFT expr2
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, ">>", asExpr($1._node), asExpr($3._node))
     }
     | expr3 LSHIFT expr2
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "<<", asExpr($1._node), asExpr($3._node))
     }
     ;

expr2: expr1
     | expr2 '+' expr1
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "+", asExpr($1._node), asExpr($3._node))
     }
     | expr2 '-' expr1
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "-", asExpr($1._node), asExpr($3._node))
     }
     ;

expr1: term
     | expr1 '*' term
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "*", asExpr($1._node), asExpr($3._node))
     }
     | expr1 '/' term
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "/", asExpr($1._node), asExpr($3._node))
     }
     | expr1 '%' term
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "%", asExpr($1._node), asExpr($3._node))
     }
     ;

term: unary
    ;

unary: PLUSPLUS unary
     {
       $$._node = ast.NewPrefixOpNode($1._token.location, "++", asExpr($2._node))
     }
     | MINUSMINUS unary
     {
       $$._node = ast.NewPrefixOpNode($1._token.location, "--", asExpr($2._node))
     }
     | '+' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "+", asExpr($2._node))
     }
     | '-' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "-", asExpr($2._node))
     }
     | '!' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "!", asExpr($2._node))
     }
     | '~' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "~", asExpr($2._node))
     }
     | SIZEOF '(' type ')'
     {
       $$._node = ast.NewSizeofTypeNode($1._token.location, asType($3._node), typesys.NewIntegerTypeRef($1._token.location, "unsigned long"))
     }
     | SIZEOF unary
     {
       $$._node = ast.NewSizeofExprNode($1._token.location, asExpr($2._node), typesys.NewIntegerTypeRef($1._token.location, "unsigned long"))
     }
     | postfix
     ;

postfix: primary
       | primary PLUSPLUS
       {
         $$._node = ast.NewSuffixOpNode($1._token.location, "++", asExpr($1._node))
       }
       | primary MINUSMINUS
       {
         $$._node = ast.NewSuffixOpNode($1._token.location, "--", asExpr($1._node))
       }
       | primary '(' args ')'
       {
         $$._node = ast.NewFuncallNode($1._token.location, asExpr($1._node), asExprs($3._nodes))
       }
       ;

name: IDENTIFIER
    ;

args:
    {
      $$._nodes = []duck.INode { }
    }
    | expr
    {
      $$._nodes = append($1._nodes, $1._node)
    }
    | args ',' expr
    {
      $$._nodes = append($1._nodes, $3._node)
    }
    ;

primary: INTEGER
       {
         $$._node = ast.NewIntegerLiteralNode($1._token.location, $1._token.literal)
       }
       | CHARACTER
       {
         // TODO: decode character literal
         $$._node = ast.NewIntegerLiteralNode($1._token.location, $1._token.literal)
       }
       | STRING
       {
         $$._node = ast.NewStringLiteralNode($1._token.location, $1._token.literal)
       }
       | IDENTIFIER
       {
         $$._node = ast.NewVariableNode($1._token.location, $1._token.literal)
       }
       | '(' expr ')'
       {
         $$._node = asExpr($2._node)
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
