package core

type IStmt interface {
  AsStmt() IStmt
  GetLocation() Location
}

type IExpr interface {
  AsExpr() IExpr
  GetType() IType
  IsVar() bool
  IsAddr() bool
  IsConstant() bool
  GetAddressNode(IType) IExpr
}
