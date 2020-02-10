package main

/*
Simple Cli created to learn how cobra framework works
*/

import (
	"os"

	"github.com/choascli/commands"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "Welcome Time Series",
	}
	cmd.AddCommand(commands.PrintTime()) //
	cmd.AddCommand(commands.PrintMonth())
	cmd.AddCommand(commands.PrintDay())
	cmd.AddCommand(commands.PrintZone())
	cmd.AddCommand(commands.CreateJsonFile())
	cmd.AddCommand(commands.Factorial())
	cmd.AddCommand(commands.FibtoN())
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
