package codegen

import (
	"magoito-compiler/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (ctx *Context) compileRuntimeHelpers() {
	ctx.printf = ctx.module.NewFunc(
		"printf", types.I32, ir.NewParam("fmt", types.I8Ptr),
	)
	ctx.printf.Sig.Variadic = true
	ctx.strcmp = ctx.module.NewFunc(
		"strcmp", types.I32, ir.NewParam("a", types.I8Ptr), ir.NewParam("b", types.I8Ptr),
	)
	powFn := ctx.module.NewFunc(
		"powi", types.I32, ir.NewParam("base", types.I32), ir.NewParam("exp", types.I32),
	)
	ctx.pow = powFn

	entry := powFn.NewBlock("entry")
	result := entry.NewAlloca(types.I32)
	eVar := entry.NewAlloca(types.I32)
	entry.NewStore(constant.NewInt(types.I32, 1), result)
	entry.NewStore(powFn.Params[1], eVar)
	loop := powFn.NewBlock("loop")
	entry.NewBr(loop)

	cond := loop.NewLoad(types.I32, eVar)
	cmp := loop.NewICmp(enum.IPredSGT, cond, constant.NewInt(types.I32, 0))
	body := powFn.NewBlock("body")
	end := powFn.NewBlock("end")
	loop.NewCondBr(cmp, body, end)

	resVal := body.NewLoad(types.I32, result)
	mul := body.NewMul(resVal, powFn.Params[0])
	body.NewStore(mul, result)
	dec := body.NewSub(cond, constant.NewInt(types.I32, 1))
	body.NewStore(dec, eVar)
	body.NewBr(loop)

	out := end.NewLoad(types.I32, result)
	end.NewRet(out)
}

func (ctx *Context) compilePrintln(v value.Value, t ast.Type) error {
	switch typ := t.(type) {
	case *ast.BasicType:
		switch typ.Name {
		case "Int":
			gl, err := ctx.getStringGlobal("%d")
			if err != nil {
				return err
			}
			ind := constant.NewInt(types.I32, 0)
			typ := types.NewArray(uint64(len("%d")+1), types.I8)
			fmtPtr := ctx.currentBlock.NewGetElementPtr(typ, gl, ind, ind)
			ctx.currentBlock.NewCall(ctx.printf, fmtPtr, v)
		case "Bool":
			gl, err := ctx.getStringGlobal("true")
			if err != nil {
				return err
			}
			ind := constant.NewInt(types.I32, 0)
			typ := types.NewArray(uint64(len("true")+1), types.I8)
			truePtr := ctx.currentBlock.NewGetElementPtr(typ, gl, ind, ind)

			gl, err = ctx.getStringGlobal("false")
			if err != nil {
				return err
			}
			typ = types.NewArray(uint64(len("false")+1), types.I8)
			falsePtr := ctx.currentBlock.NewGetElementPtr(typ, gl, ind, ind)

			choice := ctx.currentBlock.NewSelect(v, truePtr, falsePtr)
			gl, err = ctx.getStringGlobal("%s")
			if err != nil {
				return err
			}
			typ = types.NewArray(uint64(len("%s")+1), types.I8)
			fmtPtr := ctx.currentBlock.NewGetElementPtr(typ, gl, ind, ind)
			ctx.currentBlock.NewCall(ctx.printf, fmtPtr, choice)
		case "String":
			gl, err := ctx.getStringGlobal("%s")
			if err != nil {
				return err
			}
			ind := constant.NewInt(types.I32, 0)
			typ := types.NewArray(uint64(len("%s")+1), types.I8)
			fmtPtr := ctx.currentBlock.NewGetElementPtr(typ, gl, ind, ind)
			ctx.currentBlock.NewCall(ctx.printf, fmtPtr, v)
		case "Unit":
			gl, err := ctx.getStringGlobal("unit")
			if err != nil {
				return err
			}
			ind := constant.NewInt(types.I32, 0)
			typ := types.NewArray(uint64(len("unit")+1), types.I8)
			fmtPtr := ctx.currentBlock.NewGetElementPtr(typ, gl, ind, ind)
			ctx.currentBlock.NewCall(ctx.printf, fmtPtr)
			return nil
		}
	}
	return nil
}
