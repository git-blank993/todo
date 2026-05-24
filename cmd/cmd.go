package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func AddCmd(args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("No todo provided")
	}
	return args[0], nil
}

func ListCmd(args []string) string {
	if len(args) < 1 {
		return "all"
	} else {
		return args[0]
	}
}

func DeleteCmd(args []string) (int, error) {
	if len(args) < 1 {
		return -1, fmt.Errorf("No todo number provided")
	}
	if len(args) > 1 {
		return -1, fmt.Errorf("Only one element can be removed at a time")
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, fmt.Errorf("Provide a valid natural number")
	}
	return index, nil
}

func ClearCmd(args []string) (bool, error) {
	fmt.Println("Are you sure you want to clear all tasks?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val := scanner.Text()
	if err := scanner.Err(); err != nil {
		return false, err
	}
	if val == "y" || val == "yes" || val == "Yes" {
		return true, nil
	}
	return false, nil
}

func UpdateStatusCmd(args []string) (int, error) {
	if len(args) < 1 {
		return -1, fmt.Errorf("No todo number provided")
	}
	if len(args) > 1 {
		return -1, fmt.Errorf("Only one todo can be completed at a time")
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, fmt.Errorf("Provide a valid natural number")
	}
	return index, nil
}

func UpdateDescCmd(args []string) (int, string, error) {
	if len(args) < 2 {
		return -1, "", fmt.Errorf("Provide both task Id and Description. First task Id and then Task Description in order")
	}
	index, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, "", fmt.Errorf("Provide a valid natural number")
	}
	description := args[1]
	if description == "" {
		return -1, "", fmt.Errorf("Provide a valid description")
	}

	return index, description, nil
}

// func helpCmd() {
// 	fmt.Println("All commands in todo CLI:")
// 	fmt.Println("----------------------------------")
//
// 	keys := make([]string, 0)
// 	for key := range commands {
// 		keys = append(keys, key)
// 	}
// 	sort.Strings(keys)
// 	for _, key := range keys {
// 		fmt.Printf("%s: %s\n\n", commands[key].Usage, commands[key].Description)
// 	}
// }
