package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"list"},
	Short:   "List all tasks",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(List())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
