%{
package parser

import (
  "errors"
  "fmt"
  "strconv"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)
%}

%union {
  _token token

  _node core.INode
  _nodes []core.INode

  _entity core.IEntity
  _entities []core.IEntity

  _typeref core.ITypeRef
  _typerefs []core.ITypeRef
}

%token EOF
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

compilation_unit: EOF
                | import_stmts top_defs EOF
                {
                  if lex, ok := yylex.(*lex); ok {
                    var loc core.Location
                    if lex.firstToken != nil {
                      loc = lex.firstToken.location
                    }
                    lex.ast = ast.NewAST(loc, asDeclarations($2._node))
                  } else {
                    panic("parser is broken")
                  }
                }
                ;

declaration_file: import_stmts
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

top_defs: defun
        {
          $$._node = ast.NewDeclarations(
            entity.NewDefinedVariables(),
            entity.NewUndefinedVariables(),
            entity.NewDefinedFunctions(asDefinedFunction($1._entity)),
            entity.NewUndefinedFunctions(),
            entity.NewConstants(),
            ast.NewStructNodes(),
            ast.NewUnionNodes(),
            ast.NewTypedefNodes(),
          )
        }
        | defvars
        {
          $$._node = ast.NewDeclarations(
            entity.NewDefinedVariables(asDefinedVariable($1._entity)),
            entity.NewUndefinedVariables(),
            entity.NewDefinedFunctions(),
            entity.NewUndefinedFunctions(),
            entity.NewConstants(),
            ast.NewStructNodes(),
            ast.NewUnionNodes(),
            ast.NewTypedefNodes(),
          )
        }
        | defconst
        {
          $$._node = ast.NewDeclarations(
            entity.NewDefinedVariables(),
            entity.NewUndefinedVariables(),
            entity.NewDefinedFunctions(),
            entity.NewUndefinedFunctions(),
            entity.NewConstants(asConstant($1._entity)),
            ast.NewStructNodes(),
            ast.NewUnionNodes(),
            ast.NewTypedefNodes(),
          )
        }
        | defstruct
        {
          $$._node = ast.NewDeclarations(
            entity.NewDefinedVariables(),
            entity.NewUndefinedVariables(),
            entity.NewDefinedFunctions(),
            entity.NewUndefinedFunctions(),
            entity.NewConstants(),
            ast.NewStructNodes(asStructNode($1._node)),
            ast.NewUnionNodes(),
            ast.NewTypedefNodes(),
          )
        }
        | defunion
        {
          $$._node = ast.NewDeclarations(
            entity.NewDefinedVariables(),
            entity.NewUndefinedVariables(),
            entity.NewDefinedFunctions(),
            entity.NewUndefinedFunctions(),
            entity.NewConstants(),
            ast.NewStructNodes(),
            ast.NewUnionNodes(asUnionNode($1._node)),
            ast.NewTypedefNodes(),
          )
        }
        | typedef
        {
          $$._node = ast.NewDeclarations(
            entity.NewDefinedVariables(),
            entity.NewUndefinedVariables(),
            entity.NewDefinedFunctions(),
            entity.NewUndefinedFunctions(),
            entity.NewConstants(),
            ast.NewStructNodes(),
            ast.NewUnionNodes(),
            ast.NewTypedefNodes(asTypedefNode($1._node)),
          )
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

typeref_name: typeref name
            {
              $$._token = $2._token
              $$._typeref = $1._typeref
            }
            ;

static_typeref_name: STATIC typeref_name
                   {
                     $$._token = $2._token
                     $$._typeref = $2._typeref
                   }
                   ;

defvars: typeref_name '=' expr ';'
       {
         ref := $1._typeref
         $$._entity = entity.NewDefinedVariable(true, ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, asExprNode($3._node))
       }
       | static_typeref_name '=' expr ';'
       {
         ref := $1._typeref
         $$._entity = entity.NewDefinedVariable(false, ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, asExprNode($3._node))
       }
       ;

defconst: CONST typeref_name '=' expr ';'
        {
          ref := $2._typeref
          $$._entity = entity.NewConstant(ast.NewTypeNode(ref.GetLocation(), ref), $2._token.literal, asExprNode($4._node))
        }
        ;

defun: typeref_name '(' ')' block
     {
       ps := entity.NewParams($2._token.location, []*entity.Parameter { })
       t := typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = entity.NewDefinedFunction(true,
         ast.NewTypeNode(t.GetLocation(), t),
         $1._token.literal,
         ps,
         asStmtNode($4._node),
       )
     }
     | typeref_name '(' params ')' block
     {
       ps := asParams($3._entity)
       t := typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = entity.NewDefinedFunction(true, ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, asStmtNode($5._node))
     }
     | static_typeref_name '(' ')' block
     {
       ps := entity.NewParams($2._token.location, []*entity.Parameter { })
       t := typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = entity.NewDefinedFunction(false, ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, asStmtNode($4._node))
     }
     | static_typeref_name '(' params ')' block
     {
       ps := asParams($3._entity)
       t := typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = entity.NewDefinedFunction(false, ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, asStmtNode($5._node))
     }
     ;

/*
storage:
       | STATIC
       ;
 */

params: fixedparams
      {
        $$._entity = entity.NewParams($1._token.location, asParams($1._entity).GetParamDescs())
      }
      | fixedparams ',' DOTDOTDOT
      {
        $$._entity = entity.NewParams($1._token.location, asParams($1._entity).GetParamDescs())
//      $$._entity.AcceptVarArgs()
      }
      ;

fixedparams: param
           {
             $$._entity = entity.NewParams($1._token.location, []*entity.Parameter { asParameter($1._entity) })
           }
           | fixedparams ',' param
           {
             $$._entity = entity.NewParams($1._token.location, append(asParams($1._entity).GetParamDescs(), asParameter($3._entity)))
           }
           ;

param: type name
     {
       $$._entity = entity.NewParameter(asTypeNode($1._node), $2._token.literal)
     }
     ;

block: '{' defvar_list stmts '}'
     {
       $$._node = ast.NewBlockNode($1._token.location, asDefinedVariables($2._entities), asStmtNodes($3._nodes))
     }
     ;

defvar_list:
           | defvar_list defvars
           {
             $$._entities = append($1._entities, $2._entity)
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
                  $$._nodes = []core.INode { $1._node }
                }
                | member_list_body slot ';'
                {
                  $$._nodes = append($1._nodes, $2._node)
                }
                ;

slot: type name
    {
      $$._node = ast.NewSlot(asTypeNode($1._node), $2._token.literal)
    }
    ;

extern_typeref_name: EXTERN typeref_name
              {
                $$._token = $1._token
                $$._typeref = $2._typeref
              }
              ;

funcdecl: extern_typeref_name '(' ')' ';'
        {
          ps := entity.NewParams($1._typeref.GetLocation(), []*entity.Parameter { })
          ref := typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
          $$._entity = entity.NewUndefinedFunction(
            ast.NewTypeNode(ref.GetLocation(), ref),
            $1._token.literal,
            ps,
          )
        }
        | extern_typeref_name '(' params ')' ';'
        {
          ps := asParams($3._entity)
          ref := typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
          $$._entity = entity.NewUndefinedFunction(ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, ps)
        }
        ;

vardecl: extern_typeref_name ';'
       {
         ref := $1._typeref
         $$._entity = entity.NewUndefinedVariable(ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal)
       }
       ;

type: typeref
    {
      $$._node = ast.NewTypeNode($1._token.location, $1._typeref)
    }
    ;

typeref: VOID
       {
         $$._typeref = typesys.NewVoidTypeRef($1._token.location)
       }
       | CHAR
       {
         $$._typeref = typesys.NewCharTypeRef($1._token.location)
       }
       | SHORT
       {
         $$._typeref = typesys.NewShortTypeRef($1._token.location)
       }
       | INT
       {
         $$._typeref = typesys.NewIntTypeRef($1._token.location)
       }
       | LONG
       {
         $$._typeref = typesys.NewLongTypeRef($1._token.location)
       }
       | UNSIGNED CHAR
       {
         $$._typeref = typesys.NewUnsignedIntTypeRef($1._token.location)
       }
       | UNSIGNED SHORT
       {
         $$._typeref = typesys.NewUnsignedShortTypeRef($1._token.location)
       }
       | UNSIGNED INT
       {
         $$._typeref = typesys.NewUnsignedIntTypeRef($1._token.location)
       }
       | UNSIGNED LONG
       {
         $$._typeref = typesys.NewUnsignedLongTypeRef($1._token.location)
       }
       | STRUCT IDENTIFIER
       {
         $$._typeref = typesys.NewStructTypeRef($1._token.location, $2._token.literal)
       }
       | UNION IDENTIFIER
       {
         $$._typeref = typesys.NewUnionTypeRef($1._token.location, $2._token.literal)
       }
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
       | typeref '(' ')'
       {
         $$._typeref = typesys.NewFunctionTypeRef($1._typeref, typesys.NewParamTypeRefs($2._token.location, []core.ITypeRef { }, false))
       }
       | typeref '(' param_typerefs ')'
       {
         $$._typeref = typesys.NewFunctionTypeRef($1._typeref, $3._typeref)
       }
       ;

param_typerefs: typeref
               {
                 $$._typerefs = []core.ITypeRef { $1._typeref }
               }
               | param_typerefs ',' typeref
               {
                 $$._typerefs = append($1._typerefs, $3._typeref)
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
      $$._node = ast.NewExprStmtNode($1._token.location, asExprNode($1._node))
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
              $$._node = ast.NewLabelNode($1._token.location, $1._token.literal, asStmtNode($3._node))
            }
            ;

if_stmt: IF '(' expr ')' stmt ELSE stmt
       {
         $$._node = ast.NewIfNode($1._token.location, asExprNode($3._node), asStmtNode($5._node), asStmtNode($7._node))
       }
       ;

while_stmt: WHILE '(' expr ')' stmt
          {
            $$._node = ast.NewWhileNode($1._token.location, asExprNode($3._node), asStmtNode($5._node))
          }
          ;

dowhile_stmt: DO stmt WHILE '(' expr ')' ';'
            {
              $$._node = ast.NewDoWhileNode($1._token.location, asStmtNode($2._node), asExprNode($5._node))
            }
            ;

for_stmt: FOR '(' expr ';' expr ';' expr ')' stmt
        {
          $$._node = ast.NewForNode($1._token.location, asExprNode($3._node), asExprNode($5._node), asExprNode($7._node), asStmtNode($9._node))
        }
        ;

switch_stmt: SWITCH '(' expr ')' '{' case_clauses '}'
           {
             $$._node = ast.NewSwitchNode($1._token.location, asExprNode($3._node), asStmtNodes($6._nodes))
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
             $$._node = ast.NewCaseNode($1._token.location, asExprNodes($1._nodes), asStmtNode($2._node))
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
                $$._node = ast.NewCaseNode($1._token.location, []core.IExprNode { }, asStmtNode($3._node))
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
             $$._node = ast.NewReturnNode($1._token.location, asExprNode($2._node))
           }
           ;

expr: term '=' expr
    {
      $$._node = ast.NewAssignNode($1._token.location, asExprNode($1._node), asExprNode($3._node))
    }
    | term PLUSEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "+", asExprNode($1._node), asExprNode($3._node))
    }
    | term MINUSEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "-", asExprNode($1._node), asExprNode($3._node))
    }
    | term MULEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "*", asExprNode($1._node), asExprNode($3._node))
    }
    | term DIVEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "/", asExprNode($1._node), asExprNode($3._node))
    }
    | term MODEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "%", asExprNode($1._node), asExprNode($3._node))
    }
    | term ANDEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "&", asExprNode($1._node), asExprNode($3._node))
    }
    | term OREQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "|", asExprNode($1._node), asExprNode($3._node))
    }
    | term XOREQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "^", asExprNode($1._node), asExprNode($3._node))
    }
    | term LSHIFTEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, "<<", asExprNode($1._node), asExprNode($3._node))
    }
    | term RSHIFTEQ expr
    {
      $$._node = ast.NewOpAssignNode($1._token.location, ">>", asExprNode($1._node), asExprNode($3._node))
    }
    | expr10
    ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      {
        $$._node = ast.NewCondExprNode($1._token.location, asExprNode($1._node), asExprNode($3._node), asExprNode($5._node))
      }
      ;

expr9: expr8
     | expr9 OROR expr8
     {
       $$._node = ast.NewLogicalOrNode($1._token.location, asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr8: expr7
     | expr8 ANDAND expr7
     {
       $$._node = ast.NewLogicalAndNode($1._token.location, asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr7: expr6
     | expr7 '>' expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, ">", asExprNode($1._node), asExprNode($3._node))
     }
     | expr7 '<' expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "<", asExprNode($1._node), asExprNode($3._node))
     }
     | expr7 GTEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, ">=", asExprNode($1._node), asExprNode($3._node))
     }
     | expr7 LTEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "<=", asExprNode($1._node), asExprNode($3._node))
     }
     | expr7 EQEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "==", asExprNode($1._node), asExprNode($3._node))
     }
     | expr7 NEQ expr6
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "!=", asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr6: expr5
     | expr6 '|' expr5
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "|", asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr5: expr4
     | expr5 '^' expr4
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "^", asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr4: expr3
     | expr4 '&' expr3
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "&", asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr3: expr2
     | expr3 RSHIFT expr2
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, ">>", asExprNode($1._node), asExprNode($3._node))
     }
     | expr3 LSHIFT expr2
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "<<", asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr2: expr1
     | expr2 '+' expr1
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "+", asExprNode($1._node), asExprNode($3._node))
     }
     | expr2 '-' expr1
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "-", asExprNode($1._node), asExprNode($3._node))
     }
     ;

expr1: term
     | expr1 '*' term
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "*", asExprNode($1._node), asExprNode($3._node))
     }
     | expr1 '/' term
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "/", asExprNode($1._node), asExprNode($3._node))
     }
     | expr1 '%' term
     {
       $$._node = ast.NewBinaryOpNode($1._token.location, "%", asExprNode($1._node), asExprNode($3._node))
     }
     ;

term: unary
    ;

unary: PLUSPLUS unary
     {
       $$._node = ast.NewPrefixOpNode($1._token.location, "++", asExprNode($2._node))
     }
     | MINUSMINUS unary
     {
       $$._node = ast.NewPrefixOpNode($1._token.location, "--", asExprNode($2._node))
     }
     | '+' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "+", asExprNode($2._node))
     }
     | '-' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "-", asExprNode($2._node))
     }
     | '!' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "!", asExprNode($2._node))
     }
     | '~' term
     {
       $$._node = ast.NewUnaryOpNode($1._token.location, "~", asExprNode($2._node))
     }
     | SIZEOF '(' type ')'
     {
       $$._node = ast.NewSizeofTypeNode($1._token.location, asTypeNode($3._node), typesys.NewUnsignedLongTypeRef($1._token.location))
     }
     | SIZEOF unary
     {
       $$._node = ast.NewSizeofExprNode($1._token.location, asExprNode($2._node), typesys.NewUnsignedLongTypeRef($1._token.location))
     }
     | postfix
     ;

postfix: primary
       | primary PLUSPLUS
       {
         $$._node = ast.NewSuffixOpNode($1._token.location, "++", asExprNode($1._node))
       }
       | primary MINUSMINUS
       {
         $$._node = ast.NewSuffixOpNode($1._token.location, "--", asExprNode($1._node))
       }
       | primary '(' ')'
       {
         $$._node = ast.NewFuncallNode($1._token.location, asExprNode($1._node), []core.IExprNode { })
       }
       | primary '(' args ')'
       {
         $$._node = ast.NewFuncallNode($1._token.location, asExprNode($1._node), asExprNodes($3._nodes))
       }
       ;

name: IDENTIFIER
    ;

args: expr
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
         $$._node = asExprNode($2._node)
       }
       ;

%%

var Verbose = 0

func (self *lex) Lex(lval *yySymType) int {
  t := self.getNextToken()
  if t == nil {
    if self.isEOF {
      return 0
    } else {
      self.isEOF = true
      return EOF
    }
  } else {
    lval._token = *t
    return t.id
  }
}

func (self *lex) Error(s string) {
  self.error = errors.New(s)
  panic(fmt.Errorf("%s: %s", self, s))
}

func ParseExpr(s string) (*ast.AST, error) {
  yyDebug = Verbose
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
