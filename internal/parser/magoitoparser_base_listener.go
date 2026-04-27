// Code generated from MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MagoitoParser
import "github.com/antlr4-go/antlr/v4"

// BaseMagoitoParserListener is a complete listener for a parse tree produced by MagoitoParser.
type BaseMagoitoParserListener struct{}

var _ MagoitoParserListener = &BaseMagoitoParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMagoitoParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMagoitoParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMagoitoParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMagoitoParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMagoitoParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMagoitoParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseMagoitoParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseMagoitoParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterBinder is called when production binder is entered.
func (s *BaseMagoitoParserListener) EnterBinder(ctx *BinderContext) {}

// ExitBinder is called when production binder is exited.
func (s *BaseMagoitoParserListener) ExitBinder(ctx *BinderContext) {}

// EnterArrowType is called when production arrowType is entered.
func (s *BaseMagoitoParserListener) EnterArrowType(ctx *ArrowTypeContext) {}

// ExitArrowType is called when production arrowType is exited.
func (s *BaseMagoitoParserListener) ExitArrowType(ctx *ArrowTypeContext) {}

// EnterTupleArrowType is called when production tupleArrowType is entered.
func (s *BaseMagoitoParserListener) EnterTupleArrowType(ctx *TupleArrowTypeContext) {}

// ExitTupleArrowType is called when production tupleArrowType is exited.
func (s *BaseMagoitoParserListener) ExitTupleArrowType(ctx *TupleArrowTypeContext) {}

// EnterNonTupleType is called when production nonTupleType is entered.
func (s *BaseMagoitoParserListener) EnterNonTupleType(ctx *NonTupleTypeContext) {}

// ExitNonTupleType is called when production nonTupleType is exited.
func (s *BaseMagoitoParserListener) ExitNonTupleType(ctx *NonTupleTypeContext) {}

// EnterTupleType is called when production tupleType is entered.
func (s *BaseMagoitoParserListener) EnterTupleType(ctx *TupleTypeContext) {}

// ExitTupleType is called when production tupleType is exited.
func (s *BaseMagoitoParserListener) ExitTupleType(ctx *TupleTypeContext) {}

// EnterBasicType is called when production basicType is entered.
func (s *BaseMagoitoParserListener) EnterBasicType(ctx *BasicTypeContext) {}

// ExitBasicType is called when production basicType is exited.
func (s *BaseMagoitoParserListener) ExitBasicType(ctx *BasicTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BaseMagoitoParserListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BaseMagoitoParserListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterRecordTypeField is called when production recordTypeField is entered.
func (s *BaseMagoitoParserListener) EnterRecordTypeField(ctx *RecordTypeFieldContext) {}

// ExitRecordTypeField is called when production recordTypeField is exited.
func (s *BaseMagoitoParserListener) ExitRecordTypeField(ctx *RecordTypeFieldContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMagoitoParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMagoitoParserListener) ExitExpr(ctx *ExprContext) {}

// EnterSeqExpr is called when production seqExpr is entered.
func (s *BaseMagoitoParserListener) EnterSeqExpr(ctx *SeqExprContext) {}

// ExitSeqExpr is called when production seqExpr is exited.
func (s *BaseMagoitoParserListener) ExitSeqExpr(ctx *SeqExprContext) {}

// EnterControlExpr is called when production controlExpr is entered.
func (s *BaseMagoitoParserListener) EnterControlExpr(ctx *ControlExprContext) {}

// ExitControlExpr is called when production controlExpr is exited.
func (s *BaseMagoitoParserListener) ExitControlExpr(ctx *ControlExprContext) {}

// EnterVarDeclExpr is called when production varDeclExpr is entered.
func (s *BaseMagoitoParserListener) EnterVarDeclExpr(ctx *VarDeclExprContext) {}

// ExitVarDeclExpr is called when production varDeclExpr is exited.
func (s *BaseMagoitoParserListener) ExitVarDeclExpr(ctx *VarDeclExprContext) {}

// EnterWhileExpr is called when production whileExpr is entered.
func (s *BaseMagoitoParserListener) EnterWhileExpr(ctx *WhileExprContext) {}

// ExitWhileExpr is called when production whileExpr is exited.
func (s *BaseMagoitoParserListener) ExitWhileExpr(ctx *WhileExprContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BaseMagoitoParserListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BaseMagoitoParserListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterAssignExpr is called when production assignExpr is entered.
func (s *BaseMagoitoParserListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production assignExpr is exited.
func (s *BaseMagoitoParserListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseMagoitoParserListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseMagoitoParserListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseMagoitoParserListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseMagoitoParserListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseMagoitoParserListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseMagoitoParserListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterComparisonExpr is called when production comparisonExpr is entered.
func (s *BaseMagoitoParserListener) EnterComparisonExpr(ctx *ComparisonExprContext) {}

// ExitComparisonExpr is called when production comparisonExpr is exited.
func (s *BaseMagoitoParserListener) ExitComparisonExpr(ctx *ComparisonExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseMagoitoParserListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseMagoitoParserListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterMultiplicativeExpr is called when production multiplicativeExpr is entered.
func (s *BaseMagoitoParserListener) EnterMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// ExitMultiplicativeExpr is called when production multiplicativeExpr is exited.
func (s *BaseMagoitoParserListener) ExitMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// EnterPowerExpr is called when production powerExpr is entered.
func (s *BaseMagoitoParserListener) EnterPowerExpr(ctx *PowerExprContext) {}

// ExitPowerExpr is called when production powerExpr is exited.
func (s *BaseMagoitoParserListener) ExitPowerExpr(ctx *PowerExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseMagoitoParserListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseMagoitoParserListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterProjectionExpr is called when production projectionExpr is entered.
func (s *BaseMagoitoParserListener) EnterProjectionExpr(ctx *ProjectionExprContext) {}

// ExitProjectionExpr is called when production projectionExpr is exited.
func (s *BaseMagoitoParserListener) ExitProjectionExpr(ctx *ProjectionExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseMagoitoParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseMagoitoParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseMagoitoParserListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseMagoitoParserListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterCallee is called when production callee is entered.
func (s *BaseMagoitoParserListener) EnterCallee(ctx *CalleeContext) {}

// ExitCallee is called when production callee is exited.
func (s *BaseMagoitoParserListener) ExitCallee(ctx *CalleeContext) {}

// EnterRecordExpr is called when production recordExpr is entered.
func (s *BaseMagoitoParserListener) EnterRecordExpr(ctx *RecordExprContext) {}

// ExitRecordExpr is called when production recordExpr is exited.
func (s *BaseMagoitoParserListener) ExitRecordExpr(ctx *RecordExprContext) {}

// EnterRecordFieldList is called when production recordFieldList is entered.
func (s *BaseMagoitoParserListener) EnterRecordFieldList(ctx *RecordFieldListContext) {}

// ExitRecordFieldList is called when production recordFieldList is exited.
func (s *BaseMagoitoParserListener) ExitRecordFieldList(ctx *RecordFieldListContext) {}

// EnterRecordField is called when production recordField is entered.
func (s *BaseMagoitoParserListener) EnterRecordField(ctx *RecordFieldContext) {}

// ExitRecordField is called when production recordField is exited.
func (s *BaseMagoitoParserListener) ExitRecordField(ctx *RecordFieldContext) {}
