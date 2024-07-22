package api

import (
	"FirstProject/pkg/olimpiad"
	"FirstProject/pkg/person"
	"FirstProject/pkg/tableOlimpiads"
	"FirstProject/pkg/tablePeople"
	"net/http"
)

func FillEndpoints() {

	peopleDataTable := tablePeople.NewTable(make(map[person.Person]int))
	olimpDataTable := tableOlimpiads.NewTable(make(map[int]olimpiad.Olimp))

	h1 := NewHandlerStart()
	h2 := NewHandlerForAddPerson(&peopleDataTable, &olimpDataTable)
	h3 := NewHandlerForShawPeople(&peopleDataTable, &olimpDataTable)
	h4 := NewHandlerForAddOlimp(&olimpDataTable)

	http.Handle("/", h1)
	http.Handle("/addPerson", h2)
	http.Handle("/addOlimp", h4)
	http.Handle("/shawAll", h3)

}
