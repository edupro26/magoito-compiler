// Code generated from MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MagoitoParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MagoitoParser.
type MagoitoParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MagoitoParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MagoitoParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by MagoitoParser#binder.
	VisitBinder(ctx *BinderContext) interface{}

	// Visit a parse tree produced by MagoitoParser#arrowType.
	VisitArrowType(ctx *ArrowTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#tupleArrowType.
	VisitTupleArrowType(ctx *TupleArrowTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#nonTupleType.
	VisitNonTupleType(ctx *NonTupleTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#basicType.
	VisitBasicType(ctx *BasicTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordTypeField.
	VisitRecordTypeField(ctx *RecordTypeFieldContext) interface{}

	// Visit a parse tree produced by MagoitoParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#seqExpr.
	VisitSeqExpr(ctx *SeqExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#controlExpr.
	VisitControlExpr(ctx *ControlExprContext) interface{}

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

	// Visit a parse tree produced by MagoitoParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#callee.
	VisitCallee(ctx *CalleeContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordExpr.
	VisitRecordExpr(ctx *RecordExprContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordFieldList.
	VisitRecordFieldList(ctx *RecordFieldListContext) interface{}

	// Visit a parse tree produced by MagoitoParser#recordField.
	VisitRecordField(ctx *RecordFieldContext) interface{}
}
