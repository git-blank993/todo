package task

import (
	"fmt"
	"testing"
	"time"
)

func TestTaskStoreWrite(t *testing.T) {
	now := time.Now()
	var tasks = []*Task{
		{
			Id:          1,
			Description: "Buy milk",
			Status:      InProgress,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Id:          2,
			Description: "Dont forget to eat eggs",
			Status:      Todo,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	manager := &TaskStore{"../task.json", tasks}
	manager.Save()
}

func TestTaskStoreRead(t *testing.T) {
	var tasks []*Task
	var manager = &TaskStore{"../task.json", tasks}
	manager.Read()
	for _, task := range manager.tasks {
		fmt.Println(task)
	}
}
