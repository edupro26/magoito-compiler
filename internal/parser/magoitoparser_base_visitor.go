// Code generated from MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MagoitoParser
import "github.com/antlr4-go/antlr/v4"

type BaseMagoitoParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMagoitoParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitBinder(ctx *BinderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitArrowType(ctx *ArrowTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitTupleArrowType(ctx *TupleArrowTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitNonTupleType(ctx *NonTupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitBasicType(ctx *BasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitRecordType(ctx *RecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitRecordTypeField(ctx *RecordTypeFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitSeqExpr(ctx *SeqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitControlExpr(ctx *ControlExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitVarDeclExpr(ctx *VarDeclExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitWhileExpr(ctx *WhileExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitComparisonExpr(ctx *ComparisonExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitPowerExpr(ctx *PowerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitProjectionExpr(ctx *ProjectionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitCallee(ctx *CalleeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitRecordExpr(ctx *RecordExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitRecordFieldList(ctx *RecordFieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitRecordField(ctx *RecordFieldContext) interface{} {
	return v.VisitChildren(ctx)
}
