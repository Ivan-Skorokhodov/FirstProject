package api

import (
	"net/http"
)

type httpHandlerAddPerson struct {
}

func NewHandlerAddPerson() httpHandlerAddPerson {
	return httpHandlerAddPerson{}
}

func (h httpHandlerAddPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/skor/FirstProject/static/startPerson.html")
}
