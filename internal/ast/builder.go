package ast

import (
	"sort"
	"strconv"

	"magoito-compiler/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

type ASTBuilder struct {
	*parser.BaseMagoitoParserVisitor
	ast *Program
}

func Build(cst antlr.ParseTree) (program *Program) {
	b := &ASTBuilder{ast: &Program{}}
	return b.Visit(cst).(*Program)
}

func (b *ASTBuilder) Visit(tree antlr.ParseTree) any {
	if tree == nil {
		return nil
	}
	return tree.Accept(b)
}

func (b *ASTBuilder) VisitProgram(ctx *parser.ProgramContext) any {
	decls := make([]Declaration, 0, len(ctx.AllDeclaration()))
	for _, d := range ctx.AllDeclaration() {
		if decl, ok := b.Visit(d).(Declaration); ok {
			decls = append(decls, decl)
		}
	}
	return &Program{Declarations: decls}
}

func (b *ASTBuilder) VisitConstDeclaration(ctx *parser.ConstDeclarationContext) any {
	name := ""
	if binder := ctx.Binder(); binder != nil {
		name = binder.GetText()
	}
	typ := AsType(b.Visit(ctx.TypeExpr()))
	value := AsExpr(b.Visit(ctx.Expr()))
	return &ConstDecl{Name: name, Type: typ, Value: value}
}

func (b *ASTBuilder) VisitFunDeclaration(ctx *parser.FunDeclarationContext) any {
	name := ""
	if id := ctx.IDENTIFIER(); id != nil {
		name = id.GetText()
	}
	params := make([]string, 0)
	for _, p := range ctx.AllBinder() {
		params = append(params, p.GetText())
	}
	typ := AsType(b.Visit(ctx.TypeExpr()))
	body := AsExpr(b.Visit(ctx.Expr()))
	return &FunDecl{Name: name, Params: params, Type: typ, Body: body}
}

func (b *ASTBuilder) VisitArrowType(ctx *parser.ArrowTypeContext) any {
	left := AsType(b.Visit(ctx.NonTupleType()))
	if ctx.TypeExpr() == nil {
		return left
	}
	params := []Type{left}
	ret := AsType(b.Visit(ctx.TypeExpr()))
	return &FunctionType{Params: params, Return: ret}
}

func (b *ASTBuilder) VisitTupleArrowType(ctx *parser.TupleArrowTypeContext) any {
	params := make([]Type, 0)
	for _, t := range ctx.TupleType().AllTypeExpr() {
		params = append(params, AsType(b.Visit(t)))
	}
	ret := AsType(b.Visit(ctx.TypeExpr()))
	return &FunctionType{Params: params, Return: ret}
}

func (b *ASTBuilder) VisitBasicNonTupleType(ctx *parser.BasicNonTupleTypeContext) any {
	return &BasicType{Name: ctx.BasicType().GetText()}
}

func (b *ASTBuilder) VisitRecordNonTupleType(ctx *parser.RecordNonTupleTypeContext) any {
	fields := []RecordTypeField{}
	for _, f := range ctx.RecordType().AllRecordTypeField() {
		fields = append(fields, RecordTypeField{
			Label: f.IDENTIFIER().GetText(),
			Type:  AsType(b.Visit(f.TypeExpr())),
		})
	}
	return &RecordType{Fields: fields}
}

func (b *ASTBuilder) VisitParenNonTupleType(ctx *parser.ParenNonTupleTypeContext) any {
	return b.Visit(ctx.TypeExpr())
}

func (b *ASTBuilder) VisitIntBasicType(ctx *parser.IntBasicTypeContext) any {
	return &BasicType{Name: ctx.GetText()}
}

func (b *ASTBuilder) VisitBoolBasicType(ctx *parser.BoolBasicTypeContext) any {
	return &BasicType{Name: ctx.GetText()}
}

func (b *ASTBuilder) VisitStringBasicType(ctx *parser.StringBasicTypeContext) any {
	return &BasicType{Name: ctx.GetText()}
}

func (b *ASTBuilder) VisitUnitBasicType(ctx *parser.UnitBasicTypeContext) any {
	return &BasicType{Name: ctx.GetText()}
}

func (b *ASTBuilder) VisitExpr(ctx *parser.ExprContext) any {
	return b.Visit(ctx.SeqExpr())
}

func (b *ASTBuilder) VisitSeqExpr(ctx *parser.SeqExprContext) any {
	left := AsExpr(b.Visit(ctx.ControlExpr()))
	if ctx.SeqExpr() != nil {
		right := AsExpr(b.Visit(ctx.SeqExpr()))
		return &SequenceExpr{Left: left, Right: right}
	}
	return left
}

func (b *ASTBuilder) VisitVarDeclControl(ctx *parser.VarDeclControlContext) any {
	return b.Visit(ctx.VarDeclExpr())
}

func (b *ASTBuilder) VisitWhileControl(ctx *parser.WhileControlContext) any {
	return b.Visit(ctx.WhileExpr())
}

func (b *ASTBuilder) VisitIfControl(ctx *parser.IfControlContext) any {
	return b.Visit(ctx.IfExpr())
}

func (b *ASTBuilder) VisitAssignControl(ctx *parser.AssignControlContext) any {
	return b.Visit(ctx.AssignExpr())
}

func (b *ASTBuilder) VisitOrControl(ctx *parser.OrControlContext) any {
	return b.Visit(ctx.OrExpr())
}

func (b *ASTBuilder) VisitVarDeclExpr(ctx *parser.VarDeclExprContext) any {
	name := ctx.Binder().GetText()
	typ := AsType(b.Visit(ctx.TypeExpr()))
	value := AsExpr(b.Visit(ctx.ControlExpr()))
	return &VarDecl{Name: name, Type: typ, Value: value}
}

func (b *ASTBuilder) VisitWhileExpr(ctx *parser.WhileExprContext) any {
	cond := AsExpr(b.Visit(ctx.OrExpr()))
	body := AsExpr(b.Visit(ctx.ControlExpr()))
	return &WhileExpr{Condition: cond, Body: body}
}

func (b *ASTBuilder) VisitIfExpr(ctx *parser.IfExprContext) any {
	cond := AsExpr(b.Visit(ctx.OrExpr()))
	then := AsExpr(b.Visit(ctx.ControlExpr(0)))
	var els Expr
	if len(ctx.AllControlExpr()) > 1 {
		els = AsExpr(b.Visit(ctx.ControlExpr(1)))
	}
	return &IfExpr{Condition: cond, Then: then, Else: els}
}

func (b *ASTBuilder) VisitAssignExpr(ctx *parser.AssignExprContext) any {
	name := ctx.IDENTIFIER().GetText()
	value := AsExpr(b.Visit(ctx.ControlExpr()))
	return &Assignment{Name: name, Value: value}
}

func (b *ASTBuilder) VisitOrExpr(ctx *parser.OrExprContext) any {
	parts := ctx.AllAndExpr()
	if len(parts) == 0 {
		return nil
	}
	ops := ctx.AllOR()
	left := AsExpr(b.Visit(parts[0]))
	for i := 1; i < len(parts); i++ {
		right := AsExpr(b.Visit(parts[i]))
		op := ops[i-1].GetText()
		left = &BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (b *ASTBuilder) VisitAndExpr(ctx *parser.AndExprContext) any {
	parts := ctx.AllEqualityExpr()
	if len(parts) == 0 {
		return nil
	}
	ops := ctx.AllAND()
	left := AsExpr(b.Visit(parts[0]))
	for i := 1; i < len(parts); i++ {
		right := AsExpr(b.Visit(parts[i]))
		op := ops[i-1].GetText()
		left = &BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (b *ASTBuilder) VisitEqualityExpr(ctx *parser.EqualityExprContext) any {
	parts := ctx.AllComparisonExpr()
	if len(parts) == 0 {
		return nil
	}
	ops := opsFromTokens(ctx.AllEQ(), ctx.AllNEQ())
	left := AsExpr(b.Visit(parts[0]))
	for i := 1; i < len(parts); i++ {
		right := AsExpr(b.Visit(parts[i]))
		op := ops[i-1]
		left = &BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (b *ASTBuilder) VisitComparisonExpr(ctx *parser.ComparisonExprContext) any {
	parts := ctx.AllAdditiveExpr()
	if len(parts) == 0 {
		return nil
	}
	ops := opsFromTokens(ctx.AllLT(), ctx.AllLTE(), ctx.AllGT(), ctx.AllGTE())
	left := AsExpr(b.Visit(parts[0]))
	for i := 1; i < len(parts); i++ {
		right := AsExpr(b.Visit(parts[i]))
		op := ops[i-1]
		left = &BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (b *ASTBuilder) VisitAdditiveExpr(ctx *parser.AdditiveExprContext) any {
	parts := ctx.AllMultiplicativeExpr()
	if len(parts) == 0 {
		return nil
	}
	ops := opsFromTokens(ctx.AllPLUS(), ctx.AllMINUS())
	left := AsExpr(b.Visit(parts[0]))
	for i := 1; i < len(parts); i++ {
		right := AsExpr(b.Visit(parts[i]))
		op := ops[i-1]
		left = &BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (b *ASTBuilder) VisitMultiplicativeExpr(ctx *parser.MultiplicativeExprContext) any {
	parts := ctx.AllPowerExpr()
	if len(parts) == 0 {
		return nil
	}
	ops := opsFromTokens(ctx.AllSTAR(), ctx.AllSLASH(), ctx.AllPERCENT())
	left := AsExpr(b.Visit(parts[0]))
	for i := 1; i < len(parts); i++ {
		right := AsExpr(b.Visit(parts[i]))
		op := ops[i-1]
		left = &BinaryExpr{Left: left, Op: op, Right: right}
	}
	return left
}

func (b *ASTBuilder) VisitPowerExpr(ctx *parser.PowerExprContext) any {
	if ctx.PowerExpr() != nil {
		left := AsExpr(b.Visit(ctx.UnaryExpr()))
		right := AsExpr(b.Visit(ctx.PowerExpr()))
		op := ctx.POWER().GetText()
		return &BinaryExpr{Left: left, Op: op, Right: right}
	}
	return b.Visit(ctx.UnaryExpr())
}

func (b *ASTBuilder) VisitUnaryExpr(ctx *parser.UnaryExprContext) any {
	if ctx.MINUS() != nil {
		op := ctx.MINUS().GetText()
		operand := AsExpr(b.Visit(ctx.UnaryExpr()))
		return &UnaryExpr{Op: op, Operand: operand}
	}
	if ctx.NOT() != nil {
		op := ctx.NOT().GetText()
		operand := AsExpr(b.Visit(ctx.UnaryExpr()))
		return &UnaryExpr{Op: op, Operand: operand}
	}
	return b.Visit(ctx.ProjectionExpr())
}

func (b *ASTBuilder) VisitProjectionExpr(ctx *parser.ProjectionExprContext) any {
	base := AsExpr(b.Visit(ctx.PrimaryExpr()))
	for _, ident := range ctx.AllIDENTIFIER() {
		base = &ProjectionExpr{Record: base, Field: ident.GetText()}
	}
	return base
}

func (b *ASTBuilder) VisitIntLiteralPrimary(ctx *parser.IntLiteralPrimaryContext) any {
	value, _ := strconv.Atoi(ctx.INT_LITERAL().GetText())
	return &IntLiteral{Value: value}
}

func (b *ASTBuilder) VisitStringLiteralPrimary(ctx *parser.StringLiteralPrimaryContext) any {
	s, _ := strconv.Unquote(ctx.STRING_LITERAL().GetText())
	return &StringLiteral{Value: s}
}

func (b *ASTBuilder) VisitTruePrimary(ctx *parser.TruePrimaryContext) any {
	return &BoolLiteral{Value: true}
}

func (b *ASTBuilder) VisitFalsePrimary(ctx *parser.FalsePrimaryContext) any {
	return &BoolLiteral{Value: false}
}

func (b *ASTBuilder) VisitUnitPrimary(ctx *parser.UnitPrimaryContext) any {
	return &UnitLiteral{}
}

func (b *ASTBuilder) VisitCallPrimary(ctx *parser.CallPrimaryContext) any {
	return b.Visit(ctx.CallExpr())
}

func (b *ASTBuilder) VisitIdentifierPrimary(ctx *parser.IdentifierPrimaryContext) any {
	return &Identifier{Name: ctx.IDENTIFIER().GetText()}
}

func (b *ASTBuilder) VisitParenPrimary(ctx *parser.ParenPrimaryContext) any {
	return b.Visit(ctx.Expr())
}

func (b *ASTBuilder) VisitRecordPrimary(ctx *parser.RecordPrimaryContext) any {
	return b.Visit(ctx.RecordExpr())
}

func (b *ASTBuilder) VisitCallExpr(ctx *parser.CallExprContext) any {
	callee := ctx.Callee().GetText()
	args := make([]Expr, 0)
	for _, a := range ctx.AllExpr() {
		args = append(args, AsExpr(b.Visit(a)))
	}
	return &CallExpr{Callee: callee, Args: args}
}

func (b *ASTBuilder) VisitRecordExpr(ctx *parser.RecordExprContext) any {
	fields := []RecordField{}
	if list := ctx.RecordFieldList(); list != nil {
		for _, f := range list.AllRecordField() {
			label := f.IDENTIFIER().GetText()
			value := AsExpr(b.Visit(f.Expr()))
			fields = append(fields, RecordField{Label: label, Value: value})
		}
	}
	return &RecordExpr{Fields: fields}
}

func opsFromTokens(groups ...[]antlr.TerminalNode) []string {
	var nodes []antlr.TerminalNode
	for _, g := range groups {
		nodes = append(nodes, g...)
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].GetSymbol().GetTokenIndex() <
			nodes[j].GetSymbol().GetTokenIndex()
	})
	ops := make([]string, len(nodes))
	for i, n := range nodes {
		ops[i] = n.GetText()
	}
	return ops
}
