package service

import (
	"task-tracker-cli/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddTask(t *testing.T) {
	srv := New()
	id, err := srv.AddTask("Anything")
	require.NoError(t, err)

	task, err := srv.GetTask(id)
	require.NoError(t, err)

	assert.Equal(t, task.ID, id)
	assert.Equal(t, task.Description, "Anything")
	assert.Equal(t, task.Status, "todo")
}

func TestAddTask_EmptyDescription(t *testing.T) {
	srv := New()
	id, err := srv.AddTask("")
	require.Error(t, err, "description cannot be empty")
	assert.Equal(t, 0, id)
}

func TestUpdateTask(t *testing.T) {
	srv := New()
	id, _ := srv.AddTask("Anything")

	err := srv.UpdateTask(id, "Nothing")
	require.NoError(t, err)
	task, _ := srv.GetTask(id)

	assert.Equal(t, task.ID, id)
	assert.Equal(t, task.Description, "Nothing")
	assert.NotEqual(t, task.CreatedAt, task.UpdatedAt)
}

func TestDeleteTask(t *testing.T) {
	srv := New()
	id, _ := srv.AddTask("Anything")

	err := srv.DeleteTask(id)
	require.NoError(t, err)

}

func TestDeleteTask_NotExist(t *testing.T) {
	srv := New()
	err := srv.DeleteTask(1)
	require.Error(t, err)
}

func TestMarkDoneTask(t *testing.T) {
	srv := New()
	id, _ := srv.AddTask("Anything")

	err := srv.MarkDone(id)
	require.NoError(t, err)

	task, _ := srv.GetTask(id)
	assert.Equal(t, task.Status, "done")
	assert.NotEqual(t, task.CreatedAt, task.UpdatedAt)
}

func TestList(t *testing.T) {
	srv := New()
	id1, _ := srv.AddTask("Done task")
	id2, _ := srv.AddTask("In Progress task")
	srv.AddTask("Todo task")

	srv.MarkDone(id1)
	srv.MarkInProgress(id2)

	done := srv.ListTasks("done")
	assert.Len(t, done, 1)
	assert.Equal(t, done[0].Description, "Done task")

	todo := srv.ListTasks("todo")
	assert.Len(t, todo, 1)
	assert.Equal(t, todo[0].Description, "Todo task")

	inProgress := srv.ListTasks("in-progress")
	assert.Len(t, inProgress, 1)
	assert.Equal(t, inProgress[0].Description, "In Progress task")

	all := srv.ListTasks("")
	assert.Len(t, all, 3)
}

func TestLoad(t *testing.T) {
	srv := New()
	var tasks []model.Task
	time := time.Now()
	task := model.Task{
		ID:          1,
		Description: "Test",
		Status:      "todo",
		CreatedAt:   time,
		UpdatedAt:   time,
	}
	tasks = append(tasks, task)
	assert.Len(t, srv.tasks, 0)

	srv.Load(tasks)
	assert.Len(t, srv.tasks, 1)
}

func TestTasks(t *testing.T) {
	srv := New()
	var tasks []model.Task
	time := time.Now()
	task := model.Task{
		ID:          1,
		Description: "Test",
		Status:      "todo",
		CreatedAt:   time,
		UpdatedAt:   time,
	}
	tasks = append(tasks, task)
	srv.Load(tasks)
	assert.Len(t, srv.tasks, 1)

	res := srv.Tasks()
	assert.Equal(t, res, tasks)
}
