package ast

import (
  "bitbucket.org/yyuu/xtc/core"
)

func loc(lineNumber int, lineOffset int) core.Location {
  return core.NewLocation("", lineNumber, lineOffset)
}
