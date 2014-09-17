package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
)

func visitAddressNode(v ast.INodeVisitor, node *ast.AddressNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitArefNode(v ast.INodeVisitor, node *ast.ArefNode) {
  ast.VisitExpr(v, node.GetExpr())
  ast.VisitExpr(v, node.GetIndex())
}

func visitAssignNode(v ast.INodeVisitor, node *ast.AssignNode) {
  ast.VisitExpr(v, node.GetLhs())
  ast.VisitExpr(v, node.GetRhs())
}

func visitBinaryOpNode(v ast.INodeVisitor, node *ast.BinaryOpNode) {
  ast.VisitExpr(v, node.GetLeft())
  ast.VisitExpr(v, node.GetRight())
}

func visitBlockNode(v ast.INodeVisitor, node *ast.BlockNode) {
  vars := node.GetVariables()
  for i := range vars {
    ast.VisitExpr(v, vars[i].GetInitializer())
  }
  ast.VisitStmts(v, node.GetStmts())
}

func visitBreakNode(v ast.INodeVisitor, node *ast.BreakNode) {
  // nop
}

func visitCaseNode(v ast.INodeVisitor, node *ast.CaseNode) {
  ast.VisitExprs(v, node.GetValues())
  ast.VisitStmt(v, node.GetBody())
}

func visitCastNode(v ast.INodeVisitor, node *ast.CastNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitCondExprNode(v ast.INodeVisitor, node *ast.CondExprNode) {
  ast.VisitExpr(v, node.GetCond())
  ast.VisitExpr(v, node.GetThenExpr())
  if node.GetElseExpr() != nil {
    ast.VisitExpr(v, node.GetElseExpr())
  }
}

func visitContinueNode(v ast.INodeVisitor, node *ast.ContinueNode) {
  // nop
}

func visitDereferenceNode(v ast.INodeVisitor, node *ast.DereferenceNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitDoWhileNode(v ast.INodeVisitor, node *ast.DoWhileNode) {
  ast.VisitStmt(v, node.GetBody())
  ast.VisitExpr(v, node.GetCond())
}

func visitExprStmtNode(v ast.INodeVisitor, node *ast.ExprStmtNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitForNode(v ast.INodeVisitor, node *ast.ForNode) {
  ast.VisitExpr(v, node.GetInit())
  ast.VisitExpr(v, node.GetCond())
  ast.VisitExpr(v, node.GetIncr())
  ast.VisitStmt(v, node.GetBody())
}

func visitFuncallNode(v ast.INodeVisitor, node *ast.FuncallNode) {
  ast.VisitExpr(v, node.GetExpr())
  ast.VisitExprs(v, node.GetArgs())
}

func visitGotoNode(v ast.INodeVisitor, node *ast.GotoNode) {
  // nop
}

func visitIfNode(v ast.INodeVisitor, node *ast.IfNode) {
  ast.VisitExpr(v, node.GetCond())
  ast.VisitStmt(v, node.GetThenBody())
  if node.GetElseBody() != nil {
    ast.VisitStmt(v, node.GetElseBody())
  }
}

func visitIntegerLiteralNode(v ast.INodeVisitor, node *ast.IntegerLiteralNode) {
  // nop
}

func visitLabelNode(v ast.INodeVisitor, node *ast.LabelNode) {
  ast.VisitStmt(v, node.GetStmt())
}

func visitLogicalAndNode(v ast.INodeVisitor, node *ast.LogicalAndNode) {
  ast.VisitExpr(v, node.GetLeft())
  ast.VisitExpr(v, node.GetRight())
}

func visitLogicalOrNode(v ast.INodeVisitor, node *ast.LogicalOrNode) {
  ast.VisitExpr(v, node.GetLeft())
  ast.VisitExpr(v, node.GetRight())
}

func visitMemberNode(v ast.INodeVisitor, node *ast.MemberNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitOpAssignNode(v ast.INodeVisitor, node *ast.OpAssignNode) {
  ast.VisitExpr(v, node.GetLhs())
  ast.VisitExpr(v, node.GetRhs())
}

func visitPrefixOpNode(v ast.INodeVisitor, node *ast.PrefixOpNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitPtrMemberNode(v ast.INodeVisitor, node *ast.PtrMemberNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitReturnNode(v ast.INodeVisitor, node *ast.ReturnNode) {
  if node.GetExpr() != nil {
    ast.VisitExpr(v, node.GetExpr())
  }
}

func visitSizeofExprNode(v ast.INodeVisitor, node *ast.SizeofExprNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitSizeofTypeNode(v ast.INodeVisitor, node *ast.SizeofTypeNode) {
  // nop
}

func visitStringLiteralNode(v ast.INodeVisitor, node *ast.StringLiteralNode) {
  // nop
}

func visitStructNode(v ast.INodeVisitor, node *ast.StructNode) {
  // nop
}

func visitSuffixOpNode(v ast.INodeVisitor, node *ast.SuffixOpNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitSwitchNode(v ast.INodeVisitor, node *ast.SwitchNode) {
  ast.VisitExpr(v, node.GetCond())
  ast.VisitStmts(v, node.GetCases())
}

func visitTypeNode(v ast.INodeVisitor, node *ast.TypeNode) {
  // nop
}

func visitTypedefNode(v ast.INodeVisitor, node *ast.TypedefNode) {
  // nop
}

func visitUnaryOpNode(v ast.INodeVisitor, node *ast.UnaryOpNode) {
  ast.VisitExpr(v, node.GetExpr())
}

func visitUnionNode(v ast.INodeVisitor, node *ast.UnionNode) {
  // nop
}

func visitVariableNode(v ast.INodeVisitor, node *ast.VariableNode) {
  // nop
}

func visitWhileNode(v ast.INodeVisitor, node *ast.WhileNode) {
  ast.VisitExpr(v, node.GetCond())
  ast.VisitStmt(v, node.GetBody())
}

func visitNode(v ast.INodeVisitor, node core.INode) {
  switch typed := node.(type) {
    case *ast.AddressNode: visitAddressNode(v, typed)
    case *ast.ArefNode: visitArefNode(v, typed)
    case *ast.AssignNode: visitAssignNode(v, typed)
    case *ast.BinaryOpNode: visitBinaryOpNode(v, typed)
    case *ast.BlockNode: visitBlockNode(v, typed)
    case *ast.BreakNode: visitBreakNode(v, typed)
    case *ast.CaseNode: visitCaseNode(v, typed)
    case *ast.CastNode: visitCastNode(v, typed)
    case *ast.CondExprNode: visitCondExprNode(v, typed)
    case *ast.ContinueNode: visitContinueNode(v, typed)
    case *ast.DereferenceNode: visitDereferenceNode(v, typed)
    case *ast.DoWhileNode: visitDoWhileNode(v, typed)
    case *ast.ExprStmtNode: visitExprStmtNode(v, typed)
    case *ast.ForNode: visitForNode(v, typed)
    case *ast.FuncallNode: visitFuncallNode(v, typed)
    case *ast.GotoNode: visitGotoNode(v, typed)
    case *ast.IfNode: visitIfNode(v, typed)
    case *ast.IntegerLiteralNode: visitIntegerLiteralNode(v, typed)
    case *ast.LabelNode: visitLabelNode(v, typed)
    case *ast.LogicalAndNode: visitLogicalAndNode(v, typed)
    case *ast.LogicalOrNode: visitLogicalOrNode(v, typed)
    case *ast.MemberNode: visitMemberNode(v, typed)
    case *ast.OpAssignNode: visitOpAssignNode(v, typed)
    case *ast.PrefixOpNode: visitPrefixOpNode(v, typed)
    case *ast.PtrMemberNode: visitPtrMemberNode(v, typed)
    case *ast.ReturnNode: visitReturnNode(v, typed)
    case *ast.SizeofExprNode: visitSizeofExprNode(v, typed)
    case *ast.SizeofTypeNode: visitSizeofTypeNode(v, typed)
    case *ast.StringLiteralNode: visitStringLiteralNode(v, typed)
    case *ast.StructNode: visitStructNode(v, typed)
    case *ast.SuffixOpNode: visitSuffixOpNode(v, typed)
    case *ast.SwitchNode: visitSwitchNode(v, typed)
    case *ast.TypeNode: visitTypeNode(v, typed)
    case *ast.TypedefNode: visitTypedefNode(v, typed)
    case *ast.UnaryOpNode: visitUnaryOpNode(v, typed)
    case *ast.UnionNode: visitUnionNode(v, typed)
    case *ast.VariableNode: visitVariableNode(v, typed)
    case *ast.WhileNode: visitWhileNode(v, typed)
  }
}
