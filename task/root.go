package task

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "task",
	Version: version,
	Short:   "task - a simple CLI to transform and inspect strings",
	Long: `task is a super fancy CLI (kidding)
   
One can use task to manage todos by adding/accomplishing a task or preview all tasks straight from the terminal`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		return err
	}
	return nil
}
