package esgo

import (
	"time"
)

type TaskHandler interface {
	HandleTask(t *Task) error
}

type TaskListener interface {
	ListenTasks() chan *TaskWork
	Queue(t *Task) error
}

type Task struct {
	Type    string                 `json:"task"`
	EventID string                 `json:"eventid"`
	Created time.Time              `json"created"`
	Data    map[string]interface{} `json:"data"`
}

func NewTask() *Task {
	var t Task
	t.Created = time.Now().UTC()
	t.Data = make(map[string]interface{})
	return &t
}

type TaskWork struct {
	Result chan error
	Task
}

func NewTaskWork(t Task) *TaskWork {
	var tw TaskWork
	tw.Task = t
	tw.Result = make(chan error, 1)
	return &tw
}
