package main

type SQL map[string]Table

type Table [][]string

func Constructor(names []string, _ []int) SQL {
	tables := make(map[string]Table)
	for i := 0; i < len(names); i++ {
		tables[names[i]] = make(Table, 0, 16)
	}

	return tables
}

func (x SQL) InsertRow(name string, row []string) {
	x[name] = append(x[name], row)
}

func (x SQL) DeleteRow(_ string, _ int) {
}

func (x SQL) SelectCell(name string, rowId int, columnId int) string {
	return x[name][rowId-1][columnId-1]
}
