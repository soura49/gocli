package commands

import (
	"time"

	"github.com/spf13/cobra"
)

/*
Usage for the command day
Description: used to display the currnet local time
*/

func PrintDay() *cobra.Command {
	return &cobra.Command{
		Use:   "day",
		Short: "Displays the Current local Day",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(time.Now().Day())
			return nil
		},
	}
}
