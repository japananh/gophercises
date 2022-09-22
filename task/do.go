package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:     "do",
	Aliases: []string{"do"},
	Short:   "Resolve a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := Resolve(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
