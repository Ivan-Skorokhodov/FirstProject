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

	h1 := NewHandlerStartOlimp()
	h2 := NewHandlerStartPerson()
	h3 := NewHandlerForAddPerson(&peopleDataTable, &olimpDataTable)
	h4 := NewHandlerForShawPeople(&peopleDataTable, &olimpDataTable)
	h5 := NewHandlerForAddOlimp(&olimpDataTable)

	http.Handle("/startOlimp", h1)
	http.Handle("/startPerson", h2)
	http.Handle("/addPerson", h3)
	http.Handle("/addOlimp", h5)
	http.Handle("/shawAll", h4)

}
