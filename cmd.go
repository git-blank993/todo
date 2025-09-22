package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func addCmd(args []string) {
	if len(args) < 1 {
		log.Fatal("No todo provided")
	}
	todo, err := addTodo(args)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Added: %s", todo)
}

func listCmd(args []string) {
	todos, err := getTodos()
	if err != nil {
		os.Exit(1)
	}
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

func removeCmd(args []string) {
	if len(args) < 1 {
		log.Fatal("No todo number provided")
	}
	if len(args) > 1 {
		log.Fatal("Only one element can be removed at a time")
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Provide a valid natural number")
	}
	todo, err := removeTodo(index - 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Removed: %d. %s", index, todo)
}

func clearCmd(args []string) {
	filePath, err := getTodoFilePath()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(filePath, []byte(""), 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Removed all todo list")
}

func completeCmd(args []string) {
	if len(args) < 1 {
		log.Fatal("No todo number provided")
	}
	if len(args) > 1 {
		log.Fatal("Only one todo can be completed at a time")
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Provide a valid natural number")
	}
	result, err := markDoneTodo(index - 1)
	if err != nil {
		log.Fatal(err)
	}
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
