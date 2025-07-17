package cmd

import (
	"todo-list/internal/tasks"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
