package api

import (
	"net/http"
)

type httpHandlerAddOlimp struct {
}

func NewHandlerAddOlimp() httpHandlerAddOlimp {
	return httpHandlerAddOlimp{}
}

func (h httpHandlerAddOlimp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/skor/FirstProject/static/startOlimp.html")
}
