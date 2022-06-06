package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")

	result := m.Run()

	fmt.Println("Cleaning up....")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "task number 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("%s", err)
	}

	cmdPath := filepath.Join(dir, binName)

	// test add new task
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)

		if err := cmd.Run(); err != nil {
			t.Fatalf("%s", err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%s", err)
		}

		expected := "   1: " + task + "\n"

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		// create task for completion
		tasks := []string{"task 01", "task 02"}
		for _, item := range tasks {
			cmdTaskAdd := exec.Command(cmdPath, "-task", item)
			if err := cmdTaskAdd.Run(); err != nil {
				t.Fatalf("%s", err)
			}
		}

		cmdComplete := exec.Command(cmdPath, "-complete", "2")
		if err := cmdComplete.Run(); err != nil {
			t.Fatalf("%s", err)
		}

		listCmd := exec.Command(cmdPath, "-list")
		todoItemsOut, err := listCmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%s", err)
		}

		// Note: the task added during `AddTask` is first one in the list of tasks, so first listed tasks will be it
		todoList := string(todoItemsOut)
		todoItems := strings.Split(todoList, "\n")
		for _, item := range todoItems {
			if item == tasks[0] {
				t.Fatalf("The task \"%s\" should not be complete", tasks[0])
			}
		}
	})
}
