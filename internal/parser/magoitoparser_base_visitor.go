// Code generated from MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MagoitoParser
import "github.com/antlr4-go/antlr/v4"

type BaseMagoitoParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMagoitoParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitConstDeclaration(ctx *ConstDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitFunDeclaration(ctx *FunDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitIdBinder(ctx *IdBinderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitWildcardBinder(ctx *WildcardBinderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitArrowType(ctx *ArrowTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitTupleArrowType(ctx *TupleArrowTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitBasicNonTupleType(ctx *BasicNonTupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitRecordNonTupleType(ctx *RecordNonTupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitParenNonTupleType(ctx *ParenNonTupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitIntBasicType(ctx *IntBasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitBoolBasicType(ctx *BoolBasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitStringBasicType(ctx *StringBasicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitUnitBasicType(ctx *UnitBasicTypeContext) interface{} {
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

func (v *BaseMagoitoParserVisitor) VisitVarDeclControl(ctx *VarDeclControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitWhileControl(ctx *WhileControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitIfControl(ctx *IfControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitAssignControl(ctx *AssignControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitOrControl(ctx *OrControlContext) interface{} {
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

func (v *BaseMagoitoParserVisitor) VisitIntLiteralPrimary(ctx *IntLiteralPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitStringLiteralPrimary(ctx *StringLiteralPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitTruePrimary(ctx *TruePrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitFalsePrimary(ctx *FalsePrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitUnitPrimary(ctx *UnitPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitCallPrimary(ctx *CallPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitIdentifierPrimary(ctx *IdentifierPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitParenPrimary(ctx *ParenPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitRecordPrimary(ctx *RecordPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitIdentCallee(ctx *IdentCalleeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMagoitoParserVisitor) VisitPrintCallee(ctx *PrintCalleeContext) interface{} {
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
