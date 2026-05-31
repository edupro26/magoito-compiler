package main

import (
	"bytes"
	"fmt"
	"magoito-compiler/internal/ast"
	"magoito-compiler/internal/codegen"
	"magoito-compiler/internal/parser"
	"magoito-compiler/internal/semantics"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type Error struct {
	File string
	Err  error
}

type TestGroup struct {
	Name   string
	Passed int
	Failed int
	Errors []Error
}

var showFailed bool
var showAll bool
var maxShow int

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run the test suite",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !showFailed && showAll || !showFailed && maxShow != 5 {
			return fmt.Errorf("the --show-failed flag must be set to use --all or --max")
		}
		dirs := []string{
			filepath.Join("test", "invalid-syntax"),
			filepath.Join("test", "invalid-semantics"),
			filepath.Join("test", "valid"),
		}
		var results []TestGroup
		for _, dir := range dirs {
			group, err := runTests(dir)
			if err != nil {
				return err
			}
			results = append(results, group)
		}
		printSummary(results)
		if showFailed {
			printErrors(results)
		}
		return nil
	},
}

func init() {
	testCmd.Flags().BoolVar(&showFailed, "show-failed", false, "Show failing tests")
	testCmd.Flags().BoolVar(&showAll, "all", false, "Show all failing tests")
	testCmd.Flags().IntVar(&maxShow, "max", 5, "Maximum number of failing tests to display")
	rootCmd.AddCommand(testCmd)
}

func runTests(testDir string) (TestGroup, error) {
	g := TestGroup{Name: filepath.Base(testDir)}
	entries, err := os.ReadDir(testDir)
	if err != nil {
		return g, err
	}
	for _, entry := range entries {
		fileDir := filepath.Join(testDir, entry.Name())
		files, err := filepath.Glob(filepath.Join(fileDir, "*.mag"))
		if err != nil || len(files) == 0 {
			continue
		}

		passed, err := runTest(files[0], g.Name != "valid")
		if passed {
			g.Passed++
		} else {
			g.Failed++
			g.Errors = append(g.Errors, Error{File: files[0], Err: err})
		}
	}
	return g, nil
}

func runTest(path string, expectError bool) (bool, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	cst, err := parser.ParseProgram(string(bytes))
	if err == nil {
		program := ast.Build(cst)
		err = semantics.Validate(program)
		if err == nil && !expectError {
			if err := runValid(path, program); err != nil {
				return false, err
			}
		}
	}
	if (err != nil) == expectError {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, fmt.Errorf("expected an error but program succeeded")
}

func runValid(path string, program *ast.Program) error {
	module, err := codegen.Generate(program)
	if err != nil {
		return err
	}
	llPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".ll"
	if err := os.WriteFile(llPath, []byte(module.String()), 0644); err != nil {
		return err
	}
	expectPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".expect"
	expectBytes, err := os.ReadFile(expectPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("lli", llPath)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	if !bytes.Equal(output, expectBytes) {
		return fmt.Errorf("output mismatch: expected %q, got %q", string(expectBytes), string(output))
	}
	return nil
}

func printErrors(groups []TestGroup) {
	for _, g := range groups {
		if g.Failed == 0 {
			continue
		}

		var limit int
		if showAll {
			limit = len(g.Errors)
		} else if maxShow >= len(g.Errors) {
			limit = len(g.Errors)
		} else if maxShow < len(g.Errors) {
			limit = maxShow
		}
		fmt.Printf("Failing %s tests:\n", g.Name)
		for i := 0; i < limit; i++ {
			e := g.Errors[i]
			if g.Name == "valid" {
				fmt.Printf("	%s\n		%v\n", e.File, e.Err)
			} else {
				fmt.Printf("	%s\n", e.File)
			}
		}

		if limit < len(g.Errors) {
			fmt.Printf("	...and %d more errors hidden\n", len(g.Errors)-limit)
		}
	}
}

func printSummary(results []TestGroup) {
	fmt.Printf("Test Summary:\n")
	for _, r := range results {
		fmt.Printf("Passed: %d, Failed: %d (%s)\n",
			r.Passed, r.Failed, r.Name)
	}
}
