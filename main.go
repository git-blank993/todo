package main

import (
	"fmt"
	"log"
	"os"

	"github.com/git-blank993/todo/cmd"
	"github.com/git-blank993/todo/task"
)

func main() {

	if len(os.Args) < 2 {
		HelpCmd()
		log.Fatal("No command provided")
	}
	command := os.Args[1]
	store := task.NewTaskStore("./task.json")
	manager := task.NewTaskManager(store)
	store.Read()
	switch command {
	case "add":
		description, err := cmd.AddCmd(os.Args[2:])
		if err != nil {

		}
		task := task.NewTask(manager.NewId(), description)
		manager.Add(task)
	case "list":
		filter := cmd.ListCmd(os.Args[2:])
		switch filter {
		case "all":
			manager.List()
		case "done":
			manager.FilterList(task.Done)
		case "in-progress":
			manager.FilterList(task.InProgress)
		case "todo":
			manager.FilterList(task.Todo)
		}
	case "delete":
		id, err := cmd.DeleteCmd(os.Args[2:])
		if err != nil {
			fmt.Println("Error deleting task:\n", err.Error())
		}
		manager.Delete(id)
	case "update":
		id, desc, err := cmd.UpdateDescCmd(os.Args[2:])
		if err != nil {
			fmt.Println("Error updating task:\n", err.Error())
		}
		manager.UpdateDescription(id, desc)
	case "mark-todo":
		id, err := cmd.UpdateStatusCmd(os.Args[2:])
		if err != nil {
			fmt.Println("Error updating status:\n", err.Error())
		}
		manager.UpdateStatus(id, task.Todo)

	case "mark-in-progress":
		id, err := cmd.UpdateStatusCmd(os.Args[2:])
		if err != nil {
			fmt.Println("Error updating status:\n", err.Error())
		}
		manager.UpdateStatus(id, task.InProgress)

	case "mark-done":
		id, err := cmd.UpdateStatusCmd(os.Args[2:])
		if err != nil {
			fmt.Println("Error updating status:\n", err.Error())
		}
		manager.UpdateStatus(id, task.Done)

	case "clear":
		if ok, err := cmd.ClearCmd(os.Args[2:]); err != nil {
			fmt.Println("Error clearing all tasks:\n", err.Error())
		} else {
			if ok {
				manager.Clear()
			} else {
				fmt.Println("")
			}
		}

	default:
		HelpCmd()
	}

}
