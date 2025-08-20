package taskHandler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	id          int
	status      int
	isDeleted   bool
	description string
}

func LoadTasks() (tasks []Task, err error) {
	// Opening tasks data file, if not created then create
	file, err := os.OpenFile(DATA_FILENAME, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return tasks, err
	}
	file.Close()

	// Getting content from the file
	fileContent, err := os.ReadFile(DATA_FILENAME)
	if err != nil {
		return tasks, err
	}

	// Parsing data from file to string slice
	parsedData := strings.Split(string(fileContent), "\n")
	tasks = make([]Task, 0)

	// Converting strings to task objects
	for _, dataLine := range parsedData {
		splittedTask := strings.Split(dataLine, DATA_FILE_SEPARATOR)

		if len(splittedTask) < 4 {
			continue
		}

		taskId, _ := strconv.Atoi(splittedTask[0])
		taskStatus, _ := strconv.Atoi(splittedTask[1])
		taskIsDeleted, _ := strconv.ParseBool(splittedTask[2])

		tasks = append(tasks, Task{
			id:          taskId,
			status:      taskStatus,
			isDeleted:   taskIsDeleted,
			description: splittedTask[3],
		})
	}

	return tasks, nil
}

func appendToDataFile(task Task, dataFileName string) (err error) {
	// Creating a string to put in the data file
	fileDataString := fmt.Sprintf(
		"%d%s%d%s%v%s%s\n",
		task.id,
		DATA_FILE_SEPARATOR,
		task.status,
		DATA_FILE_SEPARATOR,
		task.isDeleted,
		DATA_FILE_SEPARATOR,
		task.description,
	)

	// Opening file
	file, err := os.OpenFile(dataFileName, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error: could not open file: %w", err)
	}

	// Writing to the file
	if _, err := file.WriteString(fileDataString); err != nil {
		return fmt.Errorf("error: could not write to file: %w", err)
	}

	// Successfully written, closing the file
	file.Close()
	return nil
}

func parseArgv() (operation string, param1 string, param2 string) {
	argc := len(os.Args)

	// Based on the number of argc return operation and params
	switch argc {
	case 2:
		return os.Args[1], "", ""
	case 3:
		return os.Args[1], os.Args[2], ""
	case 4:
		return os.Args[1], os.Args[2], os.Args[3]
	default:
		log.Fatal("Usage: task-tracker-cli <operation> [parameter 1] [parameter 2]")
	}

	return
}

func HandleOperation(tasks []Task) (newTasks []Task, err error) {
	operation, param1, _ := parseArgv()

	switch operation {
	case "add":
		if param1 == "" {
			log.Fatal("error: cannot add empty task")
		}

		tasks, err = handleAddTask(tasks, param1, DATA_FILENAME)
		if err != nil {
			return tasks, err
		}
	default:
		log.Fatal("Usage: task-tracker-cli <operation> [parameter 1] [parameter 2]")
	}

	return tasks, nil
}

func handleAddTask(tasks []Task, taskDescription, dataFileName string) (newTasks []Task, err error) {
	// Creating new task and adding it to tasks slice
	newTask := Task{id: len(tasks), status: TODO, description: taskDescription}
	tasks = append(tasks, newTask)

	// Writing new task to data file
	err = appendToDataFile(newTask, dataFileName)
	if err != nil {
		tasks = tasks[:len(tasks)-1]
		return tasks, fmt.Errorf("error: could not write to data file: %w", err)
	}

	fmt.Printf("Task added successfully (id: %d)\n", newTask.id)
	return tasks, nil
}
