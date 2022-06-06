package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mosteligible/todolist"
)

const todoFilename = ".todo.json"

func main() {
	// parse cli commands
	task := flag.String("task", "", "Task to be included in the ToDo list.")
	list := flag.Bool("list", false, "List the items in ToDo list.")
	complete := flag.Int("complete", 0, "Index of item to be marked complete.")

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
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag provided!")
		os.Exit(1)
	}
}
