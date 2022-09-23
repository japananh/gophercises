package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"add"},
	Short:   "Add a task",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := Add(args[0], args[1]); err != nil {
			fmt.Println("Error when add a task:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
