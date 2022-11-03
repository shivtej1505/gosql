package command

import (
	"fmt"
	"os"
)

// Give table instance
func Select(table string) {
	dir := "data"

	data, err := os.ReadFile(dir + "/" + table)
	if err != nil {
		panic(data)
	}
	fmt.Print(string(data))
}
