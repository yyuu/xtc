package ast

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
  "bitbucket.org/yyuu/bs/xt"
)

func TestBlock1(t *testing.T) {
/*
  {
    println("hello, world");
  }
 */
  x := NewBlockNode(
    loc(0,0),
    []*entity.DefinedVariable { },
    []core.IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "println"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"hello, world\"") })),
    },
  )
  s := `{
  "ClassName": "ast.BlockNode",
  "Location": "[:0,0]",
  "Variables": [],
  "Stmts": [
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "println",
          "Entity": null
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*",
              "Type": null
            },
            "Value": "\"hello, world\""
          }
        ]
      }
    }
  ]
}`
  xt.AssertStringEqualsDiff(t, "BlockNode1", xt.JSON(x), s)
}

func TestBlock2(t *testing.T) {
/*
  {
    int n = 12345;
    printf("%d", n);
  }
 */
  x := NewBlockNode(
    loc(0,0),
    []*entity.DefinedVariable {
      entity.NewDefinedVariable(
        true,
        NewTypeNode(loc(0,0), typesys.NewIntTypeRef(loc(0,0))),
        "n",
        NewIntegerLiteralNode(loc(0,0), "12345"),
      ),
    },
    []core.IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "n") })),
    },
  )
  s := `{
  "ClassName": "ast.BlockNode",
  "Location": "[:0,0]",
  "Variables": [
    {
      "ClassName": "entity.DefinedVariable",
      "Private": true,
      "Name": "n",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int",
        "Type": null
      },
      "Initializer": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "TypeNode": {
          "ClassName": "ast.TypeNode",
          "Location": "[:0,0]",
          "TypeRef": "int",
          "Type": null
        },
        "Value": 12345
      }
    }
  ],
  "Stmts": [
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "printf",
          "Entity": null
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*",
              "Type": null
            },
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "n",
            "Entity": null
          }
        ]
      }
    }
  ]
}`
  xt.AssertStringEqualsDiff(t, "BlockNode2", xt.JSON(x), s)
}

func TestBlock3(t *testing.T) {
/*
  {
    int n = 12345;
    int m = 67890;
    printf("%d", n);
    printf("%d", m);
  }
 */
  x := NewBlockNode(
    loc(0,0),
    []*entity.DefinedVariable {
      entity.NewDefinedVariable(
        true,
        NewTypeNode(loc(0,0), typesys.NewIntTypeRef(loc(0,0))),
        "n",
        NewIntegerLiteralNode(loc(0,0), "12345"),
      ),
      entity.NewDefinedVariable(
        true,
        NewTypeNode(loc(0,0), typesys.NewIntTypeRef(loc(0,0))),
        "m",
        NewIntegerLiteralNode(loc(0,0), "67890"),
      ),
    },
    []core.IStmtNode {
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "n") })),
      NewExprStmtNode(loc(0,0), NewFuncallNode(loc(0,0), NewVariableNode(loc(0,0), "printf"), []core.IExprNode { NewStringLiteralNode(loc(0,0), "\"%d\""), NewVariableNode(loc(0,0), "m") })),
    },
  )
  s := `{
  "ClassName": "ast.BlockNode",
  "Location": "[:0,0]",
  "Variables": [
    {
      "ClassName": "entity.DefinedVariable",
      "Private": true,
      "Name": "n",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int",
        "Type": null
      },
      "Initializer": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "TypeNode": {
          "ClassName": "ast.TypeNode",
          "Location": "[:0,0]",
          "TypeRef": "int",
          "Type": null
        },
        "Value": 12345
      }
    },
    {
      "ClassName": "entity.DefinedVariable",
      "Private": true,
      "Name": "m",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int",
        "Type": null
      },
      "Initializer": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "TypeNode": {
          "ClassName": "ast.TypeNode",
          "Location": "[:0,0]",
          "TypeRef": "int",
          "Type": null
        },
        "Value": 67890
      }
    }
  ],
  "Stmts": [
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "printf",
          "Entity": null
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*",
              "Type": null
            },
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "n",
            "Entity": null
          }
        ]
      }
    },
    {
      "ClassName": "ast.ExprStmtNode",
      "Location": "[:0,0]",
      "Expr": {
        "ClassName": "ast.FuncallNode",
        "Location": "[:0,0]",
        "Expr": {
          "ClassName": "ast.VariableNode",
          "Location": "[:0,0]",
          "Name": "printf",
          "Entity": null
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*",
              "Type": null
            },
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "m",
            "Entity": null
          }
        ]
      }
    }
  ]
}`
  xt.AssertStringEqualsDiff(t, "BlockNode3", xt.JSON(x), s)
}
