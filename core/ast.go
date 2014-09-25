package core

type INode interface {
  String() string
  GetLocation() Location
}

type IExprNode interface {
  INode
  AsExprNode() IExprNode
  GetType() IType
  SetType(IType)
  GetOrigType() IType
  IsConstant() bool
  IsParameter() bool
  IsLvalue() bool
  IsAssignable() bool
  IsLoadable() bool
  IsCallable() bool
  IsPointer() bool
}

type IBinaryOpNode interface {
  IExprNode
  GetOperator() string
  GetLeft() IExprNode
  GetRight() IExprNode
  SetLeft(IExprNode)
  SetRight(IExprNode)
}

type IUnaryArithmeticOpNode interface {
  IExprNode
  GetOperator() string
  GetOpType() IType
  SetOpType(IType)
  GetExpr() IExprNode
  SetExpr(IExprNode)
  GetAmount() int
  SetAmount(int)
}

type IStmtNode interface {
  INode
  AsStmtNode() IStmtNode
}

type ITypeNode interface {
  INode
  AsTypeNode() ITypeNode
  GetTypeRef() ITypeRef
  IsResolved() bool
  GetType() IType
  SetType(IType)
}

type ISlot interface {
  INode
  GetName() string
  GetOffset() int
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
  GetType() IType
}

type ITypeDefinition interface {
  INode
  IsTypeDefinition() bool
  GetName() string
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
  DefiningType() IType
}

type ICompositeTypeDefinition interface {
  ITypeDefinition
  IsCompositeType() bool
  GetMembers() []ISlot
}
