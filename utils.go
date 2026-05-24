package main

import (
	"fmt"
	"sort"
)

type Command struct {
	Usage       string
	Description string
	Function    func(args []string)
}

var commands = map[string]Command{
	"add": {
		Usage:       "add <task description>",
		Description: "Add a new task to your to-do list.",
	},
	"list": {
		Usage:       "list",
		Description: "Show all of the current tasks.",
	},
	"delete": {
		Usage:       "delete <task id>",
		Description: "Remove a task by its id.",
	},
	"mark-todo": {
		Usage:       "mark-todo <task id>",
		Description: "Mark a task as todo using its id",
	},
	"mark-in-progress": {
		Usage:       "mark-in-progress <task id>",
		Description: "Mark a task as in progress using its.",
	},
	"mark-done": {
		Usage:       "mark-done <task id>",
		Description: "Mark a task as complete using its.",
	},
	"clear": {
		Usage:       "clear",
		Description: "Remove all tasks from your list.",
	},
}

func HelpCmd() {
	fmt.Println("All commands in todo CLI:")
	fmt.Println("----------------------------------")

	keys := make([]string, 0)
	for key := range commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %s\n\n", commands[key].Usage, commands[key].Description)
	}
}
