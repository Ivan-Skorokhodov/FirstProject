package tablePeople

import (
	"FirstProject/pkg/person"
)

type TablePeople struct {
	table map[person.Person]int
}

func NewTable(table map[person.Person]int) TablePeople {
	return TablePeople{table: table}
}

func (t *TablePeople) Add(person person.Person, olimpNumber int) {
	t.table[person] = olimpNumber
}

func (t *TablePeople) GetTable() *map[person.Person]int {
	return &t.table
}
