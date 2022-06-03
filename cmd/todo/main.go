package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mosteligible/todolist"
)

const todoFilename = ".todo.json"

func main() {
	l := &todolist.List{}

	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
