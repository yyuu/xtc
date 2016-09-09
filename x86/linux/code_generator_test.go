package x86_linux

import (
  "testing"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_ir "bitbucket.org/yyuu/xtc/ir"
  "bitbucket.org/yyuu/xtc/xt"
)

func setupCodeGenerator(ir *xtc_ir.IR) *CodeGenerator {
  errorHandler := xtc_core.NewErrorHandler(xtc_core.LOG_WARN)
  options := xtc_core.NewOptions("code_generator_test.go")
  return NewCodeGenerator(errorHandler, options)
}

func TestCodeGeneratorEmpty(t *testing.T) {
  loc := xtc_core.NewLocation("foo", 0, 0)
  ir := xtc_ir.NewIR(loc,
    xtc_entity.NewDefinedVariables(),
    xtc_entity.NewDefinedFunctions(),
    xtc_entity.NewUndefinedFunctions(),
    xtc_entity.NewToplevelScope(),
    xtc_entity.NewConstantTable(),
  )
  str := `{
  "NaturalType": 2,
  "LabelSymbols": {
    "ClassName": "asm.SymbolTable",
    "Base": ".L",
    "Seq": 0
  },
  "Assemblies": [
    {
      "ClassName": "asm.Directive",
      "Content": "\t.file\t\"foo\""
    }
  ]
}`
  generator := setupCodeGenerator(ir)
  asm, err := generator.Generate(ir)
  xt.AssertNil(t, "should not be failed", err)
  xt.AssertStringEqualsDiff(t, "should return empty assembly", xt.JSON(asm), str)
}
