package api

import (
	"net/http"
)

type httpHandlerStartOlimp struct {
}

func NewHandlerStartOlimp() httpHandlerStartOlimp {
	return httpHandlerStartOlimp{}
}

func (h httpHandlerStartOlimp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/skor/FirstProject/static/startOlimp.html")
}
