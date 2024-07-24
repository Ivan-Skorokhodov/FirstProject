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

	h1 := NewHandlerAddOlimp()
	h2 := NewHandlerAddPerson()
	h3 := NewHandlerShawAddedPerson(&peopleDataTable, &olimpDataTable)
	h4 := NewHandlerShawAddedOlimp(&olimpDataTable)
	h5 := NewHandlerShawAllPeople(&peopleDataTable, &olimpDataTable)

	http.Handle("/addOlimp", h1)
	http.Handle("/addPerson", h2)
	http.Handle("/shawPerson", h3)
	http.Handle("/shawOlimp", h4)
	http.Handle("/shawAll", h5)

}
