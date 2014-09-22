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
          "Name": "println"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": {
                "ClassName": "typesys.PointerTypeRef",
                "Location": "[:0,0]",
                "BaseType": {
                  "ClassName": "typesys.IntegerTypeRef",
                  "Location": "[:0,0]",
                  "Name": "char"
                }
              }
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
        "TypeRef": {
          "ClassName": "typesys.IntegerTypeRef",
          "Location": "[:0,0]",
          "Name": "int"
        }
      },
      "Initializer": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "TypeNode": {
          "ClassName": "ast.TypeNode",
          "Location": "[:0,0]",
          "TypeRef": {
            "ClassName": "typesys.IntegerTypeRef",
            "Location": "[:0,0]",
            "Name": "int"
          }
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
          "Name": "printf"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": {
                "ClassName": "typesys.PointerTypeRef",
                "Location": "[:0,0]",
                "BaseType": {
                  "ClassName": "typesys.IntegerTypeRef",
                  "Location": "[:0,0]",
                  "Name": "char"
                }
              }
            },
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "n"
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
        "TypeRef": {
          "ClassName": "typesys.IntegerTypeRef",
          "Location": "[:0,0]",
          "Name": "int"
        }
      },
      "Initializer": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "TypeNode": {
          "ClassName": "ast.TypeNode",
          "Location": "[:0,0]",
          "TypeRef": {
            "ClassName": "typesys.IntegerTypeRef",
            "Location": "[:0,0]",
            "Name": "int"
          }
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
        "TypeRef": {
          "ClassName": "typesys.IntegerTypeRef",
          "Location": "[:0,0]",
          "Name": "int"
        }
      },
      "Initializer": {
        "ClassName": "ast.IntegerLiteralNode",
        "Location": "[:0,0]",
        "TypeNode": {
          "ClassName": "ast.TypeNode",
          "Location": "[:0,0]",
          "TypeRef": {
            "ClassName": "typesys.IntegerTypeRef",
            "Location": "[:0,0]",
            "Name": "int"
          }
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
          "Name": "printf"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": {
                "ClassName": "typesys.PointerTypeRef",
                "Location": "[:0,0]",
                "BaseType": {
                  "ClassName": "typesys.IntegerTypeRef",
                  "Location": "[:0,0]",
                  "Name": "char"
                }
              }
            },
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "n"
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
          "Name": "printf"
        },
        "Args": [
          {
            "ClassName": "ast.StringLiteralNode",
            "Location": "[:0,0]",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": {
                "ClassName": "typesys.PointerTypeRef",
                "Location": "[:0,0]",
                "BaseType": {
                  "ClassName": "typesys.IntegerTypeRef",
                  "Location": "[:0,0]",
                  "Name": "char"
                }
              }
            },
            "Value": "\"%d\""
          },
          {
            "ClassName": "ast.VariableNode",
            "Location": "[:0,0]",
            "Name": "m"
          }
        ]
      }
    }
  ]
}`
  xt.AssertStringEqualsDiff(t, "BlockNode3", xt.JSON(x), s)
}
