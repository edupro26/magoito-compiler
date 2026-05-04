package semantics

import "magoito-compiler/internal/ast"

type Kind int

const (
	CONST Kind = iota
	FUN
	VAR
)

type Symbol struct {
	Name string
	Type ast.Type
	Kind Kind
}

type SymbolTable struct {
	parent  *SymbolTable
	entries map[string]Symbol
}

func newSymbolTable(parent *SymbolTable) *SymbolTable {
	return &SymbolTable{
		parent:  parent,
		entries: make(map[string]Symbol),
	}
}

func (t *SymbolTable) define(name string, typ ast.Type, kind Kind) bool {
	if _, exists := t.entries[name]; exists {
		return false
	}
	t.entries[name] = Symbol{Name: name, Type: typ, Kind: kind}
	return true
}

func (t *SymbolTable) extend(name string, typ ast.Type, kind Kind) *SymbolTable {
	child := newSymbolTable(t)
	child.entries[name] = Symbol{Name: name, Type: typ, Kind: kind}
	return child
}

func (t *SymbolTable) lookup(name string) (Symbol, bool) {
	for table := t; table != nil; table = table.parent {
		if sym, ok := table.entries[name]; ok {
			return sym, true
		}
	}
	return Symbol{}, false
}
