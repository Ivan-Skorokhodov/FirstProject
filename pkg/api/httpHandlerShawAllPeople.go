package api

import (
	"FirstProject/pkg/tableOlimpiads"
	"FirstProject/pkg/tablePeople"
	"html/template"
	"net/http"
)

type httpHandlerShawAllPeople struct {
	peolpeDataTable *tablePeople.TablePeople
	olimpDataTable *tableOlimpiads.TableOlimpiads
}

func NewHandlerShawAllPeople(peolpeDataTable *tablePeople.TablePeople, olimpDataTable *tableOlimpiads.TableOlimpiads) httpHandlerShawAllPeople {
	return httpHandlerShawAllPeople{peolpeDataTable: peolpeDataTable, olimpDataTable: olimpDataTable}
}

type ViewData struct{
    Users []User
}
type User struct{
    Name string
	Olimp string
    Age int
}

func (h httpHandlerShawAllPeople) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data ViewData

	for person, olimpNumber := range *h.peolpeDataTable.GetTable() {
		new_element := User{Name: person.GetName(), Age: person.GetAge(), Olimp: h.olimpDataTable.GetOlimp(olimpNumber).GetName()}
		data.Users = append(data.Users, new_element)
	}

	tmpl, _ := template.ParseFiles("/home/skor/FirstProject/templates/shawAll.html")
    tmpl.Execute(w, data)
}
