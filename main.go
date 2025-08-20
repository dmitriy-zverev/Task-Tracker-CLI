package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("This is your task tracker CLI...")

	fmt.Println("Opening tasks.txt file...")
	file, err := os.OpenFile("tasks.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tasks.txt opened.")

	fmt.Println("Writing to tasks.txt...")

	if _, err := file.Write([]byte("first task is always easy\n")); err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Println("Writing successful.")

	fmt.Println("Reading tasks.txt...")

	data, err := os.ReadFile("tasks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("From tasks.txt:")
	fmt.Println(string(data))
}
