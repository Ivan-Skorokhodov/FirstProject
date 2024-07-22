package api

import (
	"net/http"
)

type httpHandlerStart struct {
}

func NewHandlerStart() httpHandlerStart {
	return httpHandlerStart{}
}

func (h httpHandlerStart) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/skor/GoStudy/FirstProject/static/start.html")
}
