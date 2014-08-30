package parser

import (
  "reflect"
  "testing"
  "bitbucket.org/yyuu/bs/ast"
)

func assertEqualsAST(t *testing.T, got ast.AST, expected ast.AST) {
  if ! reflect.DeepEqual(got, expected) {
    t.Errorf("\n;;;; expected ;;;;\n%s\n;;;; got ;;;;\n%s\n", expected, got)
    t.Fail()
  }
}

func TestParseEmpty(t *testing.T) {
  _, err := ParseExpr("")
  if err != nil {
    t.Fail()
  }
}

func loc(lineNumber int, lineOffset int) ast.Location {
  return ast.Location { SourceName: "-", LineNumber: lineNumber, LineOffset: lineOffset }
}

func TestParseFuncallWithoutArguments(t *testing.T) {
  s := `
gets( );
  `
  a, err := ParseExpr(s)
  if err != nil {
    t.Fail()
  }
  assertEqualsAST(t, *a,
    ast.AST {
      []ast.IStmtNode {
        ast.NewExprStmtNode(loc(1,1),
          ast.NewFuncallNode(loc(1,1),
                          ast.NewVariableNode(loc(1,1), "gets"),
                          []ast.IExprNode {
                          })),
      },
    },
  )
}

func TestParseFuncallWithSingleArgument(t *testing.T) {
  s := `
    println("hello, world");
  `
  a, err := ParseExpr(s)
  if err != nil {
    t.Fail()
  }
  assertEqualsAST(t, *a,
    ast.AST {
      []ast.IStmtNode {
        ast.NewExprStmtNode(loc(1,5),
          ast.NewFuncallNode(loc(1,5),
                          ast.NewVariableNode(loc(1,5), "println"),
                          []ast.IExprNode {
                            ast.NewStringLiteralNode(loc(1,13), "\"hello, world\""),
                          })),
      },
    },
  )
}

func TestParseFuncallWithMultipleArguments(t *testing.T) {
  s := `

    println(
      "hello, %s",
      "world"
    );

  `
  a, err := ParseExpr(s)
  if err != nil {
    t.Fail()
  }
  assertEqualsAST(t, *a,
    ast.AST {
      []ast.IStmtNode {
        ast.NewExprStmtNode(loc(2,5),
          ast.NewFuncallNode(loc(2,5),
                          ast.NewVariableNode(loc(2,5), "println"),
                          []ast.IExprNode {
                            ast.NewStringLiteralNode(loc(3,7), "\"hello, %s\""),
                            ast.NewStringLiteralNode(loc(4,7), "\"world\""),
                          })),
      },
    },
  )
}
