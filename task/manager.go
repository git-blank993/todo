package task

import (
	"fmt"
	"log"

	"github.com/rodaine/table"
)

func NewTaskManager(store *TaskStore) *TaskManager {
	return &TaskManager{store}
}

type TaskManager struct {
	store *TaskStore
}

func (t *TaskManager) List() {
	tbl := table.New("ID", "Description", "Status", "Created At", "Updated At")
	for _, task := range t.store.tasks {
		tbl.AddRow(task.Id, task.Description, task.Status.StatusName(), task.CreatedAt, task.UpdatedAt)
	}
	tbl.Print()
	if len(t.store.tasks) == 0 {

		fmt.Println("No tasks available")
	}
}

func (t *TaskManager) FilterList(status TaskStatus) {
	tbl := table.New("ID", "Description", "Status", "Created At", "Updated At")
	for _, task := range t.store.tasks {
		if task.Status == status {
			tbl.AddRow(task.Id, task.Description, task.Status.StatusName(), task.CreatedAt, task.UpdatedAt)
		}
	}
	tbl.Print()
	if len(t.store.tasks) == 0 {

		fmt.Println("No tasks available")
	}
}

func (t *TaskManager) Add(task Task) {
	t.store.tasks = append(t.store.tasks, &task)
	if ok, err := t.store.Save(); err != nil {
		log.Fatal("Error while creating a task: ", err.Error())
	} else {
		if ok {
			fmt.Println("Task added successfully")
		}
	}

}

func (t *TaskManager) UpdateStatus(id int, status TaskStatus) {
	if len(t.store.tasks) == 0 {

		fmt.Println("No tasks available")
		return
	}
	for _, task := range t.store.tasks {
		if task.Id == id {
			task.ChangeStatus(status)
		}
	}
	if ok, err := t.store.Save(); err != nil {
		log.Fatal("Error while creating a task: ", err.Error())
	} else {
		if ok {
			fmt.Println("Task status updated successfully")
		}
	}
}

func (t *TaskManager) UpdateDescription(id int, desc string) {
	if len(t.store.tasks) == 0 {

		fmt.Println("No tasks available")
		return
	}
	var found bool
	for _, task := range t.store.tasks {
		if task.Id == id {
			task.ChangeDescription(desc)
			found = true
		}
	}
	if !found {
		fmt.Println("No task found")
		return
	}
	if ok, err := t.store.Save(); err != nil {
		log.Fatal("Error while creating a task: ", err.Error())
	} else {
		if ok {
			fmt.Println("Task description updated successfully")
		}
	}
}

func (t *TaskManager) Delete(id int) {
	if len(t.store.tasks) == 0 {

		fmt.Println("No tasks available")
		return
	}
	var found int = -1
	for i := range t.store.tasks {
		if t.store.tasks[i].Id == id {
			found = i
		}
	}
	if found == -1 {
		fmt.Println("No task found")
		return
	}
	t.store.tasks = append(t.store.tasks[:found], t.store.tasks[found+1:]...)
	if ok, err := t.store.Save(); err != nil {
		log.Fatal("Error while creating a task: ", err.Error())
	} else {
		if ok {
			fmt.Println("Task deleted successfully")
		}
	}
}

func (t *TaskManager) Clear() {
	t.store.tasks = []*Task{}
	if ok, err := t.store.Save(); err != nil {
		log.Fatal("Error while creating a task: ", err.Error())
	} else {
		if ok {
			fmt.Println("All tasks have been cleared")
		}
	}
}

func (t *TaskManager) NewId() int {
	if len(t.store.tasks) == 0 {
		return 1
	}
	return t.store.tasks[len(t.store.tasks)-1].Id + 1
}
