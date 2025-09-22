package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	todo := strings.Join(input, " ") + "\n"
	if _, err := f.Write([]byte(todo)); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return todo
}

func removeTodo(index int) string{
	todos := getTodos()
	if index < 0 || index > len(todos) {
		log.Fatal("Provide a valid natural number")
	}
	newTodos := removeElement(todos, index)
	newTodosString := strings.Join(newTodos, "\n")
	if err := os.WriteFile("todo.txt", []byte(newTodosString), 0644); err != nil {
			log.Fatal(err)
	}
	return todos[index]
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

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No command provided")
	}
	command := os.Args[1]
	switch command {

	case "add":
		if len(os.Args) < 3 {
			log.Fatal("No todo provided")
		}
		todo := addTodo(os.Args[2:])
		fmt.Printf("Added new todo %s", todo)
	case "list":
		todos := getTodos()
		for i, items := range todos {
			if items != "" && items != "\n" {
				fmt.Printf("%d. %s\n", i+1, items)
			}
		}
	case "remove":
		
		if len(os.Args) < 3 {
			log.Fatal("No todo number provided")
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Provide a valid natural number")
		}
		todo := removeTodo(index - 1)
		fmt.Printf("Removed todo %d. %s", index, todo)
	}

}

