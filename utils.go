package main

import "fmt"

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
