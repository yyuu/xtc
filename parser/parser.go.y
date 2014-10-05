%{
package parser

import (
  "errors"
  "fmt"
  "strconv"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
)
%}

%union {
  _token *token

  _node xtc_core.INode
  _nodes []xtc_core.INode

  _entity xtc_core.IEntity
  _entities []xtc_core.IEntity

  _typeref xtc_core.ITypeRef
  _typerefs []xtc_core.ITypeRef
}

%token EOF
%token SPACES
%token BLOCK_COMMENT
%token LINE_COMMENT
%token IDENTIFIER
%token INTEGER
%token CHARACTER
%token STRING
%token TYPENAME

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
                  lex := yylex.(*lexer)
                  decl := xtc_ast.AsDeclaration($2._node)
                  for i := range $1._nodes {
                    decl.AddDeclaration(xtc_ast.AsDeclaration($1._nodes[i]))
                  }
                  lex.ast = xtc_ast.NewAST(xtc_core.NewLocation(lex.sourceName, 1, 1), decl)
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
             lex := yylex.(*lexer)
             $$._node = lex.loadLibrary($2._token.literal)
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
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(),
            xtc_entity.NewUndefinedVariables(),
            xtc_entity.NewDefinedFunctions(xtc_entity.AsDefinedFunction($1._entity)),
            xtc_entity.NewUndefinedFunctions(),
            xtc_entity.NewConstants(),
            xtc_ast.NewStructNodes(),
            xtc_ast.NewUnionNodes(),
            xtc_ast.NewTypedefNodes(),
          )
        }
        | funcdecl
        {
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(),
            xtc_entity.NewUndefinedVariables(),
            xtc_entity.NewDefinedFunctions(),
            xtc_entity.NewUndefinedFunctions(xtc_entity.AsUndefinedFunction($1._entity)),
            xtc_entity.NewConstants(),
            xtc_ast.NewStructNodes(),
            xtc_ast.NewUnionNodes(),
            xtc_ast.NewTypedefNodes(),
          )
        }
        | defvars
        {
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(xtc_entity.AsDefinedVariable($1._entity)),
            xtc_entity.NewUndefinedVariables(),
            xtc_entity.NewDefinedFunctions(),
            xtc_entity.NewUndefinedFunctions(),
            xtc_entity.NewConstants(),
            xtc_ast.NewStructNodes(),
            xtc_ast.NewUnionNodes(),
            xtc_ast.NewTypedefNodes(),
          )
        }
        | vardecl
        {
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(),
            xtc_entity.NewUndefinedVariables(xtc_entity.AsUndefinedVariable($1._entity)),
            xtc_entity.NewDefinedFunctions(),
            xtc_entity.NewUndefinedFunctions(),
            xtc_entity.NewConstants(),
            xtc_ast.NewStructNodes(),
            xtc_ast.NewUnionNodes(),
            xtc_ast.NewTypedefNodes(),
          )
        }
        | defconst
        {
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(),
            xtc_entity.NewUndefinedVariables(),
            xtc_entity.NewDefinedFunctions(),
            xtc_entity.NewUndefinedFunctions(),
            xtc_entity.NewConstants(xtc_entity.AsConstant($1._entity)),
            xtc_ast.NewStructNodes(),
            xtc_ast.NewUnionNodes(),
            xtc_ast.NewTypedefNodes(),
          )
        }
        | defstruct
        {
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(),
            xtc_entity.NewUndefinedVariables(),
            xtc_entity.NewDefinedFunctions(),
            xtc_entity.NewUndefinedFunctions(),
            xtc_entity.NewConstants(),
            xtc_ast.NewStructNodes(xtc_ast.AsStructNode($1._node)),
            xtc_ast.NewUnionNodes(),
            xtc_ast.NewTypedefNodes(),
          )
        }
        | defunion
        {
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(),
            xtc_entity.NewUndefinedVariables(),
            xtc_entity.NewDefinedFunctions(),
            xtc_entity.NewUndefinedFunctions(),
            xtc_entity.NewConstants(),
            xtc_ast.NewStructNodes(),
            xtc_ast.NewUnionNodes(xtc_ast.AsUnionNode($1._node)),
            xtc_ast.NewTypedefNodes(),
          )
        }
        | typedef
        {
          $$._node = xtc_ast.NewDeclaration(
            xtc_entity.NewDefinedVariables(),
            xtc_entity.NewUndefinedVariables(),
            xtc_entity.NewDefinedFunctions(),
            xtc_entity.NewUndefinedFunctions(),
            xtc_entity.NewConstants(),
            xtc_ast.NewStructNodes(),
            xtc_ast.NewUnionNodes(),
            xtc_ast.NewTypedefNodes(xtc_ast.AsTypedefNode($1._node)),
          )
        }
        | top_defs defun
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddDefun(xtc_entity.AsDefinedFunction($2._entity))
          $$._node = decl
        }
        | top_defs funcdecl
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddFuncdecl(xtc_entity.AsUndefinedFunction($2._entity))
          $$._node = decl
        }
        | top_defs defvars
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddDefvar(xtc_entity.AsDefinedVariable($2._entity))
          $$._node = decl
        }
        | top_defs vardecl
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddVardecl(xtc_entity.AsUndefinedVariable($2._entity))
          $$._node = decl
        }
        | top_defs defconst
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddConstant(xtc_entity.AsConstant($2._entity))
          $$._node = decl
        }
        | top_defs defstruct
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddDefstruct(xtc_ast.AsStructNode($2._node))
          $$._node = decl
        }
        | top_defs defunion
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddDefunion(xtc_ast.AsUnionNode($2._node))
          $$._node = decl
        }
        | top_defs typedef
        {
          decl := xtc_ast.AsDeclaration($1._node)
          decl.AddTypedef(xtc_ast.AsTypedefNode($2._node))
          $$._node = decl
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

defvars: typeref_name ';'
       {
         ref := $1._typeref
         $$._entity = xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, nil)
       }
       | typeref_name '=' expr ';'
       {
         ref := $1._typeref
         $$._entity = xtc_entity.NewDefinedVariable(false, xtc_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, xtc_ast.AsExprNode($3._node))
       }
       | static_typeref_name ';'
       {
         ref := $1._typeref
         $$._entity = xtc_entity.NewDefinedVariable(true, xtc_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, nil)
       }
       | static_typeref_name '=' expr ';'
       {
         ref := $1._typeref
         $$._entity = xtc_entity.NewDefinedVariable(true, xtc_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, xtc_ast.AsExprNode($3._node))
       }
       ;

defconst: CONST typeref_name '=' expr ';'
        {
          ref := $2._typeref
          $$._entity = xtc_entity.NewConstant(xtc_ast.NewTypeNode(ref.GetLocation(), ref), $2._token.literal, xtc_ast.AsExprNode($4._node))
        }
        ;

defun: typeref_name '(' ')' block
     {
       ps := xtc_entity.NewParams($2._token.location, xtc_entity.NewParameters(), false)
       t := xtc_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = xtc_entity.NewDefinedFunction(false, xtc_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, xtc_ast.AsStmtNode($4._node))
     }
     | typeref_name '(' params ')' block
     {
       ps := xtc_entity.AsParams($3._entity)
       t := xtc_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = xtc_entity.NewDefinedFunction(false, xtc_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, xtc_ast.AsStmtNode($5._node))
     }
     | static_typeref_name '(' ')' block
     {
       ps := xtc_entity.NewParams($2._token.location, xtc_entity.NewParameters(), false)
       t := xtc_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = xtc_entity.NewDefinedFunction(true, xtc_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, xtc_ast.AsStmtNode($4._node))
     }
     | static_typeref_name '(' params ')' block
     {
       ps := xtc_entity.AsParams($3._entity)
       t := xtc_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = xtc_entity.NewDefinedFunction(true, xtc_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, xtc_ast.AsStmtNode($5._node))
     }
     ;

params: fixedparams
      {
        $$._entity = xtc_entity.NewParams($1._token.location, xtc_entity.AsParams($1._entity).GetParamDescs(), false)
      }
      | fixedparams ',' DOTDOTDOT
      {
        $$._entity = xtc_entity.NewParams($1._token.location, xtc_entity.AsParams($1._entity).GetParamDescs(), true)
      }
      ;

fixedparams: param
           {
             $$._entity = xtc_entity.NewParams($1._token.location, xtc_entity.NewParameters(xtc_entity.AsParameter($1._entity)), false)
           }
           | fixedparams ',' param
           {
             $$._entity = xtc_entity.NewParams($1._token.location, append(xtc_entity.AsParams($1._entity).GetParamDescs(), xtc_entity.AsParameter($3._entity)), false)
           }
           ;

param: type name
     {
       $$._entity = xtc_entity.NewParameter(xtc_ast.AsTypeNode($1._node), $2._token.literal)
     }
     ;

block: '{' defvar_list stmts '}'
     {
       $$._node = xtc_ast.NewBlockNode($1._token.location, xtc_entity.AsDefinedVariables($2._entities), xtc_ast.AsStmtNodes($3._nodes))
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
           $$._node = xtc_ast.NewStructNode($1._token.location, xtc_typesys.NewStructTypeRef($1._token.location, $2._token.literal), $2._token.literal, xtc_ast.AsSlots($3._nodes))
         }
         ;

defunion: UNION name member_list ';'
        {
          $$._node = xtc_ast.NewUnionNode($1._token.location, xtc_typesys.NewUnionTypeRef($1._token.location, $2._token.literal), $2._token.literal, xtc_ast.AsSlots($3._nodes))
        }
        ;

member_list: '{' member_list_body '}'
           {
             $$._nodes = $2._nodes
           }
           ;

member_list_body: slot ';'
                {
                  $$._nodes = xtc_ast.NewNodes($1._node)
                }
                | member_list_body slot ';'
                {
                  $$._nodes = append($1._nodes, $2._node)
                }
                ;

slot: type name
    {
      $$._node = xtc_ast.NewSlot(xtc_ast.AsTypeNode($1._node), $2._token.literal)
    }
    ;

extern_typeref_name: EXTERN typeref_name
              {
                $$._token = $2._token
                $$._typeref = $2._typeref
              }
              ;

funcdecl: extern_typeref_name '(' ')' ';'
        {
          ps := xtc_entity.NewParams($1._typeref.GetLocation(), xtc_entity.NewParameters(), false)
          ref := xtc_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
          $$._entity = xtc_entity.NewUndefinedFunction(xtc_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, ps)
        }
        | extern_typeref_name '(' params ')' ';'
        {
          ps := xtc_entity.AsParams($3._entity)
          ref := xtc_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
          $$._entity = xtc_entity.NewUndefinedFunction(xtc_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, ps)
        }
        ;

vardecl: extern_typeref_name ';'
       {
         ref := $1._typeref
         $$._entity = xtc_entity.NewUndefinedVariable(xtc_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal)
       }
       ;

type: typeref
    {
      $$._node = xtc_ast.NewTypeNode($1._token.location, $1._typeref)
    }
    ;

typeref: VOID
       {
         $$._typeref = xtc_typesys.NewVoidTypeRef($1._token.location)
       }
       | CHAR
       {
         $$._typeref = xtc_typesys.NewCharTypeRef($1._token.location)
       }
       | SHORT
       {
         $$._typeref = xtc_typesys.NewShortTypeRef($1._token.location)
       }
       | INT
       {
         $$._typeref = xtc_typesys.NewIntTypeRef($1._token.location)
       }
       | LONG
       {
         $$._typeref = xtc_typesys.NewLongTypeRef($1._token.location)
       }
       | UNSIGNED CHAR
       {
         $$._typeref = xtc_typesys.NewUnsignedIntTypeRef($1._token.location)
       }
       | UNSIGNED SHORT
       {
         $$._typeref = xtc_typesys.NewUnsignedShortTypeRef($1._token.location)
       }
       | UNSIGNED INT
       {
         $$._typeref = xtc_typesys.NewUnsignedIntTypeRef($1._token.location)
       }
       | UNSIGNED LONG
       {
         $$._typeref = xtc_typesys.NewUnsignedLongTypeRef($1._token.location)
       }
       | STRUCT IDENTIFIER
       {
         $$._typeref = xtc_typesys.NewStructTypeRef($1._token.location, $2._token.literal)
       }
       | UNION IDENTIFIER
       {
         $$._typeref = xtc_typesys.NewUnionTypeRef($1._token.location, $2._token.literal)
       }
       | typeref '[' ']'
       {
         $$._typeref = xtc_typesys.NewArrayTypeRef($1._typeref, 0)
       }
       | typeref '[' INTEGER ']'
       {
         n, _ := strconv.Atoi($3._token.literal)
         $$._typeref = xtc_typesys.NewArrayTypeRef($1._typeref, n)
       }
       | typeref '*'
       {
         $$._typeref = xtc_typesys.NewPointerTypeRef($1._typeref)
       }
       | typeref '(' ')'
       {
         $$._typeref = xtc_typesys.NewFunctionTypeRef($1._typeref, xtc_typesys.NewParamTypeRefs($2._token.location, xtc_typesys.NewTypeRefs(), false))
       }
       | typeref '(' param_typerefs ')'
       {
         $$._typeref = xtc_typesys.NewFunctionTypeRef($1._typeref, $3._typeref)
       }
       | TYPENAME
       {
         $$._typeref = xtc_typesys.NewUserTypeRef($1._token.location, $1._token.literal)
       }
       ;

param_typerefs: typeref
              {
                $$._typerefs = xtc_typesys.NewTypeRefs($1._typeref)
              }
              | param_typerefs ',' typeref
              {
                $$._typerefs = append($1._typerefs, $3._typeref)
              }
              ;

typedef: TYPEDEF typeref IDENTIFIER ';'
       {
         lex := yylex.(*lexer)
         lex.knownTypedefs = append(lex.knownTypedefs, $3._token.literal)
         $$._node = xtc_ast.NewTypedefNode($1._token.location, $2._typeref, $3._token.literal)
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
      $$._node = xtc_ast.NewExprStmtNode($1._token.location, xtc_ast.AsExprNode($1._node))
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
              $$._node = xtc_ast.NewLabelNode($1._token.location, $1._token.literal, xtc_ast.AsStmtNode($3._node))
            }
            ;

if_stmt: IF '(' expr ')' '{' defvar_list stmts '}'
       {
         thenBody := xtc_ast.NewBlockNode($5._token.location, xtc_entity.AsDefinedVariables($6._entities), xtc_ast.AsStmtNodes($7._nodes))
         $$._node = xtc_ast.NewIfNode($1._token.location, xtc_ast.AsExprNode($3._node), thenBody, nil)
       }
       | IF '(' expr ')' '{' defvar_list stmts '}' ELSE '{' defvar_list stmts '}'
       {
         thenBody := xtc_ast.NewBlockNode($5._token.location, xtc_entity.AsDefinedVariables($6._entities), xtc_ast.AsStmtNodes($7._nodes))
         elseBody := xtc_ast.NewBlockNode($10._token.location, xtc_entity.AsDefinedVariables($11._entities), xtc_ast.AsStmtNodes($12._nodes))
         $$._node = xtc_ast.NewIfNode($1._token.location, xtc_ast.AsExprNode($3._node), thenBody, elseBody)
       }
       ;

while_stmt: WHILE '(' expr ')' stmt
          {
            $$._node = xtc_ast.NewWhileNode($1._token.location, xtc_ast.AsExprNode($3._node), xtc_ast.AsStmtNode($5._node))
          }
          ;

dowhile_stmt: DO stmt WHILE '(' expr ')' ';'
            {
              $$._node = xtc_ast.NewDoWhileNode($1._token.location, xtc_ast.AsStmtNode($2._node), xtc_ast.AsExprNode($5._node))
            }
            ;

for_stmt: FOR '(' expr ';' expr ';' expr ')' stmt
        {
          $$._node = xtc_ast.NewForNode($1._token.location, xtc_ast.AsExprNode($3._node), xtc_ast.AsExprNode($5._node), xtc_ast.AsExprNode($7._node), xtc_ast.AsStmtNode($9._node))
        }
        ;

switch_stmt: SWITCH '(' expr ')' '{' case_clauses '}'
           {
             $$._node = xtc_ast.NewSwitchNode($1._token.location, xtc_ast.AsExprNode($3._node), xtc_ast.AsStmtNodes($6._nodes))
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
             $$._node = xtc_ast.NewCaseNode($1._token.location, xtc_ast.AsExprNodes($1._nodes), xtc_ast.AsStmtNode($2._node))
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
                $$._node = xtc_ast.NewCaseNode($1._token.location, xtc_ast.NewExprNodes(), xtc_ast.AsStmtNode($3._node))
              }
              ;

case_body: stmt

goto_stmt: GOTO IDENTIFIER ';'
         {
           $$._node = xtc_ast.NewGotoNode($1._token.location, $2._token.literal)
         }
         ;


break_stmt: BREAK ';'
          {
            $$._node = xtc_ast.NewBreakNode($1._token.location)
          }
          ;

continue_stmt: CONTINUE ';'
             {
               $$._node = xtc_ast.NewContinueNode($1._token.location)
             }
             ;

return_stmt: RETURN expr ';'
           {
             $$._node = xtc_ast.NewReturnNode($1._token.location, xtc_ast.AsExprNode($2._node))
           }
           ;

expr: term '=' expr
    {
      $$._node = xtc_ast.NewAssignNode($1._token.location, xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term PLUSEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "+", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term MINUSEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "-", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term MULEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "*", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term DIVEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "/", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term MODEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "%", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term ANDEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "&", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term OREQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "|", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term XOREQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "^", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term LSHIFTEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, "<<", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | term RSHIFTEQ expr
    {
      $$._node = xtc_ast.NewOpAssignNode($1._token.location, ">>", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
    }
    | expr10
    ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      {
        $$._node = xtc_ast.NewCondExprNode($1._token.location, xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node), xtc_ast.AsExprNode($5._node))
      }
      ;

expr9: expr8
     | expr9 OROR expr8
     {
       $$._node = xtc_ast.NewLogicalOrNode($1._token.location, xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr8: expr7
     | expr8 ANDAND expr7
     {
       $$._node = xtc_ast.NewLogicalAndNode($1._token.location, xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr7: expr6
     | expr7 '>' expr6
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, ">", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr7 '<' expr6
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "<", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr7 GTEQ expr6
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, ">=", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr7 LTEQ expr6
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "<=", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr7 EQEQ expr6
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "==", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr7 NEQ expr6
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "!=", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr6: expr5
     | expr6 '|' expr5
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "|", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr5: expr4
     | expr5 '^' expr4
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "^", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr4: expr3
     | expr4 '&' expr3
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "&", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr3: expr2
     | expr3 RSHIFT expr2
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, ">>", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr3 LSHIFT expr2
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "<<", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr2: expr1
     | expr2 '+' expr1
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "+", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr2 '-' expr1
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "-", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

expr1: term
     | expr1 '*' term
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "*", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr1 '/' term
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "/", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     | expr1 '%' term
     {
       $$._node = xtc_ast.NewBinaryOpNode($1._token.location, "%", xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
     }
     ;

term: unary
    ;

unary: PLUSPLUS unary
     {
       $$._node = xtc_ast.NewPrefixOpNode($1._token.location, "++", xtc_ast.AsExprNode($2._node))
     }
     | MINUSMINUS unary
     {
       $$._node = xtc_ast.NewPrefixOpNode($1._token.location, "--", xtc_ast.AsExprNode($2._node))
     }
     | '+' term
     {
       $$._node = xtc_ast.NewUnaryOpNode($1._token.location, "+", xtc_ast.AsExprNode($2._node))
     }
     | '-' term
     {
       $$._node = xtc_ast.NewUnaryOpNode($1._token.location, "-", xtc_ast.AsExprNode($2._node))
     }
     | '!' term
     {
       $$._node = xtc_ast.NewUnaryOpNode($1._token.location, "!", xtc_ast.AsExprNode($2._node))
     }
     | '~' term
     {
       $$._node = xtc_ast.NewUnaryOpNode($1._token.location, "~", xtc_ast.AsExprNode($2._node))
     }
     | '*' term
     {
       $$._node = xtc_ast.NewDereferenceNode($1._token.location, xtc_ast.AsExprNode($2._node))
     }
     | '&' term
     {
       $$._node = xtc_ast.NewAddressNode($1._token.location, xtc_ast.AsExprNode($2._node))
     }
     | SIZEOF '(' type ')'
     {
       $$._node = xtc_ast.NewSizeofTypeNode($1._token.location, xtc_ast.AsTypeNode($3._node), xtc_typesys.NewUnsignedLongTypeRef($1._token.location))
     }
     | SIZEOF unary
     {
       $$._node = xtc_ast.NewSizeofExprNode($1._token.location, xtc_ast.AsExprNode($2._node), xtc_typesys.NewUnsignedLongTypeRef($1._token.location))
     }
     | postfix
     ;

postfix: primary
       | primary PLUSPLUS
       {
         $$._node = xtc_ast.NewSuffixOpNode($1._token.location, "++", xtc_ast.AsExprNode($1._node))
       }
       | primary MINUSMINUS
       {
         $$._node = xtc_ast.NewSuffixOpNode($1._token.location, "--", xtc_ast.AsExprNode($1._node))
       }
       | primary '[' expr ']'
       {
         $$._node = xtc_ast.NewArefNode($1._token.location, xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNode($3._node))
       }
       | primary '.' name
       {
         $$._node = xtc_ast.NewMemberNode($1._token.location, xtc_ast.AsExprNode($1._node), $3._token.literal)
       }
       | primary ARROW name
       {
         $$._node = xtc_ast.NewPtrMemberNode($1._token.location, xtc_ast.AsExprNode($1._node), $3._token.literal)
       }
       | primary '(' ')'
       {
         $$._node = xtc_ast.NewFuncallNode($1._token.location, xtc_ast.AsExprNode($1._node), xtc_ast.NewExprNodes())
       }
       | primary '(' args ')'
       {
         $$._node = xtc_ast.NewFuncallNode($1._token.location, xtc_ast.AsExprNode($1._node), xtc_ast.AsExprNodes($3._nodes))
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
         $$._node = xtc_ast.NewIntegerLiteralNode($1._token.location, $1._token.literal)
       }
       | CHARACTER
       {
         $$._node = xtc_ast.NewCharacterLiteralNode($1._token.location, $1._token.literal)
       }
       | STRING
       {
         $$._node = xtc_ast.NewStringLiteralNode($1._token.location, $1._token.literal)
       }
       | IDENTIFIER
       {
         $$._node = xtc_ast.NewVariableNode($1._token.location, $1._token.literal)
       }
       | '(' expr ')'
       {
         $$._node = xtc_ast.AsExprNode($2._node)
       }
       ;

%%

func (self *lexer) Lex(lval *yySymType) int {
  t, err := self.getNextToken()
  if err != nil {
    self.Error(err.Error())
  }
  if t == nil {
    return 0
  }
  if self.options.DumpTokens() {
    self.errorHandler.Info(t)
  }
  lval._token = t
  return t.id
}

func (self *lexer) Error(s string) {
  self.errorHandler.Error(s)
}

func ParseExpr(s string, errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options) (*xtc_ast.AST, error) {
  src, err := xtc_core.NewTemporarySourceFile("", xtc_core.EXT_PROGRAM_SOURCE, []byte(s))
  if err != nil {
    return nil, err
  }
  defer func() {
    src.Remove()
  }()
  return Parse(src, errorHandler, options)
}

func ParseFile(path string, errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options) (*xtc_ast.AST, error) {
  src := xtc_core.NewSourceFile(path, path, xtc_core.EXT_PROGRAM_SOURCE)
  return Parse(src, errorHandler, options)
}

func Parse(src *xtc_core.SourceFile, errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options) (*xtc_ast.AST, error) {
  if options.IsVerboseMode() {
    yyDebug = 4 // TODO: configurable
  }
  loader := newLibraryLoader(errorHandler, options)
  bytes, err := src.ReadAll()
  if err != nil {
    return nil, err
  }
  lex := newLexer(src.GetName(), string(bytes), loader, errorHandler, options)
  if yyParse(lex) == 0 {
    return lex.ast, nil // success
  } else {
    if errorHandler.ErrorOccured() {
      return nil, fmt.Errorf("found %d error(s).", errorHandler.GetErrors())
    } else {
      return nil, errors.New("must not happen: lexer error not recorded.")
    }
  }
}

func parametersTypeRef(params *xtc_entity.Params) *xtc_typesys.ParamTypeRefs {
  paramDescs := params.GetParamDescs()
  newParamDescs := make([]xtc_core.ITypeRef, len(paramDescs))
  for i := range paramDescs {
    newParamDescs[i] = paramDescs[i].GetTypeNode().GetTypeRef()
  }
  return xtc_typesys.NewParamTypeRefs(params.GetLocation(), newParamDescs, params.IsVararg())
}
