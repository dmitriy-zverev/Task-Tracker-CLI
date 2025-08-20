package taskHandler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	createdAt   time.Time
	updatedAt   time.Time
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

		if len(splittedTask) < 6 {
			continue
		}

		taskId, _ := strconv.Atoi(splittedTask[0])
		taskStatus, _ := strconv.Atoi(splittedTask[1])
		taskIsDeleted, _ := strconv.ParseBool(splittedTask[2])
		taskCreatedAt, _ := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", splittedTask[4])
		taskUpdatedAt, _ := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", splittedTask[5])

		tasks = append(tasks, Task{
			id:          taskId,
			status:      taskStatus,
			isDeleted:   taskIsDeleted,
			description: splittedTask[3],
			createdAt:   taskCreatedAt,
			updatedAt:   taskUpdatedAt,
		})
	}

	return tasks, nil
}

func createDataFileString(task Task) (dataFileString string) {
	return fmt.Sprintf(
		"%d%s%d%s%v%s%s%s%v%s%v\n",
		task.id,
		DATA_FILE_SEPARATOR,
		task.status,
		DATA_FILE_SEPARATOR,
		task.isDeleted,
		DATA_FILE_SEPARATOR,
		task.description,
		DATA_FILE_SEPARATOR,
		task.createdAt,
		DATA_FILE_SEPARATOR,
		task.updatedAt,
	)
}

func appendToDataFile(task Task, dataFileName string) (err error) {
	// Creating a string to put in the data file
	fileDataString := createDataFileString(task)

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

func updateDataFile(tasks []Task, dataFileName string) (err error) {
	// Creating file content from scratch
	fileContent := ""

	for _, task := range tasks {
		fileContent += createDataFileString(task)
	}

	// Overwriting the file
	data := []byte(fileContent)
	if err = os.WriteFile(dataFileName, data, 0644); err != nil {
		return fmt.Errorf("error: could not overwrite the data file: %w", err)
	}

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
	operation, param1, param2 := parseArgv()

	switch strings.ToLower(operation) {
	case ADD:
		if param1 == "" {
			log.Fatal("error: cannot add empty task")
		}

		tasks, err = handleAddTask(tasks, param1, DATA_FILENAME)
		if err != nil {
			return tasks, err
		}
	case UPDATE:
		if param1 == "" {
			log.Fatal("error: you did not specify the ID of the task")
		}
		if param2 == "" {
			log.Fatal("error: cannot update task with empty description")
		}

		taskId, _ := strconv.Atoi(param1)
		if taskId >= len(tasks) || taskId < 0 {
			log.Fatal("error: invalid ID of the task")
		}

		tasks, err = handleUpdateTask(tasks, taskId, param2, DATA_FILENAME)
		if err != nil {
			return tasks, err
		}
	case DELETE:
		if param1 == "" {
			log.Fatal("error: you did not specify the ID of the task")
		}

		taskId, _ := strconv.Atoi(param1)
		if taskId >= len(tasks) || taskId < 0 {
			log.Fatal("error: invalid ID of the task")
		}

		tasks, err = handleDeleteTask(tasks, taskId, DATA_FILENAME)
		if err != nil {
			return tasks, err
		}
	case MARK_IN_PROGRESS:
		if param1 == "" {
			log.Fatal("error: you did not specify the ID of the task")
		}

		taskId, _ := strconv.Atoi(param1)
		if taskId >= len(tasks) || taskId < 0 {
			log.Fatal("error: invalid ID of the task")
		}

		tasks, err = handleMarkTaskInProgress(tasks, taskId, DATA_FILENAME)
		if err != nil {
			return tasks, err
		}
	case MARK_DONE:
		if param1 == "" {
			log.Fatal("error: you did not specify the ID of the task")
		}

		taskId, _ := strconv.Atoi(param1)
		if taskId >= len(tasks) || taskId < 0 {
			log.Fatal("error: invalid ID of the task")
		}

		tasks, err = handleMarkTaskDone(tasks, taskId, DATA_FILENAME)
		if err != nil {
			return tasks, err
		}
	case LIST:
		// Print all tasks
		if param1 == "" {
			handleListAll(tasks)
			return tasks, nil
		}

		status := -1
		switch strings.ToLower(param1) {
		case "todo":
			status = TODO
		case "done":
			status = DONE
		case "in-progress":
			status = IN_PROGRESS
		default:
			log.Fatal("error: unknown type of tasks")
		}

		// Print only todo tasks
		if status == TODO {
			handleListTodo(tasks)
		}

		// Print only done tasks
		if status == DONE {
			handleListDone(tasks)
		}

		// Print only in progress tasks
		if status == IN_PROGRESS {
			handleListInProgress(tasks)
		}
	default:
		log.Fatal("Usage: task-tracker-cli <operation> [parameter 1] [parameter 2]")
	}

	return tasks, nil
}

func handleAddTask(tasks []Task, taskDescription, dataFileName string) (newTasks []Task, err error) {
	// Creating new task and adding it to tasks slice
	newTask := Task{
		id:          len(tasks),
		status:      TODO,
		description: taskDescription,
		createdAt:   time.Now().UTC(),
		updatedAt:   time.Now().UTC(),
	}
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

func handleUpdateTask(tasks []Task, taskId int, newDesc, dataFileName string) (newTasks []Task, err error) {
	// Updating task with new description, saving old description and updatedAt in case of errors
	oldDescription := tasks[taskId].description
	oldUpdatedAt := tasks[taskId].updatedAt

	tasks[taskId].description = newDesc
	tasks[taskId].updatedAt = time.Now().UTC()

	// Rewriting the whole file
	err = updateDataFile(tasks, dataFileName)
	if err != nil {
		tasks[taskId].description = oldDescription
		tasks[taskId].updatedAt = oldUpdatedAt
		return tasks, fmt.Errorf("error: could not write to file: %w", err)
	}

	fmt.Printf("Successfully updated task (id: %d)\n", taskId)
	return tasks, nil
}

func handleDeleteTask(tasks []Task, taskId int, dataFileName string) (newTasks []Task, err error) {
	// Change task property isDeleted to true
	tasks[taskId].isDeleted = true
	tasks[taskId].updatedAt = time.Now().UTC()

	// Rewriting the whole file
	err = updateDataFile(tasks, dataFileName)
	if err != nil {
		tasks[taskId].isDeleted = false
		return tasks, fmt.Errorf("error: could not write to file: %w", err)
	}

	fmt.Printf("Successfully deleted task (id: %d)\n", taskId)
	return tasks, nil
}

func handleMarkTaskInProgress(tasks []Task, taskId int, dataFileName string) (newTasks []Task, err error) {
	// Change task property isDeleted to true
	tasks[taskId].status = IN_PROGRESS
	tasks[taskId].updatedAt = time.Now().UTC()

	// Rewriting the whole file
	err = updateDataFile(tasks, dataFileName)
	if err != nil {
		tasks[taskId].isDeleted = false
		return tasks, fmt.Errorf("error: could not write to file: %w", err)
	}

	fmt.Printf("Marked task (id: %d) as in progress\n", taskId)
	return tasks, nil
}

func handleMarkTaskDone(tasks []Task, taskId int, dataFileName string) (newTasks []Task, err error) {
	// Change task property isDeleted to true
	oldStatus := tasks[taskId].status
	oldUpdatedAt := tasks[taskId].updatedAt
	tasks[taskId].status = DONE
	tasks[taskId].updatedAt = time.Now().UTC()

	// Rewriting the whole file
	err = updateDataFile(tasks, dataFileName)
	if err != nil {
		tasks[taskId].status = oldStatus
		tasks[taskId].updatedAt = oldUpdatedAt
		return tasks, fmt.Errorf("error: could not write to file: %w", err)
	}

	fmt.Printf("Marked task (id: %d) as done\n", taskId)
	return tasks, nil
}

func handleListAll(tasks []Task) {
	fmt.Printf("\n==========TODO LIST==========\n\n")
	fmt.Printf("——————————All Tasks——————————\n\n")

	for _, task := range tasks {
		if task.status != DONE {
			fmt.Printf("[ ]")
		} else {
			fmt.Printf("[X]")
		}

		fmt.Printf(" (id: %d)", task.id)
		fmt.Printf(" %s", task.description)

		if task.isDeleted {
			fmt.Printf(" (deleted)")
		}

		fmt.Printf("\n")
	}
}

func handleListTodo(tasks []Task) {
	fmt.Printf("\n==========TODO LIST==========\n\n")
	fmt.Printf("—————————Todo Tasks—————————\n\n")

	for _, task := range tasks {
		if task.status == TODO {
			fmt.Printf("[ ] (id: %d) %s", task.id, task.description)

			if task.isDeleted {
				fmt.Printf(" (deleted)")
			}

			fmt.Printf("\n")
		}
	}
}

func handleListDone(tasks []Task) {
	fmt.Printf("\n==========TODO LIST==========\n\n")
	fmt.Printf("—————————Done Tasks——————--—\n\n")

	for _, task := range tasks {
		if task.status == DONE {
			fmt.Printf("[X] (id: %d) %s", task.id, task.description)

			if task.isDeleted {
				fmt.Printf(" (deleted)")
			}

			fmt.Printf("\n")
		}
	}
}

func handleListInProgress(tasks []Task) {
	fmt.Printf("\n==========TODO LIST==========\n\n")
	fmt.Printf("——————In Progress Tasks——————\n\n")

	for _, task := range tasks {
		if task.status == IN_PROGRESS {
			fmt.Printf("[ ] (id: %d) %s", task.id, task.description)

			if task.isDeleted {
				fmt.Printf(" (deleted)")
			}

			fmt.Printf("\n")
		}
	}
}
