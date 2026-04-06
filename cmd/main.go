package main

import (
	"bufio"
	"fmt"
	"magoito-compiler/cmd/util"
	"os"
	"strings"
)

func main() {
	util.PrintBanner()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
		if !scanner.Scan() {
			fmt.Println()
			return
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		if args[0] == "quit" || args[0] == "q" {
			return
		}
		if args[0] == "help" {
			util.PrintHelp()
			continue
		}
		if args[0] != "mag" {
			fmt.Println(`unknown command. Try "help" for information`)
			continue
		}
		if len(args) == 1 {
			fmt.Println(`unknown command. Try "help" for information`)
			continue
		}

		switch args[1] {
		case "run":
			if err := util.RunCmd(args[2:]); err != nil {
				fmt.Println(err)
			}
		case "help":
			util.PrintHelp()
		case "quit", "q":
			return
		default:
			fmt.Println(`unknown command. Try "help" for information`)
		}
	}
}
