package commands

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func PrintTime() *cobra.Command {
	return &cobra.Command{
		Use: "time",
		RunE: func(cmd *cobra.Command, args []string) error {
			h, m, s := time.Now().Clock()
			cmd.Println(strconv.Itoa(h) + ":" + strconv.Itoa(m) + ":" + strconv.Itoa(s))
			return nil
		},
	}
}
