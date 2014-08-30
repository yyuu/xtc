package ast

import (
  "fmt"
  "strings"
)

type INode interface {
  String() string
  GetLocation() Location
}

type IExprNode interface {
  INode
  IsExpr() bool
}

type IStmtNode interface {
  INode
  IsStmt() bool
}

type ITypeNode interface {
  INode
  IsType() bool
}

type Location struct {
  SourceName string
  LineNumber int
  LineOffset int
}

type AST struct {
  Stmts []IStmtNode
}

func (self AST) String() string {
  xs := make([]string, len(self.Stmts))
  for i := range self.Stmts {
    stmt := self.Stmts[i]
    location := stmt.GetLocation()
    xs[i] = fmt.Sprintf(";; %s:%d,%d\n%s", location.SourceName, location.LineNumber+1, location.LineOffset+1, stmt)
  }
  return strings.Join(xs, "\n")
}

type ASTVisitor interface {
  // Statements
  VisitBlockNode(BlockNode)
  VisitExprStmtNode(BlockNode)
  VisitIfNode(IfNode)
  VisitSwitchNode(SwitchNode)
  VisitCaseNode(CaseNode)
  VisitWhileNode(WhileNode)
  VisitDoWhileNode(DoWhileNode)
  VisitForNode(ForNode)
  VisitBreakNode(BreakNode)
  VisitContinueNode(ContinueNode)
  VisitGotoNode(GotoNode)
  VisitLabelNode(LabelNode)
  VisitReturnNode(ReturnNode)

  // Expressions
  VisitAssignNode(AssignNode)
  VisitOpAssignNode(OpAssignNode)
  VisitCondExprNode(CondExprNode)
  VisitLogicalOrNode(LogicalOrNode)
  VisitLogicalAndNode(LogicalAndNode)
  VisitBinaryOpNode(BinaryOpNode)
  VisitUnaryOpNode(UnaryOpNode)
  VisitPrefixOpNode(PrefixOpNode)
  VisitSuffixOpNode(SuffixOpNode)
  VisitArefNode(ArefNode)
  VisitMemberNode(MemberNode)
  VisitPtrMemberNode(PtrMemberNode)
  VisitFuncallNode(FuncallNode)
  VisitDereferenceNode(DereferenceNode)
  VisitAddressNode(AddressNode)
  VisitCastNode(CastNode)
  VisitSizeofExprNode(SizeofExprNode)
  VisitSizeofTypeNode(SizeofTypeNode)
  VisitVariableNode(VariableNode)
  VisitIntegerLiteralNode(IntegerLiteralNode)
  VisitStringLiteralNode(StringLiteralNode)
}
