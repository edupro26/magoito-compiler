package main

import (
	"fmt"
	"magoito-compiler/cli/printer"
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/codegen"
	"magoito-compiler/internal/parser"
	"magoito-compiler/internal/semantics"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var printAST bool
var printPretty bool
var example bool

var runFileCmd = &cobra.Command{
	Use:   "run [file.mag]",
	Short: "Run a MAGOITO file",
	Args: func(cmd *cobra.Command, args []string) error {
		if !example && len(args) == 0 {
			return fmt.Errorf("must provide a file or --example")
		}
		if example && len(args) > 0 {
			return fmt.Errorf("cannot use both file and --example")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var file string
		if example {
			file = "example.mag"
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
	runFileCmd.Flags().BoolVar(&example, "example", false, "Run the MAGOITO example")
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
		fmt.Printf("%s: %v\n", filePath, err)
		return nil
	}
	tree := ast.Build(cst)
	if err := semantics.Validate(tree); err != nil {
		fmt.Printf("%s: %v\n", filePath, err)
		return nil
	}
	module, err := codegen.Generate(tree)
	if err != nil {
		return err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	base := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	llPath := filepath.Join(cwd, base+".ll")
	if err := os.WriteFile(llPath, []byte(module.String()), 0644); err != nil {
		return err
	}
	if printAST {
		printer.Print(tree)
	} else if printPretty {
		printer.Pretty(tree)
	}
	return nil
}
