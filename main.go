package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/shivtej1505/gosql/command"
)

func main() {
	fmt.Println("this is main")

	startPrompt()

	// insert 5 values
	// read them back
	//executeCommands()
}

func startPrompt() {
	//scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	var query string
	var delim byte
	var err error
	delim = ';'
	for {
		fmt.Printf("gosql>")
		query, err = reader.ReadString(delim)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", query)
	}
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
