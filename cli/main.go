package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "magoito [file.mag]",
	Short:                 "Magoito compiler CLI",
	Long:                  fmt.Sprintf("Magoito CLI for running programs and tests, Go version: %s\n", runtime.Version()),
	Args:                  cobra.MaximumNArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return compileProgram(args[0])
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
