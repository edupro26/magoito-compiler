// Code generated from MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MagoitoParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MagoitoParser.
type MagoitoParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MagoitoParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MagoitoParser#constDeclaration.
	VisitConstDeclaration(ctx *ConstDeclarationContext) interface{}

	// Visit a parse tree produced by MagoitoParser#funDeclaration.
	VisitFunDeclaration(ctx *FunDeclarationContext) interface{}

	// Visit a parse tree produced by MagoitoParser#idBinder.
	VisitIdBinder(ctx *IdBinderContext) interface{}

	// Visit a parse tree produced by MagoitoParser#wildcardBinder.
	VisitWildcardBinder(ctx *WildcardBinderContext) interface{}

	// Visit a parse tree produced by MagoitoParser#arrowType.
	VisitArrowType(ctx *ArrowTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#tupleArrowType.
	VisitTupleArrowType(ctx *TupleArrowTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#basicNonTupleType.
	VisitBasicNonTupleType(ctx *BasicNonTupleTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordNonTupleType.
	VisitRecordNonTupleType(ctx *RecordNonTupleTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#parenNonTupleType.
	VisitParenNonTupleType(ctx *ParenNonTupleTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#intBasicType.
	VisitIntBasicType(ctx *IntBasicTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#boolBasicType.
	VisitBoolBasicType(ctx *BoolBasicTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#stringBasicType.
	VisitStringBasicType(ctx *StringBasicTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#unitBasicType.
	VisitUnitBasicType(ctx *UnitBasicTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordTypeField.
	VisitRecordTypeField(ctx *RecordTypeFieldContext) interface{}

	// Visit a parse tree produced by MagoitoParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#seqExpr.
	VisitSeqExpr(ctx *SeqExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#varDeclControl.
	VisitVarDeclControl(ctx *VarDeclControlContext) interface{}

	// Visit a parse tree produced by MagoitoParser#whileControl.
	VisitWhileControl(ctx *WhileControlContext) interface{}

	// Visit a parse tree produced by MagoitoParser#ifControl.
	VisitIfControl(ctx *IfControlContext) interface{}

	// Visit a parse tree produced by MagoitoParser#assignControl.
	VisitAssignControl(ctx *AssignControlContext) interface{}

	// Visit a parse tree produced by MagoitoParser#orControl.
	VisitOrControl(ctx *OrControlContext) interface{}

	// Visit a parse tree produced by MagoitoParser#varDeclExpr.
	VisitVarDeclExpr(ctx *VarDeclExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#whileExpr.
	VisitWhileExpr(ctx *WhileExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#ifExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#assignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#comparisonExpr.
	VisitComparisonExpr(ctx *ComparisonExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#multiplicativeExpr.
	VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#powerExpr.
	VisitPowerExpr(ctx *PowerExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#projectionExpr.
	VisitProjectionExpr(ctx *ProjectionExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#intLiteralPrimary.
	VisitIntLiteralPrimary(ctx *IntLiteralPrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#stringLiteralPrimary.
	VisitStringLiteralPrimary(ctx *StringLiteralPrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#truePrimary.
	VisitTruePrimary(ctx *TruePrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#falsePrimary.
	VisitFalsePrimary(ctx *FalsePrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#unitPrimary.
	VisitUnitPrimary(ctx *UnitPrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#callPrimary.
	VisitCallPrimary(ctx *CallPrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#identifierPrimary.
	VisitIdentifierPrimary(ctx *IdentifierPrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#parenPrimary.
	VisitParenPrimary(ctx *ParenPrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordPrimary.
	VisitRecordPrimary(ctx *RecordPrimaryContext) interface{}

	// Visit a parse tree produced by MagoitoParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#identCallee.
	VisitIdentCallee(ctx *IdentCalleeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#printCallee.
	VisitPrintCallee(ctx *PrintCalleeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordExpr.
	VisitRecordExpr(ctx *RecordExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordFieldList.
	VisitRecordFieldList(ctx *RecordFieldListContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordField.
	VisitRecordField(ctx *RecordFieldContext) interface{}
}
