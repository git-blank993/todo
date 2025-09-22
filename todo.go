package main

import (
	"fmt"
	"os"
	"strings"
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
		Function:    addCmd,
	},
	"list": {
		Usage:       "list",
		Description: "Show all of the current tasks.",
		Function:    listCmd,
	},
	"remove": {
		Usage:       "remove <task number>",
		Description: "Remove a task by its number.",
		Function:    removeCmd,
	},
	"complete": {
		Usage:       "complete <task number>",
		Description: "Mark a os.task as complete.",
		Function:    completeCmd,
	},
	"clear": {
		Usage:       "clear",
		Description: "Remove all tasks from your list.",
		Function:    clearCmd,
	},
}

func getTodos() ([]string, error) {
	content, err := os.ReadFile("todo.txt")
	if err != nil {
		return nil, err
	}
	todos := strings.Split(string(content), "\n")
	return todos, nil
}

func addTodo(input []string) (string, error) {
	f, err := os.OpenFile("todo.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	todo_result := strings.Join(input, " ") + "\n"
	todo := "[] " + strings.Join(input, " ") + "\n"
	if _, err := f.Write([]byte(todo)); err != nil {
		f.Close()
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	return todo_result, nil
}

func removeTodo(index int) (string, error) {
	todos, err := getTodos()
	if err != nil {
		return "", err
	}
	if index < 0 || index > len(todos) {
		return "", fmt.Errorf("provide a valid natural number")
	}
	newTodo := removeElement(todos, index)
	newTodoString := strings.Join(newTodo, "\n")
	if err := os.WriteFile("todo.txt", []byte(newTodoString), 0644); err != nil {
		return "", err
	}
	return todos[index], nil
}

func markDoneTodo(index int) (string, error) {
	todos, err := getTodos()
	if err != nil {
		return "", err
	}
	if index < 0 || index > len(todos) {
		return "", fmt.Errorf("provide a valid natural number")
	}
	todoData := strings.Split(todos[index], " ")
	if todoData[0] == "[x]" {
		st := fmt.Sprintf("Task: %s is already completed", todos[index])
		return st, nil
	}

	todoData[0] = "[x]"
	todos[index] = strings.Join(todoData, " ")
	newTodoString := strings.Join(todos, "\n")
	if err := os.WriteFile("todo.txt", []byte(newTodoString), 0644); err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", err
	}
	return fmt.Sprintf("Task: %s marked as completed", todos[index]), nil
}
