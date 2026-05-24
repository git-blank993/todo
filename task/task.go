package task

import (
	"encoding/json"
	"fmt"
	"time"
)

type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

var stateName = map[TaskStatus]string{
	Todo:       "todo",
	InProgress: "in-progress",
	Done:       "done",
}

func (t TaskStatus) StatusName() string {
	return stateName[t]
}

func (t *TaskStatus) MarshalJSON() ([]byte, error) {
	s := t.StatusName()
	return json.Marshal(s)
}

func (t *TaskStatus) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s) // decode JSON string first
	if err != nil {
		return err
	}

	switch s {
	case "todo":
		*t = Todo
	case "in-progress":
		*t = InProgress
	case "done":
		*t = Done
	default:
		return fmt.Errorf("unknown status: %s", s)
	}

	return nil
}

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (t *Task) ChangeStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

func (t *Task) ChangeDescription(desc string) {
	t.Description = desc
	t.UpdatedAt = time.Now()
}

func NewTask(id int, description string) Task {
	now := time.Now()
	task := Task{id, description, Todo, now, now}
	return task
}
