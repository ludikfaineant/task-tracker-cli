package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker-cli/internal/model"
	"task-tracker-cli/internal/service"
)

var dataFile = "task.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task add|list|update|delete|mark-done|mark-in-progress")
		os.Exit(1)
	}
	srv := service.New()
	loadTasks(srv)

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task add <description>")
			os.Exit(1)
		}
		description := strings.Join(os.Args[2:], " ")
		srv.AddTask(description)

	case "list":
		by := ""
		if len(os.Args) > 2 {
			by = os.Args[2]
		}
		filter := srv.ListTasks(by)
		printTasks(filter)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id must be of type int")
			os.Exit(1)
		}
		err = srv.DeleteTask(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task update <id> <description>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id must be of type int")
			os.Exit(1)
		}
		err = srv.UpdateTask(id, strings.Join(os.Args[3:], " "))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task mark-done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id must be of type int")
			os.Exit(1)
		}
		err = srv.MarkDone(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task mark-in-progress <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id must be of type int")
			os.Exit(1)
		}
		err = srv.MarkInProgress(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
	saveTasks(srv)
}

func loadTasks(s *service.TaskService) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	var tasks []model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	s.Load(tasks)
}

func saveTasks(s *service.TaskService) {
	data, err := json.MarshalIndent(s.Tasks(), "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding tasks: %v\n", err)
		os.Exit(1)
	}
	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}
}

func printTasks(tasks []model.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("[%d] %s [%s] (created: %s; updated: %s)\n",
			task.ID, task.Description, task.Status,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}
