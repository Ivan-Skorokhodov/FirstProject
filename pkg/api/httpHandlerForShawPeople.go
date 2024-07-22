package api

import (
	"FirstProject/pkg/tableOlimpiads"
	"FirstProject/pkg/tablePeople"
	"fmt"
	"net/http"
)

type httpHandlerForShawPeople struct {
	peolpeDataTable *tablePeople.TablePeople
	olimpDataTable *tableOlimpiads.TableOlimpiads
}

func NewHandlerForShawPeople(peolpeDataTable *tablePeople.TablePeople, olimpDataTable *tableOlimpiads.TableOlimpiads) httpHandlerForShawPeople {
	return httpHandlerForShawPeople{peolpeDataTable: peolpeDataTable, olimpDataTable: olimpDataTable}
}

func (h httpHandlerForShawPeople) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for person, olimpNumber := range *h.peolpeDataTable.GetTable() {
		fmt.Fprintf(w, "Name: %s, age: %d, olimp: %s\n", person.GetName(), person.GetAge(), h.olimpDataTable.GetOlimp(olimpNumber).GetName())
	}
}
