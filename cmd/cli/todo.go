package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Task struct {
	Name string
	Done bool
}

type List []Task

func main() {
	lists := make([]Task, 0)
	var rootCmd = &cobra.Command{Use: "todolist"}
	var cmdAdd = &cobra.Command{
		Use:   "add [Task name]",
		Short: "Add task to the list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lists = append(lists, Task{
				args[0],
				false,
			})
			os.Exit(1)
		},
	}
	rootCmd.AddCommand(cmdAdd)

	var cmdList = &cobra.Command{
		Use:   "list ",
		Short: "show all list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lists)
			os.Exit(1)
		},
	}
	rootCmd.AddCommand(cmdList)

	var cmdDone = &cobra.Command{
		Use:   "done id ",
		Short: "complete the job",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lists = delarray(lists, args[0])
			os.Exit(1)
		},
	}
	rootCmd.AddCommand(cmdDone)

	// Implement other commands here

	rootCmd.Execute() // Don't change this
}
