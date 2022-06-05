package todolist_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mosteligible/todolist"
)

func TestAdd(t *testing.T) {
	taskList := todolist.List{}

	newTask := "New Task"
	taskList.Add(newTask)

	if taskList[0].Task != newTask {
		t.Errorf("Expected task: %s, but got %s", newTask, taskList[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todolist.List{}

	taskName := "Task 01"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be complete.")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be complete.")
	}
}

func TestDelete(t *testing.T) {
	l := todolist.List{}

	tasks := []string{
		"task 01",
		"task 02",
		"task 03",
	}

	for _, tsk := range tasks {
		l.Add(tsk)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, but got %q", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, but got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todolist.List{}
	l2 := todolist.List{}

	taskName := "New Task1"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating a temporary file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	if err := l1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tempFile.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}
