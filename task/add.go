package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

var onlyDigits bool
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"add"},
	Short:   "Add a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Add(args[0]))
	},
}

func init() {
	addCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(addCmd)
}
