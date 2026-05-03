package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "magoito",
	Short: "Magoito compiler CLI",
	Long:  fmt.Sprintf("Magoito CLI for running programs and tests, Go version: %s\n", runtime.Version()),
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
