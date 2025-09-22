package main

import (
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
		cmd.Function()
	}

}
