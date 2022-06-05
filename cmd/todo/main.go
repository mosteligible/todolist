package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mosteligible/todolist"
)

var todoFilename = ".todo.json"

func main() {
	if os.Getenv("TODO_FILENAME") != "" {
		todoFilename = os.Getenv("TODO_FILENAME")
	}
	// parse cli commands
	task := flag.String("task", "", "Task to be included in the ToDo list.")
	list := flag.Bool("list", false, "List the items in ToDo list.")
	complete := flag.Int("complete", 0, "Index of item to be marked complete.")
	delete := flag.Int("del", 0, "Index of item to be deleted.")

	flag.Parse()

	l := &todolist.List{}

	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)

		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		l.Delete(*delete)

		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag provided!")
		os.Exit(1)
	}
}
