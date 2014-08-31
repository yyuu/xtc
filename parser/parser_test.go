package parser

import (
  "bytes"
  "encoding/json"
  "reflect"
  "testing"
  "bitbucket.org/yyuu/bs/ast"
)

func jsonString(x interface{}) string {
  src, err := json.Marshal(x)
  if err != nil {
    panic(err)
  }
  var dst bytes.Buffer
  err = json.Indent(&dst, src, "", "  ")
  if err != nil {
    panic(err)
  }
  return dst.String()
}

func testParse(t *testing.T, s string) ast.AST {
  a, err := ParseExpr(s)
  if err != nil {
    t.Error(err)
    t.Fail()
  }
  return *a
}

func assertEqualsAST(t *testing.T, got ast.AST, expected ast.AST) {
  if ! reflect.DeepEqual(got, expected) {
    t.Errorf("\n// expected\n%s\n// got\n%s\n", jsonString(expected), jsonString(got))
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
  return ast.Location { SourceName: "", LineNumber: lineNumber, LineOffset: lineOffset }
}

func TestParseFuncallWithoutArguments(t *testing.T) {
  s := `
gets( );
  `
  assertEqualsAST(t, testParse(t, s),
    ast.AST {
      []ast.IStmtNode {
        ast.NewExprStmtNode(loc(1,1),
          ast.NewFuncallNode(loc(1,1),
            ast.NewVariableNode(loc(1,1),
              "gets",
            ),
            []ast.IExprNode {
            },
          ),
        ),
      },
    },
  )
}

func TestParseFuncallWithSingleArgument(t *testing.T) {
  s := `
    println("hello, world");
  `
  assertEqualsAST(t, testParse(t, s),
    ast.AST {
      []ast.IStmtNode {
        ast.NewExprStmtNode(loc(1,5),
          ast.NewFuncallNode(loc(1,5),
            ast.NewVariableNode(loc(1,5),
              "println",
            ),
            []ast.IExprNode {
              ast.NewStringLiteralNode(loc(1,13),
                "\"hello, world\"",
              ),
            },
          ),
        ),
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
  assertEqualsAST(t, testParse(t, s),
    ast.AST {
      []ast.IStmtNode {
        ast.NewExprStmtNode(loc(2,5),
          ast.NewFuncallNode(loc(2,5),
            ast.NewVariableNode(loc(2,5),
              "println",
            ),
            []ast.IExprNode {
              ast.NewStringLiteralNode(loc(3,7),
                "\"hello, %s\"",
              ),
              ast.NewStringLiteralNode(loc(4,7),
                "\"world\"",
              ),
            },
          ),
        ),
      },
    },
  )
}

func TestFor1(t *testing.T) {
  s := `
    for (i=0; i<100; i++) println(i);
`
  assertEqualsAST(t, testParse(t, s),
    ast.AST {
      []ast.IStmtNode {
        ast.NewForNode(loc(1,5),
          ast.NewAssignNode(loc(1,10),
            ast.NewVariableNode(loc(1,10),
              "i",
            ),
            ast.NewIntegerLiteralNode(loc(1,12),
              "0",
            ),
          ),
          ast.NewBinaryOpNode(loc(1,15),
            "<",
            ast.NewVariableNode(loc(1,15),
              "i",
            ),
            ast.NewIntegerLiteralNode(loc(1,17),
              "100",
            ),
          ),
          ast.NewSuffixOpNode(loc(1,22),
            "++",
            ast.NewVariableNode(loc(1,22),
              "i",
            ),
          ),
          ast.NewExprStmtNode(loc(1,27),
            ast.NewFuncallNode(loc(1,27),
              ast.NewVariableNode(loc(1,27),
                "println",
              ),
              []ast.IExprNode {
                ast.NewVariableNode(loc(1,35),
                  "i",
                ),
              },
            ),
          ),
        ),
      },
    },
  )
}
