package commands

import (
	"strconv"

	"github.com/spf13/cobra"
)

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

/*
Command to Return the fibonnaci of the number n
*/

func FibtoN() *cobra.Command {
	return &cobra.Command{
		Use:   "fib",
		Short: "Prints the Fibonacci of the elements",
		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Println(err)
			}
			for index := 1; index < n; index++ {
				s := strconv.Itoa(fib(index))
				cmd.Println(s + " ")
			}
			return nil
		},
	}

}
