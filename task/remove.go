package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"rm"},
	Short:   "Resolve a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := Remove(args[0]); err != nil {
			fmt.Println("Error when remove a task:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
