package commands

import (
	"time"

	"github.com/spf13/cobra"
)

func PrintDay() *cobra.Command {
	return &cobra.Command{
		Use: "day",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(time.Now().Day())
			return nil
		},
	}
}
