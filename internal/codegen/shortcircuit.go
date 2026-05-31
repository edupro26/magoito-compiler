package codegen

import (
	"fmt"
	"magoito-compiler/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (ctx *Context) compileBinaryExpr(e *ast.BinaryExpr) (value.Value, ast.Type, error) {
	if e.Op == "&&" || e.Op == "||" {
		return ctx.compileShortCircuit(e.Left, e.Right, e.Op == "&&")
	}

	leftVal, leftType, err := ctx.compileExpr(e.Left)
	if err != nil {
		return nil, nil, err
	}
	lt, ok := leftType.(*ast.BasicType)
	if !ok {
		return nil, nil, fmt.Errorf("unsupported binary type")
	}
	if e.Op != "==" && e.Op != "!=" {
		if lt.Name == "String" {
			return nil, nil, fmt.Errorf("unsupported binary op for String")
		}
	}
	rightVal, _, err := ctx.compileExpr(e.Right)
	if err != nil {
		return nil, nil, err
	}

	switch e.Op {
	case "+":
		return ctx.currentBlock.NewAdd(leftVal, rightVal), leftType, nil
	case "-":
		return ctx.currentBlock.NewSub(leftVal, rightVal), leftType, nil
	case "*":
		return ctx.currentBlock.NewMul(leftVal, rightVal), leftType, nil
	case "/":
		return ctx.currentBlock.NewSDiv(leftVal, rightVal), leftType, nil
	case "%":
		return ctx.currentBlock.NewSRem(leftVal, rightVal), leftType, nil
	case "^":
		return ctx.currentBlock.NewCall(ctx.pow, leftVal, rightVal), leftType, nil
	case "<":
		return ctx.currentBlock.NewICmp(enum.IPredSLT, leftVal, rightVal), &ast.BasicType{Name: "Bool"}, nil
	case "<=":
		return ctx.currentBlock.NewICmp(enum.IPredSLE, leftVal, rightVal), &ast.BasicType{Name: "Bool"}, nil
	case ">":
		return ctx.currentBlock.NewICmp(enum.IPredSGT, leftVal, rightVal), &ast.BasicType{Name: "Bool"}, nil
	case ">=":
		return ctx.currentBlock.NewICmp(enum.IPredSGE, leftVal, rightVal), &ast.BasicType{Name: "Bool"}, nil
	case "==", "!=":
		if lt.Name == "String" {
			cmp := ctx.currentBlock.NewCall(ctx.strcmp, leftVal, rightVal)
			pred := enum.IPredEQ
			if e.Op == "!=" {
				pred = enum.IPredNE
			}
			zero := constant.NewInt(types.I32, 0)
			return ctx.currentBlock.NewICmp(pred, cmp, zero), &ast.BasicType{Name: "Bool"}, nil
		}
		pred := enum.IPredEQ
		if e.Op == "!=" {
			pred = enum.IPredNE
		}
		return ctx.currentBlock.NewICmp(pred, leftVal, rightVal), &ast.BasicType{Name: "Bool"}, nil
	}
	return nil, nil, fmt.Errorf("unsupported binary op: %s", e.Op)
}

func (ctx *Context) compileShortCircuit(left, right ast.Expr, isAnd bool) (value.Value, ast.Type, error) {
	leftVal, _, err := ctx.compileExpr(left)
	if err != nil {
		return nil, nil, err
	}

	ctx.labelCounter++
	thenBlock := ctx.currentFunc.NewBlock(fmt.Sprintf("sc_then%d", ctx.labelCounter))
	ctx.labelCounter++
	elseBlock := ctx.currentFunc.NewBlock(fmt.Sprintf("sc_else%d", ctx.labelCounter))
	ctx.labelCounter++
	mergeBlock := ctx.currentFunc.NewBlock(fmt.Sprintf("sc_merge%d", ctx.labelCounter))

	if isAnd {
		ctx.currentBlock.NewCondBr(leftVal, thenBlock, elseBlock)
	} else {
		ctx.currentBlock.NewCondBr(leftVal, elseBlock, thenBlock)
	}

	ctx.currentBlock = thenBlock
	rightVal, _, err := ctx.compileExpr(right)
	if err != nil {
		return nil, nil, err
	}
	thenEnd := ctx.currentBlock
	if thenEnd.Term == nil {
		thenEnd.NewBr(mergeBlock)
	}

	ctx.currentBlock = elseBlock
	shortVal := constant.NewInt(types.I1, 0)
	if !isAnd {
		shortVal = constant.NewInt(types.I1, 1)
	}
	elseEnd := ctx.currentBlock
	if elseEnd.Term == nil {
		elseEnd.NewBr(mergeBlock)
	}

	ctx.currentBlock = mergeBlock
	phi := mergeBlock.NewPhi(
		ir.NewIncoming(rightVal, thenEnd),
		ir.NewIncoming(shortVal, elseEnd),
	)
	phi.Typ = types.I1
	return phi, &ast.BasicType{Name: "Bool"}, nil
}
