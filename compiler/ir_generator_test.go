package compiler

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_ir "bitbucket.org/yyuu/xtc/ir"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func setupIRGenerator(ast *xtc_ast.AST, table *xtc_typesys.TypeTable) (*xtc_ir.IR, error) {
  errorHandler := xtc_core.NewErrorHandler(xtc_core.LOG_WARN)
  options := xtc_core.NewOptions("type_checker_test.go")

  localResolver := NewLocalResolver(errorHandler, options)
  ast2, err := localResolver.Resolve(ast)
  if err != nil {
    panic("must not happen: test data is broken")
  }

  typeResolver := NewTypeResolver(errorHandler, options, table)
  ast3, err := typeResolver.Resolve(ast2)
  if err != nil {
    panic("must not happen: test data is broken")
  }

  generator := NewIRGenerator(errorHandler, options, table)
  return generator.Generate(ast3)
}

func TestIRGeneratorEmpty(t *testing.T) {
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  str := `{
  "ClassName": "ir.IR",
  "Location": "[:0,0]",
  "Defvars": [],
  "Defuns": [],
  "Funcdecls": []
}`
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  ir, err := setupIRGenerator(ast, table)
  xt.AssertNil(t, "should not be failed", err)
  xt.AssertStringEqualsDiff(t, "should return empty IR", xt.JSON(ir), str)
}

func TestIRGeneratorEmptyFunction(t *testing.T) {
/*
  void f() {
    return;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc, []xtc_core.ITypeRef { }, false),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, nil),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  str := `{
  "ClassName": "ir.IR",
  "Location": "[:0,0]",
  "Defvars": [],
  "Defuns": [
    {
      "ClassName": "entity.DefinedFunction",
      "Private": false,
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "void()",
        "Type": "void()"
      },
      "Name": "f",
      "Params": {
        "ClassName": "entity.Params",
        "Location": "[:0,0]",
        "ParamDescs": [],
        "Vararg": false
      },
      "Body": {
        "ClassName": "ast.BlockNode",
        "Location": "[:0,0]",
        "Variables": [],
        "Stmts": [
          {
            "ClassName": "ast.ReturnNode",
            "Location": "[:0,0]",
            "Expr": null
          }
        ]
      },
      "IR": [
        {
          "ClassName": "ir.Return",
          "Location": "[:0,0]",
          "Expr": null
        }
      ]
    }
  ],
  "Funcdecls": []
}`
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  ir, err := setupIRGenerator(ast, table)
  xt.AssertNil(t, "should not be failed", err)
  xt.AssertStringEqualsDiff(t, "should return IR", xt.JSON(ir), str)
}

func TestIRGeneratorHelloWorld(t *testing.T) {
/*
  extern int printf(char* format, ...);
  int main(int argc, char*[] argv) {
    printf("hello, world\n");
    return 0;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                  xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0),
                },
                false,
              ),
            ),
          ),
          "main",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "argc"),
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0)), "argv"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewExprStmtNode(loc,
                xtc_ast.NewFuncallNode(loc,
                  xtc_ast.NewVariableNode(loc, "printf"),
                  []xtc_core.IExprNode {
                    xtc_ast.NewStringLiteralNode(loc, "hello, world\n"),
                  },
                ),
              ),
              xtc_ast.NewReturnNode(loc, xtc_ast.NewIntegerLiteralNode(loc, "0")),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(
        xtc_entity.NewUndefinedFunction(
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)),
                },
                true,
              ),
            ),
          ),
          "printf",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc))), "format"),
            ),
            true,
          ),
        ),
      ),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  str := `{
  "ClassName": "ir.IR",
  "Location": "[:0,0]",
  "Defvars": [],
  "Defuns": [
    {
      "ClassName": "entity.DefinedFunction",
      "Private": false,
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int(int,char*[])",
        "Type": "int(int,char**)"
      },
      "Name": "main",
      "Params": {
        "ClassName": "entity.Params",
        "Location": "[:0,0]",
        "ParamDescs": [
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "argc",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "int",
              "Type": "int"
            },
            "Initializer": null,
            "IR": null
          },
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "argv",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*[]",
              "Type": "char**"
            },
            "Initializer": null,
            "IR": null
          }
        ],
        "Vararg": false
      },
      "Body": {
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
                "Name": "printf",
                "Entity": {
                  "ClassName": "entity.UndefinedFunction",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int(char*)",
                    "Type": "int(char*)"
                  },
                  "Name": "printf",
                  "Params": {
                    "ClassName": "entity.Params",
                    "Location": "[:0,0]",
                    "ParamDescs": [
                      {
                        "ClassName": "entity.Parameter",
                        "Private": true,
                        "Name": "format",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "char*",
                          "Type": "char*"
                        },
                        "Initializer": null,
                        "IR": null
                      }
                    ],
                    "Vararg": true
                  }
                }
              },
              "Args": [
                {
                  "ClassName": "ast.StringLiteralNode",
                  "Location": "[:0,0]",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "char*",
                    "Type": "char*"
                  },
                  "Value": "hello, world\n"
                }
              ]
            }
          },
          {
            "ClassName": "ast.ReturnNode",
            "Location": "[:0,0]",
            "Expr": {
              "ClassName": "ast.IntegerLiteralNode",
              "Location": "[:0,0]",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Value": 0
            }
          }
        ]
      },
      "IR": [
        {
          "ClassName": "ir.ExprStmt",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ir.Call",
            "TypeId": 2,
            "Expr": {
              "ClassName": "ir.Addr",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.UndefinedFunction",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int(char*)",
                  "Type": "int(char*)"
                },
                "Name": "printf",
                "Params": {
                  "ClassName": "entity.Params",
                  "Location": "[:0,0]",
                  "ParamDescs": [
                    {
                      "ClassName": "entity.Parameter",
                      "Private": true,
                      "Name": "format",
                      "TypeNode": {
                        "ClassName": "ast.TypeNode",
                        "Location": "[:0,0]",
                        "TypeRef": "char*",
                        "Type": "char*"
                      },
                      "Initializer": null,
                      "IR": null
                    }
                  ],
                  "Vararg": true
                }
              }
            },
            "Args": [
              {
                "ClassName": "ir.Str",
                "TypeId": 2,
                "Entry": {}
              }
            ]
          }
        },
        {
          "ClassName": "ir.Return",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ir.Int",
            "TypeId": 2,
            "Value": 0
          }
        }
      ]
    }
  ],
  "Funcdecls": [
    {
      "ClassName": "entity.UndefinedFunction",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int(char*)",
        "Type": "int(char*)"
      },
      "Name": "printf",
      "Params": {
        "ClassName": "entity.Params",
        "Location": "[:0,0]",
        "ParamDescs": [
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "format",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*",
              "Type": "char*"
            },
            "Initializer": null,
            "IR": null
          }
        ],
        "Vararg": true
      }
    }
  ]
}`
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  ir, err := setupIRGenerator(ast, table)
  xt.AssertNil(t, "should not be failed", err)
  xt.AssertStringEqualsDiff(t, "should return IR", xt.JSON(ir), str)
}

func TestIRGeneratorIfStatement(t *testing.T) {
/*
  extern int printf(char* format, ...);
  int main(int argc, char*[] argv) {
    if (argc % 2 == 0) {
      printf("even\n");
    } else {
      printf("odd\n");
    }
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                  xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0),
                },
                false,
              ),
            ),
          ),
          "main",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "argc"),
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewArrayTypeRef(xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)), 0)), "argv"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewIfNode(loc,
                xtc_ast.NewBinaryOpNode(loc, "==",
                  xtc_ast.NewBinaryOpNode(loc, "%",
                    xtc_ast.NewVariableNode(loc, "argc"),
                    xtc_ast.NewIntegerLiteralNode(loc, "2"),
                  ),
                  xtc_ast.NewIntegerLiteralNode(loc, "0"),
                ),
                xtc_ast.NewBlockNode(loc,
                  xtc_entity.NewDefinedVariables(),
                  []xtc_core.IStmtNode {
                    xtc_ast.NewExprStmtNode(loc,
                      xtc_ast.NewFuncallNode(loc,
                        xtc_ast.NewVariableNode(loc, "printf"),
                        []xtc_core.IExprNode {
                          xtc_ast.NewStringLiteralNode(loc, "even\n"),
                        },
                      ),
                    ),
                  },
                ),
                xtc_ast.NewBlockNode(loc,
                  xtc_entity.NewDefinedVariables(),
                  []xtc_core.IStmtNode {
                    xtc_ast.NewExprStmtNode(loc,
                      xtc_ast.NewFuncallNode(loc,
                        xtc_ast.NewVariableNode(loc, "printf"),
                        []xtc_core.IExprNode {
                          xtc_ast.NewStringLiteralNode(loc, "odd\n"),
                        },
                      ),
                    ),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(
        xtc_entity.NewUndefinedFunction(
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc)),
                },
                true,
              ),
            ),
          ),
          "printf",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc))), "format"),
            ),
            true,
          ),
        ),
      ),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  str := `{
  "ClassName": "ir.IR",
  "Location": "[:0,0]",
  "Defvars": [],
  "Defuns": [
    {
      "ClassName": "entity.DefinedFunction",
      "Private": false,
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int(int,char*[])",
        "Type": "int(int,char**)"
      },
      "Name": "main",
      "Params": {
        "ClassName": "entity.Params",
        "Location": "[:0,0]",
        "ParamDescs": [
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "argc",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "int",
              "Type": "int"
            },
            "Initializer": null,
            "IR": null
          },
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "argv",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*[]",
              "Type": "char**"
            },
            "Initializer": null,
            "IR": null
          }
        ],
        "Vararg": false
      },
      "Body": {
        "ClassName": "ast.BlockNode",
        "Location": "[:0,0]",
        "Variables": [],
        "Stmts": [
          {
            "ClassName": "ast.IfNode",
            "Location": "[:0,0]",
            "Cond": {
              "ClassName": "ast.BinaryOpNode",
              "Location": "[:0,0]",
              "Operator": "==",
              "Left": {
                "ClassName": "ast.BinaryOpNode",
                "Location": "[:0,0]",
                "Operator": "%",
                "Left": {
                  "ClassName": "ast.VariableNode",
                  "Location": "[:0,0]",
                  "Name": "argc",
                  "Entity": {
                    "ClassName": "entity.Parameter",
                    "Private": true,
                    "Name": "argc",
                    "TypeNode": {
                      "ClassName": "ast.TypeNode",
                      "Location": "[:0,0]",
                      "TypeRef": "int",
                      "Type": "int"
                    },
                    "Initializer": null,
                    "IR": null
                  }
                },
                "Right": {
                  "ClassName": "ast.IntegerLiteralNode",
                  "Location": "[:0,0]",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Value": 2
                },
                "Type": "int"
              },
              "Right": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 0
              },
              "Type": "int"
            },
            "ThenBody": {
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
                      "Name": "printf",
                      "Entity": {
                        "ClassName": "entity.UndefinedFunction",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int(char*)",
                          "Type": "int(char*)"
                        },
                        "Name": "printf",
                        "Params": {
                          "ClassName": "entity.Params",
                          "Location": "[:0,0]",
                          "ParamDescs": [
                            {
                              "ClassName": "entity.Parameter",
                              "Private": true,
                              "Name": "format",
                              "TypeNode": {
                                "ClassName": "ast.TypeNode",
                                "Location": "[:0,0]",
                                "TypeRef": "char*",
                                "Type": "char*"
                              },
                              "Initializer": null,
                              "IR": null
                            }
                          ],
                          "Vararg": true
                        }
                      }
                    },
                    "Args": [
                      {
                        "ClassName": "ast.StringLiteralNode",
                        "Location": "[:0,0]",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "char*",
                          "Type": "char*"
                        },
                        "Value": "even\n"
                      }
                    ]
                  }
                }
              ]
            },
            "ElseBody": {
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
                      "Name": "printf",
                      "Entity": {
                        "ClassName": "entity.UndefinedFunction",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int(char*)",
                          "Type": "int(char*)"
                        },
                        "Name": "printf",
                        "Params": {
                          "ClassName": "entity.Params",
                          "Location": "[:0,0]",
                          "ParamDescs": [
                            {
                              "ClassName": "entity.Parameter",
                              "Private": true,
                              "Name": "format",
                              "TypeNode": {
                                "ClassName": "ast.TypeNode",
                                "Location": "[:0,0]",
                                "TypeRef": "char*",
                                "Type": "char*"
                              },
                              "Initializer": null,
                              "IR": null
                            }
                          ],
                          "Vararg": true
                        }
                      }
                    },
                    "Args": [
                      {
                        "ClassName": "ast.StringLiteralNode",
                        "Location": "[:0,0]",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "char*",
                          "Type": "char*"
                        },
                        "Value": "odd\n"
                      }
                    ]
                  }
                }
              ]
            }
          }
        ]
      },
      "IR": [
        {
          "ClassName": "ir.CJump",
          "Location": "[:0,0]",
          "Cond": {
            "ClassName": "ir.Bin",
            "TypeId": 2,
            "Op": 13,
            "Left": {
              "ClassName": "ir.Bin",
              "TypeId": 2,
              "Op": 5,
              "Left": {
                "ClassName": "ir.Var",
                "TypeId": 2,
                "Entity": {
                  "ClassName": "entity.Parameter",
                  "Private": true,
                  "Name": "argc",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Initializer": null,
                  "IR": null
                }
              },
              "Right": {
                "ClassName": "ir.Int",
                "TypeId": 2,
                "Value": 2
              }
            },
            "Right": {
              "ClassName": "ir.Int",
              "TypeId": 2,
              "Value": 0
            }
          },
          "ThenLabel": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          },
          "ElseLabel": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.ExprStmt",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ir.Call",
            "TypeId": 2,
            "Expr": {
              "ClassName": "ir.Addr",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.UndefinedFunction",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int(char*)",
                  "Type": "int(char*)"
                },
                "Name": "printf",
                "Params": {
                  "ClassName": "entity.Params",
                  "Location": "[:0,0]",
                  "ParamDescs": [
                    {
                      "ClassName": "entity.Parameter",
                      "Private": true,
                      "Name": "format",
                      "TypeNode": {
                        "ClassName": "ast.TypeNode",
                        "Location": "[:0,0]",
                        "TypeRef": "char*",
                        "Type": "char*"
                      },
                      "Initializer": null,
                      "IR": null
                    }
                  ],
                  "Vararg": true
                }
              }
            },
            "Args": [
              {
                "ClassName": "ir.Str",
                "TypeId": 2,
                "Entry": {}
              }
            ]
          }
        },
        {
          "ClassName": "ir.Jump",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.ExprStmt",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ir.Call",
            "TypeId": 2,
            "Expr": {
              "ClassName": "ir.Addr",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.UndefinedFunction",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int(char*)",
                  "Type": "int(char*)"
                },
                "Name": "printf",
                "Params": {
                  "ClassName": "entity.Params",
                  "Location": "[:0,0]",
                  "ParamDescs": [
                    {
                      "ClassName": "entity.Parameter",
                      "Private": true,
                      "Name": "format",
                      "TypeNode": {
                        "ClassName": "ast.TypeNode",
                        "Location": "[:0,0]",
                        "TypeRef": "char*",
                        "Type": "char*"
                      },
                      "Initializer": null,
                      "IR": null
                    }
                  ],
                  "Vararg": true
                }
              }
            },
            "Args": [
              {
                "ClassName": "ir.Str",
                "TypeId": 2,
                "Entry": {}
              }
            ]
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        }
      ]
    }
  ],
  "Funcdecls": [
    {
      "ClassName": "entity.UndefinedFunction",
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int(char*)",
        "Type": "int(char*)"
      },
      "Name": "printf",
      "Params": {
        "ClassName": "entity.Params",
        "Location": "[:0,0]",
        "ParamDescs": [
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "format",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "char*",
              "Type": "char*"
            },
            "Initializer": null,
            "IR": null
          }
        ],
        "Vararg": true
      }
    }
  ]
}`
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  ir, err := setupIRGenerator(ast, table)
  xt.AssertNil(t, "should not be failed", err)
  xt.AssertStringEqualsDiff(t, "should return IR", xt.JSON(ir), str)
}

func TestIRGeneratorWhileStatement(t *testing.T) {
/*
  int fact(int n) {
    int i = 1;
    int r = 1;
    while (i <= n) {
      r *= i;
      i++;
    }
    return r;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "fact",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "n"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(
              xtc_entity.NewDefinedVariable(
                false,
                xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
                "i",
                xtc_ast.NewIntegerLiteralNode(loc, "1"),
              ),
              xtc_entity.NewDefinedVariable(
                false,
                xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
                "r",
                xtc_ast.NewIntegerLiteralNode(loc, "1"),
              ),
            ),
            []xtc_core.IStmtNode {
              xtc_ast.NewWhileNode(loc,
                xtc_ast.NewBinaryOpNode(loc, "<=",
                  xtc_ast.NewVariableNode(loc, "i"),
                  xtc_ast.NewVariableNode(loc, "n"),
                ),
                xtc_ast.NewBlockNode(loc,
                  xtc_entity.NewDefinedVariables(),
                  []xtc_core.IStmtNode {
                    xtc_ast.NewExprStmtNode(loc, xtc_ast.NewOpAssignNode(loc, "*", xtc_ast.NewVariableNode(loc, "r"), xtc_ast.NewVariableNode(loc, "i"))),
                    xtc_ast.NewExprStmtNode(loc, xtc_ast.NewSuffixOpNode(loc, "++", xtc_ast.NewVariableNode(loc, "i"))),
                    xtc_ast.NewReturnNode(loc, xtc_ast.NewVariableNode(loc, "r")),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  str := `{
  "ClassName": "ir.IR",
  "Location": "[:0,0]",
  "Defvars": [],
  "Defuns": [
    {
      "ClassName": "entity.DefinedFunction",
      "Private": false,
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int(int)",
        "Type": "int(int)"
      },
      "Name": "fact",
      "Params": {
        "ClassName": "entity.Params",
        "Location": "[:0,0]",
        "ParamDescs": [
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "n",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "int",
              "Type": "int"
            },
            "Initializer": null,
            "IR": null
          }
        ],
        "Vararg": false
      },
      "Body": {
        "ClassName": "ast.BlockNode",
        "Location": "[:0,0]",
        "Variables": [
          {
            "ClassName": "entity.DefinedVariable",
            "Private": false,
            "Name": "i",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "int",
              "Type": "int"
            },
            "Initializer": {
              "ClassName": "ast.IntegerLiteralNode",
              "Location": "[:0,0]",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Value": 1
            },
            "IR": null
          },
          {
            "ClassName": "entity.DefinedVariable",
            "Private": false,
            "Name": "r",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "int",
              "Type": "int"
            },
            "Initializer": {
              "ClassName": "ast.IntegerLiteralNode",
              "Location": "[:0,0]",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Value": 1
            },
            "IR": null
          }
        ],
        "Stmts": [
          {
            "ClassName": "ast.WhileNode",
            "Location": "[:0,0]",
            "Cond": {
              "ClassName": "ast.BinaryOpNode",
              "Location": "[:0,0]",
              "Operator": "\u003c=",
              "Left": {
                "ClassName": "ast.VariableNode",
                "Location": "[:0,0]",
                "Name": "i",
                "Entity": {
                  "ClassName": "entity.DefinedVariable",
                  "Private": false,
                  "Name": "i",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Initializer": {
                    "ClassName": "ast.IntegerLiteralNode",
                    "Location": "[:0,0]",
                    "TypeNode": {
                      "ClassName": "ast.TypeNode",
                      "Location": "[:0,0]",
                      "TypeRef": "int",
                      "Type": "int"
                    },
                    "Value": 1
                  },
                  "IR": null
                }
              },
              "Right": {
                "ClassName": "ast.VariableNode",
                "Location": "[:0,0]",
                "Name": "n",
                "Entity": {
                  "ClassName": "entity.Parameter",
                  "Private": true,
                  "Name": "n",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Initializer": null,
                  "IR": null
                }
              },
              "Type": "int"
            },
            "Body": {
              "ClassName": "ast.BlockNode",
              "Location": "[:0,0]",
              "Variables": [],
              "Stmts": [
                {
                  "ClassName": "ast.ExprStmtNode",
                  "Location": "[:0,0]",
                  "Expr": {
                    "ClassName": "ast.OpAssignNode",
                    "Location": "[:0,0]",
                    "Operator": "*",
                    "LHS": {
                      "ClassName": "ast.VariableNode",
                      "Location": "[:0,0]",
                      "Name": "r",
                      "Entity": {
                        "ClassName": "entity.DefinedVariable",
                        "Private": false,
                        "Name": "r",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int",
                          "Type": "int"
                        },
                        "Initializer": {
                          "ClassName": "ast.IntegerLiteralNode",
                          "Location": "[:0,0]",
                          "TypeNode": {
                            "ClassName": "ast.TypeNode",
                            "Location": "[:0,0]",
                            "TypeRef": "int",
                            "Type": "int"
                          },
                          "Value": 1
                        },
                        "IR": null
                      }
                    },
                    "RHS": {
                      "ClassName": "ast.VariableNode",
                      "Location": "[:0,0]",
                      "Name": "i",
                      "Entity": {
                        "ClassName": "entity.DefinedVariable",
                        "Private": false,
                        "Name": "i",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int",
                          "Type": "int"
                        },
                        "Initializer": {
                          "ClassName": "ast.IntegerLiteralNode",
                          "Location": "[:0,0]",
                          "TypeNode": {
                            "ClassName": "ast.TypeNode",
                            "Location": "[:0,0]",
                            "TypeRef": "int",
                            "Type": "int"
                          },
                          "Value": 1
                        },
                        "IR": null
                      }
                    }
                  }
                },
                {
                  "ClassName": "ast.ExprStmtNode",
                  "Location": "[:0,0]",
                  "Expr": {
                    "ClassName": "ast.SuffixOpNode",
                    "Location": "[:0,0]",
                    "Operator": "++",
                    "Expr": {
                      "ClassName": "ast.VariableNode",
                      "Location": "[:0,0]",
                      "Name": "i",
                      "Entity": {
                        "ClassName": "entity.DefinedVariable",
                        "Private": false,
                        "Name": "i",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int",
                          "Type": "int"
                        },
                        "Initializer": {
                          "ClassName": "ast.IntegerLiteralNode",
                          "Location": "[:0,0]",
                          "TypeNode": {
                            "ClassName": "ast.TypeNode",
                            "Location": "[:0,0]",
                            "TypeRef": "int",
                            "Type": "int"
                          },
                          "Value": 1
                        },
                        "IR": null
                      }
                    },
                    "Amount": 1,
                    "Type": null
                  }
                },
                {
                  "ClassName": "ast.ReturnNode",
                  "Location": "[:0,0]",
                  "Expr": {
                    "ClassName": "ast.VariableNode",
                    "Location": "[:0,0]",
                    "Name": "r",
                    "Entity": {
                      "ClassName": "entity.DefinedVariable",
                      "Private": false,
                      "Name": "r",
                      "TypeNode": {
                        "ClassName": "ast.TypeNode",
                        "Location": "[:0,0]",
                        "TypeRef": "int",
                        "Type": "int"
                      },
                      "Initializer": {
                        "ClassName": "ast.IntegerLiteralNode",
                        "Location": "[:0,0]",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int",
                          "Type": "int"
                        },
                        "Value": 1
                      },
                      "IR": null
                    }
                  }
                }
              ]
            }
          }
        ]
      },
      "IR": [
        {
          "ClassName": "ir.Assign",
          "Location": "[:0,0]",
          "LHS": {
            "ClassName": "ir.Addr",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.DefinedVariable",
              "Private": false,
              "Name": "i",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "IR": null
            }
          },
          "RHS": {
            "ClassName": "ir.Int",
            "TypeId": 2,
            "Value": 1
          }
        },
        {
          "ClassName": "ir.Assign",
          "Location": "[:0,0]",
          "LHS": {
            "ClassName": "ir.Addr",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.DefinedVariable",
              "Private": false,
              "Name": "r",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "IR": null
            }
          },
          "RHS": {
            "ClassName": "ir.Int",
            "TypeId": 2,
            "Value": 1
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.CJump",
          "Location": "[:0,0]",
          "Cond": {
            "ClassName": "ir.Bin",
            "TypeId": 2,
            "Op": 18,
            "Left": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.DefinedVariable",
                "Private": false,
                "Name": "i",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": {
                  "ClassName": "ast.IntegerLiteralNode",
                  "Location": "[:0,0]",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Value": 1
                },
                "IR": null
              }
            },
            "Right": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.Parameter",
                "Private": true,
                "Name": "n",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": null,
                "IR": null
              }
            }
          },
          "ThenLabel": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          },
          "ElseLabel": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.Assign",
          "Location": "[:0,0]",
          "LHS": {
            "ClassName": "ir.Addr",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.DefinedVariable",
              "Private": false,
              "Name": "r",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "IR": null
            }
          },
          "RHS": {
            "ClassName": "ir.Bin",
            "TypeId": 2,
            "Op": 2,
            "Left": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.DefinedVariable",
                "Private": false,
                "Name": "r",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": {
                  "ClassName": "ast.IntegerLiteralNode",
                  "Location": "[:0,0]",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Value": 1
                },
                "IR": null
              }
            },
            "Right": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.DefinedVariable",
                "Private": false,
                "Name": "i",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": {
                  "ClassName": "ast.IntegerLiteralNode",
                  "Location": "[:0,0]",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Value": 1
                },
                "IR": null
              }
            }
          }
        },
        {
          "ClassName": "ir.Assign",
          "Location": "[:0,0]",
          "LHS": {
            "ClassName": "ir.Addr",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.DefinedVariable",
              "Private": false,
              "Name": "i",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "IR": null
            }
          },
          "RHS": {
            "ClassName": "ir.Bin",
            "TypeId": 2,
            "Op": 0,
            "Left": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.DefinedVariable",
                "Private": false,
                "Name": "i",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": {
                  "ClassName": "ast.IntegerLiteralNode",
                  "Location": "[:0,0]",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Value": 1
                },
                "IR": null
              }
            },
            "Right": {
              "ClassName": "ir.Int",
              "TypeId": 2,
              "Value": 1
            }
          }
        },
        {
          "ClassName": "ir.Return",
          "Location": "[:0,0]",
          "Expr": {
            "ClassName": "ir.Var",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.DefinedVariable",
              "Private": false,
              "Name": "r",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "IR": null
            }
          }
        },
        {
          "ClassName": "ir.Jump",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        }
      ]
    }
  ],
  "Funcdecls": []
}`
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  ir, err := setupIRGenerator(ast, table)
  xt.AssertNil(t, "should not be failed", err)
  xt.AssertStringEqualsDiff(t, "should return IR", xt.JSON(ir), str)
}

func TestIRGeneratorForStatement(t *testing.T) {
/*
  int fact(int n) {
    int r = 1;
    for (n ; 1 < n; n--) {
      r *= n;
    }
    return r;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "fact",
          xtc_entity.NewParams(loc,
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)), "n"),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(
              xtc_entity.NewDefinedVariable(
                false,
                xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
                "r",
                xtc_ast.NewIntegerLiteralNode(loc, "1"),
              ),
            ),
            []xtc_core.IStmtNode {
              xtc_ast.NewForNode(loc,
                xtc_ast.NewVariableNode(loc, "n"),
                xtc_ast.NewBinaryOpNode(loc, "<",
                  xtc_ast.NewIntegerLiteralNode(loc, "1"),
                  xtc_ast.NewVariableNode(loc, "n"),
                ),
                xtc_ast.NewSuffixOpNode(loc, "--", xtc_ast.NewVariableNode(loc, "n")),
                xtc_ast.NewBlockNode(loc,
                  xtc_entity.NewDefinedVariables(),
                  []xtc_core.IStmtNode {
                    xtc_ast.NewExprStmtNode(loc, xtc_ast.NewOpAssignNode(loc, "*", xtc_ast.NewVariableNode(loc, "r"), xtc_ast.NewVariableNode(loc, "n"))),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  str := `{
  "ClassName": "ir.IR",
  "Location": "[:0,0]",
  "Defvars": [],
  "Defuns": [
    {
      "ClassName": "entity.DefinedFunction",
      "Private": false,
      "TypeNode": {
        "ClassName": "ast.TypeNode",
        "Location": "[:0,0]",
        "TypeRef": "int(int)",
        "Type": "int(int)"
      },
      "Name": "fact",
      "Params": {
        "ClassName": "entity.Params",
        "Location": "[:0,0]",
        "ParamDescs": [
          {
            "ClassName": "entity.Parameter",
            "Private": true,
            "Name": "n",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "int",
              "Type": "int"
            },
            "Initializer": null,
            "IR": null
          }
        ],
        "Vararg": false
      },
      "Body": {
        "ClassName": "ast.BlockNode",
        "Location": "[:0,0]",
        "Variables": [
          {
            "ClassName": "entity.DefinedVariable",
            "Private": false,
            "Name": "r",
            "TypeNode": {
              "ClassName": "ast.TypeNode",
              "Location": "[:0,0]",
              "TypeRef": "int",
              "Type": "int"
            },
            "Initializer": {
              "ClassName": "ast.IntegerLiteralNode",
              "Location": "[:0,0]",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Value": 1
            },
            "IR": null
          }
        ],
        "Stmts": [
          {
            "ClassName": "ast.ForNode",
            "Location": "[:0,0]",
            "Init": {
              "ClassName": "ast.VariableNode",
              "Location": "[:0,0]",
              "Name": "n",
              "Entity": {
                "ClassName": "entity.Parameter",
                "Private": true,
                "Name": "n",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": null,
                "IR": null
              }
            },
            "Cond": {
              "ClassName": "ast.BinaryOpNode",
              "Location": "[:0,0]",
              "Operator": "\u003c",
              "Left": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "Right": {
                "ClassName": "ast.VariableNode",
                "Location": "[:0,0]",
                "Name": "n",
                "Entity": {
                  "ClassName": "entity.Parameter",
                  "Private": true,
                  "Name": "n",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Initializer": null,
                  "IR": null
                }
              },
              "Type": "int"
            },
            "Incr": {
              "ClassName": "ast.SuffixOpNode",
              "Location": "[:0,0]",
              "Operator": "--",
              "Expr": {
                "ClassName": "ast.VariableNode",
                "Location": "[:0,0]",
                "Name": "n",
                "Entity": {
                  "ClassName": "entity.Parameter",
                  "Private": true,
                  "Name": "n",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Initializer": null,
                  "IR": null
                }
              },
              "Amount": 1,
              "Type": null
            },
            "Body": {
              "ClassName": "ast.BlockNode",
              "Location": "[:0,0]",
              "Variables": [],
              "Stmts": [
                {
                  "ClassName": "ast.ExprStmtNode",
                  "Location": "[:0,0]",
                  "Expr": {
                    "ClassName": "ast.OpAssignNode",
                    "Location": "[:0,0]",
                    "Operator": "*",
                    "LHS": {
                      "ClassName": "ast.VariableNode",
                      "Location": "[:0,0]",
                      "Name": "r",
                      "Entity": {
                        "ClassName": "entity.DefinedVariable",
                        "Private": false,
                        "Name": "r",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int",
                          "Type": "int"
                        },
                        "Initializer": {
                          "ClassName": "ast.IntegerLiteralNode",
                          "Location": "[:0,0]",
                          "TypeNode": {
                            "ClassName": "ast.TypeNode",
                            "Location": "[:0,0]",
                            "TypeRef": "int",
                            "Type": "int"
                          },
                          "Value": 1
                        },
                        "IR": null
                      }
                    },
                    "RHS": {
                      "ClassName": "ast.VariableNode",
                      "Location": "[:0,0]",
                      "Name": "n",
                      "Entity": {
                        "ClassName": "entity.Parameter",
                        "Private": true,
                        "Name": "n",
                        "TypeNode": {
                          "ClassName": "ast.TypeNode",
                          "Location": "[:0,0]",
                          "TypeRef": "int",
                          "Type": "int"
                        },
                        "Initializer": null,
                        "IR": null
                      }
                    }
                  }
                }
              ]
            }
          }
        ]
      },
      "IR": [
        {
          "ClassName": "ir.Assign",
          "Location": "[:0,0]",
          "LHS": {
            "ClassName": "ir.Addr",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.DefinedVariable",
              "Private": false,
              "Name": "r",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "IR": null
            }
          },
          "RHS": {
            "ClassName": "ir.Int",
            "TypeId": 2,
            "Value": 1
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.CJump",
          "Location": "[:0,0]",
          "Cond": {
            "ClassName": "ir.Bin",
            "TypeId": 2,
            "Op": 17,
            "Left": {
              "ClassName": "ir.Int",
              "TypeId": 2,
              "Value": 1
            },
            "Right": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.Parameter",
                "Private": true,
                "Name": "n",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": null,
                "IR": null
              }
            }
          },
          "ThenLabel": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          },
          "ElseLabel": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.Assign",
          "Location": "[:0,0]",
          "LHS": {
            "ClassName": "ir.Addr",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.DefinedVariable",
              "Private": false,
              "Name": "r",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": {
                "ClassName": "ast.IntegerLiteralNode",
                "Location": "[:0,0]",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Value": 1
              },
              "IR": null
            }
          },
          "RHS": {
            "ClassName": "ir.Bin",
            "TypeId": 2,
            "Op": 2,
            "Left": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.DefinedVariable",
                "Private": false,
                "Name": "r",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": {
                  "ClassName": "ast.IntegerLiteralNode",
                  "Location": "[:0,0]",
                  "TypeNode": {
                    "ClassName": "ast.TypeNode",
                    "Location": "[:0,0]",
                    "TypeRef": "int",
                    "Type": "int"
                  },
                  "Value": 1
                },
                "IR": null
              }
            },
            "Right": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.Parameter",
                "Private": true,
                "Name": "n",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": null,
                "IR": null
              }
            }
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.Assign",
          "Location": "[:0,0]",
          "LHS": {
            "ClassName": "ir.Addr",
            "TypeId": 2,
            "Entity": {
              "ClassName": "entity.Parameter",
              "Private": true,
              "Name": "n",
              "TypeNode": {
                "ClassName": "ast.TypeNode",
                "Location": "[:0,0]",
                "TypeRef": "int",
                "Type": "int"
              },
              "Initializer": null,
              "IR": null
            }
          },
          "RHS": {
            "ClassName": "ir.Bin",
            "TypeId": 2,
            "Op": 1,
            "Left": {
              "ClassName": "ir.Var",
              "TypeId": 2,
              "Entity": {
                "ClassName": "entity.Parameter",
                "Private": true,
                "Name": "n",
                "TypeNode": {
                  "ClassName": "ast.TypeNode",
                  "Location": "[:0,0]",
                  "TypeRef": "int",
                  "Type": "int"
                },
                "Initializer": null,
                "IR": null
              }
            },
            "Right": {
              "ClassName": "ir.Int",
              "TypeId": 2,
              "Value": 1
            }
          }
        },
        {
          "ClassName": "ir.Jump",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        },
        {
          "ClassName": "ir.LabelStmt",
          "Location": "[builtin:compiler/ir_generator.go:0,0]",
          "Label": {
            "ClassName": "asm.Label",
            "Symbol": {
              "ClassName": "asm.UnnamedSymbol"
            }
          }
        }
      ]
    }
  ],
  "Funcdecls": []
}`
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  ir, err := setupIRGenerator(ast, table)
  xt.AssertNil(t, "should not be failed", err)
  xt.AssertStringEqualsDiff(t, "should return IR", xt.JSON(ir), str)
}
