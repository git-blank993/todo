package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DataStore interface {
	Save()
	Read() []*Task
}

func NewTaskStore(filename string) *TaskStore {
	tasks := []*Task{}
	return &TaskStore{filename, tasks}
}

type TaskStore struct {
	filename string
	tasks    []*Task
}

func (t *TaskStore) Read() {
	data, err := os.ReadFile(t.filename)
	if err != nil {
		log.Fatal(err)
	}
	var tasks []Task
	if err = json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("No data present in json file")
		log.Fatal(err)
	}
	t.tasks = []*Task{}
	for i := range tasks {
		t.tasks = append(t.tasks, &tasks[i])
	}
}

func (t *TaskStore) Save() (bool, error) {
	f, err := os.OpenFile(t.filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()
	data, err := json.Marshal(t.tasks)
	if err != nil {
		return false, err
	}
	f.Truncate(0)
	f.Seek(0, 0)
	_, err = f.Write(data)
	if err != nil {
		return false, (err)
	}
	return true, nil
}
