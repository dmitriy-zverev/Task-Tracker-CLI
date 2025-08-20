package main

import (
	"fmt"
	"log"

	"github.com/dmitriy-zverev/task-tracker-cli/taskHandler"
)

func main() {
	// Get our tasks, if data file not created then we'll create it
	tasks, err := taskHandler.LoadTasks()
	if err != nil {
		log.Fatal(err)
	}

	// Handling the operation passed from the terminal
	tasks, err = taskHandler.HandleOperation(tasks)
	if err != nil {
		fmt.Printf("error: %w", err)
	}
}
