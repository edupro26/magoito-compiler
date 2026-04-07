package parser

import (
	"errors"
	"fmt"
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

type ParseError struct {
	Message string
	Line    int
	Col     int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error: %s", e.Line, e.Col, e.Message)
}

func Parse(tokens []lexer.Token) (program *ast.Program, err error) {
	defer func() {
		if r := recover(); r != nil {
			if errVal, ok := r.(error); ok {
				var pe *ParseError
				if errors.As(errVal, &pe) {
					err = pe
					return
				}
			}
			panic(r)
		}
	}()
	p := createParser(tokens)
	program = p.parseProgram()
	return program, nil
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{tokens: tokens}
}

// at returns the current token without consuming it.
func (p *parser) at() lexer.Token {
	return p.tokens[p.pos]
}

// peek returns the token at offset positions ahead of the current position,
// returns EOF when the offset is out of range.
func (p *parser) peek(offset int) lexer.Token {
	idx := p.pos + offset
	if idx >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}
	return p.tokens[idx]
}

// advance consumes and returns the current token.
func (p *parser) advance() lexer.Token {
	token := p.at()
	if token.Kind != lexer.EOF {
		p.pos++
	}
	return token
}

// expect consumes the current token if it matches kind; otherwise it panics.
func (p *parser) expect(kind lexer.Kind) lexer.Token {
	token := p.at()
	if token.Kind != kind {
		panic(&ParseError{
			Message: fmt.Sprintf(
				"expected '%v', got '%v'", kind.ToString(), token.Kind.ToString(),
			),
			Line: token.Line,
			Col:  token.Col,
		})
	}
	return p.advance()
}

// errorf panics with a ParseError formatted with the given message.
func (p *parser) errorf(format string, args ...any) {
	token := p.at()
	panic(&ParseError{
		Message: fmt.Sprintf(format, args...),
		Line:    token.Line,
		Col:     token.Col,
	})
}

func (p *parser) parseProgram() *ast.Program {
	var decls []ast.Declaration
	for p.at().Kind != lexer.EOF {
		decls = append(decls, p.parseDeclaration())
	}
	if len(decls) == 0 {
		p.errorf("program must contain at least one declaration")
	}
	return &ast.Program{Declarations: decls}
}

func (p *parser) parseDeclaration() ast.Declaration {
	switch p.at().Kind {
	case lexer.CONST:
		return p.parseConstDecl()
	case lexer.FUN:
		return p.parseFunDecl()
	default:
		p.errorf("expected 'const' or 'fun', got %s", p.at().Kind.ToString())
		return nil
	}
}

func (p *parser) parseConstDecl() *ast.ConstDecl {
	p.expect(lexer.CONST)
	name := p.expect(lexer.IDENTIFIER).Value
	p.expect(lexer.COLON)
	typ := p.parseType()
	p.expect(lexer.ASSIGN)
	value := p.parseExpr()
	return &ast.ConstDecl{Name: name, Type: typ, Value: value}
}

func (p *parser) parseFunDecl() *ast.FunDecl {
	p.expect(lexer.FUN)
	name := p.expect(lexer.IDENTIFIER).Value
	p.expect(lexer.LPAREN)
	params := p.parseParams()
	p.expect(lexer.RPAREN)
	p.expect(lexer.COLON)
	typ := p.parseType()
	p.expect(lexer.ASSIGN)
	body := p.parseExpr()
	return &ast.FunDecl{Name: name, Params: params, Type: typ, Body: body}
}

func (p *parser) parseParams() []string {
	params := []string{p.parseParam()}
	for p.at().Kind == lexer.COMMA {
		p.advance()
		params = append(params, p.parseParam())
	}
	return params
}

func (p *parser) parseParam() string {
	tok := p.at()
	switch tok.Kind {
	case lexer.IDENTIFIER:
		return p.advance().Value
	case lexer.WILDCARD:
		return p.advance().Value
	default:
		p.errorf("expected 'identifier' or '_' in parameter list, got %v", tok.Kind.ToString())
		return ""
	}
}
