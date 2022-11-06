package command

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type SelectContext struct {
	Table     string
	Selectors []Selector
}

type Selector struct {
	Name string
}

// Give table instance
func Select(selCtx SelectContext) {
	log.Println(selCtx.Table)
	log.Println(selCtx.Selectors)
	log.Println("-----")

	dir := "data"

	if FileExists(dir + "/" + selCtx.Table + ".meta") {
		meta, err := os.ReadFile(dir + "/" + selCtx.Table + ".meta")
		if err != nil {
			panic(meta)
		}
		fmt.Print(string(meta))

		if FileExists(dir + "/" + selCtx.Table + ".data") {
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
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
