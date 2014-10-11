package compiler

import (
  "fmt"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
)

func visitAddressNode(v xtc_ast.INodeVisitor, node *xtc_ast.AddressNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitArefNode(v xtc_ast.INodeVisitor, node *xtc_ast.ArefNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
  xtc_ast.VisitExprNode(v, node.GetIndex())
}

func visitAssignNode(v xtc_ast.INodeVisitor, node *xtc_ast.AssignNode) {
  xtc_ast.VisitExprNode(v, node.GetLHS())
  xtc_ast.VisitExprNode(v, node.GetRHS())
}

func visitBinaryOpNode(v xtc_ast.INodeVisitor, node *xtc_ast.BinaryOpNode) {
  xtc_ast.VisitExprNode(v, node.GetLeft())
  xtc_ast.VisitExprNode(v, node.GetRight())
}

func visitBlockNode(v xtc_ast.INodeVisitor, node *xtc_ast.BlockNode) {
  vars := node.GetVariables()
  for i := range vars {
    if vars[i].HasInitializer() {
      xtc_ast.VisitExprNode(v, vars[i].GetInitializer())
    }
  }
  xtc_ast.VisitStmtNodes(v, node.GetStmts())
}

func visitBreakNode(v xtc_ast.INodeVisitor, node *xtc_ast.BreakNode) {
  // nop
}

func visitCaseNode(v xtc_ast.INodeVisitor, node *xtc_ast.CaseNode) {
  xtc_ast.VisitExprNodes(v, node.GetValues())
  xtc_ast.VisitStmtNode(v, node.GetBody())
}

func visitCastNode(v xtc_ast.INodeVisitor, node *xtc_ast.CastNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitCondExprNode(v xtc_ast.INodeVisitor, node *xtc_ast.CondExprNode) {
  xtc_ast.VisitExprNode(v, node.GetCond())
  xtc_ast.VisitExprNode(v, node.GetThenExpr())
  if node.GetElseExpr() != nil {
    xtc_ast.VisitExprNode(v, node.GetElseExpr())
  }
}

func visitContinueNode(v xtc_ast.INodeVisitor, node *xtc_ast.ContinueNode) {
  // nop
}

func visitDereferenceNode(v xtc_ast.INodeVisitor, node *xtc_ast.DereferenceNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitDoWhileNode(v xtc_ast.INodeVisitor, node *xtc_ast.DoWhileNode) {
  xtc_ast.VisitStmtNode(v, node.GetBody())
  xtc_ast.VisitExprNode(v, node.GetCond())
}

func visitExprStmtNode(v xtc_ast.INodeVisitor, node *xtc_ast.ExprStmtNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitForNode(v xtc_ast.INodeVisitor, node *xtc_ast.ForNode) {
  xtc_ast.VisitExprNode(v, node.GetInit())
  xtc_ast.VisitExprNode(v, node.GetCond())
  xtc_ast.VisitExprNode(v, node.GetIncr())
  xtc_ast.VisitStmtNode(v, node.GetBody())
}

func visitFuncallNode(v xtc_ast.INodeVisitor, node *xtc_ast.FuncallNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
  xtc_ast.VisitExprNodes(v, node.GetArgs())
}

func visitGotoNode(v xtc_ast.INodeVisitor, node *xtc_ast.GotoNode) {
  // nop
}

func visitIfNode(v xtc_ast.INodeVisitor, node *xtc_ast.IfNode) {
  xtc_ast.VisitExprNode(v, node.GetCond())
  xtc_ast.VisitStmtNode(v, node.GetThenBody())
  if node.HasElseBody() {
    xtc_ast.VisitStmtNode(v, node.GetElseBody())
  }
}

func visitIntegerLiteralNode(v xtc_ast.INodeVisitor, node *xtc_ast.IntegerLiteralNode) {
  // nop
}

func visitLabelNode(v xtc_ast.INodeVisitor, node *xtc_ast.LabelNode) {
  xtc_ast.VisitStmtNode(v, node.GetStmt())
}

func visitLogicalAndNode(v xtc_ast.INodeVisitor, node *xtc_ast.LogicalAndNode) {
  xtc_ast.VisitExprNode(v, node.GetLeft())
  xtc_ast.VisitExprNode(v, node.GetRight())
}

func visitLogicalOrNode(v xtc_ast.INodeVisitor, node *xtc_ast.LogicalOrNode) {
  xtc_ast.VisitExprNode(v, node.GetLeft())
  xtc_ast.VisitExprNode(v, node.GetRight())
}

func visitMemberNode(v xtc_ast.INodeVisitor, node *xtc_ast.MemberNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitOpAssignNode(v xtc_ast.INodeVisitor, node *xtc_ast.OpAssignNode) {
  xtc_ast.VisitExprNode(v, node.GetLHS())
  xtc_ast.VisitExprNode(v, node.GetRHS())
}

func visitPrefixOpNode(v xtc_ast.INodeVisitor, node *xtc_ast.PrefixOpNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitPtrMemberNode(v xtc_ast.INodeVisitor, node *xtc_ast.PtrMemberNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitReturnNode(v xtc_ast.INodeVisitor, node *xtc_ast.ReturnNode) {
  if node.HasExpr() {
    xtc_ast.VisitExprNode(v, node.GetExpr())
  }
}

func visitSizeofExprNode(v xtc_ast.INodeVisitor, node *xtc_ast.SizeofExprNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitSizeofTypeNode(v xtc_ast.INodeVisitor, node *xtc_ast.SizeofTypeNode) {
  // nop
}

func visitStringLiteralNode(v xtc_ast.INodeVisitor, node *xtc_ast.StringLiteralNode) {
  // nop
}

func visitStructNode(v xtc_ast.INodeVisitor, node *xtc_ast.StructNode) {
  // nop
}

func visitSuffixOpNode(v xtc_ast.INodeVisitor, node *xtc_ast.SuffixOpNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitSwitchNode(v xtc_ast.INodeVisitor, node *xtc_ast.SwitchNode) {
  xtc_ast.VisitExprNode(v, node.GetCond())
  xtc_ast.VisitStmtNodes(v, node.GetCases())
}

func visitTypedefNode(v xtc_ast.INodeVisitor, node *xtc_ast.TypedefNode) {
  // nop
}

func visitUnaryOpNode(v xtc_ast.INodeVisitor, node *xtc_ast.UnaryOpNode) {
  xtc_ast.VisitExprNode(v, node.GetExpr())
}

func visitUnionNode(v xtc_ast.INodeVisitor, node *xtc_ast.UnionNode) {
  // nop
}

func visitVariableNode(v xtc_ast.INodeVisitor, node *xtc_ast.VariableNode) {
  // nop
}

func visitWhileNode(v xtc_ast.INodeVisitor, node *xtc_ast.WhileNode) {
  xtc_ast.VisitExprNode(v, node.GetCond())
  xtc_ast.VisitStmtNode(v, node.GetBody())
}

func visitStmtNode(v xtc_ast.INodeVisitor, unknown xtc_core.IStmtNode) {
  switch node := unknown.(type) {
    case *xtc_ast.BlockNode: visitBlockNode(v, node)
    case *xtc_ast.BreakNode: visitBreakNode(v, node)
    case *xtc_ast.CaseNode: visitCaseNode(v, node)
    case *xtc_ast.ContinueNode: visitContinueNode(v, node)
    case *xtc_ast.DoWhileNode: visitDoWhileNode(v, node)
    case *xtc_ast.ExprStmtNode: visitExprStmtNode(v, node)
    case *xtc_ast.ForNode: visitForNode(v, node)
    case *xtc_ast.GotoNode: visitGotoNode(v, node)
    case *xtc_ast.IfNode: visitIfNode(v, node)
    case *xtc_ast.LabelNode: visitLabelNode(v, node)
    case *xtc_ast.ReturnNode: visitReturnNode(v, node)
    case *xtc_ast.SwitchNode: visitSwitchNode(v, node)
    case *xtc_ast.WhileNode: visitWhileNode(v, node)
    default: {
      panic(fmt.Errorf("unknown stmt node: %s", unknown))
    }
  }
}

func visitExprNode(v xtc_ast.INodeVisitor, unknown xtc_core.IExprNode) {
  switch node := unknown.(type) {
    case *xtc_ast.AddressNode: visitAddressNode(v, node)
    case *xtc_ast.ArefNode: visitArefNode(v, node)
    case *xtc_ast.AssignNode: visitAssignNode(v, node)
    case *xtc_ast.BinaryOpNode: visitBinaryOpNode(v, node)
    case *xtc_ast.CondExprNode: visitCondExprNode(v, node)
    case *xtc_ast.DereferenceNode: visitDereferenceNode(v, node)
    case *xtc_ast.FuncallNode: visitFuncallNode(v, node)
    case *xtc_ast.IntegerLiteralNode: visitIntegerLiteralNode(v, node)
    case *xtc_ast.LogicalAndNode: visitLogicalAndNode(v, node)
    case *xtc_ast.LogicalOrNode: visitLogicalOrNode(v, node)
    case *xtc_ast.MemberNode: visitMemberNode(v, node)
    case *xtc_ast.OpAssignNode: visitOpAssignNode(v, node)
    case *xtc_ast.PrefixOpNode: visitPrefixOpNode(v, node)
    case *xtc_ast.PtrMemberNode: visitPtrMemberNode(v, node)
    case *xtc_ast.SizeofExprNode: visitSizeofExprNode(v, node)
    case *xtc_ast.SizeofTypeNode: visitSizeofTypeNode(v, node)
    case *xtc_ast.StringLiteralNode: visitStringLiteralNode(v, node)
    case *xtc_ast.SuffixOpNode: visitSuffixOpNode(v, node)
    case *xtc_ast.UnaryOpNode: visitUnaryOpNode(v, node)
    case *xtc_ast.VariableNode: visitVariableNode(v, node)
    default: {
      panic(fmt.Errorf("unknown expr node: %s", unknown))
    }
  }
}

func visitTypeDefinition(v xtc_ast.INodeVisitor, unknown xtc_core.ITypeDefinition) {
  switch node := unknown.(type) {
    case *xtc_ast.StructNode: visitStructNode(v, node)
    case *xtc_ast.TypedefNode: visitTypedefNode(v, node)
    case *xtc_ast.UnionNode: visitUnionNode(v, node)
    default: {
      panic(fmt.Errorf("unknown type definition: %s", unknown))
    }
  }
}

func visitEntity(v xtc_entity.IEntityVisitor, unknown xtc_core.IEntity) {
  switch unknown.(type) {
    case *xtc_entity.DefinedVariable: { }
    case *xtc_entity.UndefinedVariable: { }
    case *xtc_entity.Constant: { }
    case *xtc_entity.DefinedFunction: { }
    case *xtc_entity.UndefinedFunction: { }
    default: {
      panic(fmt.Errorf("unknown xtc_entity: %s", unknown))
    }
  }
}
