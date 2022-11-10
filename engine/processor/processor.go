package processor

import (
	"fmt"
	"strings"

	context "github.com/shivtej1505/gosql/engine/context"
	"github.com/shivtej1505/gosql/utils"
)

type Processor struct {
}

func (processor Processor) ParseQuery(query string) interface{} {
	tokens := strings.Split(strings.ToLower(query), " ")
	if len(tokens) == 0 {
		panic("invalid input")
	}

	if tokens[0] == "insert" { // insert into table (id, name) values (1, "Shivang")
		fmt.Println("parsing insert command")
		insertCtx := context.InsertContext{}
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

			var values []context.Value
			var value context.Value
			idx = 3
			for idx < 3+valueCount {
				value.Column = utils.RemoveInvalidChars(tokens[idx])
				value.Value = utils.RemoveInvalidChars(tokens[idx+valueCount+1])

				values = append(values, value)
				idx++
			}

			insertCtx.Values = values

			//command.Insert(insertCtx)
			return insertCtx
		}
	} else if tokens[0] == "select" { // select * from table
		fmt.Println("parsing select command")
		selCtx := context.SelectContext{}
		var selector context.Selector

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
		return selCtx
		//command.Select(selCtx)
	} else if tokens[0] == "create" { // create table tinku columns id int, name string;
		if tokens[1] == "table" {
			fmt.Println("parsing command create table")
			fmt.Printf("Table=%s\n", tokens[2])
			ctCtx := context.CreateTableContext{
				Table: tokens[2],
			}
			if tokens[3] == "columns" {
				var column context.Column
				idx := 4
				for idx < len(tokens) {
					column.Name = tokens[idx]
					column.Type = strings.TrimSuffix(tokens[idx+1], ",")
					ctCtx.Columns = append(ctCtx.Columns, column)
					idx += 2
				}
				//command.CreateTable(ctCtx)
				return ctCtx
			}
		}
	}
	return nil
}
