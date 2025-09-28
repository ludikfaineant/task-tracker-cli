package service

import (
	"fmt"
	"task-tracker-cli/internal/model"
	"time"
)

// TaskService manages a list of tasks
type TaskService struct {
	tasks []model.Task
}

// New returns a new TaskService with an empty task list
func New() *TaskService {
	return &TaskService{tasks: []model.Task{}}
}

// Tasks returns a copy of all tasks in the service
func (s *TaskService) Tasks() []model.Task {
	return s.tasks
}

// Load replaces the current list of tasks with the given slice
func (s *TaskService) Load(tasks []model.Task) {
	s.tasks = tasks
}

// AddTask adds a new task with the given description and returns its ID and error
func (s *TaskService) AddTask(description string) (int, error) {
	if description == "" {
		return 0, fmt.Errorf("description cannot be empty")
	}
	id := getNextID(s.tasks)
	newTask := model.Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	s.tasks = append(s.tasks, newTask)
	return id, nil
}

// UpdateTask update the desciption of a task by ID and returns error if not found
func (s *TaskService) UpdateTask(id int, description string) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks[idx].Description = description
	s.tasks[idx].UpdatedAt = time.Now().UTC()
	return nil
}

// DeleteTask removes a task by ID and returns error if not found
func (s *TaskService) DeleteTask(id int) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks = append(s.tasks[:idx], s.tasks[idx+1:]...)
	return nil
}

// ListTasks returns all tasks, optionally filtered by stasus {"done", "todo", "in-progress"}
func (s *TaskService) ListTasks(by string) []model.Task {
	var result []model.Task
	if by == "todo" || by == "done" || by == "in-progress" {
		for _, task := range s.tasks {
			if task.Status == by {
				result = append(result, task)
			}
		}
	} else {
		result = s.tasks
	}
	return result
}

// MarkDone marks the task with the given ID as done and returns error if not found
func (s *TaskService) MarkDone(id int) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks[idx].Status = "done"
	s.tasks[idx].UpdatedAt = time.Now().UTC()
	return nil
}

// MarkInProgress marks the task with the given ID as in-progress and returns error if not found
func (s *TaskService) MarkInProgress(id int) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks[idx].Status = "in-progress"
	s.tasks[idx].UpdatedAt = time.Now().UTC()
	return nil
}

// GetTask returns a pointer to the task with the given ID and error if not found
func (s *TaskService) GetTask(id int) (*model.Task, error) {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return nil, err
	}
	return &s.tasks[idx], nil
}

func getNextID(tasks []model.Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

func findID(tasks []model.Task, id int) (int, error) {
	for i, task := range tasks {
		if task.ID == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Not found task with id: %d", id)
}
