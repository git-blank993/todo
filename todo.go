package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Command struct {
	Usage       string
	Description string
	Function    func()
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
		Description: "Mark a task as complete.",
		Function:    completeCmd,
	},
	"clear": {
		Usage:       "clear",
		Description: "Remove all tasks from your list.",
		Function:    clearCmd,
	},
}

func removeElement(arr []string, index int) []string {
	a := make([]string, 0)
	if index < 0 || index > len(arr) {
		fmt.Println("Warning: Index out of bound returning the original array")
		return arr
	}
	a = append(a, arr[:index]...)
	a = append(a, arr[index+1:]...)
	return a
}

func getTodos() []string {
	content, err := os.ReadFile("todo.txt")
	if err != nil {
		log.Fatal(err)
	}
	todos := strings.Split(string(content), "\n")
	return todos
}

func addTodo(input []string) string {
	f, err := os.OpenFile("todo.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	todo_result := strings.Join(input, " ") + "\n"
	todo := "[] " + strings.Join(input, " ") + "\n"
	if _, err := f.Write([]byte(todo)); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return todo_result
}

func removeTodo(index int) string {
	todos := getTodos()
	if index < 0 || index > len(todos) {
		log.Fatal("Provide a valid natural number")
	}
	newTodo := removeElement(todos, index)
	newTodoString := strings.Join(newTodo, "\n")
	if err := os.WriteFile("todo.txt", []byte(newTodoString), 0644); err != nil {
		log.Fatal(err)
	}
	return todos[index]
}

func markDoneTodo(index int) string {
	todos := getTodos()
	if index < 0 || index > len(todos) {
		log.Fatal("Provide a valid natural number")
	}
	todoData := strings.Split(todos[index], " ")
	if todoData[0] == "[x]" {
		st := fmt.Sprintf("Task: %s is already completed", todos[index])
		return st
	}

	todoData[0] = "[x]"
	todos[index] = strings.Join(todoData, " ")
	newTodoString := strings.Join(todos, "\n")
	if err := os.WriteFile("todo.txt", []byte(newTodoString), 0644); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Task: %s marked as completed", todos[index])
}

func addCmd() {
	if len(os.Args) < 3 {
		log.Fatal("No todo provided")
	}
	todo := addTodo(os.Args[2:])
	fmt.Printf("Added: %s", todo)
}

func listCmd() {
	todos := getTodos()
	for i, items := range todos {
		if items != "" && items != "\n" {
			completed := ""
			if strings.Split(items, " ")[0] == "[x]" {
				completed = "(✓)"
			}
			item := strings.Split(items, " ")[1:]
			itemString := strings.Join(item, " ")
			fmt.Printf("%d. %s \t %s\n", i+1, itemString, completed)
		}
	}
}

func removeCmd() {
	if len(os.Args) < 3 {
		log.Fatal("No todo number provided")
	}
	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Provide a valid natural number")
	}
	todo := removeTodo(index - 1)
	fmt.Printf("Removed: %d. %s", index, todo)
}

func clearCmd() {
	if err := os.WriteFile("todo.txt", []byte(""), 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Removed all todo list")
}

func completeCmd() {
	if len(os.Args) < 3 {
		log.Fatal("No todo number provided")
	}
	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Provide a valid natural number")
	}
	result := markDoneTodo(index - 1)
	fmt.Println(result)
}

func helpCmd() {
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
