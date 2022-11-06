package command

import (
	"os"
)

type CreateTableContext struct {
	Table   string
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

func CreateTable(createTableCtx CreateTableContext) {
	dir := "data"
	err := createDir(dir)
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
}
