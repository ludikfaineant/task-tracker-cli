package service

import (
	"fmt"
	"task-tracker-cli/internal/model"
	"time"
)

type TaskService struct {
	tasks []model.Task
}

func New() *TaskService {
	return &TaskService{tasks: []model.Task{}}
}

func (s *TaskService) Tasks() []model.Task {
	return s.tasks
}

func (s *TaskService) Load(tasks []model.Task) {
	s.tasks = tasks
}

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

func (s *TaskService) UpdateTask(id int, description string) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks[idx].Description = description
	s.tasks[idx].UpdatedAt = time.Now().UTC()
	return nil
}

func (s *TaskService) DeleteTask(id int) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks = append(s.tasks[:idx], s.tasks[idx+1:]...)
	return nil
}

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

func (s *TaskService) MarkDone(id int) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks[idx].Status = "done"
	s.tasks[idx].UpdatedAt = time.Now().UTC()
	return nil
}

func (s *TaskService) MarkInProgress(id int) error {
	idx, err := findID(s.tasks, id)
	if err != nil {
		return err
	}
	s.tasks[idx].Status = "in-progress"
	s.tasks[idx].UpdatedAt = time.Now().UTC()
	return nil
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
