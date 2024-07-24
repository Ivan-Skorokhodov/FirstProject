package api

import (
	"FirstProject/pkg/tableOlimpiads"
	"FirstProject/pkg/tablePeople"
	"fmt"
	"net/http"
)

type httpHandlerShawAllPeople struct {
	peolpeDataTable *tablePeople.TablePeople
	olimpDataTable *tableOlimpiads.TableOlimpiads
}

func NewHandlerShawAllPeople(peolpeDataTable *tablePeople.TablePeople, olimpDataTable *tableOlimpiads.TableOlimpiads) httpHandlerShawAllPeople {
	return httpHandlerShawAllPeople{peolpeDataTable: peolpeDataTable, olimpDataTable: olimpDataTable}
}

func (h httpHandlerShawAllPeople) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for person, olimpNumber := range *h.peolpeDataTable.GetTable() {
		fmt.Fprintf(w, "Name: %s, age: %d, olimp: %s\n", person.GetName(), person.GetAge(), h.olimpDataTable.GetOlimp(olimpNumber).GetName())
	}
}
