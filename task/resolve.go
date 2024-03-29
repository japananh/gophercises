package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:     "do",
	Aliases: []string{"do"},
	Short:   "Resolve a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := Resolve(args[0]); err != nil {
			fmt.Println("Error when resolve a task:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(resolveCmd)
}
