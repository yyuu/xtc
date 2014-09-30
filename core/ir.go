package core

type IStmt interface {
  AsStmt() IStmt
  GetLocation() Location
}

type IExpr interface {
  AsExpr() IExpr
  GetTypeId() int
  IsVar() bool
  IsAddr() bool
  IsConstant() bool
  GetAsmValue() IImmediateValue
  GetAddress() IOperand
  GetMemref() IMemoryReference
  GetAddressNode(int) IExpr
  GetEntityForce() IEntity
}
