package parser

import (
  bs_core "bitbucket.org/yyuu/bs/core"
)

func loc(lineNumber int, lineOffset int) bs_core.Location {
  return bs_core.NewLocation("", lineNumber, lineOffset)
}
