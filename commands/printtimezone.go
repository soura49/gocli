package commands

import (
	"time"

	"github.com/spf13/cobra"
)

func PrintZone() *cobra.Command {
	return &cobra.Command{
		Use: "zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(time.Now().Zone())
			return nil
		},
	}
}
