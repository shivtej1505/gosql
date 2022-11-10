package context

type InsertContext struct {
	Table  string
	Values []Value
}

type Value struct {
	Column string
	Value  interface{}
}

type CreateTableContext struct {
	Table   string
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

type SelectContext struct {
	Table     string
	Selectors []Selector
}

type Selector struct {
	Name string
}
