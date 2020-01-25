package commands

import (
	"time"

	"github.com/spf13/cobra"
)

func PrintMonth() *cobra.Command {
	return &cobra.Command{
		Use: "month",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(time.Now().Month())
			return nil
		},
	}
}
