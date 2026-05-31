package main

import (
	"fmt"
	"magoito-compiler/cli/printer"
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/codegen"
	"magoito-compiler/internal/parser"
	"magoito-compiler/internal/semantics"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/spf13/cobra"
)

var printAST bool
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
		return runFile(file)
	},
}

func init() {
	runFileCmd.Flags().BoolVar(&printAST, "ast", false, "Print AST in raw output")
	runFileCmd.Flags().BoolVar(&example, "example", false, "Run the MAGOITO example")
	rootCmd.AddCommand(runFileCmd)
}

func runFile(filePath string) error {
	tree, err := loadProgram(filePath)
	if err != nil || tree == nil {
		return err
	}
	if printAST {
		printer.Print(tree)
		return nil
	}
	module, err := codegen.Generate(tree)
	if err != nil {
		return err
	}
	return runModule(module)
}

func compileProgram(filePath string) error {
	tree, err := loadProgram(filePath)
	if err != nil || tree == nil {
		return err
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
	return os.WriteFile(llPath, []byte(module.String()), 0644)
}

func loadProgram(filePath string) (*ast.Program, error) {
	filePath = filepath.Clean(filePath)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %q: %w", filePath, err)
	}
	cst, err := parser.ParseProgram(string(bytes))
	if err != nil {
		fmt.Printf("%s: %v\n", filePath, err)
		return nil, nil
	}
	tree := ast.Build(cst)
	if err := semantics.Validate(tree); err != nil {
		fmt.Printf("%s: %v\n", filePath, err)
		return nil, nil
	}
	return tree, nil
}

func runModule(module *ir.Module) error {
	tmp, err := os.CreateTemp("", "magoito-*.ll")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.WriteString(module.String()); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	cmd := exec.Command("lli", tmp.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("lli failed: %s", string(output))
	}
	_, _ = os.Stdout.Write(output)
	return nil
}
