package parser

import (
  xtc_core "bitbucket.org/yyuu/xtc/core"
)

func loc(lineNumber int, lineOffset int) xtc_core.Location {
  return xtc_core.NewLocation("", lineNumber, lineOffset)
}
