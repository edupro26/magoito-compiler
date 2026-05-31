package codegen

import (
	"fmt"
	"magoito-compiler/internal/ast"

	"github.com/llir/llvm/ir/types"
)

func (ctx *Context) compileType(t ast.Type) (types.Type, error) {
	switch typ := t.(type) {
	case *ast.BasicType:
		switch typ.Name {
		case "Int":
			return types.I32, nil
		case "Bool":
			return types.I1, nil
		case "Unit":
			return types.I1, nil
		case "String":
			return types.I8Ptr, nil
		}
	case *ast.FunctionType:
		params := make([]types.Type, len(typ.Params))
		for i, paramTyp := range typ.Params {
			typ, err := ctx.compileType(paramTyp)
			if err != nil {
				return nil, err
			}
			params[i] = typ
		}
		ret, err := ctx.compileType(typ.Return)
		if err != nil {
			return nil, err
		}
		return types.NewFunc(ret, params...), nil
	case *ast.RecordType:
		return nil, fmt.Errorf("records are not supported in codegen")
	}
	return nil, nil
}

func (ctx *Context) compileFunction(f *ast.FunDecl) error {
	ctx.scope = newScope(nil)
	ctx.currentFunc = ctx.functions[f.Name]
	ctx.currentBlock = ctx.currentFunc.NewBlock("entry")
	ctx.labelCounter = 0
	ft, _ := f.Type.(*ast.FunctionType)
	for i, paramTyp := range ft.Params {
		if f.Params[i] == "_" {
			continue
		}

		typ, err := ctx.compileType(paramTyp)
		if err != nil {
			return err
		}
		ptr := ctx.currentBlock.NewAlloca(typ)
		ctx.currentBlock.NewStore(ctx.currentFunc.Params[i], ptr)
		ctx.scope.define(f.Params[i], ptr, paramTyp)
	}
	val, _, err := ctx.compileExpr(f.Body)
	if err != nil {
		return err
	}
	ctx.currentBlock.NewRet(val)

	ctx.currentFunc = nil
	ctx.currentBlock = nil
	ctx.scope = nil
	return nil
}
