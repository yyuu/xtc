package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

func checkAddress(ent core.IEntity, memref core.IMemoryReference, address core.IOperand) {
  if memref == nil && address == nil {
    panic(fmt.Errorf("address did not resolved: %s", ent.GetName()))
  }
}
