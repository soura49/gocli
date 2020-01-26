package commands

import (
	"time"

	"github.com/spf13/cobra"
)

/*
Usage for the command is create
Description is need to create a aws CFT template in json format
*/

func PrintMonth() *cobra.Command {
	return &cobra.Command{
		Short: "Displays the Current month of the Year",
		Use:   "month",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(time.Now().Month())
			return nil
		},
	}
}
