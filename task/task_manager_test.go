package task

import (
	"fmt"
	"testing"
	"time"
)

func TestTaskManagerList(t *testing.T) {
	var tasks []*Task
	store := &TaskStore{"../task.json", tasks}
	store.Read()
	manager := &TaskManager{store}
	manager.List()
}

func TestTaskManagerAdd(t *testing.T) {
	var tasks []*Task
	store := &TaskStore{"../task.json", tasks}
	store.Read()
	manager := &TaskManager{store}
	now := time.Now()
	var task = Task{store.tasks[len(store.tasks)-1].Id + 1, fmt.Sprintf("Test %d", store.tasks[len(store.tasks)-1].Id+1), 0, now, now}
	manager.Add(task)
}

func TestTaskManagerUpdate(t *testing.T) {
	var tasks []*Task
	store := &TaskStore{"../task.json", tasks}
	store.Read()
	manager := &TaskManager{store}
	manager.UpdateStatus(3, 1)
	manager.UpdateDescription(3, "New Test 3")
	manager.List()
}

func TestTaskManagerDelete(t *testing.T) {
	var tasks []*Task
	store := &TaskStore{"../task.json", tasks}
	store.Read()
	manager := &TaskManager{store}
	manager.Delete(3)
}

func TestTaskManagerClear(t *testing.T) {
	var tasks []*Task
	store := &TaskStore{"../task.json", tasks}
	store.Read()
	manager := &TaskManager{store}
	manager.Clear()
}
