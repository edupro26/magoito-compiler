package parser

import (
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/lexer"
	"strconv"
)

// Operator precedence (lowest → highest):
//	;           |  sequence        |  parseSequence
//	while / if  |  control flow    |  parseControl
//	||          |  disjunction     |  parseOr
//	&&          |  conjunction     |  parseAnd
//	== !=       |  equality        |  parseEquality
//	< <= > >=   |  comparison      |  parseComparison
//	+ -         |  addition        |  parseAddition
//	* / %       |  multiplication  |  parseMultiplication
//	^           |  power           |  parsePower
//	- !         |  unary           |  parseUnary
//	.           |  projection      |  parseProjection
//	            |  primary         |  parsePrimary

func (p *parser) parseExpr() ast.Expr {
	return p.parseSequence()
}

func (p *parser) parseSequence() ast.Expr {
	left := p.parseControl()
	if p.at().Kind == lexer.SEMICOLON {
		p.advance()
		right := p.parseSequence()
		return &ast.SequenceExpr{Left: left, Right: right}
	}
	return left
}

func (p *parser) parseControl() ast.Expr {
	switch p.at().Kind {
	case lexer.VAR:
		return p.parseVarDecl()
	case lexer.WHILE:
		return p.parseWhile()
	case lexer.IF:
		return p.parseIf()
	case lexer.IDENTIFIER:
		if p.peek(1).Kind == lexer.ASSIGN {
			return p.parseAssignment()
		}
		fallthrough
	default:
		return p.parseOr()
	}
}

func (p *parser) parseVarDecl() *ast.VarDecl {
	p.expect(lexer.VAR)
	var name string
	switch p.at().Kind {
	case lexer.IDENTIFIER:
		name = p.advance().Value
	case lexer.WILDCARD:
		name = p.advance().Value
	default:
		p.errorf("expected identifier or _ after 'var', got %v", p.at().Kind.ToString())
	}
	p.expect(lexer.COLON)
	typ := p.parseType()
	p.expect(lexer.ASSIGN)
	value := p.parseControl()
	return &ast.VarDecl{Name: name, Type: typ, Value: value}
}

func (p *parser) parseWhile() *ast.WhileExpr {
	p.expect(lexer.WHILE)
	condition := p.parseOr()
	p.expect(lexer.DO)
	body := p.parseControl()
	return &ast.WhileExpr{Condition: condition, Body: body}
}

func (p *parser) parseIf() *ast.IfExpr {
	p.expect(lexer.IF)
	condition := p.parseOr()
	p.expect(lexer.THEN)
	thenBranch := p.parseControl()
	var elseBranch ast.Expr
	if p.at().Kind == lexer.ELSE {
		p.advance()
		elseBranch = p.parseControl()
	}
	return &ast.IfExpr{Condition: condition, Then: thenBranch, Else: elseBranch}
}

func (p *parser) parseAssignment() *ast.Assignment {
	name := p.expect(lexer.IDENTIFIER).Value
	p.expect(lexer.ASSIGN)
	value := p.parseControl()
	return &ast.Assignment{Name: name, Value: value}
}

func (p *parser) parseOr() ast.Expr {
	left := p.parseAnd()
	for p.at().Kind == lexer.OR {
		p.advance()
		right := p.parseAnd()
		left = &ast.BinaryExpr{Op: "||", Left: left, Right: right}
	}
	return left
}

func (p *parser) parseAnd() ast.Expr {
	left := p.parseEquality()
	for p.at().Kind == lexer.AND {
		p.advance()
		right := p.parseEquality()
		left = &ast.BinaryExpr{Op: "&&", Left: left, Right: right}
	}
	return left
}

func (p *parser) parseEquality() ast.Expr {
	left := p.parseComparison()
	for p.at().Kind == lexer.EQ || p.at().Kind == lexer.NEQ {
		op := p.advance().Value
		right := p.parseComparison()
		left = &ast.BinaryExpr{Op: op, Left: left, Right: right}
	}
	return left
}

func (p *parser) parseComparison() ast.Expr {
	left := p.parseAddition()
	for {
		switch p.at().Kind {
		case lexer.LT, lexer.LTE, lexer.GT, lexer.GTE:
			op := p.advance().Value
			right := p.parseAddition()
			left = &ast.BinaryExpr{Op: op, Left: left, Right: right}
		default:
			return left
		}
	}
}

func (p *parser) parseAddition() ast.Expr {
	left := p.parseMultiplication()
	for p.at().Kind == lexer.PLUS || p.at().Kind == lexer.MINUS {
		op := p.advance().Value
		right := p.parseMultiplication()
		left = &ast.BinaryExpr{Op: op, Left: left, Right: right}
	}
	return left
}

func (p *parser) parseMultiplication() ast.Expr {
	left := p.parsePower()
	for p.at().Kind == lexer.STAR || p.at().Kind == lexer.SLASH || p.at().Kind == lexer.PERCENT {
		op := p.advance().Value
		right := p.parsePower()
		left = &ast.BinaryExpr{Op: op, Left: left, Right: right}
	}
	return left
}

func (p *parser) parsePower() ast.Expr {
	base := p.parseUnary()
	if p.at().Kind == lexer.POWER {
		p.advance()
		exp := p.parsePower()
		return &ast.BinaryExpr{Op: "^", Left: base, Right: exp}
	}
	return base
}

func (p *parser) parseUnary() ast.Expr {
	switch p.at().Kind {
	case lexer.MINUS:
		p.advance()
		return &ast.UnaryExpr{Op: "-", Operand: p.parseUnary()}
	case lexer.NOT:
		p.advance()
		return &ast.UnaryExpr{Op: "!", Operand: p.parseUnary()}
	default:
		return p.parseProjection()
	}
}

func (p *parser) parseProjection() ast.Expr {
	expr := p.parsePrimary()
	for p.at().Kind == lexer.DOT {
		p.advance()
		field := p.expect(lexer.IDENTIFIER).Value
		expr = &ast.ProjectionExpr{Record: expr, Field: field}
	}
	return expr
}

func (p *parser) parsePrimary() ast.Expr {
	switch p.at().Kind {
	case lexer.INT:
		tok := p.advance()
		val, err := strconv.Atoi(tok.Value)
		if err != nil {
			p.errorf("invalid integer %q", tok.Value)
		}
		return &ast.IntLiteral{Value: val}
	case lexer.STRING:
		tok := p.advance()
		return &ast.StringLiteral{Value: tok.Value}
	case lexer.TRUE:
		p.advance()
		return &ast.BoolLiteral{Value: true}
	case lexer.FALSE:
		p.advance()
		return &ast.BoolLiteral{Value: false}
	case lexer.UNIT:
		p.advance()
		return &ast.UnitLiteral{}
	case lexer.PRINT:
		p.advance()
		p.expect(lexer.LPAREN)
		arg := p.parseExpr()
		p.expect(lexer.RPAREN)
		return &ast.CallExpr{Callee: "print", Args: []ast.Expr{arg}}
	case lexer.IDENTIFIER:
		name := p.advance().Value
		if p.at().Kind == lexer.LPAREN {
			p.advance()
			args := p.parseCallArgs()
			p.expect(lexer.RPAREN)
			return &ast.CallExpr{Callee: name, Args: args}
		}
		return &ast.Identifier{Name: name}
	case lexer.LPAREN:
		p.advance()
		expr := p.parseExpr()
		p.expect(lexer.RPAREN)
		return expr
	case lexer.LBRACE:
		return p.parseRecordExpr()
	default:
		p.errorf("unexpected token %v in expression", p.at().Kind.ToString())
	}
	return nil
}

func (p *parser) parseCallArgs() []ast.Expr {
	if p.at().Kind == lexer.RPAREN {
		return nil
	}
	args := []ast.Expr{p.parseExpr()}
	for p.at().Kind == lexer.COMMA {
		p.advance()
		args = append(args, p.parseExpr())
	}
	return args
}

func (p *parser) parseRecordExpr() *ast.RecordExpr {
	p.expect(lexer.LBRACE)
	var fields []ast.RecordField
	if p.at().Kind != lexer.RBRACE {
		label := p.expect(lexer.IDENTIFIER).Value
		p.expect(lexer.ASSIGN)
		value := p.parseExpr()
		fields = append(fields, ast.RecordField{Label: label, Value: value})
		for p.at().Kind == lexer.COMMA {
			p.advance()
			label = p.expect(lexer.IDENTIFIER).Value
			p.expect(lexer.ASSIGN)
			value = p.parseExpr()
			fields = append(fields, ast.RecordField{Label: label, Value: value})
		}
	}
	p.expect(lexer.RBRACE)
	return &ast.RecordExpr{Fields: fields}
}
