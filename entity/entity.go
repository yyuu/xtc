package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

func checkAddress(ent core.IEntity, memref core.IMemoryReference, address core.IOperand) {
  if memref == nil && address == nil {
    panic(fmt.Errorf("address did not resolved: %s", ent.GetName()))
  }
}
