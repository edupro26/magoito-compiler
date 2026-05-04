package semantics

import (
	"fmt"
	"magoito-compiler/internal/ast"
	"strings"
)

type SemanticError struct {
	position *ast.Position
	message  string
}

func (e *SemanticError) Error() string {
	if e.position != nil {
		return fmt.Sprintf("%d:%d: semantic error: %s",
			e.position.Line, e.position.Col, e.message)
	} else {
		return fmt.Sprintf("semantic error: %s", e.message)
	}
}

type ErrorList struct {
	Errors []SemanticError
}

func (l *ErrorList) Error() string {
	if len(l.Errors) == 0 {
		return ""
	}
	lines := make([]string, len(l.Errors))
	for i, err := range l.Errors {
		lines[i] = err.Error()
	}
	return strings.Join(lines, "\n")
}

func (l *ErrorList) Add(pos *ast.Position, msg string) {
	l.Errors = append(l.Errors, SemanticError{
		position: pos,
		message:  msg,
	})
}

func (l *ErrorList) HasErrors() bool {
	return len(l.Errors) > 0
}

const (
	DUP_TOP_DECL = "duplicate top-level declaration"
	MISSING_MAIN = "missing 'main' function"

	BAD_CONST_INIT   = "unsupported const initializer"
	TYPE_MISMATCH    = "type mismatch"
	DUP_RECORD_FIELD = "duplicate record field"
)
