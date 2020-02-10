package commands

import (
	"os/exec"

	"github.com/spf13/cobra"
)

/*
Command to Generate UUID
*/

func GenerateUUID() *cobra.Command {
	return &cobra.Command{
		Use:   "genuuid",
		Short: "To Geneate the UUID",
		RunE: func(cmd *cobra.Command, args []string) error {
			out, err := exec.Command("uuidgen").Output()
			if err != nil {
				cmd.Println(err)
			}
			cmd.Print(string(out))
			return nil
		},
	}
}
