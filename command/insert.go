package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/shivtej1505/gosql/utils"
)

type InsertContext struct {
	Table  string
	Values []Value
}

type Value struct {
	Column string
	Value  interface{}
}

// Give table instance
func Insert(insertCtx InsertContext) error {
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
