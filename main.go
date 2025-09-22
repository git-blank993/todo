package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		helpCmd()
		log.Fatal("No command provided")
	}
	command := os.Args[1]

	if cmd, found := commands[command]; found {
		cmd.Function(os.Args[2:])
	} else {
		fmt.Printf("No command named %s found\n\n", command)
		helpCmd()
	}

}
