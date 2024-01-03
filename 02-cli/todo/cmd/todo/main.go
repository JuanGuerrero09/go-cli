package main

import (
	"flag"
	"fmt"
	"go-cli/02-cli/todo"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2024\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	check := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	todoFileName := ".todo.json"

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments provided
	switch {
	// For no extra arguments, print the list
	case *list:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)

		}

	case *check > 0:
		if err := l.Complete(*check); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
