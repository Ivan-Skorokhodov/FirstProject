package api

import (
	"net/http"
)

type httpHandlerStartPerson struct {
}

func NewHandlerStartPerson() httpHandlerStartPerson {
	return httpHandlerStartPerson{}
}

func (h httpHandlerStartPerson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/skor/FirstProject/static/startPerson.html")
}
