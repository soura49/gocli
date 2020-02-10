package commands

import (
	"github.com/spf13/cobra"
)

func checkStrings(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return false
}

/*
Command to check the two strings are Equal
*/

func CheckAnagram() *cobra.Command {
	return &cobra.Command{
		Use:   "anagram",
		Short: "Two Checks Two Strings are Equal",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 2 {
				cmd.Println(checkStrings(args[0], args[1]))
			} else {
				cmd.Println("Need Exactly Two Arguments")
			}
			return nil
		},
	}
}
