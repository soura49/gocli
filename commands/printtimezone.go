package commands

import (
	"time"

	"github.com/spf13/cobra"
)

/*
Usage for the command is create
Description is need to create a aws CFT template in json format
*/

func PrintZone() *cobra.Command {
	return &cobra.Command{
		Use:   "zone",
		Short: "Displays the Current Local Time Zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(time.Now().Zone())
			return nil
		},
	}
}
