package util

import (
	"errors"
	"fmt"
	"magoito-compiler/cmd/printer"
	"magoito-compiler/internal/lexer"
	"magoito-compiler/internal/parser"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

const (
	OK   = "✅"
	FAIL = "❌"
)

func PrintBanner() {
	fmt.Printf("Magoito CLI (early version), Go version: %s\n", runtime.Version())
	fmt.Println(`Try "help" for more information, or run an example: run --example 1`)
}

func PrintHelp() {
	fmt.Println("Commands:")
	fmt.Println(" mag run [option] <file.mag>         Parse a MAGOITO source file")
	fmt.Println(" mag run [option] --example <n>      Run a MAGOITO example")
	fmt.Println(" mag test                            Run MAGOITO test suite")
	fmt.Println(" help                                Show this help message")
	fmt.Println(" quit                                Exit the CLI")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println(" --ast                               Replace default output with the AST tree")
	fmt.Println()
}

func RunCmd(args []string) error {
	if len(args) == 0 {
		return errors.New(`usage: mag run <file.mag> [--ast]`)
	}
	printAST := hasFlag(args, "--ast")
	args = filterOut(args, "--ast")

	if len(args) >= 2 && (args[0] == "--example" || args[0] == "-e") {
		n, err := strconv.Atoi(args[1])
		if err != nil || n < 1 || n > 3 {
			return errors.New("example must be 1, 2, or 3. Try: run -e 1")
		}
		file := fmt.Sprintf("example%02d.mag", n)
		return runFile(file, printAST)
	}

	if len(args) != 1 {
		return errors.New(`usage: mag run <file.mag> [--ast]`)
	}
	return runFile(args[0], printAST)
}

func TestCmd() error {
	testDirs := []string{
		filepath.Join("test", "valid"),
		filepath.Join("test", "invalid-syntax"),
		filepath.Join("test", "invalid-semantics"),
	}
	for _, dir := range testDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("cannot read test directory %q: %w", dir, err)
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			subDir := filepath.Join(dir, entry.Name())
			files, err := filepath.Glob(filepath.Join(subDir, "*.mag"))
			if err != nil {
				return err
			}
			for _, f := range files {
				emoji := OK
				content, err := os.ReadFile(f)
				if err != nil {
					emoji = FAIL
				} else {
					tokens, err := lexer.Tokenize(string(content))
					if err != nil {
						emoji = FAIL
					} else if _, err = parser.Parse(tokens); err != nil {
						emoji = FAIL
					}
				}
				fmt.Printf("%s - %s\n", f, emoji)
			}
		}
	}
	return nil
}

func runFile(filePath string, printAST bool) error {
	filePath = filepath.Clean(filePath)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", filePath, err)
	}

	tokens, err := lexer.Tokenize(string(bytes))
	if err != nil {
		return err
	}
	program, err := parser.Parse(tokens)
	if err != nil {
		return fmt.Errorf("syntax error: %w", err)
	}
	fmt.Printf("parsed %d top-level declaration(s)\n", len(program.Declarations))

	if printAST {
		fmt.Println()
		printer.Print(program)
	}
	return nil
}

func hasFlag(args []string, flag string) bool {
	for _, a := range args {
		if a == flag {
			return true
		}
	}
	return false
}

func filterOut(args []string, flag string) []string {
	out := make([]string, 0, len(args))
	for _, a := range args {
		if a != flag {
			out = append(out, a)
		}
	}
	return out
}
