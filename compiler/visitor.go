package compiler

import (
  "fmt"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
)

func visitAddressNode(v bs_ast.INodeVisitor, node *bs_ast.AddressNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitArefNode(v bs_ast.INodeVisitor, node *bs_ast.ArefNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
  bs_ast.VisitExprNode(v, node.GetIndex())
}

func visitAssignNode(v bs_ast.INodeVisitor, node *bs_ast.AssignNode) {
  bs_ast.VisitExprNode(v, node.GetLHS())
  bs_ast.VisitExprNode(v, node.GetRHS())
}

func visitBinaryOpNode(v bs_ast.INodeVisitor, node *bs_ast.BinaryOpNode) {
  bs_ast.VisitExprNode(v, node.GetLeft())
  bs_ast.VisitExprNode(v, node.GetRight())
}

func visitBlockNode(v bs_ast.INodeVisitor, node *bs_ast.BlockNode) {
  vars := node.GetVariables()
  for i := range vars {
    bs_ast.VisitExprNode(v, vars[i].GetInitializer())
  }
  bs_ast.VisitStmtNodes(v, node.GetStmts())
}

func visitBreakNode(v bs_ast.INodeVisitor, node *bs_ast.BreakNode) {
  // nop
}

func visitCaseNode(v bs_ast.INodeVisitor, node *bs_ast.CaseNode) {
  bs_ast.VisitExprNodes(v, node.GetValues())
  bs_ast.VisitStmtNode(v, node.GetBody())
}

func visitCastNode(v bs_ast.INodeVisitor, node *bs_ast.CastNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitCondExprNode(v bs_ast.INodeVisitor, node *bs_ast.CondExprNode) {
  bs_ast.VisitExprNode(v, node.GetCond())
  bs_ast.VisitExprNode(v, node.GetThenExpr())
  if node.GetElseExpr() != nil {
    bs_ast.VisitExprNode(v, node.GetElseExpr())
  }
}

func visitContinueNode(v bs_ast.INodeVisitor, node *bs_ast.ContinueNode) {
  // nop
}

func visitDereferenceNode(v bs_ast.INodeVisitor, node *bs_ast.DereferenceNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitDoWhileNode(v bs_ast.INodeVisitor, node *bs_ast.DoWhileNode) {
  bs_ast.VisitStmtNode(v, node.GetBody())
  bs_ast.VisitExprNode(v, node.GetCond())
}

func visitExprStmtNode(v bs_ast.INodeVisitor, node *bs_ast.ExprStmtNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitForNode(v bs_ast.INodeVisitor, node *bs_ast.ForNode) {
  bs_ast.VisitExprNode(v, node.GetInit())
  bs_ast.VisitExprNode(v, node.GetCond())
  bs_ast.VisitExprNode(v, node.GetIncr())
  bs_ast.VisitStmtNode(v, node.GetBody())
}

func visitFuncallNode(v bs_ast.INodeVisitor, node *bs_ast.FuncallNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
  bs_ast.VisitExprNodes(v, node.GetArgs())
}

func visitGotoNode(v bs_ast.INodeVisitor, node *bs_ast.GotoNode) {
  // nop
}

func visitIfNode(v bs_ast.INodeVisitor, node *bs_ast.IfNode) {
  bs_ast.VisitExprNode(v, node.GetCond())
  bs_ast.VisitStmtNode(v, node.GetThenBody())
  if node.HasElseBody() {
    bs_ast.VisitStmtNode(v, node.GetElseBody())
  }
}

func visitIntegerLiteralNode(v bs_ast.INodeVisitor, node *bs_ast.IntegerLiteralNode) {
  // nop
}

func visitLabelNode(v bs_ast.INodeVisitor, node *bs_ast.LabelNode) {
  bs_ast.VisitStmtNode(v, node.GetStmt())
}

func visitLogicalAndNode(v bs_ast.INodeVisitor, node *bs_ast.LogicalAndNode) {
  bs_ast.VisitExprNode(v, node.GetLeft())
  bs_ast.VisitExprNode(v, node.GetRight())
}

func visitLogicalOrNode(v bs_ast.INodeVisitor, node *bs_ast.LogicalOrNode) {
  bs_ast.VisitExprNode(v, node.GetLeft())
  bs_ast.VisitExprNode(v, node.GetRight())
}

func visitMemberNode(v bs_ast.INodeVisitor, node *bs_ast.MemberNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitOpAssignNode(v bs_ast.INodeVisitor, node *bs_ast.OpAssignNode) {
  bs_ast.VisitExprNode(v, node.GetLHS())
  bs_ast.VisitExprNode(v, node.GetRHS())
}

func visitPrefixOpNode(v bs_ast.INodeVisitor, node *bs_ast.PrefixOpNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitPtrMemberNode(v bs_ast.INodeVisitor, node *bs_ast.PtrMemberNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitReturnNode(v bs_ast.INodeVisitor, node *bs_ast.ReturnNode) {
  if node.GetExpr() != nil {
    bs_ast.VisitExprNode(v, node.GetExpr())
  }
}

func visitSizeofExprNode(v bs_ast.INodeVisitor, node *bs_ast.SizeofExprNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitSizeofTypeNode(v bs_ast.INodeVisitor, node *bs_ast.SizeofTypeNode) {
  // nop
}

func visitStringLiteralNode(v bs_ast.INodeVisitor, node *bs_ast.StringLiteralNode) {
  // nop
}

func visitStructNode(v bs_ast.INodeVisitor, node *bs_ast.StructNode) {
  // nop
}

func visitSuffixOpNode(v bs_ast.INodeVisitor, node *bs_ast.SuffixOpNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitSwitchNode(v bs_ast.INodeVisitor, node *bs_ast.SwitchNode) {
  bs_ast.VisitExprNode(v, node.GetCond())
  bs_ast.VisitStmtNodes(v, node.GetCases())
}

func visitTypedefNode(v bs_ast.INodeVisitor, node *bs_ast.TypedefNode) {
  // nop
}

func visitUnaryOpNode(v bs_ast.INodeVisitor, node *bs_ast.UnaryOpNode) {
  bs_ast.VisitExprNode(v, node.GetExpr())
}

func visitUnionNode(v bs_ast.INodeVisitor, node *bs_ast.UnionNode) {
  // nop
}

func visitVariableNode(v bs_ast.INodeVisitor, node *bs_ast.VariableNode) {
  // nop
}

func visitWhileNode(v bs_ast.INodeVisitor, node *bs_ast.WhileNode) {
  bs_ast.VisitExprNode(v, node.GetCond())
  bs_ast.VisitStmtNode(v, node.GetBody())
}

func visitStmtNode(v bs_ast.INodeVisitor, unknown bs_core.IStmtNode) {
  switch node := unknown.(type) {
    case *bs_ast.BlockNode: visitBlockNode(v, node)
    case *bs_ast.BreakNode: visitBreakNode(v, node)
    case *bs_ast.CaseNode: visitCaseNode(v, node)
    case *bs_ast.ContinueNode: visitContinueNode(v, node)
    case *bs_ast.DoWhileNode: visitDoWhileNode(v, node)
    case *bs_ast.ExprStmtNode: visitExprStmtNode(v, node)
    case *bs_ast.ForNode: visitForNode(v, node)
    case *bs_ast.GotoNode: visitGotoNode(v, node)
    case *bs_ast.IfNode: visitIfNode(v, node)
    case *bs_ast.LabelNode: visitLabelNode(v, node)
    case *bs_ast.ReturnNode: visitReturnNode(v, node)
    case *bs_ast.SwitchNode: visitSwitchNode(v, node)
    case *bs_ast.WhileNode: visitWhileNode(v, node)
    default: {
      panic(fmt.Errorf("unknown stmt node: %s", unknown))
    }
  }
}

func visitExprNode(v bs_ast.INodeVisitor, unknown bs_core.IExprNode) {
  switch node := unknown.(type) {
    case *bs_ast.AddressNode: visitAddressNode(v, node)
    case *bs_ast.ArefNode: visitArefNode(v, node)
    case *bs_ast.AssignNode: visitAssignNode(v, node)
    case *bs_ast.BinaryOpNode: visitBinaryOpNode(v, node)
    case *bs_ast.CondExprNode: visitCondExprNode(v, node)
    case *bs_ast.DereferenceNode: visitDereferenceNode(v, node)
    case *bs_ast.FuncallNode: visitFuncallNode(v, node)
    case *bs_ast.IntegerLiteralNode: visitIntegerLiteralNode(v, node)
    case *bs_ast.LogicalAndNode: visitLogicalAndNode(v, node)
    case *bs_ast.LogicalOrNode: visitLogicalOrNode(v, node)
    case *bs_ast.MemberNode: visitMemberNode(v, node)
    case *bs_ast.OpAssignNode: visitOpAssignNode(v, node)
    case *bs_ast.PrefixOpNode: visitPrefixOpNode(v, node)
    case *bs_ast.PtrMemberNode: visitPtrMemberNode(v, node)
    case *bs_ast.SizeofExprNode: visitSizeofExprNode(v, node)
    case *bs_ast.SizeofTypeNode: visitSizeofTypeNode(v, node)
    case *bs_ast.StringLiteralNode: visitStringLiteralNode(v, node)
    case *bs_ast.SuffixOpNode: visitSuffixOpNode(v, node)
    case *bs_ast.UnaryOpNode: visitUnaryOpNode(v, node)
    case *bs_ast.VariableNode: visitVariableNode(v, node)
    default: {
      panic(fmt.Errorf("unknown expr node: %s", unknown))
    }
  }
}

func visitTypeDefinition(v bs_ast.INodeVisitor, unknown bs_core.ITypeDefinition) {
  switch node := unknown.(type) {
    case *bs_ast.StructNode: visitStructNode(v, node)
    case *bs_ast.TypedefNode: visitTypedefNode(v, node)
    case *bs_ast.UnionNode: visitUnionNode(v, node)
    default: {
      panic(fmt.Errorf("unknown type definition: %s", unknown))
    }
  }
}

func visitEntity(v bs_entity.IEntityVisitor, unknown bs_core.IEntity) {
  switch unknown.(type) {
    case *bs_entity.DefinedVariable: { }
    case *bs_entity.UndefinedVariable: { }
    case *bs_entity.Constant: { }
    case *bs_entity.DefinedFunction: { }
    case *bs_entity.UndefinedFunction: { }
    default: {
      panic(fmt.Errorf("unknown bs_entity: %s", unknown))
    }
  }
}
