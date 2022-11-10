package executor

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	context "github.com/shivtej1505/gosql/engine/context"
	"github.com/shivtej1505/gosql/utils"
)

type Executor struct {
}

func NewExecutor() Executor {
	return Executor{}
}

// Give table instance
func (exe Executor) Insert(insertCtx context.InsertContext) error {
	dir := "data"
	err := utils.CreateDir(dir)
	if err != nil {
		panic(err)
	}

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(dir+"/"+insertCtx.Table+".data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// TODO: Validation
	// 1. Check if table exits
	// 2. Check if query is correct

	// TODO:
	// 1. Add default values
	// 2. Put columns in order

	var writeStream []byte
	for _, val := range insertCtx.Values {
		var valueBytes []byte
		//if val.Value.(type) ==

		switch v := val.Value.(type) {
		case int:
			// Allocate 4 bytes and store it
			valueBytes = []byte(strconv.Itoa(v))
		case string:
			// Allocate 1 byte and store it
			valueBytes = []byte(v)
		default:
			errors.New("invalid value")
		}

		fmt.Printf("%v took %v bytes\n", val.Value, len(valueBytes))
		fmt.Println("------")

		writeStream = append(writeStream, valueBytes...)
	}

	_, err = f.Write(writeStream)
	//_, err = f.Write([]byte("1"))
	if err != nil {
		return err
	}

	return nil
}

func (exe Executor) CreateTable(createTableCtx context.CreateTableContext) error {
	dir := "data"
	err := utils.CreateDir(dir)
	if err != nil {
		panic(err)
	}

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(dir+"/"+createTableCtx.Table+".meta", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	for _, column := range createTableCtx.Columns {
		_, err = f.Write([]byte(column.Name + ";" + column.Type + "\n"))
		if err != nil {
			panic(err)
		}
	}
	return nil
}

// Give table instance
func (exe Executor) Select(selCtx context.SelectContext) error {
	log.Println(selCtx.Table)
	log.Println(selCtx.Selectors)
	log.Println("-----")

	dir := "data"

	if utils.FileExists(dir + "/" + selCtx.Table + ".meta") {
		meta, err := os.ReadFile(dir + "/" + selCtx.Table + ".meta")
		if err != nil {
			panic(meta)
		}
		fmt.Print(string(meta))

		if utils.FileExists(dir + "/" + selCtx.Table + ".data") {
			// TODO: Only read columns required
			data, err := os.ReadFile(dir + "/" + selCtx.Table + ".data")
			if err != nil {
				panic(data)
			}
			fmt.Print(string(data))
		}
	} else {
		fmt.Println("no such table")
	}

	return nil
}
