package main

import (
	"fmt"
	"magoito-compiler/cmd/printer"
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/parser"
	"os"
)

func main() {
	// 1. Read file
	bytes, err := os.ReadFile("example01.mag")
	if err != nil {
		panic(err)
	}

	cst, err := parser.ParseProgram(string(bytes))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// --- Build AST and print ---
	tree := ast.Build(cst)
	printer.Print(tree)
}
