package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shivtej1505/gosql/command"
)

func main() {
	fmt.Println("this is main")

	cleanupChan := make(chan os.Signal)
	signal.Notify(cleanupChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-cleanupChan
		fmt.Println("\nbye")
		os.Exit(0)
	}()

	startPrompt()
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
		parseQuery(strings.TrimSuffix(query, ";"))
	}
}

func parseQuery(query string) {
	tokens := strings.Split(strings.ToLower(query), " ")
	if len(tokens) == 0 {
		panic("invalid input")
	}

	if tokens[0] == "insert" {
		fmt.Println("parsing insert command")
	} else if tokens[0] == "select" {
		fmt.Println("parsing select command")
	} else if tokens[0] == "create" {
		if tokens[1] == "table" {
			fmt.Println("parsing command create table")
			fmt.Printf("Table=%s\n", tokens[2])
			ctCtx := command.CreateTableContext{
				Table: tokens[2],
			}
			if tokens[3] == "columns" {
				var column command.Column
				idx := 4
				for idx < len(tokens) {
					column.Name = tokens[idx]
					column.Type = strings.TrimSuffix(tokens[idx+1], ",")
					ctCtx.Columns = append(ctCtx.Columns, column)
					idx += 2
				}
				command.CreateTable(ctCtx)
			}
		}
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
