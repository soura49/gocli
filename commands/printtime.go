package commands

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

/*
Usage for the command is create
Description is need to create a aws CFT template in json format
*/

func PrintTime() *cobra.Command {
	return &cobra.Command{
		Use:   "time",
		Short: "Displays the current local time in 24-hour format",
		RunE: func(cmd *cobra.Command, args []string) error {
			h, m, s := time.Now().Clock()
			cmd.Println(strconv.Itoa(h) + ":" + strconv.Itoa(m) + ":" + strconv.Itoa(s))
			return nil
		},
	}
}
