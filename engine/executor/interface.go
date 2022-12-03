package executor

import context "github.com/shivtej1505/gosql/engine/context"

type Executor interface {
	Insert(insertCtx context.InsertContext) error
	Select(selCtx context.SelectContext) error
	CreateTable(createTableCtx context.CreateTableContext) error
}
