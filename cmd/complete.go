package cmd

import (
	"fmt"
	"strconv"
	"todo-list/internal/tasks"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		if err := tasks.CompleteTask(id); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
