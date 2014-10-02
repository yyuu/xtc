package parser

import (
  "fmt"
  bs_core "bitbucket.org/yyuu/bs/core"
)

type token struct {
  id int
  literal string
  location bs_core.Location
}

func (self token) String() string {
  return fmt.Sprintf("%s %d %q", self.location, self.id, self.literal)
}
