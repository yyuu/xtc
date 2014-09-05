package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Location struct {
  sourceName string
  lineNumber int
  lineOffset int
}

func NewLocation(sourceName string, lineNumber int, lineOffset int) duck.ILocation {
  return Location { sourceName, lineNumber, lineOffset }
}

func (self Location) GetSourceName() string {
  return self.sourceName
}

func (self Location) GetLineNumber() int {
  return self.lineNumber
}

func (self Location) GetLineOffset() int {
  return self.lineOffset
}

func (self Location) String() string {
  return fmt.Sprintf("[%s:%d,%d]", self.sourceName, self.lineNumber, self.lineOffset)
}

func (self Location) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.String())
  return []byte(s), nil
}

type AST struct {
  location duck.ILocation
  declarations Declarations
}

func NewAST(source duck.ILocation, declarations Declarations) AST {
  return AST { source, declarations }
}

func (self AST) String() string {
  return fmt.Sprintf(";; %s\n%s", self.location, self.declarations)
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
