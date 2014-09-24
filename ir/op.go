package ir

import (
  "fmt"
)

const (
  OP_ADD = iota
  OP_SUB
  OP_MUL
  OP_S_DIV
  OP_U_DIV
  OP_S_MOD
  OP_U_MOD
  OP_BIT_AND
  OP_BIT_OR
  OP_BIT_XOR
  OP_BIT_LSHIFT
  OP_BIT_RSHIFT
  OP_ARITH_RSHIFT

  OP_EQ
  OP_NEQ
  OP_S_GT
  OP_S_GTEQ
  OP_S_LT
  OP_S_LTEQ
  OP_U_GT
  OP_U_GTEQ
  OP_U_LT
  OP_U_LTEQ

  OP_UMINUS
  OP_BIT_NOT
  OP_NOT

  OP_S_CAST
  OP_U_CAST
)

func OpInternBinary(op string, isSigned bool) int {
  switch op {
    case "+": {
      return OP_ADD
    }
    case "-": {
      return OP_SUB
    }
    case "*": {
      return OP_MUL
    }
    case "/": {
      if isSigned {
        return OP_S_DIV
      } else {
        return OP_U_DIV
      }
    }
    case "%": {
      if isSigned {
        return OP_S_MOD
      } else {
        return OP_U_MOD
      }
    }
    case "&": {
      return OP_BIT_AND
    }
    case "|": {
      return OP_BIT_OR
    }
    case "^": {
      return OP_BIT_XOR
    }
    case "<<": {
      return OP_BIT_LSHIFT
    }
    case ">>": {
      if isSigned {
        return OP_ARITH_RSHIFT
      } else {
        return OP_BIT_RSHIFT
      }
    }
    case "==": {
      return OP_EQ
    }
    case "!=": {
      return OP_NEQ
    }
    case "<": {
      if isSigned {
        return OP_S_LT
      } else {
        return OP_U_LT
      }
    }
    case "<=": {
      if isSigned {
        return OP_S_LTEQ
      } else {
        return OP_U_LTEQ
      }
    }
    case ">": {
      if isSigned {
        return OP_S_GT
      } else {
        return OP_U_GT
      }
    }
    case ">=": {
      if isSigned {
        return OP_S_GTEQ
      } else {
        return OP_U_GTEQ
      }
    }
    default: {
      panic(fmt.Errorf("unknown binary op: %s", op))
    }
  }
}

func OpInternUnary(op string) int {
  switch op {
    case "+": {
      panic("unary+ should not be in IR")
    }
    case "-": {
      return OP_UMINUS
    }
    case "~": {
      return OP_BIT_NOT
    }
    case "!": {
      return OP_NOT
    }
    default: {
      panic(fmt.Errorf("unknown unary op: %s", op))
    }
  }
}
