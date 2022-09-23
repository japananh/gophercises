package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

const isCompletedFlag = "isCompleted"

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"list"},
	Short:   "List all tasks",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		isCompleted, err := cmd.Flags().GetBool(isCompletedFlag)
		if err != nil {
			fmt.Println("Error when list task:", err)
		}
		if err := List(&Filter{IsCompleted: isCompleted}); err != nil {
			fmt.Println("Error when list task:", err)
		}
	},
}
var isCompleted bool

func init() {
	listCmd.Flags().BoolVarP(&isCompleted, isCompletedFlag, "c", false, "Show all completed tasks")
	rootCmd.AddCommand(listCmd)
}
