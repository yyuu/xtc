package asm

import (
  "fmt"
)

const (
  TYPE_INT8 = iota
  TYPE_INT16
  TYPE_INT32
  TYPE_INT64
)

func TypeGet(size int) int {
  switch size {
    case 1: return TYPE_INT8
    case 2: return TYPE_INT16
    case 4: return TYPE_INT32
    case 8: return TYPE_INT64
    default: {
      panic(fmt.Errorf("unsupported asm type size: %d", size))
    }
  }
}

func TypeSize(t int) int {
  switch t {
    case TYPE_INT8:  return 1
    case TYPE_INT16: return 2
    case TYPE_INT32: return 4
    case TYPE_INT64: return 8
    default: {
      panic("must not happen")
    }
  }
}
