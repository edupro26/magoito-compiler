package parser

import (
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/lexer"
)

var validBasicTypes = map[string]bool{
	"Int":    true,
	"Bool":   true,
	"String": true,
	"Unit":   true,
}

func (p *parser) parseType() ast.Type {
	if p.at().Kind == lexer.LPAREN {
		return p.parseGroupedType()
	}
	left := p.parseTypeAtom()
	if p.at().Kind == lexer.ARROW {
		p.advance()
		rt := p.parseType()
		return &ast.FunctionType{Params: []ast.Type{left}, Return: rt}
	}
	return left
}

func (p *parser) parseTypeAtom() ast.Type {
	tok := p.at()
	switch tok.Kind {
	case lexer.IDENTIFIER:
		name := p.advance().Value
		if !validBasicTypes[name] {
			p.errorf("unknown type '%v', expected one of Int, Bool, String, Unit", name)
		}
		return &ast.BasicType{Name: name}
	case lexer.LBRACE:
		return p.parseRecordType()
	default:
		p.errorf("expected 'type', got %v", tok.Kind.ToString())
		return nil
	}
}

func (p *parser) parseGroupedType() ast.Type {
	p.expect(lexer.LPAREN)
	first := p.parseType()
	types := []ast.Type{first}
	for p.at().Kind == lexer.COMMA {
		p.advance()
		types = append(types, p.parseType())
	}
	p.expect(lexer.RPAREN)
	if p.at().Kind == lexer.ARROW {
		p.advance()
		ret := p.parseType()
		return &ast.FunctionType{Params: types, Return: ret}
	}
	if len(types) > 1 {
		p.errorf("multi-type group must be followed by '->'")
	}
	return types[0]
}

func (p *parser) parseRecordType() *ast.RecordType {
	p.expect(lexer.LBRACE)
	var fields []ast.RecordTypeField
	if p.at().Kind != lexer.RBRACE {
		label := p.expect(lexer.IDENTIFIER).Value
		p.expect(lexer.COLON)
		typ := p.parseType()
		fields = append(fields, ast.RecordTypeField{Label: label, Type: typ})
		for p.at().Kind == lexer.COMMA {
			p.advance()
			label = p.expect(lexer.IDENTIFIER).Value
			p.expect(lexer.COLON)
			typ = p.parseType()
			fields = append(fields, ast.RecordTypeField{Label: label, Type: typ})
		}
	}
	p.expect(lexer.RBRACE)
	return &ast.RecordType{Fields: fields}
}
