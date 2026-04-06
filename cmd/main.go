package main

import (
	"bufio"
	"fmt"
	"magoito-compiler/cmd/printer"
	"magoito-compiler/internal/lexer"
	"magoito-compiler/internal/parser"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Magoito CLI")
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			fmt.Println()
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "q" || line == "quit" {
			return
		}

		parts := strings.Fields(line)
		if len(parts) == 2 && parts[0] == "run" {
			if err := runFile(parts[1]); err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("unknown command. Use: run <file>, q, or quit")
	}
}

func runFile(filePath string) error {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	tokens, err := lexer.Tokenize(string(bytes))
	if err != nil {
		return err
	}
	ast, err := parser.Parse(tokens)
	if err != nil {
		return fmt.Errorf("syntax error: %w", err)
	}

	fmt.Printf("parsed %d top-level declaration(s)\n\n", len(ast.Declarations))
	printer.Print(ast)
	return nil
}
