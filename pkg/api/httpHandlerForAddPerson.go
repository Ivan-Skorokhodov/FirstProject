package api

import (
	"FirstProject/pkg/person"
	"FirstProject/pkg/tableOlimpiads"
	"FirstProject/pkg/tablePeople"
	"fmt"
	"net/http"
	"strconv"
)

type httpHandlerForAddPerson struct {
	peolpeDataTable *tablePeople.TablePeople
	olimpDataTable *tableOlimpiads.TableOlimpiads
}

func NewHandlerForAddPerson(peolpeDataTable *tablePeople.TablePeople, olimpDataTable *tableOlimpiads.TableOlimpiads) httpHandlerForAddPerson {
	return httpHandlerForAddPerson{peolpeDataTable: peolpeDataTable, olimpDataTable: olimpDataTable}
}

func (h httpHandlerForAddPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	olimpNumber := r.FormValue("olimpNumber")

	ageInt, _ := strconv.Atoi(age)
	olimpNumberInt, _ := strconv.Atoi(olimpNumber)

	person := person.NewPerson(name, ageInt)
	h.peolpeDataTable.Add(person, olimpNumberInt)

	fmt.Fprintf(w, "We add person with name: %s, age: %s, olimpiad: %s", name, age, h.olimpDataTable.GetOlimp(olimpNumberInt).GetName())
}
