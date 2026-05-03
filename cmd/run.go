package main

import (
	"fmt"
	"magoito-compiler/cmd/printer"
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/parser"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var printAST bool
var printPretty bool
var example int

var runFileCmd = &cobra.Command{
	Use:   "run [file.mag]",
	Short: "Run a MAGOITO file",
	Args: func(cmd *cobra.Command, args []string) error {
		if example == 0 && len(args) == 0 {
			return fmt.Errorf("must provide a file or --example")
		}
		if example != 0 && len(args) > 0 {
			return fmt.Errorf("cannot use both file and --example")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var file string
		if example != 0 {
			if example < 1 || example > 3 {
				return fmt.Errorf("example must be 1–3")
			}
			file = fmt.Sprintf("example%02d.mag", example)
		} else {
			file = args[0]
		}
		if printAST && printPretty {
			return fmt.Errorf("cannot use both --ast and --pretty")
		}
		return runFile(file)
	},
}

func init() {
	runFileCmd.Flags().BoolVar(&printAST, "ast", false, "Print AST in raw output")
	runFileCmd.Flags().BoolVar(&printPretty, "pretty", false, "Print AST in pretty output")
	runFileCmd.Flags().IntVar(&example, "example", 0, "Run a MAGOITO example")
	rootCmd.AddCommand(runFileCmd)
}

func runFile(filePath string) error {
	filePath = filepath.Clean(filePath)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", filePath, err)
	}
	cst, err := parser.ParseProgram(string(bytes))
	if err != nil {
		return fmt.Errorf("%s:%w", filePath, err)
	}
	tree := ast.Build(cst)
	if printAST {
		printer.Print(tree)
	} else if printPretty {
		printer.Pretty(tree)
	}
	return nil
}
