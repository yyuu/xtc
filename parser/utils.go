package parser

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

func asExprNode(x core.INode) core.IExprNode {
  return x.(core.IExprNode)
}

func asExprNodes(xs []core.INode) []core.IExprNode {
  ys := make([]core.IExprNode, len(xs))
  for i := range xs {
    ys[i] = asExprNode(xs[i])
  }
  return ys
}

func asStmtNode(x core.INode) core.IStmtNode {
  return x.(core.IStmtNode)
}

func asStmtNodes(xs []core.INode) []core.IStmtNode {
  ys := make([]core.IStmtNode, len(xs))
  for i := range xs {
    ys[i] = asStmtNode(xs[i])
  }
  return ys
}

func asTypeNode(x core.INode) core.ITypeNode {
  return x.(core.ITypeNode)
}

func asTypeNodes(xs []core.INode) []core.ITypeNode {
  ys := make([]core.ITypeNode, len(xs))
  for i := range xs {
    ys[i] = asTypeNode(xs[i])
  }
  return ys
}

func asSlot(x core.INode) ast.Slot {
  return x.(ast.Slot)
}

func asSlots(xs []core.INode) []ast.Slot {
  ys := make([]ast.Slot, len(xs))
  for i := range xs {
    ys[i] = asSlot(xs[i])
  }
  return ys
}

func asTypeDefinition(x core.INode) core.ITypeDefinition {
  return x.(core.ITypeDefinition)
}

func asTypeDefinitions(xs []core.INode) []core.ITypeDefinition {
  ys := make([]core.ITypeDefinition, len(xs))
  for i := range xs {
    ys[i] = asTypeDefinition(xs[i])
  }
  return ys
}

func asDeclarations(x core.INode) ast.Declarations {
  return x.(ast.Declarations)
}

func asDefinedFunction(x core.IEntity) core.IDefinedFunction {
  return x.(core.IDefinedFunction)
}

func asUndefinedFunction(x core.IEntity) core.IUndefinedFunction {
  return x.(core.IUndefinedFunction)
}

func asVariable(x core.IEntity) core.IVariable {
  return x.(core.IVariable)
}

func asVariables(xs []core.IEntity) []core.IVariable {
  ys := make([]core.IVariable, len(xs))
  for i := range xs {
    ys[i] = asVariable(xs[i])
  }
  return ys
}

func asDefinedVariable(x core.IEntity) core.IDefinedVariable {
  return x.(core.IDefinedVariable)
}

func asDefinedVariables(xs []core.IEntity) []core.IDefinedVariable {
  ys := make([]core.IDefinedVariable, len(xs))
  for i := range xs {
    ys[i] = asDefinedVariable(xs[i])
  }
  return ys
}

func asUndefinedVariable(x core.IEntity) core.IUndefinedVariable {
  return x.(core.IUndefinedVariable)
}

func asConstant(x core.IEntity) core.IConstant {
  return x.(core.IConstant)
}

func asStructNode(x core.INode) ast.StructNode {
  return x.(ast.StructNode)
}

func asUnionNode(x core.INode) ast.UnionNode {
  return x.(ast.UnionNode)
}

func asTypedefNode(x core.INode) ast.TypedefNode {
  return x.(ast.TypedefNode)
}

func asParams(x core.IEntity) entity.Params {
  return x.(entity.Params)
}

func asParameter(x core.IEntity) entity.Parameter {
  return x.(entity.Parameter)
}

func parametersTypeRef(params entity.Params) typesys.ParamTypeRefs {
  paramDescs := params.GetParamDescs()
  ps := make([]core.ITypeRef, len(paramDescs))
  for i := range paramDescs {
    ps[i] = paramDescs[i].GetTypeNode().GetTypeRef()
  }
  return typesys.NewParamTypeRefs(params.GetLocation(), ps, false)
}

func defvars(xs...core.IDefinedVariable) []core.IDefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IDefinedVariable { }
  }
}

func vardecls(xs...core.IUndefinedVariable) []core.IUndefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IUndefinedVariable { }
  }
}

func defuns(xs...core.IDefinedFunction) []core.IDefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IDefinedFunction { }
  }
}

func funcdecls(xs...core.IUndefinedFunction) []core.IUndefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IUndefinedFunction { }
  }
}

func defconsts(xs...core.IConstant) []core.IConstant {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IConstant { }
  }
}

func defstructs(xs...ast.StructNode) []ast.StructNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []ast.StructNode { }
  }
}

func defunions(xs...ast.UnionNode) []ast.UnionNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []ast.UnionNode { }
  }
}

func typedefs(xs...ast.TypedefNode) []ast.TypedefNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []ast.TypedefNode { }
  }
}
