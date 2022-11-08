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
	"github.com/shivtej1505/gosql/utils"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

type Config struct {
	Directory string `json:"directory"`
}

type Engine struct {
	Config Config
}

func main() {
	cleanupChan := make(chan os.Signal)
	signal.Notify(cleanupChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-cleanupChan
		fmt.Println("\nbye")
		os.Exit(0)
	}()

	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelTrace)
	engine := initEngine()
	startPrompt(engine)
}

func initEngine() Engine {
	viper.Set("Debug", true)
	viper.Set("Info", true)
	viper.Set("Verbose", true)
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	var config Config

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	fmt.Println(viper.GetString("directory"))

	return Engine{
		Config: config,
	}
}

// TODO: Implement history
func startPrompt(engine Engine) {
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
		query = strings.TrimPrefix(query, "\n")
		query = strings.TrimSuffix(query, ";")
		parseQuery(query)
	}
}

func parseQuery(query string) {
	tokens := strings.Split(strings.ToLower(query), " ")
	if len(tokens) == 0 {
		panic("invalid input")
	}

	if tokens[0] == "insert" { // insert into table (id, name) values (1, "Shivang")
		fmt.Println("parsing insert command")
		insertCtx := command.InsertContext{}
		if tokens[1] == "into" {
			insertCtx.Table = tokens[2]

			idx := 3
			valueCount := 0
			for idx < len(tokens) {
				if tokens[idx] == "values" {
					break
				}
				idx++
				valueCount++
			}

			var values []command.Value
			var value command.Value
			idx = 3
			for idx < 3+valueCount {
				value.Column = utils.RemoveInvalidChars(tokens[idx])
				value.Value = utils.RemoveInvalidChars(tokens[idx+valueCount+1])

				values = append(values, value)
				idx++
			}

			insertCtx.Values = values

			command.Insert(insertCtx)
		}
	} else if tokens[0] == "select" { // select * from table
		fmt.Println("parsing select command")
		selCtx := command.SelectContext{}
		var selector command.Selector

		idx := 1
		for idx < len(tokens) {
			if tokens[idx] == "from" {
				break
			} else {
				selector.Name = tokens[idx]
			}
			idx += 1
			selCtx.Selectors = append(selCtx.Selectors, selector)
		}

		selCtx.Table = tokens[idx+1]
		command.Select(selCtx)
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
