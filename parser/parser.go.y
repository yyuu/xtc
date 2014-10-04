%{
package parser

import (
  "errors"
  "fmt"
  "strconv"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
)
%}

%union {
  _token *token

  _node bs_core.INode
  _nodes []bs_core.INode

  _entity bs_core.IEntity
  _entities []bs_core.IEntity

  _typeref bs_core.ITypeRef
  _typerefs []bs_core.ITypeRef
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
                  decl := bs_ast.AsDeclaration($2._node)
                  for i := range $1._nodes {
                    decl.AddDeclaration(bs_ast.AsDeclaration($1._nodes[i]))
                  }
                  lex.ast = bs_ast.NewAST(bs_core.NewLocation(lex.sourceName, 1, 1), decl)
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
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(),
            bs_entity.NewUndefinedVariables(),
            bs_entity.NewDefinedFunctions(bs_entity.AsDefinedFunction($1._entity)),
            bs_entity.NewUndefinedFunctions(),
            bs_entity.NewConstants(),
            bs_ast.NewStructNodes(),
            bs_ast.NewUnionNodes(),
            bs_ast.NewTypedefNodes(),
          )
        }
        | funcdecl
        {
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(),
            bs_entity.NewUndefinedVariables(),
            bs_entity.NewDefinedFunctions(),
            bs_entity.NewUndefinedFunctions(bs_entity.AsUndefinedFunction($1._entity)),
            bs_entity.NewConstants(),
            bs_ast.NewStructNodes(),
            bs_ast.NewUnionNodes(),
            bs_ast.NewTypedefNodes(),
          )
        }
        | defvars
        {
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(bs_entity.AsDefinedVariable($1._entity)),
            bs_entity.NewUndefinedVariables(),
            bs_entity.NewDefinedFunctions(),
            bs_entity.NewUndefinedFunctions(),
            bs_entity.NewConstants(),
            bs_ast.NewStructNodes(),
            bs_ast.NewUnionNodes(),
            bs_ast.NewTypedefNodes(),
          )
        }
        | vardecl
        {
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(),
            bs_entity.NewUndefinedVariables(bs_entity.AsUndefinedVariable($1._entity)),
            bs_entity.NewDefinedFunctions(),
            bs_entity.NewUndefinedFunctions(),
            bs_entity.NewConstants(),
            bs_ast.NewStructNodes(),
            bs_ast.NewUnionNodes(),
            bs_ast.NewTypedefNodes(),
          )
        }
        | defconst
        {
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(),
            bs_entity.NewUndefinedVariables(),
            bs_entity.NewDefinedFunctions(),
            bs_entity.NewUndefinedFunctions(),
            bs_entity.NewConstants(bs_entity.AsConstant($1._entity)),
            bs_ast.NewStructNodes(),
            bs_ast.NewUnionNodes(),
            bs_ast.NewTypedefNodes(),
          )
        }
        | defstruct
        {
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(),
            bs_entity.NewUndefinedVariables(),
            bs_entity.NewDefinedFunctions(),
            bs_entity.NewUndefinedFunctions(),
            bs_entity.NewConstants(),
            bs_ast.NewStructNodes(bs_ast.AsStructNode($1._node)),
            bs_ast.NewUnionNodes(),
            bs_ast.NewTypedefNodes(),
          )
        }
        | defunion
        {
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(),
            bs_entity.NewUndefinedVariables(),
            bs_entity.NewDefinedFunctions(),
            bs_entity.NewUndefinedFunctions(),
            bs_entity.NewConstants(),
            bs_ast.NewStructNodes(),
            bs_ast.NewUnionNodes(bs_ast.AsUnionNode($1._node)),
            bs_ast.NewTypedefNodes(),
          )
        }
        | typedef
        {
          $$._node = bs_ast.NewDeclaration(
            bs_entity.NewDefinedVariables(),
            bs_entity.NewUndefinedVariables(),
            bs_entity.NewDefinedFunctions(),
            bs_entity.NewUndefinedFunctions(),
            bs_entity.NewConstants(),
            bs_ast.NewStructNodes(),
            bs_ast.NewUnionNodes(),
            bs_ast.NewTypedefNodes(bs_ast.AsTypedefNode($1._node)),
          )
        }
        | top_defs defun
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddDefun(bs_entity.AsDefinedFunction($2._entity))
          $$._node = decl
        }
        | top_defs funcdecl
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddFuncdecl(bs_entity.AsUndefinedFunction($2._entity))
          $$._node = decl
        }
        | top_defs defvars
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddDefvar(bs_entity.AsDefinedVariable($2._entity))
          $$._node = decl
        }
        | top_defs vardecl
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddVardecl(bs_entity.AsUndefinedVariable($2._entity))
          $$._node = decl
        }
        | top_defs defconst
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddConstant(bs_entity.AsConstant($2._entity))
          $$._node = decl
        }
        | top_defs defstruct
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddDefstruct(bs_ast.AsStructNode($2._node))
          $$._node = decl
        }
        | top_defs defunion
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddDefunion(bs_ast.AsUnionNode($2._node))
          $$._node = decl
        }
        | top_defs typedef
        {
          decl := bs_ast.AsDeclaration($1._node)
          decl.AddTypedef(bs_ast.AsTypedefNode($2._node))
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

defvars: typeref_name '=' expr ';'
       {
         ref := $1._typeref
         $$._entity = bs_entity.NewDefinedVariable(false, bs_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, bs_ast.AsExprNode($3._node))
       }
       | static_typeref_name '=' expr ';'
       {
         ref := $1._typeref
         $$._entity = bs_entity.NewDefinedVariable(true, bs_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, bs_ast.AsExprNode($3._node))
       }
       ;

defconst: CONST typeref_name '=' expr ';'
        {
          ref := $2._typeref
          $$._entity = bs_entity.NewConstant(bs_ast.NewTypeNode(ref.GetLocation(), ref), $2._token.literal, bs_ast.AsExprNode($4._node))
        }
        ;

defun: typeref_name '(' ')' block
     {
       ps := bs_entity.NewParams($2._token.location, bs_entity.NewParameters(), false)
       t := bs_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = bs_entity.NewDefinedFunction(false, bs_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, bs_ast.AsStmtNode($4._node))
     }
     | typeref_name '(' params ')' block
     {
       ps := bs_entity.AsParams($3._entity)
       t := bs_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = bs_entity.NewDefinedFunction(false, bs_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, bs_ast.AsStmtNode($5._node))
     }
     | static_typeref_name '(' ')' block
     {
       ps := bs_entity.NewParams($2._token.location, bs_entity.NewParameters(), false)
       t := bs_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = bs_entity.NewDefinedFunction(true, bs_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, bs_ast.AsStmtNode($4._node))
     }
     | static_typeref_name '(' params ')' block
     {
       ps := bs_entity.AsParams($3._entity)
       t := bs_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
       $$._entity = bs_entity.NewDefinedFunction(true, bs_ast.NewTypeNode(t.GetLocation(), t), $1._token.literal, ps, bs_ast.AsStmtNode($5._node))
     }
     ;

params: fixedparams
      {
        $$._entity = bs_entity.NewParams($1._token.location, bs_entity.AsParams($1._entity).GetParamDescs(), false)
      }
      | fixedparams ',' DOTDOTDOT
      {
        $$._entity = bs_entity.NewParams($1._token.location, bs_entity.AsParams($1._entity).GetParamDescs(), true)
      }
      ;

fixedparams: param
           {
             $$._entity = bs_entity.NewParams($1._token.location, bs_entity.NewParameters(bs_entity.AsParameter($1._entity)), false)
           }
           | fixedparams ',' param
           {
             $$._entity = bs_entity.NewParams($1._token.location, append(bs_entity.AsParams($1._entity).GetParamDescs(), bs_entity.AsParameter($3._entity)), false)
           }
           ;

param: type name
     {
       $$._entity = bs_entity.NewParameter(bs_ast.AsTypeNode($1._node), $2._token.literal)
     }
     ;

block: '{' defvar_list stmts '}'
     {
       $$._node = bs_ast.NewBlockNode($1._token.location, bs_entity.AsDefinedVariables($2._entities), bs_ast.AsStmtNodes($3._nodes))
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
           $$._node = bs_ast.NewStructNode($1._token.location, bs_typesys.NewStructTypeRef($1._token.location, $2._token.literal), $2._token.literal, bs_ast.AsSlots($3._nodes))
         }
         ;

defunion: UNION name member_list ';'
        {
          $$._node = bs_ast.NewUnionNode($1._token.location, bs_typesys.NewUnionTypeRef($1._token.location, $2._token.literal), $2._token.literal, bs_ast.AsSlots($3._nodes))
        }
        ;

member_list: '{' member_list_body '}'
           {
             $$._nodes = $2._nodes
           }
           ;

member_list_body: slot ';'
                {
                  $$._nodes = bs_ast.NewNodes($1._node)
                }
                | member_list_body slot ';'
                {
                  $$._nodes = append($1._nodes, $2._node)
                }
                ;

slot: type name
    {
      $$._node = bs_ast.NewSlot(bs_ast.AsTypeNode($1._node), $2._token.literal)
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
          ps := bs_entity.NewParams($1._typeref.GetLocation(), bs_entity.NewParameters(), false)
          ref := bs_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
          $$._entity = bs_entity.NewUndefinedFunction(bs_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, ps)
        }
        | extern_typeref_name '(' params ')' ';'
        {
          ps := bs_entity.AsParams($3._entity)
          ref := bs_typesys.NewFunctionTypeRef($1._typeref, parametersTypeRef(ps))
          $$._entity = bs_entity.NewUndefinedFunction(bs_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal, ps)
        }
        ;

vardecl: extern_typeref_name ';'
       {
         ref := $1._typeref
         $$._entity = bs_entity.NewUndefinedVariable(bs_ast.NewTypeNode(ref.GetLocation(), ref), $1._token.literal)
       }
       ;

type: typeref
    {
      $$._node = bs_ast.NewTypeNode($1._token.location, $1._typeref)
    }
    ;

typeref: VOID
       {
         $$._typeref = bs_typesys.NewVoidTypeRef($1._token.location)
       }
       | CHAR
       {
         $$._typeref = bs_typesys.NewCharTypeRef($1._token.location)
       }
       | SHORT
       {
         $$._typeref = bs_typesys.NewShortTypeRef($1._token.location)
       }
       | INT
       {
         $$._typeref = bs_typesys.NewIntTypeRef($1._token.location)
       }
       | LONG
       {
         $$._typeref = bs_typesys.NewLongTypeRef($1._token.location)
       }
       | UNSIGNED CHAR
       {
         $$._typeref = bs_typesys.NewUnsignedIntTypeRef($1._token.location)
       }
       | UNSIGNED SHORT
       {
         $$._typeref = bs_typesys.NewUnsignedShortTypeRef($1._token.location)
       }
       | UNSIGNED INT
       {
         $$._typeref = bs_typesys.NewUnsignedIntTypeRef($1._token.location)
       }
       | UNSIGNED LONG
       {
         $$._typeref = bs_typesys.NewUnsignedLongTypeRef($1._token.location)
       }
       | STRUCT IDENTIFIER
       {
         $$._typeref = bs_typesys.NewStructTypeRef($1._token.location, $2._token.literal)
       }
       | UNION IDENTIFIER
       {
         $$._typeref = bs_typesys.NewUnionTypeRef($1._token.location, $2._token.literal)
       }
       | typeref '[' ']'
       {
         $$._typeref = bs_typesys.NewArrayTypeRef($1._typeref, 0)
       }
       | typeref '[' INTEGER ']'
       {
         n, _ := strconv.Atoi($3._token.literal)
         $$._typeref = bs_typesys.NewArrayTypeRef($1._typeref, n)
       }
       | typeref '*'
       {
         $$._typeref = bs_typesys.NewPointerTypeRef($1._typeref)
       }
       | typeref '(' ')'
       {
         $$._typeref = bs_typesys.NewFunctionTypeRef($1._typeref, bs_typesys.NewParamTypeRefs($2._token.location, bs_typesys.NewTypeRefs(), false))
       }
       | typeref '(' param_typerefs ')'
       {
         $$._typeref = bs_typesys.NewFunctionTypeRef($1._typeref, $3._typeref)
       }
       | TYPENAME
       {
         $$._typeref = bs_typesys.NewUserTypeRef($1._token.location, $1._token.literal)
       }
       ;

param_typerefs: typeref
              {
                $$._typerefs = bs_typesys.NewTypeRefs($1._typeref)
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
         $$._node = bs_ast.NewTypedefNode($1._token.location, $2._typeref, $3._token.literal)
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
      $$._node = bs_ast.NewExprStmtNode($1._token.location, bs_ast.AsExprNode($1._node))
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
              $$._node = bs_ast.NewLabelNode($1._token.location, $1._token.literal, bs_ast.AsStmtNode($3._node))
            }
            ;

if_stmt: IF '(' expr ')' '{' defvar_list stmts '}'
       {
         thenBody := bs_ast.NewBlockNode($5._token.location, bs_entity.AsDefinedVariables($6._entities), bs_ast.AsStmtNodes($7._nodes))
         $$._node = bs_ast.NewIfNode($1._token.location, bs_ast.AsExprNode($3._node), thenBody, nil)
       }
       | IF '(' expr ')' '{' defvar_list stmts '}' ELSE '{' defvar_list stmts '}'
       {
         thenBody := bs_ast.NewBlockNode($5._token.location, bs_entity.AsDefinedVariables($6._entities), bs_ast.AsStmtNodes($7._nodes))
         elseBody := bs_ast.NewBlockNode($10._token.location, bs_entity.AsDefinedVariables($11._entities), bs_ast.AsStmtNodes($12._nodes))
         $$._node = bs_ast.NewIfNode($1._token.location, bs_ast.AsExprNode($3._node), thenBody, elseBody)
       }
       ;

while_stmt: WHILE '(' expr ')' stmt
          {
            $$._node = bs_ast.NewWhileNode($1._token.location, bs_ast.AsExprNode($3._node), bs_ast.AsStmtNode($5._node))
          }
          ;

dowhile_stmt: DO stmt WHILE '(' expr ')' ';'
            {
              $$._node = bs_ast.NewDoWhileNode($1._token.location, bs_ast.AsStmtNode($2._node), bs_ast.AsExprNode($5._node))
            }
            ;

for_stmt: FOR '(' expr ';' expr ';' expr ')' stmt
        {
          $$._node = bs_ast.NewForNode($1._token.location, bs_ast.AsExprNode($3._node), bs_ast.AsExprNode($5._node), bs_ast.AsExprNode($7._node), bs_ast.AsStmtNode($9._node))
        }
        ;

switch_stmt: SWITCH '(' expr ')' '{' case_clauses '}'
           {
             $$._node = bs_ast.NewSwitchNode($1._token.location, bs_ast.AsExprNode($3._node), bs_ast.AsStmtNodes($6._nodes))
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
             $$._node = bs_ast.NewCaseNode($1._token.location, bs_ast.AsExprNodes($1._nodes), bs_ast.AsStmtNode($2._node))
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
                $$._node = bs_ast.NewCaseNode($1._token.location, bs_ast.NewExprNodes(), bs_ast.AsStmtNode($3._node))
              }
              ;

case_body: stmt

goto_stmt: GOTO IDENTIFIER ';'
         {
           $$._node = bs_ast.NewGotoNode($1._token.location, $2._token.literal)
         }
         ;


break_stmt: BREAK ';'
          {
            $$._node = bs_ast.NewBreakNode($1._token.location)
          }
          ;

continue_stmt: CONTINUE ';'
             {
               $$._node = bs_ast.NewContinueNode($1._token.location)
             }
             ;

return_stmt: RETURN expr ';'
           {
             $$._node = bs_ast.NewReturnNode($1._token.location, bs_ast.AsExprNode($2._node))
           }
           ;

expr: term '=' expr
    {
      $$._node = bs_ast.NewAssignNode($1._token.location, bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term PLUSEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "+", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term MINUSEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "-", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term MULEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "*", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term DIVEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "/", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term MODEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "%", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term ANDEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "&", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term OREQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "|", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term XOREQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "^", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term LSHIFTEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, "<<", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | term RSHIFTEQ expr
    {
      $$._node = bs_ast.NewOpAssignNode($1._token.location, ">>", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
    }
    | expr10
    ;

expr10: expr9
      | expr9 '?' expr ':' expr10
      {
        $$._node = bs_ast.NewCondExprNode($1._token.location, bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node), bs_ast.AsExprNode($5._node))
      }
      ;

expr9: expr8
     | expr9 OROR expr8
     {
       $$._node = bs_ast.NewLogicalOrNode($1._token.location, bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr8: expr7
     | expr8 ANDAND expr7
     {
       $$._node = bs_ast.NewLogicalAndNode($1._token.location, bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr7: expr6
     | expr7 '>' expr6
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, ">", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr7 '<' expr6
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "<", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr7 GTEQ expr6
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, ">=", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr7 LTEQ expr6
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "<=", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr7 EQEQ expr6
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "==", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr7 NEQ expr6
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "!=", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr6: expr5
     | expr6 '|' expr5
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "|", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr5: expr4
     | expr5 '^' expr4
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "^", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr4: expr3
     | expr4 '&' expr3
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "&", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr3: expr2
     | expr3 RSHIFT expr2
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, ">>", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr3 LSHIFT expr2
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "<<", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr2: expr1
     | expr2 '+' expr1
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "+", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr2 '-' expr1
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "-", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

expr1: term
     | expr1 '*' term
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "*", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr1 '/' term
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "/", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     | expr1 '%' term
     {
       $$._node = bs_ast.NewBinaryOpNode($1._token.location, "%", bs_ast.AsExprNode($1._node), bs_ast.AsExprNode($3._node))
     }
     ;

term: unary
    ;

unary: PLUSPLUS unary
     {
       $$._node = bs_ast.NewPrefixOpNode($1._token.location, "++", bs_ast.AsExprNode($2._node))
     }
     | MINUSMINUS unary
     {
       $$._node = bs_ast.NewPrefixOpNode($1._token.location, "--", bs_ast.AsExprNode($2._node))
     }
     | '+' term
     {
       $$._node = bs_ast.NewUnaryOpNode($1._token.location, "+", bs_ast.AsExprNode($2._node))
     }
     | '-' term
     {
       $$._node = bs_ast.NewUnaryOpNode($1._token.location, "-", bs_ast.AsExprNode($2._node))
     }
     | '!' term
     {
       $$._node = bs_ast.NewUnaryOpNode($1._token.location, "!", bs_ast.AsExprNode($2._node))
     }
     | '~' term
     {
       $$._node = bs_ast.NewUnaryOpNode($1._token.location, "~", bs_ast.AsExprNode($2._node))
     }
     | SIZEOF '(' type ')'
     {
       $$._node = bs_ast.NewSizeofTypeNode($1._token.location, bs_ast.AsTypeNode($3._node), bs_typesys.NewUnsignedLongTypeRef($1._token.location))
     }
     | SIZEOF unary
     {
       $$._node = bs_ast.NewSizeofExprNode($1._token.location, bs_ast.AsExprNode($2._node), bs_typesys.NewUnsignedLongTypeRef($1._token.location))
     }
     | postfix
     ;

postfix: primary
       | primary PLUSPLUS
       {
         $$._node = bs_ast.NewSuffixOpNode($1._token.location, "++", bs_ast.AsExprNode($1._node))
       }
       | primary MINUSMINUS
       {
         $$._node = bs_ast.NewSuffixOpNode($1._token.location, "--", bs_ast.AsExprNode($1._node))
       }
       | primary '(' ')'
       {
         $$._node = bs_ast.NewFuncallNode($1._token.location, bs_ast.AsExprNode($1._node), bs_ast.NewExprNodes())
       }
       | primary '(' args ')'
       {
         $$._node = bs_ast.NewFuncallNode($1._token.location, bs_ast.AsExprNode($1._node), bs_ast.AsExprNodes($3._nodes))
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
         $$._node = bs_ast.NewIntegerLiteralNode($1._token.location, $1._token.literal)
       }
       | CHARACTER
       {
         $$._node = bs_ast.NewCharacterLiteralNode($1._token.location, $1._token.literal)
       }
       | STRING
       {
         $$._node = bs_ast.NewStringLiteralNode($1._token.location, $1._token.literal)
       }
       | IDENTIFIER
       {
         $$._node = bs_ast.NewVariableNode($1._token.location, $1._token.literal)
       }
       | '(' expr ')'
       {
         $$._node = bs_ast.AsExprNode($2._node)
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

func ParseExpr(s string, errorHandler *bs_core.ErrorHandler, options *bs_core.Options) (*bs_ast.AST, error) {
  src, err := bs_core.NewTemporarySourceFile("", bs_core.EXT_PROGRAM_SOURCE, []byte(s))
  if err != nil {
    return nil, err
  }
  defer func() {
    src.Remove()
  }()
  return Parse(src, errorHandler, options)
}

func ParseFile(path string, errorHandler *bs_core.ErrorHandler, options *bs_core.Options) (*bs_ast.AST, error) {
  src := bs_core.NewSourceFile(path, path, bs_core.EXT_PROGRAM_SOURCE)
  return Parse(src, errorHandler, options)
}

func Parse(src *bs_core.SourceFile, errorHandler *bs_core.ErrorHandler, options *bs_core.Options) (*bs_ast.AST, error) {
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

func parametersTypeRef(params *bs_entity.Params) *bs_typesys.ParamTypeRefs {
  paramDescs := params.GetParamDescs()
  newParamDescs := make([]bs_core.ITypeRef, len(paramDescs))
  for i := range paramDescs {
    newParamDescs[i] = paramDescs[i].GetTypeNode().GetTypeRef()
  }
  return bs_typesys.NewParamTypeRefs(params.GetLocation(), newParamDescs, params.IsVararg())
}
