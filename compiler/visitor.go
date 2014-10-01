package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

func visitAddressNode(v ast.INodeVisitor, node *ast.AddressNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitArefNode(v ast.INodeVisitor, node *ast.ArefNode) {
  ast.VisitExprNode(v, node.GetExpr())
  ast.VisitExprNode(v, node.GetIndex())
}

func visitAssignNode(v ast.INodeVisitor, node *ast.AssignNode) {
  ast.VisitExprNode(v, node.GetLHS())
  ast.VisitExprNode(v, node.GetRHS())
}

func visitBinaryOpNode(v ast.INodeVisitor, node *ast.BinaryOpNode) {
  ast.VisitExprNode(v, node.GetLeft())
  ast.VisitExprNode(v, node.GetRight())
}

func visitBlockNode(v ast.INodeVisitor, node *ast.BlockNode) {
  vars := node.GetVariables()
  for i := range vars {
    ast.VisitExprNode(v, vars[i].GetInitializer())
  }
  ast.VisitStmtNodes(v, node.GetStmts())
}

func visitBreakNode(v ast.INodeVisitor, node *ast.BreakNode) {
  // nop
}

func visitCaseNode(v ast.INodeVisitor, node *ast.CaseNode) {
  ast.VisitExprNodes(v, node.GetValues())
  ast.VisitStmtNode(v, node.GetBody())
}

func visitCastNode(v ast.INodeVisitor, node *ast.CastNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitCondExprNode(v ast.INodeVisitor, node *ast.CondExprNode) {
  ast.VisitExprNode(v, node.GetCond())
  ast.VisitExprNode(v, node.GetThenExpr())
  if node.GetElseExpr() != nil {
    ast.VisitExprNode(v, node.GetElseExpr())
  }
}

func visitContinueNode(v ast.INodeVisitor, node *ast.ContinueNode) {
  // nop
}

func visitDereferenceNode(v ast.INodeVisitor, node *ast.DereferenceNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitDoWhileNode(v ast.INodeVisitor, node *ast.DoWhileNode) {
  ast.VisitStmtNode(v, node.GetBody())
  ast.VisitExprNode(v, node.GetCond())
}

func visitExprStmtNode(v ast.INodeVisitor, node *ast.ExprStmtNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitForNode(v ast.INodeVisitor, node *ast.ForNode) {
  ast.VisitExprNode(v, node.GetInit())
  ast.VisitExprNode(v, node.GetCond())
  ast.VisitExprNode(v, node.GetIncr())
  ast.VisitStmtNode(v, node.GetBody())
}

func visitFuncallNode(v ast.INodeVisitor, node *ast.FuncallNode) {
  ast.VisitExprNode(v, node.GetExpr())
  ast.VisitExprNodes(v, node.GetArgs())
}

func visitGotoNode(v ast.INodeVisitor, node *ast.GotoNode) {
  // nop
}

func visitIfNode(v ast.INodeVisitor, node *ast.IfNode) {
  ast.VisitExprNode(v, node.GetCond())
  ast.VisitStmtNode(v, node.GetThenBody())
  if node.HasElseBody() {
    ast.VisitStmtNode(v, node.GetElseBody())
  }
}

func visitIntegerLiteralNode(v ast.INodeVisitor, node *ast.IntegerLiteralNode) {
  // nop
}

func visitLabelNode(v ast.INodeVisitor, node *ast.LabelNode) {
  ast.VisitStmtNode(v, node.GetStmt())
}

func visitLogicalAndNode(v ast.INodeVisitor, node *ast.LogicalAndNode) {
  ast.VisitExprNode(v, node.GetLeft())
  ast.VisitExprNode(v, node.GetRight())
}

func visitLogicalOrNode(v ast.INodeVisitor, node *ast.LogicalOrNode) {
  ast.VisitExprNode(v, node.GetLeft())
  ast.VisitExprNode(v, node.GetRight())
}

func visitMemberNode(v ast.INodeVisitor, node *ast.MemberNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitOpAssignNode(v ast.INodeVisitor, node *ast.OpAssignNode) {
  ast.VisitExprNode(v, node.GetLHS())
  ast.VisitExprNode(v, node.GetRHS())
}

func visitPrefixOpNode(v ast.INodeVisitor, node *ast.PrefixOpNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitPtrMemberNode(v ast.INodeVisitor, node *ast.PtrMemberNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitReturnNode(v ast.INodeVisitor, node *ast.ReturnNode) {
  if node.GetExpr() != nil {
    ast.VisitExprNode(v, node.GetExpr())
  }
}

func visitSizeofExprNode(v ast.INodeVisitor, node *ast.SizeofExprNode) {
  ast.VisitExprNode(v, node.GetExpr())
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
  ast.VisitExprNode(v, node.GetExpr())
}

func visitSwitchNode(v ast.INodeVisitor, node *ast.SwitchNode) {
  ast.VisitExprNode(v, node.GetCond())
  ast.VisitStmtNodes(v, node.GetCases())
}

func visitTypedefNode(v ast.INodeVisitor, node *ast.TypedefNode) {
  // nop
}

func visitUnaryOpNode(v ast.INodeVisitor, node *ast.UnaryOpNode) {
  ast.VisitExprNode(v, node.GetExpr())
}

func visitUnionNode(v ast.INodeVisitor, node *ast.UnionNode) {
  // nop
}

func visitVariableNode(v ast.INodeVisitor, node *ast.VariableNode) {
  // nop
}

func visitWhileNode(v ast.INodeVisitor, node *ast.WhileNode) {
  ast.VisitExprNode(v, node.GetCond())
  ast.VisitStmtNode(v, node.GetBody())
}

func visitStmtNode(v ast.INodeVisitor, unknown core.IStmtNode) {
  switch node := unknown.(type) {
    case *ast.BlockNode: visitBlockNode(v, node)
    case *ast.BreakNode: visitBreakNode(v, node)
    case *ast.CaseNode: visitCaseNode(v, node)
    case *ast.ContinueNode: visitContinueNode(v, node)
    case *ast.DoWhileNode: visitDoWhileNode(v, node)
    case *ast.ExprStmtNode: visitExprStmtNode(v, node)
    case *ast.ForNode: visitForNode(v, node)
    case *ast.GotoNode: visitGotoNode(v, node)
    case *ast.IfNode: visitIfNode(v, node)
    case *ast.LabelNode: visitLabelNode(v, node)
    case *ast.ReturnNode: visitReturnNode(v, node)
    case *ast.SwitchNode: visitSwitchNode(v, node)
    case *ast.WhileNode: visitWhileNode(v, node)
    default: {
      panic(fmt.Errorf("unknown stmt node: %s", unknown))
    }
  }
}

func visitExprNode(v ast.INodeVisitor, unknown core.IExprNode) {
  switch node := unknown.(type) {
    case *ast.AddressNode: visitAddressNode(v, node)
    case *ast.ArefNode: visitArefNode(v, node)
    case *ast.AssignNode: visitAssignNode(v, node)
    case *ast.BinaryOpNode: visitBinaryOpNode(v, node)
    case *ast.CondExprNode: visitCondExprNode(v, node)
    case *ast.DereferenceNode: visitDereferenceNode(v, node)
    case *ast.FuncallNode: visitFuncallNode(v, node)
    case *ast.IntegerLiteralNode: visitIntegerLiteralNode(v, node)
    case *ast.LogicalAndNode: visitLogicalAndNode(v, node)
    case *ast.LogicalOrNode: visitLogicalOrNode(v, node)
    case *ast.MemberNode: visitMemberNode(v, node)
    case *ast.OpAssignNode: visitOpAssignNode(v, node)
    case *ast.PrefixOpNode: visitPrefixOpNode(v, node)
    case *ast.PtrMemberNode: visitPtrMemberNode(v, node)
    case *ast.SizeofExprNode: visitSizeofExprNode(v, node)
    case *ast.SizeofTypeNode: visitSizeofTypeNode(v, node)
    case *ast.StringLiteralNode: visitStringLiteralNode(v, node)
    case *ast.SuffixOpNode: visitSuffixOpNode(v, node)
    case *ast.UnaryOpNode: visitUnaryOpNode(v, node)
    case *ast.VariableNode: visitVariableNode(v, node)
    default: {
      panic(fmt.Errorf("unknown expr node: %s", unknown))
    }
  }
}

func visitTypeDefinition(v ast.INodeVisitor, unknown core.ITypeDefinition) {
  switch node := unknown.(type) {
    case *ast.StructNode: visitStructNode(v, node)
    case *ast.TypedefNode: visitTypedefNode(v, node)
    case *ast.UnionNode: visitUnionNode(v, node)
    default: {
      panic(fmt.Errorf("unknown type definition: %s", unknown))
    }
  }
}

func visitEntity(v entity.IEntityVisitor, unknown core.IEntity) {
  switch unknown.(type) {
    case *entity.DefinedVariable: { }
    case *entity.UndefinedVariable: { }
    case *entity.Constant: { }
    case *entity.DefinedFunction: { }
    case *entity.UndefinedFunction: { }
    default: {
      panic(fmt.Errorf("unknown entity: %s", unknown))
    }
  }
}
