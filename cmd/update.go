package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"todo-list/internal/tasks"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [new task name]",
	Short: "Update a task name",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		newName := strings.Join(args[1:], " ")
		if err := tasks.UpdateTask(id, newName); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
} 