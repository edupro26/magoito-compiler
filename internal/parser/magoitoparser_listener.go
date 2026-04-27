// Code generated from MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MagoitoParser
import "github.com/antlr4-go/antlr/v4"

// MagoitoParserListener is a complete listener for a parse tree produced by MagoitoParser.
type MagoitoParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterBinder is called when entering the binder production.
	EnterBinder(c *BinderContext)

	// EnterArrowType is called when entering the arrowType production.
	EnterArrowType(c *ArrowTypeContext)

	// EnterTupleArrowType is called when entering the tupleArrowType production.
	EnterTupleArrowType(c *TupleArrowTypeContext)

	// EnterNonTupleType is called when entering the nonTupleType production.
	EnterNonTupleType(c *NonTupleTypeContext)

	// EnterTupleType is called when entering the tupleType production.
	EnterTupleType(c *TupleTypeContext)

	// EnterBasicType is called when entering the basicType production.
	EnterBasicType(c *BasicTypeContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterRecordTypeField is called when entering the recordTypeField production.
	EnterRecordTypeField(c *RecordTypeFieldContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterSeqExpr is called when entering the seqExpr production.
	EnterSeqExpr(c *SeqExprContext)

	// EnterControlExpr is called when entering the controlExpr production.
	EnterControlExpr(c *ControlExprContext)

	// EnterVarDeclExpr is called when entering the varDeclExpr production.
	EnterVarDeclExpr(c *VarDeclExprContext)

	// EnterWhileExpr is called when entering the whileExpr production.
	EnterWhileExpr(c *WhileExprContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterAssignExpr is called when entering the assignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterComparisonExpr is called when entering the comparisonExpr production.
	EnterComparisonExpr(c *ComparisonExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterMultiplicativeExpr is called when entering the multiplicativeExpr production.
	EnterMultiplicativeExpr(c *MultiplicativeExprContext)

	// EnterPowerExpr is called when entering the powerExpr production.
	EnterPowerExpr(c *PowerExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterProjectionExpr is called when entering the projectionExpr production.
	EnterProjectionExpr(c *ProjectionExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterCallee is called when entering the callee production.
	EnterCallee(c *CalleeContext)

	// EnterRecordExpr is called when entering the recordExpr production.
	EnterRecordExpr(c *RecordExprContext)

	// EnterRecordFieldList is called when entering the recordFieldList production.
	EnterRecordFieldList(c *RecordFieldListContext)

	// EnterRecordField is called when entering the recordField production.
	EnterRecordField(c *RecordFieldContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitBinder is called when exiting the binder production.
	ExitBinder(c *BinderContext)

	// ExitArrowType is called when exiting the arrowType production.
	ExitArrowType(c *ArrowTypeContext)

	// ExitTupleArrowType is called when exiting the tupleArrowType production.
	ExitTupleArrowType(c *TupleArrowTypeContext)

	// ExitNonTupleType is called when exiting the nonTupleType production.
	ExitNonTupleType(c *NonTupleTypeContext)

	// ExitTupleType is called when exiting the tupleType production.
	ExitTupleType(c *TupleTypeContext)

	// ExitBasicType is called when exiting the basicType production.
	ExitBasicType(c *BasicTypeContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitRecordTypeField is called when exiting the recordTypeField production.
	ExitRecordTypeField(c *RecordTypeFieldContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitSeqExpr is called when exiting the seqExpr production.
	ExitSeqExpr(c *SeqExprContext)

	// ExitControlExpr is called when exiting the controlExpr production.
	ExitControlExpr(c *ControlExprContext)

	// ExitVarDeclExpr is called when exiting the varDeclExpr production.
	ExitVarDeclExpr(c *VarDeclExprContext)

	// ExitWhileExpr is called when exiting the whileExpr production.
	ExitWhileExpr(c *WhileExprContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitAssignExpr is called when exiting the assignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitComparisonExpr is called when exiting the comparisonExpr production.
	ExitComparisonExpr(c *ComparisonExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitMultiplicativeExpr is called when exiting the multiplicativeExpr production.
	ExitMultiplicativeExpr(c *MultiplicativeExprContext)

	// ExitPowerExpr is called when exiting the powerExpr production.
	ExitPowerExpr(c *PowerExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitProjectionExpr is called when exiting the projectionExpr production.
	ExitProjectionExpr(c *ProjectionExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitCallee is called when exiting the callee production.
	ExitCallee(c *CalleeContext)

	// ExitRecordExpr is called when exiting the recordExpr production.
	ExitRecordExpr(c *RecordExprContext)

	// ExitRecordFieldList is called when exiting the recordFieldList production.
	ExitRecordFieldList(c *RecordFieldListContext)

	// ExitRecordField is called when exiting the recordField production.
	ExitRecordField(c *RecordFieldContext)
}
