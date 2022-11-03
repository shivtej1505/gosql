package main

import (
	"fmt"

	"github.com/shivtej1505/gosql/command"
)

func main() {
	fmt.Println("this is main")

	// insert 5 values
	// read them back
	executeCommands()
}

func executeCommands() {
	table := "numbers"
	command.Insert(table, "1")
	command.Insert(table, "2")
	command.Insert(table, "3")
	command.Insert(table, "4")
	command.Insert(table, "5")

	command.Select(table)
}
