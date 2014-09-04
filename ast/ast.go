package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

type Location struct {
  SourceName string
  LineNumber int
  LineOffset int
}

func NewLocation(sourceName string, lineNumber int, lineOffset int) duck.ILocation {
  return Location { sourceName, lineNumber, lineOffset }
}

func (self Location) GetSourceName() string {
  return self.SourceName
}

func (self Location) GetLineNumber() int {
  return self.LineNumber
}

func (self Location) GetLineOffset() int {
  return self.LineOffset
}

func (self Location) String() string {
  return fmt.Sprintf("[%s:%d,%d]", self.SourceName, self.LineNumber, self.LineOffset)
}

func (self Location) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.String())
  return []byte(s), nil
}

type AST struct {
  Stmts []duck.IStmtNode
}

func (self AST) String() string {
  xs := make([]string, len(self.Stmts))
  for i := range self.Stmts {
    stmt := self.Stmts[i]
    xs[i] = fmt.Sprintf(";; %s\n%s", stmt.GetLocation(), stmt)
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
