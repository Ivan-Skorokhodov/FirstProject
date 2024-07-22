package tableOlimpiads

import "FirstProject/pkg/olimpiad"

type TableOlimpiads struct {
	table map[int]olimpiad.Olimp
}

func NewTable(table map[int]olimpiad.Olimp) TableOlimpiads {
	return TableOlimpiads{table: table}
}

func (t *TableOlimpiads) Add(olimp olimpiad.Olimp) {
	t.table[olimp.GetOlimpNumber()] = olimp
}

func (t *TableOlimpiads) GetTable() *map[int]olimpiad.Olimp {
	return &t.table
}

func (t *TableOlimpiads) GetOlimp(olimpNumber int) olimpiad.Olimp {
	return t.table[olimpNumber]
}