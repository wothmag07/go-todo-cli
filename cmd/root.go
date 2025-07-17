package cmd

import (
	"fmt"
	"os"
	"todo-list/internal/tasks"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-list",
	Short: "A simple CLI to-do list application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		tasks.LoadTasks()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		tasks.SaveTasks()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
