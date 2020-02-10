package commands

import (
	"strconv"

	"github.com/spf13/cobra"
)

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * (fact(n - 1))
	}
	return 0
}

/*
Command to Return the factorial of the number n
*/

func Factorial() *cobra.Command {
	return &cobra.Command{
		Use:   "fact",
		Short: "Determines the factorial of the element n",
		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Println(err)
			}
			cmd.Println(fact(n))
			return nil
		},
	}
}
